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


type BackupAmfInfo struct {


	
	// simple type
 
	

	BackupAmf string `json:"backupAmf" yaml:"backupAmf" bson:"backupAmf" mapstructure:"BackupAmf"`
 
	

	GuamiList *[]Guami `json:"guamiList,omitempty" yaml:"guamiList" bson:"guamiList" mapstructure:"GuamiList"`
}
