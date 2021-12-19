# TransactionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Err** | Pointer to **map[string]interface{}** |  | [optional] 
**Fee** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 
**LogMessages** | Pointer to **[]string** |  | [optional] 
**PreBalances** | Pointer to **[]int32** |  | [optional] 
**PostBalances** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewTransactionMeta

`func NewTransactionMeta() *TransactionMeta`

NewTransactionMeta instantiates a new TransactionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMetaWithDefaults

`func NewTransactionMetaWithDefaults() *TransactionMeta`

NewTransactionMetaWithDefaults instantiates a new TransactionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErr

`func (o *TransactionMeta) GetErr() map[string]interface{}`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *TransactionMeta) GetErrOk() (*map[string]interface{}, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *TransactionMeta) SetErr(v map[string]interface{})`

SetErr sets Err field to given value.

### HasErr

`func (o *TransactionMeta) HasErr() bool`

HasErr returns a boolean if a field has been set.

### GetFee

`func (o *TransactionMeta) GetFee() int32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionMeta) GetFeeOk() (*int32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionMeta) SetFee(v int32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransactionMeta) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionMeta) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionMeta) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionMeta) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionMeta) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogMessages

`func (o *TransactionMeta) GetLogMessages() []string`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *TransactionMeta) GetLogMessagesOk() (*[]string, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *TransactionMeta) SetLogMessages(v []string)`

SetLogMessages sets LogMessages field to given value.

### HasLogMessages

`func (o *TransactionMeta) HasLogMessages() bool`

HasLogMessages returns a boolean if a field has been set.

### GetPreBalances

`func (o *TransactionMeta) GetPreBalances() []int32`

GetPreBalances returns the PreBalances field if non-nil, zero value otherwise.

### GetPreBalancesOk

`func (o *TransactionMeta) GetPreBalancesOk() (*[]int32, bool)`

GetPreBalancesOk returns a tuple with the PreBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBalances

`func (o *TransactionMeta) SetPreBalances(v []int32)`

SetPreBalances sets PreBalances field to given value.

### HasPreBalances

`func (o *TransactionMeta) HasPreBalances() bool`

HasPreBalances returns a boolean if a field has been set.

### GetPostBalances

`func (o *TransactionMeta) GetPostBalances() []int32`

GetPostBalances returns the PostBalances field if non-nil, zero value otherwise.

### GetPostBalancesOk

`func (o *TransactionMeta) GetPostBalancesOk() (*[]int32, bool)`

GetPostBalancesOk returns a tuple with the PostBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBalances

`func (o *TransactionMeta) SetPostBalances(v []int32)`

SetPostBalances sets PostBalances field to given value.

### HasPostBalances

`func (o *TransactionMeta) HasPostBalances() bool`

HasPostBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


