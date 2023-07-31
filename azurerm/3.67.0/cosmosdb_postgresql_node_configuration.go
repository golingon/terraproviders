// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbpostgresqlnodeconfiguration "github.com/golingon/terraproviders/azurerm/3.67.0/cosmosdbpostgresqlnodeconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbPostgresqlNodeConfiguration creates a new instance of [CosmosdbPostgresqlNodeConfiguration].
func NewCosmosdbPostgresqlNodeConfiguration(name string, args CosmosdbPostgresqlNodeConfigurationArgs) *CosmosdbPostgresqlNodeConfiguration {
	return &CosmosdbPostgresqlNodeConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbPostgresqlNodeConfiguration)(nil)

// CosmosdbPostgresqlNodeConfiguration represents the Terraform resource azurerm_cosmosdb_postgresql_node_configuration.
type CosmosdbPostgresqlNodeConfiguration struct {
	Name      string
	Args      CosmosdbPostgresqlNodeConfigurationArgs
	state     *cosmosdbPostgresqlNodeConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) Type() string {
	return "azurerm_cosmosdb_postgresql_node_configuration"
}

// LocalName returns the local name for [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) LocalName() string {
	return cpnc.Name
}

// Configuration returns the configuration (args) for [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) Configuration() interface{} {
	return cpnc.Args
}

// DependOn is used for other resources to depend on [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(cpnc)
}

// Dependencies returns the list of resources [CosmosdbPostgresqlNodeConfiguration] depends_on.
func (cpnc *CosmosdbPostgresqlNodeConfiguration) Dependencies() terra.Dependencies {
	return cpnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) LifecycleManagement() *terra.Lifecycle {
	return cpnc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbPostgresqlNodeConfiguration].
func (cpnc *CosmosdbPostgresqlNodeConfiguration) Attributes() cosmosdbPostgresqlNodeConfigurationAttributes {
	return cosmosdbPostgresqlNodeConfigurationAttributes{ref: terra.ReferenceResource(cpnc)}
}

// ImportState imports the given attribute values into [CosmosdbPostgresqlNodeConfiguration]'s state.
func (cpnc *CosmosdbPostgresqlNodeConfiguration) ImportState(av io.Reader) error {
	cpnc.state = &cosmosdbPostgresqlNodeConfigurationState{}
	if err := json.NewDecoder(av).Decode(cpnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpnc.Type(), cpnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbPostgresqlNodeConfiguration] has state.
func (cpnc *CosmosdbPostgresqlNodeConfiguration) State() (*cosmosdbPostgresqlNodeConfigurationState, bool) {
	return cpnc.state, cpnc.state != nil
}

// StateMust returns the state for [CosmosdbPostgresqlNodeConfiguration]. Panics if the state is nil.
func (cpnc *CosmosdbPostgresqlNodeConfiguration) StateMust() *cosmosdbPostgresqlNodeConfigurationState {
	if cpnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpnc.Type(), cpnc.LocalName()))
	}
	return cpnc.state
}

// CosmosdbPostgresqlNodeConfigurationArgs contains the configurations for azurerm_cosmosdb_postgresql_node_configuration.
type CosmosdbPostgresqlNodeConfigurationArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbpostgresqlnodeconfiguration.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbPostgresqlNodeConfigurationAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_cosmosdb_postgresql_node_configuration.
func (cpnc cosmosdbPostgresqlNodeConfigurationAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(cpnc.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_postgresql_node_configuration.
func (cpnc cosmosdbPostgresqlNodeConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpnc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_postgresql_node_configuration.
func (cpnc cosmosdbPostgresqlNodeConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpnc.ref.Append("name"))
}

// Value returns a reference to field value of azurerm_cosmosdb_postgresql_node_configuration.
func (cpnc cosmosdbPostgresqlNodeConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cpnc.ref.Append("value"))
}

func (cpnc cosmosdbPostgresqlNodeConfigurationAttributes) Timeouts() cosmosdbpostgresqlnodeconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbpostgresqlnodeconfiguration.TimeoutsAttributes](cpnc.ref.Append("timeouts"))
}

type cosmosdbPostgresqlNodeConfigurationState struct {
	ClusterId string                                             `json:"cluster_id"`
	Id        string                                             `json:"id"`
	Name      string                                             `json:"name"`
	Value     string                                             `json:"value"`
	Timeouts  *cosmosdbpostgresqlnodeconfiguration.TimeoutsState `json:"timeouts"`
}
