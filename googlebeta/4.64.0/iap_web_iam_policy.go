// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebIamPolicy creates a new instance of [IapWebIamPolicy].
func NewIapWebIamPolicy(name string, args IapWebIamPolicyArgs) *IapWebIamPolicy {
	return &IapWebIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebIamPolicy)(nil)

// IapWebIamPolicy represents the Terraform resource google_iap_web_iam_policy.
type IapWebIamPolicy struct {
	Name      string
	Args      IapWebIamPolicyArgs
	state     *iapWebIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) Type() string {
	return "google_iap_web_iam_policy"
}

// LocalName returns the local name for [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) LocalName() string {
	return iwip.Name
}

// Configuration returns the configuration (args) for [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) Configuration() interface{} {
	return iwip.Args
}

// DependOn is used for other resources to depend on [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iwip)
}

// Dependencies returns the list of resources [IapWebIamPolicy] depends_on.
func (iwip *IapWebIamPolicy) Dependencies() terra.Dependencies {
	return iwip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return iwip.Lifecycle
}

// Attributes returns the attributes for [IapWebIamPolicy].
func (iwip *IapWebIamPolicy) Attributes() iapWebIamPolicyAttributes {
	return iapWebIamPolicyAttributes{ref: terra.ReferenceResource(iwip)}
}

// ImportState imports the given attribute values into [IapWebIamPolicy]'s state.
func (iwip *IapWebIamPolicy) ImportState(av io.Reader) error {
	iwip.state = &iapWebIamPolicyState{}
	if err := json.NewDecoder(av).Decode(iwip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwip.Type(), iwip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebIamPolicy] has state.
func (iwip *IapWebIamPolicy) State() (*iapWebIamPolicyState, bool) {
	return iwip.state, iwip.state != nil
}

// StateMust returns the state for [IapWebIamPolicy]. Panics if the state is nil.
func (iwip *IapWebIamPolicy) StateMust() *iapWebIamPolicyState {
	if iwip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwip.Type(), iwip.LocalName()))
	}
	return iwip.state
}

// IapWebIamPolicyArgs contains the configurations for google_iap_web_iam_policy.
type IapWebIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type iapWebIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_iam_policy.
func (iwip iapWebIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_iam_policy.
func (iwip iapWebIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_iam_policy.
func (iwip iapWebIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_iam_policy.
func (iwip iapWebIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("project"))
}

type iapWebIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
