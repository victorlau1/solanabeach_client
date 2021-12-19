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

// AccountSerumOrder struct for AccountSerumOrder
type AccountSerumOrder struct {
	Id *int32 `json:"id,omitempty"`
	Blocknumber *int32 `json:"blocknumber,omitempty"`
	Transactionhash *string `json:"transactionhash,omitempty"`
	Market *AccountSerumOrderMarket `json:"market,omitempty"`
	Owner *Address `json:"owner,omitempty"`
	Openorders *Address `json:"openorders,omitempty"`
	Side *string `json:"side,omitempty"`
	Type *string `json:"type,omitempty"`
	Pricelimit *int32 `json:"pricelimit,omitempty"`
	Maxquantity *int32 `json:"maxquantity,omitempty"`
	Valid *bool `json:"valid,omitempty"`
	Status *string `json:"status,omitempty"`
	Processedat *int32 `json:"processedat,omitempty"`
	Txindex *int32 `json:"txindex,omitempty"`
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	Ondemand *bool `json:"ondemand,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Qty *float32 `json:"qty,omitempty"`
}

// NewAccountSerumOrder instantiates a new AccountSerumOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSerumOrder() *AccountSerumOrder {
	this := AccountSerumOrder{}
	return &this
}

// NewAccountSerumOrderWithDefaults instantiates a new AccountSerumOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSerumOrderWithDefaults() *AccountSerumOrder {
	this := AccountSerumOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AccountSerumOrder) SetId(v int32) {
	o.Id = &v
}

// GetBlocknumber returns the Blocknumber field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetBlocknumber() int32 {
	if o == nil || o.Blocknumber == nil {
		var ret int32
		return ret
	}
	return *o.Blocknumber
}

// GetBlocknumberOk returns a tuple with the Blocknumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetBlocknumberOk() (*int32, bool) {
	if o == nil || o.Blocknumber == nil {
		return nil, false
	}
	return o.Blocknumber, true
}

// HasBlocknumber returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasBlocknumber() bool {
	if o != nil && o.Blocknumber != nil {
		return true
	}

	return false
}

// SetBlocknumber gets a reference to the given int32 and assigns it to the Blocknumber field.
func (o *AccountSerumOrder) SetBlocknumber(v int32) {
	o.Blocknumber = &v
}

// GetTransactionhash returns the Transactionhash field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetTransactionhash() string {
	if o == nil || o.Transactionhash == nil {
		var ret string
		return ret
	}
	return *o.Transactionhash
}

// GetTransactionhashOk returns a tuple with the Transactionhash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetTransactionhashOk() (*string, bool) {
	if o == nil || o.Transactionhash == nil {
		return nil, false
	}
	return o.Transactionhash, true
}

// HasTransactionhash returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasTransactionhash() bool {
	if o != nil && o.Transactionhash != nil {
		return true
	}

	return false
}

// SetTransactionhash gets a reference to the given string and assigns it to the Transactionhash field.
func (o *AccountSerumOrder) SetTransactionhash(v string) {
	o.Transactionhash = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetMarket() AccountSerumOrderMarket {
	if o == nil || o.Market == nil {
		var ret AccountSerumOrderMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetMarketOk() (*AccountSerumOrderMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AccountSerumOrderMarket and assigns it to the Market field.
func (o *AccountSerumOrder) SetMarket(v AccountSerumOrderMarket) {
	o.Market = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetOwner() Address {
	if o == nil || o.Owner == nil {
		var ret Address
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetOwnerOk() (*Address, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Address and assigns it to the Owner field.
func (o *AccountSerumOrder) SetOwner(v Address) {
	o.Owner = &v
}

// GetOpenorders returns the Openorders field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetOpenorders() Address {
	if o == nil || o.Openorders == nil {
		var ret Address
		return ret
	}
	return *o.Openorders
}

// GetOpenordersOk returns a tuple with the Openorders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetOpenordersOk() (*Address, bool) {
	if o == nil || o.Openorders == nil {
		return nil, false
	}
	return o.Openorders, true
}

// HasOpenorders returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasOpenorders() bool {
	if o != nil && o.Openorders != nil {
		return true
	}

	return false
}

// SetOpenorders gets a reference to the given Address and assigns it to the Openorders field.
func (o *AccountSerumOrder) SetOpenorders(v Address) {
	o.Openorders = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *AccountSerumOrder) SetSide(v string) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountSerumOrder) SetType(v string) {
	o.Type = &v
}

// GetPricelimit returns the Pricelimit field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetPricelimit() int32 {
	if o == nil || o.Pricelimit == nil {
		var ret int32
		return ret
	}
	return *o.Pricelimit
}

// GetPricelimitOk returns a tuple with the Pricelimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetPricelimitOk() (*int32, bool) {
	if o == nil || o.Pricelimit == nil {
		return nil, false
	}
	return o.Pricelimit, true
}

// HasPricelimit returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasPricelimit() bool {
	if o != nil && o.Pricelimit != nil {
		return true
	}

	return false
}

// SetPricelimit gets a reference to the given int32 and assigns it to the Pricelimit field.
func (o *AccountSerumOrder) SetPricelimit(v int32) {
	o.Pricelimit = &v
}

// GetMaxquantity returns the Maxquantity field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetMaxquantity() int32 {
	if o == nil || o.Maxquantity == nil {
		var ret int32
		return ret
	}
	return *o.Maxquantity
}

// GetMaxquantityOk returns a tuple with the Maxquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetMaxquantityOk() (*int32, bool) {
	if o == nil || o.Maxquantity == nil {
		return nil, false
	}
	return o.Maxquantity, true
}

// HasMaxquantity returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasMaxquantity() bool {
	if o != nil && o.Maxquantity != nil {
		return true
	}

	return false
}

// SetMaxquantity gets a reference to the given int32 and assigns it to the Maxquantity field.
func (o *AccountSerumOrder) SetMaxquantity(v int32) {
	o.Maxquantity = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *AccountSerumOrder) SetValid(v bool) {
	o.Valid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountSerumOrder) SetStatus(v string) {
	o.Status = &v
}

// GetProcessedat returns the Processedat field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetProcessedat() int32 {
	if o == nil || o.Processedat == nil {
		var ret int32
		return ret
	}
	return *o.Processedat
}

// GetProcessedatOk returns a tuple with the Processedat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetProcessedatOk() (*int32, bool) {
	if o == nil || o.Processedat == nil {
		return nil, false
	}
	return o.Processedat, true
}

// HasProcessedat returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasProcessedat() bool {
	if o != nil && o.Processedat != nil {
		return true
	}

	return false
}

// SetProcessedat gets a reference to the given int32 and assigns it to the Processedat field.
func (o *AccountSerumOrder) SetProcessedat(v int32) {
	o.Processedat = &v
}

// GetTxindex returns the Txindex field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetTxindex() int32 {
	if o == nil || o.Txindex == nil {
		var ret int32
		return ret
	}
	return *o.Txindex
}

// GetTxindexOk returns a tuple with the Txindex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetTxindexOk() (*int32, bool) {
	if o == nil || o.Txindex == nil {
		return nil, false
	}
	return o.Txindex, true
}

// HasTxindex returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasTxindex() bool {
	if o != nil && o.Txindex != nil {
		return true
	}

	return false
}

// SetTxindex gets a reference to the given int32 and assigns it to the Txindex field.
func (o *AccountSerumOrder) SetTxindex(v int32) {
	o.Txindex = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetTimestamp() Timestamp {
	if o == nil || o.Timestamp == nil {
		var ret Timestamp
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetTimestampOk() (*Timestamp, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given Timestamp and assigns it to the Timestamp field.
func (o *AccountSerumOrder) SetTimestamp(v Timestamp) {
	o.Timestamp = &v
}

// GetOndemand returns the Ondemand field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetOndemand() bool {
	if o == nil || o.Ondemand == nil {
		var ret bool
		return ret
	}
	return *o.Ondemand
}

// GetOndemandOk returns a tuple with the Ondemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetOndemandOk() (*bool, bool) {
	if o == nil || o.Ondemand == nil {
		return nil, false
	}
	return o.Ondemand, true
}

// HasOndemand returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasOndemand() bool {
	if o != nil && o.Ondemand != nil {
		return true
	}

	return false
}

// SetOndemand gets a reference to the given bool and assigns it to the Ondemand field.
func (o *AccountSerumOrder) SetOndemand(v bool) {
	o.Ondemand = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *AccountSerumOrder) SetPrice(v float32) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *AccountSerumOrder) GetQty() float32 {
	if o == nil || o.Qty == nil {
		var ret float32
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSerumOrder) GetQtyOk() (*float32, bool) {
	if o == nil || o.Qty == nil {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *AccountSerumOrder) HasQty() bool {
	if o != nil && o.Qty != nil {
		return true
	}

	return false
}

// SetQty gets a reference to the given float32 and assigns it to the Qty field.
func (o *AccountSerumOrder) SetQty(v float32) {
	o.Qty = &v
}

func (o AccountSerumOrder) MarshalJSON() ([]byte, error) {
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
	if o.Openorders != nil {
		toSerialize["openorders"] = o.Openorders
	}
	if o.Side != nil {
		toSerialize["side"] = o.Side
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Pricelimit != nil {
		toSerialize["pricelimit"] = o.Pricelimit
	}
	if o.Maxquantity != nil {
		toSerialize["maxquantity"] = o.Maxquantity
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Processedat != nil {
		toSerialize["processedat"] = o.Processedat
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
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Qty != nil {
		toSerialize["qty"] = o.Qty
	}
	return json.Marshal(toSerialize)
}

type NullableAccountSerumOrder struct {
	value *AccountSerumOrder
	isSet bool
}

func (v NullableAccountSerumOrder) Get() *AccountSerumOrder {
	return v.value
}

func (v *NullableAccountSerumOrder) Set(val *AccountSerumOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSerumOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSerumOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSerumOrder(val *AccountSerumOrder) *NullableAccountSerumOrder {
	return &NullableAccountSerumOrder{value: val, isSet: true}
}

func (v NullableAccountSerumOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSerumOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

