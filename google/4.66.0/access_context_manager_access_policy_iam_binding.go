// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanageraccesspolicyiambinding "github.com/golingon/terraproviders/google/4.66.0/accesscontextmanageraccesspolicyiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAccessPolicyIamBinding creates a new instance of [AccessContextManagerAccessPolicyIamBinding].
func NewAccessContextManagerAccessPolicyIamBinding(name string, args AccessContextManagerAccessPolicyIamBindingArgs) *AccessContextManagerAccessPolicyIamBinding {
	return &AccessContextManagerAccessPolicyIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAccessPolicyIamBinding)(nil)

// AccessContextManagerAccessPolicyIamBinding represents the Terraform resource google_access_context_manager_access_policy_iam_binding.
type AccessContextManagerAccessPolicyIamBinding struct {
	Name      string
	Args      AccessContextManagerAccessPolicyIamBindingArgs
	state     *accessContextManagerAccessPolicyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) Type() string {
	return "google_access_context_manager_access_policy_iam_binding"
}

// LocalName returns the local name for [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) LocalName() string {
	return acmapib.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) Configuration() interface{} {
	return acmapib.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(acmapib)
}

// Dependencies returns the list of resources [AccessContextManagerAccessPolicyIamBinding] depends_on.
func (acmapib *AccessContextManagerAccessPolicyIamBinding) Dependencies() terra.Dependencies {
	return acmapib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) LifecycleManagement() *terra.Lifecycle {
	return acmapib.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAccessPolicyIamBinding].
func (acmapib *AccessContextManagerAccessPolicyIamBinding) Attributes() accessContextManagerAccessPolicyIamBindingAttributes {
	return accessContextManagerAccessPolicyIamBindingAttributes{ref: terra.ReferenceResource(acmapib)}
}

// ImportState imports the given attribute values into [AccessContextManagerAccessPolicyIamBinding]'s state.
func (acmapib *AccessContextManagerAccessPolicyIamBinding) ImportState(av io.Reader) error {
	acmapib.state = &accessContextManagerAccessPolicyIamBindingState{}
	if err := json.NewDecoder(av).Decode(acmapib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmapib.Type(), acmapib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAccessPolicyIamBinding] has state.
func (acmapib *AccessContextManagerAccessPolicyIamBinding) State() (*accessContextManagerAccessPolicyIamBindingState, bool) {
	return acmapib.state, acmapib.state != nil
}

// StateMust returns the state for [AccessContextManagerAccessPolicyIamBinding]. Panics if the state is nil.
func (acmapib *AccessContextManagerAccessPolicyIamBinding) StateMust() *accessContextManagerAccessPolicyIamBindingState {
	if acmapib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmapib.Type(), acmapib.LocalName()))
	}
	return acmapib.state
}

// AccessContextManagerAccessPolicyIamBindingArgs contains the configurations for google_access_context_manager_access_policy_iam_binding.
type AccessContextManagerAccessPolicyIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *accesscontextmanageraccesspolicyiambinding.Condition `hcl:"condition,block"`
}
type accessContextManagerAccessPolicyIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_access_context_manager_access_policy_iam_binding.
func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(acmapib.ref.Append("etag"))
}

// Id returns a reference to field id of google_access_context_manager_access_policy_iam_binding.
func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmapib.ref.Append("id"))
}

// Members returns a reference to field members of google_access_context_manager_access_policy_iam_binding.
func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmapib.ref.Append("members"))
}

// Name returns a reference to field name of google_access_context_manager_access_policy_iam_binding.
func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmapib.ref.Append("name"))
}

// Role returns a reference to field role of google_access_context_manager_access_policy_iam_binding.
func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(acmapib.ref.Append("role"))
}

func (acmapib accessContextManagerAccessPolicyIamBindingAttributes) Condition() terra.ListValue[accesscontextmanageraccesspolicyiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[accesscontextmanageraccesspolicyiambinding.ConditionAttributes](acmapib.ref.Append("condition"))
}

type accessContextManagerAccessPolicyIamBindingState struct {
	Etag      string                                                      `json:"etag"`
	Id        string                                                      `json:"id"`
	Members   []string                                                    `json:"members"`
	Name      string                                                      `json:"name"`
	Role      string                                                      `json:"role"`
	Condition []accesscontextmanageraccesspolicyiambinding.ConditionState `json:"condition"`
}
