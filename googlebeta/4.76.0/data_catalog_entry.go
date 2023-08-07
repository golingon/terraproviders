// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogentry "github.com/golingon/terraproviders/googlebeta/4.76.0/datacatalogentry"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogEntry creates a new instance of [DataCatalogEntry].
func NewDataCatalogEntry(name string, args DataCatalogEntryArgs) *DataCatalogEntry {
	return &DataCatalogEntry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogEntry)(nil)

// DataCatalogEntry represents the Terraform resource google_data_catalog_entry.
type DataCatalogEntry struct {
	Name      string
	Args      DataCatalogEntryArgs
	state     *dataCatalogEntryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogEntry].
func (dce *DataCatalogEntry) Type() string {
	return "google_data_catalog_entry"
}

// LocalName returns the local name for [DataCatalogEntry].
func (dce *DataCatalogEntry) LocalName() string {
	return dce.Name
}

// Configuration returns the configuration (args) for [DataCatalogEntry].
func (dce *DataCatalogEntry) Configuration() interface{} {
	return dce.Args
}

// DependOn is used for other resources to depend on [DataCatalogEntry].
func (dce *DataCatalogEntry) DependOn() terra.Reference {
	return terra.ReferenceResource(dce)
}

// Dependencies returns the list of resources [DataCatalogEntry] depends_on.
func (dce *DataCatalogEntry) Dependencies() terra.Dependencies {
	return dce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogEntry].
func (dce *DataCatalogEntry) LifecycleManagement() *terra.Lifecycle {
	return dce.Lifecycle
}

// Attributes returns the attributes for [DataCatalogEntry].
func (dce *DataCatalogEntry) Attributes() dataCatalogEntryAttributes {
	return dataCatalogEntryAttributes{ref: terra.ReferenceResource(dce)}
}

// ImportState imports the given attribute values into [DataCatalogEntry]'s state.
func (dce *DataCatalogEntry) ImportState(av io.Reader) error {
	dce.state = &dataCatalogEntryState{}
	if err := json.NewDecoder(av).Decode(dce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dce.Type(), dce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogEntry] has state.
func (dce *DataCatalogEntry) State() (*dataCatalogEntryState, bool) {
	return dce.state, dce.state != nil
}

// StateMust returns the state for [DataCatalogEntry]. Panics if the state is nil.
func (dce *DataCatalogEntry) StateMust() *dataCatalogEntryState {
	if dce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dce.Type(), dce.LocalName()))
	}
	return dce.state
}

// DataCatalogEntryArgs contains the configurations for google_data_catalog_entry.
type DataCatalogEntryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EntryGroup: string, required
	EntryGroup terra.StringValue `hcl:"entry_group,attr" validate:"required"`
	// EntryId: string, required
	EntryId terra.StringValue `hcl:"entry_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkedResource: string, optional
	LinkedResource terra.StringValue `hcl:"linked_resource,attr"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// UserSpecifiedSystem: string, optional
	UserSpecifiedSystem terra.StringValue `hcl:"user_specified_system,attr"`
	// UserSpecifiedType: string, optional
	UserSpecifiedType terra.StringValue `hcl:"user_specified_type,attr"`
	// BigqueryDateShardedSpec: min=0
	BigqueryDateShardedSpec []datacatalogentry.BigqueryDateShardedSpec `hcl:"bigquery_date_sharded_spec,block" validate:"min=0"`
	// BigqueryTableSpec: min=0
	BigqueryTableSpec []datacatalogentry.BigqueryTableSpec `hcl:"bigquery_table_spec,block" validate:"min=0"`
	// GcsFilesetSpec: optional
	GcsFilesetSpec *datacatalogentry.GcsFilesetSpec `hcl:"gcs_fileset_spec,block"`
	// Timeouts: optional
	Timeouts *datacatalogentry.Timeouts `hcl:"timeouts,block"`
}
type dataCatalogEntryAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("display_name"))
}

// EntryGroup returns a reference to field entry_group of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) EntryGroup() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("entry_group"))
}

// EntryId returns a reference to field entry_id of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) EntryId() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("entry_id"))
}

// Id returns a reference to field id of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("id"))
}

// IntegratedSystem returns a reference to field integrated_system of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) IntegratedSystem() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("integrated_system"))
}

// LinkedResource returns a reference to field linked_resource of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) LinkedResource() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("linked_resource"))
}

// Name returns a reference to field name of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("name"))
}

// Schema returns a reference to field schema of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("schema"))
}

// Type returns a reference to field type of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("type"))
}

// UserSpecifiedSystem returns a reference to field user_specified_system of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) UserSpecifiedSystem() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("user_specified_system"))
}

// UserSpecifiedType returns a reference to field user_specified_type of google_data_catalog_entry.
func (dce dataCatalogEntryAttributes) UserSpecifiedType() terra.StringValue {
	return terra.ReferenceAsString(dce.ref.Append("user_specified_type"))
}

func (dce dataCatalogEntryAttributes) BigqueryDateShardedSpec() terra.ListValue[datacatalogentry.BigqueryDateShardedSpecAttributes] {
	return terra.ReferenceAsList[datacatalogentry.BigqueryDateShardedSpecAttributes](dce.ref.Append("bigquery_date_sharded_spec"))
}

func (dce dataCatalogEntryAttributes) BigqueryTableSpec() terra.ListValue[datacatalogentry.BigqueryTableSpecAttributes] {
	return terra.ReferenceAsList[datacatalogentry.BigqueryTableSpecAttributes](dce.ref.Append("bigquery_table_spec"))
}

func (dce dataCatalogEntryAttributes) GcsFilesetSpec() terra.ListValue[datacatalogentry.GcsFilesetSpecAttributes] {
	return terra.ReferenceAsList[datacatalogentry.GcsFilesetSpecAttributes](dce.ref.Append("gcs_fileset_spec"))
}

func (dce dataCatalogEntryAttributes) Timeouts() datacatalogentry.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacatalogentry.TimeoutsAttributes](dce.ref.Append("timeouts"))
}

type dataCatalogEntryState struct {
	Description             string                                          `json:"description"`
	DisplayName             string                                          `json:"display_name"`
	EntryGroup              string                                          `json:"entry_group"`
	EntryId                 string                                          `json:"entry_id"`
	Id                      string                                          `json:"id"`
	IntegratedSystem        string                                          `json:"integrated_system"`
	LinkedResource          string                                          `json:"linked_resource"`
	Name                    string                                          `json:"name"`
	Schema                  string                                          `json:"schema"`
	Type                    string                                          `json:"type"`
	UserSpecifiedSystem     string                                          `json:"user_specified_system"`
	UserSpecifiedType       string                                          `json:"user_specified_type"`
	BigqueryDateShardedSpec []datacatalogentry.BigqueryDateShardedSpecState `json:"bigquery_date_sharded_spec"`
	BigqueryTableSpec       []datacatalogentry.BigqueryTableSpecState       `json:"bigquery_table_spec"`
	GcsFilesetSpec          []datacatalogentry.GcsFilesetSpecState          `json:"gcs_fileset_spec"`
	Timeouts                *datacatalogentry.TimeoutsState                 `json:"timeouts"`
}
