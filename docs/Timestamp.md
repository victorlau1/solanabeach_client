# Timestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absolute** | Pointer to **int32** |  | [optional] 
**Relative** | Pointer to **int32** |  | [optional] 

## Methods

### NewTimestamp

`func NewTimestamp() *Timestamp`

NewTimestamp instantiates a new Timestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampWithDefaults

`func NewTimestampWithDefaults() *Timestamp`

NewTimestampWithDefaults instantiates a new Timestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolute

`func (o *Timestamp) GetAbsolute() int32`

GetAbsolute returns the Absolute field if non-nil, zero value otherwise.

### GetAbsoluteOk

`func (o *Timestamp) GetAbsoluteOk() (*int32, bool)`

GetAbsoluteOk returns a tuple with the Absolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolute

`func (o *Timestamp) SetAbsolute(v int32)`

SetAbsolute sets Absolute field to given value.

### HasAbsolute

`func (o *Timestamp) HasAbsolute() bool`

HasAbsolute returns a boolean if a field has been set.

### GetRelative

`func (o *Timestamp) GetRelative() int32`

GetRelative returns the Relative field if non-nil, zero value otherwise.

### GetRelativeOk

`func (o *Timestamp) GetRelativeOk() (*int32, bool)`

GetRelativeOk returns a tuple with the Relative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelative

`func (o *Timestamp) SetRelative(v int32)`

SetRelative sets Relative field to given value.

### HasRelative

`func (o *Timestamp) HasRelative() bool`

HasRelative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


