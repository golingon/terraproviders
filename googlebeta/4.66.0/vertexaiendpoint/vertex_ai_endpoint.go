// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vertexaiendpoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeployedModels struct {
	// AutomaticResources: min=0
	AutomaticResources []AutomaticResources `hcl:"automatic_resources,block" validate:"min=0"`
	// DedicatedResources: min=0
	DedicatedResources []DedicatedResources `hcl:"dedicated_resources,block" validate:"min=0"`
	// PrivateEndpoints: min=0
	PrivateEndpoints []PrivateEndpoints `hcl:"private_endpoints,block" validate:"min=0"`
}

type AutomaticResources struct{}

type DedicatedResources struct {
	// AutoscalingMetricSpecs: min=0
	AutoscalingMetricSpecs []AutoscalingMetricSpecs `hcl:"autoscaling_metric_specs,block" validate:"min=0"`
	// MachineSpec: min=0
	MachineSpec []MachineSpec `hcl:"machine_spec,block" validate:"min=0"`
}

type AutoscalingMetricSpecs struct{}

type MachineSpec struct{}

type PrivateEndpoints struct{}

type EncryptionSpec struct {
	// KmsKeyName: string, required
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DeployedModelsAttributes struct {
	ref terra.Reference
}

func (dm DeployedModelsAttributes) InternalRef() (terra.Reference, error) {
	return dm.ref, nil
}

func (dm DeployedModelsAttributes) InternalWithRef(ref terra.Reference) DeployedModelsAttributes {
	return DeployedModelsAttributes{ref: ref}
}

func (dm DeployedModelsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dm.ref.InternalTokens()
}

func (dm DeployedModelsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("create_time"))
}

func (dm DeployedModelsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("display_name"))
}

func (dm DeployedModelsAttributes) EnableAccessLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dm.ref.Append("enable_access_logging"))
}

func (dm DeployedModelsAttributes) EnableContainerLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dm.ref.Append("enable_container_logging"))
}

func (dm DeployedModelsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("id"))
}

func (dm DeployedModelsAttributes) Model() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("model"))
}

func (dm DeployedModelsAttributes) ModelVersionId() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("model_version_id"))
}

func (dm DeployedModelsAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("service_account"))
}

func (dm DeployedModelsAttributes) SharedResources() terra.StringValue {
	return terra.ReferenceAsString(dm.ref.Append("shared_resources"))
}

func (dm DeployedModelsAttributes) AutomaticResources() terra.ListValue[AutomaticResourcesAttributes] {
	return terra.ReferenceAsList[AutomaticResourcesAttributes](dm.ref.Append("automatic_resources"))
}

func (dm DeployedModelsAttributes) DedicatedResources() terra.ListValue[DedicatedResourcesAttributes] {
	return terra.ReferenceAsList[DedicatedResourcesAttributes](dm.ref.Append("dedicated_resources"))
}

func (dm DeployedModelsAttributes) PrivateEndpoints() terra.ListValue[PrivateEndpointsAttributes] {
	return terra.ReferenceAsList[PrivateEndpointsAttributes](dm.ref.Append("private_endpoints"))
}

type AutomaticResourcesAttributes struct {
	ref terra.Reference
}

func (ar AutomaticResourcesAttributes) InternalRef() (terra.Reference, error) {
	return ar.ref, nil
}

func (ar AutomaticResourcesAttributes) InternalWithRef(ref terra.Reference) AutomaticResourcesAttributes {
	return AutomaticResourcesAttributes{ref: ref}
}

func (ar AutomaticResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ar.ref.InternalTokens()
}

func (ar AutomaticResourcesAttributes) MaxReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("max_replica_count"))
}

func (ar AutomaticResourcesAttributes) MinReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("min_replica_count"))
}

type DedicatedResourcesAttributes struct {
	ref terra.Reference
}

func (dr DedicatedResourcesAttributes) InternalRef() (terra.Reference, error) {
	return dr.ref, nil
}

func (dr DedicatedResourcesAttributes) InternalWithRef(ref terra.Reference) DedicatedResourcesAttributes {
	return DedicatedResourcesAttributes{ref: ref}
}

func (dr DedicatedResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dr.ref.InternalTokens()
}

func (dr DedicatedResourcesAttributes) MaxReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dr.ref.Append("max_replica_count"))
}

func (dr DedicatedResourcesAttributes) MinReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dr.ref.Append("min_replica_count"))
}

func (dr DedicatedResourcesAttributes) AutoscalingMetricSpecs() terra.ListValue[AutoscalingMetricSpecsAttributes] {
	return terra.ReferenceAsList[AutoscalingMetricSpecsAttributes](dr.ref.Append("autoscaling_metric_specs"))
}

