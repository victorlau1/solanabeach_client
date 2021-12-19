# TokenHolderMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegate** | Pointer to [**Address**](Address.md) |  | [optional] 
**Native** | Pointer to **bool** |  | [optional] 
**DelegatedAmount** | Pointer to **int32** |  | [optional] 
**CloseAuthority** | Pointer to [**Address**](Address.md) |  | [optional] 

## Methods

### NewTokenHolderMeta

`func NewTokenHolderMeta() *TokenHolderMeta`

NewTokenHolderMeta instantiates a new TokenHolderMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHolderMetaWithDefaults

`func NewTokenHolderMetaWithDefaults() *TokenHolderMeta`

NewTokenHolderMetaWithDefaults instantiates a new TokenHolderMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegate

`func (o *TokenHolderMeta) GetDelegate() Address`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TokenHolderMeta) GetDelegateOk() (*Address, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TokenHolderMeta) SetDelegate(v Address)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *TokenHolderMeta) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetNative

`func (o *TokenHolderMeta) GetNative() bool`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *TokenHolderMeta) GetNativeOk() (*bool, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *TokenHolderMeta) SetNative(v bool)`

SetNative sets Native field to given value.

### HasNative

`func (o *TokenHolderMeta) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetDelegatedAmount

`func (o *TokenHolderMeta) GetDelegatedAmount() int32`

GetDelegatedAmount returns the DelegatedAmount field if non-nil, zero value otherwise.

### GetDelegatedAmountOk

`func (o *TokenHolderMeta) GetDelegatedAmountOk() (*int32, bool)`

GetDelegatedAmountOk returns a tuple with the DelegatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedAmount

`func (o *TokenHolderMeta) SetDelegatedAmount(v int32)`

SetDelegatedAmount sets DelegatedAmount field to given value.

### HasDelegatedAmount

`func (o *TokenHolderMeta) HasDelegatedAmount() bool`

HasDelegatedAmount returns a boolean if a field has been set.

### GetCloseAuthority

`func (o *TokenHolderMeta) GetCloseAuthority() Address`

GetCloseAuthority returns the CloseAuthority field if non-nil, zero value otherwise.

### GetCloseAuthorityOk

`func (o *TokenHolderMeta) GetCloseAuthorityOk() (*Address, bool)`

GetCloseAuthorityOk returns a tuple with the CloseAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAuthority

`func (o *TokenHolderMeta) SetCloseAuthority(v Address)`

SetCloseAuthority sets CloseAuthority field to given value.

### HasCloseAuthority

`func (o *TokenHolderMeta) HasCloseAuthority() bool`

HasCloseAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


