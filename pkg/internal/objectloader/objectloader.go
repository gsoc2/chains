/*
Copyright 2022 The Tekton Authors

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

package objectloader

import (
	"encoding/json"
	"os"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

func TaskRunFromFile(f string) (*v1beta1.TaskRun, error) {
	contents, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var tr v1beta1.TaskRun
	if err := json.Unmarshal(contents, &tr); err != nil {
		return nil, err
	}
	return &tr, nil
}

func PipelineRunFromFile(f string) (*v1beta1.PipelineRun, error) {
	contents, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var pr v1beta1.PipelineRun
	if err := json.Unmarshal(contents, &pr); err != nil {
		return nil, err
	}
	return &pr, nil
}
