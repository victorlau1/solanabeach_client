# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Cmcid** | Pointer to **int32** |  | [optional] 
**Pricedata** | Pointer to [**TokenPricedata**](TokenPricedata.md) |  | [optional] 
**Pubkey** | Pointer to **string** |  | [optional] 
**Lamports** | Pointer to **int64** |  | [optional] 
**Supply** | Pointer to **int32** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Initialized** | Pointer to **bool** |  | [optional] 
**Holders** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to [**TokenMeta**](TokenMeta.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Token) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Token) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Token) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Token) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTicker

`func (o *Token) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Token) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Token) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *Token) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetLogo

`func (o *Token) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Token) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Token) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Token) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCmcid

`func (o *Token) GetCmcid() int32`

GetCmcid returns the Cmcid field if non-nil, zero value otherwise.

### GetCmcidOk

`func (o *Token) GetCmcidOk() (*int32, bool)`

GetCmcidOk returns a tuple with the Cmcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmcid

`func (o *Token) SetCmcid(v int32)`

SetCmcid sets Cmcid field to given value.

### HasCmcid

`func (o *Token) HasCmcid() bool`

HasCmcid returns a boolean if a field has been set.

### GetPricedata

`func (o *Token) GetPricedata() TokenPricedata`

GetPricedata returns the Pricedata field if non-nil, zero value otherwise.

### GetPricedataOk

`func (o *Token) GetPricedataOk() (*TokenPricedata, bool)`

GetPricedataOk returns a tuple with the Pricedata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricedata

`func (o *Token) SetPricedata(v TokenPricedata)`

SetPricedata sets Pricedata field to given value.

### HasPricedata

`func (o *Token) HasPricedata() bool`

HasPricedata returns a boolean if a field has been set.

### GetPubkey

`func (o *Token) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Token) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Token) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *Token) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetLamports

`func (o *Token) GetLamports() int64`

GetLamports returns the Lamports field if non-nil, zero value otherwise.

### GetLamportsOk

`func (o *Token) GetLamportsOk() (*int64, bool)`

GetLamportsOk returns a tuple with the Lamports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLamports

`func (o *Token) SetLamports(v int64)`

SetLamports sets Lamports field to given value.

### HasLamports

`func (o *Token) HasLamports() bool`

HasLamports returns a boolean if a field has been set.

### GetSupply

`func (o *Token) GetSupply() int32`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *Token) GetSupplyOk() (*int32, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *Token) SetSupply(v int32)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *Token) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetDecimals

`func (o *Token) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *Token) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *Token) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *Token) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetInitialized

`func (o *Token) GetInitialized() bool`

GetInitialized returns the Initialized field if non-nil, zero value otherwise.

### GetInitializedOk

`func (o *Token) GetInitializedOk() (*bool, bool)`

GetInitializedOk returns a tuple with the Initialized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialized

`func (o *Token) SetInitialized(v bool)`

SetInitialized sets Initialized field to given value.

### HasInitialized

`func (o *Token) HasInitialized() bool`

HasInitialized returns a boolean if a field has been set.

### GetHolders

`func (o *Token) GetHolders() int32`

GetHolders returns the Holders field if non-nil, zero value otherwise.

### GetHoldersOk

`func (o *Token) GetHoldersOk() (*int32, bool)`

GetHoldersOk returns a tuple with the Holders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolders

`func (o *Token) SetHolders(v int32)`

SetHolders sets Holders field to given value.

### HasHolders

`func (o *Token) HasHolders() bool`

HasHolders returns a boolean if a field has been set.

### GetMeta

`func (o *Token) GetMeta() TokenMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Token) GetMetaOk() (*TokenMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Token) SetMeta(v TokenMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Token) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOndemand

`func (o *Token) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *Token) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *Token) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *Token) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


