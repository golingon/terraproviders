// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataprocjob

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Status struct{}

type HadoopConfig struct {
	// ArchiveUris: list of string, optional
	ArchiveUris terra.ListValue[terra.StringValue] `hcl:"archive_uris,attr"`
	// Args: list of string, optional
	Args terra.ListValue[terra.StringValue] `hcl:"args,attr"`
	// FileUris: list of string, optional
	FileUris terra.ListValue[terra.StringValue] `hcl:"file_uris,attr"`
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// MainClass: string, optional
	MainClass terra.StringValue `hcl:"main_class,attr"`
	// MainJarFileUri: string, optional
	MainJarFileUri terra.StringValue `hcl:"main_jar_file_uri,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// HadoopConfigLoggingConfig: optional
	LoggingConfig *HadoopConfigLoggingConfig `hcl:"logging_config,block"`
}

type HadoopConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type HiveConfig struct {
	// ContinueOnFailure: bool, optional
	ContinueOnFailure terra.BoolValue `hcl:"continue_on_failure,attr"`
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// QueryFileUri: string, optional
	QueryFileUri terra.StringValue `hcl:"query_file_uri,attr"`
	// QueryList: list of string, optional
	QueryList terra.ListValue[terra.StringValue] `hcl:"query_list,attr"`
	// ScriptVariables: map of string, optional
	ScriptVariables terra.MapValue[terra.StringValue] `hcl:"script_variables,attr"`
}

type PigConfig struct {
	// ContinueOnFailure: bool, optional
	ContinueOnFailure terra.BoolValue `hcl:"continue_on_failure,attr"`
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// QueryFileUri: string, optional
	QueryFileUri terra.StringValue `hcl:"query_file_uri,attr"`
	// QueryList: list of string, optional
	QueryList terra.ListValue[terra.StringValue] `hcl:"query_list,attr"`
	// ScriptVariables: map of string, optional
	ScriptVariables terra.MapValue[terra.StringValue] `hcl:"script_variables,attr"`
	// PigConfigLoggingConfig: optional
	LoggingConfig *PigConfigLoggingConfig `hcl:"logging_config,block"`
}

type PigConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type Placement struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
}

type PrestoConfig struct {
	// ClientTags: list of string, optional
	ClientTags terra.ListValue[terra.StringValue] `hcl:"client_tags,attr"`
	// ContinueOnFailure: bool, optional
	ContinueOnFailure terra.BoolValue `hcl:"continue_on_failure,attr"`
	// OutputFormat: string, optional
	OutputFormat terra.StringValue `hcl:"output_format,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// QueryFileUri: string, optional
	QueryFileUri terra.StringValue `hcl:"query_file_uri,attr"`
	// QueryList: list of string, optional
	QueryList terra.ListValue[terra.StringValue] `hcl:"query_list,attr"`
	// PrestoConfigLoggingConfig: optional
	LoggingConfig *PrestoConfigLoggingConfig `hcl:"logging_config,block"`
}

type PrestoConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type PysparkConfig struct {
	// ArchiveUris: list of string, optional
	ArchiveUris terra.ListValue[terra.StringValue] `hcl:"archive_uris,attr"`
	// Args: list of string, optional
	Args terra.ListValue[terra.StringValue] `hcl:"args,attr"`
	// FileUris: list of string, optional
	FileUris terra.ListValue[terra.StringValue] `hcl:"file_uris,attr"`
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// MainPythonFileUri: string, required
	MainPythonFileUri terra.StringValue `hcl:"main_python_file_uri,attr" validate:"required"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// PythonFileUris: list of string, optional
	PythonFileUris terra.ListValue[terra.StringValue] `hcl:"python_file_uris,attr"`
	// PysparkConfigLoggingConfig: optional
	LoggingConfig *PysparkConfigLoggingConfig `hcl:"logging_config,block"`
}

type PysparkConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type Reference struct {
	// JobId: string, optional
	JobId terra.StringValue `hcl:"job_id,attr"`
}

type Scheduling struct {
	// MaxFailuresPerHour: number, required
	MaxFailuresPerHour terra.NumberValue `hcl:"max_failures_per_hour,attr" validate:"required"`
	// MaxFailuresTotal: number, required
	MaxFailuresTotal terra.NumberValue `hcl:"max_failures_total,attr" validate:"required"`
}

type SparkConfig struct {
	// ArchiveUris: list of string, optional
	ArchiveUris terra.ListValue[terra.StringValue] `hcl:"archive_uris,attr"`
	// Args: list of string, optional
	Args terra.ListValue[terra.StringValue] `hcl:"args,attr"`
	// FileUris: list of string, optional
	FileUris terra.ListValue[terra.StringValue] `hcl:"file_uris,attr"`
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// MainClass: string, optional
	MainClass terra.StringValue `hcl:"main_class,attr"`
	// MainJarFileUri: string, optional
	MainJarFileUri terra.StringValue `hcl:"main_jar_file_uri,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// SparkConfigLoggingConfig: optional
	LoggingConfig *SparkConfigLoggingConfig `hcl:"logging_config,block"`
}

type SparkConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type SparksqlConfig struct {
	// JarFileUris: list of string, optional
	JarFileUris terra.ListValue[terra.StringValue] `hcl:"jar_file_uris,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// QueryFileUri: string, optional
	QueryFileUri terra.StringValue `hcl:"query_file_uri,attr"`
	// QueryList: list of string, optional
	QueryList terra.ListValue[terra.StringValue] `hcl:"query_list,attr"`
	// ScriptVariables: map of string, optional
	ScriptVariables terra.MapValue[terra.StringValue] `hcl:"script_variables,attr"`
	// SparksqlConfigLoggingConfig: optional
	LoggingConfig *SparksqlConfigLoggingConfig `hcl:"logging_config,block"`
}

type SparksqlConfigLoggingConfig struct {
	// DriverLogLevels: map of string, required
	DriverLogLevels terra.MapValue[terra.StringValue] `hcl:"driver_log_levels,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type StatusAttributes struct {
	ref terra.Reference
}

func (s StatusAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusAttributes) InternalWithRef(ref terra.Reference) StatusAttributes {
	return StatusAttributes{ref: ref}
}

func (s StatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusAttributes) Details() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("details"))
}

func (s StatusAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("state"))
}

func (s StatusAttributes) StateStartTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("state_start_time"))
}

func (s StatusAttributes) Substate() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("substate"))
}

type HadoopConfigAttributes struct {
	ref terra.Reference
}

func (hc HadoopConfigAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HadoopConfigAttributes) InternalWithRef(ref terra.Reference) HadoopConfigAttributes {
	return HadoopConfigAttributes{ref: ref}
}

func (hc HadoopConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HadoopConfigAttributes) ArchiveUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("archive_uris"))
}

func (hc HadoopConfigAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("args"))
}

func (hc HadoopConfigAttributes) FileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("file_uris"))
}

func (hc HadoopConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("jar_file_uris"))
}

func (hc HadoopConfigAttributes) MainClass() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("main_class"))
}

func (hc HadoopConfigAttributes) MainJarFileUri() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("main_jar_file_uri"))
}

func (hc HadoopConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("properties"))
}

func (hc HadoopConfigAttributes) LoggingConfig() terra.ListValue[HadoopConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[HadoopConfigLoggingConfigAttributes](hc.ref.Append("logging_config"))
}

type HadoopConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc HadoopConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc HadoopConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) HadoopConfigLoggingConfigAttributes {
	return HadoopConfigLoggingConfigAttributes{ref: ref}
}

func (lc HadoopConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc HadoopConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type HiveConfigAttributes struct {
	ref terra.Reference
}

func (hc HiveConfigAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HiveConfigAttributes) InternalWithRef(ref terra.Reference) HiveConfigAttributes {
	return HiveConfigAttributes{ref: ref}
}

func (hc HiveConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HiveConfigAttributes) ContinueOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(hc.ref.Append("continue_on_failure"))
}

