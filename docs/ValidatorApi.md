# \ValidatorApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchValidatorByVotepubkey**](ValidatorApi.md#FetchValidatorByVotepubkey) | **Get** /validator/{votepubkey} | Fetch validator by vote pubkey
[**FetchValidatorDelegators**](ValidatorApi.md#FetchValidatorDelegators) | **Get** /validator/{votepubkey}/delegators | Fetch validator delegators
[**FetchValidatorHistory**](ValidatorApi.md#FetchValidatorHistory) | **Get** /validator/{votepubkey}/history | Fetch validator history
[**FetchValidatorSlots**](ValidatorApi.md#FetchValidatorSlots) | **Get** /validator/{nodepubkey}/slots | Fetch validator slots



## FetchValidatorByVotepubkey

> Validator FetchValidatorByVotepubkey(ctx, votepubkey).Execute()

Fetch validator by vote pubkey



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
    votepubkey := "votepubkey_example" // string | Validator pubkey

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidatorApi.FetchValidatorByVotepubkey(context.Background(), votepubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidatorApi.FetchValidatorByVotepubkey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchValidatorByVotepubkey`: Validator
    fmt.Fprintf(os.Stdout, "Response from `ValidatorApi.FetchValidatorByVotepubkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**votepubkey** | **string** | Validator pubkey | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidatorByVotepubkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Validator**](Validator.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchValidatorDelegators

> []StakeAccount FetchValidatorDelegators(ctx, votepubkey).Limit(limit).Offset(offset).Execute()

Fetch validator delegators



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
    votepubkey := "votepubkey_example" // string | Validator's vote pubkey
    limit := int32(56) // int32 | Result limit (max 1000) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidatorApi.FetchValidatorDelegators(context.Background(), votepubkey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidatorApi.FetchValidatorDelegators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchValidatorDelegators`: []StakeAccount
    fmt.Fprintf(os.Stdout, "Response from `ValidatorApi.FetchValidatorDelegators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**votepubkey** | **string** | Validator&#39;s vote pubkey | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidatorDelegatorsRequest struct via the builder pattern


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


## FetchValidatorHistory

> InlineResponse2003 FetchValidatorHistory(ctx, votepubkey).Execute()

Fetch validator history



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
    votepubkey := "votepubkey_example" // string | Validator's vote pubkey

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidatorApi.FetchValidatorHistory(context.Background(), votepubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidatorApi.FetchValidatorHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchValidatorHistory`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `ValidatorApi.FetchValidatorHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**votepubkey** | **string** | Validator&#39;s vote pubkey | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidatorHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchValidatorSlots

> []InlineResponse2004 FetchValidatorSlots(ctx, nodepubkey).Execute()

Fetch validator slots



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
    nodepubkey := "nodepubkey_example" // string | Validator's node pubkey

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ValidatorApi.FetchValidatorSlots(context.Background(), nodepubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidatorApi.FetchValidatorSlots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchValidatorSlots`: []InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `ValidatorApi.FetchValidatorSlots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodepubkey** | **string** | Validator&#39;s node pubkey | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchValidatorSlotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse2004**](InlineResponse2004.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

