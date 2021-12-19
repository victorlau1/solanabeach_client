# TransactionHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumRequiredSignatures** | Pointer to **int32** |  | [optional] 
**NumReadonlySignedAccounts** | Pointer to **int32** |  | [optional] 
**NumReadonlyUnsignedAccounts** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransactionHeader

`func NewTransactionHeader() *TransactionHeader`

NewTransactionHeader instantiates a new TransactionHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionHeaderWithDefaults

`func NewTransactionHeaderWithDefaults() *TransactionHeader`

NewTransactionHeaderWithDefaults instantiates a new TransactionHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumRequiredSignatures

`func (o *TransactionHeader) GetNumRequiredSignatures() int32`

GetNumRequiredSignatures returns the NumRequiredSignatures field if non-nil, zero value otherwise.

### GetNumRequiredSignaturesOk

`func (o *TransactionHeader) GetNumRequiredSignaturesOk() (*int32, bool)`

GetNumRequiredSignaturesOk returns a tuple with the NumRequiredSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRequiredSignatures

`func (o *TransactionHeader) SetNumRequiredSignatures(v int32)`

SetNumRequiredSignatures sets NumRequiredSignatures field to given value.

### HasNumRequiredSignatures

`func (o *TransactionHeader) HasNumRequiredSignatures() bool`

HasNumRequiredSignatures returns a boolean if a field has been set.

### GetNumReadonlySignedAccounts

`func (o *TransactionHeader) GetNumReadonlySignedAccounts() int32`

GetNumReadonlySignedAccounts returns the NumReadonlySignedAccounts field if non-nil, zero value otherwise.

### GetNumReadonlySignedAccountsOk

`func (o *TransactionHeader) GetNumReadonlySignedAccountsOk() (*int32, bool)`

GetNumReadonlySignedAccountsOk returns a tuple with the NumReadonlySignedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReadonlySignedAccounts

`func (o *TransactionHeader) SetNumReadonlySignedAccounts(v int32)`

SetNumReadonlySignedAccounts sets NumReadonlySignedAccounts field to given value.

### HasNumReadonlySignedAccounts

`func (o *TransactionHeader) HasNumReadonlySignedAccounts() bool`

HasNumReadonlySignedAccounts returns a boolean if a field has been set.

### GetNumReadonlyUnsignedAccounts

`func (o *TransactionHeader) GetNumReadonlyUnsignedAccounts() int32`

GetNumReadonlyUnsignedAccounts returns the NumReadonlyUnsignedAccounts field if non-nil, zero value otherwise.

### GetNumReadonlyUnsignedAccountsOk

`func (o *TransactionHeader) GetNumReadonlyUnsignedAccountsOk() (*int32, bool)`

GetNumReadonlyUnsignedAccountsOk returns a tuple with the NumReadonlyUnsignedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReadonlyUnsignedAccounts

`func (o *TransactionHeader) SetNumReadonlyUnsignedAccounts(v int32)`

SetNumReadonlyUnsignedAccounts sets NumReadonlyUnsignedAccounts field to given value.

### HasNumReadonlyUnsignedAccounts

`func (o *TransactionHeader) HasNumReadonlyUnsignedAccounts() bool`

HasNumReadonlyUnsignedAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


