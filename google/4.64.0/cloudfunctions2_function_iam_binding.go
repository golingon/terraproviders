// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctions2functioniambinding "github.com/golingon/terraproviders/google/4.64.0/cloudfunctions2functioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctions2FunctionIamBinding creates a new instance of [Cloudfunctions2FunctionIamBinding].
func NewCloudfunctions2FunctionIamBinding(name string, args Cloudfunctions2FunctionIamBindingArgs) *Cloudfunctions2FunctionIamBinding {
	return &Cloudfunctions2FunctionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudfunctions2FunctionIamBinding)(nil)

// Cloudfunctions2FunctionIamBinding represents the Terraform resource google_cloudfunctions2_function_iam_binding.
type Cloudfunctions2FunctionIamBinding struct {
	Name      string
	Args      Cloudfunctions2FunctionIamBindingArgs
	state     *cloudfunctions2FunctionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) Type() string {
	return "google_cloudfunctions2_function_iam_binding"
}

// LocalName returns the local name for [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) LocalName() string {
	return cfib.Name
}

// Configuration returns the configuration (args) for [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) Configuration() interface{} {
	return cfib.Args
}

// DependOn is used for other resources to depend on [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cfib)
}

// Dependencies returns the list of resources [Cloudfunctions2FunctionIamBinding] depends_on.
func (cfib *Cloudfunctions2FunctionIamBinding) Dependencies() terra.Dependencies {
	return cfib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cfib.Lifecycle
}

// Attributes returns the attributes for [Cloudfunctions2FunctionIamBinding].
func (cfib *Cloudfunctions2FunctionIamBinding) Attributes() cloudfunctions2FunctionIamBindingAttributes {
	return cloudfunctions2FunctionIamBindingAttributes{ref: terra.ReferenceResource(cfib)}
}

// ImportState imports the given attribute values into [Cloudfunctions2FunctionIamBinding]'s state.
func (cfib *Cloudfunctions2FunctionIamBinding) ImportState(av io.Reader) error {
	cfib.state = &cloudfunctions2FunctionIamBindingState{}
	if err := json.NewDecoder(av).Decode(cfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfib.Type(), cfib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudfunctions2FunctionIamBinding] has state.
func (cfib *Cloudfunctions2FunctionIamBinding) State() (*cloudfunctions2FunctionIamBindingState, bool) {
	return cfib.state, cfib.state != nil
}

// StateMust returns the state for [Cloudfunctions2FunctionIamBinding]. Panics if the state is nil.
func (cfib *Cloudfunctions2FunctionIamBinding) StateMust() *cloudfunctions2FunctionIamBindingState {
	if cfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfib.Type(), cfib.LocalName()))
	}
	return cfib.state
}

// Cloudfunctions2FunctionIamBindingArgs contains the configurations for google_cloudfunctions2_function_iam_binding.
type Cloudfunctions2FunctionIamBindingArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
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
	// Condition: optional
	Condition *cloudfunctions2functioniambinding.Condition `hcl:"condition,block"`
}
type cloudfunctions2FunctionIamBindingAttributes struct {
	ref terra.Reference
}

// CloudFunction returns a reference to field cloud_function of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("cloud_function"))
}

// Etag returns a reference to field etag of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("location"))
}

// Members returns a reference to field members of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfib.ref.Append("members"))
}

// Project returns a reference to field project of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("project"))
}

// Role returns a reference to field role of google_cloudfunctions2_function_iam_binding.
func (cfib cloudfunctions2FunctionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cfib.ref.Append("role"))
}

func (cfib cloudfunctions2FunctionIamBindingAttributes) Condition() terra.ListValue[cloudfunctions2functioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[cloudfunctions2functioniambinding.ConditionAttributes](cfib.ref.Append("condition"))
}

type cloudfunctions2FunctionIamBindingState struct {
	CloudFunction string                                             `json:"cloud_function"`
	Etag          string                                             `json:"etag"`
	Id            string                                             `json:"id"`
	Location      string                                             `json:"location"`
	Members       []string                                           `json:"members"`
	Project       string                                             `json:"project"`
	Role          string                                             `json:"role"`
	Condition     []cloudfunctions2functioniambinding.ConditionState `json:"condition"`
}
