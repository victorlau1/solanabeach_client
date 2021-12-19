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

// InlineResponse20012 struct for InlineResponse20012
type InlineResponse20012 struct {
	Apy *int32 `json:"apy,omitempty"`
}

// NewInlineResponse20012 instantiates a new InlineResponse20012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012WithDefaults() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *InlineResponse20012) GetApy() int32 {
	if o == nil || o.Apy == nil {
		var ret int32
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetApyOk() (*int32, bool) {
	if o == nil || o.Apy == nil {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *InlineResponse20012) HasApy() bool {
	if o != nil && o.Apy != nil {
		return true
	}

	return false
}

// SetApy gets a reference to the given int32 and assigns it to the Apy field.
func (o *InlineResponse20012) SetApy(v int32) {
	o.Apy = &v
}

func (o InlineResponse20012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apy != nil {
		toSerialize["apy"] = o.Apy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012 struct {
	value *InlineResponse20012
	isSet bool
}

func (v NullableInlineResponse20012) Get() *InlineResponse20012 {
	return v.value
}

func (v *NullableInlineResponse20012) Set(val *InlineResponse20012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012(val *InlineResponse20012) *NullableInlineResponse20012 {
	return &NullableInlineResponse20012{value: val, isSet: true}
}

func (v NullableInlineResponse20012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

