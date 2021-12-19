# \TokenApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchToken**](TokenApi.md#FetchToken) | **Get** /token/{pubkey} | Fetch single token
[**FetchTokenAccounts**](TokenApi.md#FetchTokenAccounts) | **Get** /token-accounts/{owner} | Fetch token accounts owned by owner
[**FetchTokenHolders**](TokenApi.md#FetchTokenHolders) | **Get** /token/{mint}/holders | Fetch token holders
[**FetchTokenTransfers**](TokenApi.md#FetchTokenTransfers) | **Get** /token/{mint}/transfers | Fetch token transfers
[**FetchTokenTransfersByOwner**](TokenApi.md#FetchTokenTransfersByOwner) | **Get** /token-transfers/{owner} | Fetch token transfers by owner
[**FetchTokens**](TokenApi.md#FetchTokens) | **Get** /tokens | Fetch tokens



## FetchToken

> Token FetchToken(ctx, pubkey).Execute()

Fetch single token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pubkey := "pubkey_example" // string | Token address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchToken(context.Background(), pubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchToken`: Token
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Token address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Token**](Token.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenAccounts

> []TokenHolder FetchTokenAccounts(ctx, owner).Limit(limit).Offset(offset).Execute()

Fetch token accounts owned by owner



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | Owner address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchTokenAccounts(context.Background(), owner).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchTokenAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenAccounts`: []TokenHolder
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchTokenAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]TokenHolder**](TokenHolder.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenHolders

> []TokenHolder FetchTokenHolders(ctx, mint).Limit(limit).Offset(offset).Execute()

Fetch token holders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mint := "mint_example" // string | Mint address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchTokenHolders(context.Background(), mint).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchTokenHolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenHolders`: []TokenHolder
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchTokenHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mint** | **string** | Mint address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]TokenHolder**](TokenHolder.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenTransfers

> []TokenTransfer FetchTokenTransfers(ctx, mint).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch token transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    mint := "mint_example" // string | Mint address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Token transfers cursor (blocknumber,txindex) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchTokenTransfers(context.Background(), mint).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchTokenTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenTransfers`: []TokenTransfer
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchTokenTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mint** | **string** | Mint address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Token transfers cursor (blocknumber,txindex) | 

### Return type

[**[]TokenTransfer**](TokenTransfer.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenTransfersByOwner

> []TokenTransfer FetchTokenTransfersByOwner(ctx, owner).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch token transfers by owner



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    owner := "owner_example" // string | Owner address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Token transfers cursor (blocknumber,txindex) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchTokenTransfersByOwner(context.Background(), owner).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchTokenTransfersByOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenTransfersByOwner`: []TokenTransfer
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchTokenTransfersByOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**owner** | **string** | Owner address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenTransfersByOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Token transfers cursor (blocknumber,txindex) | 

### Return type

[**[]TokenTransfer**](TokenTransfer.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokens

> InlineResponse20017 FetchTokens(ctx).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()

Fetch tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    sort := "sort_example" // string | Sort by param (optional)
    dir := "dir_example" // string | Sort direction param (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenApi.FetchTokens(context.Background()).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenApi.FetchTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokens`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `TokenApi.FetchTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **sort** | **string** | Sort by param | 
 **dir** | **string** | Sort direction param | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

