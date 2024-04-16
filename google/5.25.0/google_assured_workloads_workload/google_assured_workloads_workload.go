// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_assured_workloads_workload

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_assured_workloads_workload.
type Resource struct {
	Name      string
	Args      Args
	state     *googleAssuredWorkloadsWorkloadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gaww *Resource) Type() string {
	return "google_assured_workloads_workload"
}

// LocalName returns the local name for [Resource].
func (gaww *Resource) LocalName() string {
	return gaww.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gaww *Resource) Configuration() interface{} {
	return gaww.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gaww *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gaww)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gaww *Resource) Dependencies() terra.Dependencies {
	return gaww.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gaww *Resource) LifecycleManagement() *terra.Lifecycle {
	return gaww.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gaww *Resource) Attributes() googleAssuredWorkloadsWorkloadAttributes {
	return googleAssuredWorkloadsWorkloadAttributes{ref: terra.ReferenceResource(gaww)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gaww *Resource) ImportState(state io.Reader) error {
	gaww.state = &googleAssuredWorkloadsWorkloadState{}
	if err := json.NewDecoder(state).Decode(gaww.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gaww.Type(), gaww.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gaww *Resource) State() (*googleAssuredWorkloadsWorkloadState, bool) {
	return gaww.state, gaww.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gaww *Resource) StateMust() *googleAssuredWorkloadsWorkloadState {
	if gaww.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gaww.Type(), gaww.LocalName()))
	}
	return gaww.state
}

