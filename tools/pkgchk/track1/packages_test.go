// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package track1

import (
	"path/filepath"
	"strings"
	"testing"
)

const (
	testRoot = "../../testpkgs"
)

var (
	expected = map[string]string{
		"scenrioa/foo":                    "foo",
		"scenriob/foo":                    "foo",
		"scenriob/foo/v2":                 "foo",
		"scenrioc/mgmt/2019-10-11/foo":    "foo",
		"scenriod/mgmt/2019-10-11/foo":    "foo",
		"scenriod/mgmt/2019-10-11/foo/v2": "foo",
		"scenrioe/mgmt/2019-10-11/foo":    "foo",
		"scenrioe/mgmt/2019-10-11/foo/v2": "foo",
		"scenrioe/mgmt/2019-10-11/foo/v3": "foo",
	}
)

func TestList(t *testing.T) {
	root, err := filepath.Abs(testRoot)
	if err != nil {
		t.Fatalf("failed to get absolute path: %+v", err)
	}
	pkgs, err := List(root)
	if err != nil {
		t.Fatalf("failed to get packages: %+v", err)
	}
	if len(pkgs) != len(expected) {
		t.Fatalf("expected %d packages, but got %d", len(expected), len(pkgs))
	}
	for _, pkg := range pkgs {
		if pkgName, ok := expected[pkg.Path()]; !ok {
			t.Fatalf("got pkg path '%s', but not found in expected", pkg.Path())
		} else if !strings.EqualFold(pkgName, pkg.Name()) {
			t.Fatalf("expected package of '%s' in path '%s', but got '%s'", pkgName, pkg.Path(), pkg.Name())
		}
	}
}

func TestVerifier_Verify(t *testing.T) {
	root, err := filepath.Abs(testRoot)
	if err != nil {
		t.Fatalf("failed to get absolute path: %+v", err)
	}
	pkgs, err := List(root)
	if err != nil {
		t.Fatalf("failed to get packages: %+v", err)
	}

	verifier := GetDefaultVerifier()
	for _, pkg := range pkgs {
		if errors := verifier.Verify(pkg); len(errors) != 0 {
			t.Fatalf("failed to verify packages: %+v", errors)
		}
	}
}
