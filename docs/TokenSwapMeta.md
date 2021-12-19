# TokenSwapMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supply** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**Liquidity** | Pointer to [**TokenSwapMetaLiquidity**](TokenSwapMetaLiquidity.md) |  | [optional] 
**Volume** | Pointer to [**TokenSwapMetaLiquidity**](TokenSwapMetaLiquidity.md) |  | [optional] 
**Prices** | Pointer to [**TokenSwapMetaPrices**](TokenSwapMetaPrices.md) |  | [optional] 
**Volumes** | Pointer to **[]float32** |  | [optional] 
**Volume24h** | Pointer to **float32** |  | [optional] 

## Methods

### NewTokenSwapMeta

`func NewTokenSwapMeta() *TokenSwapMeta`

NewTokenSwapMeta instantiates a new TokenSwapMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSwapMetaWithDefaults

`func NewTokenSwapMetaWithDefaults() *TokenSwapMeta`

NewTokenSwapMetaWithDefaults instantiates a new TokenSwapMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupply

`func (o *TokenSwapMeta) GetSupply() Amount`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *TokenSwapMeta) GetSupplyOk() (*Amount, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *TokenSwapMeta) SetSupply(v Amount)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *TokenSwapMeta) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetLiquidity

`func (o *TokenSwapMeta) GetLiquidity() TokenSwapMetaLiquidity`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *TokenSwapMeta) GetLiquidityOk() (*TokenSwapMetaLiquidity, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *TokenSwapMeta) SetLiquidity(v TokenSwapMetaLiquidity)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *TokenSwapMeta) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetVolume

`func (o *TokenSwapMeta) GetVolume() TokenSwapMetaLiquidity`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TokenSwapMeta) GetVolumeOk() (*TokenSwapMetaLiquidity, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TokenSwapMeta) SetVolume(v TokenSwapMetaLiquidity)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *TokenSwapMeta) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetPrices

`func (o *TokenSwapMeta) GetPrices() TokenSwapMetaPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *TokenSwapMeta) GetPricesOk() (*TokenSwapMetaPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *TokenSwapMeta) SetPrices(v TokenSwapMetaPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *TokenSwapMeta) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetVolumes

`func (o *TokenSwapMeta) GetVolumes() []float32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *TokenSwapMeta) GetVolumesOk() (*[]float32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *TokenSwapMeta) SetVolumes(v []float32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *TokenSwapMeta) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetVolume24h

`func (o *TokenSwapMeta) GetVolume24h() float32`

GetVolume24h returns the Volume24h field if non-nil, zero value otherwise.

### GetVolume24hOk

`func (o *TokenSwapMeta) GetVolume24hOk() (*float32, bool)`

GetVolume24hOk returns a tuple with the Volume24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume24h

`func (o *TokenSwapMeta) SetVolume24h(v float32)`

SetVolume24h sets Volume24h field to given value.

### HasVolume24h

`func (o *TokenSwapMeta) HasVolume24h() bool`

HasVolume24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


