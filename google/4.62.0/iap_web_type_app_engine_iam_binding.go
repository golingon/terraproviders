// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapwebtypeappengineiambinding "github.com/golingon/terraproviders/google/4.62.0/iapwebtypeappengineiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebTypeAppEngineIamBinding creates a new instance of [IapWebTypeAppEngineIamBinding].
func NewIapWebTypeAppEngineIamBinding(name string, args IapWebTypeAppEngineIamBindingArgs) *IapWebTypeAppEngineIamBinding {
	return &IapWebTypeAppEngineIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebTypeAppEngineIamBinding)(nil)

// IapWebTypeAppEngineIamBinding represents the Terraform resource google_iap_web_type_app_engine_iam_binding.
type IapWebTypeAppEngineIamBinding struct {
	Name      string
	Args      IapWebTypeAppEngineIamBindingArgs
	state     *iapWebTypeAppEngineIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) Type() string {
	return "google_iap_web_type_app_engine_iam_binding"
}

// LocalName returns the local name for [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) LocalName() string {
	return iwtaeib.Name
}

// Configuration returns the configuration (args) for [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) Configuration() interface{} {
	return iwtaeib.Args
}

// DependOn is used for other resources to depend on [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iwtaeib)
}

// Dependencies returns the list of resources [IapWebTypeAppEngineIamBinding] depends_on.
func (iwtaeib *IapWebTypeAppEngineIamBinding) Dependencies() terra.Dependencies {
	return iwtaeib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iwtaeib.Lifecycle
}

// Attributes returns the attributes for [IapWebTypeAppEngineIamBinding].
func (iwtaeib *IapWebTypeAppEngineIamBinding) Attributes() iapWebTypeAppEngineIamBindingAttributes {
	return iapWebTypeAppEngineIamBindingAttributes{ref: terra.ReferenceResource(iwtaeib)}
}

// ImportState imports the given attribute values into [IapWebTypeAppEngineIamBinding]'s state.
func (iwtaeib *IapWebTypeAppEngineIamBinding) ImportState(av io.Reader) error {
	iwtaeib.state = &iapWebTypeAppEngineIamBindingState{}
	if err := json.NewDecoder(av).Decode(iwtaeib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwtaeib.Type(), iwtaeib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebTypeAppEngineIamBinding] has state.
func (iwtaeib *IapWebTypeAppEngineIamBinding) State() (*iapWebTypeAppEngineIamBindingState, bool) {
	return iwtaeib.state, iwtaeib.state != nil
}

// StateMust returns the state for [IapWebTypeAppEngineIamBinding]. Panics if the state is nil.
func (iwtaeib *IapWebTypeAppEngineIamBinding) StateMust() *iapWebTypeAppEngineIamBindingState {
	if iwtaeib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwtaeib.Type(), iwtaeib.LocalName()))
	}
	return iwtaeib.state
}

// IapWebTypeAppEngineIamBindingArgs contains the configurations for google_iap_web_type_app_engine_iam_binding.
type IapWebTypeAppEngineIamBindingArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebtypeappengineiambinding.Condition `hcl:"condition,block"`
}
type iapWebTypeAppEngineIamBindingAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iwtaeib.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtaeib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtaeib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iwtaeib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtaeib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_type_app_engine_iam_binding.
func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwtaeib.ref.Append("role"))
}

func (iwtaeib iapWebTypeAppEngineIamBindingAttributes) Condition() terra.ListValue[iapwebtypeappengineiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebtypeappengineiambinding.ConditionAttributes](iwtaeib.ref.Append("condition"))
}

type iapWebTypeAppEngineIamBindingState struct {
	AppId     string                                         `json:"app_id"`
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Members   []string                                       `json:"members"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Condition []iapwebtypeappengineiambinding.ConditionState `json:"condition"`
}
