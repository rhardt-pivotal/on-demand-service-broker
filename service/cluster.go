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
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/craigfurman/herottp"
	"github.com/pivotal-cf/on-demand-service-broker/authorizationheader"
	"github.com/pivotal-cf/on-demand-service-broker/config"
)

type Cluster struct {
	GUID         string `json:"service_instance_id"`
	PlanUniqueID string `json:"plan_id"`
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/fake_instance_lister.go . InstanceLister
// InstanceLister provides a interface to query service instances present in the platform
type IClusterLister interface {
	Clusters(string) (string, error)
}


func BuildClusterLister(
	siapiConfig config.ServiceInstancesAPI, logger *log.Logger,
) (IClusterLister, error) {

	url := siapiConfig.URL
	url = url[0:strings.LastIndex(url, "/")]+"/v1/clusters"

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("error getting a certificate pool to append our trusted cert to: %s", err)
	}
	cert := siapiConfig.RootCACert
	certPool.AppendCertsFromPEM([]byte(cert))

	httpClient := herottp.New(herottp.Config{
		Timeout:                           30 * time.Second,
		RootCAs:                           certPool,
		DisableTLSCertificateVerification: siapiConfig.DisableSSLCertVerification,
	})

	authHeaderBuilder := authorizationheader.NewBasicAuthHeaderBuilder(
		siapiConfig.Authentication.Basic.Username,
		siapiConfig.Authentication.Basic.Password,
	)

	return NewClusterLister(httpClient, authHeaderBuilder, url, true, logger), nil
}
