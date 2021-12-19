# MarketMetaOrderBooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to **[][]float32** |  | [optional] 
**Bids** | Pointer to **[][]float32** |  | [optional] 

## Methods

### NewMarketMetaOrderBooks

`func NewMarketMetaOrderBooks() *MarketMetaOrderBooks`

NewMarketMetaOrderBooks instantiates a new MarketMetaOrderBooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketMetaOrderBooksWithDefaults

`func NewMarketMetaOrderBooksWithDefaults() *MarketMetaOrderBooks`

NewMarketMetaOrderBooksWithDefaults instantiates a new MarketMetaOrderBooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *MarketMetaOrderBooks) GetAsks() [][]float32`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *MarketMetaOrderBooks) GetAsksOk() (*[][]float32, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *MarketMetaOrderBooks) SetAsks(v [][]float32)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *MarketMetaOrderBooks) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *MarketMetaOrderBooks) GetBids() [][]float32`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *MarketMetaOrderBooks) GetBidsOk() (*[][]float32, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *MarketMetaOrderBooks) SetBids(v [][]float32)`

SetBids sets Bids field to given value.

### HasBids

`func (o *MarketMetaOrderBooks) HasBids() bool`

HasBids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


