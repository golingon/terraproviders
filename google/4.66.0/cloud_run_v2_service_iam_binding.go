// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunv2serviceiambinding "github.com/golingon/terraproviders/google/4.66.0/cloudrunv2serviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2ServiceIamBinding creates a new instance of [CloudRunV2ServiceIamBinding].
func NewCloudRunV2ServiceIamBinding(name string, args CloudRunV2ServiceIamBindingArgs) *CloudRunV2ServiceIamBinding {
	return &CloudRunV2ServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2ServiceIamBinding)(nil)

// CloudRunV2ServiceIamBinding represents the Terraform resource google_cloud_run_v2_service_iam_binding.
type CloudRunV2ServiceIamBinding struct {
	Name      string
	Args      CloudRunV2ServiceIamBindingArgs
	state     *cloudRunV2ServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) Type() string {
	return "google_cloud_run_v2_service_iam_binding"
}

// LocalName returns the local name for [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) LocalName() string {
	return crvsib.Name
}

// Configuration returns the configuration (args) for [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) Configuration() interface{} {
	return crvsib.Args
}

// DependOn is used for other resources to depend on [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crvsib)
}

// Dependencies returns the list of resources [CloudRunV2ServiceIamBinding] depends_on.
func (crvsib *CloudRunV2ServiceIamBinding) Dependencies() terra.Dependencies {
	return crvsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return crvsib.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2ServiceIamBinding].
func (crvsib *CloudRunV2ServiceIamBinding) Attributes() cloudRunV2ServiceIamBindingAttributes {
	return cloudRunV2ServiceIamBindingAttributes{ref: terra.ReferenceResource(crvsib)}
}

// ImportState imports the given attribute values into [CloudRunV2ServiceIamBinding]'s state.
func (crvsib *CloudRunV2ServiceIamBinding) ImportState(av io.Reader) error {
	crvsib.state = &cloudRunV2ServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(crvsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvsib.Type(), crvsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2ServiceIamBinding] has state.
func (crvsib *CloudRunV2ServiceIamBinding) State() (*cloudRunV2ServiceIamBindingState, bool) {
	return crvsib.state, crvsib.state != nil
}

// StateMust returns the state for [CloudRunV2ServiceIamBinding]. Panics if the state is nil.
func (crvsib *CloudRunV2ServiceIamBinding) StateMust() *cloudRunV2ServiceIamBindingState {
	if crvsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvsib.Type(), crvsib.LocalName()))
	}
	return crvsib.state
}

// CloudRunV2ServiceIamBindingArgs contains the configurations for google_cloud_run_v2_service_iam_binding.
type CloudRunV2ServiceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudrunv2serviceiambinding.Condition `hcl:"condition,block"`
}
type cloudRunV2ServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crvsib.ref.Append("members"))
}

// Name returns a reference to field name of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_v2_service_iam_binding.
func (crvsib cloudRunV2ServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crvsib.ref.Append("role"))
}

func (crvsib cloudRunV2ServiceIamBindingAttributes) Condition() terra.ListValue[cloudrunv2serviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2serviceiambinding.ConditionAttributes](crvsib.ref.Append("condition"))
}

type cloudRunV2ServiceIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Location  string                                       `json:"location"`
	Members   []string                                     `json:"members"`
	Name      string                                       `json:"name"`
	Project   string                                       `json:"project"`
	Role      string                                       `json:"role"`
	Condition []cloudrunv2serviceiambinding.ConditionState `json:"condition"`
}
