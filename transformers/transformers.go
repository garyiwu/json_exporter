// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transformers

import "fmt"

type Transformer interface {
	Transform(data []byte) ([]byte, error)
}

type TransformationConfig struct {
	Type  string
	Query string
}

func NewTransformer(config TransformationConfig) (Transformer, error) {
	switch config.Type {
	case "jq":
		return NewJQTransformer(config.Query), nil
	default:
		return nil, fmt.Errorf("unsupported transformer type: %s", config.Type)
	}
}
