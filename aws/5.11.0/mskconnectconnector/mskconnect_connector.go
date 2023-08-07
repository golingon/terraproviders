// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mskconnectconnector

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Capacity struct {
	// Autoscaling: optional
	Autoscaling *Autoscaling `hcl:"autoscaling,block"`
	// ProvisionedCapacity: optional
	ProvisionedCapacity *ProvisionedCapacity `hcl:"provisioned_capacity,block"`
}

type Autoscaling struct {
	// MaxWorkerCount: number, required
	MaxWorkerCount terra.NumberValue `hcl:"max_worker_count,attr" validate:"required"`
	// McuCount: number, optional
	McuCount terra.NumberValue `hcl:"mcu_count,attr"`
	// MinWorkerCount: number, required
	MinWorkerCount terra.NumberValue `hcl:"min_worker_count,attr" validate:"required"`
	// ScaleInPolicy: optional
	ScaleInPolicy *ScaleInPolicy `hcl:"scale_in_policy,block"`
	// ScaleOutPolicy: optional
	ScaleOutPolicy *ScaleOutPolicy `hcl:"scale_out_policy,block"`
}

type ScaleInPolicy struct {
	// CpuUtilizationPercentage: number, optional
	CpuUtilizationPercentage terra.NumberValue `hcl:"cpu_utilization_percentage,attr"`
}

type ScaleOutPolicy struct {
	// CpuUtilizationPercentage: number, optional
	CpuUtilizationPercentage terra.NumberValue `hcl:"cpu_utilization_percentage,attr"`
}

type ProvisionedCapacity struct {
	// McuCount: number, optional
	McuCount terra.NumberValue `hcl:"mcu_count,attr"`
	// WorkerCount: number, required
	WorkerCount terra.NumberValue `hcl:"worker_count,attr" validate:"required"`
}

type KafkaCluster struct {
	// ApacheKafkaCluster: required
	ApacheKafkaCluster *ApacheKafkaCluster `hcl:"apache_kafka_cluster,block" validate:"required"`
}

type ApacheKafkaCluster struct {
	// BootstrapServers: string, required
	BootstrapServers terra.StringValue `hcl:"bootstrap_servers,attr" validate:"required"`
	// Vpc: required
	Vpc *Vpc `hcl:"vpc,block" validate:"required"`
}

type Vpc struct {
	// SecurityGroups: set of string, required
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr" validate:"required"`
	// Subnets: set of string, required
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr" validate:"required"`
}

type KafkaClusterClientAuthentication struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
}

type KafkaClusterEncryptionInTransit struct {
	// EncryptionType: string, optional
	EncryptionType terra.StringValue `hcl:"encryption_type,attr"`
}

type LogDelivery struct {
	// WorkerLogDelivery: required
	WorkerLogDelivery *WorkerLogDelivery `hcl:"worker_log_delivery,block" validate:"required"`
}

