// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnspolicy "github.com/golingon/terraproviders/google/4.63.1/dnspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsPolicy creates a new instance of [DnsPolicy].
func NewDnsPolicy(name string, args DnsPolicyArgs) *DnsPolicy {
	return &DnsPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsPolicy)(nil)

// DnsPolicy represents the Terraform resource google_dns_policy.
type DnsPolicy struct {
	Name      string
	Args      DnsPolicyArgs
	state     *dnsPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsPolicy].
func (dp *DnsPolicy) Type() string {
	return "google_dns_policy"
}

// LocalName returns the local name for [DnsPolicy].
func (dp *DnsPolicy) LocalName() string {
	return dp.Name
}

// Configuration returns the configuration (args) for [DnsPolicy].
func (dp *DnsPolicy) Configuration() interface{} {
	return dp.Args
}

// DependOn is used for other resources to depend on [DnsPolicy].
func (dp *DnsPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dp)
}

// Dependencies returns the list of resources [DnsPolicy] depends_on.
func (dp *DnsPolicy) Dependencies() terra.Dependencies {
	return dp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsPolicy].
func (dp *DnsPolicy) LifecycleManagement() *terra.Lifecycle {
	return dp.Lifecycle
}

// Attributes returns the attributes for [DnsPolicy].
func (dp *DnsPolicy) Attributes() dnsPolicyAttributes {
	return dnsPolicyAttributes{ref: terra.ReferenceResource(dp)}
}

// ImportState imports the given attribute values into [DnsPolicy]'s state.
func (dp *DnsPolicy) ImportState(av io.Reader) error {
	dp.state = &dnsPolicyState{}
	if err := json.NewDecoder(av).Decode(dp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dp.Type(), dp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsPolicy] has state.
func (dp *DnsPolicy) State() (*dnsPolicyState, bool) {
	return dp.state, dp.state != nil
}

// StateMust returns the state for [DnsPolicy]. Panics if the state is nil.
func (dp *DnsPolicy) StateMust() *dnsPolicyState {
	if dp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dp.Type(), dp.LocalName()))
	}
	return dp.state
}

// DnsPolicyArgs contains the configurations for google_dns_policy.
type DnsPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableInboundForwarding: bool, optional
	EnableInboundForwarding terra.BoolValue `hcl:"enable_inbound_forwarding,attr"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AlternativeNameServerConfig: optional
	AlternativeNameServerConfig *dnspolicy.AlternativeNameServerConfig `hcl:"alternative_name_server_config,block"`
	// Networks: min=0
	Networks []dnspolicy.Networks `hcl:"networks,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dnspolicy.Timeouts `hcl:"timeouts,block"`
}
type dnsPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dns_policy.
func (dp dnsPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("description"))
}

// EnableInboundForwarding returns a reference to field enable_inbound_forwarding of google_dns_policy.
func (dp dnsPolicyAttributes) EnableInboundForwarding() terra.BoolValue {
	return terra.ReferenceAsBool(dp.ref.Append("enable_inbound_forwarding"))
}

// EnableLogging returns a reference to field enable_logging of google_dns_policy.
func (dp dnsPolicyAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dp.ref.Append("enable_logging"))
}

// Id returns a reference to field id of google_dns_policy.
func (dp dnsPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("id"))
}

// Name returns a reference to field name of google_dns_policy.
func (dp dnsPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

// Project returns a reference to field project of google_dns_policy.
func (dp dnsPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("project"))
}

func (dp dnsPolicyAttributes) AlternativeNameServerConfig() terra.ListValue[dnspolicy.AlternativeNameServerConfigAttributes] {
	return terra.ReferenceAsList[dnspolicy.AlternativeNameServerConfigAttributes](dp.ref.Append("alternative_name_server_config"))
}

func (dp dnsPolicyAttributes) Networks() terra.SetValue[dnspolicy.NetworksAttributes] {
	return terra.ReferenceAsSet[dnspolicy.NetworksAttributes](dp.ref.Append("networks"))
}

func (dp dnsPolicyAttributes) Timeouts() dnspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnspolicy.TimeoutsAttributes](dp.ref.Append("timeouts"))
}

type dnsPolicyState struct {
	Description                 string                                       `json:"description"`
	EnableInboundForwarding     bool                                         `json:"enable_inbound_forwarding"`
	EnableLogging               bool                                         `json:"enable_logging"`
	Id                          string                                       `json:"id"`
	Name                        string                                       `json:"name"`
	Project                     string                                       `json:"project"`
	AlternativeNameServerConfig []dnspolicy.AlternativeNameServerConfigState `json:"alternative_name_server_config"`
	Networks                    []dnspolicy.NetworksState                    `json:"networks"`
	Timeouts                    *dnspolicy.TimeoutsState                     `json:"timeouts"`
}
