// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamonitordatacollectionrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DataFlow struct{}

type DataSources struct {
	// DataImport: min=0
	DataImport []DataImport `hcl:"data_import,block" validate:"min=0"`
	// Extension: min=0
	Extension []Extension `hcl:"extension,block" validate:"min=0"`
	// IisLog: min=0
	IisLog []IisLog `hcl:"iis_log,block" validate:"min=0"`
	// LogFile: min=0
	LogFile []LogFile `hcl:"log_file,block" validate:"min=0"`
	// PerformanceCounter: min=0
	PerformanceCounter []PerformanceCounter `hcl:"performance_counter,block" validate:"min=0"`
	// PlatformTelemetry: min=0
	PlatformTelemetry []PlatformTelemetry `hcl:"platform_telemetry,block" validate:"min=0"`
	// PrometheusForwarder: min=0
	PrometheusForwarder []PrometheusForwarder `hcl:"prometheus_forwarder,block" validate:"min=0"`
	// Syslog: min=0
	Syslog []Syslog `hcl:"syslog,block" validate:"min=0"`
	// WindowsEventLog: min=0
	WindowsEventLog []WindowsEventLog `hcl:"windows_event_log,block" validate:"min=0"`
	// WindowsFirewallLog: min=0
	WindowsFirewallLog []WindowsFirewallLog `hcl:"windows_firewall_log,block" validate:"min=0"`
}

type DataImport struct {
	// EventHubDataSource: min=0
	EventHubDataSource []EventHubDataSource `hcl:"event_hub_data_source,block" validate:"min=0"`
}

type EventHubDataSource struct{}

type Extension struct{}

type IisLog struct{}

type LogFile struct {
	// Settings: min=0
	Settings []Settings `hcl:"settings,block" validate:"min=0"`
}

type Settings struct {
	// Text: min=0
	Text []Text `hcl:"text,block" validate:"min=0"`
}

type Text struct{}

type PerformanceCounter struct{}

type PlatformTelemetry struct{}

type PrometheusForwarder struct {
	// LabelIncludeFilter: min=0
	LabelIncludeFilter []LabelIncludeFilter `hcl:"label_include_filter,block" validate:"min=0"`
}

type LabelIncludeFilter struct{}

type Syslog struct{}

type WindowsEventLog struct{}

type WindowsFirewallLog struct{}

type Destinations struct {
	// AzureMonitorMetrics: min=0
	AzureMonitorMetrics []AzureMonitorMetrics `hcl:"azure_monitor_metrics,block" validate:"min=0"`
	// EventHub: min=0
	EventHub []EventHub `hcl:"event_hub,block" validate:"min=0"`
	// EventHubDirect: min=0
	EventHubDirect []EventHubDirect `hcl:"event_hub_direct,block" validate:"min=0"`
	// LogAnalytics: min=0
	LogAnalytics []LogAnalytics `hcl:"log_analytics,block" validate:"min=0"`
	// MonitorAccount: min=0
	MonitorAccount []MonitorAccount `hcl:"monitor_account,block" validate:"min=0"`
	// StorageBlob: min=0
	StorageBlob []StorageBlob `hcl:"storage_blob,block" validate:"min=0"`
	// StorageBlobDirect: min=0
	StorageBlobDirect []StorageBlobDirect `hcl:"storage_blob_direct,block" validate:"min=0"`
	// StorageTableDirect: min=0
	StorageTableDirect []StorageTableDirect `hcl:"storage_table_direct,block" validate:"min=0"`
}

type AzureMonitorMetrics struct{}

type EventHub struct{}

type EventHubDirect struct{}

type LogAnalytics struct{}

type MonitorAccount struct{}

type StorageBlob struct{}

type StorageBlobDirect struct{}

type StorageTableDirect struct{}

type Identity struct{}

type StreamDeclaration struct {
	// Column: min=0
	Column []Column `hcl:"column,block" validate:"min=0"`
}

type Column struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataFlowAttributes struct {
	ref terra.Reference
}

func (df DataFlowAttributes) InternalRef() (terra.Reference, error) {
	return df.ref, nil
}

