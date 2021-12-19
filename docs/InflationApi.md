# \InflationApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInflationInfo**](InflationApi.md#FetchInflationInfo) | **Get** /inflation | Fetch inflation
[**FetchSupplyInfo**](InflationApi.md#FetchSupplyInfo) | **Get** /supply | Fetch supply



## FetchInflationInfo

> InlineResponse2007 FetchInflationInfo(ctx).Execute()

Fetch inflation



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
    resp, r, err := api_client.InflationApi.FetchInflationInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InflationApi.FetchInflationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchInflationInfo`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `InflationApi.FetchInflationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchInflationInfoRequest struct via the builder pattern


### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSupplyInfo

> InlineResponse2008 FetchSupplyInfo(ctx).Execute()

Fetch supply



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
    resp, r, err := api_client.InflationApi.FetchSupplyInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InflationApi.FetchSupplyInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSupplyInfo`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `InflationApi.FetchSupplyInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchSupplyInfoRequest struct via the builder pattern


### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

