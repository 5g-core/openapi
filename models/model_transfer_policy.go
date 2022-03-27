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


// TransferPolicy - Describes a transfer policy.
type TransferPolicy struct {


	
	// simple type
 
	

	MaxBitRateDl *string `json:"maxBitRateDl,omitempty" yaml:"maxBitRateDl" bson:"maxBitRateDl" mapstructure:"MaxBitRateDl"`
 
	

	MaxBitRateUl *string `json:"maxBitRateUl,omitempty" yaml:"maxBitRateUl" bson:"maxBitRateUl" mapstructure:"MaxBitRateUl"`

	// Indicates a rating group for the recommended time window. 
	

	RatingGroup int32 `json:"ratingGroup" yaml:"ratingGroup" bson:"ratingGroup" mapstructure:"RatingGroup"`
 
	

	RecTimeInt TimeWindow `json:"recTimeInt" yaml:"recTimeInt" bson:"recTimeInt" mapstructure:"RecTimeInt"`

	// Contains an identity of a transfer policy. 
	

	TransPolicyId int32 `json:"transPolicyId" yaml:"transPolicyId" bson:"transPolicyId" mapstructure:"TransPolicyId"`
}