func (df DataFlowAttributes) InternalWithRef(ref terra.Reference) DataFlowAttributes {
	return DataFlowAttributes{ref: ref}
}

func (df DataFlowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return df.ref.InternalTokens()
}

func (df DataFlowAttributes) BuiltInTransform() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("built_in_transform"))
}

func (df DataFlowAttributes) Destinations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](df.ref.Append("destinations"))
}

func (df DataFlowAttributes) OutputStream() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("output_stream"))
}

func (df DataFlowAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](df.ref.Append("streams"))
}

func (df DataFlowAttributes) TransformKql() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("transform_kql"))
}

type DataSourcesAttributes struct {
	ref terra.Reference
}

func (ds DataSourcesAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds DataSourcesAttributes) InternalWithRef(ref terra.Reference) DataSourcesAttributes {
	return DataSourcesAttributes{ref: ref}
}

func (ds DataSourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds DataSourcesAttributes) DataImport() terra.ListValue[DataImportAttributes] {
	return terra.ReferenceAsList[DataImportAttributes](ds.ref.Append("data_import"))
}

func (ds DataSourcesAttributes) Extension() terra.ListValue[ExtensionAttributes] {
	return terra.ReferenceAsList[ExtensionAttributes](ds.ref.Append("extension"))
}

func (ds DataSourcesAttributes) IisLog() terra.ListValue[IisLogAttributes] {
	return terra.ReferenceAsList[IisLogAttributes](ds.ref.Append("iis_log"))
}

func (ds DataSourcesAttributes) LogFile() terra.ListValue[LogFileAttributes] {
	return terra.ReferenceAsList[LogFileAttributes](ds.ref.Append("log_file"))
}

func (ds DataSourcesAttributes) PerformanceCounter() terra.ListValue[PerformanceCounterAttributes] {
	return terra.ReferenceAsList[PerformanceCounterAttributes](ds.ref.Append("performance_counter"))
}

func (ds DataSourcesAttributes) PlatformTelemetry() terra.ListValue[PlatformTelemetryAttributes] {
	return terra.ReferenceAsList[PlatformTelemetryAttributes](ds.ref.Append("platform_telemetry"))
}

func (ds DataSourcesAttributes) PrometheusForwarder() terra.ListValue[PrometheusForwarderAttributes] {
	return terra.ReferenceAsList[PrometheusForwarderAttributes](ds.ref.Append("prometheus_forwarder"))
}

func (ds DataSourcesAttributes) Syslog() terra.ListValue[SyslogAttributes] {
	return terra.ReferenceAsList[SyslogAttributes](ds.ref.Append("syslog"))
}

func (ds DataSourcesAttributes) WindowsEventLog() terra.ListValue[WindowsEventLogAttributes] {
	return terra.ReferenceAsList[WindowsEventLogAttributes](ds.ref.Append("windows_event_log"))
}

func (ds DataSourcesAttributes) WindowsFirewallLog() terra.ListValue[WindowsFirewallLogAttributes] {
	return terra.ReferenceAsList[WindowsFirewallLogAttributes](ds.ref.Append("windows_firewall_log"))
}

type DataImportAttributes struct {
	ref terra.Reference
}

func (di DataImportAttributes) InternalRef() (terra.Reference, error) {
	return di.ref, nil
}

func (di DataImportAttributes) InternalWithRef(ref terra.Reference) DataImportAttributes {
	return DataImportAttributes{ref: ref}
}

func (di DataImportAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return di.ref.InternalTokens()
}

func (di DataImportAttributes) EventHubDataSource() terra.ListValue[EventHubDataSourceAttributes] {
	return terra.ReferenceAsList[EventHubDataSourceAttributes](di.ref.Append("event_hub_data_source"))
}

type EventHubDataSourceAttributes struct {
	ref terra.Reference
}

func (ehds EventHubDataSourceAttributes) InternalRef() (terra.Reference, error) {
	return ehds.ref, nil
}

func (ehds EventHubDataSourceAttributes) InternalWithRef(ref terra.Reference) EventHubDataSourceAttributes {
	return EventHubDataSourceAttributes{ref: ref}
}

