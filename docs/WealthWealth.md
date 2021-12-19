# WealthWealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to [**WealthWealthDistribution**](WealthWealthDistribution.md) |  | [optional] 
**GroupedBalances** | Pointer to [**[]WealthWealthGroupedBalances**](WealthWealthGroupedBalances.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 

## Methods

### NewWealthWealth

`func NewWealthWealth() *WealthWealth`

NewWealthWealth instantiates a new WealthWealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWealthWealthWithDefaults

`func NewWealthWealthWithDefaults() *WealthWealth`

NewWealthWealthWithDefaults instantiates a new WealthWealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *WealthWealth) GetDistribution() WealthWealthDistribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *WealthWealth) GetDistributionOk() (*WealthWealthDistribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *WealthWealth) SetDistribution(v WealthWealthDistribution)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *WealthWealth) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetGroupedBalances

`func (o *WealthWealth) GetGroupedBalances() []WealthWealthGroupedBalances`

GetGroupedBalances returns the GroupedBalances field if non-nil, zero value otherwise.

### GetGroupedBalancesOk

`func (o *WealthWealth) GetGroupedBalancesOk() (*[]WealthWealthGroupedBalances, bool)`

GetGroupedBalancesOk returns a tuple with the GroupedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedBalances

`func (o *WealthWealth) SetGroupedBalances(v []WealthWealthGroupedBalances)`

SetGroupedBalances sets GroupedBalances field to given value.

### HasGroupedBalances

`func (o *WealthWealth) HasGroupedBalances() bool`

HasGroupedBalances returns a boolean if a field has been set.

### GetPrice

`func (o *WealthWealth) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *WealthWealth) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *WealthWealth) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *WealthWealth) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


