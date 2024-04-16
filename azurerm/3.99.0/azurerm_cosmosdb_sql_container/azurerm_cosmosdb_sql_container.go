// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cosmosdb_sql_container

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

// Resource represents the Terraform resource azurerm_cosmosdb_sql_container.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCosmosdbSqlContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acsc *Resource) Type() string {
	return "azurerm_cosmosdb_sql_container"
}

// LocalName returns the local name for [Resource].
func (acsc *Resource) LocalName() string {
	return acsc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acsc *Resource) Configuration() interface{} {
	return acsc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acsc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acsc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acsc *Resource) Dependencies() terra.Dependencies {
	return acsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acsc *Resource) LifecycleManagement() *terra.Lifecycle {
	return acsc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acsc *Resource) Attributes() azurermCosmosdbSqlContainerAttributes {
	return azurermCosmosdbSqlContainerAttributes{ref: terra.ReferenceResource(acsc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acsc *Resource) ImportState(state io.Reader) error {
	acsc.state = &azurermCosmosdbSqlContainerState{}
	if err := json.NewDecoder(state).Decode(acsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acsc.Type(), acsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acsc *Resource) State() (*azurermCosmosdbSqlContainerState, bool) {
	return acsc.state, acsc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acsc *Resource) StateMust() *azurermCosmosdbSqlContainerState {
	if acsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acsc.Type(), acsc.LocalName()))
	}
	return acsc.state
}

// Args contains the configurations for azurerm_cosmosdb_sql_container.
type Args struct {
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
	AutoscaleSettings *AutoscaleSettings `hcl:"autoscale_settings,block"`
	// ConflictResolutionPolicy: optional
	ConflictResolutionPolicy *ConflictResolutionPolicy `hcl:"conflict_resolution_policy,block"`
	// IndexingPolicy: optional
	IndexingPolicy *IndexingPolicy `hcl:"indexing_policy,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// UniqueKey: min=0
	UniqueKey []UniqueKey `hcl:"unique_key,block" validate:"min=0"`
}

type azurermCosmosdbSqlContainerAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("account_name"))
}

// AnalyticalStorageTtl returns a reference to field analytical_storage_ttl of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) AnalyticalStorageTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(acsc.ref.Append("analytical_storage_ttl"))
}

// DatabaseName returns a reference to field database_name of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("database_name"))
}

// DefaultTtl returns a reference to field default_ttl of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) DefaultTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(acsc.ref.Append("default_ttl"))
}

// Id returns a reference to field id of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("name"))
}

// PartitionKeyPath returns a reference to field partition_key_path of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) PartitionKeyPath() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("partition_key_path"))
}

// PartitionKeyVersion returns a reference to field partition_key_version of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) PartitionKeyVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(acsc.ref.Append("partition_key_version"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(acsc.ref.Append("resource_group_name"))
}

// Throughput returns a reference to field throughput of azurerm_cosmosdb_sql_container.
func (acsc azurermCosmosdbSqlContainerAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(acsc.ref.Append("throughput"))
}

func (acsc azurermCosmosdbSqlContainerAttributes) AutoscaleSettings() terra.ListValue[AutoscaleSettingsAttributes] {
	return terra.ReferenceAsList[AutoscaleSettingsAttributes](acsc.ref.Append("autoscale_settings"))
}

func (acsc azurermCosmosdbSqlContainerAttributes) ConflictResolutionPolicy() terra.ListValue[ConflictResolutionPolicyAttributes] {
	return terra.ReferenceAsList[ConflictResolutionPolicyAttributes](acsc.ref.Append("conflict_resolution_policy"))
}

func (acsc azurermCosmosdbSqlContainerAttributes) IndexingPolicy() terra.ListValue[IndexingPolicyAttributes] {
	return terra.ReferenceAsList[IndexingPolicyAttributes](acsc.ref.Append("indexing_policy"))
}

func (acsc azurermCosmosdbSqlContainerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acsc.ref.Append("timeouts"))
}

func (acsc azurermCosmosdbSqlContainerAttributes) UniqueKey() terra.SetValue[UniqueKeyAttributes] {
	return terra.ReferenceAsSet[UniqueKeyAttributes](acsc.ref.Append("unique_key"))
}

type azurermCosmosdbSqlContainerState struct {
	AccountName              string                          `json:"account_name"`
	AnalyticalStorageTtl     float64                         `json:"analytical_storage_ttl"`
	DatabaseName             string                          `json:"database_name"`
	DefaultTtl               float64                         `json:"default_ttl"`
	Id                       string                          `json:"id"`
	Name                     string                          `json:"name"`
	PartitionKeyPath         string                          `json:"partition_key_path"`
	PartitionKeyVersion      float64                         `json:"partition_key_version"`
	ResourceGroupName        string                          `json:"resource_group_name"`
	Throughput               float64                         `json:"throughput"`
	AutoscaleSettings        []AutoscaleSettingsState        `json:"autoscale_settings"`
	ConflictResolutionPolicy []ConflictResolutionPolicyState `json:"conflict_resolution_policy"`
	IndexingPolicy           []IndexingPolicyState           `json:"indexing_policy"`
	Timeouts                 *TimeoutsState                  `json:"timeouts"`
	UniqueKey                []UniqueKeyState                `json:"unique_key"`
}
