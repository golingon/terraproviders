// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigquerytable "github.com/golingon/terraproviders/googlebeta/4.75.1/bigquerytable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryTable creates a new instance of [BigqueryTable].
func NewBigqueryTable(name string, args BigqueryTableArgs) *BigqueryTable {
	return &BigqueryTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryTable)(nil)

// BigqueryTable represents the Terraform resource google_bigquery_table.
type BigqueryTable struct {
	Name      string
	Args      BigqueryTableArgs
	state     *bigqueryTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryTable].
func (bt *BigqueryTable) Type() string {
	return "google_bigquery_table"
}

// LocalName returns the local name for [BigqueryTable].
func (bt *BigqueryTable) LocalName() string {
	return bt.Name
}

// Configuration returns the configuration (args) for [BigqueryTable].
func (bt *BigqueryTable) Configuration() interface{} {
	return bt.Args
}

// DependOn is used for other resources to depend on [BigqueryTable].
func (bt *BigqueryTable) DependOn() terra.Reference {
	return terra.ReferenceResource(bt)
}

// Dependencies returns the list of resources [BigqueryTable] depends_on.
func (bt *BigqueryTable) Dependencies() terra.Dependencies {
	return bt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryTable].
func (bt *BigqueryTable) LifecycleManagement() *terra.Lifecycle {
	return bt.Lifecycle
}

// Attributes returns the attributes for [BigqueryTable].
func (bt *BigqueryTable) Attributes() bigqueryTableAttributes {
	return bigqueryTableAttributes{ref: terra.ReferenceResource(bt)}
}

// ImportState imports the given attribute values into [BigqueryTable]'s state.
func (bt *BigqueryTable) ImportState(av io.Reader) error {
	bt.state = &bigqueryTableState{}
	if err := json.NewDecoder(av).Decode(bt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bt.Type(), bt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryTable] has state.
func (bt *BigqueryTable) State() (*bigqueryTableState, bool) {
	return bt.state, bt.state != nil
}

// StateMust returns the state for [BigqueryTable]. Panics if the state is nil.
func (bt *BigqueryTable) StateMust() *bigqueryTableState {
	if bt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bt.Type(), bt.LocalName()))
	}
	return bt.state
}

