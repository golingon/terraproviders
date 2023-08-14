// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadatashare "github.com/golingon/terraproviders/azurerm/3.69.0/datadatashare"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataShare creates a new instance of [DataDataShare].
func NewDataDataShare(name string, args DataDataShareArgs) *DataDataShare {
	return &DataDataShare{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataShare)(nil)

// DataDataShare represents the Terraform data resource azurerm_data_share.
type DataDataShare struct {
	Name string
	Args DataDataShareArgs
}

// DataSource returns the Terraform object type for [DataDataShare].
func (ds *DataDataShare) DataSource() string {
	return "azurerm_data_share"
}

// LocalName returns the local name for [DataDataShare].
func (ds *DataDataShare) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DataDataShare].
func (ds *DataDataShare) Configuration() interface{} {
	return ds.Args
}

// Attributes returns the attributes for [DataDataShare].
func (ds *DataDataShare) Attributes() dataDataShareAttributes {
	return dataDataShareAttributes{ref: terra.ReferenceDataResource(ds)}
}

// DataDataShareArgs contains the configurations for azurerm_data_share.
type DataDataShareArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SnapshotSchedule: min=0
	SnapshotSchedule []datadatashare.SnapshotSchedule `hcl:"snapshot_schedule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadatashare.Timeouts `hcl:"timeouts,block"`
}
type dataDataShareAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of azurerm_data_share.
func (ds dataDataShareAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("account_id"))
}

// Description returns a reference to field description of azurerm_data_share.
func (ds dataDataShareAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_share.
func (ds dataDataShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_data_share.
func (ds dataDataShareAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("kind"))
}

// Name returns a reference to field name of azurerm_data_share.
func (ds dataDataShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("name"))
}

// Terms returns a reference to field terms of azurerm_data_share.
func (ds dataDataShareAttributes) Terms() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("terms"))
}

func (ds dataDataShareAttributes) SnapshotSchedule() terra.ListValue[datadatashare.SnapshotScheduleAttributes] {
	return terra.ReferenceAsList[datadatashare.SnapshotScheduleAttributes](ds.ref.Append("snapshot_schedule"))
}

func (ds dataDataShareAttributes) Timeouts() datadatashare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatashare.TimeoutsAttributes](ds.ref.Append("timeouts"))
}