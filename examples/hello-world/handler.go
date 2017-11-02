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

// TODO: Is this still true?
// Currently, if processor.yaml is not provided, nuclio-build will look for
// a Handler() function in the handler package. Future implementations
// will alleviate this limitation

package handler

import (
	"github.com/nuclio/nuclio-sdk"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.Info("Event received")

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/text",
		Body:        []byte("Response from handler"),
	}, nil
}
