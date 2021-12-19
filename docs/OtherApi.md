# \OtherApi

All URIs are relative to *https://api.solanabeach.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAlias**](OtherApi.md#FetchAlias) | **Get** /alias | Fetch alias
[**FetchClusterTime**](OtherApi.md#FetchClusterTime) | **Get** /cluster-time | Fetch current cluster time
[**FetchEpochHistory**](OtherApi.md#FetchEpochHistory) | **Get** /epoch-history | Fetch epoch history
[**FetchHealth**](OtherApi.md#FetchHealth) | **Get** /health | Fetch performance info
[**FetchMarketChardData**](OtherApi.md#FetchMarketChardData) | **Get** /market-chart-data | Fetch market chart data
[**FetchNetworkStatus**](OtherApi.md#FetchNetworkStatus) | **Get** /network-status | Fetch network status
[**FetchNonValidators**](OtherApi.md#FetchNonValidators) | **Get** /non-validators | Fetch non validator nodes
[**FetchStakeHistory**](OtherApi.md#FetchStakeHistory) | **Get** /stake-history | Fetch stake history
[**FetchStakingAPY**](OtherApi.md#FetchStakingAPY) | **Get** /staking-apy | Fetch staking APY



## FetchAlias

> InlineResponse20014 FetchAlias(ctx).Execute()

Fetch alias



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
    resp, r, err := api_client.OtherApi.FetchAlias(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAlias`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchAlias`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAliasRequest struct via the builder pattern


### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchClusterTime

> int32 FetchClusterTime(ctx).Execute()

Fetch current cluster time



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
    resp, r, err := api_client.OtherApi.FetchClusterTime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchClusterTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchClusterTime`: int32
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchClusterTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchClusterTimeRequest struct via the builder pattern


### Return type

**int32**

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEpochHistory

> InlineResponse20013 FetchEpochHistory(ctx).Execute()

Fetch epoch history



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
    resp, r, err := api_client.OtherApi.FetchEpochHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchEpochHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEpochHistory`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchEpochHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchEpochHistoryRequest struct via the builder pattern


### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchHealth

> InlineResponse20010 FetchHealth(ctx).Execute()

Fetch performance info



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
    resp, r, err := api_client.OtherApi.FetchHealth(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchHealth`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchHealthRequest struct via the builder pattern


### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMarketChardData

> []InlineResponse20016 FetchMarketChardData(ctx).Execute()

Fetch market chart data



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
    resp, r, err := api_client.OtherApi.FetchMarketChardData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchMarketChardData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMarketChardData`: []InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchMarketChardData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchMarketChardDataRequest struct via the builder pattern


### Return type

[**[]InlineResponse20016**](InlineResponse20016.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkStatus

> InlineResponse20011 FetchNetworkStatus(ctx).Execute()

Fetch network status



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
    resp, r, err := api_client.OtherApi.FetchNetworkStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchNetworkStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNetworkStatus`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchNetworkStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchNetworkStatusRequest struct via the builder pattern


### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNonValidators

> []InlineResponse20015 FetchNonValidators(ctx).Execute()

Fetch non validator nodes



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
    resp, r, err := api_client.OtherApi.FetchNonValidators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchNonValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNonValidators`: []InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchNonValidators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchNonValidatorsRequest struct via the builder pattern


### Return type

[**[]InlineResponse20015**](InlineResponse20015.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStakeHistory

> []InlineResponse2009 FetchStakeHistory(ctx).Limit(limit).Offset(offset).Execute()

Fetch stake history



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
    limit := int32(56) // int32 | Result limit (max 512) (optional)
    offset := int32(56) // int32 | Result offset (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OtherApi.FetchStakeHistory(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchStakeHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStakeHistory`: []InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchStakeHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchStakeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Result limit (max 512) | 
 **offset** | **int32** | Result offset | 

### Return type

[**[]InlineResponse2009**](InlineResponse2009.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStakingAPY

> InlineResponse20012 FetchStakingAPY(ctx).Execute()

Fetch staking APY



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
    resp, r, err := api_client.OtherApi.FetchStakingAPY(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherApi.FetchStakingAPY``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStakingAPY`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OtherApi.FetchStakingAPY`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFetchStakingAPYRequest struct via the builder pattern


### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[apiAuth](../README.md#apiAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

