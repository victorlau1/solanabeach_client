# TokenPricedata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** |  | [optional] 
**Volume24h** | Pointer to **float32** |  | [optional] 
**PercentChange1h** | Pointer to **float32** |  | [optional] 
**LastUpdated** | Pointer to **int32** |  | [optional] 

## Methods

### NewTokenPricedata

`func NewTokenPricedata() *TokenPricedata`

NewTokenPricedata instantiates a new TokenPricedata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPricedataWithDefaults

`func NewTokenPricedataWithDefaults() *TokenPricedata`

NewTokenPricedataWithDefaults instantiates a new TokenPricedata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *TokenPricedata) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TokenPricedata) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TokenPricedata) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TokenPricedata) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetVolume24h

`func (o *TokenPricedata) GetVolume24h() float32`

GetVolume24h returns the Volume24h field if non-nil, zero value otherwise.

### GetVolume24hOk

`func (o *TokenPricedata) GetVolume24hOk() (*float32, bool)`

GetVolume24hOk returns a tuple with the Volume24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume24h

`func (o *TokenPricedata) SetVolume24h(v float32)`

SetVolume24h sets Volume24h field to given value.

### HasVolume24h

`func (o *TokenPricedata) HasVolume24h() bool`

HasVolume24h returns a boolean if a field has been set.

### GetPercentChange1h

`func (o *TokenPricedata) GetPercentChange1h() float32`

GetPercentChange1h returns the PercentChange1h field if non-nil, zero value otherwise.

### GetPercentChange1hOk

`func (o *TokenPricedata) GetPercentChange1hOk() (*float32, bool)`

GetPercentChange1hOk returns a tuple with the PercentChange1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange1h

`func (o *TokenPricedata) SetPercentChange1h(v float32)`

SetPercentChange1h sets PercentChange1h field to given value.

### HasPercentChange1h

`func (o *TokenPricedata) HasPercentChange1h() bool`

HasPercentChange1h returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TokenPricedata) GetLastUpdated() int32`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TokenPricedata) GetLastUpdatedOk() (*int32, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TokenPricedata) SetLastUpdated(v int32)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TokenPricedata) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


