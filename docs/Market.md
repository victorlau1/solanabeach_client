# Market

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**Address**](Address.md) |  | [optional] 
**Basemint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Quotemint** | Pointer to [**Address**](Address.md) |  | [optional] 
**Meta** | Pointer to [**MarketMeta**](MarketMeta.md) |  | [optional] 

## Methods

### NewMarket

`func NewMarket() *Market`

NewMarket instantiates a new Market object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketWithDefaults

`func NewMarketWithDefaults() *Market`

NewMarketWithDefaults instantiates a new Market object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *Market) GetPubkey() Address`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *Market) GetPubkeyOk() (*Address, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *Market) SetPubkey(v Address)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *Market) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetBasemint

`func (o *Market) GetBasemint() Address`

GetBasemint returns the Basemint field if non-nil, zero value otherwise.

### GetBasemintOk

`func (o *Market) GetBasemintOk() (*Address, bool)`

GetBasemintOk returns a tuple with the Basemint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasemint

`func (o *Market) SetBasemint(v Address)`

SetBasemint sets Basemint field to given value.

### HasBasemint

`func (o *Market) HasBasemint() bool`

HasBasemint returns a boolean if a field has been set.

### GetQuotemint

`func (o *Market) GetQuotemint() Address`

GetQuotemint returns the Quotemint field if non-nil, zero value otherwise.

### GetQuotemintOk

`func (o *Market) GetQuotemintOk() (*Address, bool)`

GetQuotemintOk returns a tuple with the Quotemint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotemint

`func (o *Market) SetQuotemint(v Address)`

SetQuotemint sets Quotemint field to given value.

### HasQuotemint

`func (o *Market) HasQuotemint() bool`

HasQuotemint returns a boolean if a field has been set.

### GetMeta

`func (o *Market) GetMeta() MarketMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Market) GetMetaOk() (*MarketMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Market) SetMeta(v MarketMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Market) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


