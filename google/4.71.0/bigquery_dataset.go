// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerydataset "github.com/golingon/terraproviders/google/4.71.0/bigquerydataset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDataset creates a new instance of [BigqueryDataset].
func NewBigqueryDataset(name string, args BigqueryDatasetArgs) *BigqueryDataset {
	return &BigqueryDataset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDataset)(nil)

// BigqueryDataset represents the Terraform resource google_bigquery_dataset.
type BigqueryDataset struct {
	Name      string
	Args      BigqueryDatasetArgs
	state     *bigqueryDatasetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDataset].
func (bd *BigqueryDataset) Type() string {
	return "google_bigquery_dataset"
}

// LocalName returns the local name for [BigqueryDataset].
func (bd *BigqueryDataset) LocalName() string {
	return bd.Name
}

// Configuration returns the configuration (args) for [BigqueryDataset].
func (bd *BigqueryDataset) Configuration() interface{} {
	return bd.Args
}

// DependOn is used for other resources to depend on [BigqueryDataset].
func (bd *BigqueryDataset) DependOn() terra.Reference {
	return terra.ReferenceResource(bd)
}

// Dependencies returns the list of resources [BigqueryDataset] depends_on.
func (bd *BigqueryDataset) Dependencies() terra.Dependencies {
	return bd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDataset].
func (bd *BigqueryDataset) LifecycleManagement() *terra.Lifecycle {
	return bd.Lifecycle
}

// Attributes returns the attributes for [BigqueryDataset].
func (bd *BigqueryDataset) Attributes() bigqueryDatasetAttributes {
	return bigqueryDatasetAttributes{ref: terra.ReferenceResource(bd)}
}

// ImportState imports the given attribute values into [BigqueryDataset]'s state.
func (bd *BigqueryDataset) ImportState(av io.Reader) error {
	bd.state = &bigqueryDatasetState{}
	if err := json.NewDecoder(av).Decode(bd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bd.Type(), bd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDataset] has state.
func (bd *BigqueryDataset) State() (*bigqueryDatasetState, bool) {
	return bd.state, bd.state != nil
}

// StateMust returns the state for [BigqueryDataset]. Panics if the state is nil.
func (bd *BigqueryDataset) StateMust() *bigqueryDatasetState {
	if bd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bd.Type(), bd.LocalName()))
	}
	return bd.state
}

