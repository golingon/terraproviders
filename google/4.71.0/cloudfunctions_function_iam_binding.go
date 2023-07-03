// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctionsfunctioniambinding "github.com/golingon/terraproviders/google/4.71.0/cloudfunctionsfunctioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctionsFunctionIamBinding creates a new instance of [CloudfunctionsFunctionIamBinding].
func NewCloudfunctionsFunctionIamBinding(name string, args CloudfunctionsFunctionIamBindingArgs) *CloudfunctionsFunctionIamBinding {
	return &CloudfunctionsFunctionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfunctionsFunctionIamBinding)(nil)

// CloudfunctionsFunctionIamBinding represents the Terraform resource google_cloudfunctions_function_iam_binding.
type CloudfunctionsFunctionIamBinding struct {
	Name      string
	Args      CloudfunctionsFunctionIamBindingArgs
	state     *cloudfunctionsFunctionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) Type() string {
	return "google_cloudfunctions_function_iam_binding"
}

// LocalName returns the local name for [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) LocalName() string {
	return cfib.Name
}

// Configuration returns the configuration (args) for [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) Configuration() interface{} {
	return cfib.Args
}

// DependOn is used for other resources to depend on [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cfib)
}

// Dependencies returns the list of resources [CloudfunctionsFunctionIamBinding] depends_on.
func (cfib *CloudfunctionsFunctionIamBinding) Dependencies() terra.Dependencies {
	return cfib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cfib.Lifecycle
}

// Attributes returns the attributes for [CloudfunctionsFunctionIamBinding].
func (cfib *CloudfunctionsFunctionIamBinding) Attributes() cloudfunctionsFunctionIamBindingAttributes {
	return cloudfunctionsFunctionIamBindingAttributes{ref: terra.ReferenceResource(cfib)}
}

// ImportState imports the given attribute values into [CloudfunctionsFunctionIamBinding]'s state.
func (cfib *CloudfunctionsFunctionIamBinding) ImportState(av io.Reader) error {
	cfib.state = &cloudfunctionsFunctionIamBindingState{}
	if err := json.NewDecoder(av).Decode(cfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfib.Type(), cfib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfunctionsFunctionIamBinding] has state.
func (cfib *CloudfunctionsFunctionIamBinding) State() (*cloudfunctionsFunctionIamBindingState, bool) {
	return cfib.state, cfib.state != nil
}

// StateMust returns the state for [CloudfunctionsFunctionIamBinding]. Panics if the state is nil.
func (cfib *CloudfunctionsFunctionIamBinding) StateMust() *cloudfunctionsFunctionIamBindingState {
	if cfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfib.Type(), cfib.LocalName()))
	}
	return cfib.state
}

// CloudfunctionsFunctionIamBindingArgs contains the configurations for google_cloudfunctions_function_iam_binding.
type CloudfunctionsFunctionIamBindingArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudfunctionsfunctioniambinding.Condition `hcl:"condition,block"`
}
type cloudfunctionsFunctionIamBindingAttributes struct {
	ref terra.Reference
}

// CloudFunction returns a reference to field cloud_function of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("cloud_function"))
}

// Etag returns a reference to field etag of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("id"))
}

// Members returns a reference to field members of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfib.ref.Append("members"))
}

// Project returns a reference to field project of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("region"))
}

// Role returns a reference to field role of google_cloudfunctions_function_iam_binding.
func (cfib cloudfunctionsFunctionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("role"))
}

func (cfib cloudfunctionsFunctionIamBindingAttributes) Condition() terra.ListValue[cloudfunctionsfunctioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunctioniambinding.ConditionAttributes](cfib.ref.Append("condition"))
}

type cloudfunctionsFunctionIamBindingState struct {
	CloudFunction string                                            `json:"cloud_function"`
	Etag          string                                            `json:"etag"`
	Id            string                                            `json:"id"`
	Members       []string                                          `json:"members"`
	Project       string                                            `json:"project"`
	Region        string                                            `json:"region"`
	Role          string                                            `json:"role"`
	Condition     []cloudfunctionsfunctioniambinding.ConditionState `json:"condition"`
}
