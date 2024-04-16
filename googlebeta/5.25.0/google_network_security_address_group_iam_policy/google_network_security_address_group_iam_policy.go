// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_network_security_address_group_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_network_security_address_group_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNetworkSecurityAddressGroupIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnsagip *Resource) Type() string {
	return "google_network_security_address_group_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gnsagip *Resource) LocalName() string {
	return gnsagip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnsagip *Resource) Configuration() interface{} {
	return gnsagip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnsagip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnsagip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnsagip *Resource) Dependencies() terra.Dependencies {
	return gnsagip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnsagip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnsagip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnsagip *Resource) Attributes() googleNetworkSecurityAddressGroupIamPolicyAttributes {
	return googleNetworkSecurityAddressGroupIamPolicyAttributes{ref: terra.ReferenceResource(gnsagip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnsagip *Resource) ImportState(state io.Reader) error {
	gnsagip.state = &googleNetworkSecurityAddressGroupIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gnsagip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnsagip.Type(), gnsagip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnsagip *Resource) State() (*googleNetworkSecurityAddressGroupIamPolicyState, bool) {
	return gnsagip.state, gnsagip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnsagip *Resource) StateMust() *googleNetworkSecurityAddressGroupIamPolicyState {
	if gnsagip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnsagip.Type(), gnsagip.LocalName()))
	}
	return gnsagip.state
}

// Args contains the configurations for google_network_security_address_group_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type googleNetworkSecurityAddressGroupIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("etag"))
}

// Id returns a reference to field id of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_network_security_address_group_iam_policy.
func (gnsagip googleNetworkSecurityAddressGroupIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnsagip.ref.Append("project"))
}

type googleNetworkSecurityAddressGroupIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
