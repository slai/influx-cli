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

// FunctionUpdateRequest struct for FunctionUpdateRequest
type FunctionUpdateRequest struct {
	Name        *string `json:"name,omitempty" yaml:"name,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// script is script to be executed
	Script *string `json:"script,omitempty" yaml:"script,omitempty"`
}

// NewFunctionUpdateRequest instantiates a new FunctionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionUpdateRequest() *FunctionUpdateRequest {
	this := FunctionUpdateRequest{}
	return &this
}

// NewFunctionUpdateRequestWithDefaults instantiates a new FunctionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionUpdateRequestWithDefaults() *FunctionUpdateRequest {
	this := FunctionUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FunctionUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FunctionUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FunctionUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FunctionUpdateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FunctionUpdateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FunctionUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *FunctionUpdateRequest) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionUpdateRequest) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *FunctionUpdateRequest) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *FunctionUpdateRequest) SetScript(v string) {
	o.Script = &v
}

func (o FunctionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionUpdateRequest struct {
	value *FunctionUpdateRequest
	isSet bool
}

func (v NullableFunctionUpdateRequest) Get() *FunctionUpdateRequest {
	return v.value
}

func (v *NullableFunctionUpdateRequest) Set(val *FunctionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionUpdateRequest(val *FunctionUpdateRequest) *NullableFunctionUpdateRequest {
	return &NullableFunctionUpdateRequest{value: val, isSet: true}
}

func (v NullableFunctionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}