func (hc HiveConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("jar_file_uris"))
}

func (hc HiveConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("properties"))
}

func (hc HiveConfigAttributes) QueryFileUri() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("query_file_uri"))
}

func (hc HiveConfigAttributes) QueryList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](hc.ref.Append("query_list"))
}

func (hc HiveConfigAttributes) ScriptVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("script_variables"))
}

type PigConfigAttributes struct {
	ref terra.Reference
}

func (pc PigConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PigConfigAttributes) InternalWithRef(ref terra.Reference) PigConfigAttributes {
	return PigConfigAttributes{ref: ref}
}

func (pc PigConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PigConfigAttributes) ContinueOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("continue_on_failure"))
}

func (pc PigConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("jar_file_uris"))
}

func (pc PigConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("properties"))
}

func (pc PigConfigAttributes) QueryFileUri() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("query_file_uri"))
}

func (pc PigConfigAttributes) QueryList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("query_list"))
}

func (pc PigConfigAttributes) ScriptVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("script_variables"))
}

func (pc PigConfigAttributes) LoggingConfig() terra.ListValue[PigConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[PigConfigLoggingConfigAttributes](pc.ref.Append("logging_config"))
}

type PigConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc PigConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc PigConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) PigConfigLoggingConfigAttributes {
	return PigConfigLoggingConfigAttributes{ref: ref}
}

func (lc PigConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc PigConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type PlacementAttributes struct {
	ref terra.Reference
}

func (p PlacementAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlacementAttributes) InternalWithRef(ref terra.Reference) PlacementAttributes {
	return PlacementAttributes{ref: ref}
}

func (p PlacementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlacementAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("cluster_name"))
}

func (p PlacementAttributes) ClusterUuid() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("cluster_uuid"))
}

type PrestoConfigAttributes struct {
	ref terra.Reference
}

func (pc PrestoConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PrestoConfigAttributes) InternalWithRef(ref terra.Reference) PrestoConfigAttributes {
	return PrestoConfigAttributes{ref: ref}
}

func (pc PrestoConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PrestoConfigAttributes) ClientTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("client_tags"))
}

func (pc PrestoConfigAttributes) ContinueOnFailure() terra.BoolValue {
	return terra.ReferenceAsBool(pc.ref.Append("continue_on_failure"))
}

func (pc PrestoConfigAttributes) OutputFormat() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("output_format"))
}

func (pc PrestoConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("properties"))
}

func (pc PrestoConfigAttributes) QueryFileUri() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("query_file_uri"))
}

func (pc PrestoConfigAttributes) QueryList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("query_list"))
}

func (pc PrestoConfigAttributes) LoggingConfig() terra.ListValue[PrestoConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[PrestoConfigLoggingConfigAttributes](pc.ref.Append("logging_config"))
}

type PrestoConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc PrestoConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc PrestoConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) PrestoConfigLoggingConfigAttributes {
	return PrestoConfigLoggingConfigAttributes{ref: ref}
}

func (lc PrestoConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc PrestoConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type PysparkConfigAttributes struct {
	ref terra.Reference
}

func (pc PysparkConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PysparkConfigAttributes) InternalWithRef(ref terra.Reference) PysparkConfigAttributes {
	return PysparkConfigAttributes{ref: ref}
}

func (pc PysparkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PysparkConfigAttributes) ArchiveUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("archive_uris"))
}

func (pc PysparkConfigAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("args"))
}

func (pc PysparkConfigAttributes) FileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("file_uris"))
}

func (pc PysparkConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("jar_file_uris"))
}

func (pc PysparkConfigAttributes) MainPythonFileUri() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("main_python_file_uri"))
}

func (pc PysparkConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("properties"))
}

