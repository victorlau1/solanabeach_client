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

// NonvalidatorsPubkey struct for NonvalidatorsPubkey
type NonvalidatorsPubkey struct {
	Address *string `json:"address,omitempty"`
}

// NewNonvalidatorsPubkey instantiates a new NonvalidatorsPubkey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonvalidatorsPubkey() *NonvalidatorsPubkey {
	this := NonvalidatorsPubkey{}
	return &this
}

// NewNonvalidatorsPubkeyWithDefaults instantiates a new NonvalidatorsPubkey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonvalidatorsPubkeyWithDefaults() *NonvalidatorsPubkey {
	this := NonvalidatorsPubkey{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NonvalidatorsPubkey) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonvalidatorsPubkey) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NonvalidatorsPubkey) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NonvalidatorsPubkey) SetAddress(v string) {
	o.Address = &v
}

func (o NonvalidatorsPubkey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableNonvalidatorsPubkey struct {
	value *NonvalidatorsPubkey
	isSet bool
}

func (v NullableNonvalidatorsPubkey) Get() *NonvalidatorsPubkey {
	return v.value
}

func (v *NullableNonvalidatorsPubkey) Set(val *NonvalidatorsPubkey) {
	v.value = val
	v.isSet = true
}

func (v NullableNonvalidatorsPubkey) IsSet() bool {
	return v.isSet
}

func (v *NullableNonvalidatorsPubkey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonvalidatorsPubkey(val *NonvalidatorsPubkey) *NullableNonvalidatorsPubkey {
	return &NullableNonvalidatorsPubkey{value: val, isSet: true}
}

func (v NullableNonvalidatorsPubkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonvalidatorsPubkey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


