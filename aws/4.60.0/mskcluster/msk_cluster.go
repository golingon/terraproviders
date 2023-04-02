// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mskcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BrokerNodeGroupInfo struct {
	// AzDistribution: string, optional
	AzDistribution terra.StringValue `hcl:"az_distribution,attr"`
	// ClientSubnets: set of string, required
	ClientSubnets terra.SetValue[terra.StringValue] `hcl:"client_subnets,attr" validate:"required"`
	// EbsVolumeSize: number, optional
	EbsVolumeSize terra.NumberValue `hcl:"ebs_volume_size,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// SecurityGroups: set of string, required
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr" validate:"required"`
	// ConnectivityInfo: optional
	ConnectivityInfo *ConnectivityInfo `hcl:"connectivity_info,block"`
	// StorageInfo: optional
	StorageInfo *StorageInfo `hcl:"storage_info,block"`
}

type ConnectivityInfo struct {
	// PublicAccess: optional
	PublicAccess *PublicAccess `hcl:"public_access,block"`
}

type PublicAccess struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type StorageInfo struct {
	// EbsStorageInfo: optional
	EbsStorageInfo *EbsStorageInfo `hcl:"ebs_storage_info,block"`
}

type EbsStorageInfo struct {
	// VolumeSize: number, optional
	VolumeSize terra.NumberValue `hcl:"volume_size,attr"`
	// ProvisionedThroughput: optional
	ProvisionedThroughput *ProvisionedThroughput `hcl:"provisioned_throughput,block"`
}

type ProvisionedThroughput struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// VolumeThroughput: number, optional
	VolumeThroughput terra.NumberValue `hcl:"volume_throughput,attr"`
}

type ClientAuthentication struct {
	// Unauthenticated: bool, optional
	Unauthenticated terra.BoolValue `hcl:"unauthenticated,attr"`
	// Sasl: optional
	Sasl *Sasl `hcl:"sasl,block"`
	// Tls: optional
	Tls *Tls `hcl:"tls,block"`
}

type Sasl struct {
	// Iam: bool, optional
	Iam terra.BoolValue `hcl:"iam,attr"`
	// Scram: bool, optional
	Scram terra.BoolValue `hcl:"scram,attr"`
}

type Tls struct {
	// CertificateAuthorityArns: set of string, optional
	CertificateAuthorityArns terra.SetValue[terra.StringValue] `hcl:"certificate_authority_arns,attr"`
}

type ConfigurationInfo struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Revision: number, required
	Revision terra.NumberValue `hcl:"revision,attr" validate:"required"`
}

type EncryptionInfo struct {
	// EncryptionAtRestKmsKeyArn: string, optional
	EncryptionAtRestKmsKeyArn terra.StringValue `hcl:"encryption_at_rest_kms_key_arn,attr"`
	// EncryptionInTransit: optional
	EncryptionInTransit *EncryptionInTransit `hcl:"encryption_in_transit,block"`
}

type EncryptionInTransit struct {
	// ClientBroker: string, optional
	ClientBroker terra.StringValue `hcl:"client_broker,attr"`
	// InCluster: bool, optional
	InCluster terra.BoolValue `hcl:"in_cluster,attr"`
}

type LoggingInfo struct {
	// BrokerLogs: required
	BrokerLogs *BrokerLogs `hcl:"broker_logs,block" validate:"required"`
}

type BrokerLogs struct {
	// CloudwatchLogs: optional
	CloudwatchLogs *CloudwatchLogs `hcl:"cloudwatch_logs,block"`
	// Firehose: optional
	Firehose *Firehose `hcl:"firehose,block"`
	// S3: optional
	S3 *S3 `hcl:"s3,block"`
}

type CloudwatchLogs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// LogGroup: string, optional
	LogGroup terra.StringValue `hcl:"log_group,attr"`
}

type Firehose struct {
	// DeliveryStream: string, optional
	DeliveryStream terra.StringValue `hcl:"delivery_stream,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type S3 struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}

type OpenMonitoring struct {
	// Prometheus: required
	Prometheus *Prometheus `hcl:"prometheus,block" validate:"required"`
}

type Prometheus struct {
	// JmxExporter: optional
	JmxExporter *JmxExporter `hcl:"jmx_exporter,block"`
	// NodeExporter: optional
	NodeExporter *NodeExporter `hcl:"node_exporter,block"`
}

type JmxExporter struct {
	// EnabledInBroker: bool, required
	EnabledInBroker terra.BoolValue `hcl:"enabled_in_broker,attr" validate:"required"`
}

type NodeExporter struct {
	// EnabledInBroker: bool, required
	EnabledInBroker terra.BoolValue `hcl:"enabled_in_broker,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type BrokerNodeGroupInfoAttributes struct {
	ref terra.Reference
}

