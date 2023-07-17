// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineVersionIamPolicy creates a new instance of [IapAppEngineVersionIamPolicy].
func NewIapAppEngineVersionIamPolicy(name string, args IapAppEngineVersionIamPolicyArgs) *IapAppEngineVersionIamPolicy {
	return &IapAppEngineVersionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineVersionIamPolicy)(nil)

// IapAppEngineVersionIamPolicy represents the Terraform resource google_iap_app_engine_version_iam_policy.
type IapAppEngineVersionIamPolicy struct {
	Name      string
	Args      IapAppEngineVersionIamPolicyArgs
	state     *iapAppEngineVersionIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) Type() string {
	return "google_iap_app_engine_version_iam_policy"
}

// LocalName returns the local name for [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) LocalName() string {
	return iaevip.Name
}

// Configuration returns the configuration (args) for [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) Configuration() interface{} {
	return iaevip.Args
}

// DependOn is used for other resources to depend on [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iaevip)
}

// Dependencies returns the list of resources [IapAppEngineVersionIamPolicy] depends_on.
func (iaevip *IapAppEngineVersionIamPolicy) Dependencies() terra.Dependencies {
	return iaevip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return iaevip.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineVersionIamPolicy].
func (iaevip *IapAppEngineVersionIamPolicy) Attributes() iapAppEngineVersionIamPolicyAttributes {
	return iapAppEngineVersionIamPolicyAttributes{ref: terra.ReferenceResource(iaevip)}
}

// ImportState imports the given attribute values into [IapAppEngineVersionIamPolicy]'s state.
func (iaevip *IapAppEngineVersionIamPolicy) ImportState(av io.Reader) error {
	iaevip.state = &iapAppEngineVersionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(iaevip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaevip.Type(), iaevip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineVersionIamPolicy] has state.
func (iaevip *IapAppEngineVersionIamPolicy) State() (*iapAppEngineVersionIamPolicyState, bool) {
	return iaevip.state, iaevip.state != nil
}

// StateMust returns the state for [IapAppEngineVersionIamPolicy]. Panics if the state is nil.
func (iaevip *IapAppEngineVersionIamPolicy) StateMust() *iapAppEngineVersionIamPolicyState {
	if iaevip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaevip.Type(), iaevip.LocalName()))
	}
	return iaevip.state
}

// IapAppEngineVersionIamPolicyArgs contains the configurations for google_iap_app_engine_version_iam_policy.
type IapAppEngineVersionIamPolicyArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// VersionId: string, required
	VersionId terra.StringValue `hcl:"version_id,attr" validate:"required"`
}
type iapAppEngineVersionIamPolicyAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("project"))
}

// Service returns a reference to field service of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("service"))
}

// VersionId returns a reference to field version_id of google_iap_app_engine_version_iam_policy.
func (iaevip iapAppEngineVersionIamPolicyAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(iaevip.ref.Append("version_id"))
}

type iapAppEngineVersionIamPolicyState struct {
	AppId      string `json:"app_id"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Service    string `json:"service"`
	VersionId  string `json:"version_id"`
}