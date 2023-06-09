// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapTunnelIamPolicy creates a new instance of [IapTunnelIamPolicy].
func NewIapTunnelIamPolicy(name string, args IapTunnelIamPolicyArgs) *IapTunnelIamPolicy {
	return &IapTunnelIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelIamPolicy)(nil)

// IapTunnelIamPolicy represents the Terraform resource google_iap_tunnel_iam_policy.
type IapTunnelIamPolicy struct {
	Name      string
	Args      IapTunnelIamPolicyArgs
	state     *iapTunnelIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) Type() string {
	return "google_iap_tunnel_iam_policy"
}

// LocalName returns the local name for [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) LocalName() string {
	return itip.Name
}

// Configuration returns the configuration (args) for [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) Configuration() interface{} {
	return itip.Args
}

// DependOn is used for other resources to depend on [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(itip)
}

// Dependencies returns the list of resources [IapTunnelIamPolicy] depends_on.
func (itip *IapTunnelIamPolicy) Dependencies() terra.Dependencies {
	return itip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return itip.Lifecycle
}

// Attributes returns the attributes for [IapTunnelIamPolicy].
func (itip *IapTunnelIamPolicy) Attributes() iapTunnelIamPolicyAttributes {
	return iapTunnelIamPolicyAttributes{ref: terra.ReferenceResource(itip)}
}

// ImportState imports the given attribute values into [IapTunnelIamPolicy]'s state.
func (itip *IapTunnelIamPolicy) ImportState(av io.Reader) error {
	itip.state = &iapTunnelIamPolicyState{}
	if err := json.NewDecoder(av).Decode(itip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itip.Type(), itip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelIamPolicy] has state.
func (itip *IapTunnelIamPolicy) State() (*iapTunnelIamPolicyState, bool) {
	return itip.state, itip.state != nil
}

// StateMust returns the state for [IapTunnelIamPolicy]. Panics if the state is nil.
func (itip *IapTunnelIamPolicy) StateMust() *iapTunnelIamPolicyState {
	if itip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itip.Type(), itip.LocalName()))
	}
	return itip.state
}

// IapTunnelIamPolicyArgs contains the configurations for google_iap_tunnel_iam_policy.
type IapTunnelIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type iapTunnelIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_tunnel_iam_policy.
func (itip iapTunnelIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(itip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_tunnel_iam_policy.
func (itip iapTunnelIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_tunnel_iam_policy.
func (itip iapTunnelIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(itip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_tunnel_iam_policy.
func (itip iapTunnelIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itip.ref.Append("project"))
}

type iapTunnelIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
