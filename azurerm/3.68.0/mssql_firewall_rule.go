// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlfirewallrule "github.com/golingon/terraproviders/azurerm/3.68.0/mssqlfirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlFirewallRule creates a new instance of [MssqlFirewallRule].
func NewMssqlFirewallRule(name string, args MssqlFirewallRuleArgs) *MssqlFirewallRule {
	return &MssqlFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlFirewallRule)(nil)

// MssqlFirewallRule represents the Terraform resource azurerm_mssql_firewall_rule.
type MssqlFirewallRule struct {
	Name      string
	Args      MssqlFirewallRuleArgs
	state     *mssqlFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) Type() string {
	return "azurerm_mssql_firewall_rule"
}

// LocalName returns the local name for [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) LocalName() string {
	return mfr.Name
}

// Configuration returns the configuration (args) for [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) Configuration() interface{} {
	return mfr.Args
}

// DependOn is used for other resources to depend on [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(mfr)
}

// Dependencies returns the list of resources [MssqlFirewallRule] depends_on.
func (mfr *MssqlFirewallRule) Dependencies() terra.Dependencies {
	return mfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return mfr.Lifecycle
}

// Attributes returns the attributes for [MssqlFirewallRule].
func (mfr *MssqlFirewallRule) Attributes() mssqlFirewallRuleAttributes {
	return mssqlFirewallRuleAttributes{ref: terra.ReferenceResource(mfr)}
}

// ImportState imports the given attribute values into [MssqlFirewallRule]'s state.
func (mfr *MssqlFirewallRule) ImportState(av io.Reader) error {
	mfr.state = &mssqlFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(mfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfr.Type(), mfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlFirewallRule] has state.
func (mfr *MssqlFirewallRule) State() (*mssqlFirewallRuleState, bool) {
	return mfr.state, mfr.state != nil
}

// StateMust returns the state for [MssqlFirewallRule]. Panics if the state is nil.
func (mfr *MssqlFirewallRule) StateMust() *mssqlFirewallRuleState {
	if mfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfr.Type(), mfr.LocalName()))
	}
	return mfr.state
}

// MssqlFirewallRuleArgs contains the configurations for azurerm_mssql_firewall_rule.
type MssqlFirewallRuleArgs struct {
	// EndIpAddress: string, required
	EndIpAddress terra.StringValue `hcl:"end_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// StartIpAddress: string, required
	StartIpAddress terra.StringValue `hcl:"start_ip_address,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlfirewallrule.Timeouts `hcl:"timeouts,block"`
}
type mssqlFirewallRuleAttributes struct {
	ref terra.Reference
}

// EndIpAddress returns a reference to field end_ip_address of azurerm_mssql_firewall_rule.
func (mfr mssqlFirewallRuleAttributes) EndIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("end_ip_address"))
}

// Id returns a reference to field id of azurerm_mssql_firewall_rule.
func (mfr mssqlFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mssql_firewall_rule.
func (mfr mssqlFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("name"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_firewall_rule.
func (mfr mssqlFirewallRuleAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("server_id"))
}

// StartIpAddress returns a reference to field start_ip_address of azurerm_mssql_firewall_rule.
func (mfr mssqlFirewallRuleAttributes) StartIpAddress() terra.StringValue {
	return terra.ReferenceAsString(mfr.ref.Append("start_ip_address"))
}

func (mfr mssqlFirewallRuleAttributes) Timeouts() mssqlfirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlfirewallrule.TimeoutsAttributes](mfr.ref.Append("timeouts"))
}

type mssqlFirewallRuleState struct {
	EndIpAddress   string                           `json:"end_ip_address"`
	Id             string                           `json:"id"`
	Name           string                           `json:"name"`
	ServerId       string                           `json:"server_id"`
	StartIpAddress string                           `json:"start_ip_address"`
	Timeouts       *mssqlfirewallrule.TimeoutsState `json:"timeouts"`
}
