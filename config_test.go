/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/HotelsDotCom/go-logger/loggertest"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var TestEnv map[string]string

func BeforeConfig() {

	loggertest.Init("DEBUG")
	TestEnv = make(map[string]string)
	lookupEnv = func(k string) (string, bool) {
		v, ok := TestEnv[k]
		return v, ok
	}
}

func AfterConfig() {

	lookupEnv = os.LookupEnv
	loggertest.Reset()
}

func TestApiHostEnv(t *testing.T) {

	BeforeConfig()
	defer AfterConfig()

	TestEnv["FLYTE_API"] = "http://test_api:8080"

	url := apiHost()
	assert.Equal(t, "http://test_api:8080", url.String())
}

func TestPackName(t *testing.T) {
	BeforeConfig()
	defer AfterConfig()

	TestEnv["PACK_NAME"] = "Slack2"
	assert.Equal(t, "Slack2", packName())
}

func TestApiHostEnvNotSet(t *testing.T) {

	BeforeConfig()
	defer AfterConfig()

	assert.Panics(t, func() { apiHost() })
}

func TestApiHostEnvInvalidUrl(t *testing.T) {

	BeforeConfig()
	defer AfterConfig()

	TestEnv["FLYTE_API"] = ":/invalid url"

	loggertest.Init("DEBUG")
	defer loggertest.Reset()

	assert.Panics(t, func() { apiHost() })
}

func TestSlackTokenEnv(t *testing.T) {

	BeforeConfig()
	defer AfterConfig()

	TestEnv["FLYTE_SLACK_TOKEN"] = "abc"

	assert.Equal(t, "abc", slackToken())
}

func TestSlackTokenEnvNotSet(t *testing.T) {

	BeforeConfig()
	defer AfterConfig()

	assert.Panics(t, func() { slackToken() })
}
