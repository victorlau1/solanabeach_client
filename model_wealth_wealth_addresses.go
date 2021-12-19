/*
Solanabeach Backend API

Solanabeach Backend REST API documentation.  ## Rate limiting Current API rate limit per IP is 100 requests per second.    ## Pagination Most of the endpoints returning array data support pagination. The API uses two types of pagination, depending on the returned data. Some endpoints support both, some are limited to just one type.    ## Supported pagination types ### Offset / limit Offset / limit pagination is the simplest form of pagination supported by the API. Offset parameter represents the number of results to skip before returning the data, and the limit parameter limits the number of returned results.   To prevent overloading the API, all limit params have a max value. Each API endpoint has its own max value.  ### Cursor Cursor pagination is more complex than the offset / limit one, but, it comes naturally for some blockchain data (such as blocks, transactions, token transfers, etc). Cursors contain data like blocknumber, transaction index, ... and they're described on their respective API endpoints. Limit parameter works exactly the same way as it does in the offset / limit pagination.  ## Authentication The public API uses a Bearer OAuth authentication method, and the API key should be provided in the `Authorization` header in each request. API keys are issued on request. 

API version: 0.0.1
Contact: andrej@vgng.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WealthWealthAddresses struct for WealthWealthAddresses
type WealthWealthAddresses struct {
	Count *int32 `json:"count,omitempty"`
	Percentage *string `json:"percentage,omitempty"`
}

// NewWealthWealthAddresses instantiates a new WealthWealthAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWealthWealthAddresses() *WealthWealthAddresses {
	this := WealthWealthAddresses{}
	return &this
}

// NewWealthWealthAddressesWithDefaults instantiates a new WealthWealthAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWealthWealthAddressesWithDefaults() *WealthWealthAddresses {
	this := WealthWealthAddresses{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WealthWealthAddresses) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WealthWealthAddresses) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WealthWealthAddresses) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *WealthWealthAddresses) SetCount(v int32) {
	o.Count = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *WealthWealthAddresses) GetPercentage() string {
	if o == nil || o.Percentage == nil {
		var ret string
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WealthWealthAddresses) GetPercentageOk() (*string, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *WealthWealthAddresses) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given string and assigns it to the Percentage field.
func (o *WealthWealthAddresses) SetPercentage(v string) {
	o.Percentage = &v
}

func (o WealthWealthAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}
	return json.Marshal(toSerialize)
}

type NullableWealthWealthAddresses struct {
	value *WealthWealthAddresses
	isSet bool
}

func (v NullableWealthWealthAddresses) Get() *WealthWealthAddresses {
	return v.value
}

func (v *NullableWealthWealthAddresses) Set(val *WealthWealthAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableWealthWealthAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableWealthWealthAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWealthWealthAddresses(val *WealthWealthAddresses) *NullableWealthWealthAddresses {
	return &NullableWealthWealthAddresses{value: val, isSet: true}
}

func (v NullableWealthWealthAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWealthWealthAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


