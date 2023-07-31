// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsresponsepolicyrule "github.com/golingon/terraproviders/google/4.75.1/dnsresponsepolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsResponsePolicyRule creates a new instance of [DnsResponsePolicyRule].
func NewDnsResponsePolicyRule(name string, args DnsResponsePolicyRuleArgs) *DnsResponsePolicyRule {
	return &DnsResponsePolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsResponsePolicyRule)(nil)

// DnsResponsePolicyRule represents the Terraform resource google_dns_response_policy_rule.
type DnsResponsePolicyRule struct {
	Name      string
	Args      DnsResponsePolicyRuleArgs
	state     *dnsResponsePolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) Type() string {
	return "google_dns_response_policy_rule"
}

// LocalName returns the local name for [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) LocalName() string {
	return drpr.Name
}

// Configuration returns the configuration (args) for [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) Configuration() interface{} {
	return drpr.Args
}

// DependOn is used for other resources to depend on [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(drpr)
}

// Dependencies returns the list of resources [DnsResponsePolicyRule] depends_on.
func (drpr *DnsResponsePolicyRule) Dependencies() terra.Dependencies {
	return drpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) LifecycleManagement() *terra.Lifecycle {
	return drpr.Lifecycle
}

// Attributes returns the attributes for [DnsResponsePolicyRule].
func (drpr *DnsResponsePolicyRule) Attributes() dnsResponsePolicyRuleAttributes {
	return dnsResponsePolicyRuleAttributes{ref: terra.ReferenceResource(drpr)}
}

// ImportState imports the given attribute values into [DnsResponsePolicyRule]'s state.
func (drpr *DnsResponsePolicyRule) ImportState(av io.Reader) error {
	drpr.state = &dnsResponsePolicyRuleState{}
	if err := json.NewDecoder(av).Decode(drpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", drpr.Type(), drpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsResponsePolicyRule] has state.
func (drpr *DnsResponsePolicyRule) State() (*dnsResponsePolicyRuleState, bool) {
	return drpr.state, drpr.state != nil
}

// StateMust returns the state for [DnsResponsePolicyRule]. Panics if the state is nil.
func (drpr *DnsResponsePolicyRule) StateMust() *dnsResponsePolicyRuleState {
	if drpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", drpr.Type(), drpr.LocalName()))
	}
	return drpr.state
}

// DnsResponsePolicyRuleArgs contains the configurations for google_dns_response_policy_rule.
type DnsResponsePolicyRuleArgs struct {
	// DnsName: string, required
	DnsName terra.StringValue `hcl:"dns_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResponsePolicy: string, required
	ResponsePolicy terra.StringValue `hcl:"response_policy,attr" validate:"required"`
	// RuleName: string, required
	RuleName terra.StringValue `hcl:"rule_name,attr" validate:"required"`
	// LocalData: optional
	LocalData *dnsresponsepolicyrule.LocalData `hcl:"local_data,block"`
	// Timeouts: optional
	Timeouts *dnsresponsepolicyrule.Timeouts `hcl:"timeouts,block"`
}
type dnsResponsePolicyRuleAttributes struct {
	ref terra.Reference
}

// DnsName returns a reference to field dns_name of google_dns_response_policy_rule.
func (drpr dnsResponsePolicyRuleAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(drpr.ref.Append("dns_name"))
}

// Id returns a reference to field id of google_dns_response_policy_rule.
func (drpr dnsResponsePolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drpr.ref.Append("id"))
}

// Project returns a reference to field project of google_dns_response_policy_rule.
func (drpr dnsResponsePolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(drpr.ref.Append("project"))
}

// ResponsePolicy returns a reference to field response_policy of google_dns_response_policy_rule.
func (drpr dnsResponsePolicyRuleAttributes) ResponsePolicy() terra.StringValue {
	return terra.ReferenceAsString(drpr.ref.Append("response_policy"))
}

// RuleName returns a reference to field rule_name of google_dns_response_policy_rule.
func (drpr dnsResponsePolicyRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(drpr.ref.Append("rule_name"))
}

func (drpr dnsResponsePolicyRuleAttributes) LocalData() terra.ListValue[dnsresponsepolicyrule.LocalDataAttributes] {
	return terra.ReferenceAsList[dnsresponsepolicyrule.LocalDataAttributes](drpr.ref.Append("local_data"))
}

func (drpr dnsResponsePolicyRuleAttributes) Timeouts() dnsresponsepolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsresponsepolicyrule.TimeoutsAttributes](drpr.ref.Append("timeouts"))
}

type dnsResponsePolicyRuleState struct {
	DnsName        string                                 `json:"dns_name"`
	Id             string                                 `json:"id"`
	Project        string                                 `json:"project"`
	ResponsePolicy string                                 `json:"response_policy"`
	RuleName       string                                 `json:"rule_name"`
	LocalData      []dnsresponsepolicyrule.LocalDataState `json:"local_data"`
	Timeouts       *dnsresponsepolicyrule.TimeoutsState   `json:"timeouts"`
}
