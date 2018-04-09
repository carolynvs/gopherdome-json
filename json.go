package gopherson

import (
	"bytes"
	"encoding/json"
)

func ToJson(thing interface{}) []byte {
	result, err := json.Marshal(thing)
	if err != nil {
		panic(err)
	}

	return result
}

func FromJson(value []byte, thing interface{}) {
	err := json.Unmarshal(value, &thing)
	if err != nil {
		panic(err)
	}
}

func Pretty(thing interface{}) string {
	value := ToJson(thing)

	var out bytes.Buffer
	json.Indent(&out, value, "", "\t")
	return out.String()
}
