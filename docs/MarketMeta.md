# MarketMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to [**Address**](Address.md) |  | [optional] 
**Bids** | Pointer to [**Address**](Address.md) |  | [optional] 
**BaseLotSize** | Pointer to **string** |  | [optional] 
**QuoteLotSize** | Pointer to **string** |  | [optional] 
**BaseTokenDecimals** | Pointer to **int32** |  | [optional] 
**QuoteTokenDecimals** | Pointer to **int32** |  | [optional] 
**BaseVault** | Pointer to [**Address**](Address.md) |  | [optional] 
**QuoteVault** | Pointer to [**Address**](Address.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MarketPrice** | Pointer to **float32** |  | [optional] 
**OrderBooks** | Pointer to [**MarketMetaOrderBooks**](MarketMetaOrderBooks.md) |  | [optional] 
**Liquidity** | Pointer to [**MarketMetaLiquidity**](MarketMetaLiquidity.md) |  | [optional] 
**Volume** | Pointer to [**MarketMetaLiquidity**](MarketMetaLiquidity.md) |  | [optional] 
**Volumes** | Pointer to **[]float32** |  | [optional] 
**Volume24h** | Pointer to **float32** |  | [optional] 

## Methods

### NewMarketMeta

`func NewMarketMeta() *MarketMeta`

NewMarketMeta instantiates a new MarketMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMetaWithDefaults

`func NewMarketMetaWithDefaults() *MarketMeta`

NewMarketMetaWithDefaults instantiates a new MarketMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *MarketMeta) GetAsks() Address`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *MarketMeta) GetAsksOk() (*Address, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *MarketMeta) SetAsks(v Address)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *MarketMeta) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *MarketMeta) GetBids() Address`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *MarketMeta) GetBidsOk() (*Address, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *MarketMeta) SetBids(v Address)`

SetBids sets Bids field to given value.

### HasBids

`func (o *MarketMeta) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetBaseLotSize

`func (o *MarketMeta) GetBaseLotSize() string`

GetBaseLotSize returns the BaseLotSize field if non-nil, zero value otherwise.

### GetBaseLotSizeOk

`func (o *MarketMeta) GetBaseLotSizeOk() (*string, bool)`

GetBaseLotSizeOk returns a tuple with the BaseLotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLotSize

`func (o *MarketMeta) SetBaseLotSize(v string)`

SetBaseLotSize sets BaseLotSize field to given value.

### HasBaseLotSize

`func (o *MarketMeta) HasBaseLotSize() bool`

HasBaseLotSize returns a boolean if a field has been set.

### GetQuoteLotSize

`func (o *MarketMeta) GetQuoteLotSize() string`

GetQuoteLotSize returns the QuoteLotSize field if non-nil, zero value otherwise.

### GetQuoteLotSizeOk

`func (o *MarketMeta) GetQuoteLotSizeOk() (*string, bool)`

GetQuoteLotSizeOk returns a tuple with the QuoteLotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteLotSize

`func (o *MarketMeta) SetQuoteLotSize(v string)`

SetQuoteLotSize sets QuoteLotSize field to given value.

### HasQuoteLotSize

`func (o *MarketMeta) HasQuoteLotSize() bool`

HasQuoteLotSize returns a boolean if a field has been set.

### GetBaseTokenDecimals

`func (o *MarketMeta) GetBaseTokenDecimals() int32`

GetBaseTokenDecimals returns the BaseTokenDecimals field if non-nil, zero value otherwise.

### GetBaseTokenDecimalsOk

`func (o *MarketMeta) GetBaseTokenDecimalsOk() (*int32, bool)`

GetBaseTokenDecimalsOk returns a tuple with the BaseTokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTokenDecimals

`func (o *MarketMeta) SetBaseTokenDecimals(v int32)`

SetBaseTokenDecimals sets BaseTokenDecimals field to given value.

### HasBaseTokenDecimals

`func (o *MarketMeta) HasBaseTokenDecimals() bool`

HasBaseTokenDecimals returns a boolean if a field has been set.

### GetQuoteTokenDecimals

`func (o *MarketMeta) GetQuoteTokenDecimals() int32`

GetQuoteTokenDecimals returns the QuoteTokenDecimals field if non-nil, zero value otherwise.

### GetQuoteTokenDecimalsOk

`func (o *MarketMeta) GetQuoteTokenDecimalsOk() (*int32, bool)`

