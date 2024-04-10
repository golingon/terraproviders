// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cosmosdbsqlcontainer "github.com/golingon/terraproviders/azurerm/3.98.0/cosmosdbsqlcontainer"
	"io"
)

// NewCosmosdbSqlContainer creates a new instance of [CosmosdbSqlContainer].
func NewCosmosdbSqlContainer(name string, args CosmosdbSqlContainerArgs) *CosmosdbSqlContainer {
	return &CosmosdbSqlContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbSqlContainer)(nil)

// CosmosdbSqlContainer represents the Terraform resource azurerm_cosmosdb_sql_container.
type CosmosdbSqlContainer struct {
	Name      string
	Args      CosmosdbSqlContainerArgs
	state     *cosmosdbSqlContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) Type() string {
	return "azurerm_cosmosdb_sql_container"
}

// LocalName returns the local name for [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) LocalName() string {
	return csc.Name
}

// Configuration returns the configuration (args) for [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) Configuration() interface{} {
	return csc.Args
}

// DependOn is used for other resources to depend on [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) DependOn() terra.Reference {
	return terra.ReferenceResource(csc)
}

// Dependencies returns the list of resources [CosmosdbSqlContainer] depends_on.
func (csc *CosmosdbSqlContainer) Dependencies() terra.Dependencies {
	return csc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) LifecycleManagement() *terra.Lifecycle {
	return csc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbSqlContainer].
func (csc *CosmosdbSqlContainer) Attributes() cosmosdbSqlContainerAttributes {
	return cosmosdbSqlContainerAttributes{ref: terra.ReferenceResource(csc)}
}

// ImportState imports the given attribute values into [CosmosdbSqlContainer]'s state.
func (csc *CosmosdbSqlContainer) ImportState(av io.Reader) error {
	csc.state = &cosmosdbSqlContainerState{}
	if err := json.NewDecoder(av).Decode(csc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csc.Type(), csc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbSqlContainer] has state.
func (csc *CosmosdbSqlContainer) State() (*cosmosdbSqlContainerState, bool) {
	return csc.state, csc.state != nil
}

// StateMust returns the state for [CosmosdbSqlContainer]. Panics if the state is nil.
func (csc *CosmosdbSqlContainer) StateMust() *cosmosdbSqlContainerState {
	if csc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csc.Type(), csc.LocalName()))
	}
	return csc.state
}

// CosmosdbSqlContainerArgs contains the configurations for azurerm_cosmosdb_sql_container.
type CosmosdbSqlContainerArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AnalyticalStorageTtl: number, optional
	AnalyticalStorageTtl terra.NumberValue `hcl:"analytical_storage_ttl,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// DefaultTtl: number, optional
	DefaultTtl terra.NumberValue `hcl:"default_ttl,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKeyPath: string, required
	PartitionKeyPath terra.StringValue `hcl:"partition_key_path,attr" validate:"required"`
	// PartitionKeyVersion: number, optional
	PartitionKeyVersion terra.NumberValue `hcl:"partition_key_version,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// AutoscaleSettings: optional
	AutoscaleSettings *cosmosdbsqlcontainer.AutoscaleSettings `hcl:"autoscale_settings,block"`
	// ConflictResolutionPolicy: optional
	ConflictResolutionPolicy *cosmosdbsqlcontainer.ConflictResolutionPolicy `hcl:"conflict_resolution_policy,block"`
	// IndexingPolicy: optional
	IndexingPolicy *cosmosdbsqlcontainer.IndexingPolicy `hcl:"indexing_policy,block"`
	// Timeouts: optional
	Timeouts *cosmosdbsqlcontainer.Timeouts `hcl:"timeouts,block"`
	// UniqueKey: min=0
	UniqueKey []cosmosdbsqlcontainer.UniqueKey `hcl:"unique_key,block" validate:"min=0"`
}
type cosmosdbSqlContainerAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("account_name"))
}

// AnalyticalStorageTtl returns a reference to field analytical_storage_ttl of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(csc.ref.Append("analytical_storage_ttl"))
}

// DatabaseName returns a reference to field database_name of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("database_name"))
}

// DefaultTtl returns a reference to field default_ttl of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(csc.ref.Append("default_ttl"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("name"))
}

// PartitionKeyPath returns a reference to field partition_key_path of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) PartitionKeyPath() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("partition_key_path"))
}

// PartitionKeyVersion returns a reference to field partition_key_version of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) PartitionKeyVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(csc.ref.Append("partition_key_version"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(csc.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_sql_container.
func (csc cosmosdbSqlContainerAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(csc.ref.Append("throughput"))
}

func (csc cosmosdbSqlContainerAttributes) AutoscaleSettings() terra.ListValue[cosmosdbsqlcontainer.AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[cosmosdbsqlcontainer.AutoscaleSettingsAttributes](csc.ref.Append("autoscale_settings"))
}

func (csc cosmosdbSqlContainerAttributes) ConflictResolutionPolicy() terra.ListValue[cosmosdbsqlcontainer.ConflictResolutionPolicyAttributes] {
	return terra.ReferenceAsList[cosmosdbsqlcontainer.ConflictResolutionPolicyAttributes](csc.ref.Append("conflict_resolution_policy"))
}

func (csc cosmosdbSqlContainerAttributes) IndexingPolicy() terra.ListValue[cosmosdbsqlcontainer.IndexingPolicyAttributes] {
	return terra.ReferenceAsList[cosmosdbsqlcontainer.IndexingPolicyAttributes](csc.ref.Append("indexing_policy"))
}

func (csc cosmosdbSqlContainerAttributes) Timeouts() cosmosdbsqlcontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbsqlcontainer.TimeoutsAttributes](csc.ref.Append("timeouts"))
}

func (csc cosmosdbSqlContainerAttributes) UniqueKey() terra.SetValue[cosmosdbsqlcontainer.UniqueKeyAttributes] {
	return terra.ReferenceAsSet[cosmosdbsqlcontainer.UniqueKeyAttributes](csc.ref.Append("unique_key"))
}

type cosmosdbSqlContainerState struct {
	AccountName              string                                               `json:"account_name"`
	AnalyticalStorageTtl     float64                                              `json:"analytical_storage_ttl"`
	DatabaseName             string                                               `json:"database_name"`
	DefaultTtl               float64                                              `json:"default_ttl"`
	Id                       string                                               `json:"id"`
	Name                     string                                               `json:"name"`
	PartitionKeyPath         string                                               `json:"partition_key_path"`
	PartitionKeyVersion      float64                                              `json:"partition_key_version"`
	ResourceGroupName        string                                               `json:"resource_group_name"`
	Throughput               float64                                              `json:"throughput"`
	AutoscaleSettings        []cosmosdbsqlcontainer.AutoscaleSettingsState        `json:"autoscale_settings"`
	ConflictResolutionPolicy []cosmosdbsqlcontainer.ConflictResolutionPolicyState `json:"conflict_resolution_policy"`
	IndexingPolicy           []cosmosdbsqlcontainer.IndexingPolicyState           `json:"indexing_policy"`
	Timeouts                 *cosmosdbsqlcontainer.TimeoutsState                  `json:"timeouts"`
	UniqueKey                []cosmosdbsqlcontainer.UniqueKeyState                `json:"unique_key"`
}
