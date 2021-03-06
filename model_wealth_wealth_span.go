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

// WealthWealthSpan struct for WealthWealthSpan
type WealthWealthSpan struct {
	From *int32 `json:"from,omitempty"`
	To *int32 `json:"to,omitempty"`
}

// NewWealthWealthSpan instantiates a new WealthWealthSpan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWealthWealthSpan() *WealthWealthSpan {
	this := WealthWealthSpan{}
	return &this
}

// NewWealthWealthSpanWithDefaults instantiates a new WealthWealthSpan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWealthWealthSpanWithDefaults() *WealthWealthSpan {
	this := WealthWealthSpan{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *WealthWealthSpan) GetFrom() int32 {
	if o == nil || o.From == nil {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WealthWealthSpan) GetFromOk() (*int32, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *WealthWealthSpan) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *WealthWealthSpan) SetFrom(v int32) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *WealthWealthSpan) GetTo() int32 {
	if o == nil || o.To == nil {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WealthWealthSpan) GetToOk() (*int32, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *WealthWealthSpan) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *WealthWealthSpan) SetTo(v int32) {
	o.To = &v
}

func (o WealthWealthSpan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableWealthWealthSpan struct {
	value *WealthWealthSpan
	isSet bool
}

func (v NullableWealthWealthSpan) Get() *WealthWealthSpan {
	return v.value
}

func (v *NullableWealthWealthSpan) Set(val *WealthWealthSpan) {
	v.value = val
	v.isSet = true
}

func (v NullableWealthWealthSpan) IsSet() bool {
	return v.isSet
}

func (v *NullableWealthWealthSpan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWealthWealthSpan(val *WealthWealthSpan) *NullableWealthWealthSpan {
	return &NullableWealthWealthSpan{value: val, isSet: true}
}

func (v NullableWealthWealthSpan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWealthWealthSpan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


