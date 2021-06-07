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

// HealthCheck struct for HealthCheck
type HealthCheck struct {
	Name    string            `json:"name"`
	Message *string           `json:"message,omitempty"`
	Checks  *[]HealthCheck    `json:"checks,omitempty"`
	Status  HealthCheckStatus `json:"status"`
	Version *string           `json:"version,omitempty"`
	Commit  *string           `json:"commit,omitempty"`
}

// NewHealthCheck instantiates a new HealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheck(name string, status HealthCheckStatus) *HealthCheck {
	this := HealthCheck{}
	this.Name = name
	this.Status = status
	return &this
}

// NewHealthCheckWithDefaults instantiates a new HealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckWithDefaults() *HealthCheck {
	this := HealthCheck{}
	return &this
}

// GetName returns the Name field value
func (o *HealthCheck) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HealthCheck) SetName(v string) {
	o.Name = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HealthCheck) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HealthCheck) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HealthCheck) SetMessage(v string) {
	o.Message = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *HealthCheck) GetChecks() []HealthCheck {
	if o == nil || o.Checks == nil {
		var ret []HealthCheck
		return ret
	}
	return *o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetChecksOk() (*[]HealthCheck, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *HealthCheck) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []HealthCheck and assigns it to the Checks field.
func (o *HealthCheck) SetChecks(v []HealthCheck) {
	o.Checks = &v
}

// GetStatus returns the Status field value
func (o *HealthCheck) GetStatus() HealthCheckStatus {
	if o == nil {
		var ret HealthCheckStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetStatusOk() (*HealthCheckStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HealthCheck) SetStatus(v HealthCheckStatus) {
	o.Status = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HealthCheck) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HealthCheck) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HealthCheck) SetVersion(v string) {
	o.Version = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *HealthCheck) GetCommit() string {
	if o == nil || o.Commit == nil {
		var ret string
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetCommitOk() (*string, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *HealthCheck) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given string and assigns it to the Commit field.
func (o *HealthCheck) SetCommit(v string) {
	o.Commit = &v
}

func (o HealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	return json.Marshal(toSerialize)
}

type NullableHealthCheck struct {
	value *HealthCheck
	isSet bool
}

func (v NullableHealthCheck) Get() *HealthCheck {
	return v.value
}

func (v *NullableHealthCheck) Set(val *HealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheck(val *HealthCheck) *NullableHealthCheck {
	return &NullableHealthCheck{value: val, isSet: true}
}

func (v NullableHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}