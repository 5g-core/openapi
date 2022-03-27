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


type ExternalUnrelatedClass struct {


	
	// simple type
 
	

	LcsClientExternals *[]LcsClientExternal `json:"lcsClientExternals,omitempty" yaml:"lcsClientExternals" bson:"lcsClientExternals" mapstructure:"LcsClientExternals"`
 
	

	AfExternals *[]AfExternal `json:"afExternals,omitempty" yaml:"afExternals" bson:"afExternals" mapstructure:"AfExternals"`
 
	

	LcsClientGroupExternals *[]LcsClientGroupExternal `json:"lcsClientGroupExternals,omitempty" yaml:"lcsClientGroupExternals" bson:"lcsClientGroupExternals" mapstructure:"LcsClientGroupExternals"`
}