func (bngi BrokerNodeGroupInfoAttributes) InternalRef() (terra.Reference, error) {
	return bngi.ref, nil
}

func (bngi BrokerNodeGroupInfoAttributes) InternalWithRef(ref terra.Reference) BrokerNodeGroupInfoAttributes {
	return BrokerNodeGroupInfoAttributes{ref: ref}
}

func (bngi BrokerNodeGroupInfoAttributes) InternalTokens() hclwrite.Tokens {
	return bngi.ref.InternalTokens()
}

func (bngi BrokerNodeGroupInfoAttributes) AzDistribution() terra.StringValue {
	return terra.ReferenceAsString(bngi.ref.Append("az_distribution"))
}

func (bngi BrokerNodeGroupInfoAttributes) ClientSubnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bngi.ref.Append("client_subnets"))
}

func (bngi BrokerNodeGroupInfoAttributes) EbsVolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(bngi.ref.Append("ebs_volume_size"))
}

func (bngi BrokerNodeGroupInfoAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(bngi.ref.Append("instance_type"))
}

func (bngi BrokerNodeGroupInfoAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bngi.ref.Append("security_groups"))
}

func (bngi BrokerNodeGroupInfoAttributes) ConnectivityInfo() terra.ListValue[ConnectivityInfoAttributes] {
	return terra.ReferenceAsList[ConnectivityInfoAttributes](bngi.ref.Append("connectivity_info"))
}

func (bngi BrokerNodeGroupInfoAttributes) StorageInfo() terra.ListValue[StorageInfoAttributes] {
	return terra.ReferenceAsList[StorageInfoAttributes](bngi.ref.Append("storage_info"))
}

type ConnectivityInfoAttributes struct {
	ref terra.Reference
}

func (ci ConnectivityInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci ConnectivityInfoAttributes) InternalWithRef(ref terra.Reference) ConnectivityInfoAttributes {
	return ConnectivityInfoAttributes{ref: ref}
}

func (ci ConnectivityInfoAttributes) InternalTokens() hclwrite.Tokens {
	return ci.ref.InternalTokens()
}

func (ci ConnectivityInfoAttributes) PublicAccess() terra.ListValue[PublicAccessAttributes] {
	return terra.ReferenceAsList[PublicAccessAttributes](ci.ref.Append("public_access"))
}

type PublicAccessAttributes struct {
	ref terra.Reference
}

func (pa PublicAccessAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PublicAccessAttributes) InternalWithRef(ref terra.Reference) PublicAccessAttributes {
	return PublicAccessAttributes{ref: ref}
}

func (pa PublicAccessAttributes) InternalTokens() hclwrite.Tokens {
	return pa.ref.InternalTokens()
}

func (pa PublicAccessAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("type"))
}

type StorageInfoAttributes struct {
	ref terra.Reference
}

func (si StorageInfoAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si StorageInfoAttributes) InternalWithRef(ref terra.Reference) StorageInfoAttributes {
	return StorageInfoAttributes{ref: ref}
}

func (si StorageInfoAttributes) InternalTokens() hclwrite.Tokens {
	return si.ref.InternalTokens()
}

func (si StorageInfoAttributes) EbsStorageInfo() terra.ListValue[EbsStorageInfoAttributes] {
	return terra.ReferenceAsList[EbsStorageInfoAttributes](si.ref.Append("ebs_storage_info"))
}

type EbsStorageInfoAttributes struct {
	ref terra.Reference
}

func (esi EbsStorageInfoAttributes) InternalRef() (terra.Reference, error) {
	return esi.ref, nil
}

func (esi EbsStorageInfoAttributes) InternalWithRef(ref terra.Reference) EbsStorageInfoAttributes {
	return EbsStorageInfoAttributes{ref: ref}
}

func (esi EbsStorageInfoAttributes) InternalTokens() hclwrite.Tokens {
	return esi.ref.InternalTokens()
}

func (esi EbsStorageInfoAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(esi.ref.Append("volume_size"))
}

func (esi EbsStorageInfoAttributes) ProvisionedThroughput() terra.ListValue[ProvisionedThroughputAttributes] {
	return terra.ReferenceAsList[ProvisionedThroughputAttributes](esi.ref.Append("provisioned_throughput"))
}

type ProvisionedThroughputAttributes struct {
	ref terra.Reference
}

