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

// AccountToken struct for AccountToken
type AccountToken struct {
	Mint *Address `json:"mint,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	Address *Address `json:"address,omitempty"`
	Decimals *int32 `json:"decimals,omitempty"`
}

// NewAccountToken instantiates a new AccountToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountToken() *AccountToken {
	this := AccountToken{}
	return &this
}

// NewAccountTokenWithDefaults instantiates a new AccountToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTokenWithDefaults() *AccountToken {
	this := AccountToken{}
	return &this
}

// GetMint returns the Mint field value if set, zero value otherwise.
func (o *AccountToken) GetMint() Address {
	if o == nil || o.Mint == nil {
		var ret Address
		return ret
	}
	return *o.Mint
}

// GetMintOk returns a tuple with the Mint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetMintOk() (*Address, bool) {
	if o == nil || o.Mint == nil {
		return nil, false
	}
	return o.Mint, true
}

// HasMint returns a boolean if a field has been set.
func (o *AccountToken) HasMint() bool {
	if o != nil && o.Mint != nil {
		return true
	}

	return false
}

// SetMint gets a reference to the given Address and assigns it to the Mint field.
func (o *AccountToken) SetMint(v Address) {
	o.Mint = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AccountToken) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AccountToken) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *AccountToken) SetAmount(v int32) {
	o.Amount = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AccountToken) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountToken) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *AccountToken) SetAddress(v Address) {
	o.Address = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *AccountToken) GetDecimals() int32 {
	if o == nil || o.Decimals == nil {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountToken) GetDecimalsOk() (*int32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *AccountToken) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *AccountToken) SetDecimals(v int32) {
	o.Decimals = &v
}

func (o AccountToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mint != nil {
		toSerialize["mint"] = o.Mint
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	return json.Marshal(toSerialize)
}

type NullableAccountToken struct {
	value *AccountToken
	isSet bool
}

func (v NullableAccountToken) Get() *AccountToken {
	return v.value
}

func (v *NullableAccountToken) Set(val *AccountToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountToken(val *AccountToken) *NullableAccountToken {
	return &NullableAccountToken{value: val, isSet: true}
}

func (v NullableAccountToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


