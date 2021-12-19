# \TransactionApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchLatestTransactions**](TransactionApi.md#FetchLatestTransactions) | **Get** /latest-transactions | Fetch 50 latest transactions
[**FetchTransactionByHash**](TransactionApi.md#FetchTransactionByHash) | **Get** /transaction/{hash} | Fetch transaction by transaction hash
[**FetchTransactionHashesByAddress**](TransactionApi.md#FetchTransactionHashesByAddress) | **Get** /transaction-hashes/{address} | Fetch latest transaction hashes
[**FetchTransactionsByAddress**](TransactionApi.md#FetchTransactionsByAddress) | **Get** /transactions/{address} | Fetch latest transactions
[**FetchTransactionsByBlockNumber**](TransactionApi.md#FetchTransactionsByBlockNumber) | **Get** /block-transactions/{number} | Fetch latest transactions in a block



## FetchLatestTransactions

> []Transaction FetchLatestTransactions(ctx).Limit(limit).Cursor(cursor).Execute()

Fetch 50 latest transactions



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
    cursor := "cursor_example" // string | Transaction cursor (blocknumber,index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.FetchLatestTransactions(context.Background()).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.FetchLatestTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchLatestTransactions`: []Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.FetchLatestTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchLatestTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 100) | 
 **cursor** | **string** | Transaction cursor (blocknumber,index) | 

### Return type

[**[]Transaction**](Transaction.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTransactionByHash

> Transaction FetchTransactionByHash(ctx, hash).Execute()

Fetch transaction by transaction hash



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
    hash := "hash_example" // string | Transaction hash

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.FetchTransactionByHash(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.FetchTransactionByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTransactionByHash`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.FetchTransactionByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Transaction hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTransactionByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTransactionHashesByAddress

> []InlineResponse2001 FetchTransactionHashesByAddress(ctx, address).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch latest transaction hashes



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
    address := "address_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Transaction cursor (blocknumber,index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.FetchTransactionHashesByAddress(context.Background(), address).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.FetchTransactionHashesByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTransactionHashesByAddress`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.FetchTransactionHashesByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTransactionHashesByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Transaction cursor (blocknumber,index) | 

### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTransactionsByAddress

> []Transaction FetchTransactionsByAddress(ctx, address).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch latest transactions



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
    address := "address_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Transaction cursor (blocknumber,index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.FetchTransactionsByAddress(context.Background(), address).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.FetchTransactionsByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTransactionsByAddress`: []Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.FetchTransactionsByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTransactionsByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Transaction cursor (blocknumber,index) | 

### Return type

[**[]Transaction**](Transaction.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTransactionsByBlockNumber

> InlineResponse2002 FetchTransactionsByBlockNumber(ctx, number).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch latest transactions in a block



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
    number := int32(56) // int32 | Block number
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Transaction cursor (index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionApi.FetchTransactionsByBlockNumber(context.Background(), number).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.FetchTransactionsByBlockNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTransactionsByBlockNumber`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.FetchTransactionsByBlockNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**number** | **int32** | Block number | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchTransactionsByBlockNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Transaction cursor (index) | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

