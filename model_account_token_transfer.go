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

// AccountTokenTransfer struct for AccountTokenTransfer
type AccountTokenTransfer struct {
	Blocknumber *int32 `json:"blocknumber,omitempty"`
	Txhash *string `json:"txhash,omitempty"`
	Mint *Address `json:"mint,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	Source *Address `json:"source,omitempty"`
	Destination *Address `json:"destination,omitempty"`
	Inner *bool `json:"inner,omitempty"`
	Txindex *int32 `json:"txindex,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	Decimals *int32 `json:"decimals,omitempty"`
}

// NewAccountTokenTransfer instantiates a new AccountTokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTokenTransfer() *AccountTokenTransfer {
	this := AccountTokenTransfer{}
	return &this
}

// NewAccountTokenTransferWithDefaults instantiates a new AccountTokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTokenTransferWithDefaults() *AccountTokenTransfer {
	this := AccountTokenTransfer{}
	return &this
}

// GetBlocknumber returns the Blocknumber field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetBlocknumber() int32 {
	if o == nil || o.Blocknumber == nil {
		var ret int32
		return ret
	}
	return *o.Blocknumber
}

// GetBlocknumberOk returns a tuple with the Blocknumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetBlocknumberOk() (*int32, bool) {
	if o == nil || o.Blocknumber == nil {
		return nil, false
	}
	return o.Blocknumber, true
}

// HasBlocknumber returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasBlocknumber() bool {
	if o != nil && o.Blocknumber != nil {
		return true
	}

	return false
}

// SetBlocknumber gets a reference to the given int32 and assigns it to the Blocknumber field.
func (o *AccountTokenTransfer) SetBlocknumber(v int32) {
	o.Blocknumber = &v
}

// GetTxhash returns the Txhash field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetTxhash() string {
	if o == nil || o.Txhash == nil {
		var ret string
		return ret
	}
	return *o.Txhash
}

// GetTxhashOk returns a tuple with the Txhash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetTxhashOk() (*string, bool) {
	if o == nil || o.Txhash == nil {
		return nil, false
	}
	return o.Txhash, true
}

// HasTxhash returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasTxhash() bool {
	if o != nil && o.Txhash != nil {
		return true
	}

	return false
}

// SetTxhash gets a reference to the given string and assigns it to the Txhash field.
func (o *AccountTokenTransfer) SetTxhash(v string) {
	o.Txhash = &v
}

// GetMint returns the Mint field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetMint() Address {
	if o == nil || o.Mint == nil {
		var ret Address
		return ret
	}
	return *o.Mint
}

// GetMintOk returns a tuple with the Mint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetMintOk() (*Address, bool) {
	if o == nil || o.Mint == nil {
		return nil, false
	}
	return o.Mint, true
}

// HasMint returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasMint() bool {
	if o != nil && o.Mint != nil {
		return true
	}

	return false
}

// SetMint gets a reference to the given Address and assigns it to the Mint field.
func (o *AccountTokenTransfer) SetMint(v Address) {
	o.Mint = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *AccountTokenTransfer) SetValid(v bool) {
	o.Valid = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *AccountTokenTransfer) SetAmount(v int32) {
	o.Amount = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetSource() Address {
	if o == nil || o.Source == nil {
		var ret Address
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetSourceOk() (*Address, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Address and assigns it to the Source field.
func (o *AccountTokenTransfer) SetSource(v Address) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetDestination() Address {
	if o == nil || o.Destination == nil {
		var ret Address
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetDestinationOk() (*Address, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Address and assigns it to the Destination field.
func (o *AccountTokenTransfer) SetDestination(v Address) {
	o.Destination = &v
}

// GetInner returns the Inner field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetInner() bool {
	if o == nil || o.Inner == nil {
		var ret bool
		return ret
	}
	return *o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetInnerOk() (*bool, bool) {
	if o == nil || o.Inner == nil {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasInner() bool {
	if o != nil && o.Inner != nil {
		return true
	}

	return false
}

// SetInner gets a reference to the given bool and assigns it to the Inner field.
func (o *AccountTokenTransfer) SetInner(v bool) {
	o.Inner = &v
}

// GetTxindex returns the Txindex field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetTxindex() int32 {
	if o == nil || o.Txindex == nil {
		var ret int32
		return ret
	}
	return *o.Txindex
}

// GetTxindexOk returns a tuple with the Txindex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetTxindexOk() (*int32, bool) {
	if o == nil || o.Txindex == nil {
		return nil, false
	}
	return o.Txindex, true
}

// HasTxindex returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasTxindex() bool {
	if o != nil && o.Txindex != nil {
		return true
	}

	return false
}

// SetTxindex gets a reference to the given int32 and assigns it to the Txindex field.
func (o *AccountTokenTransfer) SetTxindex(v int32) {
	o.Txindex = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetTimestamp() Timestamp {
	if o == nil || o.Timestamp == nil {
		var ret Timestamp
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetTimestampOk() (*Timestamp, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given Timestamp and assigns it to the Timestamp field.
func (o *AccountTokenTransfer) SetTimestamp(v Timestamp) {
	o.Timestamp = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *AccountTokenTransfer) GetDecimals() int32 {
	if o == nil || o.Decimals == nil {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTokenTransfer) GetDecimalsOk() (*int32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *AccountTokenTransfer) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *AccountTokenTransfer) SetDecimals(v int32) {
	o.Decimals = &v
}

func (o AccountTokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Blocknumber != nil {
		toSerialize["blocknumber"] = o.Blocknumber
	}
	if o.Txhash != nil {
		toSerialize["txhash"] = o.Txhash
	}
	if o.Mint != nil {
		toSerialize["mint"] = o.Mint
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Inner != nil {
		toSerialize["inner"] = o.Inner
	}
	if o.Txindex != nil {
		toSerialize["txindex"] = o.Txindex
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	return json.Marshal(toSerialize)
}

type NullableAccountTokenTransfer struct {
	value *AccountTokenTransfer
	isSet bool
}

func (v NullableAccountTokenTransfer) Get() *AccountTokenTransfer {
	return v.value
}

func (v *NullableAccountTokenTransfer) Set(val *AccountTokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTokenTransfer(val *AccountTokenTransfer) *NullableAccountTokenTransfer {
	return &NullableAccountTokenTransfer{value: val, isSet: true}
}

func (v NullableAccountTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


