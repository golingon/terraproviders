// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlmanagedinstancefailovergroup "github.com/golingon/terraproviders/azurerm/3.66.0/sqlmanagedinstancefailovergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlManagedInstanceFailoverGroup creates a new instance of [SqlManagedInstanceFailoverGroup].
func NewSqlManagedInstanceFailoverGroup(name string, args SqlManagedInstanceFailoverGroupArgs) *SqlManagedInstanceFailoverGroup {
	return &SqlManagedInstanceFailoverGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlManagedInstanceFailoverGroup)(nil)

// SqlManagedInstanceFailoverGroup represents the Terraform resource azurerm_sql_managed_instance_failover_group.
type SqlManagedInstanceFailoverGroup struct {
	Name      string
	Args      SqlManagedInstanceFailoverGroupArgs
	state     *sqlManagedInstanceFailoverGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) Type() string {
	return "azurerm_sql_managed_instance_failover_group"
}

// LocalName returns the local name for [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) LocalName() string {
	return smifg.Name
}

// Configuration returns the configuration (args) for [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) Configuration() interface{} {
	return smifg.Args
}

// DependOn is used for other resources to depend on [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(smifg)
}

// Dependencies returns the list of resources [SqlManagedInstanceFailoverGroup] depends_on.
func (smifg *SqlManagedInstanceFailoverGroup) Dependencies() terra.Dependencies {
	return smifg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) LifecycleManagement() *terra.Lifecycle {
	return smifg.Lifecycle
}

// Attributes returns the attributes for [SqlManagedInstanceFailoverGroup].
func (smifg *SqlManagedInstanceFailoverGroup) Attributes() sqlManagedInstanceFailoverGroupAttributes {
	return sqlManagedInstanceFailoverGroupAttributes{ref: terra.ReferenceResource(smifg)}
}

// ImportState imports the given attribute values into [SqlManagedInstanceFailoverGroup]'s state.
func (smifg *SqlManagedInstanceFailoverGroup) ImportState(av io.Reader) error {
	smifg.state = &sqlManagedInstanceFailoverGroupState{}
	if err := json.NewDecoder(av).Decode(smifg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smifg.Type(), smifg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlManagedInstanceFailoverGroup] has state.
func (smifg *SqlManagedInstanceFailoverGroup) State() (*sqlManagedInstanceFailoverGroupState, bool) {
	return smifg.state, smifg.state != nil
}

// StateMust returns the state for [SqlManagedInstanceFailoverGroup]. Panics if the state is nil.
func (smifg *SqlManagedInstanceFailoverGroup) StateMust() *sqlManagedInstanceFailoverGroupState {
	if smifg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smifg.Type(), smifg.LocalName()))
	}
	return smifg.state
}

// SqlManagedInstanceFailoverGroupArgs contains the configurations for azurerm_sql_managed_instance_failover_group.
type SqlManagedInstanceFailoverGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedInstanceName: string, required
	ManagedInstanceName terra.StringValue `hcl:"managed_instance_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartnerManagedInstanceId: string, required
	PartnerManagedInstanceId terra.StringValue `hcl:"partner_managed_instance_id,attr" validate:"required"`
	// ReadonlyEndpointFailoverPolicyEnabled: bool, optional
	ReadonlyEndpointFailoverPolicyEnabled terra.BoolValue `hcl:"readonly_endpoint_failover_policy_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// PartnerRegion: min=0
	PartnerRegion []sqlmanagedinstancefailovergroup.PartnerRegion `hcl:"partner_region,block" validate:"min=0"`
	// ReadWriteEndpointFailoverPolicy: required
	ReadWriteEndpointFailoverPolicy *sqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicy `hcl:"read_write_endpoint_failover_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *sqlmanagedinstancefailovergroup.Timeouts `hcl:"timeouts,block"`
}
type sqlManagedInstanceFailoverGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("location"))
}

// ManagedInstanceName returns a reference to field managed_instance_name of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) ManagedInstanceName() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("managed_instance_name"))
}

// Name returns a reference to field name of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("name"))
}

// PartnerManagedInstanceId returns a reference to field partner_managed_instance_id of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) PartnerManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("partner_managed_instance_id"))
}

// ReadonlyEndpointFailoverPolicyEnabled returns a reference to field readonly_endpoint_failover_policy_enabled of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) ReadonlyEndpointFailoverPolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(smifg.ref.Append("readonly_endpoint_failover_policy_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("resource_group_name"))
}

// Role returns a reference to field role of azurerm_sql_managed_instance_failover_group.
func (smifg sqlManagedInstanceFailoverGroupAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(smifg.ref.Append("role"))
}

func (smifg sqlManagedInstanceFailoverGroupAttributes) PartnerRegion() terra.ListValue[sqlmanagedinstancefailovergroup.PartnerRegionAttributes] {
	return terra.ReferenceAsList[sqlmanagedinstancefailovergroup.PartnerRegionAttributes](smifg.ref.Append("partner_region"))
}

func (smifg sqlManagedInstanceFailoverGroupAttributes) ReadWriteEndpointFailoverPolicy() terra.ListValue[sqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[sqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyAttributes](smifg.ref.Append("read_write_endpoint_failover_policy"))
}

func (smifg sqlManagedInstanceFailoverGroupAttributes) Timeouts() sqlmanagedinstancefailovergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlmanagedinstancefailovergroup.TimeoutsAttributes](smifg.ref.Append("timeouts"))
}

type sqlManagedInstanceFailoverGroupState struct {
	Id                                    string                                                                 `json:"id"`
	Location                              string                                                                 `json:"location"`
	ManagedInstanceName                   string                                                                 `json:"managed_instance_name"`
	Name                                  string                                                                 `json:"name"`
	PartnerManagedInstanceId              string                                                                 `json:"partner_managed_instance_id"`
	ReadonlyEndpointFailoverPolicyEnabled bool                                                                   `json:"readonly_endpoint_failover_policy_enabled"`
	ResourceGroupName                     string                                                                 `json:"resource_group_name"`
	Role                                  string                                                                 `json:"role"`
	PartnerRegion                         []sqlmanagedinstancefailovergroup.PartnerRegionState                   `json:"partner_region"`
	ReadWriteEndpointFailoverPolicy       []sqlmanagedinstancefailovergroup.ReadWriteEndpointFailoverPolicyState `json:"read_write_endpoint_failover_policy"`
	Timeouts                              *sqlmanagedinstancefailovergroup.TimeoutsState                         `json:"timeouts"`
}
