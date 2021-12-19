# StakeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to **string** |  | [optional] 
**Lamports** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**StakeAccountData**](StakeAccountData.md) |  | [optional] 

## Methods

### NewStakeAccount

`func NewStakeAccount() *StakeAccount`

NewStakeAccount instantiates a new StakeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeAccountWithDefaults

`func NewStakeAccountWithDefaults() *StakeAccount`

NewStakeAccountWithDefaults instantiates a new StakeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *StakeAccount) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *StakeAccount) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *StakeAccount) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *StakeAccount) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetLamports

`func (o *StakeAccount) GetLamports() int64`

GetLamports returns the Lamports field if non-nil, zero value otherwise.

### GetLamportsOk

`func (o *StakeAccount) GetLamportsOk() (*int64, bool)`

GetLamportsOk returns a tuple with the Lamports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLamports

`func (o *StakeAccount) SetLamports(v int64)`

SetLamports sets Lamports field to given value.

### HasLamports

`func (o *StakeAccount) HasLamports() bool`

HasLamports returns a boolean if a field has been set.

### GetData

`func (o *StakeAccount) GetData() StakeAccountData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StakeAccount) GetDataOk() (*StakeAccountData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StakeAccount) SetData(v StakeAccountData)`

SetData sets Data field to given value.

### HasData

`func (o *StakeAccount) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


