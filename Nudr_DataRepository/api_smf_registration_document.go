/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudr_DataRepository

import (
	"github.com/5g-core/openapi"
	"github.com/5g-core/openapi/models"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SMFRegistrationDocumentApiService service

/*
SMFRegistrationDocumentApiService To create an individual SMF context data of a UE in the UDR
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId UE id
 * @param pduSessionId PDU session id
 * @param optional nil or *CreateSmfContextNon3gppParamOpts - Optional Parameters:
 * @param "SmfRegistration" (optional.Interface of SmfRegistration) -
*/

type CreateSmfContextNon3gppParamOpts struct {
	SmfRegistration optional.Interface
}

func (a *SMFRegistrationDocumentApiService) CreateSmfContextNon3gpp(ctx context.Context, ueId string, pduSessionId int32, localVarOptionals *CreateSmfContextNon3gppParamOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", fmt.Sprintf("%v", pduSessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	if localVarOptionals != nil && localVarOptionals.SmfRegistration.IsSet() {
		localVarOptionalSmfRegistration, localVarOptionalSmfRegistrationok := localVarOptionals.SmfRegistration.Value().(models.SmfRegistration)
		if !localVarOptionalSmfRegistrationok {
			return nil, openapi.ReportError("smfRegistration should be SmfRegistration")
		}
		localVarPostBody = &localVarOptionalSmfRegistration
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 201:
		return localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	}
}

/*
SMFRegistrationDocumentApiService To remove an individual SMF context data of a UE the UDR
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId UE id
 * @param pduSessionId PDU session id
*/

func (a *SMFRegistrationDocumentApiService) DeleteSmfContext(ctx context.Context, ueId string, pduSessionId string) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", fmt.Sprintf("%v", pduSessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	}
}

/*
SMFRegistrationDocumentApiService Retrieves the individual SMF registration of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId UE id
 * @param pduSessionId PDU session id
 * @param optional nil or *QuerySmfRegistrationParamOpts - Optional Parameters:
 * @param "Fields" (optional.Interface of []string) -  attributes to be retrieved
 * @param "SupportedFeatures" (optional.String) -  Supported Features
@return models.SmfRegistration
*/

type QuerySmfRegistrationParamOpts struct {
	Fields            optional.Interface
	SupportedFeatures optional.String
}

func (a *SMFRegistrationDocumentApiService) QuerySmfRegistration(ctx context.Context, ueId string, pduSessionId string, localVarOptionals *QuerySmfRegistrationParamOpts) (models.SmfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.SmfRegistration
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pduSessionId"+"}", fmt.Sprintf("%v", pduSessionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", openapi.ParameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SupportedFeatures.IsSet() {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(localVarOptionals.SupportedFeatures.Value(), ""))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, apiError
	}
}
