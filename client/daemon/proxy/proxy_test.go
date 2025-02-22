/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"d7y.io/dragonfly/v2/client/config"
)

type testItem struct {
	URL      string
	Direct   bool
	UseHTTPS bool
	Redirect string
}

type testCase struct {
	Error          error
	Rules          []*config.Proxy
	RegistryMirror *config.RegistryMirror
	Items          []testItem
}

func newTestCase() *testCase {
	return &testCase{}
}

func (tc *testCase) WithRule(regx string, direct bool, useHTTPS bool, redirect string) *testCase {
	if tc.Error != nil {
		return tc
	}

	var r *config.Proxy
	r, tc.Error = config.NewProxy(regx, useHTTPS, direct, redirect)
	tc.Rules = append(tc.Rules, r)
	return tc
}

func (tc *testCase) WithRegistryMirror(rawURL string, direct bool, dynamic bool) *testCase {
	if tc.Error != nil {
		return tc
	}

	var u *url.URL
	u, tc.Error = url.Parse(rawURL)
	tc.RegistryMirror = &config.RegistryMirror{
		Remote:        &config.URL{URL: u},
		DynamicRemote: dynamic,
		Direct:        direct,
	}
	return tc
}

func (tc *testCase) WithTest(url string, direct bool, useHTTPS bool, redirect string) *testCase {
	tc.Items = append(tc.Items, testItem{url, direct, useHTTPS, redirect})
	return tc
}

func (tc *testCase) Test(t *testing.T) {
	a := assert.New(t)
	if !a.Nil(tc.Error) {
		return
	}
	tp, err := NewProxy(WithRules(tc.Rules))
	if !a.Nil(err) {
		return
	}
	for _, item := range tc.Items {
		req, err := http.NewRequest("GET", item.URL, nil)
		if !a.Nil(err) {
			continue
		}
		if !a.Equal(tp.shouldUseDragonfly(req), !item.Direct) {
			fmt.Println(item.URL)
		}
		if item.UseHTTPS {
			a.Equal(req.URL.Scheme, "https")
		} else {
			a.Equal(req.URL.Scheme, "http")
		}
		if item.Redirect != "" {
			a.Equal(item.Redirect, req.Host)
		}
	}
}

func (tc *testCase) TestMirror(t *testing.T) {
	a := assert.New(t)
	if !a.Nil(tc.Error) {
		return
	}
	tp, err := NewProxy(WithRegistryMirror(tc.RegistryMirror))
	if !a.Nil(err) {
		return
	}
	for _, item := range tc.Items {
		req, err := http.NewRequest("GET", item.URL, nil)
		if !a.Nil(err) {
			continue
		}
		if !a.Equal(tp.shouldUseDragonflyForMirror(req), !item.Direct) {
			fmt.Println(item.URL)
		}
	}
}

func TestMatch(t *testing.T) {
	newTestCase().
		WithRule("/blobs/sha256/", false, false, "").
		WithTest("http://index.docker.io/v2/blobs/sha256/xxx", false, false, "").
		WithTest("http://index.docker.io/v2/auth", true, false, "").
		Test(t)

	newTestCase().
		WithRule("/a/d", false, true, "").
		WithRule("/a/b", true, false, "").
		WithRule("/a", false, false, "").
		WithRule("/a/c", true, false, "").
		WithRule("/a/e", false, true, "").
		WithTest("http://h/a", false, false, "").   // should match /a
		WithTest("http://h/a/b", true, false, "").  // should match /a/b
		WithTest("http://h/a/c", false, false, ""). // should match /a, not /a/c
		WithTest("http://h/a/d", false, true, "").  // should match /a/d and use https
		WithTest("http://h/a/e", false, false, ""). // should match /a, not /a/e
		Test(t)

	newTestCase().
		WithRule("/a/f", false, false, "r").
		WithTest("http://h/a/f", false, false, "r"). // should match /a/f and redirect
		Test(t)

	newTestCase().
		WithTest("http://h/a", true, false, "").
		TestMirror(t)

	newTestCase().
		WithRegistryMirror("http://index.docker.io", false, false).
		WithTest("http://h/a", true, false, "").
		TestMirror(t)

	newTestCase().
		WithRegistryMirror("http://index.docker.io", false, false).
		WithTest("http://index.docker.io/v2/blobs/sha256/xxx", false, false, "").
		TestMirror(t)

	newTestCase().
		WithRegistryMirror("http://index.docker.io", true, false).
		WithTest("http://index.docker.io/v2/blobs/sha256/xxx", true, false, "").
		TestMirror(t)
}
