// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cosmosdbpostgresqlcoordinatorconfiguration "github.com/golingon/terraproviders/azurerm/3.98.0/cosmosdbpostgresqlcoordinatorconfiguration"
	"io"
)

// NewCosmosdbPostgresqlCoordinatorConfiguration creates a new instance of [CosmosdbPostgresqlCoordinatorConfiguration].
func NewCosmosdbPostgresqlCoordinatorConfiguration(name string, args CosmosdbPostgresqlCoordinatorConfigurationArgs) *CosmosdbPostgresqlCoordinatorConfiguration {
	return &CosmosdbPostgresqlCoordinatorConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbPostgresqlCoordinatorConfiguration)(nil)

// CosmosdbPostgresqlCoordinatorConfiguration represents the Terraform resource azurerm_cosmosdb_postgresql_coordinator_configuration.
type CosmosdbPostgresqlCoordinatorConfiguration struct {
	Name      string
	Args      CosmosdbPostgresqlCoordinatorConfigurationArgs
	state     *cosmosdbPostgresqlCoordinatorConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) Type() string {
	return "azurerm_cosmosdb_postgresql_coordinator_configuration"
}

// LocalName returns the local name for [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) LocalName() string {
	return cpcc.Name
}

// Configuration returns the configuration (args) for [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) Configuration() interface{} {
	return cpcc.Args
}

// DependOn is used for other resources to depend on [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(cpcc)
}

// Dependencies returns the list of resources [CosmosdbPostgresqlCoordinatorConfiguration] depends_on.
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) Dependencies() terra.Dependencies {
	return cpcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) LifecycleManagement() *terra.Lifecycle {
	return cpcc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbPostgresqlCoordinatorConfiguration].
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) Attributes() cosmosdbPostgresqlCoordinatorConfigurationAttributes {
	return cosmosdbPostgresqlCoordinatorConfigurationAttributes{ref: terra.ReferenceResource(cpcc)}
}

// ImportState imports the given attribute values into [CosmosdbPostgresqlCoordinatorConfiguration]'s state.
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) ImportState(av io.Reader) error {
	cpcc.state = &cosmosdbPostgresqlCoordinatorConfigurationState{}
	if err := json.NewDecoder(av).Decode(cpcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpcc.Type(), cpcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbPostgresqlCoordinatorConfiguration] has state.
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) State() (*cosmosdbPostgresqlCoordinatorConfigurationState, bool) {
	return cpcc.state, cpcc.state != nil
}

// StateMust returns the state for [CosmosdbPostgresqlCoordinatorConfiguration]. Panics if the state is nil.
func (cpcc *CosmosdbPostgresqlCoordinatorConfiguration) StateMust() *cosmosdbPostgresqlCoordinatorConfigurationState {
	if cpcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpcc.Type(), cpcc.LocalName()))
	}
	return cpcc.state
}

// CosmosdbPostgresqlCoordinatorConfigurationArgs contains the configurations for azurerm_cosmosdb_postgresql_coordinator_configuration.
type CosmosdbPostgresqlCoordinatorConfigurationArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbpostgresqlcoordinatorconfiguration.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbPostgresqlCoordinatorConfigurationAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_cosmosdb_postgresql_coordinator_configuration.
func (cpcc cosmosdbPostgresqlCoordinatorConfigurationAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(cpcc.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_postgresql_coordinator_configuration.
func (cpcc cosmosdbPostgresqlCoordinatorConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpcc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_postgresql_coordinator_configuration.
func (cpcc cosmosdbPostgresqlCoordinatorConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpcc.ref.Append("name"))
}

// Value returns a reference to field value of azurerm_cosmosdb_postgresql_coordinator_configuration.
func (cpcc cosmosdbPostgresqlCoordinatorConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cpcc.ref.Append("value"))
}

func (cpcc cosmosdbPostgresqlCoordinatorConfigurationAttributes) Timeouts() cosmosdbpostgresqlcoordinatorconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbpostgresqlcoordinatorconfiguration.TimeoutsAttributes](cpcc.ref.Append("timeouts"))
}

type cosmosdbPostgresqlCoordinatorConfigurationState struct {
	ClusterId string                                                    `json:"cluster_id"`
	Id        string                                                    `json:"id"`
	Name      string                                                    `json:"name"`
	Value     string                                                    `json:"value"`
	Timeouts  *cosmosdbpostgresqlcoordinatorconfiguration.TimeoutsState `json:"timeouts"`
}
