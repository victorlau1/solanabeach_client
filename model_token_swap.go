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

// TokenSwap struct for TokenSwap
type TokenSwap struct {
	Pubkey *Address `json:"pubkey,omitempty"`
	Name *string `json:"name,omitempty"`
	Poolmint *Address `json:"poolmint,omitempty"`
	Tokenamint *Address `json:"tokenamint,omitempty"`
	Tokenbmint *Address `json:"tokenbmint,omitempty"`
	Tokena *Address `json:"tokena,omitempty"`
	Tokenb *Address `json:"tokenb,omitempty"`
	Tokenprogram *Address `json:"tokenprogram,omitempty"`
	Meta *TokenSwapMeta `json:"meta,omitempty"`
}

// NewTokenSwap instantiates a new TokenSwap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenSwap() *TokenSwap {
	this := TokenSwap{}
	return &this
}

// NewTokenSwapWithDefaults instantiates a new TokenSwap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenSwapWithDefaults() *TokenSwap {
	this := TokenSwap{}
	return &this
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *TokenSwap) GetPubkey() Address {
	if o == nil || o.Pubkey == nil {
		var ret Address
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetPubkeyOk() (*Address, bool) {
	if o == nil || o.Pubkey == nil {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *TokenSwap) HasPubkey() bool {
	if o != nil && o.Pubkey != nil {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given Address and assigns it to the Pubkey field.
func (o *TokenSwap) SetPubkey(v Address) {
	o.Pubkey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenSwap) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenSwap) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenSwap) SetName(v string) {
	o.Name = &v
}

// GetPoolmint returns the Poolmint field value if set, zero value otherwise.
func (o *TokenSwap) GetPoolmint() Address {
	if o == nil || o.Poolmint == nil {
		var ret Address
		return ret
	}
	return *o.Poolmint
}

// GetPoolmintOk returns a tuple with the Poolmint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetPoolmintOk() (*Address, bool) {
	if o == nil || o.Poolmint == nil {
		return nil, false
	}
	return o.Poolmint, true
}

// HasPoolmint returns a boolean if a field has been set.
func (o *TokenSwap) HasPoolmint() bool {
	if o != nil && o.Poolmint != nil {
		return true
	}

	return false
}

// SetPoolmint gets a reference to the given Address and assigns it to the Poolmint field.
func (o *TokenSwap) SetPoolmint(v Address) {
	o.Poolmint = &v
}

// GetTokenamint returns the Tokenamint field value if set, zero value otherwise.
func (o *TokenSwap) GetTokenamint() Address {
	if o == nil || o.Tokenamint == nil {
		var ret Address
		return ret
	}
	return *o.Tokenamint
}

// GetTokenamintOk returns a tuple with the Tokenamint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetTokenamintOk() (*Address, bool) {
	if o == nil || o.Tokenamint == nil {
		return nil, false
	}
	return o.Tokenamint, true
}

// HasTokenamint returns a boolean if a field has been set.
func (o *TokenSwap) HasTokenamint() bool {
	if o != nil && o.Tokenamint != nil {
		return true
	}

	return false
}

// SetTokenamint gets a reference to the given Address and assigns it to the Tokenamint field.
func (o *TokenSwap) SetTokenamint(v Address) {
	o.Tokenamint = &v
}

// GetTokenbmint returns the Tokenbmint field value if set, zero value otherwise.
func (o *TokenSwap) GetTokenbmint() Address {
	if o == nil || o.Tokenbmint == nil {
		var ret Address
		return ret
	}
	return *o.Tokenbmint
}

// GetTokenbmintOk returns a tuple with the Tokenbmint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetTokenbmintOk() (*Address, bool) {
	if o == nil || o.Tokenbmint == nil {
		return nil, false
	}
	return o.Tokenbmint, true
}

// HasTokenbmint returns a boolean if a field has been set.
func (o *TokenSwap) HasTokenbmint() bool {
	if o != nil && o.Tokenbmint != nil {
		return true
	}

	return false
}

// SetTokenbmint gets a reference to the given Address and assigns it to the Tokenbmint field.
func (o *TokenSwap) SetTokenbmint(v Address) {
	o.Tokenbmint = &v
}

