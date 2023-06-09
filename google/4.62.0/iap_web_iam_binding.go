// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebiambinding "github.com/golingon/terraproviders/google/4.62.0/iapwebiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebIamBinding creates a new instance of [IapWebIamBinding].
func NewIapWebIamBinding(name string, args IapWebIamBindingArgs) *IapWebIamBinding {
	return &IapWebIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebIamBinding)(nil)

// IapWebIamBinding represents the Terraform resource google_iap_web_iam_binding.
type IapWebIamBinding struct {
	Name      string
	Args      IapWebIamBindingArgs
	state     *iapWebIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebIamBinding].
func (iwib *IapWebIamBinding) Type() string {
	return "google_iap_web_iam_binding"
}

// LocalName returns the local name for [IapWebIamBinding].
func (iwib *IapWebIamBinding) LocalName() string {
	return iwib.Name
}

// Configuration returns the configuration (args) for [IapWebIamBinding].
func (iwib *IapWebIamBinding) Configuration() interface{} {
	return iwib.Args
}

// DependOn is used for other resources to depend on [IapWebIamBinding].
func (iwib *IapWebIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iwib)
}

// Dependencies returns the list of resources [IapWebIamBinding] depends_on.
func (iwib *IapWebIamBinding) Dependencies() terra.Dependencies {
	return iwib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebIamBinding].
func (iwib *IapWebIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iwib.Lifecycle
}

// Attributes returns the attributes for [IapWebIamBinding].
func (iwib *IapWebIamBinding) Attributes() iapWebIamBindingAttributes {
	return iapWebIamBindingAttributes{ref: terra.ReferenceResource(iwib)}
}

// ImportState imports the given attribute values into [IapWebIamBinding]'s state.
func (iwib *IapWebIamBinding) ImportState(av io.Reader) error {
	iwib.state = &iapWebIamBindingState{}
	if err := json.NewDecoder(av).Decode(iwib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwib.Type(), iwib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebIamBinding] has state.
func (iwib *IapWebIamBinding) State() (*iapWebIamBindingState, bool) {
	return iwib.state, iwib.state != nil
}

// StateMust returns the state for [IapWebIamBinding]. Panics if the state is nil.
func (iwib *IapWebIamBinding) StateMust() *iapWebIamBindingState {
	if iwib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwib.Type(), iwib.LocalName()))
	}
	return iwib.state
}

// IapWebIamBindingArgs contains the configurations for google_iap_web_iam_binding.
type IapWebIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebiambinding.Condition `hcl:"condition,block"`
}
type iapWebIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_iam_binding.
func (iwib iapWebIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_iam_binding.
func (iwib iapWebIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_web_iam_binding.
func (iwib iapWebIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iwib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_web_iam_binding.
func (iwib iapWebIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_iam_binding.
func (iwib iapWebIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwib.ref.Append("role"))
}

func (iwib iapWebIamBindingAttributes) Condition() terra.ListValue[iapwebiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebiambinding.ConditionAttributes](iwib.ref.Append("condition"))
}

type iapWebIamBindingState struct {
	Etag      string                            `json:"etag"`
	Id        string                            `json:"id"`
	Members   []string                          `json:"members"`
	Project   string                            `json:"project"`
	Role      string                            `json:"role"`
	Condition []iapwebiambinding.ConditionState `json:"condition"`
}