func (ehds EventHubDataSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ehds.ref.InternalTokens()
}

func (ehds EventHubDataSourceAttributes) ConsumerGroup() terra.StringValue {
	return terra.ReferenceAsString(ehds.ref.Append("consumer_group"))
}

func (ehds EventHubDataSourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ehds.ref.Append("name"))
}

func (ehds EventHubDataSourceAttributes) Stream() terra.StringValue {
	return terra.ReferenceAsString(ehds.ref.Append("stream"))
}

type ExtensionAttributes struct {
	ref terra.Reference
}

func (e ExtensionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExtensionAttributes) InternalWithRef(ref terra.Reference) ExtensionAttributes {
	return ExtensionAttributes{ref: ref}
}

func (e ExtensionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExtensionAttributes) ExtensionJson() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("extension_json"))
}

func (e ExtensionAttributes) ExtensionName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("extension_name"))
}

func (e ExtensionAttributes) InputDataSources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("input_data_sources"))
}

func (e ExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

func (e ExtensionAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("streams"))
}

type IisLogAttributes struct {
	ref terra.Reference
}

func (il IisLogAttributes) InternalRef() (terra.Reference, error) {
	return il.ref, nil
}

func (il IisLogAttributes) InternalWithRef(ref terra.Reference) IisLogAttributes {
	return IisLogAttributes{ref: ref}
}

func (il IisLogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return il.ref.InternalTokens()
}

func (il IisLogAttributes) LogDirectories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](il.ref.Append("log_directories"))
}

func (il IisLogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(il.ref.Append("name"))
}

func (il IisLogAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](il.ref.Append("streams"))
}

type LogFileAttributes struct {
	ref terra.Reference
}

func (lf LogFileAttributes) InternalRef() (terra.Reference, error) {
	return lf.ref, nil
}

func (lf LogFileAttributes) InternalWithRef(ref terra.Reference) LogFileAttributes {
	return LogFileAttributes{ref: ref}
}

func (lf LogFileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lf.ref.InternalTokens()
}

func (lf LogFileAttributes) FilePatterns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lf.ref.Append("file_patterns"))
}

func (lf LogFileAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("format"))
}

func (lf LogFileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lf.ref.Append("name"))
}

func (lf LogFileAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lf.ref.Append("streams"))
}

func (lf LogFileAttributes) Settings() terra.ListValue[SettingsAttributes] {
	return terra.ReferenceAsList[SettingsAttributes](lf.ref.Append("settings"))
}

type SettingsAttributes struct {
	ref terra.Reference
}

func (s SettingsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingsAttributes) InternalWithRef(ref terra.Reference) SettingsAttributes {
	return SettingsAttributes{ref: ref}
}

func (s SettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingsAttributes) Text() terra.ListValue[TextAttributes] {
	return terra.ReferenceAsList[TextAttributes](s.ref.Append("text"))
}

type TextAttributes struct {
	ref terra.Reference
}

func (t TextAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TextAttributes) InternalWithRef(ref terra.Reference) TextAttributes {
	return TextAttributes{ref: ref}
}

func (t TextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TextAttributes) RecordStartTimestampFormat() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("record_start_timestamp_format"))
}

type PerformanceCounterAttributes struct {
	ref terra.Reference
}

func (pc PerformanceCounterAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PerformanceCounterAttributes) InternalWithRef(ref terra.Reference) PerformanceCounterAttributes {
	return PerformanceCounterAttributes{ref: ref}
}

func (pc PerformanceCounterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PerformanceCounterAttributes) CounterSpecifiers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("counter_specifiers"))
}

func (pc PerformanceCounterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("name"))
}

func (pc PerformanceCounterAttributes) SamplingFrequencyInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("sampling_frequency_in_seconds"))
}

func (pc PerformanceCounterAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pc.ref.Append("streams"))
}

type PlatformTelemetryAttributes struct {
	ref terra.Reference
}

func (pt PlatformTelemetryAttributes) InternalRef() (terra.Reference, error) {
	return pt.ref, nil
}

