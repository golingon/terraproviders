// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebtypecomputeiambinding "github.com/golingon/terraproviders/google/4.77.0/iapwebtypecomputeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebTypeComputeIamBinding creates a new instance of [IapWebTypeComputeIamBinding].
func NewIapWebTypeComputeIamBinding(name string, args IapWebTypeComputeIamBindingArgs) *IapWebTypeComputeIamBinding {
	return &IapWebTypeComputeIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebTypeComputeIamBinding)(nil)

// IapWebTypeComputeIamBinding represents the Terraform resource google_iap_web_type_compute_iam_binding.
type IapWebTypeComputeIamBinding struct {
	Name      string
	Args      IapWebTypeComputeIamBindingArgs
	state     *iapWebTypeComputeIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) Type() string {
	return "google_iap_web_type_compute_iam_binding"
}

// LocalName returns the local name for [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) LocalName() string {
	return iwtcib.Name
}

// Configuration returns the configuration (args) for [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) Configuration() interface{} {
	return iwtcib.Args
}

// DependOn is used for other resources to depend on [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iwtcib)
}

// Dependencies returns the list of resources [IapWebTypeComputeIamBinding] depends_on.
func (iwtcib *IapWebTypeComputeIamBinding) Dependencies() terra.Dependencies {
	return iwtcib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iwtcib.Lifecycle
}

// Attributes returns the attributes for [IapWebTypeComputeIamBinding].
func (iwtcib *IapWebTypeComputeIamBinding) Attributes() iapWebTypeComputeIamBindingAttributes {
	return iapWebTypeComputeIamBindingAttributes{ref: terra.ReferenceResource(iwtcib)}
}

// ImportState imports the given attribute values into [IapWebTypeComputeIamBinding]'s state.
func (iwtcib *IapWebTypeComputeIamBinding) ImportState(av io.Reader) error {
	iwtcib.state = &iapWebTypeComputeIamBindingState{}
	if err := json.NewDecoder(av).Decode(iwtcib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwtcib.Type(), iwtcib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebTypeComputeIamBinding] has state.
func (iwtcib *IapWebTypeComputeIamBinding) State() (*iapWebTypeComputeIamBindingState, bool) {
	return iwtcib.state, iwtcib.state != nil
}

// StateMust returns the state for [IapWebTypeComputeIamBinding]. Panics if the state is nil.
func (iwtcib *IapWebTypeComputeIamBinding) StateMust() *iapWebTypeComputeIamBindingState {
	if iwtcib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwtcib.Type(), iwtcib.LocalName()))
	}
	return iwtcib.state
}

// IapWebTypeComputeIamBindingArgs contains the configurations for google_iap_web_type_compute_iam_binding.
type IapWebTypeComputeIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebtypecomputeiambinding.Condition `hcl:"condition,block"`
}
type iapWebTypeComputeIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_type_compute_iam_binding.
func (iwtcib iapWebTypeComputeIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtcib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_compute_iam_binding.
func (iwtcib iapWebTypeComputeIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtcib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_web_type_compute_iam_binding.
func (iwtcib iapWebTypeComputeIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iwtcib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_web_type_compute_iam_binding.
func (iwtcib iapWebTypeComputeIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtcib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_type_compute_iam_binding.
func (iwtcib iapWebTypeComputeIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwtcib.ref.Append("role"))
}

func (iwtcib iapWebTypeComputeIamBindingAttributes) Condition() terra.ListValue[iapwebtypecomputeiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebtypecomputeiambinding.ConditionAttributes](iwtcib.ref.Append("condition"))
}

type iapWebTypeComputeIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Members   []string                                     `json:"members"`
	Project   string                                       `json:"project"`
	Role      string                                       `json:"role"`
	Condition []iapwebtypecomputeiambinding.ConditionState `json:"condition"`
}
