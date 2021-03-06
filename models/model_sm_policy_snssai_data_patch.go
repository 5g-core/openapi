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


// SmPolicySnssaiDataPatch - Contains the SM policy data for a given subscriber and S-NSSAI.
type SmPolicySnssaiDataPatch struct {


	
	// simple type
 
	

	Snssai Snssai `json:"snssai" yaml:"snssai" bson:"snssai" mapstructure:"Snssai"`
 
	

	SmPolicyDnnData *map[string]SmPolicyDnnDataPatch `json:"smPolicyDnnData,omitempty" yaml:"smPolicyDnnData" bson:"smPolicyDnnData" mapstructure:"SmPolicyDnnData"`
}
