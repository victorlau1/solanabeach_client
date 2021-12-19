# TokenHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**Address**](Address.md) |  | [optional] 
**Lamports** | Pointer to **int64** |  | [optional] 
**Mint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**TokenHolderMeta**](TokenHolderMeta.md) |  | [optional] 

## Methods

### NewTokenHolder

`func NewTokenHolder() *TokenHolder`

NewTokenHolder instantiates a new TokenHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHolderWithDefaults

`func NewTokenHolderWithDefaults() *TokenHolder`

NewTokenHolderWithDefaults instantiates a new TokenHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *TokenHolder) GetPubkey() Address`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *TokenHolder) GetPubkeyOk() (*Address, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *TokenHolder) SetPubkey(v Address)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *TokenHolder) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetLamports

`func (o *TokenHolder) GetLamports() int64`

GetLamports returns the Lamports field if non-nil, zero value otherwise.

### GetLamportsOk

`func (o *TokenHolder) GetLamportsOk() (*int64, bool)`

GetLamportsOk returns a tuple with the Lamports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLamports

`func (o *TokenHolder) SetLamports(v int64)`

SetLamports sets Lamports field to given value.

### HasLamports

`func (o *TokenHolder) HasLamports() bool`

HasLamports returns a boolean if a field has been set.

### GetMint

`func (o *TokenHolder) GetMint() Address`

GetMint returns the Mint field if non-nil, zero value otherwise.

### GetMintOk

`func (o *TokenHolder) GetMintOk() (*Address, bool)`

GetMintOk returns a tuple with the Mint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMint

`func (o *TokenHolder) SetMint(v Address)`

SetMint sets Mint field to given value.

### HasMint

`func (o *TokenHolder) HasMint() bool`

HasMint returns a boolean if a field has been set.

### GetOwner

`func (o *TokenHolder) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TokenHolder) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TokenHolder) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TokenHolder) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAmount

`func (o *TokenHolder) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenHolder) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenHolder) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TokenHolder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetState

`func (o *TokenHolder) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TokenHolder) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TokenHolder) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *TokenHolder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMeta

`func (o *TokenHolder) GetMeta() TokenHolderMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TokenHolder) GetMetaOk() (*TokenHolderMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TokenHolder) SetMeta(v TokenHolderMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TokenHolder) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


