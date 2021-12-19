# AccountSwapInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Transactionhash** | Pointer to **string** |  | [optional] 
**Tokenswap** | Pointer to [**AccountSwapInstructionTokenswap**](AccountSwapInstructionTokenswap.md) |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 
**Instruction** | Pointer to **string** |  | [optional] 
**Instructiondata** | Pointer to **map[string]interface{}** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Txindex** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountSwapInstruction

`func NewAccountSwapInstruction() *AccountSwapInstruction`

NewAccountSwapInstruction instantiates a new AccountSwapInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSwapInstructionWithDefaults

`func NewAccountSwapInstructionWithDefaults() *AccountSwapInstruction`

NewAccountSwapInstructionWithDefaults instantiates a new AccountSwapInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSwapInstruction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSwapInstruction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSwapInstruction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSwapInstruction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlocknumber

`func (o *AccountSwapInstruction) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *AccountSwapInstruction) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *AccountSwapInstruction) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *AccountSwapInstruction) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetTransactionhash

`func (o *AccountSwapInstruction) GetTransactionhash() string`

GetTransactionhash returns the Transactionhash field if non-nil, zero value otherwise.

### GetTransactionhashOk

`func (o *AccountSwapInstruction) GetTransactionhashOk() (*string, bool)`

GetTransactionhashOk returns a tuple with the Transactionhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionhash

`func (o *AccountSwapInstruction) SetTransactionhash(v string)`

SetTransactionhash sets Transactionhash field to given value.

### HasTransactionhash

`func (o *AccountSwapInstruction) HasTransactionhash() bool`

HasTransactionhash returns a boolean if a field has been set.

### GetTokenswap

`func (o *AccountSwapInstruction) GetTokenswap() AccountSwapInstructionTokenswap`

GetTokenswap returns the Tokenswap field if non-nil, zero value otherwise.

### GetTokenswapOk

`func (o *AccountSwapInstruction) GetTokenswapOk() (*AccountSwapInstructionTokenswap, bool)`

GetTokenswapOk returns a tuple with the Tokenswap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenswap

`func (o *AccountSwapInstruction) SetTokenswap(v AccountSwapInstructionTokenswap)`

SetTokenswap sets Tokenswap field to given value.

### HasTokenswap

`func (o *AccountSwapInstruction) HasTokenswap() bool`

HasTokenswap returns a boolean if a field has been set.

### GetOwner

`func (o *AccountSwapInstruction) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountSwapInstruction) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountSwapInstruction) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountSwapInstruction) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetInstruction

`func (o *AccountSwapInstruction) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AccountSwapInstruction) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AccountSwapInstruction) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *AccountSwapInstruction) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetInstructiondata

`func (o *AccountSwapInstruction) GetInstructiondata() map[string]interface{}`

GetInstructiondata returns the Instructiondata field if non-nil, zero value otherwise.

### GetInstructiondataOk

`func (o *AccountSwapInstruction) GetInstructiondataOk() (*map[string]interface{}, bool)`

GetInstructiondataOk returns a tuple with the Instructiondata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructiondata

`func (o *AccountSwapInstruction) SetInstructiondata(v map[string]interface{})`

SetInstructiondata sets Instructiondata field to given value.

### HasInstructiondata

`func (o *AccountSwapInstruction) HasInstructiondata() bool`

HasInstructiondata returns a boolean if a field has been set.

### GetValid

`func (o *AccountSwapInstruction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AccountSwapInstruction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AccountSwapInstruction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AccountSwapInstruction) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetIndex

`func (o *AccountSwapInstruction) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AccountSwapInstruction) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AccountSwapInstruction) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AccountSwapInstruction) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetTxindex

`func (o *AccountSwapInstruction) GetTxindex() int32`

GetTxindex returns the Txindex field if non-nil, zero value otherwise.

### GetTxindexOk

`func (o *AccountSwapInstruction) GetTxindexOk() (*int32, bool)`

GetTxindexOk returns a tuple with the Txindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxindex

`func (o *AccountSwapInstruction) SetTxindex(v int32)`

SetTxindex sets Txindex field to given value.

### HasTxindex

`func (o *AccountSwapInstruction) HasTxindex() bool`

HasTxindex returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountSwapInstruction) GetTimestamp() Timestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountSwapInstruction) GetTimestampOk() (*Timestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountSwapInstruction) SetTimestamp(v Timestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountSwapInstruction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOndemand

`func (o *AccountSwapInstruction) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *AccountSwapInstruction) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *AccountSwapInstruction) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *AccountSwapInstruction) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


