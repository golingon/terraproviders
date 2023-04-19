// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datasharedatasetkustodatabase "github.com/golingon/terraproviders/azurerm/3.52.0/datasharedatasetkustodatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataShareDatasetKustoDatabase creates a new instance of [DataShareDatasetKustoDatabase].
func NewDataShareDatasetKustoDatabase(name string, args DataShareDatasetKustoDatabaseArgs) *DataShareDatasetKustoDatabase {
	return &DataShareDatasetKustoDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShareDatasetKustoDatabase)(nil)

// DataShareDatasetKustoDatabase represents the Terraform resource azurerm_data_share_dataset_kusto_database.
type DataShareDatasetKustoDatabase struct {
	Name      string
	Args      DataShareDatasetKustoDatabaseArgs
	state     *dataShareDatasetKustoDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) Type() string {
	return "azurerm_data_share_dataset_kusto_database"
}

// LocalName returns the local name for [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) LocalName() string {
	return dsdkd.Name
}

// Configuration returns the configuration (args) for [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) Configuration() interface{} {
	return dsdkd.Args
}

// DependOn is used for other resources to depend on [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(dsdkd)
}

// Dependencies returns the list of resources [DataShareDatasetKustoDatabase] depends_on.
func (dsdkd *DataShareDatasetKustoDatabase) Dependencies() terra.Dependencies {
	return dsdkd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) LifecycleManagement() *terra.Lifecycle {
	return dsdkd.Lifecycle
}

// Attributes returns the attributes for [DataShareDatasetKustoDatabase].
func (dsdkd *DataShareDatasetKustoDatabase) Attributes() dataShareDatasetKustoDatabaseAttributes {
	return dataShareDatasetKustoDatabaseAttributes{ref: terra.ReferenceResource(dsdkd)}
}

// ImportState imports the given attribute values into [DataShareDatasetKustoDatabase]'s state.
func (dsdkd *DataShareDatasetKustoDatabase) ImportState(av io.Reader) error {
	dsdkd.state = &dataShareDatasetKustoDatabaseState{}
	if err := json.NewDecoder(av).Decode(dsdkd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsdkd.Type(), dsdkd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataShareDatasetKustoDatabase] has state.
func (dsdkd *DataShareDatasetKustoDatabase) State() (*dataShareDatasetKustoDatabaseState, bool) {
	return dsdkd.state, dsdkd.state != nil
}

// StateMust returns the state for [DataShareDatasetKustoDatabase]. Panics if the state is nil.
func (dsdkd *DataShareDatasetKustoDatabase) StateMust() *dataShareDatasetKustoDatabaseState {
	if dsdkd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsdkd.Type(), dsdkd.LocalName()))
	}
	return dsdkd.state
}

// DataShareDatasetKustoDatabaseArgs contains the configurations for azurerm_data_share_dataset_kusto_database.
type DataShareDatasetKustoDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KustoDatabaseId: string, required
	KustoDatabaseId terra.StringValue `hcl:"kusto_database_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareId: string, required
	ShareId terra.StringValue `hcl:"share_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasharedatasetkustodatabase.Timeouts `hcl:"timeouts,block"`
}
type dataShareDatasetKustoDatabaseAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("id"))
}

// KustoClusterLocation returns a reference to field kusto_cluster_location of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) KustoClusterLocation() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("kusto_cluster_location"))
}

// KustoDatabaseId returns a reference to field kusto_database_id of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) KustoDatabaseId() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("kusto_database_id"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("name"))
}

// ShareId returns a reference to field share_id of azurerm_data_share_dataset_kusto_database.
func (dsdkd dataShareDatasetKustoDatabaseAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(dsdkd.ref.Append("share_id"))
}

func (dsdkd dataShareDatasetKustoDatabaseAttributes) Timeouts() datasharedatasetkustodatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedatasetkustodatabase.TimeoutsAttributes](dsdkd.ref.Append("timeouts"))
}

type dataShareDatasetKustoDatabaseState struct {
	DisplayName          string                                       `json:"display_name"`
	Id                   string                                       `json:"id"`
	KustoClusterLocation string                                       `json:"kusto_cluster_location"`
	KustoDatabaseId      string                                       `json:"kusto_database_id"`
	Name                 string                                       `json:"name"`
	ShareId              string                                       `json:"share_id"`
	Timeouts             *datasharedatasetkustodatabase.TimeoutsState `json:"timeouts"`
}