// GetTokena returns the Tokena field value if set, zero value otherwise.
func (o *TokenSwap) GetTokena() Address {
	if o == nil || o.Tokena == nil {
		var ret Address
		return ret
	}
	return *o.Tokena
}

// GetTokenaOk returns a tuple with the Tokena field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetTokenaOk() (*Address, bool) {
	if o == nil || o.Tokena == nil {
		return nil, false
	}
	return o.Tokena, true
}

// HasTokena returns a boolean if a field has been set.
func (o *TokenSwap) HasTokena() bool {
	if o != nil && o.Tokena != nil {
		return true
	}

	return false
}

// SetTokena gets a reference to the given Address and assigns it to the Tokena field.
func (o *TokenSwap) SetTokena(v Address) {
	o.Tokena = &v
}

// GetTokenb returns the Tokenb field value if set, zero value otherwise.
func (o *TokenSwap) GetTokenb() Address {
	if o == nil || o.Tokenb == nil {
		var ret Address
		return ret
	}
	return *o.Tokenb
}

// GetTokenbOk returns a tuple with the Tokenb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetTokenbOk() (*Address, bool) {
	if o == nil || o.Tokenb == nil {
		return nil, false
	}
	return o.Tokenb, true
}

// HasTokenb returns a boolean if a field has been set.
func (o *TokenSwap) HasTokenb() bool {
	if o != nil && o.Tokenb != nil {
		return true
	}

	return false
}

// SetTokenb gets a reference to the given Address and assigns it to the Tokenb field.
func (o *TokenSwap) SetTokenb(v Address) {
	o.Tokenb = &v
}

// GetTokenprogram returns the Tokenprogram field value if set, zero value otherwise.
func (o *TokenSwap) GetTokenprogram() Address {
	if o == nil || o.Tokenprogram == nil {
		var ret Address
		return ret
	}
	return *o.Tokenprogram
}

// GetTokenprogramOk returns a tuple with the Tokenprogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetTokenprogramOk() (*Address, bool) {
	if o == nil || o.Tokenprogram == nil {
		return nil, false
	}
	return o.Tokenprogram, true
}

// HasTokenprogram returns a boolean if a field has been set.
func (o *TokenSwap) HasTokenprogram() bool {
	if o != nil && o.Tokenprogram != nil {
		return true
	}

	return false
}

// SetTokenprogram gets a reference to the given Address and assigns it to the Tokenprogram field.
func (o *TokenSwap) SetTokenprogram(v Address) {
	o.Tokenprogram = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TokenSwap) GetMeta() TokenSwapMeta {
	if o == nil || o.Meta == nil {
		var ret TokenSwapMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenSwap) GetMetaOk() (*TokenSwapMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TokenSwap) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given TokenSwapMeta and assigns it to the Meta field.
func (o *TokenSwap) SetMeta(v TokenSwapMeta) {
	o.Meta = &v
}

func (o TokenSwap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pubkey != nil {
		toSerialize["pubkey"] = o.Pubkey
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Poolmint != nil {
		toSerialize["poolmint"] = o.Poolmint
	}
	if o.Tokenamint != nil {
		toSerialize["tokenamint"] = o.Tokenamint
	}
	if o.Tokenbmint != nil {
		toSerialize["tokenbmint"] = o.Tokenbmint
	}
	if o.Tokena != nil {
		toSerialize["tokena"] = o.Tokena
	}
	if o.Tokenb != nil {
		toSerialize["tokenb"] = o.Tokenb
	}
	if o.Tokenprogram != nil {
		toSerialize["tokenprogram"] = o.Tokenprogram
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableTokenSwap struct {
	value *TokenSwap
	isSet bool
}

func (v NullableTokenSwap) Get() *TokenSwap {
	return v.value
}

func (v *NullableTokenSwap) Set(val *TokenSwap) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenSwap) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenSwap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenSwap(val *TokenSwap) *NullableTokenSwap {
	return &NullableTokenSwap{value: val, isSet: true}
}

func (v NullableTokenSwap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenSwap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

