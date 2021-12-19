# TransactionAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**Address**](Address.md) |  | [optional] 
**Writable** | Pointer to **bool** |  | [optional] 
**Signer** | Pointer to **bool** |  | [optional] 
**Program** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransactionAccounts

`func NewTransactionAccounts() *TransactionAccounts`

NewTransactionAccounts instantiates a new TransactionAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAccountsWithDefaults

`func NewTransactionAccountsWithDefaults() *TransactionAccounts`

NewTransactionAccountsWithDefaults instantiates a new TransactionAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TransactionAccounts) GetAccount() Address`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionAccounts) GetAccountOk() (*Address, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionAccounts) SetAccount(v Address)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TransactionAccounts) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetWritable

`func (o *TransactionAccounts) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *TransactionAccounts) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *TransactionAccounts) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *TransactionAccounts) HasWritable() bool`

HasWritable returns a boolean if a field has been set.

### GetSigner

`func (o *TransactionAccounts) GetSigner() bool`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *TransactionAccounts) GetSignerOk() (*bool, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *TransactionAccounts) SetSigner(v bool)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *TransactionAccounts) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetProgram

`func (o *TransactionAccounts) GetProgram() bool`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *TransactionAccounts) GetProgramOk() (*bool, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *TransactionAccounts) SetProgram(v bool)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *TransactionAccounts) HasProgram() bool`

HasProgram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


