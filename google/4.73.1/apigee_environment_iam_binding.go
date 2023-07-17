// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvironmentiambinding "github.com/golingon/terraproviders/google/4.73.1/apigeeenvironmentiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvironmentIamBinding creates a new instance of [ApigeeEnvironmentIamBinding].
func NewApigeeEnvironmentIamBinding(name string, args ApigeeEnvironmentIamBindingArgs) *ApigeeEnvironmentIamBinding {
	return &ApigeeEnvironmentIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvironmentIamBinding)(nil)

// ApigeeEnvironmentIamBinding represents the Terraform resource google_apigee_environment_iam_binding.
type ApigeeEnvironmentIamBinding struct {
	Name      string
	Args      ApigeeEnvironmentIamBindingArgs
	state     *apigeeEnvironmentIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) Type() string {
	return "google_apigee_environment_iam_binding"
}

// LocalName returns the local name for [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) LocalName() string {
	return aeib.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) Configuration() interface{} {
	return aeib.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(aeib)
}

// Dependencies returns the list of resources [ApigeeEnvironmentIamBinding] depends_on.
func (aeib *ApigeeEnvironmentIamBinding) Dependencies() terra.Dependencies {
	return aeib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) LifecycleManagement() *terra.Lifecycle {
	return aeib.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvironmentIamBinding].
func (aeib *ApigeeEnvironmentIamBinding) Attributes() apigeeEnvironmentIamBindingAttributes {
	return apigeeEnvironmentIamBindingAttributes{ref: terra.ReferenceResource(aeib)}
}

// ImportState imports the given attribute values into [ApigeeEnvironmentIamBinding]'s state.
func (aeib *ApigeeEnvironmentIamBinding) ImportState(av io.Reader) error {
	aeib.state = &apigeeEnvironmentIamBindingState{}
	if err := json.NewDecoder(av).Decode(aeib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeib.Type(), aeib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvironmentIamBinding] has state.
func (aeib *ApigeeEnvironmentIamBinding) State() (*apigeeEnvironmentIamBindingState, bool) {
	return aeib.state, aeib.state != nil
}

// StateMust returns the state for [ApigeeEnvironmentIamBinding]. Panics if the state is nil.
func (aeib *ApigeeEnvironmentIamBinding) StateMust() *apigeeEnvironmentIamBindingState {
	if aeib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeib.Type(), aeib.LocalName()))
	}
	return aeib.state
}

// ApigeeEnvironmentIamBindingArgs contains the configurations for google_apigee_environment_iam_binding.
type ApigeeEnvironmentIamBindingArgs struct {
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigeeenvironmentiambinding.Condition `hcl:"condition,block"`
}
type apigeeEnvironmentIamBindingAttributes struct {
	ref terra.Reference
}

// EnvId returns a reference to field env_id of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(aeib.ref.Append("env_id"))
}

// Etag returns a reference to field etag of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(aeib.ref.Append("etag"))
}

// Id returns a reference to field id of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeib.ref.Append("id"))
}

// Members returns a reference to field members of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aeib.ref.Append("members"))
}

// OrgId returns a reference to field org_id of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(aeib.ref.Append("org_id"))
}

// Role returns a reference to field role of google_apigee_environment_iam_binding.
func (aeib apigeeEnvironmentIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(aeib.ref.Append("role"))
}

func (aeib apigeeEnvironmentIamBindingAttributes) Condition() terra.ListValue[apigeeenvironmentiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[apigeeenvironmentiambinding.ConditionAttributes](aeib.ref.Append("condition"))
}

type apigeeEnvironmentIamBindingState struct {
	EnvId     string                                       `json:"env_id"`
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Members   []string                                     `json:"members"`
	OrgId     string                                       `json:"org_id"`
	Role      string                                       `json:"role"`
	Condition []apigeeenvironmentiambinding.ConditionState `json:"condition"`
}
