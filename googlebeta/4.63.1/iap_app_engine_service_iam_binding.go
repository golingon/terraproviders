// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iapappengineserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/iapappengineserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineServiceIamBinding creates a new instance of [IapAppEngineServiceIamBinding].
func NewIapAppEngineServiceIamBinding(name string, args IapAppEngineServiceIamBindingArgs) *IapAppEngineServiceIamBinding {
	return &IapAppEngineServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineServiceIamBinding)(nil)

// IapAppEngineServiceIamBinding represents the Terraform resource google_iap_app_engine_service_iam_binding.
type IapAppEngineServiceIamBinding struct {
	Name      string
	Args      IapAppEngineServiceIamBindingArgs
	state     *iapAppEngineServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) Type() string {
	return "google_iap_app_engine_service_iam_binding"
}

// LocalName returns the local name for [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) LocalName() string {
	return iaesib.Name
}

// Configuration returns the configuration (args) for [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) Configuration() interface{} {
	return iaesib.Args
}

// DependOn is used for other resources to depend on [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iaesib)
}

// Dependencies returns the list of resources [IapAppEngineServiceIamBinding] depends_on.
func (iaesib *IapAppEngineServiceIamBinding) Dependencies() terra.Dependencies {
	return iaesib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iaesib.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineServiceIamBinding].
func (iaesib *IapAppEngineServiceIamBinding) Attributes() iapAppEngineServiceIamBindingAttributes {
	return iapAppEngineServiceIamBindingAttributes{ref: terra.ReferenceResource(iaesib)}
}

// ImportState imports the given attribute values into [IapAppEngineServiceIamBinding]'s state.
func (iaesib *IapAppEngineServiceIamBinding) ImportState(av io.Reader) error {
	iaesib.state = &iapAppEngineServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(iaesib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaesib.Type(), iaesib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineServiceIamBinding] has state.
func (iaesib *IapAppEngineServiceIamBinding) State() (*iapAppEngineServiceIamBindingState, bool) {
	return iaesib.state, iaesib.state != nil
}

// StateMust returns the state for [IapAppEngineServiceIamBinding]. Panics if the state is nil.
func (iaesib *IapAppEngineServiceIamBinding) StateMust() *iapAppEngineServiceIamBindingState {
	if iaesib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaesib.Type(), iaesib.LocalName()))
	}
	return iaesib.state
}

// IapAppEngineServiceIamBindingArgs contains the configurations for google_iap_app_engine_service_iam_binding.
type IapAppEngineServiceIamBindingArgs struct {
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
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Condition: optional
	Condition *iapappengineserviceiambinding.Condition `hcl:"condition,block"`
}
type iapAppEngineServiceIamBindingAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iaesib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("role"))
}

// Service returns a reference to field service of google_iap_app_engine_service_iam_binding.
func (iaesib iapAppEngineServiceIamBindingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaesib.ref.Append("service"))
}

func (iaesib iapAppEngineServiceIamBindingAttributes) Condition() terra.ListValue[iapappengineserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapappengineserviceiambinding.ConditionAttributes](iaesib.ref.Append("condition"))
}

type iapAppEngineServiceIamBindingState struct {
	AppId     string                                         `json:"app_id"`
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Members   []string                                       `json:"members"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Service   string                                         `json:"service"`
	Condition []iapappengineserviceiambinding.ConditionState `json:"condition"`
}
