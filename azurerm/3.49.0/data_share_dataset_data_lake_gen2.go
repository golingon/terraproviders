// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datasharedatasetdatalakegen2 "github.com/golingon/terraproviders/azurerm/3.49.0/datasharedatasetdatalakegen2"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataShareDatasetDataLakeGen2(name string, args DataShareDatasetDataLakeGen2Args) *DataShareDatasetDataLakeGen2 {
	return &DataShareDatasetDataLakeGen2{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShareDatasetDataLakeGen2)(nil)

type DataShareDatasetDataLakeGen2 struct {
	Name  string
	Args  DataShareDatasetDataLakeGen2Args
	state *dataShareDatasetDataLakeGen2State
}

func (dsddlg *DataShareDatasetDataLakeGen2) Type() string {
	return "azurerm_data_share_dataset_data_lake_gen2"
}

func (dsddlg *DataShareDatasetDataLakeGen2) LocalName() string {
	return dsddlg.Name
}

func (dsddlg *DataShareDatasetDataLakeGen2) Configuration() interface{} {
	return dsddlg.Args
}

func (dsddlg *DataShareDatasetDataLakeGen2) Attributes() dataShareDatasetDataLakeGen2Attributes {
	return dataShareDatasetDataLakeGen2Attributes{ref: terra.ReferenceResource(dsddlg)}
}

func (dsddlg *DataShareDatasetDataLakeGen2) ImportState(av io.Reader) error {
	dsddlg.state = &dataShareDatasetDataLakeGen2State{}
	if err := json.NewDecoder(av).Decode(dsddlg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsddlg.Type(), dsddlg.LocalName(), err)
	}
	return nil
}

func (dsddlg *DataShareDatasetDataLakeGen2) State() (*dataShareDatasetDataLakeGen2State, bool) {
	return dsddlg.state, dsddlg.state != nil
}

func (dsddlg *DataShareDatasetDataLakeGen2) StateMust() *dataShareDatasetDataLakeGen2State {
	if dsddlg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsddlg.Type(), dsddlg.LocalName()))
	}
	return dsddlg.state
}

func (dsddlg *DataShareDatasetDataLakeGen2) DependOn() terra.Reference {
	return terra.ReferenceResource(dsddlg)
}

type DataShareDatasetDataLakeGen2Args struct {
	// FilePath: string, optional
	FilePath terra.StringValue `hcl:"file_path,attr"`
	// FileSystemName: string, required
	FileSystemName terra.StringValue `hcl:"file_system_name,attr" validate:"required"`
	// FolderPath: string, optional
	FolderPath terra.StringValue `hcl:"folder_path,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareId: string, required
	ShareId terra.StringValue `hcl:"share_id,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasharedatasetdatalakegen2.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DataShareDatasetDataLakeGen2 depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataShareDatasetDataLakeGen2Attributes struct {
	ref terra.Reference
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("display_name"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) FilePath() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("file_path"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) FileSystemName() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("file_system_name"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) FolderPath() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("folder_path"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) Id() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("id"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) Name() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("name"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) ShareId() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("share_id"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceString(dsddlg.ref.Append("storage_account_id"))
}

func (dsddlg dataShareDatasetDataLakeGen2Attributes) Timeouts() datasharedatasetdatalakegen2.TimeoutsAttributes {
	return terra.ReferenceSingle[datasharedatasetdatalakegen2.TimeoutsAttributes](dsddlg.ref.Append("timeouts"))
}

type dataShareDatasetDataLakeGen2State struct {
	DisplayName      string                                      `json:"display_name"`
	FilePath         string                                      `json:"file_path"`
	FileSystemName   string                                      `json:"file_system_name"`
	FolderPath       string                                      `json:"folder_path"`
	Id               string                                      `json:"id"`
	Name             string                                      `json:"name"`
	ShareId          string                                      `json:"share_id"`
	StorageAccountId string                                      `json:"storage_account_id"`
	Timeouts         *datasharedatasetdatalakegen2.TimeoutsState `json:"timeouts"`
}
