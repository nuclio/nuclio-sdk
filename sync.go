/*
Copyright 2017 The Nuclio Authors.
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
package nuclio

type Sync interface {
	Event
	GetHostAddress() string
	GetRemoteAddress() string
	GetWorkflowStep() string
	GetQuery() map[string]interface{}
}

type AbstractSync struct {
	AbstractEvent
}

func (as *AbstractSync) GetHostAddress() string {
	return ""
}

func (as *AbstractSync) GetRemoteAddress() string {
	return ""
}

func (as *AbstractSync) GetWorkflowStep() string {
	return ""
}

func (as *AbstractSync) GetQuery() map[string]interface{} {
	return map[string]interface{}{}
}
