// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogentrygroup "github.com/golingon/terraproviders/googlebeta/4.64.0/datacatalogentrygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogEntryGroup creates a new instance of [DataCatalogEntryGroup].
func NewDataCatalogEntryGroup(name string, args DataCatalogEntryGroupArgs) *DataCatalogEntryGroup {
	return &DataCatalogEntryGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogEntryGroup)(nil)

// DataCatalogEntryGroup represents the Terraform resource google_data_catalog_entry_group.
type DataCatalogEntryGroup struct {
	Name      string
	Args      DataCatalogEntryGroupArgs
	state     *dataCatalogEntryGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) Type() string {
	return "google_data_catalog_entry_group"
}

// LocalName returns the local name for [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) LocalName() string {
	return dceg.Name
}

// Configuration returns the configuration (args) for [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) Configuration() interface{} {
	return dceg.Args
}

// DependOn is used for other resources to depend on [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dceg)
}

// Dependencies returns the list of resources [DataCatalogEntryGroup] depends_on.
func (dceg *DataCatalogEntryGroup) Dependencies() terra.Dependencies {
	return dceg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) LifecycleManagement() *terra.Lifecycle {
	return dceg.Lifecycle
}

// Attributes returns the attributes for [DataCatalogEntryGroup].
func (dceg *DataCatalogEntryGroup) Attributes() dataCatalogEntryGroupAttributes {
	return dataCatalogEntryGroupAttributes{ref: terra.ReferenceResource(dceg)}
}

// ImportState imports the given attribute values into [DataCatalogEntryGroup]'s state.
func (dceg *DataCatalogEntryGroup) ImportState(av io.Reader) error {
	dceg.state = &dataCatalogEntryGroupState{}
	if err := json.NewDecoder(av).Decode(dceg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dceg.Type(), dceg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogEntryGroup] has state.
func (dceg *DataCatalogEntryGroup) State() (*dataCatalogEntryGroupState, bool) {
	return dceg.state, dceg.state != nil
}

// StateMust returns the state for [DataCatalogEntryGroup]. Panics if the state is nil.
func (dceg *DataCatalogEntryGroup) StateMust() *dataCatalogEntryGroupState {
	if dceg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dceg.Type(), dceg.LocalName()))
	}
	return dceg.state
}

// DataCatalogEntryGroupArgs contains the configurations for google_data_catalog_entry_group.
type DataCatalogEntryGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EntryGroupId: string, required
	EntryGroupId terra.StringValue `hcl:"entry_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *datacatalogentrygroup.Timeouts `hcl:"timeouts,block"`
}
type dataCatalogEntryGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("display_name"))
}

// EntryGroupId returns a reference to field entry_group_id of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) EntryGroupId() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("entry_group_id"))
}

// Id returns a reference to field id of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("id"))
}

// Name returns a reference to field name of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("name"))
}

// Project returns a reference to field project of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_entry_group.
func (dceg dataCatalogEntryGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dceg.ref.Append("region"))
}

func (dceg dataCatalogEntryGroupAttributes) Timeouts() datacatalogentrygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacatalogentrygroup.TimeoutsAttributes](dceg.ref.Append("timeouts"))
}

type dataCatalogEntryGroupState struct {
	Description  string                               `json:"description"`
	DisplayName  string                               `json:"display_name"`
	EntryGroupId string                               `json:"entry_group_id"`
	Id           string                               `json:"id"`
	Name         string                               `json:"name"`
	Project      string                               `json:"project"`
	Region       string                               `json:"region"`
	Timeouts     *datacatalogentrygroup.TimeoutsState `json:"timeouts"`
}
