// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskencryptionset "github.com/golingon/terraproviders/azurerm/3.52.0/diskencryptionset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskEncryptionSet creates a new instance of [DiskEncryptionSet].
func NewDiskEncryptionSet(name string, args DiskEncryptionSetArgs) *DiskEncryptionSet {
	return &DiskEncryptionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskEncryptionSet)(nil)

// DiskEncryptionSet represents the Terraform resource azurerm_disk_encryption_set.
type DiskEncryptionSet struct {
	Name      string
	Args      DiskEncryptionSetArgs
	state     *diskEncryptionSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskEncryptionSet].
func (des *DiskEncryptionSet) Type() string {
	return "azurerm_disk_encryption_set"
}

// LocalName returns the local name for [DiskEncryptionSet].
func (des *DiskEncryptionSet) LocalName() string {
	return des.Name
}

// Configuration returns the configuration (args) for [DiskEncryptionSet].
func (des *DiskEncryptionSet) Configuration() interface{} {
	return des.Args
}

// DependOn is used for other resources to depend on [DiskEncryptionSet].
func (des *DiskEncryptionSet) DependOn() terra.Reference {
	return terra.ReferenceResource(des)
}

// Dependencies returns the list of resources [DiskEncryptionSet] depends_on.
func (des *DiskEncryptionSet) Dependencies() terra.Dependencies {
	return des.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskEncryptionSet].
func (des *DiskEncryptionSet) LifecycleManagement() *terra.Lifecycle {
	return des.Lifecycle
}

// Attributes returns the attributes for [DiskEncryptionSet].
func (des *DiskEncryptionSet) Attributes() diskEncryptionSetAttributes {
	return diskEncryptionSetAttributes{ref: terra.ReferenceResource(des)}
}

// ImportState imports the given attribute values into [DiskEncryptionSet]'s state.
func (des *DiskEncryptionSet) ImportState(av io.Reader) error {
	des.state = &diskEncryptionSetState{}
	if err := json.NewDecoder(av).Decode(des.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", des.Type(), des.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskEncryptionSet] has state.
func (des *DiskEncryptionSet) State() (*diskEncryptionSetState, bool) {
	return des.state, des.state != nil
}

// StateMust returns the state for [DiskEncryptionSet]. Panics if the state is nil.
func (des *DiskEncryptionSet) StateMust() *diskEncryptionSetState {
	if des.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", des.Type(), des.LocalName()))
	}
	return des.state
}

// DiskEncryptionSetArgs contains the configurations for azurerm_disk_encryption_set.
type DiskEncryptionSetArgs struct {
	// AutoKeyRotationEnabled: bool, optional
	AutoKeyRotationEnabled terra.BoolValue `hcl:"auto_key_rotation_enabled,attr"`
	// EncryptionType: string, optional
	EncryptionType terra.StringValue `hcl:"encryption_type,attr"`
	// FederatedClientId: string, optional
	FederatedClientId terra.StringValue `hcl:"federated_client_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: required
	Identity *diskencryptionset.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *diskencryptionset.Timeouts `hcl:"timeouts,block"`
}
type diskEncryptionSetAttributes struct {
	ref terra.Reference
}

// AutoKeyRotationEnabled returns a reference to field auto_key_rotation_enabled of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) AutoKeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(des.ref.Append("auto_key_rotation_enabled"))
}

// EncryptionType returns a reference to field encryption_type of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("encryption_type"))
}

// FederatedClientId returns a reference to field federated_client_id of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) FederatedClientId() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("federated_client_id"))
}

// Id returns a reference to field id of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("key_vault_key_id"))
}

// Location returns a reference to field location of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_disk_encryption_set.
func (des diskEncryptionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags"))
}

func (des diskEncryptionSetAttributes) Identity() terra.ListValue[diskencryptionset.IdentityAttributes] {
	return terra.ReferenceAsList[diskencryptionset.IdentityAttributes](des.ref.Append("identity"))
}

func (des diskEncryptionSetAttributes) Timeouts() diskencryptionset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskencryptionset.TimeoutsAttributes](des.ref.Append("timeouts"))
}

type diskEncryptionSetState struct {
	AutoKeyRotationEnabled bool                              `json:"auto_key_rotation_enabled"`
	EncryptionType         string                            `json:"encryption_type"`
	FederatedClientId      string                            `json:"federated_client_id"`
	Id                     string                            `json:"id"`
	KeyVaultKeyId          string                            `json:"key_vault_key_id"`
	Location               string                            `json:"location"`
	Name                   string                            `json:"name"`
	ResourceGroupName      string                            `json:"resource_group_name"`
	Tags                   map[string]string                 `json:"tags"`
	Identity               []diskencryptionset.IdentityState `json:"identity"`
	Timeouts               *diskencryptionset.TimeoutsState  `json:"timeouts"`
}
