// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlfailovergroup "github.com/golingon/terraproviders/azurerm/3.67.0/sqlfailovergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlFailoverGroup creates a new instance of [SqlFailoverGroup].
func NewSqlFailoverGroup(name string, args SqlFailoverGroupArgs) *SqlFailoverGroup {
	return &SqlFailoverGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlFailoverGroup)(nil)

// SqlFailoverGroup represents the Terraform resource azurerm_sql_failover_group.
type SqlFailoverGroup struct {
	Name      string
	Args      SqlFailoverGroupArgs
	state     *sqlFailoverGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) Type() string {
	return "azurerm_sql_failover_group"
}

// LocalName returns the local name for [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) LocalName() string {
	return sfg.Name
}

// Configuration returns the configuration (args) for [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) Configuration() interface{} {
	return sfg.Args
}

// DependOn is used for other resources to depend on [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(sfg)
}

// Dependencies returns the list of resources [SqlFailoverGroup] depends_on.
func (sfg *SqlFailoverGroup) Dependencies() terra.Dependencies {
	return sfg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) LifecycleManagement() *terra.Lifecycle {
	return sfg.Lifecycle
}

// Attributes returns the attributes for [SqlFailoverGroup].
func (sfg *SqlFailoverGroup) Attributes() sqlFailoverGroupAttributes {
	return sqlFailoverGroupAttributes{ref: terra.ReferenceResource(sfg)}
}

// ImportState imports the given attribute values into [SqlFailoverGroup]'s state.
func (sfg *SqlFailoverGroup) ImportState(av io.Reader) error {
	sfg.state = &sqlFailoverGroupState{}
	if err := json.NewDecoder(av).Decode(sfg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfg.Type(), sfg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlFailoverGroup] has state.
func (sfg *SqlFailoverGroup) State() (*sqlFailoverGroupState, bool) {
	return sfg.state, sfg.state != nil
}

// StateMust returns the state for [SqlFailoverGroup]. Panics if the state is nil.
func (sfg *SqlFailoverGroup) StateMust() *sqlFailoverGroupState {
	if sfg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfg.Type(), sfg.LocalName()))
	}
	return sfg.state
}

// SqlFailoverGroupArgs contains the configurations for azurerm_sql_failover_group.
type SqlFailoverGroupArgs struct {
	// Databases: set of string, optional
	Databases terra.SetValue[terra.StringValue] `hcl:"databases,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PartnerServers: min=1
	PartnerServers []sqlfailovergroup.PartnerServers `hcl:"partner_servers,block" validate:"min=1"`
	// ReadWriteEndpointFailoverPolicy: required
	ReadWriteEndpointFailoverPolicy *sqlfailovergroup.ReadWriteEndpointFailoverPolicy `hcl:"read_write_endpoint_failover_policy,block" validate:"required"`
	// ReadonlyEndpointFailoverPolicy: optional
	ReadonlyEndpointFailoverPolicy *sqlfailovergroup.ReadonlyEndpointFailoverPolicy `hcl:"readonly_endpoint_failover_policy,block"`
	// Timeouts: optional
	Timeouts *sqlfailovergroup.Timeouts `hcl:"timeouts,block"`
}
type sqlFailoverGroupAttributes struct {
	ref terra.Reference
}

// Databases returns a reference to field databases of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Databases() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sfg.ref.Append("databases"))
}

// Id returns a reference to field id of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("resource_group_name"))
}

// Role returns a reference to field role of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("role"))
}

// ServerName returns a reference to field server_name of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("server_name"))
}

// Tags returns a reference to field tags of azurerm_sql_failover_group.
func (sfg sqlFailoverGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfg.ref.Append("tags"))
}

func (sfg sqlFailoverGroupAttributes) PartnerServers() terra.ListValue[sqlfailovergroup.PartnerServersAttributes] {
	return terra.ReferenceAsList[sqlfailovergroup.PartnerServersAttributes](sfg.ref.Append("partner_servers"))
}

func (sfg sqlFailoverGroupAttributes) ReadWriteEndpointFailoverPolicy() terra.ListValue[sqlfailovergroup.ReadWriteEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[sqlfailovergroup.ReadWriteEndpointFailoverPolicyAttributes](sfg.ref.Append("read_write_endpoint_failover_policy"))
}

func (sfg sqlFailoverGroupAttributes) ReadonlyEndpointFailoverPolicy() terra.ListValue[sqlfailovergroup.ReadonlyEndpointFailoverPolicyAttributes] {
	return terra.ReferenceAsList[sqlfailovergroup.ReadonlyEndpointFailoverPolicyAttributes](sfg.ref.Append("readonly_endpoint_failover_policy"))
}

func (sfg sqlFailoverGroupAttributes) Timeouts() sqlfailovergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlfailovergroup.TimeoutsAttributes](sfg.ref.Append("timeouts"))
}

type sqlFailoverGroupState struct {
	Databases                       []string                                                `json:"databases"`
	Id                              string                                                  `json:"id"`
	Location                        string                                                  `json:"location"`
	Name                            string                                                  `json:"name"`
	ResourceGroupName               string                                                  `json:"resource_group_name"`
	Role                            string                                                  `json:"role"`
	ServerName                      string                                                  `json:"server_name"`
	Tags                            map[string]string                                       `json:"tags"`
	PartnerServers                  []sqlfailovergroup.PartnerServersState                  `json:"partner_servers"`
	ReadWriteEndpointFailoverPolicy []sqlfailovergroup.ReadWriteEndpointFailoverPolicyState `json:"read_write_endpoint_failover_policy"`
	ReadonlyEndpointFailoverPolicy  []sqlfailovergroup.ReadonlyEndpointFailoverPolicyState  `json:"readonly_endpoint_failover_policy"`
	Timeouts                        *sqlfailovergroup.TimeoutsState                         `json:"timeouts"`
}
