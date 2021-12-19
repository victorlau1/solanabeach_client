# \AccountApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAccount**](AccountApi.md#FetchAccount) | **Get** /account/{pubkey} | Fetch account data
[**FetchAccountSerumInstructions**](AccountApi.md#FetchAccountSerumInstructions) | **Get** /account/{pubkey}/serum-instructions | Fetch account serum instructions
[**FetchAccountSerumOrders**](AccountApi.md#FetchAccountSerumOrders) | **Get** /account/{pubkey}/serum-orders | Fetch account serum orders
[**FetchAccountSwapInstructions**](AccountApi.md#FetchAccountSwapInstructions) | **Get** /account/{pubkey}/swap-instructions | Fetch account swap instructions
[**FetchAccountTokenTransfers**](AccountApi.md#FetchAccountTokenTransfers) | **Get** /account/{pubkey}/token-transfers | Fetch account token transfers
[**FetchAccountTokens**](AccountApi.md#FetchAccountTokens) | **Get** /account/{pubkey}/tokens | Fetch account tokens
[**FetchAccountTransactions**](AccountApi.md#FetchAccountTransactions) | **Get** /account/{pubkey}/transactions | Fetch account transactions
[**FetchAccounts**](AccountApi.md#FetchAccounts) | **Get** /accounts | Fetch accounts
[**FetchStakeAccountRewards**](AccountApi.md#FetchStakeAccountRewards) | **Get** /account/{stake_pubkey}/stake-rewards | Fetch stake account rewards
[**FetchStakeAccounts**](AccountApi.md#FetchStakeAccounts) | **Get** /account/{pubkey}/stakes | Fetch stake accounts
[**FetchWealthMetrics**](AccountApi.md#FetchWealthMetrics) | **Get** /wealth | Fetch wealth distribution metrics



## FetchAccount

> Account FetchAccount(ctx, pubkey).Execute()

Fetch account data



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
    pubkey := "pubkey_example" // string | Account address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccount(context.Background(), pubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountSerumInstructions

> []AccountSerumInstruction FetchAccountSerumInstructions(ctx, pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch account serum instructions



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Serum instruction cursor (blocknumber,txindex,index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountSerumInstructions(context.Background(), pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountSerumInstructions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountSerumInstructions`: []AccountSerumInstruction
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountSerumInstructions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountSerumInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Serum instruction cursor (blocknumber,txindex,index) | 

### Return type

[**[]AccountSerumInstruction**](AccountSerumInstruction.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountSerumOrders

> []AccountSerumOrder FetchAccountSerumOrders(ctx, pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch account serum orders



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Serum orders cursor (blocknumber,txindex) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountSerumOrders(context.Background(), pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountSerumOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountSerumOrders`: []AccountSerumOrder
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountSerumOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountSerumOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Serum orders cursor (blocknumber,txindex) | 

### Return type

[**[]AccountSerumOrder**](AccountSerumOrder.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountSwapInstructions

> []AccountSwapInstruction FetchAccountSwapInstructions(ctx, pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch account swap instructions



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Swap instruction cursor (blocknumber,txindex,index) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountSwapInstructions(context.Background(), pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountSwapInstructions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountSwapInstructions`: []AccountSwapInstruction
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountSwapInstructions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountSwapInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Swap instruction cursor (blocknumber,txindex,index) | 

### Return type

[**[]AccountSwapInstruction**](AccountSwapInstruction.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountTokenTransfers

> []AccountTokenTransfer FetchAccountTokenTransfers(ctx, pubkey).Limit(limit).Offset(offset).Cursor(cursor).Inner(inner).Execute()

Fetch account token transfers



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Token transfers cursor (blocknumber,txindex) (optional)
    inner := "inner_example" // bool | Filter inner instructions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountTokenTransfers(context.Background(), pubkey).Limit(limit).Offset(offset).Cursor(cursor).Inner(inner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountTokenTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountTokenTransfers`: []AccountTokenTransfer
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountTokenTransfers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountTokenTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Token transfers cursor (blocknumber,txindex) | 
 **inner** | **bool** | Filter inner instructions | 

### Return type

[**[]AccountTokenTransfer**](AccountTokenTransfer.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountTokens

> []AccountToken FetchAccountTokens(ctx, pubkey).Limit(limit).Offset(offset).Execute()

Fetch account tokens



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountTokens(context.Background(), pubkey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountTokens`: []AccountToken
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]AccountToken**](AccountToken.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountTransactions

> []Transaction FetchAccountTransactions(ctx, pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()

Fetch account transactions



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 1000) (optional)
    offset := int32(56) // int32 | Result offset (optional)
    cursor := "cursor_example" // string | Transaction cursor (blocknumber,txindex) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccountTransactions(context.Background(), pubkey).Limit(limit).Offset(offset).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccountTransactions`: []Transaction
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 1000) | 
 **offset** | **int32** | Result offset | 
 **cursor** | **string** | Transaction cursor (blocknumber,txindex) | 

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


## FetchAccounts

> []SimpleAccount FetchAccounts(ctx).Limit(limit).Offset(offset).Execute()

Fetch accounts



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
    limit := int32(56) // int32 | Result limit (max 1000) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchAccounts(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccounts`: []SimpleAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 1000) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]SimpleAccount**](SimpleAccount.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStakeAccountRewards

> []StakeAccountReward FetchStakeAccountRewards(ctx, stakePubkey).Cursor(cursor).Execute()

Fetch stake account rewards



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
    stakePubkey := "stakePubkey_example" // string | Stake account address
    cursor := int32(56) // int32 | Epoch cursor (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchStakeAccountRewards(context.Background(), stakePubkey).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchStakeAccountRewards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStakeAccountRewards`: []StakeAccountReward
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchStakeAccountRewards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stakePubkey** | **string** | Stake account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchStakeAccountRewardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **int32** | Epoch cursor | 

### Return type

[**[]StakeAccountReward**](StakeAccountReward.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStakeAccounts

> []StakeAccount FetchStakeAccounts(ctx, pubkey).Limit(limit).Offset(offset).Execute()

Fetch stake accounts



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
    pubkey := "pubkey_example" // string | Account address
    limit := int32(56) // int32 | Result limit (max 1000) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.FetchStakeAccounts(context.Background(), pubkey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchStakeAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStakeAccounts`: []StakeAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchStakeAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Account address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchStakeAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 1000) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]StakeAccount**](StakeAccount.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWealthMetrics

> Wealth FetchWealthMetrics(ctx).Execute()

Fetch wealth distribution metrics



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
    resp, r, err := api_client.AccountApi.FetchWealthMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.FetchWealthMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWealthMetrics`: Wealth
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.FetchWealthMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchWealthMetricsRequest struct via the builder pattern


### Return type

[**Wealth**](Wealth.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

