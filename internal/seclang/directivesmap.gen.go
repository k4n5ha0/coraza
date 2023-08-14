// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by internal/seclang/generator DO NOT EDIT.

package seclang

var (
	_ directive = directiveSecComponentSignature
	_ directive = directiveSecMarker
	_ directive = directiveSecAction
	_ directive = directiveSecRule
	_ directive = directiveSecResponseBodyAccess
	_ directive = directiveSecRequestBodyLimit
	_ directive = directiveSecRequestBodyAccess
	_ directive = directiveSecRuleEngine
	_ directive = directiveSecWebAppID
	_ directive = directiveSecServerSignature
	_ directive = directiveSecRuleRemoveByTag
	_ directive = directiveSecRuleRemoveByMsg
	_ directive = directiveSecRuleRemoveByID
	_ directive = directiveSecResponseBodyMimeTypesClear
	_ directive = directiveSecResponseBodyMimeType
	_ directive = directiveSecResponseBodyLimitAction
	_ directive = directiveSecResponseBodyLimit
	_ directive = directiveSecRequestBodyLimitAction
	_ directive = directiveSecRequestBodyInMemoryLimit
	_ directive = directiveSecPersistenceEngine
	_ directive = directiveSecRemoteRulesFailAction
	_ directive = directiveSecRemoteRules
	_ directive = directiveSecConnWriteStateLimit
	_ directive = directiveSecSensorID
	_ directive = directiveSecConnReadStateLimit
	_ directive = directiveSecPcreMatchLimitRecursion
	_ directive = directiveSecPcreMatchLimit
	_ directive = directiveSecHTTPBlKey
	_ directive = directiveSecGsbLookupDb
	_ directive = directiveSecHashMethodPm
	_ directive = directiveSecHashMethodRx
	_ directive = directiveSecHashParam
	_ directive = directiveSecHashKey
	_ directive = directiveSecHashEngine
	_ directive = directiveSecDefaultAction
	_ directive = directiveSecConnEngine
	_ directive = directiveSecCollectionTimeout
	_ directive = directiveSecAuditLog
	_ directive = directiveSecAuditLogType
	_ directive = directiveSecAuditLogFormat
	_ directive = directiveSecAuditLogDir
	_ directive = directiveSecAuditLogDirMode
	_ directive = directiveSecAuditLogFileMode
	_ directive = directiveSecAuditLogRelevantStatus
	_ directive = directiveSecAuditLogParts
	_ directive = directiveSecAuditEngine
	_ directive = directiveSecDataDir
	_ directive = directiveSecUploadKeepFiles
	_ directive = directiveSecUploadFileMode
	_ directive = directiveSecUploadFileLimit
	_ directive = directiveSecUploadDir
	_ directive = directiveSecRequestBodyNoFilesLimit
	_ directive = directiveSecDebugLog
	_ directive = directiveSecDebugLogLevel
	_ directive = directiveSecRuleUpdateTargetByID
	_ directive = directiveSecIgnoreRuleCompilationErrors
	_ directive = directiveSecDataset
	_ directive = directiveSecArgumentsLimit
)

var directivesMap = map[string]directive{
	"seccomponentsignature":          directiveSecComponentSignature,
	"secmarker":                      directiveSecMarker,
	"secaction":                      directiveSecAction,
	"secrule":                        directiveSecRule,
	"secresponsebodyaccess":          directiveSecResponseBodyAccess,
	"secrequestbodylimit":            directiveSecRequestBodyLimit,
	"secrequestbodyaccess":           directiveSecRequestBodyAccess,
	"secruleengine":                  directiveSecRuleEngine,
	"secwebappid":                    directiveSecWebAppID,
	"secserversignature":             directiveSecServerSignature,
	"secruleremovebytag":             directiveSecRuleRemoveByTag,
	"secruleremovebymsg":             directiveSecRuleRemoveByMsg,
	"secruleremovebyid":              directiveSecRuleRemoveByID,
	"secresponsebodymimetypesclear":  directiveSecResponseBodyMimeTypesClear,
	"secresponsebodymimetype":        directiveSecResponseBodyMimeType,
	"secresponsebodylimitaction":     directiveSecResponseBodyLimitAction,
	"secresponsebodylimit":           directiveSecResponseBodyLimit,
	"secrequestbodylimitaction":      directiveSecRequestBodyLimitAction,
	"secrequestbodyinmemorylimit":    directiveSecRequestBodyInMemoryLimit,
	"secpersistenceengine":           directiveSecPersistenceEngine,
	"secremoterulesfailaction":       directiveSecRemoteRulesFailAction,
	"secremoterules":                 directiveSecRemoteRules,
	"secconnwritestatelimit":         directiveSecConnWriteStateLimit,
	"secsensorid":                    directiveSecSensorID,
	"secconnreadstatelimit":          directiveSecConnReadStateLimit,
	"secpcrematchlimitrecursion":     directiveSecPcreMatchLimitRecursion,
	"secpcrematchlimit":              directiveSecPcreMatchLimit,
	"sechttpblkey":                   directiveSecHTTPBlKey,
	"secgsblookupdb":                 directiveSecGsbLookupDb,
	"sechashmethodpm":                directiveSecHashMethodPm,
	"sechashmethodrx":                directiveSecHashMethodRx,
	"sechashparam":                   directiveSecHashParam,
	"sechashkey":                     directiveSecHashKey,
	"sechashengine":                  directiveSecHashEngine,
	"secdefaultaction":               directiveSecDefaultAction,
	"secconnengine":                  directiveSecConnEngine,
	"seccollectiontimeout":           directiveSecCollectionTimeout,
	"secauditlog":                    directiveSecAuditLog,
	"secauditlogtype":                directiveSecAuditLogType,
	"secauditlogformat":              directiveSecAuditLogFormat,
	"secauditlogdir":                 directiveSecAuditLogDir,
	"secauditlogdirmode":             directiveSecAuditLogDirMode,
	"secauditlogfilemode":            directiveSecAuditLogFileMode,
	"secauditlogrelevantstatus":      directiveSecAuditLogRelevantStatus,
	"secauditlogparts":               directiveSecAuditLogParts,
	"secauditengine":                 directiveSecAuditEngine,
	"secdatadir":                     directiveSecDataDir,
	"secuploadkeepfiles":             directiveSecUploadKeepFiles,
	"secuploadfilemode":              directiveSecUploadFileMode,
	"secuploadfilelimit":             directiveSecUploadFileLimit,
	"secuploaddir":                   directiveSecUploadDir,
	"secrequestbodynofileslimit":     directiveSecRequestBodyNoFilesLimit,
	"secdebuglog":                    directiveSecDebugLog,
	"secdebugloglevel":               directiveSecDebugLogLevel,
	"secruleupdatetargetbyid":        directiveSecRuleUpdateTargetByID,
	"secignorerulecompilationerrors": directiveSecIgnoreRuleCompilationErrors,
	"secdataset":                     directiveSecDataset,
	"secargumentslimit":              directiveSecArgumentsLimit,

	// Unsupported directives
	"secargumentseparator":     directiveUnsupported,
	"seccookieformat":          directiveUnsupported,
	"secruleupdatetargetbytag": directiveUnsupported,
	"secruleupdatetargetbymsg": directiveUnsupported,
	"secruleupdateactionbyid":  directiveUnsupported,
	"secrulescript":            directiveUnsupported,
	"secruleperftime":          directiveUnsupported,
	"secunicodemap":            directiveUnsupported,
	"sectmpdir":                directiveUnsupported,
}
