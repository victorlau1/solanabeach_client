# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocknumber** | Pointer to **int64** |  | [optional] 
**Blockhash** | Pointer to **string** |  | [optional] 
**Previousblockhash** | Pointer to **string** |  | [optional] 
**Parentslot** | Pointer to **int64** |  | [optional] 
**Blocktime** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**Metrics** | Pointer to [**BlockMetrics**](BlockMetrics.md) |  | [optional] 
**Programstats** | Pointer to [**[]BlockProgramstats**](BlockProgramstats.md) |  | [optional] 
**Rewards** | Pointer to **map[string]interface{}** |  | [optional] 
**Proposer** | Pointer to **string** |  | [optional] 
**ProposerData** | Pointer to [**BlockProposerData**](BlockProposerData.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 

## Methods

### NewBlock

`func NewBlock() *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocknumber

`func (o *Block) GetBlocknumber() int64`

GetBlocknumber returns the Blocknumber field if non-nil, zero value otherwise.

### GetBlocknumberOk

`func (o *Block) GetBlocknumberOk() (*int64, bool)`

GetBlocknumberOk returns a tuple with the Blocknumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocknumber

`func (o *Block) SetBlocknumber(v int64)`

SetBlocknumber sets Blocknumber field to given value.

### HasBlocknumber

`func (o *Block) HasBlocknumber() bool`

HasBlocknumber returns a boolean if a field has been set.

### GetBlockhash

`func (o *Block) GetBlockhash() string`

GetBlockhash returns the Blockhash field if non-nil, zero value otherwise.

### GetBlockhashOk

`func (o *Block) GetBlockhashOk() (*string, bool)`

GetBlockhashOk returns a tuple with the Blockhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockhash

`func (o *Block) SetBlockhash(v string)`

SetBlockhash sets Blockhash field to given value.

### HasBlockhash

`func (o *Block) HasBlockhash() bool`

HasBlockhash returns a boolean if a field has been set.

### GetPreviousblockhash

`func (o *Block) GetPreviousblockhash() string`

GetPreviousblockhash returns the Previousblockhash field if non-nil, zero value otherwise.

### GetPreviousblockhashOk

`func (o *Block) GetPreviousblockhashOk() (*string, bool)`

GetPreviousblockhashOk returns a tuple with the Previousblockhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousblockhash

`func (o *Block) SetPreviousblockhash(v string)`

SetPreviousblockhash sets Previousblockhash field to given value.

### HasPreviousblockhash

`func (o *Block) HasPreviousblockhash() bool`

HasPreviousblockhash returns a boolean if a field has been set.

### GetParentslot

`func (o *Block) GetParentslot() int64`

GetParentslot returns the Parentslot field if non-nil, zero value otherwise.

### GetParentslotOk

`func (o *Block) GetParentslotOk() (*int64, bool)`

GetParentslotOk returns a tuple with the Parentslot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentslot

`func (o *Block) SetParentslot(v int64)`

SetParentslot sets Parentslot field to given value.

### HasParentslot

`func (o *Block) HasParentslot() bool`

HasParentslot returns a boolean if a field has been set.

### GetBlocktime

`func (o *Block) GetBlocktime() Timestamp`

GetBlocktime returns the Blocktime field if non-nil, zero value otherwise.

### GetBlocktimeOk

`func (o *Block) GetBlocktimeOk() (*Timestamp, bool)`

GetBlocktimeOk returns a tuple with the Blocktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocktime

`func (o *Block) SetBlocktime(v Timestamp)`

SetBlocktime sets Blocktime field to given value.

### HasBlocktime

`func (o *Block) HasBlocktime() bool`

HasBlocktime returns a boolean if a field has been set.

### GetMetrics

`func (o *Block) GetMetrics() BlockMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Block) GetMetricsOk() (*BlockMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Block) SetMetrics(v BlockMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Block) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProgramstats

`func (o *Block) GetProgramstats() []BlockProgramstats`

GetProgramstats returns the Programstats field if non-nil, zero value otherwise.

### GetProgramstatsOk

`func (o *Block) GetProgramstatsOk() (*[]BlockProgramstats, bool)`

GetProgramstatsOk returns a tuple with the Programstats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramstats

`func (o *Block) SetProgramstats(v []BlockProgramstats)`

SetProgramstats sets Programstats field to given value.

### HasProgramstats

`func (o *Block) HasProgramstats() bool`

HasProgramstats returns a boolean if a field has been set.

### GetRewards

`func (o *Block) GetRewards() map[string]interface{}`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *Block) GetRewardsOk() (*map[string]interface{}, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *Block) SetRewards(v map[string]interface{})`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *Block) HasRewards() bool`

HasRewards returns a boolean if a field has been set.

### GetProposer

`func (o *Block) GetProposer() string`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *Block) GetProposerOk() (*string, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *Block) SetProposer(v string)`

SetProposer sets Proposer field to given value.

### HasProposer

`func (o *Block) HasProposer() bool`

HasProposer returns a boolean if a field has been set.

### GetProposerData

`func (o *Block) GetProposerData() BlockProposerData`

GetProposerData returns the ProposerData field if non-nil, zero value otherwise.

### GetProposerDataOk

`func (o *Block) GetProposerDataOk() (*BlockProposerData, bool)`

GetProposerDataOk returns a tuple with the ProposerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposerData

`func (o *Block) SetProposerData(v BlockProposerData)`

SetProposerData sets ProposerData field to given value.

### HasProposerData

`func (o *Block) HasProposerData() bool`

HasProposerData returns a boolean if a field has been set.

### GetOndemand

`func (o *Block) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *Block) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *Block) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *Block) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


