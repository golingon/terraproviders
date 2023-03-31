// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkfirewallfirewallpolicy "github.com/golingon/terraproviders/aws/4.60.0/networkfirewallfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkfirewallFirewallPolicy creates a new instance of [NetworkfirewallFirewallPolicy].
func NewNetworkfirewallFirewallPolicy(name string, args NetworkfirewallFirewallPolicyArgs) *NetworkfirewallFirewallPolicy {
	return &NetworkfirewallFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkfirewallFirewallPolicy)(nil)

// NetworkfirewallFirewallPolicy represents the Terraform resource aws_networkfirewall_firewall_policy.
type NetworkfirewallFirewallPolicy struct {
	Name      string
	Args      NetworkfirewallFirewallPolicyArgs
	state     *networkfirewallFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) Type() string {
	return "aws_networkfirewall_firewall_policy"
}

// LocalName returns the local name for [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) LocalName() string {
	return nfp.Name
}

// Configuration returns the configuration (args) for [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) Configuration() interface{} {
	return nfp.Args
}

// DependOn is used for other resources to depend on [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(nfp)
}

// Dependencies returns the list of resources [NetworkfirewallFirewallPolicy] depends_on.
func (nfp *NetworkfirewallFirewallPolicy) Dependencies() terra.Dependencies {
	return nfp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return nfp.Lifecycle
}

// Attributes returns the attributes for [NetworkfirewallFirewallPolicy].
func (nfp *NetworkfirewallFirewallPolicy) Attributes() networkfirewallFirewallPolicyAttributes {
	return networkfirewallFirewallPolicyAttributes{ref: terra.ReferenceResource(nfp)}
}

// ImportState imports the given attribute values into [NetworkfirewallFirewallPolicy]'s state.
func (nfp *NetworkfirewallFirewallPolicy) ImportState(av io.Reader) error {
	nfp.state = &networkfirewallFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(nfp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nfp.Type(), nfp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkfirewallFirewallPolicy] has state.
func (nfp *NetworkfirewallFirewallPolicy) State() (*networkfirewallFirewallPolicyState, bool) {
	return nfp.state, nfp.state != nil
}

// StateMust returns the state for [NetworkfirewallFirewallPolicy]. Panics if the state is nil.
func (nfp *NetworkfirewallFirewallPolicy) StateMust() *networkfirewallFirewallPolicyState {
	if nfp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nfp.Type(), nfp.LocalName()))
	}
	return nfp.state
}

// NetworkfirewallFirewallPolicyArgs contains the configurations for aws_networkfirewall_firewall_policy.
type NetworkfirewallFirewallPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EncryptionConfiguration: optional
	EncryptionConfiguration *networkfirewallfirewallpolicy.EncryptionConfiguration `hcl:"encryption_configuration,block"`
	// FirewallPolicy: required
	FirewallPolicy *networkfirewallfirewallpolicy.FirewallPolicy `hcl:"firewall_policy,block" validate:"required"`
}
type networkfirewallFirewallPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("description"))
}

// Id returns a reference to field id of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("id"))
}

// Name returns a reference to field name of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nfp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nfp.ref.Append("tags_all"))
}

// UpdateToken returns a reference to field update_token of aws_networkfirewall_firewall_policy.
func (nfp networkfirewallFirewallPolicyAttributes) UpdateToken() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("update_token"))
}

func (nfp networkfirewallFirewallPolicyAttributes) EncryptionConfiguration() terra.ListValue[networkfirewallfirewallpolicy.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[networkfirewallfirewallpolicy.EncryptionConfigurationAttributes](nfp.ref.Append("encryption_configuration"))
}

func (nfp networkfirewallFirewallPolicyAttributes) FirewallPolicy() terra.ListValue[networkfirewallfirewallpolicy.FirewallPolicyAttributes] {
	return terra.ReferenceAsList[networkfirewallfirewallpolicy.FirewallPolicyAttributes](nfp.ref.Append("firewall_policy"))
}

type networkfirewallFirewallPolicyState struct {
	Arn                     string                                                       `json:"arn"`
	Description             string                                                       `json:"description"`
	Id                      string                                                       `json:"id"`
	Name                    string                                                       `json:"name"`
	Tags                    map[string]string                                            `json:"tags"`
	TagsAll                 map[string]string                                            `json:"tags_all"`
	UpdateToken             string                                                       `json:"update_token"`
	EncryptionConfiguration []networkfirewallfirewallpolicy.EncryptionConfigurationState `json:"encryption_configuration"`
	FirewallPolicy          []networkfirewallfirewallpolicy.FirewallPolicyState          `json:"firewall_policy"`
}
