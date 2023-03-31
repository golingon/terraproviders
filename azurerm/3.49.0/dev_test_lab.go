// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestlab "github.com/golingon/terraproviders/azurerm/3.49.0/devtestlab"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDevTestLab(name string, args DevTestLabArgs) *DevTestLab {
	return &DevTestLab{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestLab)(nil)

type DevTestLab struct {
	Name  string
	Args  DevTestLabArgs
	state *devTestLabState
}

func (dtl *DevTestLab) Type() string {
	return "azurerm_dev_test_lab"
}

func (dtl *DevTestLab) LocalName() string {
	return dtl.Name
}

func (dtl *DevTestLab) Configuration() interface{} {
	return dtl.Args
}

func (dtl *DevTestLab) Attributes() devTestLabAttributes {
	return devTestLabAttributes{ref: terra.ReferenceResource(dtl)}
}

func (dtl *DevTestLab) ImportState(av io.Reader) error {
	dtl.state = &devTestLabState{}
	if err := json.NewDecoder(av).Decode(dtl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtl.Type(), dtl.LocalName(), err)
	}
	return nil
}

func (dtl *DevTestLab) State() (*devTestLabState, bool) {
	return dtl.state, dtl.state != nil
}

func (dtl *DevTestLab) StateMust() *devTestLabState {
	if dtl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtl.Type(), dtl.LocalName()))
	}
	return dtl.state
}

func (dtl *DevTestLab) DependOn() terra.Reference {
	return terra.ReferenceResource(dtl)
}

type DevTestLabArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *devtestlab.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DevTestLab depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type devTestLabAttributes struct {
	ref terra.Reference
}

func (dtl devTestLabAttributes) ArtifactsStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("artifacts_storage_account_id"))
}

func (dtl devTestLabAttributes) DefaultPremiumStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("default_premium_storage_account_id"))
}

func (dtl devTestLabAttributes) DefaultStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("default_storage_account_id"))
}

func (dtl devTestLabAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("id"))
}

func (dtl devTestLabAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("key_vault_id"))
}

func (dtl devTestLabAttributes) Location() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("location"))
}

func (dtl devTestLabAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("name"))
}

func (dtl devTestLabAttributes) PremiumDataDiskStorageAccountId() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("premium_data_disk_storage_account_id"))
}

func (dtl devTestLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("resource_group_name"))
}

func (dtl devTestLabAttributes) StorageType() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("storage_type"))
}

func (dtl devTestLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dtl.ref.Append("tags"))
}

func (dtl devTestLabAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceString(dtl.ref.Append("unique_identifier"))
}

func (dtl devTestLabAttributes) Timeouts() devtestlab.TimeoutsAttributes {
	return terra.ReferenceSingle[devtestlab.TimeoutsAttributes](dtl.ref.Append("timeouts"))
}

type devTestLabState struct {
	ArtifactsStorageAccountId       string                    `json:"artifacts_storage_account_id"`
	DefaultPremiumStorageAccountId  string                    `json:"default_premium_storage_account_id"`
	DefaultStorageAccountId         string                    `json:"default_storage_account_id"`
	Id                              string                    `json:"id"`
	KeyVaultId                      string                    `json:"key_vault_id"`
	Location                        string                    `json:"location"`
	Name                            string                    `json:"name"`
	PremiumDataDiskStorageAccountId string                    `json:"premium_data_disk_storage_account_id"`
	ResourceGroupName               string                    `json:"resource_group_name"`
	StorageType                     string                    `json:"storage_type"`
	Tags                            map[string]string         `json:"tags"`
	UniqueIdentifier                string                    `json:"unique_identifier"`
	Timeouts                        *devtestlab.TimeoutsState `json:"timeouts"`
}
