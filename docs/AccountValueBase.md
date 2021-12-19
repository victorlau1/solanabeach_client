# AccountValueBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Balance** | Pointer to **int32** |  | [optional] 
**Executable** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**Address**](Address.md) |  | [optional] 
**RentEpoch** | Pointer to **int32** |  | [optional] 
**DataSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccountValueBase

`func NewAccountValueBase() *AccountValueBase`

NewAccountValueBase instantiates a new AccountValueBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountValueBaseWithDefaults

`func NewAccountValueBaseWithDefaults() *AccountValueBase`

NewAccountValueBaseWithDefaults instantiates a new AccountValueBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountValueBase) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountValueBase) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountValueBase) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountValueBase) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalance

`func (o *AccountValueBase) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountValueBase) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountValueBase) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountValueBase) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetExecutable

`func (o *AccountValueBase) GetExecutable() bool`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *AccountValueBase) GetExecutableOk() (*bool, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *AccountValueBase) SetExecutable(v bool)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *AccountValueBase) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetOwner

`func (o *AccountValueBase) GetOwner() Address`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountValueBase) GetOwnerOk() (*Address, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountValueBase) SetOwner(v Address)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccountValueBase) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRentEpoch

`func (o *AccountValueBase) GetRentEpoch() int32`

GetRentEpoch returns the RentEpoch field if non-nil, zero value otherwise.

### GetRentEpochOk

`func (o *AccountValueBase) GetRentEpochOk() (*int32, bool)`

GetRentEpochOk returns a tuple with the RentEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentEpoch

`func (o *AccountValueBase) SetRentEpoch(v int32)`

SetRentEpoch sets RentEpoch field to given value.

### HasRentEpoch

`func (o *AccountValueBase) HasRentEpoch() bool`

HasRentEpoch returns a boolean if a field has been set.

### GetDataSize

`func (o *AccountValueBase) GetDataSize() int32`

GetDataSize returns the DataSize field if non-nil, zero value otherwise.

### GetDataSizeOk

`func (o *AccountValueBase) GetDataSizeOk() (*int32, bool)`

GetDataSizeOk returns a tuple with the DataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSize

`func (o *AccountValueBase) SetDataSize(v int32)`

SetDataSize sets DataSize field to given value.

### HasDataSize

`func (o *AccountValueBase) HasDataSize() bool`

HasDataSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


