/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RequestedQos struct {
	Var5qi int32  `json:"5qi" yaml:"5qi" bson:"5qi" mapstructure:"Var5qi"`
	GbrUl  string `json:"gbrUl,omitempty" yaml:"gbrUl" bson:"gbrUl" mapstructure:"GbrUl"`
	GbrDl  string `json:"gbrDl,omitempty" yaml:"gbrDl" bson:"gbrDl" mapstructure:"GbrDl"`
}