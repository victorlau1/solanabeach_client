# ValidatorLatestBlocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocknumber** | Pointer to **int32** |  | [optional] 
**Blocktime** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Metrics** | Pointer to [**ValidatorMetrics**](ValidatorMetrics.md) |  | [optional] 
**Proposer** | Pointer to **string** |  | [optional] 

## Methods

### NewValidatorLatestBlocks

`func NewValidatorLatestBlocks() *ValidatorLatestBlocks`

NewValidatorLatestBlocks instantiates a new ValidatorLatestBlocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorLatestBlocksWithDefaults

`func NewValidatorLatestBlocksWithDefaults() *ValidatorLatestBlocks`

NewValidatorLatestBlocksWithDefaults instantiates a new ValidatorLatestBlocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocknumber

`func (o *ValidatorLatestBlocks) GetBlocknumber() int32`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *ValidatorLatestBlocks) GetBlocknumberOk() (*int32, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *ValidatorLatestBlocks) SetBlocknumber(v int32)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *ValidatorLatestBlocks) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetBlocktime

`func (o *ValidatorLatestBlocks) GetBlocktime() Timestamp`

GetBlocktime returns the Blocktime field if non-nil, zero value otherwise.

### GetBlocktimeOk

`func (o *ValidatorLatestBlocks) GetBlocktimeOk() (*Timestamp, bool)`

GetBlocktimeOk returns a tuple with the Blocktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocktime

`func (o *ValidatorLatestBlocks) SetBlocktime(v Timestamp)`

SetBlocktime sets Blocktime field to given value.

### HasBlocktime

`func (o *ValidatorLatestBlocks) HasBlocktime() bool`

HasBlocktime returns a boolean if a field has been set.

### GetMetrics

`func (o *ValidatorLatestBlocks) GetMetrics() ValidatorMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ValidatorLatestBlocks) GetMetricsOk() (*ValidatorMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ValidatorLatestBlocks) SetMetrics(v ValidatorMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ValidatorLatestBlocks) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProposer

`func (o *ValidatorLatestBlocks) GetProposer() string`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *ValidatorLatestBlocks) GetProposerOk() (*string, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *ValidatorLatestBlocks) SetProposer(v string)`

SetProposer sets Proposer field to given value.

### HasProposer

`func (o *ValidatorLatestBlocks) HasProposer() bool`

HasProposer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


