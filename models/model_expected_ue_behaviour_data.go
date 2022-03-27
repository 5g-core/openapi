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


type ExpectedUeBehaviourData struct {


	
	// simple type
 
	

	StationaryIndication *StationaryIndication `json:"stationaryIndication,omitempty" yaml:"stationaryIndication" bson:"stationaryIndication" mapstructure:"StationaryIndication"`
 
	

	CommunicationDurationTime *int32 `json:"communicationDurationTime,omitempty" yaml:"communicationDurationTime" bson:"communicationDurationTime" mapstructure:"CommunicationDurationTime"`
 
	

	PeriodicTime *int32 `json:"periodicTime,omitempty" yaml:"periodicTime" bson:"periodicTime" mapstructure:"PeriodicTime"`
 
	

	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty" yaml:"scheduledCommunicationTime" bson:"scheduledCommunicationTime" mapstructure:"ScheduledCommunicationTime"`
 
	

	ScheduledCommunicationType *ScheduledCommunicationType `json:"scheduledCommunicationType,omitempty" yaml:"scheduledCommunicationType" bson:"scheduledCommunicationType" mapstructure:"ScheduledCommunicationType"`

	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G. 
	

	ExpectedUmts *[]LocationArea `json:"expectedUmts,omitempty" yaml:"expectedUmts" bson:"expectedUmts" mapstructure:"ExpectedUmts"`
 
	

	TrafficProfile *TrafficProfile `json:"trafficProfile,omitempty" yaml:"trafficProfile" bson:"trafficProfile" mapstructure:"TrafficProfile"`
 
	

	BatteryIndication *BatteryIndication `json:"batteryIndication,omitempty" yaml:"batteryIndication" bson:"batteryIndication" mapstructure:"BatteryIndication"`
 
	

	ValidityTime *time.Time `json:"validityTime,omitempty" yaml:"validityTime" bson:"validityTime" mapstructure:"ValidityTime"`
}
