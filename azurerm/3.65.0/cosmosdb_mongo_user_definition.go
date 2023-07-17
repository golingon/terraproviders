// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbmongouserdefinition "github.com/golingon/terraproviders/azurerm/3.65.0/cosmosdbmongouserdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbMongoUserDefinition creates a new instance of [CosmosdbMongoUserDefinition].
func NewCosmosdbMongoUserDefinition(name string, args CosmosdbMongoUserDefinitionArgs) *CosmosdbMongoUserDefinition {
	return &CosmosdbMongoUserDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbMongoUserDefinition)(nil)

// CosmosdbMongoUserDefinition represents the Terraform resource azurerm_cosmosdb_mongo_user_definition.
type CosmosdbMongoUserDefinition struct {
	Name      string
	Args      CosmosdbMongoUserDefinitionArgs
	state     *cosmosdbMongoUserDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) Type() string {
	return "azurerm_cosmosdb_mongo_user_definition"
}

// LocalName returns the local name for [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) LocalName() string {
	return cmud.Name
}

// Configuration returns the configuration (args) for [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) Configuration() interface{} {
	return cmud.Args
}

// DependOn is used for other resources to depend on [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(cmud)
}

// Dependencies returns the list of resources [CosmosdbMongoUserDefinition] depends_on.
func (cmud *CosmosdbMongoUserDefinition) Dependencies() terra.Dependencies {
	return cmud.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) LifecycleManagement() *terra.Lifecycle {
	return cmud.Lifecycle
}

// Attributes returns the attributes for [CosmosdbMongoUserDefinition].
func (cmud *CosmosdbMongoUserDefinition) Attributes() cosmosdbMongoUserDefinitionAttributes {
	return cosmosdbMongoUserDefinitionAttributes{ref: terra.ReferenceResource(cmud)}
}

// ImportState imports the given attribute values into [CosmosdbMongoUserDefinition]'s state.
func (cmud *CosmosdbMongoUserDefinition) ImportState(av io.Reader) error {
	cmud.state = &cosmosdbMongoUserDefinitionState{}
	if err := json.NewDecoder(av).Decode(cmud.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmud.Type(), cmud.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbMongoUserDefinition] has state.
func (cmud *CosmosdbMongoUserDefinition) State() (*cosmosdbMongoUserDefinitionState, bool) {
	return cmud.state, cmud.state != nil
}

// StateMust returns the state for [CosmosdbMongoUserDefinition]. Panics if the state is nil.
func (cmud *CosmosdbMongoUserDefinition) StateMust() *cosmosdbMongoUserDefinitionState {
	if cmud.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmud.Type(), cmud.LocalName()))
	}
	return cmud.state
}

// CosmosdbMongoUserDefinitionArgs contains the configurations for azurerm_cosmosdb_mongo_user_definition.
type CosmosdbMongoUserDefinitionArgs struct {
	// CosmosMongoDatabaseId: string, required
	CosmosMongoDatabaseId terra.StringValue `hcl:"cosmos_mongo_database_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InheritedRoleNames: list of string, optional
	InheritedRoleNames terra.ListValue[terra.StringValue] `hcl:"inherited_role_names,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cosmosdbmongouserdefinition.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbMongoUserDefinitionAttributes struct {
	ref terra.Reference
}

// CosmosMongoDatabaseId returns a reference to field cosmos_mongo_database_id of azurerm_cosmosdb_mongo_user_definition.
func (cmud cosmosdbMongoUserDefinitionAttributes) CosmosMongoDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(cmud.ref.Append("cosmos_mongo_database_id"))
}

// Id returns a reference to field id of azurerm_cosmosdb_mongo_user_definition.
func (cmud cosmosdbMongoUserDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmud.ref.Append("id"))
}

// InheritedRoleNames returns a reference to field inherited_role_names of azurerm_cosmosdb_mongo_user_definition.
func (cmud cosmosdbMongoUserDefinitionAttributes) InheritedRoleNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmud.ref.Append("inherited_role_names"))
}

// Password returns a reference to field password of azurerm_cosmosdb_mongo_user_definition.
func (cmud cosmosdbMongoUserDefinitionAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(cmud.ref.Append("password"))
}

// Username returns a reference to field username of azurerm_cosmosdb_mongo_user_definition.
func (cmud cosmosdbMongoUserDefinitionAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(cmud.ref.Append("username"))
}

func (cmud cosmosdbMongoUserDefinitionAttributes) Timeouts() cosmosdbmongouserdefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbmongouserdefinition.TimeoutsAttributes](cmud.ref.Append("timeouts"))
}

type cosmosdbMongoUserDefinitionState struct {
	CosmosMongoDatabaseId string                                     `json:"cosmos_mongo_database_id"`
	Id                    string                                     `json:"id"`
	InheritedRoleNames    []string                                   `json:"inherited_role_names"`
	Password              string                                     `json:"password"`
	Username              string                                     `json:"username"`
	Timeouts              *cosmosdbmongouserdefinition.TimeoutsState `json:"timeouts"`
}
