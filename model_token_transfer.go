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

// TokenTransfer struct for TokenTransfer
type TokenTransfer struct {
	Id *int32 `json:"id,omitempty"`
	Blocknumber *int32 `json:"blocknumber,omitempty"`
	Transactionhash *string `json:"transactionhash,omitempty"`
	Mint *Address `json:"mint,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	Source *Address `json:"source,omitempty"`
	Destination *Address `json:"destination,omitempty"`
	Decimals *int32 `json:"decimals,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Innerinstruction *bool `json:"innerinstruction,omitempty"`
	Index *int32 `json:"index,omitempty"`
	Owner *Address `json:"owner,omitempty"`
}

// NewTokenTransfer instantiates a new TokenTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenTransfer() *TokenTransfer {
	this := TokenTransfer{}
	return &this
}

// NewTokenTransferWithDefaults instantiates a new TokenTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTransferWithDefaults() *TokenTransfer {
	this := TokenTransfer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenTransfer) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenTransfer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TokenTransfer) SetId(v int32) {
	o.Id = &v
}

// GetBlocknumber returns the Blocknumber field value if set, zero value otherwise.
func (o *TokenTransfer) GetBlocknumber() int32 {
	if o == nil || o.Blocknumber == nil {
		var ret int32
		return ret
	}
	return *o.Blocknumber
}

// GetBlocknumberOk returns a tuple with the Blocknumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetBlocknumberOk() (*int32, bool) {
	if o == nil || o.Blocknumber == nil {
		return nil, false
	}
	return o.Blocknumber, true
}

// HasBlocknumber returns a boolean if a field has been set.
func (o *TokenTransfer) HasBlocknumber() bool {
	if o != nil && o.Blocknumber != nil {
		return true
	}

	return false
}

// SetBlocknumber gets a reference to the given int32 and assigns it to the Blocknumber field.
func (o *TokenTransfer) SetBlocknumber(v int32) {
	o.Blocknumber = &v
}

// GetTransactionhash returns the Transactionhash field value if set, zero value otherwise.
func (o *TokenTransfer) GetTransactionhash() string {
	if o == nil || o.Transactionhash == nil {
		var ret string
		return ret
	}
	return *o.Transactionhash
}

// GetTransactionhashOk returns a tuple with the Transactionhash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetTransactionhashOk() (*string, bool) {
	if o == nil || o.Transactionhash == nil {
		return nil, false
	}
	return o.Transactionhash, true
}

// HasTransactionhash returns a boolean if a field has been set.
func (o *TokenTransfer) HasTransactionhash() bool {
	if o != nil && o.Transactionhash != nil {
		return true
	}

	return false
}

// SetTransactionhash gets a reference to the given string and assigns it to the Transactionhash field.
func (o *TokenTransfer) SetTransactionhash(v string) {
	o.Transactionhash = &v
}

// GetMint returns the Mint field value if set, zero value otherwise.
func (o *TokenTransfer) GetMint() Address {
	if o == nil || o.Mint == nil {
		var ret Address
		return ret
	}
	return *o.Mint
}

// GetMintOk returns a tuple with the Mint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetMintOk() (*Address, bool) {
	if o == nil || o.Mint == nil {
		return nil, false
	}
	return o.Mint, true
}

// HasMint returns a boolean if a field has been set.
func (o *TokenTransfer) HasMint() bool {
	if o != nil && o.Mint != nil {
		return true
	}

	return false
}

// SetMint gets a reference to the given Address and assigns it to the Mint field.
func (o *TokenTransfer) SetMint(v Address) {
	o.Mint = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TokenTransfer) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TokenTransfer) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *TokenTransfer) SetAmount(v int32) {
	o.Amount = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TokenTransfer) GetSource() Address {
	if o == nil || o.Source == nil {
		var ret Address
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetSourceOk() (*Address, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TokenTransfer) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Address and assigns it to the Source field.
func (o *TokenTransfer) SetSource(v Address) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TokenTransfer) GetDestination() Address {
	if o == nil || o.Destination == nil {
		var ret Address
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetDestinationOk() (*Address, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TokenTransfer) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given Address and assigns it to the Destination field.
func (o *TokenTransfer) SetDestination(v Address) {
	o.Destination = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *TokenTransfer) GetDecimals() int32 {
	if o == nil || o.Decimals == nil {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetDecimalsOk() (*int32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenTransfer) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *TokenTransfer) SetDecimals(v int32) {
	o.Decimals = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TokenTransfer) GetTimestamp() Timestamp {
	if o == nil || o.Timestamp == nil {
		var ret Timestamp
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetTimestampOk() (*Timestamp, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TokenTransfer) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given Timestamp and assigns it to the Timestamp field.
func (o *TokenTransfer) SetTimestamp(v Timestamp) {
	o.Timestamp = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *TokenTransfer) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *TokenTransfer) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *TokenTransfer) SetValid(v bool) {
	o.Valid = &v
}

// GetInnerinstruction returns the Innerinstruction field value if set, zero value otherwise.
func (o *TokenTransfer) GetInnerinstruction() bool {
	if o == nil || o.Innerinstruction == nil {
		var ret bool
		return ret
	}
	return *o.Innerinstruction
}

// GetInnerinstructionOk returns a tuple with the Innerinstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetInnerinstructionOk() (*bool, bool) {
	if o == nil || o.Innerinstruction == nil {
		return nil, false
	}
	return o.Innerinstruction, true
}

// HasInnerinstruction returns a boolean if a field has been set.
func (o *TokenTransfer) HasInnerinstruction() bool {
	if o != nil && o.Innerinstruction != nil {
		return true
	}

	return false
}

// SetInnerinstruction gets a reference to the given bool and assigns it to the Innerinstruction field.
func (o *TokenTransfer) SetInnerinstruction(v bool) {
	o.Innerinstruction = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TokenTransfer) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TokenTransfer) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TokenTransfer) SetIndex(v int32) {
	o.Index = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *TokenTransfer) GetOwner() Address {
	if o == nil || o.Owner == nil {
		var ret Address
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenTransfer) GetOwnerOk() (*Address, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *TokenTransfer) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Address and assigns it to the Owner field.
func (o *TokenTransfer) SetOwner(v Address) {
	o.Owner = &v
}

func (o TokenTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Blocknumber != nil {
		toSerialize["blocknumber"] = o.Blocknumber
	}
	if o.Transactionhash != nil {
		toSerialize["transactionhash"] = o.Transactionhash
	}
	if o.Mint != nil {
		toSerialize["mint"] = o.Mint
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
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.Innerinstruction != nil {
		toSerialize["innerinstruction"] = o.Innerinstruction
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableTokenTransfer struct {
	value *TokenTransfer
	isSet bool
}

func (v NullableTokenTransfer) Get() *TokenTransfer {
	return v.value
}

func (v *NullableTokenTransfer) Set(val *TokenTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenTransfer(val *TokenTransfer) *NullableTokenTransfer {
	return &NullableTokenTransfer{value: val, isSet: true}
}

func (v NullableTokenTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


