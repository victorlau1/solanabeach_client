# AccountSerumOrderMarket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**Address**](Address.md) |  | [optional] 
**QuoteMint** | Pointer to [**Address**](Address.md) |  | [optional] 
**BaseMint** | Pointer to [**Address**](Address.md) |  | [optional] 
**CurrentPrice** | Pointer to **float32** |  | [optional] 
**BaseDecimals** | Pointer to **int32** |  | [optional] 
**QuoteDecimals** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountSerumOrderMarket

`func NewAccountSerumOrderMarket() *AccountSerumOrderMarket`

NewAccountSerumOrderMarket instantiates a new AccountSerumOrderMarket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSerumOrderMarketWithDefaults

`func NewAccountSerumOrderMarketWithDefaults() *AccountSerumOrderMarket`

NewAccountSerumOrderMarketWithDefaults instantiates a new AccountSerumOrderMarket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *AccountSerumOrderMarket) GetPubkey() Address`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *AccountSerumOrderMarket) GetPubkeyOk() (*Address, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *AccountSerumOrderMarket) SetPubkey(v Address)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *AccountSerumOrderMarket) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetQuoteMint

`func (o *AccountSerumOrderMarket) GetQuoteMint() Address`

GetQuoteMint returns the QuoteMint field if non-nil, zero value otherwise.

### GetQuoteMintOk

`func (o *AccountSerumOrderMarket) GetQuoteMintOk() (*Address, bool)`

GetQuoteMintOk returns a tuple with the QuoteMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteMint

`func (o *AccountSerumOrderMarket) SetQuoteMint(v Address)`

SetQuoteMint sets QuoteMint field to given value.

### HasQuoteMint

`func (o *AccountSerumOrderMarket) HasQuoteMint() bool`

HasQuoteMint returns a boolean if a field has been set.

### GetBaseMint

`func (o *AccountSerumOrderMarket) GetBaseMint() Address`

GetBaseMint returns the BaseMint field if non-nil, zero value otherwise.

### GetBaseMintOk

`func (o *AccountSerumOrderMarket) GetBaseMintOk() (*Address, bool)`

GetBaseMintOk returns a tuple with the BaseMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMint

`func (o *AccountSerumOrderMarket) SetBaseMint(v Address)`

SetBaseMint sets BaseMint field to given value.

### HasBaseMint

`func (o *AccountSerumOrderMarket) HasBaseMint() bool`

HasBaseMint returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *AccountSerumOrderMarket) GetCurrentPrice() float32`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *AccountSerumOrderMarket) GetCurrentPriceOk() (*float32, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *AccountSerumOrderMarket) SetCurrentPrice(v float32)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *AccountSerumOrderMarket) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetBaseDecimals

`func (o *AccountSerumOrderMarket) GetBaseDecimals() int32`

GetBaseDecimals returns the BaseDecimals field if non-nil, zero value otherwise.

### GetBaseDecimalsOk

`func (o *AccountSerumOrderMarket) GetBaseDecimalsOk() (*int32, bool)`

GetBaseDecimalsOk returns a tuple with the BaseDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDecimals

`func (o *AccountSerumOrderMarket) SetBaseDecimals(v int32)`

SetBaseDecimals sets BaseDecimals field to given value.

### HasBaseDecimals

`func (o *AccountSerumOrderMarket) HasBaseDecimals() bool`

HasBaseDecimals returns a boolean if a field has been set.

### GetQuoteDecimals

`func (o *AccountSerumOrderMarket) GetQuoteDecimals() int32`

GetQuoteDecimals returns the QuoteDecimals field if non-nil, zero value otherwise.

### GetQuoteDecimalsOk

`func (o *AccountSerumOrderMarket) GetQuoteDecimalsOk() (*int32, bool)`

GetQuoteDecimalsOk returns a tuple with the QuoteDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteDecimals

`func (o *AccountSerumOrderMarket) SetQuoteDecimals(v int32)`

SetQuoteDecimals sets QuoteDecimals field to given value.

### HasQuoteDecimals

`func (o *AccountSerumOrderMarket) HasQuoteDecimals() bool`

HasQuoteDecimals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


