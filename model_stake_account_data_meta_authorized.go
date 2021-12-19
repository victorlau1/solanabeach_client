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

// StakeAccountDataMetaAuthorized struct for StakeAccountDataMetaAuthorized
type StakeAccountDataMetaAuthorized struct {
	Staker *string `json:"staker,omitempty"`
	Withdrawer *string `json:"withdrawer,omitempty"`
}

// NewStakeAccountDataMetaAuthorized instantiates a new StakeAccountDataMetaAuthorized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakeAccountDataMetaAuthorized() *StakeAccountDataMetaAuthorized {
	this := StakeAccountDataMetaAuthorized{}
	return &this
}

// NewStakeAccountDataMetaAuthorizedWithDefaults instantiates a new StakeAccountDataMetaAuthorized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakeAccountDataMetaAuthorizedWithDefaults() *StakeAccountDataMetaAuthorized {
	this := StakeAccountDataMetaAuthorized{}
	return &this
}

// GetStaker returns the Staker field value if set, zero value otherwise.
func (o *StakeAccountDataMetaAuthorized) GetStaker() string {
	if o == nil || o.Staker == nil {
		var ret string
		return ret
	}
	return *o.Staker
}

// GetStakerOk returns a tuple with the Staker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountDataMetaAuthorized) GetStakerOk() (*string, bool) {
	if o == nil || o.Staker == nil {
		return nil, false
	}
	return o.Staker, true
}

// HasStaker returns a boolean if a field has been set.
func (o *StakeAccountDataMetaAuthorized) HasStaker() bool {
	if o != nil && o.Staker != nil {
		return true
	}

	return false
}

// SetStaker gets a reference to the given string and assigns it to the Staker field.
func (o *StakeAccountDataMetaAuthorized) SetStaker(v string) {
	o.Staker = &v
}

// GetWithdrawer returns the Withdrawer field value if set, zero value otherwise.
func (o *StakeAccountDataMetaAuthorized) GetWithdrawer() string {
	if o == nil || o.Withdrawer == nil {
		var ret string
		return ret
	}
	return *o.Withdrawer
}

// GetWithdrawerOk returns a tuple with the Withdrawer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountDataMetaAuthorized) GetWithdrawerOk() (*string, bool) {
	if o == nil || o.Withdrawer == nil {
		return nil, false
	}
	return o.Withdrawer, true
}

// HasWithdrawer returns a boolean if a field has been set.
func (o *StakeAccountDataMetaAuthorized) HasWithdrawer() bool {
	if o != nil && o.Withdrawer != nil {
		return true
	}

	return false
}

// SetWithdrawer gets a reference to the given string and assigns it to the Withdrawer field.
func (o *StakeAccountDataMetaAuthorized) SetWithdrawer(v string) {
	o.Withdrawer = &v
}

func (o StakeAccountDataMetaAuthorized) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Staker != nil {
		toSerialize["staker"] = o.Staker
	}
	if o.Withdrawer != nil {
		toSerialize["withdrawer"] = o.Withdrawer
	}
	return json.Marshal(toSerialize)
}

type NullableStakeAccountDataMetaAuthorized struct {
	value *StakeAccountDataMetaAuthorized
	isSet bool
}

func (v NullableStakeAccountDataMetaAuthorized) Get() *StakeAccountDataMetaAuthorized {
	return v.value
}

func (v *NullableStakeAccountDataMetaAuthorized) Set(val *StakeAccountDataMetaAuthorized) {
	v.value = val
	v.isSet = true
}

func (v NullableStakeAccountDataMetaAuthorized) IsSet() bool {
	return v.isSet
}

func (v *NullableStakeAccountDataMetaAuthorized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakeAccountDataMetaAuthorized(val *StakeAccountDataMetaAuthorized) *NullableStakeAccountDataMetaAuthorized {
	return &NullableStakeAccountDataMetaAuthorized{value: val, isSet: true}
}

func (v NullableStakeAccountDataMetaAuthorized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakeAccountDataMetaAuthorized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


