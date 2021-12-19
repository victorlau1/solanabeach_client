# InlineResponse20015

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pubkey** | Pointer to [**NonvalidatorsPubkey**](NonvalidatorsPubkey.md) |  | [optional] 
**FeatureSet** | Pointer to **int32** |  | [optional] 
**Gossip** | Pointer to **string** |  | [optional] 
**Rpc** | Pointer to **string** |  | [optional] 
**ShredVersion** | Pointer to **int32** |  | [optional] 
**Tpu** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**NonvalidatorsLocation**](NonvalidatorsLocation.md) |  | [optional] 
**Asn** | Pointer to [**NonvalidatorsAsn**](NonvalidatorsAsn.md) |  | [optional] 
**Validator** | Pointer to **bool** |  | [optional] 

## Methods

### NewInlineResponse20015

`func NewInlineResponse20015() *InlineResponse20015`

NewInlineResponse20015 instantiates a new InlineResponse20015 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20015WithDefaults

`func NewInlineResponse20015WithDefaults() *InlineResponse20015`

NewInlineResponse20015WithDefaults instantiates a new InlineResponse20015 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubkey

`func (o *InlineResponse20015) GetPubkey() NonvalidatorsPubkey`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *InlineResponse20015) GetPubkeyOk() (*NonvalidatorsPubkey, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *InlineResponse20015) SetPubkey(v NonvalidatorsPubkey)`

SetPubkey sets Pubkey field to given value.

### HasPubkey

`func (o *InlineResponse20015) HasPubkey() bool`

HasPubkey returns a boolean if a field has been set.

### GetFeatureSet

`func (o *InlineResponse20015) GetFeatureSet() int32`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *InlineResponse20015) GetFeatureSetOk() (*int32, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *InlineResponse20015) SetFeatureSet(v int32)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *InlineResponse20015) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetGossip

`func (o *InlineResponse20015) GetGossip() string`

GetGossip returns the Gossip field if non-nil, zero value otherwise.

### GetGossipOk

`func (o *InlineResponse20015) GetGossipOk() (*string, bool)`

GetGossipOk returns a tuple with the Gossip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGossip

`func (o *InlineResponse20015) SetGossip(v string)`

SetGossip sets Gossip field to given value.

### HasGossip

`func (o *InlineResponse20015) HasGossip() bool`

HasGossip returns a boolean if a field has been set.

### GetRpc

`func (o *InlineResponse20015) GetRpc() string`

GetRpc returns the Rpc field if non-nil, zero value otherwise.

### GetRpcOk

`func (o *InlineResponse20015) GetRpcOk() (*string, bool)`

GetRpcOk returns a tuple with the Rpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpc

`func (o *InlineResponse20015) SetRpc(v string)`

SetRpc sets Rpc field to given value.

### HasRpc

`func (o *InlineResponse20015) HasRpc() bool`

HasRpc returns a boolean if a field has been set.

### GetShredVersion

`func (o *InlineResponse20015) GetShredVersion() int32`

GetShredVersion returns the ShredVersion field if non-nil, zero value otherwise.

### GetShredVersionOk

`func (o *InlineResponse20015) GetShredVersionOk() (*int32, bool)`

GetShredVersionOk returns a tuple with the ShredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShredVersion

`func (o *InlineResponse20015) SetShredVersion(v int32)`

SetShredVersion sets ShredVersion field to given value.

### HasShredVersion

`func (o *InlineResponse20015) HasShredVersion() bool`

HasShredVersion returns a boolean if a field has been set.

### GetTpu

`func (o *InlineResponse20015) GetTpu() string`

GetTpu returns the Tpu field if non-nil, zero value otherwise.

### GetTpuOk

`func (o *InlineResponse20015) GetTpuOk() (*string, bool)`

GetTpuOk returns a tuple with the Tpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpu

`func (o *InlineResponse20015) SetTpu(v string)`

SetTpu sets Tpu field to given value.

### HasTpu

`func (o *InlineResponse20015) HasTpu() bool`

HasTpu returns a boolean if a field has been set.

### GetVersion

`func (o *InlineResponse20015) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineResponse20015) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineResponse20015) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineResponse20015) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLocation

`func (o *InlineResponse20015) GetLocation() NonvalidatorsLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse20015) GetLocationOk() (*NonvalidatorsLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse20015) SetLocation(v NonvalidatorsLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse20015) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAsn

`func (o *InlineResponse20015) GetAsn() NonvalidatorsAsn`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *InlineResponse20015) GetAsnOk() (*NonvalidatorsAsn, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *InlineResponse20015) SetAsn(v NonvalidatorsAsn)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *InlineResponse20015) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetValidator

`func (o *InlineResponse20015) GetValidator() bool`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *InlineResponse20015) GetValidatorOk() (*bool, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *InlineResponse20015) SetValidator(v bool)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *InlineResponse20015) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


