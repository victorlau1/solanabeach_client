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

// AccountSerumInstruction struct for AccountSerumInstruction
type AccountSerumInstruction struct {
	Id *int32 `json:"id,omitempty"`
	Blocknumber *int32 `json:"blocknumber,omitempty"`
	Transactionhash *string `json:"transactionhash,omitempty"`
	Market *AccountSerumInstructionMarket `json:"market,omitempty"`
	Owner *Address `json:"owner,omitempty"`
	Instruction *string `json:"instruction,omitempty"`
	Instructiondata *map[string]interface{} `json:"instructiondata,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Index *int32 `json:"index,omitempty"`
	Txindex *int32 `json:"txindex,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	Ondemand *bool `json:"ondemand,omitempty"`
}

// NewAccountSerumInstruction instantiates a new AccountSerumInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSerumInstruction() *AccountSerumInstruction {
	this := AccountSerumInstruction{}
	return &this
}

// NewAccountSerumInstructionWithDefaults instantiates a new AccountSerumInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSerumInstructionWithDefaults() *AccountSerumInstruction {
	this := AccountSerumInstruction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AccountSerumInstruction) SetId(v int32) {
	o.Id = &v
}

// GetBlocknumber returns the Blocknumber field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetBlocknumber() int32 {
	if o == nil || o.Blocknumber == nil {
		var ret int32
		return ret
	}
	return *o.Blocknumber
}

// GetBlocknumberOk returns a tuple with the Blocknumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetBlocknumberOk() (*int32, bool) {
	if o == nil || o.Blocknumber == nil {
		return nil, false
	}
	return o.Blocknumber, true
}

// HasBlocknumber returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasBlocknumber() bool {
	if o != nil && o.Blocknumber != nil {
		return true
	}

	return false
}

// SetBlocknumber gets a reference to the given int32 and assigns it to the Blocknumber field.
func (o *AccountSerumInstruction) SetBlocknumber(v int32) {
	o.Blocknumber = &v
}

// GetTransactionhash returns the Transactionhash field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetTransactionhash() string {
	if o == nil || o.Transactionhash == nil {
		var ret string
		return ret
	}
	return *o.Transactionhash
}

// GetTransactionhashOk returns a tuple with the Transactionhash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetTransactionhashOk() (*string, bool) {
	if o == nil || o.Transactionhash == nil {
		return nil, false
	}
	return o.Transactionhash, true
}

// HasTransactionhash returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasTransactionhash() bool {
	if o != nil && o.Transactionhash != nil {
		return true
	}

	return false
}

// SetTransactionhash gets a reference to the given string and assigns it to the Transactionhash field.
func (o *AccountSerumInstruction) SetTransactionhash(v string) {
	o.Transactionhash = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetMarket() AccountSerumInstructionMarket {
	if o == nil || o.Market == nil {
		var ret AccountSerumInstructionMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetMarketOk() (*AccountSerumInstructionMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AccountSerumInstructionMarket and assigns it to the Market field.
func (o *AccountSerumInstruction) SetMarket(v AccountSerumInstructionMarket) {
	o.Market = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetOwner() Address {
	if o == nil || o.Owner == nil {
		var ret Address
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetOwnerOk() (*Address, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Address and assigns it to the Owner field.
func (o *AccountSerumInstruction) SetOwner(v Address) {
	o.Owner = &v
}

// GetInstruction returns the Instruction field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetInstruction() string {
	if o == nil || o.Instruction == nil {
		var ret string
		return ret
	}
	return *o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetInstructionOk() (*string, bool) {
	if o == nil || o.Instruction == nil {
		return nil, false
	}
	return o.Instruction, true
}

// HasInstruction returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasInstruction() bool {
	if o != nil && o.Instruction != nil {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given string and assigns it to the Instruction field.
func (o *AccountSerumInstruction) SetInstruction(v string) {
	o.Instruction = &v
}

// GetInstructiondata returns the Instructiondata field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetInstructiondata() map[string]interface{} {
	if o == nil || o.Instructiondata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Instructiondata
}

// GetInstructiondataOk returns a tuple with the Instructiondata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetInstructiondataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Instructiondata == nil {
		return nil, false
	}
	return o.Instructiondata, true
}

// HasInstructiondata returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasInstructiondata() bool {
	if o != nil && o.Instructiondata != nil {
		return true
	}

	return false
}

// SetInstructiondata gets a reference to the given map[string]interface{} and assigns it to the Instructiondata field.
func (o *AccountSerumInstruction) SetInstructiondata(v map[string]interface{}) {
	o.Instructiondata = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *AccountSerumInstruction) SetValid(v bool) {
	o.Valid = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *AccountSerumInstruction) SetIndex(v int32) {
	o.Index = &v
}

// GetTxindex returns the Txindex field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetTxindex() int32 {
	if o == nil || o.Txindex == nil {
		var ret int32
		return ret
	}
	return *o.Txindex
}

// GetTxindexOk returns a tuple with the Txindex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetTxindexOk() (*int32, bool) {
	if o == nil || o.Txindex == nil {
		return nil, false
	}
	return o.Txindex, true
}

// HasTxindex returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasTxindex() bool {
	if o != nil && o.Txindex != nil {
		return true
	}

	return false
}

// SetTxindex gets a reference to the given int32 and assigns it to the Txindex field.
func (o *AccountSerumInstruction) SetTxindex(v int32) {
	o.Txindex = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetTimestamp() Timestamp {
	if o == nil || o.Timestamp == nil {
		var ret Timestamp
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetTimestampOk() (*Timestamp, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given Timestamp and assigns it to the Timestamp field.
func (o *AccountSerumInstruction) SetTimestamp(v Timestamp) {
	o.Timestamp = &v
}

// GetOndemand returns the Ondemand field value if set, zero value otherwise.
func (o *AccountSerumInstruction) GetOndemand() bool {
	if o == nil || o.Ondemand == nil {
		var ret bool
		return ret
	}
	return *o.Ondemand
}

// GetOndemandOk returns a tuple with the Ondemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumInstruction) GetOndemandOk() (*bool, bool) {
	if o == nil || o.Ondemand == nil {
		return nil, false
	}
	return o.Ondemand, true
}

// HasOndemand returns a boolean if a field has been set.
func (o *AccountSerumInstruction) HasOndemand() bool {
	if o != nil && o.Ondemand != nil {
		return true
	}

	return false
}

// SetOndemand gets a reference to the given bool and assigns it to the Ondemand field.
func (o *AccountSerumInstruction) SetOndemand(v bool) {
	o.Ondemand = &v
}

func (o AccountSerumInstruction) MarshalJSON() ([]byte, error) {
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
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Instruction != nil {
		toSerialize["instruction"] = o.Instruction
	}
	if o.Instructiondata != nil {
		toSerialize["instructiondata"] = o.Instructiondata
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Txindex != nil {
		toSerialize["txindex"] = o.Txindex
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Ondemand != nil {
		toSerialize["ondemand"] = o.Ondemand
	}
	return json.Marshal(toSerialize)
}

type NullableAccountSerumInstruction struct {
	value *AccountSerumInstruction
	isSet bool
}

func (v NullableAccountSerumInstruction) Get() *AccountSerumInstruction {
	return v.value
}

func (v *NullableAccountSerumInstruction) Set(val *AccountSerumInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSerumInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSerumInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSerumInstruction(val *AccountSerumInstruction) *NullableAccountSerumInstruction {
	return &NullableAccountSerumInstruction{value: val, isSet: true}
}

func (v NullableAccountSerumInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSerumInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

