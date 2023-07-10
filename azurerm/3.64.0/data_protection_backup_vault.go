// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackupvault "github.com/golingon/terraproviders/azurerm/3.64.0/dataprotectionbackupvault"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupVault creates a new instance of [DataProtectionBackupVault].
func NewDataProtectionBackupVault(name string, args DataProtectionBackupVaultArgs) *DataProtectionBackupVault {
	return &DataProtectionBackupVault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupVault)(nil)

// DataProtectionBackupVault represents the Terraform resource azurerm_data_protection_backup_vault.
type DataProtectionBackupVault struct {
	Name      string
	Args      DataProtectionBackupVaultArgs
	state     *dataProtectionBackupVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) Type() string {
	return "azurerm_data_protection_backup_vault"
}

// LocalName returns the local name for [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) LocalName() string {
	return dpbv.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) Configuration() interface{} {
	return dpbv.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbv)
}

// Dependencies returns the list of resources [DataProtectionBackupVault] depends_on.
func (dpbv *DataProtectionBackupVault) Dependencies() terra.Dependencies {
	return dpbv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) LifecycleManagement() *terra.Lifecycle {
	return dpbv.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupVault].
func (dpbv *DataProtectionBackupVault) Attributes() dataProtectionBackupVaultAttributes {
	return dataProtectionBackupVaultAttributes{ref: terra.ReferenceResource(dpbv)}
}

// ImportState imports the given attribute values into [DataProtectionBackupVault]'s state.
func (dpbv *DataProtectionBackupVault) ImportState(av io.Reader) error {
	dpbv.state = &dataProtectionBackupVaultState{}
	if err := json.NewDecoder(av).Decode(dpbv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbv.Type(), dpbv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupVault] has state.
func (dpbv *DataProtectionBackupVault) State() (*dataProtectionBackupVaultState, bool) {
	return dpbv.state, dpbv.state != nil
}

// StateMust returns the state for [DataProtectionBackupVault]. Panics if the state is nil.
func (dpbv *DataProtectionBackupVault) StateMust() *dataProtectionBackupVaultState {
	if dpbv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbv.Type(), dpbv.LocalName()))
	}
	return dpbv.state
}

// DataProtectionBackupVaultArgs contains the configurations for azurerm_data_protection_backup_vault.
type DataProtectionBackupVaultArgs struct {
	// DatastoreType: string, required
	DatastoreType terra.StringValue `hcl:"datastore_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Redundancy: string, required
	Redundancy terra.StringValue `hcl:"redundancy,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *dataprotectionbackupvault.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *dataprotectionbackupvault.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupVaultAttributes struct {
	ref terra.Reference
}

// DatastoreType returns a reference to field datastore_type of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) DatastoreType() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("datastore_type"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("name"))
}

// Redundancy returns a reference to field redundancy of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) Redundancy() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("redundancy"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_protection_backup_vault.
func (dpbv dataProtectionBackupVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpbv.ref.Append("tags"))
}

func (dpbv dataProtectionBackupVaultAttributes) Identity() terra.ListValue[dataprotectionbackupvault.IdentityAttributes] {
	return terra.ReferenceAsList[dataprotectionbackupvault.IdentityAttributes](dpbv.ref.Append("identity"))
}

func (dpbv dataProtectionBackupVaultAttributes) Timeouts() dataprotectionbackupvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackupvault.TimeoutsAttributes](dpbv.ref.Append("timeouts"))
}

type dataProtectionBackupVaultState struct {
	DatastoreType     string                                    `json:"datastore_type"`
	Id                string                                    `json:"id"`
	Location          string                                    `json:"location"`
	Name              string                                    `json:"name"`
	Redundancy        string                                    `json:"redundancy"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	Tags              map[string]string                         `json:"tags"`
	Identity          []dataprotectionbackupvault.IdentityState `json:"identity"`
	Timeouts          *dataprotectionbackupvault.TimeoutsState  `json:"timeouts"`
}
