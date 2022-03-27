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

	"time"
)


// UsageMonDataLimit - Contains usage monitoring control data for a subscriber.
type UsageMonDataLimit struct {


	
	// simple type
 
	

	LimitId string `json:"limitId" yaml:"limitId" bson:"limitId" mapstructure:"LimitId"`
 
	

	Scopes *map[string]UsageMonDataScope `json:"scopes,omitempty" yaml:"scopes" bson:"scopes" mapstructure:"Scopes"`
 
	

	UmLevel *UsageMonLevel `json:"umLevel,omitempty" yaml:"umLevel" bson:"umLevel" mapstructure:"UmLevel"`
 
	

	StartDate *time.Time `json:"startDate,omitempty" yaml:"startDate" bson:"startDate" mapstructure:"StartDate"`
 
	

	EndDate *time.Time `json:"endDate,omitempty" yaml:"endDate" bson:"endDate" mapstructure:"EndDate"`
 
	

	UsageLimit *UsageThreshold `json:"usageLimit,omitempty" yaml:"usageLimit" bson:"usageLimit" mapstructure:"UsageLimit"`
 
	

	ResetPeriod *TimePeriod `json:"resetPeriod,omitempty" yaml:"resetPeriod" bson:"resetPeriod" mapstructure:"ResetPeriod"`
}
