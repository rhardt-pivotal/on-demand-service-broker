// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/pivotal-cf/on-demand-service-broker/authorizationheader"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/fake_doer.go . Doer

type ClusterLister struct {
	authHeaderBuilder authorizationheader.AuthHeaderBuilder
	baseURL           string
	client            Doer
	configured        bool
	logger            *log.Logger
}

var (
	ClusterNotFound = errors.New("Cluster not found")
)

func NewClusterLister(
	client Doer,
	authHeaderBuilder authorizationheader.AuthHeaderBuilder,
	baseURL string,
	configured bool,
	logger *log.Logger) *ClusterLister {
	return &ClusterLister{
		authHeaderBuilder: authHeaderBuilder,
		baseURL:           baseURL,
		client:            client,
		configured:        configured,
		logger:            logger,
	}
}

func (s *ClusterLister) Clusters(name string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, s.baseURL, nil)
	if err != nil {
		return "", err
	}

	if len(name) > 0 {
		request.URL.Path = request.URL.Path + "/" + name
	}

	err = s.authHeaderBuilder.AddAuthHeader(request, s.logger)
	if err != nil {
		return "", err
	}

	response, err := s.client.Do(request)

	if err != nil {
		return s.clusterListerError(response, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		genericJson := map[string]string{}
		err := json.NewDecoder(response.Body).Decode(&genericJson)
		body := ""
		if err == nil {
			body = genericJson["description"]
		}
		return s.clusterListerError(response, fmt.Errorf("HTTP response status: %s. %s", response.Status, body))
	}

	// var instances []Instance
	//err = json.NewDecoder(response.Body).Decode(&instances)
	//if err != nil {
	//	return instances, err
	//}
	//return instances, nil

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}


func (s *ClusterLister) clusterListerError(response *http.Response, err error) (string, error) {
	if s.configured {
		if urlError, ok := err.(*url.Error); ok {
			if urlError.Err != nil && urlError.URL != "" {
				switch urlError.Err.(type) {
				case x509.UnknownAuthorityError:
					return "", fmt.Errorf(
						"SSL validation error for `service_instances_api.url`: %s. Please configure a `service_instances_api.root_ca_cert` and use a valid SSL certificate",
						urlError.URL,
					)
				default:
					return "", fmt.Errorf("error communicating with service_instances_api (%s): %s", urlError.URL, err.Error())
				}
			}
		}

		if response != nil {
			return "", fmt.Errorf("error communicating with service_instances_api (%s): %s", response.Request.URL, err.Error())
		}
	}
	return "", err
}
