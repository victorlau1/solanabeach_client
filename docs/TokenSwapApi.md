# \TokenSwapApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTokenSwap**](TokenSwapApi.md#FetchTokenSwap) | **Get** /token-swap/{pubkey} | Fetch token swap
[**FetchTokenSwaps**](TokenSwapApi.md#FetchTokenSwaps) | **Get** /token-swaps | Fetch token swaps
[**FetchTokenSwapsByMint**](TokenSwapApi.md#FetchTokenSwapsByMint) | **Get** /token-swaps/{mint} | Fetch token swaps by mint



## FetchTokenSwap

> TokenSwap FetchTokenSwap(ctx, pubkey).Execute()

Fetch token swap



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
    pubkey := "pubkey_example" // string | Token swap address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenSwapApi.FetchTokenSwap(context.Background(), pubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenSwapApi.FetchTokenSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenSwap`: TokenSwap
    fmt.Fprintf(os.Stdout, "Response from `TokenSwapApi.FetchTokenSwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Token swap address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenSwap**](TokenSwap.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenSwaps

> InlineResponse2006 FetchTokenSwaps(ctx).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()

Fetch token swaps



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
    resp, r, err := api_client.TokenSwapApi.FetchTokenSwaps(context.Background()).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenSwapApi.FetchTokenSwaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenSwaps`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `TokenSwapApi.FetchTokenSwaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenSwapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **sort** | **string** | Sort by param | 
 **dir** | **string** | Sort direction param | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTokenSwapsByMint

> []TokenSwap FetchTokenSwapsByMint(ctx, mint).Limit(limit).Offset(offset).Execute()

Fetch token swaps by mint



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
    resp, r, err := api_client.TokenSwapApi.FetchTokenSwapsByMint(context.Background(), mint).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenSwapApi.FetchTokenSwapsByMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTokenSwapsByMint`: []TokenSwap
    fmt.Fprintf(os.Stdout, "Response from `TokenSwapApi.FetchTokenSwapsByMint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mint** | **string** | Mint address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTokenSwapsByMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]TokenSwap**](TokenSwap.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

