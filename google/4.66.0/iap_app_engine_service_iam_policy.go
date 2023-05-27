// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapAppEngineServiceIamPolicy creates a new instance of [IapAppEngineServiceIamPolicy].
func NewIapAppEngineServiceIamPolicy(name string, args IapAppEngineServiceIamPolicyArgs) *IapAppEngineServiceIamPolicy {
	return &IapAppEngineServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapAppEngineServiceIamPolicy)(nil)

// IapAppEngineServiceIamPolicy represents the Terraform resource google_iap_app_engine_service_iam_policy.
type IapAppEngineServiceIamPolicy struct {
	Name      string
	Args      IapAppEngineServiceIamPolicyArgs
	state     *iapAppEngineServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) Type() string {
	return "google_iap_app_engine_service_iam_policy"
}

// LocalName returns the local name for [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) LocalName() string {
	return iaesip.Name
}

// Configuration returns the configuration (args) for [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) Configuration() interface{} {
	return iaesip.Args
}

// DependOn is used for other resources to depend on [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iaesip)
}

// Dependencies returns the list of resources [IapAppEngineServiceIamPolicy] depends_on.
func (iaesip *IapAppEngineServiceIamPolicy) Dependencies() terra.Dependencies {
	return iaesip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return iaesip.Lifecycle
}

// Attributes returns the attributes for [IapAppEngineServiceIamPolicy].
func (iaesip *IapAppEngineServiceIamPolicy) Attributes() iapAppEngineServiceIamPolicyAttributes {
	return iapAppEngineServiceIamPolicyAttributes{ref: terra.ReferenceResource(iaesip)}
}

// ImportState imports the given attribute values into [IapAppEngineServiceIamPolicy]'s state.
func (iaesip *IapAppEngineServiceIamPolicy) ImportState(av io.Reader) error {
	iaesip.state = &iapAppEngineServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(iaesip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaesip.Type(), iaesip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapAppEngineServiceIamPolicy] has state.
func (iaesip *IapAppEngineServiceIamPolicy) State() (*iapAppEngineServiceIamPolicyState, bool) {
	return iaesip.state, iaesip.state != nil
}

// StateMust returns the state for [IapAppEngineServiceIamPolicy]. Panics if the state is nil.
func (iaesip *IapAppEngineServiceIamPolicy) StateMust() *iapAppEngineServiceIamPolicyState {
	if iaesip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaesip.Type(), iaesip.LocalName()))
	}
	return iaesip.state
}

// IapAppEngineServiceIamPolicyArgs contains the configurations for google_iap_app_engine_service_iam_policy.
type IapAppEngineServiceIamPolicyArgs struct {
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
}
type iapAppEngineServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("project"))
}

// Service returns a reference to field service of google_iap_app_engine_service_iam_policy.
func (iaesip iapAppEngineServiceIamPolicyAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(iaesip.ref.Append("service"))
}

type iapAppEngineServiceIamPolicyState struct {
	AppId      string `json:"app_id"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Service    string `json:"service"`
}
