/*
Solanabeach Backend API

Solanabeach Backend REST API documentation.  ## Rate limiting Current API rate limit per IP is 100 requests per second.    ## Pagination Most of the endpoints returning array data support pagination. The API uses two types of pagination, depending on the returned data. Some endpoints support both, some are limited to just one type.    ## Supported pagination types ### Offset / limit Offset / limit pagination is the simplest form of pagination supported by the API. Offset parameter represents the number of results to skip before returning the data, and the limit parameter limits the number of returned results.   To prevent overloading the API, all limit params have a max value. Each API endpoint has its own max value.  ### Cursor Cursor pagination is more complex than the offset / limit one, but, it comes naturally for some blockchain data (such as blocks, transactions, token transfers, etc). Cursors contain data like blocknumber, transaction index, ... and they're described on their respective API endpoints. Limit parameter works exactly the same way as it does in the offset / limit pagination.  ## Authentication The public API uses a Bearer OAuth authentication method, and the API key should be provided in the `Authorization` header in each request. API keys are issued on request. 

API version: 0.0.1
Contact: andrej@vgng.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Epoch *int32 `json:"epoch,omitempty"`
	Foundation *float32 `json:"foundation,omitempty"`
	Total *float32 `json:"total,omitempty"`
	Validator *float32 `json:"validator,omitempty"`
}

// NewInlineResponse2007 instantiates a new InlineResponse2007 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007WithDefaults() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// GetEpoch returns the Epoch field value if set, zero value otherwise.
func (o *InlineResponse2007) GetEpoch() int32 {
	if o == nil || o.Epoch == nil {
		var ret int32
		return ret
	}
	return *o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetEpochOk() (*int32, bool) {
	if o == nil || o.Epoch == nil {
		return nil, false
	}
	return o.Epoch, true
}

// HasEpoch returns a boolean if a field has been set.
func (o *InlineResponse2007) HasEpoch() bool {
	if o != nil && o.Epoch != nil {
		return true
	}

	return false
}

// SetEpoch gets a reference to the given int32 and assigns it to the Epoch field.
func (o *InlineResponse2007) SetEpoch(v int32) {
	o.Epoch = &v
}

// GetFoundation returns the Foundation field value if set, zero value otherwise.
func (o *InlineResponse2007) GetFoundation() float32 {
	if o == nil || o.Foundation == nil {
		var ret float32
		return ret
	}
	return *o.Foundation
}

// GetFoundationOk returns a tuple with the Foundation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetFoundationOk() (*float32, bool) {
	if o == nil || o.Foundation == nil {
		return nil, false
	}
	return o.Foundation, true
}

// HasFoundation returns a boolean if a field has been set.
func (o *InlineResponse2007) HasFoundation() bool {
	if o != nil && o.Foundation != nil {
		return true
	}

	return false
}

// SetFoundation gets a reference to the given float32 and assigns it to the Foundation field.
func (o *InlineResponse2007) SetFoundation(v float32) {
	o.Foundation = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse2007) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse2007) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *InlineResponse2007) SetTotal(v float32) {
	o.Total = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *InlineResponse2007) GetValidator() float32 {
	if o == nil || o.Validator == nil {
		var ret float32
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetValidatorOk() (*float32, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *InlineResponse2007) HasValidator() bool {
	if o != nil && o.Validator != nil {
		return true
	}

	return false
}

// SetValidator gets a reference to the given float32 and assigns it to the Validator field.
func (o *InlineResponse2007) SetValidator(v float32) {
	o.Validator = &v
}

func (o InlineResponse2007) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Epoch != nil {
		toSerialize["epoch"] = o.Epoch
	}
	if o.Foundation != nil {
		toSerialize["foundation"] = o.Foundation
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007 struct {
	value *InlineResponse2007
	isSet bool
}

func (v NullableInlineResponse2007) Get() *InlineResponse2007 {
	return v.value
}

func (v *NullableInlineResponse2007) Set(val *InlineResponse2007) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007(val *InlineResponse2007) *NullableInlineResponse2007 {
	return &NullableInlineResponse2007{value: val, isSet: true}
}

func (v NullableInlineResponse2007) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


