/*
Solanabeach Backend API

Solanabeach Backend REST API documentation.  ## Rate limiting Current API rate limit per IP is 100 requests per second.    ## Pagination Most of the endpoints returning array data support pagination. The API uses two types of pagination, depending on the returned data. Some endpoints support both, some are limited to just one type.    ## Supported pagination types ### Offset / limit Offset / limit pagination is the simplest form of pagination supported by the API. Offset parameter represents the number of results to skip before returning the data, and the limit parameter limits the number of returned results.   To prevent overloading the API, all limit params have a max value. Each API endpoint has its own max value.  ### Cursor Cursor pagination is more complex than the offset / limit one, but, it comes naturally for some blockchain data (such as blocks, transactions, token transfers, etc). Cursors contain data like blocknumber, transaction index, ... and they're described on their respective API endpoints. Limit parameter works exactly the same way as it does in the offset / limit pagination.  ## Authentication The public API uses a Bearer OAuth authentication method, and the API key should be provided in the `Authorization` header in each request. API keys are issued on request. 

API version: 0.0.1
Contact: andrej@vgng.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TokenApiService TokenApi service
type TokenApiService service

type ApiFetchTokenRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	pubkey string
}


func (r ApiFetchTokenRequest) Execute() (Token, *_nethttp.Response, error) {
	return r.ApiService.FetchTokenExecute(r)
}

/*
FetchToken Fetch single token

Fetch token by pubkey

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pubkey Token address
 @return ApiFetchTokenRequest
*/
func (a *TokenApiService) FetchToken(ctx _context.Context, pubkey string) ApiFetchTokenRequest {
	return ApiFetchTokenRequest{
		ApiService: a,
		ctx: ctx,
		pubkey: pubkey,
	}
}

