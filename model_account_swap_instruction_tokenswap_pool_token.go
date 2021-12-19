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

// AccountSwapInstructionTokenswapPoolToken struct for AccountSwapInstructionTokenswapPoolToken
type AccountSwapInstructionTokenswapPoolToken struct {
	Decimals *int32 `json:"decimals,omitempty"`
}

// NewAccountSwapInstructionTokenswapPoolToken instantiates a new AccountSwapInstructionTokenswapPoolToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSwapInstructionTokenswapPoolToken() *AccountSwapInstructionTokenswapPoolToken {
	this := AccountSwapInstructionTokenswapPoolToken{}
	return &this
}

// NewAccountSwapInstructionTokenswapPoolTokenWithDefaults instantiates a new AccountSwapInstructionTokenswapPoolToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSwapInstructionTokenswapPoolTokenWithDefaults() *AccountSwapInstructionTokenswapPoolToken {
	this := AccountSwapInstructionTokenswapPoolToken{}
	return &this
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *AccountSwapInstructionTokenswapPoolToken) GetDecimals() int32 {
	if o == nil || o.Decimals == nil {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSwapInstructionTokenswapPoolToken) GetDecimalsOk() (*int32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *AccountSwapInstructionTokenswapPoolToken) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *AccountSwapInstructionTokenswapPoolToken) SetDecimals(v int32) {
	o.Decimals = &v
}

func (o AccountSwapInstructionTokenswapPoolToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	return json.Marshal(toSerialize)
}

type NullableAccountSwapInstructionTokenswapPoolToken struct {
	value *AccountSwapInstructionTokenswapPoolToken
	isSet bool
}

func (v NullableAccountSwapInstructionTokenswapPoolToken) Get() *AccountSwapInstructionTokenswapPoolToken {
	return v.value
}

func (v *NullableAccountSwapInstructionTokenswapPoolToken) Set(val *AccountSwapInstructionTokenswapPoolToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSwapInstructionTokenswapPoolToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSwapInstructionTokenswapPoolToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSwapInstructionTokenswapPoolToken(val *AccountSwapInstructionTokenswapPoolToken) *NullableAccountSwapInstructionTokenswapPoolToken {
	return &NullableAccountSwapInstructionTokenswapPoolToken{value: val, isSet: true}
}

func (v NullableAccountSwapInstructionTokenswapPoolToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSwapInstructionTokenswapPoolToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