func (dr DedicatedResourcesAttributes) MachineSpec() terra.ListValue[MachineSpecAttributes] {
	return terra.ReferenceAsList[MachineSpecAttributes](dr.ref.Append("machine_spec"))
}

type AutoscalingMetricSpecsAttributes struct {
	ref terra.Reference
}

func (ams AutoscalingMetricSpecsAttributes) InternalRef() (terra.Reference, error) {
	return ams.ref, nil
}

func (ams AutoscalingMetricSpecsAttributes) InternalWithRef(ref terra.Reference) AutoscalingMetricSpecsAttributes {
	return AutoscalingMetricSpecsAttributes{ref: ref}
}

func (ams AutoscalingMetricSpecsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ams.ref.InternalTokens()
}

func (ams AutoscalingMetricSpecsAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("metric_name"))
}

func (ams AutoscalingMetricSpecsAttributes) Target() terra.NumberValue {
	return terra.ReferenceAsNumber(ams.ref.Append("target"))
}

type MachineSpecAttributes struct {
	ref terra.Reference
}

func (ms MachineSpecAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MachineSpecAttributes) InternalWithRef(ref terra.Reference) MachineSpecAttributes {
	return MachineSpecAttributes{ref: ref}
}

func (ms MachineSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MachineSpecAttributes) AcceleratorCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("accelerator_count"))
}

func (ms MachineSpecAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("accelerator_type"))
}

func (ms MachineSpecAttributes) MachineType() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("machine_type"))
}

type PrivateEndpointsAttributes struct {
	ref terra.Reference
}

func (pe PrivateEndpointsAttributes) InternalRef() (terra.Reference, error) {
	return pe.ref, nil
}

func (pe PrivateEndpointsAttributes) InternalWithRef(ref terra.Reference) PrivateEndpointsAttributes {
	return PrivateEndpointsAttributes{ref: ref}
}

func (pe PrivateEndpointsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pe.ref.InternalTokens()
}

func (pe PrivateEndpointsAttributes) ExplainHttpUri() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("explain_http_uri"))
}

func (pe PrivateEndpointsAttributes) HealthHttpUri() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("health_http_uri"))
}

func (pe PrivateEndpointsAttributes) PredictHttpUri() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("predict_http_uri"))
}

func (pe PrivateEndpointsAttributes) ServiceAttachment() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("service_attachment"))
}

type EncryptionSpecAttributes struct {
	ref terra.Reference
}

func (es EncryptionSpecAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EncryptionSpecAttributes) InternalWithRef(ref terra.Reference) EncryptionSpecAttributes {
	return EncryptionSpecAttributes{ref: ref}
}

func (es EncryptionSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EncryptionSpecAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("kms_key_name"))
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

type DeployedModelsState struct {
	CreateTime             string                    `json:"create_time"`
	DisplayName            string                    `json:"display_name"`
	EnableAccessLogging    bool                      `json:"enable_access_logging"`
	EnableContainerLogging bool                      `json:"enable_container_logging"`
	Id                     string                    `json:"id"`
	Model                  string                    `json:"model"`
	ModelVersionId         string                    `json:"model_version_id"`
	ServiceAccount         string                    `json:"service_account"`
	SharedResources        string                    `json:"shared_resources"`
	AutomaticResources     []AutomaticResourcesState `json:"automatic_resources"`
	DedicatedResources     []DedicatedResourcesState `json:"dedicated_resources"`
	PrivateEndpoints       []PrivateEndpointsState   `json:"private_endpoints"`
}

type AutomaticResourcesState struct {
	MaxReplicaCount float64 `json:"max_replica_count"`
	MinReplicaCount float64 `json:"min_replica_count"`
}

type DedicatedResourcesState struct {
	MaxReplicaCount        float64                       `json:"max_replica_count"`
	MinReplicaCount        float64                       `json:"min_replica_count"`
	AutoscalingMetricSpecs []AutoscalingMetricSpecsState `json:"autoscaling_metric_specs"`
	MachineSpec            []MachineSpecState            `json:"machine_spec"`
}

type AutoscalingMetricSpecsState struct {
	MetricName string  `json:"metric_name"`
	Target     float64 `json:"target"`
}

type MachineSpecState struct {
	AcceleratorCount float64 `json:"accelerator_count"`
	AcceleratorType  string  `json:"accelerator_type"`
	MachineType      string  `json:"machine_type"`
}

type PrivateEndpointsState struct {
	ExplainHttpUri    string `json:"explain_http_uri"`
	HealthHttpUri     string `json:"health_http_uri"`
	PredictHttpUri    string `json:"predict_http_uri"`
	ServiceAttachment string `json:"service_attachment"`
}

type EncryptionSpecState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}