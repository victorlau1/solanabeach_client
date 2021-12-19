# SimpleAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**Address**](Address.md) |  | [optional] 
**Balance** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimpleAccount

`func NewSimpleAccount() *SimpleAccount`

NewSimpleAccount instantiates a new SimpleAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleAccountWithDefaults

`func NewSimpleAccountWithDefaults() *SimpleAccount`

NewSimpleAccountWithDefaults instantiates a new SimpleAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *SimpleAccount) GetPubkey() Address`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *SimpleAccount) GetPubkeyOk() (*Address, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *SimpleAccount) SetPubkey(v Address)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *SimpleAccount) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetBalance

`func (o *SimpleAccount) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SimpleAccount) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SimpleAccount) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *SimpleAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


