// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlmanagedinstancefailovergroup "github.com/golingon/terraproviders/azurerm/3.64.0/mssqlmanagedinstancefailovergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlManagedInstanceFailoverGroup creates a new instance of [MssqlManagedInstanceFailoverGroup].
func NewMssqlManagedInstanceFailoverGroup(name string, args MssqlManagedInstanceFailoverGroupArgs) *MssqlManagedInstanceFailoverGroup {
	return &MssqlManagedInstanceFailoverGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlManagedInstanceFailoverGroup)(nil)

// MssqlManagedInstanceFailoverGroup represents the Terraform resource azurerm_mssql_managed_instance_failover_group.
type MssqlManagedInstanceFailoverGroup struct {
	Name      string
	Args      MssqlManagedInstanceFailoverGroupArgs
	state     *mssqlManagedInstanceFailoverGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) Type() string {
	return "azurerm_mssql_managed_instance_failover_group"
}

// LocalName returns the local name for [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) LocalName() string {
	return mmifg.Name
}

// Configuration returns the configuration (args) for [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) Configuration() interface{} {
	return mmifg.Args
}

// DependOn is used for other resources to depend on [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mmifg)
}

// Dependencies returns the list of resources [MssqlManagedInstanceFailoverGroup] depends_on.
func (mmifg *MssqlManagedInstanceFailoverGroup) Dependencies() terra.Dependencies {
	return mmifg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) LifecycleManagement() *terra.Lifecycle {
	return mmifg.Lifecycle
}

// Attributes returns the attributes for [MssqlManagedInstanceFailoverGroup].
func (mmifg *MssqlManagedInstanceFailoverGroup) Attributes() mssqlManagedInstanceFailoverGroupAttributes {
	return mssqlManagedInstanceFailoverGroupAttributes{ref: terra.ReferenceResource(mmifg)}
}

// ImportState imports the given attribute values into [MssqlManagedInstanceFailoverGroup]'s state.
func (mmifg *MssqlManagedInstanceFailoverGroup) ImportState(av io.Reader) error {
	mmifg.state = &mssqlManagedInstanceFailoverGroupState{}
	if err := json.NewDecoder(av).Decode(mmifg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmifg.Type(), mmifg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlManagedInstanceFailoverGroup] has state.
func (mmifg *MssqlManagedInstanceFailoverGroup) State() (*mssqlManagedInstanceFailoverGroupState, bool) {
	return mmifg.state, mmifg.state != nil
}

// StateMust returns the state for [MssqlManagedInstanceFailoverGroup]. Panics if the state is nil.
func (mmifg *MssqlManagedInstanceFailoverGroup) StateMust() *mssqlManagedInstanceFailoverGroupState {
	if mmifg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmifg.Type(), mmifg.LocalName()))
	}
	return mmifg.state
}

// MssqlManagedInstanceFailoverGroupArgs contains the configurations for azurerm_mssql_managed_instance_failover_group.
type MssqlManagedInstanceFailoverGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedInstanceId: string, required
	ManagedInstanceId terra.StringValue `hcl:"managed_instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartnerManagedInstanceId: string, required
	PartnerManagedInstanceId terra.StringValue `hcl:"partner_managed_instance_id,attr" validate:"required"`
	// ReadonlyEndpointFailoverPolicyEnabled: bool, optional
	ReadonlyEndpointFailoverPolicyEnabled terra.BoolValue `hcl:"readonly_endpoint_failover_policy_enabled,attr"`
	// PartnerRegion: min=0
	PartnerRegion []mssqlmanagedinstancefailovergroup.PartnerRegion `hcl:"partner_region,block" validate:"min=0"`
	// ReadWriteEndpointFailoverPolicy: required
	ReadWriteEndpointFailoverPolicy *mssqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicy `hcl:"read_write_endpoint_failover_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlmanagedinstancefailovergroup.Timeouts `hcl:"timeouts,block"`
}
type mssqlManagedInstanceFailoverGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("location"))
}

// ManagedInstanceId returns a reference to field managed_instance_id of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) ManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("managed_instance_id"))
}

// Name returns a reference to field name of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("name"))
}

// PartnerManagedInstanceId returns a reference to field partner_managed_instance_id of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) PartnerManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("partner_managed_instance_id"))
}

// ReadonlyEndpointFailoverPolicyEnabled returns a reference to field readonly_endpoint_failover_policy_enabled of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) ReadonlyEndpointFailoverPolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mmifg.ref.Append("readonly_endpoint_failover_policy_enabled"))
}

// Role returns a reference to field role of azurerm_mssql_managed_instance_failover_group.
func (mmifg mssqlManagedInstanceFailoverGroupAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(mmifg.ref.Append("role"))
}

func (mmifg mssqlManagedInstanceFailoverGroupAttributes) PartnerRegion() terra.ListValue[mssqlmanagedinstancefailovergroup.PartnerRegionAttributes] {
	return terra.ReferenceAsList[mssqlmanagedinstancefailovergroup.PartnerRegionAttributes](mmifg.ref.Append("partner_region"))
}

func (mmifg mssqlManagedInstanceFailoverGroupAttributes) ReadWriteEndpointFailoverPolicy() terra.ListValue[mssqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[mssqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyAttributes](mmifg.ref.Append("read_write_endpoint_failover_policy"))
}

func (mmifg mssqlManagedInstanceFailoverGroupAttributes) Timeouts() mssqlmanagedinstancefailovergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlmanagedinstancefailovergroup.TimeoutsAttributes](mmifg.ref.Append("timeouts"))
}

type mssqlManagedInstanceFailoverGroupState struct {
	Id                                    string                                                                   `json:"id"`
	Location                              string                                                                   `json:"location"`
	ManagedInstanceId                     string                                                                   `json:"managed_instance_id"`
	Name                                  string                                                                   `json:"name"`
	PartnerManagedInstanceId              string                                                                   `json:"partner_managed_instance_id"`
	ReadonlyEndpointFailoverPolicyEnabled bool                                                                     `json:"readonly_endpoint_failover_policy_enabled"`
	Role                                  string                                                                   `json:"role"`
	PartnerRegion                         []mssqlmanagedinstancefailovergroup.PartnerRegionState                   `json:"partner_region"`
	ReadWriteEndpointFailoverPolicy       []mssqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyState `json:"read_write_endpoint_failover_policy"`
	Timeouts                              *mssqlmanagedinstancefailovergroup.TimeoutsState                         `json:"timeouts"`
}