// Args contains the configurations for google_assured_workloads_workload.
type Args struct {
	// BillingAccount: string, optional
	BillingAccount terra.StringValue `hcl:"billing_account,attr"`
	// ComplianceRegime: string, required
	ComplianceRegime terra.StringValue `hcl:"compliance_regime,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableSovereignControls: bool, optional
	EnableSovereignControls terra.BoolValue `hcl:"enable_sovereign_controls,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Partner: string, optional
	Partner terra.StringValue `hcl:"partner,attr"`
	// ProvisionedResourcesParent: string, optional
	ProvisionedResourcesParent terra.StringValue `hcl:"provisioned_resources_parent,attr"`
	// ViolationNotificationsEnabled: bool, optional
	ViolationNotificationsEnabled terra.BoolValue `hcl:"violation_notifications_enabled,attr"`
	// KmsSettings: optional
	KmsSettings *KmsSettings `hcl:"kms_settings,block"`
	// PartnerPermissions: optional
	PartnerPermissions *PartnerPermissions `hcl:"partner_permissions,block"`
	// ResourceSettings: min=0
	ResourceSettings []ResourceSettings `hcl:"resource_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleAssuredWorkloadsWorkloadAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("billing_account"))
}

// ComplianceRegime returns a reference to field compliance_regime of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) ComplianceRegime() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("compliance_regime"))
}

// CompliantButDisallowedServices returns a reference to field compliant_but_disallowed_services of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) CompliantButDisallowedServices() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gaww.ref.Append("compliant_but_disallowed_services"))
}

// CreateTime returns a reference to field create_time of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaww.ref.Append("effective_labels"))
}

// EnableSovereignControls returns a reference to field enable_sovereign_controls of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) EnableSovereignControls() terra.BoolValue {
	return terra.ReferenceAsBool(gaww.ref.Append("enable_sovereign_controls"))
}

// Id returns a reference to field id of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("id"))
}

// KajEnrollmentState returns a reference to field kaj_enrollment_state of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) KajEnrollmentState() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("kaj_enrollment_state"))
}

// Labels returns a reference to field labels of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaww.ref.Append("labels"))
}

// Location returns a reference to field location of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("location"))
}

// Name returns a reference to field name of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("name"))
}

// Organization returns a reference to field organization of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("organization"))
}

// Partner returns a reference to field partner of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) Partner() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("partner"))
}

// ProvisionedResourcesParent returns a reference to field provisioned_resources_parent of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) ProvisionedResourcesParent() terra.StringValue {
	return terra.ReferenceAsString(gaww.ref.Append("provisioned_resources_parent"))
}

// TerraformLabels returns a reference to field terraform_labels of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaww.ref.Append("terraform_labels"))
}

// ViolationNotificationsEnabled returns a reference to field violation_notifications_enabled of google_assured_workloads_workload.
func (gaww googleAssuredWorkloadsWorkloadAttributes) ViolationNotificationsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(gaww.ref.Append("violation_notifications_enabled"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) ComplianceStatus() terra.ListValue[ComplianceStatusAttributes] {
	return terra.ReferenceAsList[ComplianceStatusAttributes](gaww.ref.Append("compliance_status"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) EkmProvisioningResponse() terra.ListValue[EkmProvisioningResponseAttributes] {
	return terra.ReferenceAsList[EkmProvisioningResponseAttributes](gaww.ref.Append("ekm_provisioning_response"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) Resources() terra.ListValue[ResourcesAttributes] {
	return terra.ReferenceAsList[ResourcesAttributes](gaww.ref.Append("resources"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) SaaEnrollmentResponse() terra.ListValue[SaaEnrollmentResponseAttributes] {
	return terra.ReferenceAsList[SaaEnrollmentResponseAttributes](gaww.ref.Append("saa_enrollment_response"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) KmsSettings() terra.ListValue[KmsSettingsAttributes] {
	return terra.ReferenceAsList[KmsSettingsAttributes](gaww.ref.Append("kms_settings"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) PartnerPermissions() terra.ListValue[PartnerPermissionsAttributes] {
	return terra.ReferenceAsList[PartnerPermissionsAttributes](gaww.ref.Append("partner_permissions"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) ResourceSettings() terra.ListValue[ResourceSettingsAttributes] {
	return terra.ReferenceAsList[ResourceSettingsAttributes](gaww.ref.Append("resource_settings"))
}

func (gaww googleAssuredWorkloadsWorkloadAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gaww.ref.Append("timeouts"))
}

type googleAssuredWorkloadsWorkloadState struct {
	BillingAccount                 string                         `json:"billing_account"`
	ComplianceRegime               string                         `json:"compliance_regime"`
	CompliantButDisallowedServices []string                       `json:"compliant_but_disallowed_services"`
	CreateTime                     string                         `json:"create_time"`
	DisplayName                    string                         `json:"display_name"`
	EffectiveLabels                map[string]string              `json:"effective_labels"`
	EnableSovereignControls        bool                           `json:"enable_sovereign_controls"`
	Id                             string                         `json:"id"`
	KajEnrollmentState             string                         `json:"kaj_enrollment_state"`
	Labels                         map[string]string              `json:"labels"`
	Location                       string                         `json:"location"`
	Name                           string                         `json:"name"`
	Organization                   string                         `json:"organization"`
	Partner                        string                         `json:"partner"`
	ProvisionedResourcesParent     string                         `json:"provisioned_resources_parent"`
	TerraformLabels                map[string]string              `json:"terraform_labels"`
	ViolationNotificationsEnabled  bool                           `json:"violation_notifications_enabled"`
	ComplianceStatus               []ComplianceStatusState        `json:"compliance_status"`
	EkmProvisioningResponse        []EkmProvisioningResponseState `json:"ekm_provisioning_response"`
	Resources                      []ResourcesState               `json:"resources"`
	SaaEnrollmentResponse          []SaaEnrollmentResponseState   `json:"saa_enrollment_response"`
	KmsSettings                    []KmsSettingsState             `json:"kms_settings"`
	PartnerPermissions             []PartnerPermissionsState      `json:"partner_permissions"`
	ResourceSettings               []ResourceSettingsState        `json:"resource_settings"`
	Timeouts                       *TimeoutsState                 `json:"timeouts"`
}
