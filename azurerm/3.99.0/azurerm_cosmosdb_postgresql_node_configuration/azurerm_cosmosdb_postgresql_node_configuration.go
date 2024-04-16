// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cosmosdb_postgresql_node_configuration

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

// Resource represents the Terraform resource azurerm_cosmosdb_postgresql_node_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCosmosdbPostgresqlNodeConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acpnc *Resource) Type() string {
	return "azurerm_cosmosdb_postgresql_node_configuration"
}

// LocalName returns the local name for [Resource].
func (acpnc *Resource) LocalName() string {
	return acpnc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acpnc *Resource) Configuration() interface{} {
	return acpnc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acpnc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acpnc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acpnc *Resource) Dependencies() terra.Dependencies {
	return acpnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acpnc *Resource) LifecycleManagement() *terra.Lifecycle {
	return acpnc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acpnc *Resource) Attributes() azurermCosmosdbPostgresqlNodeConfigurationAttributes {
	return azurermCosmosdbPostgresqlNodeConfigurationAttributes{ref: terra.ReferenceResource(acpnc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acpnc *Resource) ImportState(state io.Reader) error {
	acpnc.state = &azurermCosmosdbPostgresqlNodeConfigurationState{}
	if err := json.NewDecoder(state).Decode(acpnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acpnc.Type(), acpnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acpnc *Resource) State() (*azurermCosmosdbPostgresqlNodeConfigurationState, bool) {
	return acpnc.state, acpnc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acpnc *Resource) StateMust() *azurermCosmosdbPostgresqlNodeConfigurationState {
	if acpnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acpnc.Type(), acpnc.LocalName()))
	}
	return acpnc.state
}

// Args contains the configurations for azurerm_cosmosdb_postgresql_node_configuration.
type Args struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCosmosdbPostgresqlNodeConfigurationAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_cosmosdb_postgresql_node_configuration.
func (acpnc azurermCosmosdbPostgresqlNodeConfigurationAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(acpnc.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_postgresql_node_configuration.
func (acpnc azurermCosmosdbPostgresqlNodeConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acpnc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_postgresql_node_configuration.
func (acpnc azurermCosmosdbPostgresqlNodeConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acpnc.ref.Append("name"))
}

// Value returns a reference to field value of azurerm_cosmosdb_postgresql_node_configuration.
func (acpnc azurermCosmosdbPostgresqlNodeConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(acpnc.ref.Append("value"))
}

func (acpnc azurermCosmosdbPostgresqlNodeConfigurationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acpnc.ref.Append("timeouts"))
}

type azurermCosmosdbPostgresqlNodeConfigurationState struct {
	ClusterId string         `json:"cluster_id"`
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Value     string         `json:"value"`
	Timeouts  *TimeoutsState `json:"timeouts"`
}
