// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkfirewallrulegroup "github.com/golingon/terraproviders/aws/5.10.0/networkfirewallrulegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkfirewallRuleGroup creates a new instance of [NetworkfirewallRuleGroup].
func NewNetworkfirewallRuleGroup(name string, args NetworkfirewallRuleGroupArgs) *NetworkfirewallRuleGroup {
	return &NetworkfirewallRuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkfirewallRuleGroup)(nil)

// NetworkfirewallRuleGroup represents the Terraform resource aws_networkfirewall_rule_group.
type NetworkfirewallRuleGroup struct {
	Name      string
	Args      NetworkfirewallRuleGroupArgs
	state     *networkfirewallRuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) Type() string {
	return "aws_networkfirewall_rule_group"
}

// LocalName returns the local name for [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) LocalName() string {
	return nrg.Name
}

// Configuration returns the configuration (args) for [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) Configuration() interface{} {
	return nrg.Args
}

// DependOn is used for other resources to depend on [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(nrg)
}

// Dependencies returns the list of resources [NetworkfirewallRuleGroup] depends_on.
func (nrg *NetworkfirewallRuleGroup) Dependencies() terra.Dependencies {
	return nrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) LifecycleManagement() *terra.Lifecycle {
	return nrg.Lifecycle
}

// Attributes returns the attributes for [NetworkfirewallRuleGroup].
func (nrg *NetworkfirewallRuleGroup) Attributes() networkfirewallRuleGroupAttributes {
	return networkfirewallRuleGroupAttributes{ref: terra.ReferenceResource(nrg)}
}

// ImportState imports the given attribute values into [NetworkfirewallRuleGroup]'s state.
func (nrg *NetworkfirewallRuleGroup) ImportState(av io.Reader) error {
	nrg.state = &networkfirewallRuleGroupState{}
	if err := json.NewDecoder(av).Decode(nrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrg.Type(), nrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkfirewallRuleGroup] has state.
func (nrg *NetworkfirewallRuleGroup) State() (*networkfirewallRuleGroupState, bool) {
	return nrg.state, nrg.state != nil
}

// StateMust returns the state for [NetworkfirewallRuleGroup]. Panics if the state is nil.
func (nrg *NetworkfirewallRuleGroup) StateMust() *networkfirewallRuleGroupState {
	if nrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrg.Type(), nrg.LocalName()))
	}
	return nrg.state
}

// NetworkfirewallRuleGroupArgs contains the configurations for aws_networkfirewall_rule_group.
type NetworkfirewallRuleGroupArgs struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Rules: string, optional
	Rules terra.StringValue `hcl:"rules,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// EncryptionConfiguration: optional
	EncryptionConfiguration *networkfirewallrulegroup.EncryptionConfiguration `hcl:"encryption_configuration,block"`
	// RuleGroup: optional
	RuleGroup *networkfirewallrulegroup.RuleGroup `hcl:"rule_group,block"`
}
type networkfirewallRuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("arn"))
}

// Capacity returns a reference to field capacity of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(nrg.ref.Append("capacity"))
}

// Description returns a reference to field description of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("description"))
}

// Id returns a reference to field id of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("id"))
}

// Name returns a reference to field name of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("name"))
}

// Rules returns a reference to field rules of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Rules() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("rules"))
}

// Tags returns a reference to field tags of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nrg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nrg.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("type"))
}

// UpdateToken returns a reference to field update_token of aws_networkfirewall_rule_group.
func (nrg networkfirewallRuleGroupAttributes) UpdateToken() terra.StringValue {
	return terra.ReferenceAsString(nrg.ref.Append("update_token"))
}

func (nrg networkfirewallRuleGroupAttributes) EncryptionConfiguration() terra.ListValue[networkfirewallrulegroup.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[networkfirewallrulegroup.EncryptionConfigurationAttributes](nrg.ref.Append("encryption_configuration"))
}

func (nrg networkfirewallRuleGroupAttributes) RuleGroup() terra.ListValue[networkfirewallrulegroup.RuleGroupAttributes] {
	return terra.ReferenceAsList[networkfirewallrulegroup.RuleGroupAttributes](nrg.ref.Append("rule_group"))
}

type networkfirewallRuleGroupState struct {
	Arn                     string                                                  `json:"arn"`
	Capacity                float64                                                 `json:"capacity"`
	Description             string                                                  `json:"description"`
	Id                      string                                                  `json:"id"`
	Name                    string                                                  `json:"name"`
	Rules                   string                                                  `json:"rules"`
	Tags                    map[string]string                                       `json:"tags"`
	TagsAll                 map[string]string                                       `json:"tags_all"`
	Type                    string                                                  `json:"type"`
	UpdateToken             string                                                  `json:"update_token"`
	EncryptionConfiguration []networkfirewallrulegroup.EncryptionConfigurationState `json:"encryption_configuration"`
	RuleGroup               []networkfirewallrulegroup.RuleGroupState               `json:"rule_group"`
}
