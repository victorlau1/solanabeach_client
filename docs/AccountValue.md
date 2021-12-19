# AccountValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to [**AccountValueBase**](AccountValueBase.md) |  | [optional] 
**Extended** | Pointer to **map[string]interface{}** | Depends on the type of object | [optional] 

## Methods

### NewAccountValue

`func NewAccountValue() *AccountValue`

NewAccountValue instantiates a new AccountValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountValueWithDefaults

`func NewAccountValueWithDefaults() *AccountValue`

NewAccountValueWithDefaults instantiates a new AccountValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *AccountValue) GetBase() AccountValueBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AccountValue) GetBaseOk() (*AccountValueBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AccountValue) SetBase(v AccountValueBase)`

SetBase sets Base field to given value.

### HasBase

`func (o *AccountValue) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetExtended

`func (o *AccountValue) GetExtended() map[string]interface{}`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *AccountValue) GetExtendedOk() (*map[string]interface{}, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *AccountValue) SetExtended(v map[string]interface{})`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *AccountValue) HasExtended() bool`

HasExtended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