// BigqueryTableArgs contains the configurations for google_bigquery_table.
type BigqueryTableArgs struct {
	// Clustering: list of string, optional
	Clustering terra.ListValue[terra.StringValue] `hcl:"clustering,attr"`
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExpirationTime: number, optional
	ExpirationTime terra.NumberValue `hcl:"expiration_time,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
	// EncryptionConfiguration: optional
	EncryptionConfiguration *bigquerytable.EncryptionConfiguration `hcl:"encryption_configuration,block"`
	// ExternalDataConfiguration: optional
	ExternalDataConfiguration *bigquerytable.ExternalDataConfiguration `hcl:"external_data_configuration,block"`
	// MaterializedView: optional
	MaterializedView *bigquerytable.MaterializedView `hcl:"materialized_view,block"`
	// RangePartitioning: optional
	RangePartitioning *bigquerytable.RangePartitioning `hcl:"range_partitioning,block"`
	// TimePartitioning: optional
	TimePartitioning *bigquerytable.TimePartitioning `hcl:"time_partitioning,block"`
	// View: optional
	View *bigquerytable.View `hcl:"view,block"`
}
type bigqueryTableAttributes struct {
	ref terra.Reference
}

// Clustering returns a reference to field clustering of google_bigquery_table.
func (bt bigqueryTableAttributes) Clustering() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bt.ref.Append("clustering"))
}

// CreationTime returns a reference to field creation_time of google_bigquery_table.
func (bt bigqueryTableAttributes) CreationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("creation_time"))
}

// DatasetId returns a reference to field dataset_id of google_bigquery_table.
func (bt bigqueryTableAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("dataset_id"))
}

// DeletionProtection returns a reference to field deletion_protection of google_bigquery_table.
func (bt bigqueryTableAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(bt.ref.Append("deletion_protection"))
}

// Description returns a reference to field description of google_bigquery_table.
func (bt bigqueryTableAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("description"))
}

// Etag returns a reference to field etag of google_bigquery_table.
func (bt bigqueryTableAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("etag"))
}

// ExpirationTime returns a reference to field expiration_time of google_bigquery_table.
func (bt bigqueryTableAttributes) ExpirationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("expiration_time"))
}

// FriendlyName returns a reference to field friendly_name of google_bigquery_table.
func (bt bigqueryTableAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("friendly_name"))
}

// Id returns a reference to field id of google_bigquery_table.
func (bt bigqueryTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("id"))
}

// Labels returns a reference to field labels of google_bigquery_table.
func (bt bigqueryTableAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bt.ref.Append("labels"))
}

// LastModifiedTime returns a reference to field last_modified_time of google_bigquery_table.
func (bt bigqueryTableAttributes) LastModifiedTime() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("last_modified_time"))
}

// Location returns a reference to field location of google_bigquery_table.
func (bt bigqueryTableAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("location"))
}

// NumBytes returns a reference to field num_bytes of google_bigquery_table.
func (bt bigqueryTableAttributes) NumBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("num_bytes"))
}

// NumLongTermBytes returns a reference to field num_long_term_bytes of google_bigquery_table.
func (bt bigqueryTableAttributes) NumLongTermBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("num_long_term_bytes"))
}

// NumRows returns a reference to field num_rows of google_bigquery_table.
func (bt bigqueryTableAttributes) NumRows() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("num_rows"))
}

// Project returns a reference to field project of google_bigquery_table.
func (bt bigqueryTableAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("project"))
}

// Schema returns a reference to field schema of google_bigquery_table.
func (bt bigqueryTableAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("schema"))
}

// SelfLink returns a reference to field self_link of google_bigquery_table.
func (bt bigqueryTableAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("self_link"))
}

// TableId returns a reference to field table_id of google_bigquery_table.
func (bt bigqueryTableAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("table_id"))
}

// Type returns a reference to field type of google_bigquery_table.
func (bt bigqueryTableAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bt.ref.Append("type"))
}

func (bt bigqueryTableAttributes) EncryptionConfiguration() terra.ListValue[bigquerytable.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[bigquerytable.EncryptionConfigurationAttributes](bt.ref.Append("encryption_configuration"))
}

func (bt bigqueryTableAttributes) ExternalDataConfiguration() terra.ListValue[bigquerytable.ExternalDataConfigurationAttributes] {
	return terra.ReferenceAsList[bigquerytable.ExternalDataConfigurationAttributes](bt.ref.Append("external_data_configuration"))
}

func (bt bigqueryTableAttributes) MaterializedView() terra.ListValue[bigquerytable.MaterializedViewAttributes] {
	return terra.ReferenceAsList[bigquerytable.MaterializedViewAttributes](bt.ref.Append("materialized_view"))
}

func (bt bigqueryTableAttributes) RangePartitioning() terra.ListValue[bigquerytable.RangePartitioningAttributes] {
	return terra.ReferenceAsList[bigquerytable.RangePartitioningAttributes](bt.ref.Append("range_partitioning"))
}

func (bt bigqueryTableAttributes) TimePartitioning() terra.ListValue[bigquerytable.TimePartitioningAttributes] {
	return terra.ReferenceAsList[bigquerytable.TimePartitioningAttributes](bt.ref.Append("time_partitioning"))
}

func (bt bigqueryTableAttributes) View() terra.ListValue[bigquerytable.ViewAttributes] {
	return terra.ReferenceAsList[bigquerytable.ViewAttributes](bt.ref.Append("view"))
}

type bigqueryTableState struct {
	Clustering                []string                                       `json:"clustering"`
	CreationTime              float64                                        `json:"creation_time"`
	DatasetId                 string                                         `json:"dataset_id"`
	DeletionProtection        bool                                           `json:"deletion_protection"`
	Description               string                                         `json:"description"`
	Etag                      string                                         `json:"etag"`
	ExpirationTime            float64                                        `json:"expiration_time"`
	FriendlyName              string                                         `json:"friendly_name"`
	Id                        string                                         `json:"id"`
	Labels                    map[string]string                              `json:"labels"`
	LastModifiedTime          float64                                        `json:"last_modified_time"`
	Location                  string                                         `json:"location"`
	NumBytes                  float64                                        `json:"num_bytes"`
	NumLongTermBytes          float64                                        `json:"num_long_term_bytes"`
	NumRows                   float64                                        `json:"num_rows"`
	Project                   string                                         `json:"project"`
	Schema                    string                                         `json:"schema"`
	SelfLink                  string                                         `json:"self_link"`
	TableId                   string                                         `json:"table_id"`
	Type                      string                                         `json:"type"`
	EncryptionConfiguration   []bigquerytable.EncryptionConfigurationState   `json:"encryption_configuration"`
	ExternalDataConfiguration []bigquerytable.ExternalDataConfigurationState `json:"external_data_configuration"`
	MaterializedView          []bigquerytable.MaterializedViewState          `json:"materialized_view"`
	RangePartitioning         []bigquerytable.RangePartitioningState         `json:"range_partitioning"`
	TimePartitioning          []bigquerytable.TimePartitioningState          `json:"time_partitioning"`
	View                      []bigquerytable.ViewState                      `json:"view"`
}