func (pt PlatformTelemetryAttributes) InternalWithRef(ref terra.Reference) PlatformTelemetryAttributes {
	return PlatformTelemetryAttributes{ref: ref}
}

func (pt PlatformTelemetryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pt.ref.InternalTokens()
}

func (pt PlatformTelemetryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("name"))
}

func (pt PlatformTelemetryAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pt.ref.Append("streams"))
}

type PrometheusForwarderAttributes struct {
	ref terra.Reference
}

func (pf PrometheusForwarderAttributes) InternalRef() (terra.Reference, error) {
	return pf.ref, nil
}

func (pf PrometheusForwarderAttributes) InternalWithRef(ref terra.Reference) PrometheusForwarderAttributes {
	return PrometheusForwarderAttributes{ref: ref}
}

func (pf PrometheusForwarderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pf.ref.InternalTokens()
}

func (pf PrometheusForwarderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("name"))
}

func (pf PrometheusForwarderAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pf.ref.Append("streams"))
}

func (pf PrometheusForwarderAttributes) LabelIncludeFilter() terra.ListValue[LabelIncludeFilterAttributes] {
	return terra.ReferenceAsList[LabelIncludeFilterAttributes](pf.ref.Append("label_include_filter"))
}

type LabelIncludeFilterAttributes struct {
	ref terra.Reference
}

func (lif LabelIncludeFilterAttributes) InternalRef() (terra.Reference, error) {
	return lif.ref, nil
}

func (lif LabelIncludeFilterAttributes) InternalWithRef(ref terra.Reference) LabelIncludeFilterAttributes {
	return LabelIncludeFilterAttributes{ref: ref}
}

func (lif LabelIncludeFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lif.ref.InternalTokens()
}

func (lif LabelIncludeFilterAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(lif.ref.Append("label"))
}

func (lif LabelIncludeFilterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(lif.ref.Append("value"))
}

type SyslogAttributes struct {
	ref terra.Reference
}

func (s SyslogAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SyslogAttributes) InternalWithRef(ref terra.Reference) SyslogAttributes {
	return SyslogAttributes{ref: ref}
}

func (s SyslogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SyslogAttributes) FacilityNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("facility_names"))
}

func (s SyslogAttributes) LogLevels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("log_levels"))
}

func (s SyslogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SyslogAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("streams"))
}

type WindowsEventLogAttributes struct {
	ref terra.Reference
}

func (wel WindowsEventLogAttributes) InternalRef() (terra.Reference, error) {
	return wel.ref, nil
}

func (wel WindowsEventLogAttributes) InternalWithRef(ref terra.Reference) WindowsEventLogAttributes {
	return WindowsEventLogAttributes{ref: ref}
}

func (wel WindowsEventLogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wel.ref.InternalTokens()
}

func (wel WindowsEventLogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wel.ref.Append("name"))
}

func (wel WindowsEventLogAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wel.ref.Append("streams"))
}

func (wel WindowsEventLogAttributes) XPathQueries() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wel.ref.Append("x_path_queries"))
}

type WindowsFirewallLogAttributes struct {
	ref terra.Reference
}

func (wfl WindowsFirewallLogAttributes) InternalRef() (terra.Reference, error) {
	return wfl.ref, nil
}

func (wfl WindowsFirewallLogAttributes) InternalWithRef(ref terra.Reference) WindowsFirewallLogAttributes {
	return WindowsFirewallLogAttributes{ref: ref}
}

func (wfl WindowsFirewallLogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wfl.ref.InternalTokens()
}

func (wfl WindowsFirewallLogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wfl.ref.Append("name"))
}

func (wfl WindowsFirewallLogAttributes) Streams() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wfl.ref.Append("streams"))
}

type DestinationsAttributes struct {
	ref terra.Reference
}

func (d DestinationsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationsAttributes) InternalWithRef(ref terra.Reference) DestinationsAttributes {
	return DestinationsAttributes{ref: ref}
}

func (d DestinationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationsAttributes) AzureMonitorMetrics() terra.ListValue[AzureMonitorMetricsAttributes] {
	return terra.ReferenceAsList[AzureMonitorMetricsAttributes](d.ref.Append("azure_monitor_metrics"))
}

