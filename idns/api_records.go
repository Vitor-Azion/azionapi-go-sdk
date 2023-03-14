/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RecordsApiService RecordsApi service
type RecordsApiService service

type ApiDeleteZoneRecordRequest struct {
	ctx        context.Context
	ApiService *RecordsApiService
	zoneId     int32
	recordId   int32
}

func (r ApiDeleteZoneRecordRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DeleteZoneRecordExecute(r)
}

/*
DeleteZoneRecord Remove an Intelligent DNS zone record

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param zoneId The hosted zone id
	@param recordId The zone record id
	@return ApiDeleteZoneRecordRequest
*/
func (a *RecordsApiService) DeleteZoneRecord(ctx context.Context, zoneId int32, recordId int32) ApiDeleteZoneRecordRequest {
	return ApiDeleteZoneRecordRequest{
		ApiService: a,
		ctx:        ctx,
		zoneId:     zoneId,
		recordId:   recordId,
	}
}

// Execute executes the request
//
//	@return string
func (a *RecordsApiService) DeleteZoneRecordExecute(r ApiDeleteZoneRecordRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.DeleteZoneRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intelligent_dns/{zone_id}/records/{record_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_id"+"}", url.PathEscape(parameterValueToString(r.zoneId, "zoneId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"record_id"+"}", url.PathEscape(parameterValueToString(r.recordId, "recordId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneId < 1 {
		return localVarReturnValue, nil, reportError("zoneId must be greater than 1")
	}
	if r.recordId < 1 {
		return localVarReturnValue, nil, reportError("recordId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetZoneRecordsRequest struct {
	ctx        context.Context
	ApiService *RecordsApiService
	zoneId     int32
}

func (r ApiGetZoneRecordsRequest) Execute() (*GetRecordsResponse, *http.Response, error) {
	return r.ApiService.GetZoneRecordsExecute(r)
}

/*
GetZoneRecords Get a collection of Intelligent DNS zone records

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param zoneId The hosted zone id
	@return ApiGetZoneRecordsRequest
*/
func (a *RecordsApiService) GetZoneRecords(ctx context.Context, zoneId int32) ApiGetZoneRecordsRequest {
	return ApiGetZoneRecordsRequest{
		ApiService: a,
		ctx:        ctx,
		zoneId:     zoneId,
	}
}

// Execute executes the request
//
//	@return GetRecordsResponse
func (a *RecordsApiService) GetZoneRecordsExecute(r ApiGetZoneRecordsRequest) (*GetRecordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRecordsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.GetZoneRecords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intelligent_dns/{zone_id}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_id"+"}", url.PathEscape(parameterValueToString(r.zoneId, "zoneId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneId < 1 {
		return localVarReturnValue, nil, reportError("zoneId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostZoneRecordRequest struct {
	ctx             context.Context
	ApiService      *RecordsApiService
	zoneId          int32
	recordPostOrPut *RecordPostOrPut
}

func (r ApiPostZoneRecordRequest) RecordPostOrPut(recordPostOrPut RecordPostOrPut) ApiPostZoneRecordRequest {
	r.recordPostOrPut = &recordPostOrPut
	return r
}

func (r ApiPostZoneRecordRequest) Execute() (*PostOrPutRecordResponse, *http.Response, error) {
	return r.ApiService.PostZoneRecordExecute(r)
}

/*
PostZoneRecord Create a new Intelligent DNS zone record

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param zoneId The hosted zone id
	@return ApiPostZoneRecordRequest
*/
func (a *RecordsApiService) PostZoneRecord(ctx context.Context, zoneId int32) ApiPostZoneRecordRequest {
	return ApiPostZoneRecordRequest{
		ApiService: a,
		ctx:        ctx,
		zoneId:     zoneId,
	}
}

// Execute executes the request
//
//	@return PostOrPutRecordResponse
func (a *RecordsApiService) PostZoneRecordExecute(r ApiPostZoneRecordRequest) (*PostOrPutRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PostOrPutRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.PostZoneRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intelligent_dns/{zone_id}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_id"+"}", url.PathEscape(parameterValueToString(r.zoneId, "zoneId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneId < 1 {
		return localVarReturnValue, nil, reportError("zoneId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordPostOrPut
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutZoneRecordRequest struct {
	ctx             context.Context
	ApiService      *RecordsApiService
	zoneId          int32
	recordId        int32
	recordPostOrPut *RecordPostOrPut
}

func (r ApiPutZoneRecordRequest) RecordPostOrPut(recordPostOrPut RecordPostOrPut) ApiPutZoneRecordRequest {
	r.recordPostOrPut = &recordPostOrPut
	return r
}

func (r ApiPutZoneRecordRequest) Execute() (*PostOrPutRecordResponse, *http.Response, error) {
	return r.ApiService.PutZoneRecordExecute(r)
}

/*
PutZoneRecord Update an Intelligent DNS zone record

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param zoneId The hosted zone id
	@param recordId The zone record id
	@return ApiPutZoneRecordRequest
*/
func (a *RecordsApiService) PutZoneRecord(ctx context.Context, zoneId int32, recordId int32) ApiPutZoneRecordRequest {
	return ApiPutZoneRecordRequest{
		ApiService: a,
		ctx:        ctx,
		zoneId:     zoneId,
		recordId:   recordId,
	}
}

// Execute executes the request
//
//	@return PostOrPutRecordResponse
func (a *RecordsApiService) PutZoneRecordExecute(r ApiPutZoneRecordRequest) (*PostOrPutRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PostOrPutRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsApiService.PutZoneRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/intelligent_dns/{zone_id}/records/{record_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_id"+"}", url.PathEscape(parameterValueToString(r.zoneId, "zoneId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"record_id"+"}", url.PathEscape(parameterValueToString(r.recordId, "recordId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.zoneId < 1 {
		return localVarReturnValue, nil, reportError("zoneId must be greater than 1")
	}
	if r.recordId < 1 {
		return localVarReturnValue, nil, reportError("recordId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordPostOrPut
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
