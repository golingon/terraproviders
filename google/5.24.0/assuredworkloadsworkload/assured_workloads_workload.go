// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package assuredworkloadsworkload

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ComplianceStatus struct{}

type EkmProvisioningResponse struct{}

type Resources struct{}

type SaaEnrollmentResponse struct{}

type KmsSettings struct {
	// NextRotationTime: string, required
	NextRotationTime terra.StringValue `hcl:"next_rotation_time,attr" validate:"required"`
	// RotationPeriod: string, required
	RotationPeriod terra.StringValue `hcl:"rotation_period,attr" validate:"required"`
}

type PartnerPermissions struct {
	// AssuredWorkloadsMonitoring: bool, optional
	AssuredWorkloadsMonitoring terra.BoolValue `hcl:"assured_workloads_monitoring,attr"`
	// DataLogsViewer: bool, optional
	DataLogsViewer terra.BoolValue `hcl:"data_logs_viewer,attr"`
	// ServiceAccessApprover: bool, optional
	ServiceAccessApprover terra.BoolValue `hcl:"service_access_approver,attr"`
}

type ResourceSettings struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// ResourceId: string, optional
	ResourceId terra.StringValue `hcl:"resource_id,attr"`
	// ResourceType: string, optional
	ResourceType terra.StringValue `hcl:"resource_type,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ComplianceStatusAttributes struct {
	ref terra.Reference
}

func (cs ComplianceStatusAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs ComplianceStatusAttributes) InternalWithRef(ref terra.Reference) ComplianceStatusAttributes {
	return ComplianceStatusAttributes{ref: ref}
}

func (cs ComplianceStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs ComplianceStatusAttributes) AcknowledgedViolationCount() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](cs.ref.Append("acknowledged_violation_count"))
}

func (cs ComplianceStatusAttributes) ActiveViolationCount() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](cs.ref.Append("active_violation_count"))
}

type EkmProvisioningResponseAttributes struct {
	ref terra.Reference
}

func (epr EkmProvisioningResponseAttributes) InternalRef() (terra.Reference, error) {
	return epr.ref, nil
}

func (epr EkmProvisioningResponseAttributes) InternalWithRef(ref terra.Reference) EkmProvisioningResponseAttributes {
	return EkmProvisioningResponseAttributes{ref: ref}
}

func (epr EkmProvisioningResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return epr.ref.InternalTokens()
}

func (epr EkmProvisioningResponseAttributes) EkmProvisioningErrorDomain() terra.StringValue {
	return terra.ReferenceAsString(epr.ref.Append("ekm_provisioning_error_domain"))
}

func (epr EkmProvisioningResponseAttributes) EkmProvisioningErrorMapping() terra.StringValue {
	return terra.ReferenceAsString(epr.ref.Append("ekm_provisioning_error_mapping"))
}

func (epr EkmProvisioningResponseAttributes) EkmProvisioningState() terra.StringValue {
	return terra.ReferenceAsString(epr.ref.Append("ekm_provisioning_state"))
}

type ResourcesAttributes struct {
	ref terra.Reference
}

func (r ResourcesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) ResourceId() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("resource_id"))
}

func (r ResourcesAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("resource_type"))
}

type SaaEnrollmentResponseAttributes struct {
	ref terra.Reference
}

func (ser SaaEnrollmentResponseAttributes) InternalRef() (terra.Reference, error) {
	return ser.ref, nil
}

func (ser SaaEnrollmentResponseAttributes) InternalWithRef(ref terra.Reference) SaaEnrollmentResponseAttributes {
	return SaaEnrollmentResponseAttributes{ref: ref}
}

func (ser SaaEnrollmentResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ser.ref.InternalTokens()
}

func (ser SaaEnrollmentResponseAttributes) SetupErrors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ser.ref.Append("setup_errors"))
}

func (ser SaaEnrollmentResponseAttributes) SetupStatus() terra.StringValue {
	return terra.ReferenceAsString(ser.ref.Append("setup_status"))
}

type KmsSettingsAttributes struct {
	ref terra.Reference
}

func (ks KmsSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ks.ref, nil
}

func (ks KmsSettingsAttributes) InternalWithRef(ref terra.Reference) KmsSettingsAttributes {
	return KmsSettingsAttributes{ref: ref}
}

func (ks KmsSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ks.ref.InternalTokens()
}

func (ks KmsSettingsAttributes) NextRotationTime() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("next_rotation_time"))
}

func (ks KmsSettingsAttributes) RotationPeriod() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("rotation_period"))
}

type PartnerPermissionsAttributes struct {
	ref terra.Reference
}

func (pp PartnerPermissionsAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp PartnerPermissionsAttributes) InternalWithRef(ref terra.Reference) PartnerPermissionsAttributes {
	return PartnerPermissionsAttributes{ref: ref}
}

func (pp PartnerPermissionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp PartnerPermissionsAttributes) AssuredWorkloadsMonitoring() terra.BoolValue {
	return terra.ReferenceAsBool(pp.ref.Append("assured_workloads_monitoring"))
}

func (pp PartnerPermissionsAttributes) DataLogsViewer() terra.BoolValue {
	return terra.ReferenceAsBool(pp.ref.Append("data_logs_viewer"))
}

func (pp PartnerPermissionsAttributes) ServiceAccessApprover() terra.BoolValue {
	return terra.ReferenceAsBool(pp.ref.Append("service_access_approver"))
}

type ResourceSettingsAttributes struct {
	ref terra.Reference
}

func (rs ResourceSettingsAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceSettingsAttributes) InternalWithRef(ref terra.Reference) ResourceSettingsAttributes {
	return ResourceSettingsAttributes{ref: ref}
}

func (rs ResourceSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceSettingsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("display_name"))
}

func (rs ResourceSettingsAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("resource_id"))
}

func (rs ResourceSettingsAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("resource_type"))
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

type ComplianceStatusState struct {
	AcknowledgedViolationCount []float64 `json:"acknowledged_violation_count"`
	ActiveViolationCount       []float64 `json:"active_violation_count"`
}

type EkmProvisioningResponseState struct {
	EkmProvisioningErrorDomain  string `json:"ekm_provisioning_error_domain"`
	EkmProvisioningErrorMapping string `json:"ekm_provisioning_error_mapping"`
	EkmProvisioningState        string `json:"ekm_provisioning_state"`
}

type ResourcesState struct {
	ResourceId   float64 `json:"resource_id"`
	ResourceType string  `json:"resource_type"`
}

type SaaEnrollmentResponseState struct {
	SetupErrors []string `json:"setup_errors"`
	SetupStatus string   `json:"setup_status"`
}

type KmsSettingsState struct {
	NextRotationTime string `json:"next_rotation_time"`
	RotationPeriod   string `json:"rotation_period"`
}

type PartnerPermissionsState struct {
	AssuredWorkloadsMonitoring bool `json:"assured_workloads_monitoring"`
	DataLogsViewer             bool `json:"data_logs_viewer"`
	ServiceAccessApprover      bool `json:"service_access_approver"`
}

type ResourceSettingsState struct {
	DisplayName  string `json:"display_name"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
