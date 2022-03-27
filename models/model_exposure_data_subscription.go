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


type ExposureDataSubscription struct {


	
	// simple type
 
	

	NotificationUri string `json:"notificationUri" yaml:"notificationUri" bson:"notificationUri" mapstructure:"NotificationUri"`
 
	

	MonitoredResourceUris []string `json:"monitoredResourceUris" yaml:"monitoredResourceUris" bson:"monitoredResourceUris" mapstructure:"MonitoredResourceUris"`
 
	

	Expiry *time.Time `json:"expiry,omitempty" yaml:"expiry" bson:"expiry" mapstructure:"Expiry"`
 
	

	SupportedFeatures *string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