func (pt ProvisionedThroughputAttributes) InternalRef() (terra.Reference, error) {
	return pt.ref, nil
}

func (pt ProvisionedThroughputAttributes) InternalWithRef(ref terra.Reference) ProvisionedThroughputAttributes {
	return ProvisionedThroughputAttributes{ref: ref}
}

func (pt ProvisionedThroughputAttributes) InternalTokens() hclwrite.Tokens {
	return pt.ref.InternalTokens()
}

func (pt ProvisionedThroughputAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pt.ref.Append("enabled"))
}

func (pt ProvisionedThroughputAttributes) VolumeThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(pt.ref.Append("volume_throughput"))
}

type ClientAuthenticationAttributes struct {
	ref terra.Reference
}

func (ca ClientAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca ClientAuthenticationAttributes) InternalWithRef(ref terra.Reference) ClientAuthenticationAttributes {
	return ClientAuthenticationAttributes{ref: ref}
}

func (ca ClientAuthenticationAttributes) InternalTokens() hclwrite.Tokens {
	return ca.ref.InternalTokens()
}

func (ca ClientAuthenticationAttributes) Unauthenticated() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("unauthenticated"))
}

func (ca ClientAuthenticationAttributes) Sasl() terra.ListValue[SaslAttributes] {
	return terra.ReferenceAsList[SaslAttributes](ca.ref.Append("sasl"))
}

func (ca ClientAuthenticationAttributes) Tls() terra.ListValue[TlsAttributes] {
	return terra.ReferenceAsList[TlsAttributes](ca.ref.Append("tls"))
}

type SaslAttributes struct {
	ref terra.Reference
}

func (s SaslAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SaslAttributes) InternalWithRef(ref terra.Reference) SaslAttributes {
	return SaslAttributes{ref: ref}
}

func (s SaslAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SaslAttributes) Iam() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("iam"))
}

func (s SaslAttributes) Scram() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("scram"))
}

type TlsAttributes struct {
	ref terra.Reference
}

func (t TlsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TlsAttributes) InternalWithRef(ref terra.Reference) TlsAttributes {
	return TlsAttributes{ref: ref}
}

func (t TlsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TlsAttributes) CertificateAuthorityArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](t.ref.Append("certificate_authority_arns"))
}

type ConfigurationInfoAttributes struct {
	ref terra.Reference
}

func (ci ConfigurationInfoAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci ConfigurationInfoAttributes) InternalWithRef(ref terra.Reference) ConfigurationInfoAttributes {
	return ConfigurationInfoAttributes{ref: ref}
}

func (ci ConfigurationInfoAttributes) InternalTokens() hclwrite.Tokens {
	return ci.ref.InternalTokens()
}

func (ci ConfigurationInfoAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("arn"))
}

func (ci ConfigurationInfoAttributes) Revision() terra.NumberValue {
	return terra.ReferenceAsNumber(ci.ref.Append("revision"))
}

type EncryptionInfoAttributes struct {
	ref terra.Reference
}

func (ei EncryptionInfoAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei EncryptionInfoAttributes) InternalWithRef(ref terra.Reference) EncryptionInfoAttributes {
	return EncryptionInfoAttributes{ref: ref}
}

func (ei EncryptionInfoAttributes) InternalTokens() hclwrite.Tokens {
	return ei.ref.InternalTokens()
}

func (ei EncryptionInfoAttributes) EncryptionAtRestKmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("encryption_at_rest_kms_key_arn"))
}

func (ei EncryptionInfoAttributes) EncryptionInTransit() terra.ListValue[EncryptionInTransitAttributes] {
	return terra.ReferenceAsList[EncryptionInTransitAttributes](ei.ref.Append("encryption_in_transit"))
}

type EncryptionInTransitAttributes struct {
	ref terra.Reference
}

func (eit EncryptionInTransitAttributes) InternalRef() (terra.Reference, error) {
	return eit.ref, nil
}

func (eit EncryptionInTransitAttributes) InternalWithRef(ref terra.Reference) EncryptionInTransitAttributes {
	return EncryptionInTransitAttributes{ref: ref}
}

func (eit EncryptionInTransitAttributes) InternalTokens() hclwrite.Tokens {
	return eit.ref.InternalTokens()
}

func (eit EncryptionInTransitAttributes) ClientBroker() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("client_broker"))
}

func (eit EncryptionInTransitAttributes) InCluster() terra.BoolValue {
	return terra.ReferenceAsBool(eit.ref.Append("in_cluster"))
}

type LoggingInfoAttributes struct {
	ref terra.Reference
}

