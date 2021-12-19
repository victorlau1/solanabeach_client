# \MarketApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMarket**](MarketApi.md#FetchMarket) | **Get** /market/{pubkey} | Fetch market
[**FetchMarkets**](MarketApi.md#FetchMarkets) | **Get** /markets | Fetch markets
[**FetchMarketsByBaseMint**](MarketApi.md#FetchMarketsByBaseMint) | **Get** /markets/base/{basemint} | Fetch markets by base mint
[**FetchMarketsByQuoteMint**](MarketApi.md#FetchMarketsByQuoteMint) | **Get** /markets/quote/{quotemint} | Fetch markets by quote mint



## FetchMarket

> Market FetchMarket(ctx, pubkey).Execute()

Fetch market



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
    pubkey := "pubkey_example" // string | Market address

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FetchMarket(context.Background(), pubkey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FetchMarket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMarket`: Market
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FetchMarket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkey** | **string** | Market address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMarketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Market**](Market.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMarkets

> InlineResponse2005 FetchMarkets(ctx).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()

Fetch markets



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
    resp, r, err := api_client.MarketApi.FetchMarkets(context.Background()).Limit(limit).Offset(offset).Sort(sort).Dir(dir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FetchMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMarkets`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FetchMarkets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchMarketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 
 **sort** | **string** | Sort by param | 
 **dir** | **string** | Sort direction param | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMarketsByBaseMint

> []Market FetchMarketsByBaseMint(ctx, basemint).Limit(limit).Offset(offset).Execute()

Fetch markets by base mint



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
    basemint := "basemint_example" // string | Base mint address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FetchMarketsByBaseMint(context.Background(), basemint).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FetchMarketsByBaseMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMarketsByBaseMint`: []Market
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FetchMarketsByBaseMint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**basemint** | **string** | Base mint address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMarketsByBaseMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]Market**](Market.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMarketsByQuoteMint

> []Market FetchMarketsByQuoteMint(ctx, quotemint).Limit(limit).Offset(offset).Execute()

Fetch markets by quote mint



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
    quotemint := "quotemint_example" // string | Quote mint address
    limit := int32(56) // int32 | Result limit (max 100) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketApi.FetchMarketsByQuoteMint(context.Background(), quotemint).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketApi.FetchMarketsByQuoteMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMarketsByQuoteMint`: []Market
    fmt.Fprintf(os.Stdout, "Response from `MarketApi.FetchMarketsByQuoteMint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotemint** | **string** | Quote mint address | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMarketsByQuoteMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Result limit (max 100) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]Market**](Market.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

