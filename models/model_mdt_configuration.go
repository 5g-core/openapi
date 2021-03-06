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


type MdtConfiguration struct {


	
	// simple type
 
	

	JobType JobType `json:"jobType" yaml:"jobType" bson:"jobType" mapstructure:"JobType"`
 
	

	ReportType *ReportTypeMdt `json:"reportType,omitempty" yaml:"reportType" bson:"reportType" mapstructure:"ReportType"`
 
	

	AreaScope *AreaScope `json:"areaScope,omitempty" yaml:"areaScope" bson:"areaScope" mapstructure:"AreaScope"`
 
	

	MeasurementLteList *[]MeasurementLteForMdt `json:"measurementLteList,omitempty" yaml:"measurementLteList" bson:"measurementLteList" mapstructure:"MeasurementLteList"`
 
	

	MeasurementNrList *[]MeasurementNrForMdt `json:"measurementNrList,omitempty" yaml:"measurementNrList" bson:"measurementNrList" mapstructure:"MeasurementNrList"`
 
	

	SensorMeasurementList *[]SensorMeasurement `json:"sensorMeasurementList,omitempty" yaml:"sensorMeasurementList" bson:"sensorMeasurementList" mapstructure:"SensorMeasurementList"`
 
	

	ReportingTriggerList *[]ReportingTrigger `json:"reportingTriggerList,omitempty" yaml:"reportingTriggerList" bson:"reportingTriggerList" mapstructure:"ReportingTriggerList"`
 
	

	ReportInterval *ReportIntervalMdt `json:"reportInterval,omitempty" yaml:"reportInterval" bson:"reportInterval" mapstructure:"ReportInterval"`
 
	

	ReportIntervalNr *ReportIntervalNrMdt `json:"reportIntervalNr,omitempty" yaml:"reportIntervalNr" bson:"reportIntervalNr" mapstructure:"ReportIntervalNr"`
 
	

	ReportAmount *ReportAmountMdt `json:"reportAmount,omitempty" yaml:"reportAmount" bson:"reportAmount" mapstructure:"ReportAmount"`
 
	

	EventThresholdRsrp *int32 `json:"eventThresholdRsrp,omitempty" yaml:"eventThresholdRsrp" bson:"eventThresholdRsrp" mapstructure:"EventThresholdRsrp"`
 
	

	EventThresholdRsrpNr *int32 `json:"eventThresholdRsrpNr,omitempty" yaml:"eventThresholdRsrpNr" bson:"eventThresholdRsrpNr" mapstructure:"EventThresholdRsrpNr"`
 
	

	EventThresholdRsrq *int32 `json:"eventThresholdRsrq,omitempty" yaml:"eventThresholdRsrq" bson:"eventThresholdRsrq" mapstructure:"EventThresholdRsrq"`
 
	

	EventThresholdRsrqNr *int32 `json:"eventThresholdRsrqNr,omitempty" yaml:"eventThresholdRsrqNr" bson:"eventThresholdRsrqNr" mapstructure:"EventThresholdRsrqNr"`
 
	

	EventList *[]EventForMdt `json:"eventList,omitempty" yaml:"eventList" bson:"eventList" mapstructure:"EventList"`
 
	

	LoggingInterval *LoggingIntervalMdt `json:"loggingInterval,omitempty" yaml:"loggingInterval" bson:"loggingInterval" mapstructure:"LoggingInterval"`
 
	

	LoggingIntervalNr *LoggingIntervalNrMdt `json:"loggingIntervalNr,omitempty" yaml:"loggingIntervalNr" bson:"loggingIntervalNr" mapstructure:"LoggingIntervalNr"`
 
	

	LoggingDuration *LoggingDurationMdt `json:"loggingDuration,omitempty" yaml:"loggingDuration" bson:"loggingDuration" mapstructure:"LoggingDuration"`
 
	

	LoggingDurationNr *LoggingDurationNrMdt `json:"loggingDurationNr,omitempty" yaml:"loggingDurationNr" bson:"loggingDurationNr" mapstructure:"LoggingDurationNr"`
 
	

	PositioningMethod *PositioningMethodMdt `json:"positioningMethod,omitempty" yaml:"positioningMethod" bson:"positioningMethod" mapstructure:"PositioningMethod"`
 
	

	AddPositioningMethodList *[]PositioningMethodMdt `json:"addPositioningMethodList,omitempty" yaml:"addPositioningMethodList" bson:"addPositioningMethodList" mapstructure:"AddPositioningMethodList"`
 
	

	CollectionPeriodRmmLte *CollectionPeriodRmmLteMdt `json:"collectionPeriodRmmLte,omitempty" yaml:"collectionPeriodRmmLte" bson:"collectionPeriodRmmLte" mapstructure:"CollectionPeriodRmmLte"`
 
	

	CollectionPeriodRmmNr *CollectionPeriodRmmNrMdt `json:"collectionPeriodRmmNr,omitempty" yaml:"collectionPeriodRmmNr" bson:"collectionPeriodRmmNr" mapstructure:"CollectionPeriodRmmNr"`
 
	

	MeasurementPeriodLte *MeasurementPeriodLteMdt `json:"measurementPeriodLte,omitempty" yaml:"measurementPeriodLte" bson:"measurementPeriodLte" mapstructure:"MeasurementPeriodLte"`
 
	

	MdtAllowedPlmnIdList *[]PlmnId `json:"mdtAllowedPlmnIdList,omitempty" yaml:"mdtAllowedPlmnIdList" bson:"mdtAllowedPlmnIdList" mapstructure:"MdtAllowedPlmnIdList"`
 
	

	MbsfnAreaList *[]MbsfnArea `json:"mbsfnAreaList,omitempty" yaml:"mbsfnAreaList" bson:"mbsfnAreaList" mapstructure:"MbsfnAreaList"`
 
	

	InterFreqTargetList *[]InterFreqTargetInfo `json:"interFreqTargetList,omitempty" yaml:"interFreqTargetList" bson:"interFreqTargetList" mapstructure:"InterFreqTargetList"`
}