type WorkerLogDelivery struct {
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

type Plugin struct {
	// CustomPlugin: required
	CustomPlugin *CustomPlugin `hcl:"custom_plugin,block" validate:"required"`
}

type CustomPlugin struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Revision: number, required
	Revision terra.NumberValue `hcl:"revision,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WorkerConfiguration struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Revision: number, required
	Revision terra.NumberValue `hcl:"revision,attr" validate:"required"`
}

type CapacityAttributes struct {
	ref terra.Reference
}

func (c CapacityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapacityAttributes) InternalWithRef(ref terra.Reference) CapacityAttributes {
	return CapacityAttributes{ref: ref}
}

func (c CapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapacityAttributes) Autoscaling() terra.ListValue[AutoscalingAttributes] {
	return terra.ReferenceAsList[AutoscalingAttributes](c.ref.Append("autoscaling"))
}

func (c CapacityAttributes) ProvisionedCapacity() terra.ListValue[ProvisionedCapacityAttributes] {
	return terra.ReferenceAsList[ProvisionedCapacityAttributes](c.ref.Append("provisioned_capacity"))
}

type AutoscalingAttributes struct {
	ref terra.Reference
}

func (a AutoscalingAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AutoscalingAttributes) InternalWithRef(ref terra.Reference) AutoscalingAttributes {
	return AutoscalingAttributes{ref: ref}
}

func (a AutoscalingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AutoscalingAttributes) MaxWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("max_worker_count"))
}

func (a AutoscalingAttributes) McuCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("mcu_count"))
}

func (a AutoscalingAttributes) MinWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("min_worker_count"))
}

func (a AutoscalingAttributes) ScaleInPolicy() terra.ListValue[ScaleInPolicyAttributes] {
	return terra.ReferenceAsList[ScaleInPolicyAttributes](a.ref.Append("scale_in_policy"))
}

func (a AutoscalingAttributes) ScaleOutPolicy() terra.ListValue[ScaleOutPolicyAttributes] {
	return terra.ReferenceAsList[ScaleOutPolicyAttributes](a.ref.Append("scale_out_policy"))
}

type ScaleInPolicyAttributes struct {
	ref terra.Reference
}

func (sip ScaleInPolicyAttributes) InternalRef() (terra.Reference, error) {
	return sip.ref, nil
}

func (sip ScaleInPolicyAttributes) InternalWithRef(ref terra.Reference) ScaleInPolicyAttributes {
	return ScaleInPolicyAttributes{ref: ref}
}

func (sip ScaleInPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sip.ref.InternalTokens()
}

func (sip ScaleInPolicyAttributes) CpuUtilizationPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sip.ref.Append("cpu_utilization_percentage"))
}

type ScaleOutPolicyAttributes struct {
	ref terra.Reference
}

func (sop ScaleOutPolicyAttributes) InternalRef() (terra.Reference, error) {
	return sop.ref, nil
}

func (sop ScaleOutPolicyAttributes) InternalWithRef(ref terra.Reference) ScaleOutPolicyAttributes {
	return ScaleOutPolicyAttributes{ref: ref}
}

func (sop ScaleOutPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sop.ref.InternalTokens()
}

func (sop ScaleOutPolicyAttributes) CpuUtilizationPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sop.ref.Append("cpu_utilization_percentage"))
}

type ProvisionedCapacityAttributes struct {
	ref terra.Reference
}

func (pc ProvisionedCapacityAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ProvisionedCapacityAttributes) InternalWithRef(ref terra.Reference) ProvisionedCapacityAttributes {
	return ProvisionedCapacityAttributes{ref: ref}
}

func (pc ProvisionedCapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ProvisionedCapacityAttributes) McuCount() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("mcu_count"))
}

func (pc ProvisionedCapacityAttributes) WorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(pc.ref.Append("worker_count"))
}

type KafkaClusterAttributes struct {
	ref terra.Reference
}

func (kc KafkaClusterAttributes) InternalRef() (terra.Reference, error) {
	return kc.ref, nil
}

func (kc KafkaClusterAttributes) InternalWithRef(ref terra.Reference) KafkaClusterAttributes {
	return KafkaClusterAttributes{ref: ref}
}

func (kc KafkaClusterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kc.ref.InternalTokens()
}

func (kc KafkaClusterAttributes) ApacheKafkaCluster() terra.ListValue[ApacheKafkaClusterAttributes] {
	return terra.ReferenceAsList[ApacheKafkaClusterAttributes](kc.ref.Append("apache_kafka_cluster"))
}

type ApacheKafkaClusterAttributes struct {
	ref terra.Reference
}

func (akc ApacheKafkaClusterAttributes) InternalRef() (terra.Reference, error) {
	return akc.ref, nil
}

func (akc ApacheKafkaClusterAttributes) InternalWithRef(ref terra.Reference) ApacheKafkaClusterAttributes {
	return ApacheKafkaClusterAttributes{ref: ref}
}

func (akc ApacheKafkaClusterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return akc.ref.InternalTokens()
}

func (akc ApacheKafkaClusterAttributes) BootstrapServers() terra.StringValue {
	return terra.ReferenceAsString(akc.ref.Append("bootstrap_servers"))
}

func (akc ApacheKafkaClusterAttributes) Vpc() terra.ListValue[VpcAttributes] {
	return terra.ReferenceAsList[VpcAttributes](akc.ref.Append("vpc"))
}

type VpcAttributes struct {
	ref terra.Reference
}

func (v VpcAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VpcAttributes) InternalWithRef(ref terra.Reference) VpcAttributes {
	return VpcAttributes{ref: ref}
}

func (v VpcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VpcAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](v.ref.Append("security_groups"))
}

func (v VpcAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](v.ref.Append("subnets"))
}

type KafkaClusterClientAuthenticationAttributes struct {
	ref terra.Reference
}

func (kcca KafkaClusterClientAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return kcca.ref, nil
}

func (kcca KafkaClusterClientAuthenticationAttributes) InternalWithRef(ref terra.Reference) KafkaClusterClientAuthenticationAttributes {
	return KafkaClusterClientAuthenticationAttributes{ref: ref}
}

func (kcca KafkaClusterClientAuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kcca.ref.InternalTokens()
}

func (kcca KafkaClusterClientAuthenticationAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(kcca.ref.Append("authentication_type"))
}

type KafkaClusterEncryptionInTransitAttributes struct {
	ref terra.Reference
}

func (kceit KafkaClusterEncryptionInTransitAttributes) InternalRef() (terra.Reference, error) {
	return kceit.ref, nil
}

func (kceit KafkaClusterEncryptionInTransitAttributes) InternalWithRef(ref terra.Reference) KafkaClusterEncryptionInTransitAttributes {
	return KafkaClusterEncryptionInTransitAttributes{ref: ref}
}

func (kceit KafkaClusterEncryptionInTransitAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kceit.ref.InternalTokens()
}

func (kceit KafkaClusterEncryptionInTransitAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(kceit.ref.Append("encryption_type"))
}

type LogDeliveryAttributes struct {
	ref terra.Reference
}

func (ld LogDeliveryAttributes) InternalRef() (terra.Reference, error) {
	return ld.ref, nil
}

func (ld LogDeliveryAttributes) InternalWithRef(ref terra.Reference) LogDeliveryAttributes {
	return LogDeliveryAttributes{ref: ref}
}

func (ld LogDeliveryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ld.ref.InternalTokens()
}

func (ld LogDeliveryAttributes) WorkerLogDelivery() terra.ListValue[WorkerLogDeliveryAttributes] {
	return terra.ReferenceAsList[WorkerLogDeliveryAttributes](ld.ref.Append("worker_log_delivery"))
}

type WorkerLogDeliveryAttributes struct {
	ref terra.Reference
}

func (wld WorkerLogDeliveryAttributes) InternalRef() (terra.Reference, error) {
	return wld.ref, nil
}

func (wld WorkerLogDeliveryAttributes) InternalWithRef(ref terra.Reference) WorkerLogDeliveryAttributes {
	return WorkerLogDeliveryAttributes{ref: ref}
}

func (wld WorkerLogDeliveryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wld.ref.InternalTokens()
}

func (wld WorkerLogDeliveryAttributes) CloudwatchLogs() terra.ListValue[CloudwatchLogsAttributes] {
	return terra.ReferenceAsList[CloudwatchLogsAttributes](wld.ref.Append("cloudwatch_logs"))
}

func (wld WorkerLogDeliveryAttributes) Firehose() terra.ListValue[FirehoseAttributes] {
	return terra.ReferenceAsList[FirehoseAttributes](wld.ref.Append("firehose"))
}

func (wld WorkerLogDeliveryAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](wld.ref.Append("s3"))
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

func (cl CloudwatchLogsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (f FirehoseAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (s S3Attributes) InternalTokens() (hclwrite.Tokens, error) {
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

type PluginAttributes struct {
	ref terra.Reference
}

func (p PluginAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PluginAttributes) InternalWithRef(ref terra.Reference) PluginAttributes {
	return PluginAttributes{ref: ref}
}

func (p PluginAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PluginAttributes) CustomPlugin() terra.ListValue[CustomPluginAttributes] {
	return terra.ReferenceAsList[CustomPluginAttributes](p.ref.Append("custom_plugin"))
}

type CustomPluginAttributes struct {
	ref terra.Reference
}

func (cp CustomPluginAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp CustomPluginAttributes) InternalWithRef(ref terra.Reference) CustomPluginAttributes {
	return CustomPluginAttributes{ref: ref}
}

func (cp CustomPluginAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp CustomPluginAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("arn"))
}

func (cp CustomPluginAttributes) Revision() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("revision"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type WorkerConfigurationAttributes struct {
	ref terra.Reference
}

func (wc WorkerConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return wc.ref, nil
}

func (wc WorkerConfigurationAttributes) InternalWithRef(ref terra.Reference) WorkerConfigurationAttributes {
	return WorkerConfigurationAttributes{ref: ref}
}

func (wc WorkerConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wc.ref.InternalTokens()
}

func (wc WorkerConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wc.ref.Append("arn"))
}

func (wc WorkerConfigurationAttributes) Revision() terra.NumberValue {
	return terra.ReferenceAsNumber(wc.ref.Append("revision"))
}

type CapacityState struct {
	Autoscaling         []AutoscalingState         `json:"autoscaling"`
	ProvisionedCapacity []ProvisionedCapacityState `json:"provisioned_capacity"`
}

type AutoscalingState struct {
	MaxWorkerCount float64               `json:"max_worker_count"`
	McuCount       float64               `json:"mcu_count"`
	MinWorkerCount float64               `json:"min_worker_count"`
	ScaleInPolicy  []ScaleInPolicyState  `json:"scale_in_policy"`
	ScaleOutPolicy []ScaleOutPolicyState `json:"scale_out_policy"`
}

type ScaleInPolicyState struct {
	CpuUtilizationPercentage float64 `json:"cpu_utilization_percentage"`
}

type ScaleOutPolicyState struct {
	CpuUtilizationPercentage float64 `json:"cpu_utilization_percentage"`
}

type ProvisionedCapacityState struct {
	McuCount    float64 `json:"mcu_count"`
	WorkerCount float64 `json:"worker_count"`
}

type KafkaClusterState struct {
	ApacheKafkaCluster []ApacheKafkaClusterState `json:"apache_kafka_cluster"`
}

type ApacheKafkaClusterState struct {
	BootstrapServers string     `json:"bootstrap_servers"`
	Vpc              []VpcState `json:"vpc"`
}

type VpcState struct {
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
}

type KafkaClusterClientAuthenticationState struct {
	AuthenticationType string `json:"authentication_type"`
}

type KafkaClusterEncryptionInTransitState struct {
	EncryptionType string `json:"encryption_type"`
}

type LogDeliveryState struct {
	WorkerLogDelivery []WorkerLogDeliveryState `json:"worker_log_delivery"`
}

type WorkerLogDeliveryState struct {
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

type PluginState struct {
	CustomPlugin []CustomPluginState `json:"custom_plugin"`
}

type CustomPluginState struct {
	Arn      string  `json:"arn"`
	Revision float64 `json:"revision"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WorkerConfigurationState struct {
	Arn      string  `json:"arn"`
	Revision float64 `json:"revision"`
}
