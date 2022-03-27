/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (

)


type Pp5gVnGroupProfileData struct {


	
	// simple type

	// A map (list of key-value pairs where external VN group identifier serves as key) of AllowedMtcProviderInfo lists. In addition to defined external VN group identifier, the key value \"ALL\" may be used to identify a map entry which contains a list of AllowedMtcProviderInfo that are allowed operating all the external group identifiers. 
	

	AllowedMtcProviders *map[string][]AllowedMtcProviderInfo `json:"allowedMtcProviders,omitempty" yaml:"allowedMtcProviders" bson:"allowedMtcProviders" mapstructure:"AllowedMtcProviders"`
 
	

	SupportedFeatures *string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
