# AccountSerumOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Transactionhash** | Pointer to **string** |  | [optional] 
**Market** | Pointer to [**AccountSerumOrderMarket**](AccountSerumOrderMarket.md) |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 
**Openorders** | Pointer to [**Address**](Address.md) |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Pricelimit** | Pointer to **int32** |  | [optional] 
**Maxquantity** | Pointer to **int32** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Processedat** | Pointer to **int32** |  | [optional] 
**Txindex** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Qty** | Pointer to **float32** |  | [optional] 

## Methods

### NewAccountSerumOrder

`func NewAccountSerumOrder() *AccountSerumOrder`

NewAccountSerumOrder instantiates a new AccountSerumOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSerumOrderWithDefaults

`func NewAccountSerumOrderWithDefaults() *AccountSerumOrder`

NewAccountSerumOrderWithDefaults instantiates a new AccountSerumOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSerumOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSerumOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSerumOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSerumOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlocknumber

`func (o *AccountSerumOrder) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *AccountSerumOrder) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *AccountSerumOrder) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *AccountSerumOrder) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetTransactionhash

`func (o *AccountSerumOrder) GetTransactionhash() string`

GetTransactionhash returns the Transactionhash field if non-nil, zero value otherwise.

### GetTransactionhashOk

`func (o *AccountSerumOrder) GetTransactionhashOk() (*string, bool)`

GetTransactionhashOk returns a tuple with the Transactionhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionhash

`func (o *AccountSerumOrder) SetTransactionhash(v string)`

SetTransactionhash sets Transactionhash field to given value.

### HasTransactionhash

`func (o *AccountSerumOrder) HasTransactionhash() bool`

HasTransactionhash returns a boolean if a field has been set.

### GetMarket

`func (o *AccountSerumOrder) GetMarket() AccountSerumOrderMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AccountSerumOrder) GetMarketOk() (*AccountSerumOrderMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AccountSerumOrder) SetMarket(v AccountSerumOrderMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AccountSerumOrder) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOwner

`func (o *AccountSerumOrder) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountSerumOrder) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountSerumOrder) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountSerumOrder) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOpenorders

`func (o *AccountSerumOrder) GetOpenorders() Address`

GetOpenorders returns the Openorders field if non-nil, zero value otherwise.

### GetOpenordersOk

`func (o *AccountSerumOrder) GetOpenordersOk() (*Address, bool)`

GetOpenordersOk returns a tuple with the Openorders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenorders

`func (o *AccountSerumOrder) SetOpenorders(v Address)`

SetOpenorders sets Openorders field to given value.

### HasOpenorders

`func (o *AccountSerumOrder) HasOpenorders() bool`

HasOpenorders returns a boolean if a field has been set.

### GetSide

`func (o *AccountSerumOrder) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *AccountSerumOrder) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *AccountSerumOrder) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *AccountSerumOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *AccountSerumOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountSerumOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountSerumOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountSerumOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPricelimit

`func (o *AccountSerumOrder) GetPricelimit() int32`

GetPricelimit returns the Pricelimit field if non-nil, zero value otherwise.

### GetPricelimitOk

`func (o *AccountSerumOrder) GetPricelimitOk() (*int32, bool)`

GetPricelimitOk returns a tuple with the Pricelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricelimit

`func (o *AccountSerumOrder) SetPricelimit(v int32)`

SetPricelimit sets Pricelimit field to given value.

### HasPricelimit

`func (o *AccountSerumOrder) HasPricelimit() bool`

HasPricelimit returns a boolean if a field has been set.

### GetMaxquantity

`func (o *AccountSerumOrder) GetMaxquantity() int32`

GetMaxquantity returns the Maxquantity field if non-nil, zero value otherwise.

### GetMaxquantityOk

`func (o *AccountSerumOrder) GetMaxquantityOk() (*int32, bool)`

GetMaxquantityOk returns a tuple with the Maxquantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxquantity

`func (o *AccountSerumOrder) SetMaxquantity(v int32)`

SetMaxquantity sets Maxquantity field to given value.

### HasMaxquantity

`func (o *AccountSerumOrder) HasMaxquantity() bool`

HasMaxquantity returns a boolean if a field has been set.

### GetValid

`func (o *AccountSerumOrder) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AccountSerumOrder) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AccountSerumOrder) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AccountSerumOrder) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetStatus

`func (o *AccountSerumOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountSerumOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountSerumOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountSerumOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProcessedat

`func (o *AccountSerumOrder) GetProcessedat() int32`

GetProcessedat returns the Processedat field if non-nil, zero value otherwise.

### GetProcessedatOk

`func (o *AccountSerumOrder) GetProcessedatOk() (*int32, bool)`

GetProcessedatOk returns a tuple with the Processedat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedat

`func (o *AccountSerumOrder) SetProcessedat(v int32)`

SetProcessedat sets Processedat field to given value.

### HasProcessedat

`func (o *AccountSerumOrder) HasProcessedat() bool`

HasProcessedat returns a boolean if a field has been set.

### GetTxindex

`func (o *AccountSerumOrder) GetTxindex() int32`

GetTxindex returns the Txindex field if non-nil, zero value otherwise.

### GetTxindexOk

`func (o *AccountSerumOrder) GetTxindexOk() (*int32, bool)`

GetTxindexOk returns a tuple with the Txindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxindex

`func (o *AccountSerumOrder) SetTxindex(v int32)`

SetTxindex sets Txindex field to given value.

### HasTxindex

`func (o *AccountSerumOrder) HasTxindex() bool`

HasTxindex returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountSerumOrder) GetTimestamp() Timestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountSerumOrder) GetTimestampOk() (*Timestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountSerumOrder) SetTimestamp(v Timestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountSerumOrder) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOndemand

`func (o *AccountSerumOrder) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *AccountSerumOrder) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *AccountSerumOrder) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *AccountSerumOrder) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.

### GetPrice

`func (o *AccountSerumOrder) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AccountSerumOrder) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AccountSerumOrder) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AccountSerumOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *AccountSerumOrder) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *AccountSerumOrder) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *AccountSerumOrder) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *AccountSerumOrder) HasQty() bool`

HasQty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


