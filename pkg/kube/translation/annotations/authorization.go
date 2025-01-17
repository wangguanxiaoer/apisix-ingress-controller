// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package annotations

import (
	apisixv1 "github.com/apache/apisix-ingress-controller/pkg/types/apisix/v1"
)

const (
	// auth-type: keyAuth | basicAuth
	_authType = AnnotationsPrefix + "auth-type"
)

type basicAuth struct{}

// NewkeyBasicHandler creates a handler to convert
// annotations about basicAuth control to APISIX basic-auth plugin.
func NewBasicAuthHandler() Handler {
	return &basicAuth{}
}

func (b *basicAuth) PluginName() string {
	return "basic-auth"
}

func (b *basicAuth) Handle(e Extractor) (interface{}, error) {
	if e.GetStringAnnotation(_authType) != "basicAuth" {
		return nil, nil
	}
	plugin := apisixv1.BasicAuthConfig{}
	return &plugin, nil
}

type keyAuth struct{}

// NewkeyAuthHandler creates a handler to convert
// annotations about keyAuth control to APISIX key-auth plugin.
func NewKeyAuthHandler() Handler {
	return &keyAuth{}
}

func (k *keyAuth) PluginName() string {
	return "key-auth"
}

func (k *keyAuth) Handle(e Extractor) (interface{}, error) {
	if e.GetStringAnnotation(_authType) != "keyAuth" {
		return nil, nil
	}
	plugin := apisixv1.KeyAuthConfig{}
	return &plugin, nil
}
