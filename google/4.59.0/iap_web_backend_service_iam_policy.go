// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebBackendServiceIamPolicy creates a new instance of [IapWebBackendServiceIamPolicy].
func NewIapWebBackendServiceIamPolicy(name string, args IapWebBackendServiceIamPolicyArgs) *IapWebBackendServiceIamPolicy {
	return &IapWebBackendServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebBackendServiceIamPolicy)(nil)

// IapWebBackendServiceIamPolicy represents the Terraform resource google_iap_web_backend_service_iam_policy.
type IapWebBackendServiceIamPolicy struct {
	Name      string
	Args      IapWebBackendServiceIamPolicyArgs
	state     *iapWebBackendServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) Type() string {
	return "google_iap_web_backend_service_iam_policy"
}

// LocalName returns the local name for [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) LocalName() string {
	return iwbsip.Name
}

// Configuration returns the configuration (args) for [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) Configuration() interface{} {
	return iwbsip.Args
}

// DependOn is used for other resources to depend on [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iwbsip)
}

// Dependencies returns the list of resources [IapWebBackendServiceIamPolicy] depends_on.
func (iwbsip *IapWebBackendServiceIamPolicy) Dependencies() terra.Dependencies {
	return iwbsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return iwbsip.Lifecycle
}

// Attributes returns the attributes for [IapWebBackendServiceIamPolicy].
func (iwbsip *IapWebBackendServiceIamPolicy) Attributes() iapWebBackendServiceIamPolicyAttributes {
	return iapWebBackendServiceIamPolicyAttributes{ref: terra.ReferenceResource(iwbsip)}
}

// ImportState imports the given attribute values into [IapWebBackendServiceIamPolicy]'s state.
func (iwbsip *IapWebBackendServiceIamPolicy) ImportState(av io.Reader) error {
	iwbsip.state = &iapWebBackendServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(iwbsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwbsip.Type(), iwbsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebBackendServiceIamPolicy] has state.
func (iwbsip *IapWebBackendServiceIamPolicy) State() (*iapWebBackendServiceIamPolicyState, bool) {
	return iwbsip.state, iwbsip.state != nil
}

// StateMust returns the state for [IapWebBackendServiceIamPolicy]. Panics if the state is nil.
func (iwbsip *IapWebBackendServiceIamPolicy) StateMust() *iapWebBackendServiceIamPolicyState {
	if iwbsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwbsip.Type(), iwbsip.LocalName()))
	}
	return iwbsip.state
}

// IapWebBackendServiceIamPolicyArgs contains the configurations for google_iap_web_backend_service_iam_policy.
type IapWebBackendServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WebBackendService: string, required
	WebBackendService terra.StringValue `hcl:"web_backend_service,attr" validate:"required"`
}
type iapWebBackendServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_backend_service_iam_policy.
func (iwbsip iapWebBackendServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_backend_service_iam_policy.
func (iwbsip iapWebBackendServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_backend_service_iam_policy.
func (iwbsip iapWebBackendServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_backend_service_iam_policy.
func (iwbsip iapWebBackendServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("project"))
}

// WebBackendService returns a reference to field web_backend_service of google_iap_web_backend_service_iam_policy.
func (iwbsip iapWebBackendServiceIamPolicyAttributes) WebBackendService() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("web_backend_service"))
}

type iapWebBackendServiceIamPolicyState struct {
	Etag              string `json:"etag"`
	Id                string `json:"id"`
	PolicyData        string `json:"policy_data"`
	Project           string `json:"project"`
	WebBackendService string `json:"web_backend_service"`
}
