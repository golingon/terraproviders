// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebbackendserviceiambinding "github.com/golingon/terraproviders/google/4.64.0/iapwebbackendserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebBackendServiceIamBinding creates a new instance of [IapWebBackendServiceIamBinding].
func NewIapWebBackendServiceIamBinding(name string, args IapWebBackendServiceIamBindingArgs) *IapWebBackendServiceIamBinding {
	return &IapWebBackendServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebBackendServiceIamBinding)(nil)

// IapWebBackendServiceIamBinding represents the Terraform resource google_iap_web_backend_service_iam_binding.
type IapWebBackendServiceIamBinding struct {
	Name      string
	Args      IapWebBackendServiceIamBindingArgs
	state     *iapWebBackendServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) Type() string {
	return "google_iap_web_backend_service_iam_binding"
}

// LocalName returns the local name for [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) LocalName() string {
	return iwbsib.Name
}

// Configuration returns the configuration (args) for [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) Configuration() interface{} {
	return iwbsib.Args
}

// DependOn is used for other resources to depend on [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iwbsib)
}

// Dependencies returns the list of resources [IapWebBackendServiceIamBinding] depends_on.
func (iwbsib *IapWebBackendServiceIamBinding) Dependencies() terra.Dependencies {
	return iwbsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iwbsib.Lifecycle
}

// Attributes returns the attributes for [IapWebBackendServiceIamBinding].
func (iwbsib *IapWebBackendServiceIamBinding) Attributes() iapWebBackendServiceIamBindingAttributes {
	return iapWebBackendServiceIamBindingAttributes{ref: terra.ReferenceResource(iwbsib)}
}

// ImportState imports the given attribute values into [IapWebBackendServiceIamBinding]'s state.
func (iwbsib *IapWebBackendServiceIamBinding) ImportState(av io.Reader) error {
	iwbsib.state = &iapWebBackendServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(iwbsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwbsib.Type(), iwbsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebBackendServiceIamBinding] has state.
func (iwbsib *IapWebBackendServiceIamBinding) State() (*iapWebBackendServiceIamBindingState, bool) {
	return iwbsib.state, iwbsib.state != nil
}

// StateMust returns the state for [IapWebBackendServiceIamBinding]. Panics if the state is nil.
func (iwbsib *IapWebBackendServiceIamBinding) StateMust() *iapWebBackendServiceIamBindingState {
	if iwbsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwbsib.Type(), iwbsib.LocalName()))
	}
	return iwbsib.state
}

// IapWebBackendServiceIamBindingArgs contains the configurations for google_iap_web_backend_service_iam_binding.
type IapWebBackendServiceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// WebBackendService: string, required
	WebBackendService terra.StringValue `hcl:"web_backend_service,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebbackendserviceiambinding.Condition `hcl:"condition,block"`
}
type iapWebBackendServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwbsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwbsib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iwbsib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwbsib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwbsib.ref.Append("role"))
}

// WebBackendService returns a reference to field web_backend_service of google_iap_web_backend_service_iam_binding.
func (iwbsib iapWebBackendServiceIamBindingAttributes) WebBackendService() terra.StringValue {
	return terra.ReferenceAsString(iwbsib.ref.Append("web_backend_service"))
}

func (iwbsib iapWebBackendServiceIamBindingAttributes) Condition() terra.ListValue[iapwebbackendserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebbackendserviceiambinding.ConditionAttributes](iwbsib.ref.Append("condition"))
}

type iapWebBackendServiceIamBindingState struct {
	Etag              string                                          `json:"etag"`
	Id                string                                          `json:"id"`
	Members           []string                                        `json:"members"`
	Project           string                                          `json:"project"`
	Role              string                                          `json:"role"`
	WebBackendService string                                          `json:"web_backend_service"`
	Condition         []iapwebbackendserviceiambinding.ConditionState `json:"condition"`
}
