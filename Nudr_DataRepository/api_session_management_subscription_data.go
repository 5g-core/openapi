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

type SessionManagementSubscriptionDataApiService service

/*
SessionManagementSubscriptionDataApiService Retrieves the Session Management subscription data of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId UE id
 * @param servingPlmnId PLMN ID
 * @param optional nil or *QuerySmDataParamOpts - Optional Parameters:
 * @param "SingleNssai" (optional.Interface of models.VarSnssai) -  single NSSAI
 * @param "Dnn" (optional.String) -  DNN
 * @param "Fields" (optional.Interface of []string) -  attributes to be retrieved
 * @param "SupportedFeatures" (optional.String) -  Supported Features
 * @param "IfNoneMatch" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.2
 * @param "IfModifiedSince" (optional.String) -  Validator for conditional requests, as described in RFC 7232, 3.3
@return []models.SessionManagementSubscriptionData
*/

type QuerySmDataParamOpts struct {
	SingleNssai       optional.Interface
	Dnn               optional.String
	Fields            optional.Interface
	SupportedFeatures optional.String
	IfNoneMatch       optional.String
	IfModifiedSince   optional.String
}

func (a *SessionManagementSubscriptionDataApiService) QuerySmData(ctx context.Context, ueId string, servingPlmnId string, localVarOptionals *QuerySmDataParamOpts) ([]models.SessionManagementSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []models.SessionManagementSubscriptionData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/{servingPlmnId}/provisioned-data/sm-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingPlmnId"+"}", fmt.Sprintf("%v", servingPlmnId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SingleNssai.IsSet() {
		localVarQueryParams.Add("single-nssai", openapi.ParameterToString(localVarOptionals.SingleNssai.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Dnn.IsSet() {
		localVarQueryParams.Add("dnn", openapi.ParameterToString(localVarOptionals.Dnn.Value(), ""))
	}
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

	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.IfModifiedSince.IsSet() {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(localVarOptionals.IfModifiedSince.Value(), "")
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
