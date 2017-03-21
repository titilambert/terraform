package kubernetes

import (
	"encoding/json"
	"reflect"
)

func diffStringMap(pathPrefix string, oldV, newV map[string]interface{}) PatchOperations {
	ops := make([]PatchOperation, 0, 0)

	for k, _ := range oldV {
		if _, ok := newV[k]; ok {
			continue
		}
		ops = append(ops, &RemoveOperation{Path: pathPrefix + k})
	}

	for k, v := range newV {
		newValue := v.(string)

		if oldValue, ok := oldV[k].(string); ok {
			if oldValue == newValue {
				continue
			}

			ops = append(ops, &ReplaceOperation{
				Path:  pathPrefix + k,
				Value: v.(string),
			})
			continue
		}

		ops = append(ops, &AddOperation{
			Path:  pathPrefix + k,
			Value: newValue,
		})
	}

	return ops
}

type PatchOperations []PatchOperation

func (po PatchOperations) MarshalJSON() ([]byte, error) {
	var v []PatchOperation = po
	return json.Marshal(v)
}

func (po PatchOperations) Equal(ops []PatchOperation) bool {
	var v []PatchOperation = po
	return reflect.DeepEqual(v, ops)
}

type PatchOperation interface {
	MarshalJSON() ([]byte, error)
}

type ReplaceOperation struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
	Op    string      `json:"op"`
}

func (o *ReplaceOperation) MarshalJSON() ([]byte, error) {
	o.Op = "replace"
	return json.Marshal(*o)
}

func (o *ReplaceOperation) String() string {
	b, _ := o.MarshalJSON()
	return string(b)
}

type AddOperation struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
	Op    string      `json:"op"`
}

func (o *AddOperation) MarshalJSON() ([]byte, error) {
	o.Op = "add"
	return json.Marshal(*o)
}

func (o *AddOperation) String() string {
	b, _ := o.MarshalJSON()
	return string(b)
}

type RemoveOperation struct {
	Path string `json:"path"`
	Op   string `json:"op"`
}

func (o *RemoveOperation) MarshalJSON() ([]byte, error) {
	o.Op = "remove"
	return json.Marshal(*o)
}

func (o *RemoveOperation) String() string {
	b, _ := o.MarshalJSON()
	return string(b)
}