func (d DestinationsAttributes) EventHub() terra.ListValue[EventHubAttributes] {
	return terra.ReferenceAsList[EventHubAttributes](d.ref.Append("event_hub"))
}

func (d DestinationsAttributes) EventHubDirect() terra.ListValue[EventHubDirectAttributes] {
	return terra.ReferenceAsList[EventHubDirectAttributes](d.ref.Append("event_hub_direct"))
}

func (d DestinationsAttributes) LogAnalytics() terra.ListValue[LogAnalyticsAttributes] {
	return terra.ReferenceAsList[LogAnalyticsAttributes](d.ref.Append("log_analytics"))
}

func (d DestinationsAttributes) MonitorAccount() terra.ListValue[MonitorAccountAttributes] {
	return terra.ReferenceAsList[MonitorAccountAttributes](d.ref.Append("monitor_account"))
}

func (d DestinationsAttributes) StorageBlob() terra.ListValue[StorageBlobAttributes] {
	return terra.ReferenceAsList[StorageBlobAttributes](d.ref.Append("storage_blob"))
}

func (d DestinationsAttributes) StorageBlobDirect() terra.ListValue[StorageBlobDirectAttributes] {
	return terra.ReferenceAsList[StorageBlobDirectAttributes](d.ref.Append("storage_blob_direct"))
}

func (d DestinationsAttributes) StorageTableDirect() terra.ListValue[StorageTableDirectAttributes] {
	return terra.ReferenceAsList[StorageTableDirectAttributes](d.ref.Append("storage_table_direct"))
}

type AzureMonitorMetricsAttributes struct {
	ref terra.Reference
}

func (amm AzureMonitorMetricsAttributes) InternalRef() (terra.Reference, error) {
	return amm.ref, nil
}

func (amm AzureMonitorMetricsAttributes) InternalWithRef(ref terra.Reference) AzureMonitorMetricsAttributes {
	return AzureMonitorMetricsAttributes{ref: ref}
}

func (amm AzureMonitorMetricsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amm.ref.InternalTokens()
}

func (amm AzureMonitorMetricsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amm.ref.Append("name"))
}

type EventHubAttributes struct {
	ref terra.Reference
}

func (eh EventHubAttributes) InternalRef() (terra.Reference, error) {
	return eh.ref, nil
}

func (eh EventHubAttributes) InternalWithRef(ref terra.Reference) EventHubAttributes {
	return EventHubAttributes{ref: ref}
}

func (eh EventHubAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eh.ref.InternalTokens()
}

func (eh EventHubAttributes) EventHubId() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("event_hub_id"))
}

func (eh EventHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("name"))
}

type EventHubDirectAttributes struct {
	ref terra.Reference
}

func (ehd EventHubDirectAttributes) InternalRef() (terra.Reference, error) {
	return ehd.ref, nil
}

func (ehd EventHubDirectAttributes) InternalWithRef(ref terra.Reference) EventHubDirectAttributes {
	return EventHubDirectAttributes{ref: ref}
}

func (ehd EventHubDirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ehd.ref.InternalTokens()
}

func (ehd EventHubDirectAttributes) EventHubId() terra.StringValue {
	return terra.ReferenceAsString(ehd.ref.Append("event_hub_id"))
}

func (ehd EventHubDirectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ehd.ref.Append("name"))
}

type LogAnalyticsAttributes struct {
	ref terra.Reference
}

func (la LogAnalyticsAttributes) InternalRef() (terra.Reference, error) {
	return la.ref, nil
}

func (la LogAnalyticsAttributes) InternalWithRef(ref terra.Reference) LogAnalyticsAttributes {
	return LogAnalyticsAttributes{ref: ref}
}

func (la LogAnalyticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return la.ref.InternalTokens()
}

func (la LogAnalyticsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("name"))
}

func (la LogAnalyticsAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("workspace_resource_id"))
}

type MonitorAccountAttributes struct {
	ref terra.Reference
}

func (ma MonitorAccountAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MonitorAccountAttributes) InternalWithRef(ref terra.Reference) MonitorAccountAttributes {
	return MonitorAccountAttributes{ref: ref}
}

