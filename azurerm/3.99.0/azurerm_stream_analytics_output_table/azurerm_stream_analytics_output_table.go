// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_stream_analytics_output_table

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

// Resource represents the Terraform resource azurerm_stream_analytics_output_table.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStreamAnalyticsOutputTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asaot *Resource) Type() string {
	return "azurerm_stream_analytics_output_table"
}

// LocalName returns the local name for [Resource].
func (asaot *Resource) LocalName() string {
	return asaot.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asaot *Resource) Configuration() interface{} {
	return asaot.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asaot *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asaot)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asaot *Resource) Dependencies() terra.Dependencies {
	return asaot.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asaot *Resource) LifecycleManagement() *terra.Lifecycle {
	return asaot.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asaot *Resource) Attributes() azurermStreamAnalyticsOutputTableAttributes {
	return azurermStreamAnalyticsOutputTableAttributes{ref: terra.ReferenceResource(asaot)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asaot *Resource) ImportState(state io.Reader) error {
	asaot.state = &azurermStreamAnalyticsOutputTableState{}
	if err := json.NewDecoder(state).Decode(asaot.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asaot.Type(), asaot.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asaot *Resource) State() (*azurermStreamAnalyticsOutputTableState, bool) {
	return asaot.state, asaot.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asaot *Resource) StateMust() *azurermStreamAnalyticsOutputTableState {
	if asaot.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asaot.Type(), asaot.LocalName()))
	}
	return asaot.state
}

// Args contains the configurations for azurerm_stream_analytics_output_table.
type Args struct {
	// BatchSize: number, required
	BatchSize terra.NumberValue `hcl:"batch_size,attr" validate:"required"`
	// ColumnsToRemove: list of string, optional
	ColumnsToRemove terra.ListValue[terra.StringValue] `hcl:"columns_to_remove,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKey: string, required
	PartitionKey terra.StringValue `hcl:"partition_key,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RowKey: string, required
	RowKey terra.StringValue `hcl:"row_key,attr" validate:"required"`
	// StorageAccountKey: string, required
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStreamAnalyticsOutputTableAttributes struct {
	ref terra.Reference
}

// BatchSize returns a reference to field batch_size of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) BatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(asaot.ref.Append("batch_size"))
}

// ColumnsToRemove returns a reference to field columns_to_remove of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) ColumnsToRemove() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asaot.ref.Append("columns_to_remove"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("name"))
}

// PartitionKey returns a reference to field partition_key of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("partition_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("resource_group_name"))
}

// RowKey returns a reference to field row_key of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) RowKey() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("row_key"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("storage_account_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("storage_account_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("stream_analytics_job_name"))
}

// Table returns a reference to field table of azurerm_stream_analytics_output_table.
func (asaot azurermStreamAnalyticsOutputTableAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(asaot.ref.Append("table"))
}

func (asaot azurermStreamAnalyticsOutputTableAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asaot.ref.Append("timeouts"))
}

type azurermStreamAnalyticsOutputTableState struct {
	BatchSize              float64        `json:"batch_size"`
	ColumnsToRemove        []string       `json:"columns_to_remove"`
	Id                     string         `json:"id"`
	Name                   string         `json:"name"`
	PartitionKey           string         `json:"partition_key"`
	ResourceGroupName      string         `json:"resource_group_name"`
	RowKey                 string         `json:"row_key"`
	StorageAccountKey      string         `json:"storage_account_key"`
	StorageAccountName     string         `json:"storage_account_name"`
	StreamAnalyticsJobName string         `json:"stream_analytics_job_name"`
	Table                  string         `json:"table"`
	Timeouts               *TimeoutsState `json:"timeouts"`
}
