# TransactionRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to **[]int32** |  | [optional] 
**ProgramIdIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransactionRaw

`func NewTransactionRaw() *TransactionRaw`

NewTransactionRaw instantiates a new TransactionRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRawWithDefaults

`func NewTransactionRawWithDefaults() *TransactionRaw`

NewTransactionRawWithDefaults instantiates a new TransactionRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionRaw) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionRaw) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionRaw) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionRaw) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAccounts

`func (o *TransactionRaw) GetAccounts() []int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *TransactionRaw) GetAccountsOk() (*[]int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *TransactionRaw) SetAccounts(v []int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *TransactionRaw) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetProgramIdIndex

`func (o *TransactionRaw) GetProgramIdIndex() int32`

GetProgramIdIndex returns the ProgramIdIndex field if non-nil, zero value otherwise.

### GetProgramIdIndexOk

`func (o *TransactionRaw) GetProgramIdIndexOk() (*int32, bool)`

GetProgramIdIndexOk returns a tuple with the ProgramIdIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramIdIndex

`func (o *TransactionRaw) SetProgramIdIndex(v int32)`

SetProgramIdIndex sets ProgramIdIndex field to given value.

### HasProgramIdIndex

`func (o *TransactionRaw) HasProgramIdIndex() bool`

HasProgramIdIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