func (li LoggingInfoAttributes) InternalRef() (terra.Reference, error) {
	return li.ref, nil
}

func (li LoggingInfoAttributes) InternalWithRef(ref terra.Reference) LoggingInfoAttributes {
	return LoggingInfoAttributes{ref: ref}
}

func (li LoggingInfoAttributes) InternalTokens() hclwrite.Tokens {
	return li.ref.InternalTokens()
}

func (li LoggingInfoAttributes) BrokerLogs() terra.ListValue[BrokerLogsAttributes] {
	return terra.ReferenceAsList[BrokerLogsAttributes](li.ref.Append("broker_logs"))
}

type BrokerLogsAttributes struct {
	ref terra.Reference
}

func (bl BrokerLogsAttributes) InternalRef() (terra.Reference, error) {
	return bl.ref, nil
}

func (bl BrokerLogsAttributes) InternalWithRef(ref terra.Reference) BrokerLogsAttributes {
	return BrokerLogsAttributes{ref: ref}
}

func (bl BrokerLogsAttributes) InternalTokens() hclwrite.Tokens {
	return bl.ref.InternalTokens()
}

func (bl BrokerLogsAttributes) CloudwatchLogs() terra.ListValue[CloudwatchLogsAttributes] {
	return terra.ReferenceAsList[CloudwatchLogsAttributes](bl.ref.Append("cloudwatch_logs"))
}

func (bl BrokerLogsAttributes) Firehose() terra.ListValue[FirehoseAttributes] {
	return terra.ReferenceAsList[FirehoseAttributes](bl.ref.Append("firehose"))
}

func (bl BrokerLogsAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](bl.ref.Append("s3"))
}

type CloudwatchLogsAttributes struct {
	ref terra.Reference
}

func (cl CloudwatchLogsAttributes) InternalRef() (terra.Reference, error) {
	return cl.ref, nil
}

func (cl CloudwatchLogsAttributes) InternalWithRef(ref terra.Reference) CloudwatchLogsAttributes {
	return CloudwatchLogsAttributes{ref: ref}
}

func (cl CloudwatchLogsAttributes) InternalTokens() hclwrite.Tokens {
	return cl.ref.InternalTokens()
}

func (cl CloudwatchLogsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cl.ref.Append("enabled"))
}

func (cl CloudwatchLogsAttributes) LogGroup() terra.StringValue {
	return terra.ReferenceAsString(cl.ref.Append("log_group"))
}

type FirehoseAttributes struct {
	ref terra.Reference
}

func (f FirehoseAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FirehoseAttributes) InternalWithRef(ref terra.Reference) FirehoseAttributes {
	return FirehoseAttributes{ref: ref}
}

func (f FirehoseAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FirehoseAttributes) DeliveryStream() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("delivery_stream"))
}

func (f FirehoseAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("enabled"))
}

type S3Attributes struct {
	ref terra.Reference
}

func (s S3Attributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s S3Attributes) InternalWithRef(ref terra.Reference) S3Attributes {
	return S3Attributes{ref: ref}
}

func (s S3Attributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s S3Attributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("bucket"))
}

func (s S3Attributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

func (s S3Attributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("prefix"))
}

type OpenMonitoringAttributes struct {
	ref terra.Reference
}

func (om OpenMonitoringAttributes) InternalRef() (terra.Reference, error) {
	return om.ref, nil
}

func (om OpenMonitoringAttributes) InternalWithRef(ref terra.Reference) OpenMonitoringAttributes {
	return OpenMonitoringAttributes{ref: ref}
}

func (om OpenMonitoringAttributes) InternalTokens() hclwrite.Tokens {
	return om.ref.InternalTokens()
}

func (om OpenMonitoringAttributes) Prometheus() terra.ListValue[PrometheusAttributes] {
	return terra.ReferenceAsList[PrometheusAttributes](om.ref.Append("prometheus"))
}

type PrometheusAttributes struct {
	ref terra.Reference
}

func (p PrometheusAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrometheusAttributes) InternalWithRef(ref terra.Reference) PrometheusAttributes {
	return PrometheusAttributes{ref: ref}
}

func (p PrometheusAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PrometheusAttributes) JmxExporter() terra.ListValue[JmxExporterAttributes] {
	return terra.ReferenceAsList[JmxExporterAttributes](p.ref.Append("jmx_exporter"))
}

func (p PrometheusAttributes) NodeExporter() terra.ListValue[NodeExporterAttributes] {
	return terra.ReferenceAsList[NodeExporterAttributes](p.ref.Append("node_exporter"))
}

type JmxExporterAttributes struct {
	ref terra.Reference
}

