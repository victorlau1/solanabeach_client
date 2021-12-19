# BlockProgramstats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**ProgramId** | Pointer to [**Address**](Address.md) |  | [optional] 
**Instructions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBlockProgramstats

`func NewBlockProgramstats() *BlockProgramstats`

NewBlockProgramstats instantiates a new BlockProgramstats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockProgramstatsWithDefaults

`func NewBlockProgramstatsWithDefaults() *BlockProgramstats`

NewBlockProgramstatsWithDefaults instantiates a new BlockProgramstats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BlockProgramstats) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BlockProgramstats) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BlockProgramstats) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BlockProgramstats) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetProgramId

`func (o *BlockProgramstats) GetProgramId() Address`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BlockProgramstats) GetProgramIdOk() (*Address, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BlockProgramstats) SetProgramId(v Address)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BlockProgramstats) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetInstructions

`func (o *BlockProgramstats) GetInstructions() map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *BlockProgramstats) GetInstructionsOk() (*map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *BlockProgramstats) SetInstructions(v map[string]interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *BlockProgramstats) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


