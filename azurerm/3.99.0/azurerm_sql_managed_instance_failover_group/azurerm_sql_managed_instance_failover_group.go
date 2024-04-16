// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sql_managed_instance_failover_group

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

// Resource represents the Terraform resource azurerm_sql_managed_instance_failover_group.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSqlManagedInstanceFailoverGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asmifg *Resource) Type() string {
	return "azurerm_sql_managed_instance_failover_group"
}

// LocalName returns the local name for [Resource].
func (asmifg *Resource) LocalName() string {
	return asmifg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asmifg *Resource) Configuration() interface{} {
	return asmifg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asmifg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asmifg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asmifg *Resource) Dependencies() terra.Dependencies {
	return asmifg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asmifg *Resource) LifecycleManagement() *terra.Lifecycle {
	return asmifg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asmifg *Resource) Attributes() azurermSqlManagedInstanceFailoverGroupAttributes {
	return azurermSqlManagedInstanceFailoverGroupAttributes{ref: terra.ReferenceResource(asmifg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asmifg *Resource) ImportState(state io.Reader) error {
	asmifg.state = &azurermSqlManagedInstanceFailoverGroupState{}
	if err := json.NewDecoder(state).Decode(asmifg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asmifg.Type(), asmifg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asmifg *Resource) State() (*azurermSqlManagedInstanceFailoverGroupState, bool) {
	return asmifg.state, asmifg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asmifg *Resource) StateMust() *azurermSqlManagedInstanceFailoverGroupState {
	if asmifg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asmifg.Type(), asmifg.LocalName()))
	}
	return asmifg.state
}

// Args contains the configurations for azurerm_sql_managed_instance_failover_group.
type Args struct {
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
	// ReadWriteEndpointFailoverPolicy: required
	ReadWriteEndpointFailoverPolicy *ReadWriteEndpointFailoverPolicy `hcl:"read_write_endpoint_failover_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSqlManagedInstanceFailoverGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("location"))
}

// ManagedInstanceName returns a reference to field managed_instance_name of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) ManagedInstanceName() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("managed_instance_name"))
}

// Name returns a reference to field name of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("name"))
}

// PartnerManagedInstanceId returns a reference to field partner_managed_instance_id of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) PartnerManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("partner_managed_instance_id"))
}

// ReadonlyEndpointFailoverPolicyEnabled returns a reference to field readonly_endpoint_failover_policy_enabled of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) ReadonlyEndpointFailoverPolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asmifg.ref.Append("readonly_endpoint_failover_policy_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("resource_group_name"))
}

// Role returns a reference to field role of azurerm_sql_managed_instance_failover_group.
func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(asmifg.ref.Append("role"))
}

func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) PartnerRegion() terra.ListValue[PartnerRegionAttributes] {
	return terra.ReferenceAsList[PartnerRegionAttributes](asmifg.ref.Append("partner_region"))
}

func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) ReadWriteEndpointFailoverPolicy() terra.ListValue[ReadWriteEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[ReadWriteEndpointFailoverPolicyAttributes](asmifg.ref.Append("read_write_endpoint_failover_policy"))
}

func (asmifg azurermSqlManagedInstanceFailoverGroupAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asmifg.ref.Append("timeouts"))
}

type azurermSqlManagedInstanceFailoverGroupState struct {
	Id                                    string                                 `json:"id"`
	Location                              string                                 `json:"location"`
	ManagedInstanceName                   string                                 `json:"managed_instance_name"`
	Name                                  string                                 `json:"name"`
	PartnerManagedInstanceId              string                                 `json:"partner_managed_instance_id"`
	ReadonlyEndpointFailoverPolicyEnabled bool                                   `json:"readonly_endpoint_failover_policy_enabled"`
	ResourceGroupName                     string                                 `json:"resource_group_name"`
	Role                                  string                                 `json:"role"`
	PartnerRegion                         []PartnerRegionState                   `json:"partner_region"`
	ReadWriteEndpointFailoverPolicy       []ReadWriteEndpointFailoverPolicyState `json:"read_write_endpoint_failover_policy"`
	Timeouts                              *TimeoutsState                         `json:"timeouts"`
}