func (pc PysparkConfigAttributes) PythonFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("python_file_uris"))
}

func (pc PysparkConfigAttributes) LoggingConfig() terra.ListValue[PysparkConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[PysparkConfigLoggingConfigAttributes](pc.ref.Append("logging_config"))
}

type PysparkConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc PysparkConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc PysparkConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) PysparkConfigLoggingConfigAttributes {
	return PysparkConfigLoggingConfigAttributes{ref: ref}
}

func (lc PysparkConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc PysparkConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type ReferenceAttributes struct {
	ref terra.Reference
}

func (r ReferenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ReferenceAttributes) InternalWithRef(ref terra.Reference) ReferenceAttributes {
	return ReferenceAttributes{ref: ref}
}

func (r ReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ReferenceAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("job_id"))
}

type SchedulingAttributes struct {
	ref terra.Reference
}

func (s SchedulingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SchedulingAttributes) InternalWithRef(ref terra.Reference) SchedulingAttributes {
	return SchedulingAttributes{ref: ref}
}

func (s SchedulingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SchedulingAttributes) MaxFailuresPerHour() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("max_failures_per_hour"))
}

func (s SchedulingAttributes) MaxFailuresTotal() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("max_failures_total"))
}

type SparkConfigAttributes struct {
	ref terra.Reference
}

func (sc SparkConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SparkConfigAttributes) InternalWithRef(ref terra.Reference) SparkConfigAttributes {
	return SparkConfigAttributes{ref: ref}
}

func (sc SparkConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SparkConfigAttributes) ArchiveUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("archive_uris"))
}

func (sc SparkConfigAttributes) Args() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("args"))
}

func (sc SparkConfigAttributes) FileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("file_uris"))
}

func (sc SparkConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("jar_file_uris"))
}

func (sc SparkConfigAttributes) MainClass() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("main_class"))
}

func (sc SparkConfigAttributes) MainJarFileUri() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("main_jar_file_uri"))
}

func (sc SparkConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("properties"))
}

func (sc SparkConfigAttributes) LoggingConfig() terra.ListValue[SparkConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[SparkConfigLoggingConfigAttributes](sc.ref.Append("logging_config"))
}

type SparkConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc SparkConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc SparkConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) SparkConfigLoggingConfigAttributes {
	return SparkConfigLoggingConfigAttributes{ref: ref}
}

func (lc SparkConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc SparkConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type SparksqlConfigAttributes struct {
	ref terra.Reference
}

func (sc SparksqlConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SparksqlConfigAttributes) InternalWithRef(ref terra.Reference) SparksqlConfigAttributes {
	return SparksqlConfigAttributes{ref: ref}
}

func (sc SparksqlConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SparksqlConfigAttributes) JarFileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("jar_file_uris"))
}

func (sc SparksqlConfigAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("properties"))
}

func (sc SparksqlConfigAttributes) QueryFileUri() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("query_file_uri"))
}

func (sc SparksqlConfigAttributes) QueryList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("query_list"))
}

func (sc SparksqlConfigAttributes) ScriptVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("script_variables"))
}

func (sc SparksqlConfigAttributes) LoggingConfig() terra.ListValue[SparksqlConfigLoggingConfigAttributes] {
	return terra.ReferenceAsList[SparksqlConfigLoggingConfigAttributes](sc.ref.Append("logging_config"))
}

type SparksqlConfigLoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc SparksqlConfigLoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc SparksqlConfigLoggingConfigAttributes) InternalWithRef(ref terra.Reference) SparksqlConfigLoggingConfigAttributes {
	return SparksqlConfigLoggingConfigAttributes{ref: ref}
}

