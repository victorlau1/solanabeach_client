# AccountSerumInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Transactionhash** | Pointer to **string** |  | [optional] 
**Market** | Pointer to [**AccountSerumInstructionMarket**](AccountSerumInstructionMarket.md) |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 
**Instruction** | Pointer to **string** |  | [optional] 
**Instructiondata** | Pointer to **map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Txindex** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountSerumInstruction

`func NewAccountSerumInstruction() *AccountSerumInstruction`

NewAccountSerumInstruction instantiates a new AccountSerumInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSerumInstructionWithDefaults

`func NewAccountSerumInstructionWithDefaults() *AccountSerumInstruction`

NewAccountSerumInstructionWithDefaults instantiates a new AccountSerumInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSerumInstruction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSerumInstruction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSerumInstruction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSerumInstruction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlocknumber

`func (o *AccountSerumInstruction) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *AccountSerumInstruction) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *AccountSerumInstruction) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *AccountSerumInstruction) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetTransactionhash

`func (o *AccountSerumInstruction) GetTransactionhash() string`

GetTransactionhash returns the Transactionhash field if non-nil, zero value otherwise.

### GetTransactionhashOk

`func (o *AccountSerumInstruction) GetTransactionhashOk() (*string, bool)`

GetTransactionhashOk returns a tuple with the Transactionhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionhash

`func (o *AccountSerumInstruction) SetTransactionhash(v string)`

SetTransactionhash sets Transactionhash field to given value.

### HasTransactionhash

`func (o *AccountSerumInstruction) HasTransactionhash() bool`

HasTransactionhash returns a boolean if a field has been set.

### GetMarket

`func (o *AccountSerumInstruction) GetMarket() AccountSerumInstructionMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AccountSerumInstruction) GetMarketOk() (*AccountSerumInstructionMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AccountSerumInstruction) SetMarket(v AccountSerumInstructionMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *AccountSerumInstruction) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOwner

`func (o *AccountSerumInstruction) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountSerumInstruction) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountSerumInstruction) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountSerumInstruction) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetInstruction

`func (o *AccountSerumInstruction) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AccountSerumInstruction) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AccountSerumInstruction) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *AccountSerumInstruction) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetInstructiondata

`func (o *AccountSerumInstruction) GetInstructiondata() map[string]interface{}`

GetInstructiondata returns the Instructiondata field if non-nil, zero value otherwise.

### GetInstructiondataOk

`func (o *AccountSerumInstruction) GetInstructiondataOk() (*map[string]interface{}, bool)`

GetInstructiondataOk returns a tuple with the Instructiondata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructiondata

`func (o *AccountSerumInstruction) SetInstructiondata(v map[string]interface{})`

SetInstructiondata sets Instructiondata field to given value.

### HasInstructiondata

`func (o *AccountSerumInstruction) HasInstructiondata() bool`

HasInstructiondata returns a boolean if a field has been set.

### GetValid

`func (o *AccountSerumInstruction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AccountSerumInstruction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AccountSerumInstruction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AccountSerumInstruction) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetIndex

`func (o *AccountSerumInstruction) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AccountSerumInstruction) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AccountSerumInstruction) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AccountSerumInstruction) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetTxindex

`func (o *AccountSerumInstruction) GetTxindex() int32`

GetTxindex returns the Txindex field if non-nil, zero value otherwise.

### GetTxindexOk

`func (o *AccountSerumInstruction) GetTxindexOk() (*int32, bool)`

GetTxindexOk returns a tuple with the Txindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxindex

`func (o *AccountSerumInstruction) SetTxindex(v int32)`

SetTxindex sets Txindex field to given value.

### HasTxindex

`func (o *AccountSerumInstruction) HasTxindex() bool`

HasTxindex returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountSerumInstruction) GetTimestamp() Timestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountSerumInstruction) GetTimestampOk() (*Timestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountSerumInstruction) SetTimestamp(v Timestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountSerumInstruction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOndemand

`func (o *AccountSerumInstruction) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *AccountSerumInstruction) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *AccountSerumInstruction) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *AccountSerumInstruction) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


