// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIapTunnelInstanceIamPolicy creates a new instance of [IapTunnelInstanceIamPolicy].
func NewIapTunnelInstanceIamPolicy(name string, args IapTunnelInstanceIamPolicyArgs) *IapTunnelInstanceIamPolicy {
	return &IapTunnelInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelInstanceIamPolicy)(nil)

// IapTunnelInstanceIamPolicy represents the Terraform resource google_iap_tunnel_instance_iam_policy.
type IapTunnelInstanceIamPolicy struct {
	Name      string
	Args      IapTunnelInstanceIamPolicyArgs
	state     *iapTunnelInstanceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) Type() string {
	return "google_iap_tunnel_instance_iam_policy"
}

// LocalName returns the local name for [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) LocalName() string {
	return itiip.Name
}

// Configuration returns the configuration (args) for [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) Configuration() interface{} {
	return itiip.Args
}

// DependOn is used for other resources to depend on [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(itiip)
}

// Dependencies returns the list of resources [IapTunnelInstanceIamPolicy] depends_on.
func (itiip *IapTunnelInstanceIamPolicy) Dependencies() terra.Dependencies {
	return itiip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return itiip.Lifecycle
}

// Attributes returns the attributes for [IapTunnelInstanceIamPolicy].
func (itiip *IapTunnelInstanceIamPolicy) Attributes() iapTunnelInstanceIamPolicyAttributes {
	return iapTunnelInstanceIamPolicyAttributes{ref: terra.ReferenceResource(itiip)}
}

// ImportState imports the given attribute values into [IapTunnelInstanceIamPolicy]'s state.
func (itiip *IapTunnelInstanceIamPolicy) ImportState(av io.Reader) error {
	itiip.state = &iapTunnelInstanceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(itiip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itiip.Type(), itiip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelInstanceIamPolicy] has state.
func (itiip *IapTunnelInstanceIamPolicy) State() (*iapTunnelInstanceIamPolicyState, bool) {
	return itiip.state, itiip.state != nil
}

// StateMust returns the state for [IapTunnelInstanceIamPolicy]. Panics if the state is nil.
func (itiip *IapTunnelInstanceIamPolicy) StateMust() *iapTunnelInstanceIamPolicyState {
	if itiip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itiip.Type(), itiip.LocalName()))
	}
	return itiip.state
}

// IapTunnelInstanceIamPolicyArgs contains the configurations for google_iap_tunnel_instance_iam_policy.
type IapTunnelInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}
type iapTunnelInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("project"))
}

// Zone returns a reference to field zone of google_iap_tunnel_instance_iam_policy.
func (itiip iapTunnelInstanceIamPolicyAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(itiip.ref.Append("zone"))
}

type iapTunnelInstanceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Instance   string `json:"instance"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Zone       string `json:"zone"`
}