func (ma MonitorAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MonitorAccountAttributes) MonitorAccountId() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("monitor_account_id"))
}

func (ma MonitorAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("name"))
}

type StorageBlobAttributes struct {
	ref terra.Reference
}

func (sb StorageBlobAttributes) InternalRef() (terra.Reference, error) {
	return sb.ref, nil
}

func (sb StorageBlobAttributes) InternalWithRef(ref terra.Reference) StorageBlobAttributes {
	return StorageBlobAttributes{ref: ref}
}

func (sb StorageBlobAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sb.ref.InternalTokens()
}

func (sb StorageBlobAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("container_name"))
}

func (sb StorageBlobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
}

func (sb StorageBlobAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("storage_account_id"))
}

type StorageBlobDirectAttributes struct {
	ref terra.Reference
}

func (sbd StorageBlobDirectAttributes) InternalRef() (terra.Reference, error) {
	return sbd.ref, nil
}

func (sbd StorageBlobDirectAttributes) InternalWithRef(ref terra.Reference) StorageBlobDirectAttributes {
	return StorageBlobDirectAttributes{ref: ref}
}

func (sbd StorageBlobDirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sbd.ref.InternalTokens()
}

func (sbd StorageBlobDirectAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(sbd.ref.Append("container_name"))
}

func (sbd StorageBlobDirectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbd.ref.Append("name"))
}

func (sbd StorageBlobDirectAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sbd.ref.Append("storage_account_id"))
}

type StorageTableDirectAttributes struct {
	ref terra.Reference
}

func (std StorageTableDirectAttributes) InternalRef() (terra.Reference, error) {
	return std.ref, nil
}

func (std StorageTableDirectAttributes) InternalWithRef(ref terra.Reference) StorageTableDirectAttributes {
	return StorageTableDirectAttributes{ref: ref}
}

func (std StorageTableDirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return std.ref.InternalTokens()
}

func (std StorageTableDirectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("name"))
}

func (std StorageTableDirectAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("storage_account_id"))
}

func (std StorageTableDirectAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("table_name"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type StreamDeclarationAttributes struct {
	ref terra.Reference
}

func (sd StreamDeclarationAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd StreamDeclarationAttributes) InternalWithRef(ref terra.Reference) StreamDeclarationAttributes {
	return StreamDeclarationAttributes{ref: ref}
}

func (sd StreamDeclarationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd StreamDeclarationAttributes) StreamName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("stream_name"))
}

func (sd StreamDeclarationAttributes) Column() terra.ListValue[ColumnAttributes] {
	return terra.ReferenceAsList[ColumnAttributes](sd.ref.Append("column"))
}

type ColumnAttributes struct {
	ref terra.Reference
}

func (c ColumnAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ColumnAttributes) InternalWithRef(ref terra.Reference) ColumnAttributes {
	return ColumnAttributes{ref: ref}
}

