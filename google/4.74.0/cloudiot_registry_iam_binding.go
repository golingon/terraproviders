// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudiotregistryiambinding "github.com/golingon/terraproviders/google/4.74.0/cloudiotregistryiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudiotRegistryIamBinding creates a new instance of [CloudiotRegistryIamBinding].
func NewCloudiotRegistryIamBinding(name string, args CloudiotRegistryIamBindingArgs) *CloudiotRegistryIamBinding {
	return &CloudiotRegistryIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudiotRegistryIamBinding)(nil)

// CloudiotRegistryIamBinding represents the Terraform resource google_cloudiot_registry_iam_binding.
type CloudiotRegistryIamBinding struct {
	Name      string
	Args      CloudiotRegistryIamBindingArgs
	state     *cloudiotRegistryIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) Type() string {
	return "google_cloudiot_registry_iam_binding"
}

// LocalName returns the local name for [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) LocalName() string {
	return crib.Name
}

// Configuration returns the configuration (args) for [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) Configuration() interface{} {
	return crib.Args
}

// DependOn is used for other resources to depend on [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crib)
}

// Dependencies returns the list of resources [CloudiotRegistryIamBinding] depends_on.
func (crib *CloudiotRegistryIamBinding) Dependencies() terra.Dependencies {
	return crib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) LifecycleManagement() *terra.Lifecycle {
	return crib.Lifecycle
}

// Attributes returns the attributes for [CloudiotRegistryIamBinding].
func (crib *CloudiotRegistryIamBinding) Attributes() cloudiotRegistryIamBindingAttributes {
	return cloudiotRegistryIamBindingAttributes{ref: terra.ReferenceResource(crib)}
}

// ImportState imports the given attribute values into [CloudiotRegistryIamBinding]'s state.
func (crib *CloudiotRegistryIamBinding) ImportState(av io.Reader) error {
	crib.state = &cloudiotRegistryIamBindingState{}
	if err := json.NewDecoder(av).Decode(crib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crib.Type(), crib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudiotRegistryIamBinding] has state.
func (crib *CloudiotRegistryIamBinding) State() (*cloudiotRegistryIamBindingState, bool) {
	return crib.state, crib.state != nil
}

// StateMust returns the state for [CloudiotRegistryIamBinding]. Panics if the state is nil.
func (crib *CloudiotRegistryIamBinding) StateMust() *cloudiotRegistryIamBindingState {
	if crib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crib.Type(), crib.LocalName()))
	}
	return crib.state
}

// CloudiotRegistryIamBindingArgs contains the configurations for google_cloudiot_registry_iam_binding.
type CloudiotRegistryIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudiotregistryiambinding.Condition `hcl:"condition,block"`
}
type cloudiotRegistryIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("id"))
}

// Members returns a reference to field members of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crib.ref.Append("members"))
}

// Name returns a reference to field name of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("region"))
}

// Role returns a reference to field role of google_cloudiot_registry_iam_binding.
func (crib cloudiotRegistryIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crib.ref.Append("role"))
}

func (crib cloudiotRegistryIamBindingAttributes) Condition() terra.ListValue[cloudiotregistryiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudiotregistryiambinding.ConditionAttributes](crib.ref.Append("condition"))
}

type cloudiotRegistryIamBindingState struct {
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Members   []string                                    `json:"members"`
	Name      string                                      `json:"name"`
	Project   string                                      `json:"project"`
	Region    string                                      `json:"region"`
	Role      string                                      `json:"role"`
	Condition []cloudiotregistryiambinding.ConditionState `json:"condition"`
}
