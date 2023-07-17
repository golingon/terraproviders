// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlfailovergroup "github.com/golingon/terraproviders/azurerm/3.65.0/mssqlfailovergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlFailoverGroup creates a new instance of [MssqlFailoverGroup].
func NewMssqlFailoverGroup(name string, args MssqlFailoverGroupArgs) *MssqlFailoverGroup {
	return &MssqlFailoverGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlFailoverGroup)(nil)

// MssqlFailoverGroup represents the Terraform resource azurerm_mssql_failover_group.
type MssqlFailoverGroup struct {
	Name      string
	Args      MssqlFailoverGroupArgs
	state     *mssqlFailoverGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) Type() string {
	return "azurerm_mssql_failover_group"
}

// LocalName returns the local name for [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) LocalName() string {
	return mfg.Name
}

// Configuration returns the configuration (args) for [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) Configuration() interface{} {
	return mfg.Args
}

// DependOn is used for other resources to depend on [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mfg)
}

// Dependencies returns the list of resources [MssqlFailoverGroup] depends_on.
func (mfg *MssqlFailoverGroup) Dependencies() terra.Dependencies {
	return mfg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) LifecycleManagement() *terra.Lifecycle {
	return mfg.Lifecycle
}

// Attributes returns the attributes for [MssqlFailoverGroup].
func (mfg *MssqlFailoverGroup) Attributes() mssqlFailoverGroupAttributes {
	return mssqlFailoverGroupAttributes{ref: terra.ReferenceResource(mfg)}
}

// ImportState imports the given attribute values into [MssqlFailoverGroup]'s state.
func (mfg *MssqlFailoverGroup) ImportState(av io.Reader) error {
	mfg.state = &mssqlFailoverGroupState{}
	if err := json.NewDecoder(av).Decode(mfg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfg.Type(), mfg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlFailoverGroup] has state.
func (mfg *MssqlFailoverGroup) State() (*mssqlFailoverGroupState, bool) {
	return mfg.state, mfg.state != nil
}

// StateMust returns the state for [MssqlFailoverGroup]. Panics if the state is nil.
func (mfg *MssqlFailoverGroup) StateMust() *mssqlFailoverGroupState {
	if mfg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfg.Type(), mfg.LocalName()))
	}
	return mfg.state
}

// MssqlFailoverGroupArgs contains the configurations for azurerm_mssql_failover_group.
type MssqlFailoverGroupArgs struct {
	// Databases: set of string, optional
	Databases terra.SetValue[terra.StringValue] `hcl:"databases,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReadonlyEndpointFailoverPolicyEnabled: bool, optional
	ReadonlyEndpointFailoverPolicyEnabled terra.BoolValue `hcl:"readonly_endpoint_failover_policy_enabled,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PartnerServer: min=1
	PartnerServer []mssqlfailovergroup.PartnerServer `hcl:"partner_server,block" validate:"min=1"`
	// ReadWriteEndpointFailoverPolicy: required
	ReadWriteEndpointFailoverPolicy *mssqlfailovergroup.ReadWriteEndpointFailoverPolicy `hcl:"read_write_endpoint_failover_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlfailovergroup.Timeouts `hcl:"timeouts,block"`
}
type mssqlFailoverGroupAttributes struct {
	ref terra.Reference
}

// Databases returns a reference to field databases of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) Databases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mfg.ref.Append("databases"))
}

// Id returns a reference to field id of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfg.ref.Append("name"))
}

// ReadonlyEndpointFailoverPolicyEnabled returns a reference to field readonly_endpoint_failover_policy_enabled of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) ReadonlyEndpointFailoverPolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mfg.ref.Append("readonly_endpoint_failover_policy_enabled"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mfg.ref.Append("server_id"))
}

// Tags returns a reference to field tags of azurerm_mssql_failover_group.
func (mfg mssqlFailoverGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mfg.ref.Append("tags"))
}

func (mfg mssqlFailoverGroupAttributes) PartnerServer() terra.ListValue[mssqlfailovergroup.PartnerServerAttributes] {
	return terra.ReferenceAsList[mssqlfailovergroup.PartnerServerAttributes](mfg.ref.Append("partner_server"))
}

func (mfg mssqlFailoverGroupAttributes) ReadWriteEndpointFailoverPolicy() terra.ListValue[mssqlfailovergroup.ReadWriteEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[mssqlfailovergroup.ReadWriteEndpointFailoverPolicyAttributes](mfg.ref.Append("read_write_endpoint_failover_policy"))
}

func (mfg mssqlFailoverGroupAttributes) Timeouts() mssqlfailovergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlfailovergroup.TimeoutsAttributes](mfg.ref.Append("timeouts"))
}

type mssqlFailoverGroupState struct {
	Databases                             []string                                                  `json:"databases"`
	Id                                    string                                                    `json:"id"`
	Name                                  string                                                    `json:"name"`
	ReadonlyEndpointFailoverPolicyEnabled bool                                                      `json:"readonly_endpoint_failover_policy_enabled"`
	ServerId                              string                                                    `json:"server_id"`
	Tags                                  map[string]string                                         `json:"tags"`
	PartnerServer                         []mssqlfailovergroup.PartnerServerState                   `json:"partner_server"`
	ReadWriteEndpointFailoverPolicy       []mssqlfailovergroup.ReadWriteEndpointFailoverPolicyState `json:"read_write_endpoint_failover_policy"`
	Timeouts                              *mssqlfailovergroup.TimeoutsState                         `json:"timeouts"`
}
