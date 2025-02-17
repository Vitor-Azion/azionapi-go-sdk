# Go API client for idns

Azion Intelligent DNS API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import idns "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), idns.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), idns.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), idns.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), idns.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.azionapi.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DNSSECApi* | [**GetZoneDnsSec**](docs/DNSSECApi.md#getzonednssec) | **Get** /intelligent_dns/{zone_id}/dnssec | Retrieve the DNSSEC zone status
*DNSSECApi* | [**PutZoneDnsSec**](docs/DNSSECApi.md#putzonednssec) | **Patch** /intelligent_dns/{zone_id}/dnssec | Update the DNSSEC zone
*RecordsApi* | [**DeleteZoneRecord**](docs/RecordsApi.md#deletezonerecord) | **Delete** /intelligent_dns/{zone_id}/records/{record_id} | Remove an Intelligent DNS zone record
*RecordsApi* | [**GetZoneRecords**](docs/RecordsApi.md#getzonerecords) | **Get** /intelligent_dns/{zone_id}/records | Get a collection of Intelligent DNS zone records
*RecordsApi* | [**PostZoneRecord**](docs/RecordsApi.md#postzonerecord) | **Post** /intelligent_dns/{zone_id}/records | Create a new Intelligent DNS zone record
*RecordsApi* | [**PutZoneRecord**](docs/RecordsApi.md#putzonerecord) | **Put** /intelligent_dns/{zone_id}/records/{record_id} | Update an Intelligent DNS zone record
*ZonesApi* | [**DeleteZone**](docs/ZonesApi.md#deletezone) | **Delete** /intelligent_dns/{zone_id} | Remove an Intelligent DNS hosted zone
*ZonesApi* | [**GetZone**](docs/ZonesApi.md#getzone) | **Get** /intelligent_dns/{zone_id} | Get an Intelligent DNS hosted zone
*ZonesApi* | [**GetZones**](docs/ZonesApi.md#getzones) | **Get** /intelligent_dns | Get a collection of Intelligent DNS zones
*ZonesApi* | [**PostZone**](docs/ZonesApi.md#postzone) | **Post** /intelligent_dns | Add a new Intelligent DNS zone
*ZonesApi* | [**PutZone**](docs/ZonesApi.md#putzone) | **Put** /intelligent_dns/{zone_id} | Update an Intelligent DNS hosted zone


## Documentation For Models

 - [DnsSec](docs/DnsSec.md)
 - [DnsSecDelegationSigner](docs/DnsSecDelegationSigner.md)
 - [DnsSecDelegationSignerDigestType](docs/DnsSecDelegationSignerDigestType.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorsResponse](docs/ErrorsResponse.md)
 - [GetOrPatchDnsSecResponse](docs/GetOrPatchDnsSecResponse.md)
 - [GetRecordsResponse](docs/GetRecordsResponse.md)
 - [GetRecordsResponseResults](docs/GetRecordsResponseResults.md)
 - [GetZoneResponse](docs/GetZoneResponse.md)
 - [GetZonesResponse](docs/GetZonesResponse.md)
 - [GetZonesResponseLinks](docs/GetZonesResponseLinks.md)
 - [PostOrPutRecordResponse](docs/PostOrPutRecordResponse.md)
 - [PostOrPutZoneResponse](docs/PostOrPutZoneResponse.md)
 - [RecordGet](docs/RecordGet.md)
 - [RecordPostOrPut](docs/RecordPostOrPut.md)
 - [Zone](docs/Zone.md)


## Documentation For Authorization



### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



