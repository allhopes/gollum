// Copyright 2015-2018 trivago N.V.
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

package format

import (
	"fmt"

	"github.com/trivago/gollum/core"
)

// ToArray formatter plugin
//
// ToArray converts a set of metadata keys to CSV and applies it to ApplyTo.
//
// Parameters
//
// - Keys: List of strings specifying the keys to write as CSV.
// Note that these keys can be paths.
// See https://godoc.org/github.com/trivago/tgo/tcontainer#MarshalMap.Value.
//
// - Separator: The delimited string to insert between each value in the generated
// string. By default this parameter is set to ",".
//
// Examples
//
// This example get sthe `foo` and `bar` keys from the metdata of a message
// and set this as the new payload.
//
//  exampleProducer:
//    Type: producer.Console
//    Streams: "*"
//    Modulators:
//    - format.ToArray:
//        Separator: ';'
//        Keys:
//        - 'foo'
//        - 'bar'
type ToArray struct {
	core.SimpleFormatter `gollumdoc:"embed_type"`
	separator            string   `config:"Separator" default:","`
	keys                 []string `config:"Keys"`
}

func init() {
	core.TypeRegistry.Register(ToArray{})
}

// Configure initializes this formatter with values from a plugin config.
func (format *ToArray) Configure(conf core.PluginConfigReader) {
}

// ApplyFormatter update message payload
func (format *ToArray) ApplyFormatter(msg *core.Message) error {

	csv := ""
	metadata := msg.GetMetadata()
	for _, key := range format.keys {
		if value, ok := metadata.Value(key); ok {
			switch v := value.(type) {
			case bool:
				csv = fmt.Sprintf("%s%t%s", csv, v, format.separator)
			case int8, int16, int32, int64:
				csv = fmt.Sprintf("%s%d%s", csv, v, format.separator)
			case uint8, uint16, uint32, uint64:
				csv = fmt.Sprintf("%s%d%s", csv, v, format.separator)
			case float32, float64:
				csv = fmt.Sprintf("%s%f%s", csv, v, format.separator)
			case string:
				csv = fmt.Sprintf("%s%s%s", csv, v, format.separator)
			default:
				format.Logger.WithField("key", key).Warning("unsupported datatype")
				csv += format.separator
			}
		} else {
			format.Logger.WithField("key", key).Warning("key not found")
			csv += format.separator
		}
	}

	// Remove last separator
	if len(csv) >= len(format.separator) {
		csv = csv[:len(csv)-len(format.separator)]
	}

	format.SetAppliedContent(msg, csv)
	return nil
}