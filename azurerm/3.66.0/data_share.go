// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datashare "github.com/golingon/terraproviders/azurerm/3.66.0/datashare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataShare creates a new instance of [DataShare].
func NewDataShare(name string, args DataShareArgs) *DataShare {
	return &DataShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShare)(nil)

// DataShare represents the Terraform resource azurerm_data_share.
type DataShare struct {
	Name      string
	Args      DataShareArgs
	state     *dataShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataShare].
func (ds *DataShare) Type() string {
	return "azurerm_data_share"
}

// LocalName returns the local name for [DataShare].
func (ds *DataShare) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DataShare].
func (ds *DataShare) Configuration() interface{} {
	return ds.Args
}

// DependOn is used for other resources to depend on [DataShare].
func (ds *DataShare) DependOn() terra.Reference {
	return terra.ReferenceResource(ds)
}

// Dependencies returns the list of resources [DataShare] depends_on.
func (ds *DataShare) Dependencies() terra.Dependencies {
	return ds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataShare].
func (ds *DataShare) LifecycleManagement() *terra.Lifecycle {
	return ds.Lifecycle
}

// Attributes returns the attributes for [DataShare].
func (ds *DataShare) Attributes() dataShareAttributes {
	return dataShareAttributes{ref: terra.ReferenceResource(ds)}
}

// ImportState imports the given attribute values into [DataShare]'s state.
func (ds *DataShare) ImportState(av io.Reader) error {
	ds.state = &dataShareState{}
	if err := json.NewDecoder(av).Decode(ds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ds.Type(), ds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataShare] has state.
func (ds *DataShare) State() (*dataShareState, bool) {
	return ds.state, ds.state != nil
}

// StateMust returns the state for [DataShare]. Panics if the state is nil.
func (ds *DataShare) StateMust() *dataShareState {
	if ds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ds.Type(), ds.LocalName()))
	}
	return ds.state
}

// DataShareArgs contains the configurations for azurerm_data_share.
type DataShareArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Terms: string, optional
	Terms terra.StringValue `hcl:"terms,attr"`
	// SnapshotSchedule: optional
	SnapshotSchedule *datashare.SnapshotSchedule `hcl:"snapshot_schedule,block"`
	// Timeouts: optional
	Timeouts *datashare.Timeouts `hcl:"timeouts,block"`
}
type dataShareAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of azurerm_data_share.
func (ds dataShareAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("account_id"))
}

// Description returns a reference to field description of azurerm_data_share.
func (ds dataShareAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_data_share.
func (ds dataShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_data_share.
func (ds dataShareAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("kind"))
}

// Name returns a reference to field name of azurerm_data_share.
func (ds dataShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("name"))
}

// Terms returns a reference to field terms of azurerm_data_share.
func (ds dataShareAttributes) Terms() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("terms"))
}

func (ds dataShareAttributes) SnapshotSchedule() terra.ListValue[datashare.SnapshotScheduleAttributes] {
	return terra.ReferenceAsList[datashare.SnapshotScheduleAttributes](ds.ref.Append("snapshot_schedule"))
}

func (ds dataShareAttributes) Timeouts() datashare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datashare.TimeoutsAttributes](ds.ref.Append("timeouts"))
}

type dataShareState struct {
	AccountId        string                            `json:"account_id"`
	Description      string                            `json:"description"`
	Id               string                            `json:"id"`
	Kind             string                            `json:"kind"`
	Name             string                            `json:"name"`
	Terms            string                            `json:"terms"`
	SnapshotSchedule []datashare.SnapshotScheduleState `json:"snapshot_schedule"`
	Timeouts         *datashare.TimeoutsState          `json:"timeouts"`
}
