# \BlockApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchBlockByBlockHash**](BlockApi.md#FetchBlockByBlockHash) | **Get** /block-hash/{hash} | Fetch block by block hash
[**FetchBlockByBlockNumber**](BlockApi.md#FetchBlockByBlockNumber) | **Get** /block/{number} | Fetch block by block number
[**FetchLatestBlocks**](BlockApi.md#FetchLatestBlocks) | **Get** /latest-blocks | Fetch 50 latest blocks
[**FetchTopPrograms**](BlockApi.md#FetchTopPrograms) | **Get** /top-programs | Fetch top program stats for last 1000 blocks



## FetchBlockByBlockHash

> Block FetchBlockByBlockHash(ctx, hash).Execute()

Fetch block by block hash



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
    hash := "hash_example" // string | Block hash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.FetchBlockByBlockHash(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.FetchBlockByBlockHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBlockByBlockHash`: Block
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.FetchBlockByBlockHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Block hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchBlockByBlockHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBlockByBlockNumber

> Block FetchBlockByBlockNumber(ctx, number).Execute()

Fetch block by block number



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
    number := "number_example" // string | Block number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.FetchBlockByBlockNumber(context.Background(), number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.FetchBlockByBlockNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBlockByBlockNumber`: Block
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.FetchBlockByBlockNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **string** | Block number | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchBlockByBlockNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLatestBlocks

> []Block FetchLatestBlocks(ctx).Limit(limit).Cursor(cursor).Execute()

Fetch 50 latest blocks



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
    cursor := "cursor_example" // string | Block cursor (blocknumber) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.FetchLatestBlocks(context.Background()).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.FetchLatestBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchLatestBlocks`: []Block
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.FetchLatestBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchLatestBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 100) | 
 **cursor** | **string** | Block cursor (blocknumber) | 

### Return type

[**[]Block**](Block.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTopPrograms

> []InlineResponse200 FetchTopPrograms(ctx).Execute()

Fetch top program stats for last 1000 blocks



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlockApi.FetchTopPrograms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.FetchTopPrograms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTopPrograms`: []InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `BlockApi.FetchTopPrograms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTopProgramsRequest struct via the builder pattern


### Return type

[**[]InlineResponse200**](InlineResponse200.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

