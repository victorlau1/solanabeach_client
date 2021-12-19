# AccountToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountToken

`func NewAccountToken() *AccountToken`

NewAccountToken instantiates a new AccountToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenWithDefaults

`func NewAccountTokenWithDefaults() *AccountToken`

NewAccountTokenWithDefaults instantiates a new AccountToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMint

`func (o *AccountToken) GetMint() Address`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *AccountToken) GetMintOk() (*Address, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *AccountToken) SetMint(v Address)`

SetMint sets Mint field to given value.

### HasMint

`func (o *AccountToken) HasMint() bool`

HasMint returns a boolean if a field has been set.

### GetAmount

`func (o *AccountToken) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountToken) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountToken) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountToken) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAddress

`func (o *AccountToken) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountToken) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountToken) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountToken) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDecimals

`func (o *AccountToken) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AccountToken) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AccountToken) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *AccountToken) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