func (je JmxExporterAttributes) InternalRef() (terra.Reference, error) {
	return je.ref, nil
}

func (je JmxExporterAttributes) InternalWithRef(ref terra.Reference) JmxExporterAttributes {
	return JmxExporterAttributes{ref: ref}
}

func (je JmxExporterAttributes) InternalTokens() hclwrite.Tokens {
	return je.ref.InternalTokens()
}

func (je JmxExporterAttributes) EnabledInBroker() terra.BoolValue {
	return terra.ReferenceAsBool(je.ref.Append("enabled_in_broker"))
}

type NodeExporterAttributes struct {
	ref terra.Reference
}

func (ne NodeExporterAttributes) InternalRef() (terra.Reference, error) {
	return ne.ref, nil
}

func (ne NodeExporterAttributes) InternalWithRef(ref terra.Reference) NodeExporterAttributes {
	return NodeExporterAttributes{ref: ref}
}

func (ne NodeExporterAttributes) InternalTokens() hclwrite.Tokens {
	return ne.ref.InternalTokens()
}

func (ne NodeExporterAttributes) EnabledInBroker() terra.BoolValue {
	return terra.ReferenceAsBool(ne.ref.Append("enabled_in_broker"))
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

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type BrokerNodeGroupInfoState struct {
	AzDistribution   string                  `json:"az_distribution"`
	ClientSubnets    []string                `json:"client_subnets"`
	EbsVolumeSize    float64                 `json:"ebs_volume_size"`
	InstanceType     string                  `json:"instance_type"`
	SecurityGroups   []string                `json:"security_groups"`
	ConnectivityInfo []ConnectivityInfoState `json:"connectivity_info"`
	StorageInfo      []StorageInfoState      `json:"storage_info"`
}

type ConnectivityInfoState struct {
	PublicAccess []PublicAccessState `json:"public_access"`
}

type PublicAccessState struct {
	Type string `json:"type"`
}

type StorageInfoState struct {
	EbsStorageInfo []EbsStorageInfoState `json:"ebs_storage_info"`
}

type EbsStorageInfoState struct {
	VolumeSize            float64                      `json:"volume_size"`
	ProvisionedThroughput []ProvisionedThroughputState `json:"provisioned_throughput"`
}

type ProvisionedThroughputState struct {
	Enabled          bool    `json:"enabled"`
	VolumeThroughput float64 `json:"volume_throughput"`
}

type ClientAuthenticationState struct {
	Unauthenticated bool        `json:"unauthenticated"`
	Sasl            []SaslState `json:"sasl"`
	Tls             []TlsState  `json:"tls"`
}

type SaslState struct {
	Iam   bool `json:"iam"`
	Scram bool `json:"scram"`
}

type TlsState struct {
	CertificateAuthorityArns []string `json:"certificate_authority_arns"`
}

type ConfigurationInfoState struct {
	Arn      string  `json:"arn"`
	Revision float64 `json:"revision"`
}

type EncryptionInfoState struct {
	EncryptionAtRestKmsKeyArn string                     `json:"encryption_at_rest_kms_key_arn"`
	EncryptionInTransit       []EncryptionInTransitState `json:"encryption_in_transit"`
}

type EncryptionInTransitState struct {
	ClientBroker string `json:"client_broker"`
	InCluster    bool   `json:"in_cluster"`
}

type LoggingInfoState struct {
	BrokerLogs []BrokerLogsState `json:"broker_logs"`
}

type BrokerLogsState struct {
	CloudwatchLogs []CloudwatchLogsState `json:"cloudwatch_logs"`
	Firehose       []FirehoseState       `json:"firehose"`
	S3             []S3State             `json:"s3"`
}

type CloudwatchLogsState struct {
	Enabled  bool   `json:"enabled"`
	LogGroup string `json:"log_group"`
}

type FirehoseState struct {
	DeliveryStream string `json:"delivery_stream"`
	Enabled        bool   `json:"enabled"`
}

type S3State struct {
	Bucket  string `json:"bucket"`
	Enabled bool   `json:"enabled"`
	Prefix  string `json:"prefix"`
}

type OpenMonitoringState struct {
	Prometheus []PrometheusState `json:"prometheus"`
}

type PrometheusState struct {
	JmxExporter  []JmxExporterState  `json:"jmx_exporter"`
	NodeExporter []NodeExporterState `json:"node_exporter"`
}

type JmxExporterState struct {
	EnabledInBroker bool `json:"enabled_in_broker"`
}

type NodeExporterState struct {
	EnabledInBroker bool `json:"enabled_in_broker"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
