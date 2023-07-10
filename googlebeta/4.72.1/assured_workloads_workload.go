// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	assuredworkloadsworkload "github.com/golingon/terraproviders/googlebeta/4.72.1/assuredworkloadsworkload"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAssuredWorkloadsWorkload creates a new instance of [AssuredWorkloadsWorkload].
func NewAssuredWorkloadsWorkload(name string, args AssuredWorkloadsWorkloadArgs) *AssuredWorkloadsWorkload {
	return &AssuredWorkloadsWorkload{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AssuredWorkloadsWorkload)(nil)

// AssuredWorkloadsWorkload represents the Terraform resource google_assured_workloads_workload.
type AssuredWorkloadsWorkload struct {
	Name      string
	Args      AssuredWorkloadsWorkloadArgs
	state     *assuredWorkloadsWorkloadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) Type() string {
	return "google_assured_workloads_workload"
}

// LocalName returns the local name for [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) LocalName() string {
	return aww.Name
}

// Configuration returns the configuration (args) for [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) Configuration() interface{} {
	return aww.Args
}

// DependOn is used for other resources to depend on [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) DependOn() terra.Reference {
	return terra.ReferenceResource(aww)
}

// Dependencies returns the list of resources [AssuredWorkloadsWorkload] depends_on.
func (aww *AssuredWorkloadsWorkload) Dependencies() terra.Dependencies {
	return aww.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) LifecycleManagement() *terra.Lifecycle {
	return aww.Lifecycle
}

// Attributes returns the attributes for [AssuredWorkloadsWorkload].
func (aww *AssuredWorkloadsWorkload) Attributes() assuredWorkloadsWorkloadAttributes {
	return assuredWorkloadsWorkloadAttributes{ref: terra.ReferenceResource(aww)}
}

// ImportState imports the given attribute values into [AssuredWorkloadsWorkload]'s state.
func (aww *AssuredWorkloadsWorkload) ImportState(av io.Reader) error {
	aww.state = &assuredWorkloadsWorkloadState{}
	if err := json.NewDecoder(av).Decode(aww.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aww.Type(), aww.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AssuredWorkloadsWorkload] has state.
func (aww *AssuredWorkloadsWorkload) State() (*assuredWorkloadsWorkloadState, bool) {
	return aww.state, aww.state != nil
}

// StateMust returns the state for [AssuredWorkloadsWorkload]. Panics if the state is nil.
func (aww *AssuredWorkloadsWorkload) StateMust() *assuredWorkloadsWorkloadState {
	if aww.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aww.Type(), aww.LocalName()))
	}
	return aww.state
}

// AssuredWorkloadsWorkloadArgs contains the configurations for google_assured_workloads_workload.
type AssuredWorkloadsWorkloadArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// ComplianceRegime: string, required
	ComplianceRegime terra.StringValue `hcl:"compliance_regime,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// ProvisionedResourcesParent: string, optional
	ProvisionedResourcesParent terra.StringValue `hcl:"provisioned_resources_parent,attr"`
	// Resources: min=0
	Resources []assuredworkloadsworkload.Resources `hcl:"resources,block" validate:"min=0"`
	// KmsSettings: optional
	KmsSettings *assuredworkloadsworkload.KmsSettings `hcl:"kms_settings,block"`
	// ResourceSettings: min=0
	ResourceSettings []assuredworkloadsworkload.ResourceSettings `hcl:"resource_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *assuredworkloadsworkload.Timeouts `hcl:"timeouts,block"`
}
type assuredWorkloadsWorkloadAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("billing_account"))
}

// ComplianceRegime returns a reference to field compliance_regime of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) ComplianceRegime() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("compliance_regime"))
}

// CreateTime returns a reference to field create_time of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("display_name"))
}

// Id returns a reference to field id of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("id"))
}

// Labels returns a reference to field labels of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aww.ref.Append("labels"))
}

// Location returns a reference to field location of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("location"))
}

// Name returns a reference to field name of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("name"))
}

// Organization returns a reference to field organization of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("organization"))
}

// ProvisionedResourcesParent returns a reference to field provisioned_resources_parent of google_assured_workloads_workload.
func (aww assuredWorkloadsWorkloadAttributes) ProvisionedResourcesParent() terra.StringValue {
	return terra.ReferenceAsString(aww.ref.Append("provisioned_resources_parent"))
}

func (aww assuredWorkloadsWorkloadAttributes) Resources() terra.ListValue[assuredworkloadsworkload.ResourcesAttributes] {
	return terra.ReferenceAsList[assuredworkloadsworkload.ResourcesAttributes](aww.ref.Append("resources"))
}

func (aww assuredWorkloadsWorkloadAttributes) KmsSettings() terra.ListValue[assuredworkloadsworkload.KmsSettingsAttributes] {
	return terra.ReferenceAsList[assuredworkloadsworkload.KmsSettingsAttributes](aww.ref.Append("kms_settings"))
}

func (aww assuredWorkloadsWorkloadAttributes) ResourceSettings() terra.ListValue[assuredworkloadsworkload.ResourceSettingsAttributes] {
	return terra.ReferenceAsList[assuredworkloadsworkload.ResourceSettingsAttributes](aww.ref.Append("resource_settings"))
}

func (aww assuredWorkloadsWorkloadAttributes) Timeouts() assuredworkloadsworkload.TimeoutsAttributes {
	return terra.ReferenceAsSingle[assuredworkloadsworkload.TimeoutsAttributes](aww.ref.Append("timeouts"))
}

type assuredWorkloadsWorkloadState struct {
	BillingAccount             string                                           `json:"billing_account"`
	ComplianceRegime           string                                           `json:"compliance_regime"`
	CreateTime                 string                                           `json:"create_time"`
	DisplayName                string                                           `json:"display_name"`
	Id                         string                                           `json:"id"`
	Labels                     map[string]string                                `json:"labels"`
	Location                   string                                           `json:"location"`
	Name                       string                                           `json:"name"`
	Organization               string                                           `json:"organization"`
	ProvisionedResourcesParent string                                           `json:"provisioned_resources_parent"`
	Resources                  []assuredworkloadsworkload.ResourcesState        `json:"resources"`
	KmsSettings                []assuredworkloadsworkload.KmsSettingsState      `json:"kms_settings"`
	ResourceSettings           []assuredworkloadsworkload.ResourceSettingsState `json:"resource_settings"`
	Timeouts                   *assuredworkloadsworkload.TimeoutsState          `json:"timeouts"`
}
