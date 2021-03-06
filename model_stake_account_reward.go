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

// StakeAccountReward struct for StakeAccountReward
type StakeAccountReward struct {
	Epoch *int32 `json:"epoch,omitempty"`
	EffectiveSlot *int32 `json:"effectiveSlot,omitempty"`
	Amount *string `json:"amount,omitempty"`
	PostBalance *string `json:"postBalance,omitempty"`
	PercentChange *int32 `json:"percentChange,omitempty"`
	Apr *int32 `json:"apr,omitempty"`
}

// NewStakeAccountReward instantiates a new StakeAccountReward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakeAccountReward() *StakeAccountReward {
	this := StakeAccountReward{}
	return &this
}

// NewStakeAccountRewardWithDefaults instantiates a new StakeAccountReward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakeAccountRewardWithDefaults() *StakeAccountReward {
	this := StakeAccountReward{}
	return &this
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *StakeAccountReward) GetEpoch() int32 {
	if o == nil || o.Epoch == nil {
		var ret int32
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetEpochOk() (*int32, bool) {
	if o == nil || o.Epoch == nil {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *StakeAccountReward) HasEpoch() bool {
	if o != nil && o.Epoch != nil {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given int32 and assigns it to the Epoch field.
func (o *StakeAccountReward) SetEpoch(v int32) {
	o.Epoch = &v
}

// GetEffectiveSlot returns the EffectiveSlot field value if set, zero value otherwise.
func (o *StakeAccountReward) GetEffectiveSlot() int32 {
	if o == nil || o.EffectiveSlot == nil {
		var ret int32
		return ret
	}
	return *o.EffectiveSlot
}

// GetEffectiveSlotOk returns a tuple with the EffectiveSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetEffectiveSlotOk() (*int32, bool) {
	if o == nil || o.EffectiveSlot == nil {
		return nil, false
	}
	return o.EffectiveSlot, true
}

// HasEffectiveSlot returns a boolean if a field has been set.
func (o *StakeAccountReward) HasEffectiveSlot() bool {
	if o != nil && o.EffectiveSlot != nil {
		return true
	}

	return false
}

// SetEffectiveSlot gets a reference to the given int32 and assigns it to the EffectiveSlot field.
func (o *StakeAccountReward) SetEffectiveSlot(v int32) {
	o.EffectiveSlot = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *StakeAccountReward) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *StakeAccountReward) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *StakeAccountReward) SetAmount(v string) {
	o.Amount = &v
}

// GetPostBalance returns the PostBalance field value if set, zero value otherwise.
func (o *StakeAccountReward) GetPostBalance() string {
	if o == nil || o.PostBalance == nil {
		var ret string
		return ret
	}
	return *o.PostBalance
}

// GetPostBalanceOk returns a tuple with the PostBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetPostBalanceOk() (*string, bool) {
	if o == nil || o.PostBalance == nil {
		return nil, false
	}
	return o.PostBalance, true
}

// HasPostBalance returns a boolean if a field has been set.
func (o *StakeAccountReward) HasPostBalance() bool {
	if o != nil && o.PostBalance != nil {
		return true
	}

	return false
}

// SetPostBalance gets a reference to the given string and assigns it to the PostBalance field.
func (o *StakeAccountReward) SetPostBalance(v string) {
	o.PostBalance = &v
}

// GetPercentChange returns the PercentChange field value if set, zero value otherwise.
func (o *StakeAccountReward) GetPercentChange() int32 {
	if o == nil || o.PercentChange == nil {
		var ret int32
		return ret
	}
	return *o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetPercentChangeOk() (*int32, bool) {
	if o == nil || o.PercentChange == nil {
		return nil, false
	}
	return o.PercentChange, true
}

// HasPercentChange returns a boolean if a field has been set.
func (o *StakeAccountReward) HasPercentChange() bool {
	if o != nil && o.PercentChange != nil {
		return true
	}

	return false
}

// SetPercentChange gets a reference to the given int32 and assigns it to the PercentChange field.
func (o *StakeAccountReward) SetPercentChange(v int32) {
	o.PercentChange = &v
}

// GetApr returns the Apr field value if set, zero value otherwise.
func (o *StakeAccountReward) GetApr() int32 {
	if o == nil || o.Apr == nil {
		var ret int32
		return ret
	}
	return *o.Apr
}

// GetAprOk returns a tuple with the Apr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StakeAccountReward) GetAprOk() (*int32, bool) {
	if o == nil || o.Apr == nil {
		return nil, false
	}
	return o.Apr, true
}

// HasApr returns a boolean if a field has been set.
func (o *StakeAccountReward) HasApr() bool {
	if o != nil && o.Apr != nil {
		return true
	}

	return false
}

// SetApr gets a reference to the given int32 and assigns it to the Apr field.
func (o *StakeAccountReward) SetApr(v int32) {
	o.Apr = &v
}

func (o StakeAccountReward) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Epoch != nil {
		toSerialize["epoch"] = o.Epoch
	}
	if o.EffectiveSlot != nil {
		toSerialize["effectiveSlot"] = o.EffectiveSlot
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.PostBalance != nil {
		toSerialize["postBalance"] = o.PostBalance
	}
	if o.PercentChange != nil {
		toSerialize["percentChange"] = o.PercentChange
	}
	if o.Apr != nil {
		toSerialize["apr"] = o.Apr
	}
	return json.Marshal(toSerialize)
}

type NullableStakeAccountReward struct {
	value *StakeAccountReward
	isSet bool
}

func (v NullableStakeAccountReward) Get() *StakeAccountReward {
	return v.value
}

func (v *NullableStakeAccountReward) Set(val *StakeAccountReward) {
	v.value = val
	v.isSet = true
}

func (v NullableStakeAccountReward) IsSet() bool {
	return v.isSet
}

func (v *NullableStakeAccountReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakeAccountReward(val *StakeAccountReward) *NullableStakeAccountReward {
	return &NullableStakeAccountReward{value: val, isSet: true}
}

func (v NullableStakeAccountReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakeAccountReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


