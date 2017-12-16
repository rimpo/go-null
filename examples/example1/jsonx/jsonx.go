package jsonx

//Auto-generated code; DONT EDIT THIS CODE

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/rimpo/go-null/examples/example1/null"
)

const (
	ignoreTagName = "___"
)

//start-------- Code copied from golang's encoding/json libaray --------------

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}

//end -----------------------------------------------------------------------

//Types Interface used for omitting field when empty
type INull interface {
	IsNull() bool
}

func truncateLastComma(jsonData *bytes.Buffer) {
	if jsonData.Len() > 1 && string(jsonData.Bytes()[jsonData.Len()-1:]) == "," {
		jsonData.Truncate(jsonData.Len() - 1)
	}
}

func writeValue(jsonData *bytes.Buffer, tag string, v interface{}) error {
	m, ok := v.(json.Marshaler)
	if !ok {
		return errors.New("Failed interface conversion to json.Marshaler")
	}
	jsonData.WriteString(`"`)
	jsonData.WriteString(tag)
	jsonData.WriteString(`":`)
	b, err := m.MarshalJSON()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to MarshalJSON %!v(MISSING)", err))
	}
	jsonData.Write(b)
	return nil
}

func writeValueWithoutTag(jsonData *bytes.Buffer, v interface{}) error {
	m, ok := v.(json.Marshaler)
	if !ok {
		return errors.New("Failed interface conversion to json.Marshaler")
	}
	b, err := m.MarshalJSON()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to MarshalJSON %!v(MISSING)", err))
	}
	jsonData.Write(b)
	return nil
}

func Marshal(src interface{}, jsonData *bytes.Buffer) {
	getJSONLoop(reflect.ValueOf(src).Elem(), jsonData, ignoreTagName)
}

func getJSONLoop(v reflect.Value, jsonData *bytes.Buffer, jsonTag string) error {
	switch v.Type().Kind() {
	case reflect.Slice:
		tag, tagOpt := parseTag(jsonTag)
		//only if json tag exist processing is done
		if len(tag) > 0 {
			var jsonDataStruct bytes.Buffer

			//write the json key except for ignore tag
			var jsonKey bytes.Buffer
			if tag == ignoreTagName {
				jsonKey.WriteString("[")
			} else {
				jsonKey.WriteString(`"`)
				jsonKey.WriteString(tag)
				jsonKey.WriteString(`":[`)
			}
			jsonDataStruct.WriteString(jsonKey.String())

			for i := 0; i < v.Len(); i++ {
				getJSONLoop(v.Index(i), &jsonDataStruct, ignoreTagName)
				jsonDataStruct.WriteString(",")
			}
			truncateLastComma(&jsonDataStruct)

			jsonDataStruct.WriteString("]")
			if tagOpt.contains("omitempty") && jsonDataStruct.Len() == jsonKey.Len()+1 {
				//ignore if empty
			} else {
				jsonData.Write(jsonDataStruct.Bytes())
			}

		}
	case reflect.Struct:
		_, ok := null.AllTypes[v.Type().Name()]
		if ok {
			//field type
			tag, tagOpt := parseTag(jsonTag)
			if tag == ignoreTagName {
				//slice element (without tag printing is needed for slice)
				writeValueWithoutTag(jsonData, v.Addr().Interface())
			} else if tagOpt.contains("omitempty") {
				//struct field - with tag name and omitempty
				valid, ok := v.Addr().Interface().(INull)
				if ok && !valid.IsNull() {
					writeValue(jsonData, tag, v.Addr().Interface())
				}
			} else if len(tag) > 0 {
				//struct field - with only tag name
				writeValue(jsonData, tag, v.Addr().Interface())
			}
			// Note: struct field without tag will be ignored
		} else {
			//struct
			tag, tagOpt := parseTag(jsonTag)
			//only if json tag exist processing is done
			if len(tag) > 0 {
				var jsonDataStruct bytes.Buffer
				//write key for json
				var jsonKey bytes.Buffer
				if tag == ignoreTagName {
					jsonKey.WriteString("{")
				} else {
					jsonKey.WriteString(`"`)
					jsonKey.WriteString(tag)
					jsonKey.WriteString(`":{`)
				}
				jsonDataStruct.WriteString(jsonKey.String())
				//each field
				var initialLen int
				t := v.Type()
				for i := 0; i < t.NumField(); i++ {
					initialLen = jsonDataStruct.Len()
					getJSONLoop(v.Field(i), &jsonDataStruct, t.Field(i).Tag.Get("json"))
					//change in json data then only output comma
					if initialLen != jsonDataStruct.Len() {
						jsonDataStruct.WriteString(",")
					}
				}
				truncateLastComma(&jsonDataStruct)

				jsonDataStruct.WriteString("}")

				//omitempty struct should not show if the struct is empty
				if tagOpt.contains("omitempty") && jsonDataStruct.Len() == jsonKey.Len()+1 {
					//ignore if empty
				} else {
					jsonData.Write(jsonDataStruct.Bytes())
				}
			}

		}
	}
	return nil
}
