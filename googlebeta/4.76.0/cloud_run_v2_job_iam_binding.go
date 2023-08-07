// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudrunv2jobiambinding "github.com/golingon/terraproviders/googlebeta/4.76.0/cloudrunv2jobiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2JobIamBinding creates a new instance of [CloudRunV2JobIamBinding].
func NewCloudRunV2JobIamBinding(name string, args CloudRunV2JobIamBindingArgs) *CloudRunV2JobIamBinding {
	return &CloudRunV2JobIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2JobIamBinding)(nil)

// CloudRunV2JobIamBinding represents the Terraform resource google_cloud_run_v2_job_iam_binding.
type CloudRunV2JobIamBinding struct {
	Name      string
	Args      CloudRunV2JobIamBindingArgs
	state     *cloudRunV2JobIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) Type() string {
	return "google_cloud_run_v2_job_iam_binding"
}

// LocalName returns the local name for [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) LocalName() string {
	return crvjib.Name
}

// Configuration returns the configuration (args) for [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) Configuration() interface{} {
	return crvjib.Args
}

// DependOn is used for other resources to depend on [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crvjib)
}

// Dependencies returns the list of resources [CloudRunV2JobIamBinding] depends_on.
func (crvjib *CloudRunV2JobIamBinding) Dependencies() terra.Dependencies {
	return crvjib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) LifecycleManagement() *terra.Lifecycle {
	return crvjib.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2JobIamBinding].
func (crvjib *CloudRunV2JobIamBinding) Attributes() cloudRunV2JobIamBindingAttributes {
	return cloudRunV2JobIamBindingAttributes{ref: terra.ReferenceResource(crvjib)}
}

// ImportState imports the given attribute values into [CloudRunV2JobIamBinding]'s state.
func (crvjib *CloudRunV2JobIamBinding) ImportState(av io.Reader) error {
	crvjib.state = &cloudRunV2JobIamBindingState{}
	if err := json.NewDecoder(av).Decode(crvjib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvjib.Type(), crvjib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2JobIamBinding] has state.
func (crvjib *CloudRunV2JobIamBinding) State() (*cloudRunV2JobIamBindingState, bool) {
	return crvjib.state, crvjib.state != nil
}

// StateMust returns the state for [CloudRunV2JobIamBinding]. Panics if the state is nil.
func (crvjib *CloudRunV2JobIamBinding) StateMust() *cloudRunV2JobIamBindingState {
	if crvjib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvjib.Type(), crvjib.LocalName()))
	}
	return crvjib.state
}

// CloudRunV2JobIamBindingArgs contains the configurations for google_cloud_run_v2_job_iam_binding.
type CloudRunV2JobIamBindingArgs struct {
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
	Condition *cloudrunv2jobiambinding.Condition `hcl:"condition,block"`
}
type cloudRunV2JobIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crvjib.ref.Append("members"))
}

// Name returns a reference to field name of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_v2_job_iam_binding.
func (crvjib cloudRunV2JobIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crvjib.ref.Append("role"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Condition() terra.ListValue[cloudrunv2jobiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2jobiambinding.ConditionAttributes](crvjib.ref.Append("condition"))
}

type cloudRunV2JobIamBindingState struct {
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Location  string                                   `json:"location"`
	Members   []string                                 `json:"members"`
	Name      string                                   `json:"name"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	Condition []cloudrunv2jobiambinding.ConditionState `json:"condition"`
}
