# Validator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Validator** | Pointer to [**ValidatorValidator**](ValidatorValidator.md) |  | [optional] 
**Slots** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**Historic** | Pointer to [**[]ValidatorHistoric**](ValidatorHistoric.md) |  | [optional] 
**LatestBlocks** | Pointer to [**[]ValidatorLatestBlocks**](ValidatorLatestBlocks.md) |  | [optional] 

## Methods

### NewValidator

`func NewValidator() *Validator`

NewValidator instantiates a new Validator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorWithDefaults

`func NewValidatorWithDefaults() *Validator`

NewValidatorWithDefaults instantiates a new Validator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidator

`func (o *Validator) GetValidator() ValidatorValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *Validator) GetValidatorOk() (*ValidatorValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *Validator) SetValidator(v ValidatorValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *Validator) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetSlots

`func (o *Validator) GetSlots() [][]map[string]interface{}`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *Validator) GetSlotsOk() (*[][]map[string]interface{}, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *Validator) SetSlots(v [][]map[string]interface{})`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *Validator) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetHistoric

`func (o *Validator) GetHistoric() []ValidatorHistoric`

GetHistoric returns the Historic field if non-nil, zero value otherwise.

### GetHistoricOk

`func (o *Validator) GetHistoricOk() (*[]ValidatorHistoric, bool)`

GetHistoricOk returns a tuple with the Historic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoric

`func (o *Validator) SetHistoric(v []ValidatorHistoric)`

SetHistoric sets Historic field to given value.

### HasHistoric

`func (o *Validator) HasHistoric() bool`

HasHistoric returns a boolean if a field has been set.

### GetLatestBlocks

`func (o *Validator) GetLatestBlocks() []ValidatorLatestBlocks`

GetLatestBlocks returns the LatestBlocks field if non-nil, zero value otherwise.

### GetLatestBlocksOk

`func (o *Validator) GetLatestBlocksOk() (*[]ValidatorLatestBlocks, bool)`

GetLatestBlocksOk returns a tuple with the LatestBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBlocks

`func (o *Validator) SetLatestBlocks(v []ValidatorLatestBlocks)`

SetLatestBlocks sets LatestBlocks field to given value.

### HasLatestBlocks

`func (o *Validator) HasLatestBlocks() bool`

HasLatestBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