func (c ColumnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ColumnAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ColumnAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataFlowState struct {
	BuiltInTransform string   `json:"built_in_transform"`
	Destinations     []string `json:"destinations"`
	OutputStream     string   `json:"output_stream"`
	Streams          []string `json:"streams"`
	TransformKql     string   `json:"transform_kql"`
}

type DataSourcesState struct {
	DataImport          []DataImportState          `json:"data_import"`
	Extension           []ExtensionState           `json:"extension"`
	IisLog              []IisLogState              `json:"iis_log"`
	LogFile             []LogFileState             `json:"log_file"`
	PerformanceCounter  []PerformanceCounterState  `json:"performance_counter"`
	PlatformTelemetry   []PlatformTelemetryState   `json:"platform_telemetry"`
	PrometheusForwarder []PrometheusForwarderState `json:"prometheus_forwarder"`
	Syslog              []SyslogState              `json:"syslog"`
	WindowsEventLog     []WindowsEventLogState     `json:"windows_event_log"`
	WindowsFirewallLog  []WindowsFirewallLogState  `json:"windows_firewall_log"`
}

type DataImportState struct {
	EventHubDataSource []EventHubDataSourceState `json:"event_hub_data_source"`
}

type EventHubDataSourceState struct {
	ConsumerGroup string `json:"consumer_group"`
	Name          string `json:"name"`
	Stream        string `json:"stream"`
}

type ExtensionState struct {
	ExtensionJson    string   `json:"extension_json"`
	ExtensionName    string   `json:"extension_name"`
	InputDataSources []string `json:"input_data_sources"`
	Name             string   `json:"name"`
	Streams          []string `json:"streams"`
}

type IisLogState struct {
	LogDirectories []string `json:"log_directories"`
	Name           string   `json:"name"`
	Streams        []string `json:"streams"`
}

type LogFileState struct {
	FilePatterns []string        `json:"file_patterns"`
	Format       string          `json:"format"`
	Name         string          `json:"name"`
	Streams      []string        `json:"streams"`
	Settings     []SettingsState `json:"settings"`
}

type SettingsState struct {
	Text []TextState `json:"text"`
}

type TextState struct {
	RecordStartTimestampFormat string `json:"record_start_timestamp_format"`
}

type PerformanceCounterState struct {
	CounterSpecifiers          []string `json:"counter_specifiers"`
	Name                       string   `json:"name"`
	SamplingFrequencyInSeconds float64  `json:"sampling_frequency_in_seconds"`
	Streams                    []string `json:"streams"`
}

type PlatformTelemetryState struct {
	Name    string   `json:"name"`
	Streams []string `json:"streams"`
}

type PrometheusForwarderState struct {
	Name               string                    `json:"name"`
	Streams            []string                  `json:"streams"`
	LabelIncludeFilter []LabelIncludeFilterState `json:"label_include_filter"`
}

type LabelIncludeFilterState struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SyslogState struct {
	FacilityNames []string `json:"facility_names"`
	LogLevels     []string `json:"log_levels"`
	Name          string   `json:"name"`
	Streams       []string `json:"streams"`
}

type WindowsEventLogState struct {
	Name         string   `json:"name"`
	Streams      []string `json:"streams"`
	XPathQueries []string `json:"x_path_queries"`
}

type WindowsFirewallLogState struct {
	Name    string   `json:"name"`
	Streams []string `json:"streams"`
}

type DestinationsState struct {
	AzureMonitorMetrics []AzureMonitorMetricsState `json:"azure_monitor_metrics"`
	EventHub            []EventHubState            `json:"event_hub"`
	EventHubDirect      []EventHubDirectState      `json:"event_hub_direct"`
	LogAnalytics        []LogAnalyticsState        `json:"log_analytics"`
	MonitorAccount      []MonitorAccountState      `json:"monitor_account"`
	StorageBlob         []StorageBlobState         `json:"storage_blob"`
	StorageBlobDirect   []StorageBlobDirectState   `json:"storage_blob_direct"`
	StorageTableDirect  []StorageTableDirectState  `json:"storage_table_direct"`
}

type AzureMonitorMetricsState struct {
	Name string `json:"name"`
}

type EventHubState struct {
	EventHubId string `json:"event_hub_id"`
	Name       string `json:"name"`
}

type EventHubDirectState struct {
	EventHubId string `json:"event_hub_id"`
	Name       string `json:"name"`
}

type LogAnalyticsState struct {
	Name                string `json:"name"`
	WorkspaceResourceId string `json:"workspace_resource_id"`
}

type MonitorAccountState struct {
	MonitorAccountId string `json:"monitor_account_id"`
	Name             string `json:"name"`
}

type StorageBlobState struct {
	ContainerName    string `json:"container_name"`
	Name             string `json:"name"`
	StorageAccountId string `json:"storage_account_id"`
}

type StorageBlobDirectState struct {
	ContainerName    string `json:"container_name"`
	Name             string `json:"name"`
	StorageAccountId string `json:"storage_account_id"`
}

type StorageTableDirectState struct {
	Name             string `json:"name"`
	StorageAccountId string `json:"storage_account_id"`
	TableName        string `json:"table_name"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type StreamDeclarationState struct {
	StreamName string        `json:"stream_name"`
	Column     []ColumnState `json:"column"`
}

type ColumnState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
