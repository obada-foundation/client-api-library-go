/*
 * OBADA Client Helper API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Contact: techops@obada.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ObitApiService ObitApi service
type ObitApiService service

type ApiDownloadObitFromChainRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	obitDid *ObitDid
}

func (r ApiDownloadObitFromChainRequest) ObitDid(obitDid ObitDid) ApiDownloadObitFromChainRequest {
	r.obitDid = &obitDid
	return r
}

func (r ApiDownloadObitFromChainRequest) Execute() (ClientObitResponse, *_nethttp.Response, error) {
	return r.ApiService.DownloadObitFromChainExecute(r)
}

/*
 * DownloadObitFromChain Download Obit from Blockchain
 * Downloads the Obit information from the blockchain to the client.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDownloadObitFromChainRequest
 */
func (a *ObitApiService) DownloadObitFromChain(ctx _context.Context) ApiDownloadObitFromChainRequest {
	return ApiDownloadObitFromChainRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ClientObitResponse
 */
func (a *ObitApiService) DownloadObitFromChainExecute(r ApiDownloadObitFromChainRequest) (ClientObitResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientObitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.DownloadObitFromChain")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server/obit/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.obitDid
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchObitFromChainRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	obitDid string
}


func (r ApiFetchObitFromChainRequest) Execute() (BlockChainObitResponse, *_nethttp.Response, error) {
	return r.ApiService.FetchObitFromChainExecute(r)
}

/*
 * FetchObitFromChain Get Obit From Blockchain
 * Retrieves Obit information from blockchain but does not download it.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param obitDid Required.
 * @return ApiFetchObitFromChainRequest
 */
func (a *ObitApiService) FetchObitFromChain(ctx _context.Context, obitDid string) ApiFetchObitFromChainRequest {
	return ApiFetchObitFromChainRequest{
		ApiService: a,
		ctx: ctx,
		obitDid: obitDid,
	}
}

/*
 * Execute executes the request
 * @return BlockChainObitResponse
 */
func (a *ObitApiService) FetchObitFromChainExecute(r ApiFetchObitFromChainRequest) (BlockChainObitResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BlockChainObitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.FetchObitFromChain")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server/obit/{obit_did}"
	localVarPath = strings.Replace(localVarPath, "{"+"obit_did"+"}", _neturl.PathEscape(parameterToString(r.obitDid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGenerateObitDefRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	manufacturer *string
	partNumber *string
	serialNumber *string
}

func (r ApiGenerateObitDefRequest) Manufacturer(manufacturer string) ApiGenerateObitDefRequest {
	r.manufacturer = &manufacturer
	return r
}
func (r ApiGenerateObitDefRequest) PartNumber(partNumber string) ApiGenerateObitDefRequest {
	r.partNumber = &partNumber
	return r
}
func (r ApiGenerateObitDefRequest) SerialNumber(serialNumber string) ApiGenerateObitDefRequest {
	r.serialNumber = &serialNumber
	return r
}

func (r ApiGenerateObitDefRequest) Execute() (ObitDefinitionResponse, *_nethttp.Response, error) {
	return r.ApiService.GenerateObitDefExecute(r)
}

/*
 * GenerateObitDef Generate Obit Definition
 * Returns the Obit Definition for a given device_id, part_number and serial_number input.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGenerateObitDefRequest
 */
func (a *ObitApiService) GenerateObitDef(ctx _context.Context) ApiGenerateObitDefRequest {
	return ApiGenerateObitDefRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ObitDefinitionResponse
 */
func (a *ObitApiService) GenerateObitDefExecute(r ApiGenerateObitDefRequest) (ObitDefinitionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ObitDefinitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.GenerateObitDef")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/obit/generate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.manufacturer == nil {
		return localVarReturnValue, nil, reportError("manufacturer is required and must be specified")
	}
	if r.partNumber == nil {
		return localVarReturnValue, nil, reportError("partNumber is required and must be specified")
	}
	if r.serialNumber == nil {
		return localVarReturnValue, nil, reportError("serialNumber is required and must be specified")
	}

	localVarQueryParams.Add("manufacturer", parameterToString(*r.manufacturer, ""))
	localVarQueryParams.Add("part_number", parameterToString(*r.partNumber, ""))
	localVarQueryParams.Add("serial_number", parameterToString(*r.serialNumber, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientObitRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	obitDid string
}


func (r ApiGetClientObitRequest) Execute() (ClientObitResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientObitExecute(r)
}

/*
 * GetClientObit Get Client Obit
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param obitDid Required.
 * @return ApiGetClientObitRequest
 */
func (a *ObitApiService) GetClientObit(ctx _context.Context, obitDid string) ApiGetClientObitRequest {
	return ApiGetClientObitRequest{
		ApiService: a,
		ctx: ctx,
		obitDid: obitDid,
	}
}

/*
 * Execute executes the request
 * @return ClientObitResponse
 */
func (a *ObitApiService) GetClientObitExecute(r ApiGetClientObitRequest) (ClientObitResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientObitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.GetClientObit")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client/obit/{obit_did}"
	localVarPath = strings.Replace(localVarPath, "{"+"obit_did"+"}", _neturl.PathEscape(parameterToString(r.obitDid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSaveClientObitRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	localObit *LocalObit
}

func (r ApiSaveClientObitRequest) LocalObit(localObit LocalObit) ApiSaveClientObitRequest {
	r.localObit = &localObit
	return r
}

func (r ApiSaveClientObitRequest) Execute() (ClientObitResponse, *_nethttp.Response, error) {
	return r.ApiService.SaveClientObitExecute(r)
}

/*
 * SaveClientObit Save Client Obit
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSaveClientObitRequest
 */
func (a *ObitApiService) SaveClientObit(ctx _context.Context) ApiSaveClientObitRequest {
	return ApiSaveClientObitRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ClientObitResponse
 */
func (a *ObitApiService) SaveClientObitExecute(r ApiSaveClientObitRequest) (ClientObitResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientObitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.SaveClientObit")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/client/obit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.localObit
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUploadObitRequest struct {
	ctx _context.Context
	ApiService *ObitApiService
	obitDid *ObitDid
}

func (r ApiUploadObitRequest) ObitDid(obitDid ObitDid) ApiUploadObitRequest {
	r.obitDid = &obitDid
	return r
}

func (r ApiUploadObitRequest) Execute() (BaseResponse, *_nethttp.Response, error) {
	return r.ApiService.UploadObitExecute(r)
}

/*
 * UploadObit Upload Obit to Blockchain
 * Uploads Obit from client to the Blockchain
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUploadObitRequest
 */
func (a *ObitApiService) UploadObit(ctx _context.Context) ApiUploadObitRequest {
	return ApiUploadObitRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return BaseResponse
 */
func (a *ObitApiService) UploadObitExecute(r ApiUploadObitRequest) (BaseResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObitApiService.UploadObit")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/server/obit/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.obitDid
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
