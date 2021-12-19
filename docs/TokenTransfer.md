# TokenTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Transactionhash** | Pointer to **string** |  | [optional] 
**Mint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to [**Address**](Address.md) |  | [optional] 
**Destination** | Pointer to [**Address**](Address.md) |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Innerinstruction** | Pointer to **bool** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 

## Methods

### NewTokenTransfer

`func NewTokenTransfer() *TokenTransfer`

NewTokenTransfer instantiates a new TokenTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferWithDefaults

`func NewTokenTransferWithDefaults() *TokenTransfer`

NewTokenTransferWithDefaults instantiates a new TokenTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenTransfer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenTransfer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenTransfer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TokenTransfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlocknumber

`func (o *TokenTransfer) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *TokenTransfer) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *TokenTransfer) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *TokenTransfer) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetTransactionhash

`func (o *TokenTransfer) GetTransactionhash() string`

GetTransactionhash returns the Transactionhash field if non-nil, zero value otherwise.

### GetTransactionhashOk

`func (o *TokenTransfer) GetTransactionhashOk() (*string, bool)`

GetTransactionhashOk returns a tuple with the Transactionhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionhash

`func (o *TokenTransfer) SetTransactionhash(v string)`

SetTransactionhash sets Transactionhash field to given value.

### HasTransactionhash

`func (o *TokenTransfer) HasTransactionhash() bool`

HasTransactionhash returns a boolean if a field has been set.

### GetMint

`func (o *TokenTransfer) GetMint() Address`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *TokenTransfer) GetMintOk() (*Address, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *TokenTransfer) SetMint(v Address)`

SetMint sets Mint field to given value.

### HasMint

`func (o *TokenTransfer) HasMint() bool`

HasMint returns a boolean if a field has been set.

### GetAmount

`func (o *TokenTransfer) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenTransfer) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenTransfer) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TokenTransfer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSource

`func (o *TokenTransfer) GetSource() Address`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenTransfer) GetSourceOk() (*Address, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenTransfer) SetSource(v Address)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenTransfer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *TokenTransfer) GetDestination() Address`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TokenTransfer) GetDestinationOk() (*Address, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TokenTransfer) SetDestination(v Address)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *TokenTransfer) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenTransfer) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenTransfer) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenTransfer) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenTransfer) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetTimestamp

`func (o *TokenTransfer) GetTimestamp() Timestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TokenTransfer) GetTimestampOk() (*Timestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TokenTransfer) SetTimestamp(v Timestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TokenTransfer) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValid

`func (o *TokenTransfer) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *TokenTransfer) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *TokenTransfer) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *TokenTransfer) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetInnerinstruction

`func (o *TokenTransfer) GetInnerinstruction() bool`

GetInnerinstruction returns the Innerinstruction field if non-nil, zero value otherwise.

### GetInnerinstructionOk

`func (o *TokenTransfer) GetInnerinstructionOk() (*bool, bool)`

GetInnerinstructionOk returns a tuple with the Innerinstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerinstruction

`func (o *TokenTransfer) SetInnerinstruction(v bool)`

SetInnerinstruction sets Innerinstruction field to given value.

### HasInnerinstruction

`func (o *TokenTransfer) HasInnerinstruction() bool`

HasInnerinstruction returns a boolean if a field has been set.

### GetIndex

`func (o *TokenTransfer) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TokenTransfer) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TokenTransfer) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TokenTransfer) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetOwner

`func (o *TokenTransfer) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TokenTransfer) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TokenTransfer) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TokenTransfer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


