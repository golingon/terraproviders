// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestlab "github.com/golingon/terraproviders/azurerm/3.67.0/devtestlab"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestLab creates a new instance of [DevTestLab].
func NewDevTestLab(name string, args DevTestLabArgs) *DevTestLab {
	return &DevTestLab{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestLab)(nil)

// DevTestLab represents the Terraform resource azurerm_dev_test_lab.
type DevTestLab struct {
	Name      string
	Args      DevTestLabArgs
	state     *devTestLabState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestLab].
func (dtl *DevTestLab) Type() string {
	return "azurerm_dev_test_lab"
}

// LocalName returns the local name for [DevTestLab].
func (dtl *DevTestLab) LocalName() string {
	return dtl.Name
}

// Configuration returns the configuration (args) for [DevTestLab].
func (dtl *DevTestLab) Configuration() interface{} {
	return dtl.Args
}

// DependOn is used for other resources to depend on [DevTestLab].
func (dtl *DevTestLab) DependOn() terra.Reference {
	return terra.ReferenceResource(dtl)
}

// Dependencies returns the list of resources [DevTestLab] depends_on.
func (dtl *DevTestLab) Dependencies() terra.Dependencies {
	return dtl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestLab].
func (dtl *DevTestLab) LifecycleManagement() *terra.Lifecycle {
	return dtl.Lifecycle
}

// Attributes returns the attributes for [DevTestLab].
func (dtl *DevTestLab) Attributes() devTestLabAttributes {
	return devTestLabAttributes{ref: terra.ReferenceResource(dtl)}
}

// ImportState imports the given attribute values into [DevTestLab]'s state.
func (dtl *DevTestLab) ImportState(av io.Reader) error {
	dtl.state = &devTestLabState{}
	if err := json.NewDecoder(av).Decode(dtl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtl.Type(), dtl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestLab] has state.
func (dtl *DevTestLab) State() (*devTestLabState, bool) {
	return dtl.state, dtl.state != nil
}

// StateMust returns the state for [DevTestLab]. Panics if the state is nil.
func (dtl *DevTestLab) StateMust() *devTestLabState {
	if dtl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtl.Type(), dtl.LocalName()))
	}
	return dtl.state
}

// DevTestLabArgs contains the configurations for azurerm_dev_test_lab.
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
}
type devTestLabAttributes struct {
	ref terra.Reference
}

// ArtifactsStorageAccountId returns a reference to field artifacts_storage_account_id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) ArtifactsStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("artifacts_storage_account_id"))
}

// DefaultPremiumStorageAccountId returns a reference to field default_premium_storage_account_id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) DefaultPremiumStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("default_premium_storage_account_id"))
}

// DefaultStorageAccountId returns a reference to field default_storage_account_id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) DefaultStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("default_storage_account_id"))
}

// Id returns a reference to field id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("key_vault_id"))
}

// Location returns a reference to field location of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("name"))
}

// PremiumDataDiskStorageAccountId returns a reference to field premium_data_disk_storage_account_id of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) PremiumDataDiskStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("premium_data_disk_storage_account_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("resource_group_name"))
}

// StorageType returns a reference to field storage_type of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtl.ref.Append("tags"))
}

// UniqueIdentifier returns a reference to field unique_identifier of azurerm_dev_test_lab.
func (dtl devTestLabAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dtl.ref.Append("unique_identifier"))
}

func (dtl devTestLabAttributes) Timeouts() devtestlab.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestlab.TimeoutsAttributes](dtl.ref.Append("timeouts"))
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
