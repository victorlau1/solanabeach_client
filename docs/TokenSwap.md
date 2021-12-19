# TokenSwap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**Address**](Address.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Poolmint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Tokenamint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Tokenbmint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Tokena** | Pointer to [**Address**](Address.md) |  | [optional] 
**Tokenb** | Pointer to [**Address**](Address.md) |  | [optional] 
**Tokenprogram** | Pointer to [**Address**](Address.md) |  | [optional] 
**Meta** | Pointer to [**TokenSwapMeta**](TokenSwapMeta.md) |  | [optional] 

## Methods

### NewTokenSwap

`func NewTokenSwap() *TokenSwap`

NewTokenSwap instantiates a new TokenSwap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSwapWithDefaults

`func NewTokenSwapWithDefaults() *TokenSwap`

NewTokenSwapWithDefaults instantiates a new TokenSwap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *TokenSwap) GetPubkey() Address`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TokenSwap) GetPubkeyOk() (*Address, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TokenSwap) SetPubkey(v Address)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *TokenSwap) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetName

`func (o *TokenSwap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenSwap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenSwap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenSwap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolmint

`func (o *TokenSwap) GetPoolmint() Address`

GetPoolmint returns the Poolmint field if non-nil, zero value otherwise.

### GetPoolmintOk

`func (o *TokenSwap) GetPoolmintOk() (*Address, bool)`

GetPoolmintOk returns a tuple with the Poolmint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolmint

`func (o *TokenSwap) SetPoolmint(v Address)`

SetPoolmint sets Poolmint field to given value.

### HasPoolmint

`func (o *TokenSwap) HasPoolmint() bool`

HasPoolmint returns a boolean if a field has been set.

### GetTokenamint

`func (o *TokenSwap) GetTokenamint() Address`

GetTokenamint returns the Tokenamint field if non-nil, zero value otherwise.

### GetTokenamintOk

`func (o *TokenSwap) GetTokenamintOk() (*Address, bool)`

GetTokenamintOk returns a tuple with the Tokenamint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenamint

`func (o *TokenSwap) SetTokenamint(v Address)`

SetTokenamint sets Tokenamint field to given value.

### HasTokenamint

`func (o *TokenSwap) HasTokenamint() bool`

HasTokenamint returns a boolean if a field has been set.

### GetTokenbmint

`func (o *TokenSwap) GetTokenbmint() Address`

GetTokenbmint returns the Tokenbmint field if non-nil, zero value otherwise.

### GetTokenbmintOk

`func (o *TokenSwap) GetTokenbmintOk() (*Address, bool)`

GetTokenbmintOk returns a tuple with the Tokenbmint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenbmint

`func (o *TokenSwap) SetTokenbmint(v Address)`

SetTokenbmint sets Tokenbmint field to given value.

### HasTokenbmint

`func (o *TokenSwap) HasTokenbmint() bool`

HasTokenbmint returns a boolean if a field has been set.

### GetTokena

`func (o *TokenSwap) GetTokena() Address`

GetTokena returns the Tokena field if non-nil, zero value otherwise.

### GetTokenaOk

`func (o *TokenSwap) GetTokenaOk() (*Address, bool)`

GetTokenaOk returns a tuple with the Tokena field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokena

`func (o *TokenSwap) SetTokena(v Address)`

SetTokena sets Tokena field to given value.

### HasTokena

`func (o *TokenSwap) HasTokena() bool`

HasTokena returns a boolean if a field has been set.

### GetTokenb

`func (o *TokenSwap) GetTokenb() Address`

GetTokenb returns the Tokenb field if non-nil, zero value otherwise.

### GetTokenbOk

`func (o *TokenSwap) GetTokenbOk() (*Address, bool)`

GetTokenbOk returns a tuple with the Tokenb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenb

`func (o *TokenSwap) SetTokenb(v Address)`

SetTokenb sets Tokenb field to given value.

### HasTokenb

`func (o *TokenSwap) HasTokenb() bool`

HasTokenb returns a boolean if a field has been set.

### GetTokenprogram

`func (o *TokenSwap) GetTokenprogram() Address`

GetTokenprogram returns the Tokenprogram field if non-nil, zero value otherwise.

### GetTokenprogramOk

`func (o *TokenSwap) GetTokenprogramOk() (*Address, bool)`

GetTokenprogramOk returns a tuple with the Tokenprogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenprogram

`func (o *TokenSwap) SetTokenprogram(v Address)`

SetTokenprogram sets Tokenprogram field to given value.

### HasTokenprogram

`func (o *TokenSwap) HasTokenprogram() bool`

HasTokenprogram returns a boolean if a field has been set.

### GetMeta

`func (o *TokenSwap) GetMeta() TokenSwapMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TokenSwap) GetMetaOk() (*TokenSwapMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TokenSwap) SetMeta(v TokenSwapMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TokenSwap) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


