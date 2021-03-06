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


// ResourceItem - Identifies a subscription to policy data change notification when the change occurs in a fragment (subset of resource data) of a given resource.
type ResourceItem struct {


	
	// simple type
 
	

	MonResourceUri string `json:"monResourceUri" yaml:"monResourceUri" bson:"monResourceUri" mapstructure:"MonResourceUri"`
 
	

	Items []string `json:"items" yaml:"items" bson:"items" mapstructure:"Items"`
}