func (lc SparksqlConfigLoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc SparksqlConfigLoggingConfigAttributes) DriverLogLevels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("driver_log_levels"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type StatusState struct {
	Details        string `json:"details"`
	State          string `json:"state"`
	StateStartTime string `json:"state_start_time"`
	Substate       string `json:"substate"`
}

type HadoopConfigState struct {
	ArchiveUris    []string                         `json:"archive_uris"`
	Args           []string                         `json:"args"`
	FileUris       []string                         `json:"file_uris"`
	JarFileUris    []string                         `json:"jar_file_uris"`
	MainClass      string                           `json:"main_class"`
	MainJarFileUri string                           `json:"main_jar_file_uri"`
	Properties     map[string]string                `json:"properties"`
	LoggingConfig  []HadoopConfigLoggingConfigState `json:"logging_config"`
}

type HadoopConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type HiveConfigState struct {
	ContinueOnFailure bool              `json:"continue_on_failure"`
	JarFileUris       []string          `json:"jar_file_uris"`
	Properties        map[string]string `json:"properties"`
	QueryFileUri      string            `json:"query_file_uri"`
	QueryList         []string          `json:"query_list"`
	ScriptVariables   map[string]string `json:"script_variables"`
}

type PigConfigState struct {
	ContinueOnFailure bool                          `json:"continue_on_failure"`
	JarFileUris       []string                      `json:"jar_file_uris"`
	Properties        map[string]string             `json:"properties"`
	QueryFileUri      string                        `json:"query_file_uri"`
	QueryList         []string                      `json:"query_list"`
	ScriptVariables   map[string]string             `json:"script_variables"`
	LoggingConfig     []PigConfigLoggingConfigState `json:"logging_config"`
}

type PigConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type PlacementState struct {
	ClusterName string `json:"cluster_name"`
	ClusterUuid string `json:"cluster_uuid"`
}

type PrestoConfigState struct {
	ClientTags        []string                         `json:"client_tags"`
	ContinueOnFailure bool                             `json:"continue_on_failure"`
	OutputFormat      string                           `json:"output_format"`
	Properties        map[string]string                `json:"properties"`
	QueryFileUri      string                           `json:"query_file_uri"`
	QueryList         []string                         `json:"query_list"`
	LoggingConfig     []PrestoConfigLoggingConfigState `json:"logging_config"`
}

type PrestoConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type PysparkConfigState struct {
	ArchiveUris       []string                          `json:"archive_uris"`
	Args              []string                          `json:"args"`
	FileUris          []string                          `json:"file_uris"`
	JarFileUris       []string                          `json:"jar_file_uris"`
	MainPythonFileUri string                            `json:"main_python_file_uri"`
	Properties        map[string]string                 `json:"properties"`
	PythonFileUris    []string                          `json:"python_file_uris"`
	LoggingConfig     []PysparkConfigLoggingConfigState `json:"logging_config"`
}

type PysparkConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type ReferenceState struct {
	JobId string `json:"job_id"`
}

type SchedulingState struct {
	MaxFailuresPerHour float64 `json:"max_failures_per_hour"`
	MaxFailuresTotal   float64 `json:"max_failures_total"`
}

type SparkConfigState struct {
	ArchiveUris    []string                        `json:"archive_uris"`
	Args           []string                        `json:"args"`
	FileUris       []string                        `json:"file_uris"`
	JarFileUris    []string                        `json:"jar_file_uris"`
	MainClass      string                          `json:"main_class"`
	MainJarFileUri string                          `json:"main_jar_file_uri"`
	Properties     map[string]string               `json:"properties"`
	LoggingConfig  []SparkConfigLoggingConfigState `json:"logging_config"`
}

type SparkConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type SparksqlConfigState struct {
	JarFileUris     []string                           `json:"jar_file_uris"`
	Properties      map[string]string                  `json:"properties"`
	QueryFileUri    string                             `json:"query_file_uri"`
	QueryList       []string                           `json:"query_list"`
	ScriptVariables map[string]string                  `json:"script_variables"`
	LoggingConfig   []SparksqlConfigLoggingConfigState `json:"logging_config"`
}

type SparksqlConfigLoggingConfigState struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
