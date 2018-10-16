/*
Copyright 2018 The Kubernetes Authors.

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
	"testing"
)

type testCase struct {
	options   options
	expectErr bool
}

func TestValidate(t *testing.T) {
	cases := []testCase{{
		options: options{
			schemaPath:   "testdata/schema.yaml",
			validatePath: "testdata/schema.yaml",
		},
	}, {
		options: options{
			schemaPath:   "testdata/schema.yaml",
			validatePath: "testdata/bad-schema.yaml",
		},
		expectErr: true,
	}}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.options.validatePath, func(t *testing.T) {
			op, err := tt.options.resolve()
			if err != nil {
				t.Fatal(err)
			}
			err = op.execute()
			if tt.expectErr {
				if err == nil {
					t.Error("unexpected success")
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
