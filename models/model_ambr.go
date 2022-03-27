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


type Ambr struct {


	
	// simple type
 
	

	Uplink string `json:"uplink" yaml:"uplink" bson:"uplink" mapstructure:"Uplink"`
 
	

	Downlink string `json:"downlink" yaml:"downlink" bson:"downlink" mapstructure:"Downlink"`
}
