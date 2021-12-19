# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionHash** | Pointer to **string** |  | [optional] 
**BlockNumber** | Pointer to **int64** |  | [optional] 
**Index** | Pointer to **int64** |  | [optional] 
**Accounts** | Pointer to [**[]TransactionAccounts**](TransactionAccounts.md) |  | [optional] 
**Header** | Pointer to [**TransactionHeader**](TransactionHeader.md) |  | [optional] 
**Instructions** | Pointer to [**[]TransactionInstructions**](TransactionInstructions.md) |  | [optional] 
**RecentBlockhash** | Pointer to **string** |  | [optional] 
**Signatures** | Pointer to **[]string** |  | [optional] 
**Meta** | Pointer to [**TransactionMeta**](TransactionMeta.md) |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Blocktime** | Pointer to [**Timestamp**](Timestamp.md) |  | [optional] 
**MostImportantInstruction** | Pointer to [**TransactionMostImportantInstruction**](TransactionMostImportantInstruction.md) |  | [optional] 
**Smart** | Pointer to [**[]TransactionSmart**](TransactionSmart.md) |  | [optional] 
**Ondemand** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *Transaction) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *Transaction) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Transaction) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Transaction) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *Transaction) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetIndex

`func (o *Transaction) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Transaction) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Transaction) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Transaction) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetAccounts

`func (o *Transaction) GetAccounts() []TransactionAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Transaction) GetAccountsOk() (*[]TransactionAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Transaction) SetAccounts(v []TransactionAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Transaction) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetHeader

`func (o *Transaction) GetHeader() TransactionHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *Transaction) GetHeaderOk() (*TransactionHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *Transaction) SetHeader(v TransactionHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *Transaction) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetInstructions

`func (o *Transaction) GetInstructions() []TransactionInstructions`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Transaction) GetInstructionsOk() (*[]TransactionInstructions, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Transaction) SetInstructions(v []TransactionInstructions)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *Transaction) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetRecentBlockhash

`func (o *Transaction) GetRecentBlockhash() string`

GetRecentBlockhash returns the RecentBlockhash field if non-nil, zero value otherwise.

### GetRecentBlockhashOk

`func (o *Transaction) GetRecentBlockhashOk() (*string, bool)`

GetRecentBlockhashOk returns a tuple with the RecentBlockhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentBlockhash

`func (o *Transaction) SetRecentBlockhash(v string)`

SetRecentBlockhash sets RecentBlockhash field to given value.

### HasRecentBlockhash

`func (o *Transaction) HasRecentBlockhash() bool`

HasRecentBlockhash returns a boolean if a field has been set.

### GetSignatures

`func (o *Transaction) GetSignatures() []string`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *Transaction) GetSignaturesOk() (*[]string, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *Transaction) SetSignatures(v []string)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *Transaction) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.

### GetMeta

`func (o *Transaction) GetMeta() TransactionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Transaction) GetMetaOk() (*TransactionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Transaction) SetMeta(v TransactionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Transaction) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValid

`func (o *Transaction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Transaction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Transaction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *Transaction) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetBlocktime

`func (o *Transaction) GetBlocktime() Timestamp`

GetBlocktime returns the Blocktime field if non-nil, zero value otherwise.

### GetBlocktimeOk

`func (o *Transaction) GetBlocktimeOk() (*Timestamp, bool)`

GetBlocktimeOk returns a tuple with the Blocktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocktime

`func (o *Transaction) SetBlocktime(v Timestamp)`

SetBlocktime sets Blocktime field to given value.

### HasBlocktime

`func (o *Transaction) HasBlocktime() bool`

HasBlocktime returns a boolean if a field has been set.

### GetMostImportantInstruction

`func (o *Transaction) GetMostImportantInstruction() TransactionMostImportantInstruction`

GetMostImportantInstruction returns the MostImportantInstruction field if non-nil, zero value otherwise.

### GetMostImportantInstructionOk

`func (o *Transaction) GetMostImportantInstructionOk() (*TransactionMostImportantInstruction, bool)`

GetMostImportantInstructionOk returns a tuple with the MostImportantInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostImportantInstruction

`func (o *Transaction) SetMostImportantInstruction(v TransactionMostImportantInstruction)`

SetMostImportantInstruction sets MostImportantInstruction field to given value.

### HasMostImportantInstruction

`func (o *Transaction) HasMostImportantInstruction() bool`

HasMostImportantInstruction returns a boolean if a field has been set.

### GetSmart

`func (o *Transaction) GetSmart() []TransactionSmart`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *Transaction) GetSmartOk() (*[]TransactionSmart, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *Transaction) SetSmart(v []TransactionSmart)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *Transaction) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetOndemand

`func (o *Transaction) GetOndemand() bool`

GetOndemand returns the Ondemand field if non-nil, zero value otherwise.

### GetOndemandOk

`func (o *Transaction) GetOndemandOk() (*bool, bool)`

GetOndemandOk returns a tuple with the Ondemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndemand

`func (o *Transaction) SetOndemand(v bool)`

SetOndemand sets Ondemand field to given value.

### HasOndemand

`func (o *Transaction) HasOndemand() bool`

HasOndemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