// Execute executes the request
//  @return Token
func (a *TokenApiService) FetchTokenExecute(r ApiFetchTokenRequest) (Token, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Token
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token/{pubkey}"
	localVarPath = strings.Replace(localVarPath, "{"+"pubkey"+"}", _neturl.PathEscape(parameterToString(r.pubkey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchTokenAccountsRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	owner string
	limit *int32
	offset *int32
}

// Result limit (max 100)
func (r ApiFetchTokenAccountsRequest) Limit(limit int32) ApiFetchTokenAccountsRequest {
	r.limit = &limit
	return r
}
// Result offset
func (r ApiFetchTokenAccountsRequest) Offset(offset int32) ApiFetchTokenAccountsRequest {
	r.offset = &offset
	return r
}

func (r ApiFetchTokenAccountsRequest) Execute() ([]TokenHolder, *_nethttp.Response, error) {
	return r.ApiService.FetchTokenAccountsExecute(r)
}

/*
FetchTokenAccounts Fetch token accounts owned by owner

Fetch token accounts owned by owner ordered by amount

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Owner address
 @return ApiFetchTokenAccountsRequest
*/
func (a *TokenApiService) FetchTokenAccounts(ctx _context.Context, owner string) ApiFetchTokenAccountsRequest {
	return ApiFetchTokenAccountsRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
	}
}

// Execute executes the request
//  @return []TokenHolder
func (a *TokenApiService) FetchTokenAccountsExecute(r ApiFetchTokenAccountsRequest) ([]TokenHolder, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TokenHolder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchTokenAccounts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-accounts/{owner}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", _neturl.PathEscape(parameterToString(r.owner, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchTokenHoldersRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	mint string
	limit *int32
	offset *int32
}

// Result limit (max 100)
func (r ApiFetchTokenHoldersRequest) Limit(limit int32) ApiFetchTokenHoldersRequest {
	r.limit = &limit
	return r
}
// Result offset
func (r ApiFetchTokenHoldersRequest) Offset(offset int32) ApiFetchTokenHoldersRequest {
	r.offset = &offset
	return r
}

func (r ApiFetchTokenHoldersRequest) Execute() ([]TokenHolder, *_nethttp.Response, error) {
	return r.ApiService.FetchTokenHoldersExecute(r)
}

/*
FetchTokenHolders Fetch token holders

Fetch token holders ordered by amount. Please note that this endpoint only returns appropriate holders for known tokens (the ones in the official Solana token list).

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mint Mint address
 @return ApiFetchTokenHoldersRequest
*/
func (a *TokenApiService) FetchTokenHolders(ctx _context.Context, mint string) ApiFetchTokenHoldersRequest {
	return ApiFetchTokenHoldersRequest{
		ApiService: a,
		ctx: ctx,
		mint: mint,
	}
}

// Execute executes the request
//  @return []TokenHolder
func (a *TokenApiService) FetchTokenHoldersExecute(r ApiFetchTokenHoldersRequest) ([]TokenHolder, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TokenHolder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchTokenHolders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token/{mint}/holders"
	localVarPath = strings.Replace(localVarPath, "{"+"mint"+"}", _neturl.PathEscape(parameterToString(r.mint, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchTokenTransfersRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	mint string
	limit *int32
	offset *int32
	cursor *string
}

// Result limit (max 100)
func (r ApiFetchTokenTransfersRequest) Limit(limit int32) ApiFetchTokenTransfersRequest {
	r.limit = &limit
	return r
}
// Result offset
func (r ApiFetchTokenTransfersRequest) Offset(offset int32) ApiFetchTokenTransfersRequest {
	r.offset = &offset
	return r
}
// Token transfers cursor (blocknumber,txindex)
func (r ApiFetchTokenTransfersRequest) Cursor(cursor string) ApiFetchTokenTransfersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchTokenTransfersRequest) Execute() ([]TokenTransfer, *_nethttp.Response, error) {
	return r.ApiService.FetchTokenTransfersExecute(r)
}

/*
FetchTokenTransfers Fetch token transfers

Fetch token transfers ordered by id

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mint Mint address
 @return ApiFetchTokenTransfersRequest
*/
func (a *TokenApiService) FetchTokenTransfers(ctx _context.Context, mint string) ApiFetchTokenTransfersRequest {
	return ApiFetchTokenTransfersRequest{
		ApiService: a,
		ctx: ctx,
		mint: mint,
	}
}

// Execute executes the request
//  @return []TokenTransfer
func (a *TokenApiService) FetchTokenTransfersExecute(r ApiFetchTokenTransfersRequest) ([]TokenTransfer, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TokenTransfer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchTokenTransfers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token/{mint}/transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"mint"+"}", _neturl.PathEscape(parameterToString(r.mint, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchTokenTransfersByOwnerRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	owner string
	limit *int32
	offset *int32
	cursor *string
}

// Result limit (max 100)
func (r ApiFetchTokenTransfersByOwnerRequest) Limit(limit int32) ApiFetchTokenTransfersByOwnerRequest {
	r.limit = &limit
	return r
}
// Result offset
func (r ApiFetchTokenTransfersByOwnerRequest) Offset(offset int32) ApiFetchTokenTransfersByOwnerRequest {
	r.offset = &offset
	return r
}
// Token transfers cursor (blocknumber,txindex)
func (r ApiFetchTokenTransfersByOwnerRequest) Cursor(cursor string) ApiFetchTokenTransfersByOwnerRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFetchTokenTransfersByOwnerRequest) Execute() ([]TokenTransfer, *_nethttp.Response, error) {
	return r.ApiService.FetchTokenTransfersByOwnerExecute(r)
}

/*
FetchTokenTransfersByOwner Fetch token transfers by owner

Fetch token transfers by owner ordered by id

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Owner address
 @return ApiFetchTokenTransfersByOwnerRequest
*/
func (a *TokenApiService) FetchTokenTransfersByOwner(ctx _context.Context, owner string) ApiFetchTokenTransfersByOwnerRequest {
	return ApiFetchTokenTransfersByOwnerRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
	}
}

// Execute executes the request
//  @return []TokenTransfer
func (a *TokenApiService) FetchTokenTransfersByOwnerExecute(r ApiFetchTokenTransfersByOwnerRequest) ([]TokenTransfer, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TokenTransfer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchTokenTransfersByOwner")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/token-transfers/{owner}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", _neturl.PathEscape(parameterToString(r.owner, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchTokensRequest struct {
	ctx _context.Context
	ApiService *TokenApiService
	limit *int32
	offset *int32
	sort *string
	dir *string
}

// Result limit (max 100)
func (r ApiFetchTokensRequest) Limit(limit int32) ApiFetchTokensRequest {
	r.limit = &limit
	return r
}
// Result offset
func (r ApiFetchTokensRequest) Offset(offset int32) ApiFetchTokensRequest {
	r.offset = &offset
	return r
}
// Sort by param
func (r ApiFetchTokensRequest) Sort(sort string) ApiFetchTokensRequest {
	r.sort = &sort
	return r
}
// Sort direction param
func (r ApiFetchTokensRequest) Dir(dir string) ApiFetchTokensRequest {
	r.dir = &dir
	return r
}

func (r ApiFetchTokensRequest) Execute() (InlineResponse20017, *_nethttp.Response, error) {
	return r.ApiService.FetchTokensExecute(r)
}

/*
FetchTokens Fetch tokens

Fetch tokens ordered by holders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchTokensRequest
*/
func (a *TokenApiService) FetchTokens(ctx _context.Context) ApiFetchTokensRequest {
	return ApiFetchTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InlineResponse20017
func (a *TokenApiService) FetchTokensExecute(r ApiFetchTokensRequest) (InlineResponse20017, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20017
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenApiService.FetchTokens")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.dir != nil {
		localVarQueryParams.Add("dir", parameterToString(*r.dir, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}