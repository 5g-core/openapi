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


type ServiceName string
// enum type
// List of ServiceName 
const (
	SERVICENAME_NNRF_NFM ServiceName = "nnrf-nfm"
	SERVICENAME_NNRF_DISC ServiceName = "nnrf-disc"
	SERVICENAME_NNRF_OAUTH2 ServiceName = "nnrf-oauth2"
	SERVICENAME_NUDM_SDM ServiceName = "nudm-sdm"
	SERVICENAME_NUDM_UECM ServiceName = "nudm-uecm"
	SERVICENAME_NUDM_UEAU ServiceName = "nudm-ueau"
	SERVICENAME_NUDM_EE ServiceName = "nudm-ee"
	SERVICENAME_NUDM_PP ServiceName = "nudm-pp"
	SERVICENAME_NUDM_NIDDAU ServiceName = "nudm-niddau"
	SERVICENAME_NUDM_MT ServiceName = "nudm-mt"
	SERVICENAME_NAMF_COMM ServiceName = "namf-comm"
	SERVICENAME_NAMF_EVTS ServiceName = "namf-evts"
	SERVICENAME_NAMF_MT ServiceName = "namf-mt"
	SERVICENAME_NAMF_LOC ServiceName = "namf-loc"
	SERVICENAME_NSMF_PDUSESSION ServiceName = "nsmf-pdusession"
	SERVICENAME_NSMF_EVENT_EXPOSURE ServiceName = "nsmf-event-exposure"
	SERVICENAME_NSMF_NIDD ServiceName = "nsmf-nidd"
	SERVICENAME_NAUSF_AUTH ServiceName = "nausf-auth"
	SERVICENAME_NAUSF_SORPROTECTION ServiceName = "nausf-sorprotection"
	SERVICENAME_NAUSF_UPUPROTECTION ServiceName = "nausf-upuprotection"
	SERVICENAME_NNEF_PFDMANAGEMENT ServiceName = "nnef-pfdmanagement"
	SERVICENAME_NNEF_SMCONTEXT ServiceName = "nnef-smcontext"
	SERVICENAME_NNEF_EVENTEXPOSURE ServiceName = "nnef-eventexposure"
	SERVICENAME_NPCF_AM_POLICY_CONTROL ServiceName = "npcf-am-policy-control"
	SERVICENAME_NPCF_SMPOLICYCONTROL ServiceName = "npcf-smpolicycontrol"
	SERVICENAME_NPCF_POLICYAUTHORIZATION ServiceName = "npcf-policyauthorization"
	SERVICENAME_NPCF_BDTPOLICYCONTROL ServiceName = "npcf-bdtpolicycontrol"
	SERVICENAME_NPCF_EVENTEXPOSURE ServiceName = "npcf-eventexposure"
	SERVICENAME_NPCF_UE_POLICY_CONTROL ServiceName = "npcf-ue-policy-control"
	SERVICENAME_NSMSF_SMS ServiceName = "nsmsf-sms"
	SERVICENAME_NNSSF_NSSELECTION ServiceName = "nnssf-nsselection"
	SERVICENAME_NNSSF_NSSAIAVAILABILITY ServiceName = "nnssf-nssaiavailability"
	SERVICENAME_NUDR_DR ServiceName = "nudr-dr"
	SERVICENAME_NUDR_GROUP_ID_MAP ServiceName = "nudr-group-id-map"
	SERVICENAME_NLMF_LOC ServiceName = "nlmf-loc"
	SERVICENAME_N5G_EIR_EIC ServiceName = "n5g-eir-eic"
	SERVICENAME_NBSF_MANAGEMENT ServiceName = "nbsf-management"
	SERVICENAME_NCHF_SPENDINGLIMITCONTROL ServiceName = "nchf-spendinglimitcontrol"
	SERVICENAME_NCHF_CONVERGEDCHARGING ServiceName = "nchf-convergedcharging"
	SERVICENAME_NCHF_OFFLINEONLYCHARGING ServiceName = "nchf-offlineonlycharging"
	SERVICENAME_NNWDAF_EVENTSSUBSCRIPTION ServiceName = "nnwdaf-eventssubscription"
	SERVICENAME_NNWDAF_ANALYTICSINFO ServiceName = "nnwdaf-analyticsinfo"
	SERVICENAME_NGMLC_LOC ServiceName = "ngmlc-loc"
	SERVICENAME_NUCMF_PROVISIONING ServiceName = "nucmf-provisioning"
	SERVICENAME_NUCMF_UECAPABILITYMANAGEMENT ServiceName = "nucmf-uecapabilitymanagement"
	SERVICENAME_NHSS_SDM ServiceName = "nhss-sdm"
	SERVICENAME_NHSS_UECM ServiceName = "nhss-uecm"
	SERVICENAME_NHSS_UEAU ServiceName = "nhss-ueau"
	SERVICENAME_NHSS_EE ServiceName = "nhss-ee"
	SERVICENAME_NHSS_IMS_SDM ServiceName = "nhss-ims-sdm"
	SERVICENAME_NHSS_IMS_UECM ServiceName = "nhss-ims-uecm"
	SERVICENAME_NHSS_IMS_UEAU ServiceName = "nhss-ims-ueau"
	SERVICENAME_NSEPP_TELESCOPIC ServiceName = "nsepp-telescopic"
	SERVICENAME_NSORAF_SOR ServiceName = "nsoraf-sor"
	SERVICENAME_NSPAF_SECURED_PACKET ServiceName = "nspaf-secured-packet"
	SERVICENAME_NUDSF_DR ServiceName = "nudsf-dr"
	SERVICENAME_NNSSAAF_NSSAA ServiceName = "nnssaaf-nssaa"
)
