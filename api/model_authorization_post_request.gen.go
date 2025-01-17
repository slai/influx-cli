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

// AuthorizationPostRequest struct for AuthorizationPostRequest
type AuthorizationPostRequest struct {
	// If inactive the token is inactive and requests using the token will be rejected.
	Status *string `json:"status,omitempty" yaml:"status,omitempty"`
	// A description of the token.
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// ID of org that authorization is scoped to.
	OrgID string `json:"orgID" yaml:"orgID"`
	// ID of user that authorization is scoped to.
	UserID *string `json:"userID,omitempty" yaml:"userID,omitempty"`
	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions []Permission `json:"permissions" yaml:"permissions"`
}

// NewAuthorizationPostRequest instantiates a new AuthorizationPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationPostRequest(orgID string, permissions []Permission) *AuthorizationPostRequest {
	this := AuthorizationPostRequest{}
	var status string = "active"
	this.Status = &status
	this.OrgID = orgID
	this.Permissions = permissions
	return &this
}

// NewAuthorizationPostRequestWithDefaults instantiates a new AuthorizationPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationPostRequestWithDefaults() *AuthorizationPostRequest {
	this := AuthorizationPostRequest{}
	var status string = "active"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizationPostRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizationPostRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizationPostRequest) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizationPostRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizationPostRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizationPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOrgID returns the OrgID field value
func (o *AuthorizationPostRequest) GetOrgID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequest) GetOrgIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgID, true
}

// SetOrgID sets field value
func (o *AuthorizationPostRequest) SetOrgID(v string) {
	o.OrgID = v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *AuthorizationPostRequest) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequest) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *AuthorizationPostRequest) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *AuthorizationPostRequest) SetUserID(v string) {
	o.UserID = &v
}

// GetPermissions returns the Permissions field value
func (o *AuthorizationPostRequest) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AuthorizationPostRequest) GetPermissionsOk() (*[]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *AuthorizationPostRequest) SetPermissions(v []Permission) {
	o.Permissions = v
}

func (o AuthorizationPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["orgID"] = o.OrgID
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationPostRequest struct {
	value *AuthorizationPostRequest
	isSet bool
}

func (v NullableAuthorizationPostRequest) Get() *AuthorizationPostRequest {
	return v.value
}

func (v *NullableAuthorizationPostRequest) Set(val *AuthorizationPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationPostRequest(val *AuthorizationPostRequest) *NullableAuthorizationPostRequest {
	return &NullableAuthorizationPostRequest{value: val, isSet: true}
}

func (v NullableAuthorizationPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
