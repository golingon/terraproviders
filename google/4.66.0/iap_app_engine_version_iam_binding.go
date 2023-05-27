// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapappengineversioniambinding "github.com/golingon/terraproviders/google/4.66.0/iapappengineversioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineVersionIamBinding creates a new instance of [IapAppEngineVersionIamBinding].
func NewIapAppEngineVersionIamBinding(name string, args IapAppEngineVersionIamBindingArgs) *IapAppEngineVersionIamBinding {
	return &IapAppEngineVersionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineVersionIamBinding)(nil)

// IapAppEngineVersionIamBinding represents the Terraform resource google_iap_app_engine_version_iam_binding.
type IapAppEngineVersionIamBinding struct {
	Name      string
	Args      IapAppEngineVersionIamBindingArgs
	state     *iapAppEngineVersionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) Type() string {
	return "google_iap_app_engine_version_iam_binding"
}

// LocalName returns the local name for [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) LocalName() string {
	return iaevib.Name
}

// Configuration returns the configuration (args) for [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) Configuration() interface{} {
	return iaevib.Args
}

// DependOn is used for other resources to depend on [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(iaevib)
}

// Dependencies returns the list of resources [IapAppEngineVersionIamBinding] depends_on.
func (iaevib *IapAppEngineVersionIamBinding) Dependencies() terra.Dependencies {
	return iaevib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return iaevib.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineVersionIamBinding].
func (iaevib *IapAppEngineVersionIamBinding) Attributes() iapAppEngineVersionIamBindingAttributes {
	return iapAppEngineVersionIamBindingAttributes{ref: terra.ReferenceResource(iaevib)}
}

// ImportState imports the given attribute values into [IapAppEngineVersionIamBinding]'s state.
func (iaevib *IapAppEngineVersionIamBinding) ImportState(av io.Reader) error {
	iaevib.state = &iapAppEngineVersionIamBindingState{}
	if err := json.NewDecoder(av).Decode(iaevib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaevib.Type(), iaevib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineVersionIamBinding] has state.
func (iaevib *IapAppEngineVersionIamBinding) State() (*iapAppEngineVersionIamBindingState, bool) {
	return iaevib.state, iaevib.state != nil
}

// StateMust returns the state for [IapAppEngineVersionIamBinding]. Panics if the state is nil.
func (iaevib *IapAppEngineVersionIamBinding) StateMust() *iapAppEngineVersionIamBindingState {
	if iaevib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaevib.Type(), iaevib.LocalName()))
	}
	return iaevib.state
}

// IapAppEngineVersionIamBindingArgs contains the configurations for google_iap_app_engine_version_iam_binding.
type IapAppEngineVersionIamBindingArgs struct {
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
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
	// Condition: optional
	Condition *iapappengineversioniambinding.Condition `hcl:"condition,block"`
}
type iapAppEngineVersionIamBindingAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iaevib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("role"))
}

// Service returns a reference to field service of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("service"))
}

// VersionId returns a reference to field version_id of google_iap_app_engine_version_iam_binding.
func (iaevib iapAppEngineVersionIamBindingAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(iaevib.ref.Append("version_id"))
}

func (iaevib iapAppEngineVersionIamBindingAttributes) Condition() terra.ListValue[iapappengineversioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[iapappengineversioniambinding.ConditionAttributes](iaevib.ref.Append("condition"))
}

type iapAppEngineVersionIamBindingState struct {
	AppId     string                                         `json:"app_id"`
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Members   []string                                       `json:"members"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Service   string                                         `json:"service"`
	VersionId string                                         `json:"version_id"`
	Condition []iapappengineversioniambinding.ConditionState `json:"condition"`
}
