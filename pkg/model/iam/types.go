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

package iam

import (
	"encoding/json"
	"fmt"
)

// ParseStatements parses JSON into a list of Statements
func ParseStatements(policy string) ([]*Statement, error) {
	statements := make([]*Statement, 0)
	if err := json.Unmarshal([]byte(policy), &statements); err != nil {
		return nil, fmt.Errorf("error parsing IAM statements: %v", err)
	}
	return statements, nil
}
