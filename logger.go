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

type Logger interface {

	// emit a log entry of a given verbosity. the first argument may be an object, a string
	// or a format string. in case of the latter, the following varargs are passed
	// to a formatter (e.g. fmt.Sprintf)
	Error(format interface{}, vars ...interface{})
	Warn(format interface{}, vars ...interface{})
	Info(format interface{}, vars ...interface{})
	Debug(format interface{}, vars ...interface{})

	// emit a structured log entry. example:
	//
	// l.InfoWith("The message",
	// 	"first-key", "first-value",
	// 	"second-key", 2)
	//
	ErrorWith(format interface{}, vars ...interface{})
	WarnWith(format interface{}, vars ...interface{})
	InfoWith(format interface{}, vars ...interface{})
	DebugWith(format interface{}, vars ...interface{})

	// flushes buffered logs, if applicable
	Flush()

	// returns a child logger, if underlying logger supports hierarchal logging
	GetChild(name string) Logger
}
