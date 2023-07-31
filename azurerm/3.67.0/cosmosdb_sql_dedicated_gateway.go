// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbsqldedicatedgateway "github.com/golingon/terraproviders/azurerm/3.67.0/cosmosdbsqldedicatedgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbSqlDedicatedGateway creates a new instance of [CosmosdbSqlDedicatedGateway].
func NewCosmosdbSqlDedicatedGateway(name string, args CosmosdbSqlDedicatedGatewayArgs) *CosmosdbSqlDedicatedGateway {
	return &CosmosdbSqlDedicatedGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlDedicatedGateway)(nil)

// CosmosdbSqlDedicatedGateway represents the Terraform resource azurerm_cosmosdb_sql_dedicated_gateway.
type CosmosdbSqlDedicatedGateway struct {
	Name      string
	Args      CosmosdbSqlDedicatedGatewayArgs
	state     *cosmosdbSqlDedicatedGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) Type() string {
	return "azurerm_cosmosdb_sql_dedicated_gateway"
}

// LocalName returns the local name for [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) LocalName() string {
	return csdg.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) Configuration() interface{} {
	return csdg.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(csdg)
}

// Dependencies returns the list of resources [CosmosdbSqlDedicatedGateway] depends_on.
func (csdg *CosmosdbSqlDedicatedGateway) Dependencies() terra.Dependencies {
	return csdg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) LifecycleManagement() *terra.Lifecycle {
	return csdg.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlDedicatedGateway].
func (csdg *CosmosdbSqlDedicatedGateway) Attributes() cosmosdbSqlDedicatedGatewayAttributes {
	return cosmosdbSqlDedicatedGatewayAttributes{ref: terra.ReferenceResource(csdg)}
}

// ImportState imports the given attribute values into [CosmosdbSqlDedicatedGateway]'s state.
func (csdg *CosmosdbSqlDedicatedGateway) ImportState(av io.Reader) error {
	csdg.state = &cosmosdbSqlDedicatedGatewayState{}
	if err := json.NewDecoder(av).Decode(csdg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csdg.Type(), csdg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlDedicatedGateway] has state.
func (csdg *CosmosdbSqlDedicatedGateway) State() (*cosmosdbSqlDedicatedGatewayState, bool) {
	return csdg.state, csdg.state != nil
}

// StateMust returns the state for [CosmosdbSqlDedicatedGateway]. Panics if the state is nil.
func (csdg *CosmosdbSqlDedicatedGateway) StateMust() *cosmosdbSqlDedicatedGatewayState {
	if csdg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csdg.Type(), csdg.LocalName()))
	}
	return csdg.state
}

// CosmosdbSqlDedicatedGatewayArgs contains the configurations for azurerm_cosmosdb_sql_dedicated_gateway.
type CosmosdbSqlDedicatedGatewayArgs struct {
	// CosmosdbAccountId: string, required
	CosmosdbAccountId terra.StringValue `hcl:"cosmosdb_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceCount: number, required
	InstanceCount terra.NumberValue `hcl:"instance_count,attr" validate:"required"`
	// InstanceSize: string, required
	InstanceSize terra.StringValue `hcl:"instance_size,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbsqldedicatedgateway.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbSqlDedicatedGatewayAttributes struct {
	ref terra.Reference
}

// CosmosdbAccountId returns a reference to field cosmosdb_account_id of azurerm_cosmosdb_sql_dedicated_gateway.
func (csdg cosmosdbSqlDedicatedGatewayAttributes) CosmosdbAccountId() terra.StringValue {
	return terra.ReferenceAsString(csdg.ref.Append("cosmosdb_account_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_dedicated_gateway.
func (csdg cosmosdbSqlDedicatedGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csdg.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of azurerm_cosmosdb_sql_dedicated_gateway.
func (csdg cosmosdbSqlDedicatedGatewayAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(csdg.ref.Append("instance_count"))
}

// InstanceSize returns a reference to field instance_size of azurerm_cosmosdb_sql_dedicated_gateway.
func (csdg cosmosdbSqlDedicatedGatewayAttributes) InstanceSize() terra.StringValue {
	return terra.ReferenceAsString(csdg.ref.Append("instance_size"))
}

func (csdg cosmosdbSqlDedicatedGatewayAttributes) Timeouts() cosmosdbsqldedicatedgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqldedicatedgateway.TimeoutsAttributes](csdg.ref.Append("timeouts"))
}

type cosmosdbSqlDedicatedGatewayState struct {
	CosmosdbAccountId string                                     `json:"cosmosdb_account_id"`
	Id                string                                     `json:"id"`
	InstanceCount     float64                                    `json:"instance_count"`
	InstanceSize      string                                     `json:"instance_size"`
	Timeouts          *cosmosdbsqldedicatedgateway.TimeoutsState `json:"timeouts"`
}
