package creds

import (
	"encoding/json"

	"github.com/concourse/concourse/v7/vars"
	"sigs.k8s.io/yaml"
)

func evaluate(variablesResolver vars.Variables, in, out interface{}) error {
	byteParams, err := json.Marshal(in)
	if err != nil {
		return err
	}

	tpl := vars.NewTemplate(byteParams)

	bytes, err := tpl.Evaluate(variablesResolver, vars.EvaluateOpts{
		ExpectAllKeys: true,
	})
	if err != nil {
		return err
	}

	return yaml.Unmarshal(bytes, out, useJSONNumber)
}

func useJSONNumber(decoder *json.Decoder) *json.Decoder {
	decoder.UseNumber()
	return decoder
}
