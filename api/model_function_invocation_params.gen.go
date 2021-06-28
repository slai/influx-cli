/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FunctionInvocationParams struct for FunctionInvocationParams
type FunctionInvocationParams struct {
	Params *map[string]interface{} `json:"params,omitempty" yaml:"params,omitempty"`
}

// NewFunctionInvocationParams instantiates a new FunctionInvocationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionInvocationParams() *FunctionInvocationParams {
	this := FunctionInvocationParams{}
	return &this
}

// NewFunctionInvocationParamsWithDefaults instantiates a new FunctionInvocationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionInvocationParamsWithDefaults() *FunctionInvocationParams {
	this := FunctionInvocationParams{}
	return &this
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *FunctionInvocationParams) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionInvocationParams) GetParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *FunctionInvocationParams) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *FunctionInvocationParams) SetParams(v map[string]interface{}) {
	o.Params = &v
}

func (o FunctionInvocationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionInvocationParams struct {
	value *FunctionInvocationParams
	isSet bool
}

func (v NullableFunctionInvocationParams) Get() *FunctionInvocationParams {
	return v.value
}

func (v *NullableFunctionInvocationParams) Set(val *FunctionInvocationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionInvocationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionInvocationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionInvocationParams(val *FunctionInvocationParams) *NullableFunctionInvocationParams {
	return &NullableFunctionInvocationParams{value: val, isSet: true}
}

func (v NullableFunctionInvocationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionInvocationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}