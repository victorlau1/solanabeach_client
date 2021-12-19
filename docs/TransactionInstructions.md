# TransactionInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parsed** | Pointer to **map[string]interface{}** | Depends on the type of instruction | [optional] 
**Raw** | Pointer to [**TransactionRaw**](TransactionRaw.md) |  | [optional] 

## Methods

### NewTransactionInstructions

`func NewTransactionInstructions() *TransactionInstructions`

NewTransactionInstructions instantiates a new TransactionInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInstructionsWithDefaults

`func NewTransactionInstructionsWithDefaults() *TransactionInstructions`

NewTransactionInstructionsWithDefaults instantiates a new TransactionInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParsed

`func (o *TransactionInstructions) GetParsed() map[string]interface{}`

GetParsed returns the Parsed field if non-nil, zero value otherwise.

### GetParsedOk

`func (o *TransactionInstructions) GetParsedOk() (*map[string]interface{}, bool)`

GetParsedOk returns a tuple with the Parsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsed

`func (o *TransactionInstructions) SetParsed(v map[string]interface{})`

SetParsed sets Parsed field to given value.

### HasParsed

`func (o *TransactionInstructions) HasParsed() bool`

HasParsed returns a boolean if a field has been set.

### GetRaw

`func (o *TransactionInstructions) GetRaw() TransactionRaw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *TransactionInstructions) GetRawOk() (*TransactionRaw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *TransactionInstructions) SetRaw(v TransactionRaw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *TransactionInstructions) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