GetQuoteTokenDecimalsOk returns a tuple with the QuoteTokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTokenDecimals

`func (o *MarketMeta) SetQuoteTokenDecimals(v int32)`

SetQuoteTokenDecimals sets QuoteTokenDecimals field to given value.

### HasQuoteTokenDecimals

`func (o *MarketMeta) HasQuoteTokenDecimals() bool`

HasQuoteTokenDecimals returns a boolean if a field has been set.

### GetBaseVault

`func (o *MarketMeta) GetBaseVault() Address`

GetBaseVault returns the BaseVault field if non-nil, zero value otherwise.

### GetBaseVaultOk

`func (o *MarketMeta) GetBaseVaultOk() (*Address, bool)`

GetBaseVaultOk returns a tuple with the BaseVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVault

`func (o *MarketMeta) SetBaseVault(v Address)`

SetBaseVault sets BaseVault field to given value.

### HasBaseVault

`func (o *MarketMeta) HasBaseVault() bool`

HasBaseVault returns a boolean if a field has been set.

### GetQuoteVault

`func (o *MarketMeta) GetQuoteVault() Address`

GetQuoteVault returns the QuoteVault field if non-nil, zero value otherwise.

### GetQuoteVaultOk

`func (o *MarketMeta) GetQuoteVaultOk() (*Address, bool)`

GetQuoteVaultOk returns a tuple with the QuoteVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVault

`func (o *MarketMeta) SetQuoteVault(v Address)`

SetQuoteVault sets QuoteVault field to given value.

### HasQuoteVault

`func (o *MarketMeta) HasQuoteVault() bool`

HasQuoteVault returns a boolean if a field has been set.

### GetName

`func (o *MarketMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMarketPrice

`func (o *MarketMeta) GetMarketPrice() float32`

GetMarketPrice returns the MarketPrice field if non-nil, zero value otherwise.

### GetMarketPriceOk

`func (o *MarketMeta) GetMarketPriceOk() (*float32, bool)`

GetMarketPriceOk returns a tuple with the MarketPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketPrice

`func (o *MarketMeta) SetMarketPrice(v float32)`

SetMarketPrice sets MarketPrice field to given value.

### HasMarketPrice

`func (o *MarketMeta) HasMarketPrice() bool`

HasMarketPrice returns a boolean if a field has been set.

### GetOrderBooks

`func (o *MarketMeta) GetOrderBooks() MarketMetaOrderBooks`

GetOrderBooks returns the OrderBooks field if non-nil, zero value otherwise.

### GetOrderBooksOk

`func (o *MarketMeta) GetOrderBooksOk() (*MarketMetaOrderBooks, bool)`

GetOrderBooksOk returns a tuple with the OrderBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBooks

`func (o *MarketMeta) SetOrderBooks(v MarketMetaOrderBooks)`

SetOrderBooks sets OrderBooks field to given value.

### HasOrderBooks

`func (o *MarketMeta) HasOrderBooks() bool`

HasOrderBooks returns a boolean if a field has been set.

### GetLiquidity

`func (o *MarketMeta) GetLiquidity() MarketMetaLiquidity`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *MarketMeta) GetLiquidityOk() (*MarketMetaLiquidity, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *MarketMeta) SetLiquidity(v MarketMetaLiquidity)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *MarketMeta) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetVolume

`func (o *MarketMeta) GetVolume() MarketMetaLiquidity`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketMeta) GetVolumeOk() (*MarketMetaLiquidity, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketMeta) SetVolume(v MarketMetaLiquidity)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketMeta) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumes

`func (o *MarketMeta) GetVolumes() []float32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *MarketMeta) GetVolumesOk() (*[]float32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *MarketMeta) SetVolumes(v []float32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *MarketMeta) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetVolume24h

`func (o *MarketMeta) GetVolume24h() float32`

GetVolume24h returns the Volume24h field if non-nil, zero value otherwise.

### GetVolume24hOk

`func (o *MarketMeta) GetVolume24hOk() (*float32, bool)`

GetVolume24hOk returns a tuple with the Volume24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume24h

`func (o *MarketMeta) SetVolume24h(v float32)`

SetVolume24h sets Volume24h field to given value.

### HasVolume24h

`func (o *MarketMeta) HasVolume24h() bool`

HasVolume24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


