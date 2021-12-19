# AccountTokenTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Txhash** | Pointer to **string** |  | [optional] 
**Mint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to [**Address**](Address.md) |  | [optional] 
**Destination** | Pointer to [**Address**](Address.md) |  | [optional] 
**Inner** | Pointer to **bool** |  | [optional] 
**Txindex** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountTokenTransfer

`func NewAccountTokenTransfer() *AccountTokenTransfer`

NewAccountTokenTransfer instantiates a new AccountTokenTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenTransferWithDefaults

`func NewAccountTokenTransferWithDefaults() *AccountTokenTransfer`

NewAccountTokenTransferWithDefaults instantiates a new AccountTokenTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocknumber

`func (o *AccountTokenTransfer) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *AccountTokenTransfer) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *AccountTokenTransfer) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *AccountTokenTransfer) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetTxhash

`func (o *AccountTokenTransfer) GetTxhash() string`

GetTxhash returns the Txhash field if non-nil, zero value otherwise.

### GetTxhashOk

`func (o *AccountTokenTransfer) GetTxhashOk() (*string, bool)`

GetTxhashOk returns a tuple with the Txhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxhash

`func (o *AccountTokenTransfer) SetTxhash(v string)`

SetTxhash sets Txhash field to given value.

### HasTxhash

`func (o *AccountTokenTransfer) HasTxhash() bool`

HasTxhash returns a boolean if a field has been set.

### GetMint

`func (o *AccountTokenTransfer) GetMint() Address`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *AccountTokenTransfer) GetMintOk() (*Address, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *AccountTokenTransfer) SetMint(v Address)`

SetMint sets Mint field to given value.

### HasMint

`func (o *AccountTokenTransfer) HasMint() bool`

HasMint returns a boolean if a field has been set.

### GetValid

`func (o *AccountTokenTransfer) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AccountTokenTransfer) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AccountTokenTransfer) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AccountTokenTransfer) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetAmount

`func (o *AccountTokenTransfer) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountTokenTransfer) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountTokenTransfer) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountTokenTransfer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSource

`func (o *AccountTokenTransfer) GetSource() Address`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccountTokenTransfer) GetSourceOk() (*Address, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccountTokenTransfer) SetSource(v Address)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccountTokenTransfer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *AccountTokenTransfer) GetDestination() Address`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AccountTokenTransfer) GetDestinationOk() (*Address, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AccountTokenTransfer) SetDestination(v Address)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AccountTokenTransfer) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetInner

`func (o *AccountTokenTransfer) GetInner() bool`

GetInner returns the Inner field if non-nil, zero value otherwise.

### GetInnerOk

`func (o *AccountTokenTransfer) GetInnerOk() (*bool, bool)`

GetInnerOk returns a tuple with the Inner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInner

`func (o *AccountTokenTransfer) SetInner(v bool)`

SetInner sets Inner field to given value.

### HasInner

`func (o *AccountTokenTransfer) HasInner() bool`

HasInner returns a boolean if a field has been set.

### GetTxindex

`func (o *AccountTokenTransfer) GetTxindex() int32`

GetTxindex returns the Txindex field if non-nil, zero value otherwise.

### GetTxindexOk

`func (o *AccountTokenTransfer) GetTxindexOk() (*int32, bool)`

GetTxindexOk returns a tuple with the Txindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxindex

`func (o *AccountTokenTransfer) SetTxindex(v int32)`

SetTxindex sets Txindex field to given value.

### HasTxindex

`func (o *AccountTokenTransfer) HasTxindex() bool`

HasTxindex returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountTokenTransfer) GetTimestamp() Timestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountTokenTransfer) GetTimestampOk() (*Timestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountTokenTransfer) SetTimestamp(v Timestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountTokenTransfer) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDecimals

`func (o *AccountTokenTransfer) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AccountTokenTransfer) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AccountTokenTransfer) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *AccountTokenTransfer) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


