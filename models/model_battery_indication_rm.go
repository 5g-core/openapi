/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data). The API version is defined in 3GPP TS 29.504. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type BatteryIndicationRm struct {

	BatteryInd bool `json:"batteryInd,omitempty" yaml:"batteryInd" bson:"batteryInd" mapstructure:"BatteryInd"`

	ReplaceableInd bool `json:"replaceableInd,omitempty" yaml:"replaceableInd" bson:"replaceableInd" mapstructure:"ReplaceableInd"`

	RechargeableInd bool `json:"rechargeableInd,omitempty" yaml:"rechargeableInd" bson:"rechargeableInd" mapstructure:"RechargeableInd"`
}
