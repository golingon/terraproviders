// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudrunserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.73.1/cloudrunserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunServiceIamBinding creates a new instance of [CloudRunServiceIamBinding].
func NewCloudRunServiceIamBinding(name string, args CloudRunServiceIamBindingArgs) *CloudRunServiceIamBinding {
	return &CloudRunServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunServiceIamBinding)(nil)

// CloudRunServiceIamBinding represents the Terraform resource google_cloud_run_service_iam_binding.
type CloudRunServiceIamBinding struct {
	Name      string
	Args      CloudRunServiceIamBindingArgs
	state     *cloudRunServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) Type() string {
	return "google_cloud_run_service_iam_binding"
}

// LocalName returns the local name for [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) LocalName() string {
	return crsib.Name
}

// Configuration returns the configuration (args) for [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) Configuration() interface{} {
	return crsib.Args
}

// DependOn is used for other resources to depend on [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crsib)
}

// Dependencies returns the list of resources [CloudRunServiceIamBinding] depends_on.
func (crsib *CloudRunServiceIamBinding) Dependencies() terra.Dependencies {
	return crsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return crsib.Lifecycle
}

// Attributes returns the attributes for [CloudRunServiceIamBinding].
func (crsib *CloudRunServiceIamBinding) Attributes() cloudRunServiceIamBindingAttributes {
	return cloudRunServiceIamBindingAttributes{ref: terra.ReferenceResource(crsib)}
}

// ImportState imports the given attribute values into [CloudRunServiceIamBinding]'s state.
func (crsib *CloudRunServiceIamBinding) ImportState(av io.Reader) error {
	crsib.state = &cloudRunServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(crsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsib.Type(), crsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunServiceIamBinding] has state.
func (crsib *CloudRunServiceIamBinding) State() (*cloudRunServiceIamBindingState, bool) {
	return crsib.state, crsib.state != nil
}

// StateMust returns the state for [CloudRunServiceIamBinding]. Panics if the state is nil.
func (crsib *CloudRunServiceIamBinding) StateMust() *cloudRunServiceIamBindingState {
	if crsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsib.Type(), crsib.LocalName()))
	}
	return crsib.state
}

// CloudRunServiceIamBindingArgs contains the configurations for google_cloud_run_service_iam_binding.
type CloudRunServiceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Condition: optional
	Condition *cloudrunserviceiambinding.Condition `hcl:"condition,block"`
}
type cloudRunServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crsib.ref.Append("members"))
}

// Project returns a reference to field project of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("role"))
}

// Service returns a reference to field service of google_cloud_run_service_iam_binding.
func (crsib cloudRunServiceIamBindingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(crsib.ref.Append("service"))
}

func (crsib cloudRunServiceIamBindingAttributes) Condition() terra.ListValue[cloudrunserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudrunserviceiambinding.ConditionAttributes](crsib.ref.Append("condition"))
}

type cloudRunServiceIamBindingState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Location  string                                     `json:"location"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Service   string                                     `json:"service"`
	Condition []cloudrunserviceiambinding.ConditionState `json:"condition"`
}