// BigqueryDatasetArgs contains the configurations for google_bigquery_dataset.
type BigqueryDatasetArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// DefaultCollation: string, optional
	DefaultCollation terra.StringValue `hcl:"default_collation,attr"`
	// DefaultPartitionExpirationMs: number, optional
	DefaultPartitionExpirationMs terra.NumberValue `hcl:"default_partition_expiration_ms,attr"`
	// DefaultTableExpirationMs: number, optional
	DefaultTableExpirationMs terra.NumberValue `hcl:"default_table_expiration_ms,attr"`
	// DeleteContentsOnDestroy: bool, optional
	DeleteContentsOnDestroy terra.BoolValue `hcl:"delete_contents_on_destroy,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsCaseInsensitive: bool, optional
	IsCaseInsensitive terra.BoolValue `hcl:"is_case_insensitive,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// MaxTimeTravelHours: string, optional
	MaxTimeTravelHours terra.StringValue `hcl:"max_time_travel_hours,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Access: min=0
	Access []bigquerydataset.Access `hcl:"access,block" validate:"min=0"`
	// DefaultEncryptionConfiguration: optional
	DefaultEncryptionConfiguration *bigquerydataset.DefaultEncryptionConfiguration `hcl:"default_encryption_configuration,block"`
	// Timeouts: optional
	Timeouts *bigquerydataset.Timeouts `hcl:"timeouts,block"`
}
type bigqueryDatasetAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) CreationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(bd.ref.Append("creation_time"))
}

// DatasetId returns a reference to field dataset_id of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("dataset_id"))
}

// DefaultCollation returns a reference to field default_collation of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) DefaultCollation() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("default_collation"))
}

// DefaultPartitionExpirationMs returns a reference to field default_partition_expiration_ms of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) DefaultPartitionExpirationMs() terra.NumberValue {
	return terra.ReferenceAsNumber(bd.ref.Append("default_partition_expiration_ms"))
}

// DefaultTableExpirationMs returns a reference to field default_table_expiration_ms of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) DefaultTableExpirationMs() terra.NumberValue {
	return terra.ReferenceAsNumber(bd.ref.Append("default_table_expiration_ms"))
}

// DeleteContentsOnDestroy returns a reference to field delete_contents_on_destroy of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) DeleteContentsOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(bd.ref.Append("delete_contents_on_destroy"))
}

// Description returns a reference to field description of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("description"))
}

// Etag returns a reference to field etag of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("etag"))
}

// FriendlyName returns a reference to field friendly_name of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("friendly_name"))
}

// Id returns a reference to field id of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("id"))
}

// IsCaseInsensitive returns a reference to field is_case_insensitive of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) IsCaseInsensitive() terra.BoolValue {
	return terra.ReferenceAsBool(bd.ref.Append("is_case_insensitive"))
}

// Labels returns a reference to field labels of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bd.ref.Append("labels"))
}

// LastModifiedTime returns a reference to field last_modified_time of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) LastModifiedTime() terra.NumberValue {
	return terra.ReferenceAsNumber(bd.ref.Append("last_modified_time"))
}

// Location returns a reference to field location of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("location"))
}

// MaxTimeTravelHours returns a reference to field max_time_travel_hours of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) MaxTimeTravelHours() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("max_time_travel_hours"))
}

// Project returns a reference to field project of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_bigquery_dataset.
func (bd bigqueryDatasetAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("self_link"))
}

func (bd bigqueryDatasetAttributes) Access() terra.SetValue[bigquerydataset.AccessAttributes] {
	return terra.ReferenceAsSet[bigquerydataset.AccessAttributes](bd.ref.Append("access"))
}

func (bd bigqueryDatasetAttributes) DefaultEncryptionConfiguration() terra.ListValue[bigquerydataset.DefaultEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[bigquerydataset.DefaultEncryptionConfigurationAttributes](bd.ref.Append("default_encryption_configuration"))
}

func (bd bigqueryDatasetAttributes) Timeouts() bigquerydataset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigquerydataset.TimeoutsAttributes](bd.ref.Append("timeouts"))
}

type bigqueryDatasetState struct {
	CreationTime                   float64                                               `json:"creation_time"`
	DatasetId                      string                                                `json:"dataset_id"`
	DefaultCollation               string                                                `json:"default_collation"`
	DefaultPartitionExpirationMs   float64                                               `json:"default_partition_expiration_ms"`
	DefaultTableExpirationMs       float64                                               `json:"default_table_expiration_ms"`
	DeleteContentsOnDestroy        bool                                                  `json:"delete_contents_on_destroy"`
	Description                    string                                                `json:"description"`
	Etag                           string                                                `json:"etag"`
	FriendlyName                   string                                                `json:"friendly_name"`
	Id                             string                                                `json:"id"`
	IsCaseInsensitive              bool                                                  `json:"is_case_insensitive"`
	Labels                         map[string]string                                     `json:"labels"`
	LastModifiedTime               float64                                               `json:"last_modified_time"`
	Location                       string                                                `json:"location"`
	MaxTimeTravelHours             string                                                `json:"max_time_travel_hours"`
	Project                        string                                                `json:"project"`
	SelfLink                       string                                                `json:"self_link"`
	Access                         []bigquerydataset.AccessState                         `json:"access"`
	DefaultEncryptionConfiguration []bigquerydataset.DefaultEncryptionConfigurationState `json:"default_encryption_configuration"`
	Timeouts                       *bigquerydataset.TimeoutsState                        `json:"timeouts"`
}
