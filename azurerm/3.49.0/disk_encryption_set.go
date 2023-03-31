// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskencryptionset "github.com/golingon/terraproviders/azurerm/3.49.0/diskencryptionset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDiskEncryptionSet(name string, args DiskEncryptionSetArgs) *DiskEncryptionSet {
	return &DiskEncryptionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskEncryptionSet)(nil)

type DiskEncryptionSet struct {
	Name  string
	Args  DiskEncryptionSetArgs
	state *diskEncryptionSetState
}

func (des *DiskEncryptionSet) Type() string {
	return "azurerm_disk_encryption_set"
}

func (des *DiskEncryptionSet) LocalName() string {
	return des.Name
}

func (des *DiskEncryptionSet) Configuration() interface{} {
	return des.Args
}

func (des *DiskEncryptionSet) Attributes() diskEncryptionSetAttributes {
	return diskEncryptionSetAttributes{ref: terra.ReferenceResource(des)}
}

func (des *DiskEncryptionSet) ImportState(av io.Reader) error {
	des.state = &diskEncryptionSetState{}
	if err := json.NewDecoder(av).Decode(des.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", des.Type(), des.LocalName(), err)
	}
	return nil
}

func (des *DiskEncryptionSet) State() (*diskEncryptionSetState, bool) {
	return des.state, des.state != nil
}

func (des *DiskEncryptionSet) StateMust() *diskEncryptionSetState {
	if des.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", des.Type(), des.LocalName()))
	}
	return des.state
}

func (des *DiskEncryptionSet) DependOn() terra.Reference {
	return terra.ReferenceResource(des)
}

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
	// DependsOn contains resources that DiskEncryptionSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type diskEncryptionSetAttributes struct {
	ref terra.Reference
}

func (des diskEncryptionSetAttributes) AutoKeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceBool(des.ref.Append("auto_key_rotation_enabled"))
}

func (des diskEncryptionSetAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("encryption_type"))
}

func (des diskEncryptionSetAttributes) FederatedClientId() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("federated_client_id"))
}

func (des diskEncryptionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("id"))
}

func (des diskEncryptionSetAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("key_vault_key_id"))
}

func (des diskEncryptionSetAttributes) Location() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("location"))
}

func (des diskEncryptionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("name"))
}

func (des diskEncryptionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(des.ref.Append("resource_group_name"))
}

func (des diskEncryptionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](des.ref.Append("tags"))
}

func (des diskEncryptionSetAttributes) Identity() terra.ListValue[diskencryptionset.IdentityAttributes] {
	return terra.ReferenceList[diskencryptionset.IdentityAttributes](des.ref.Append("identity"))
}

func (des diskEncryptionSetAttributes) Timeouts() diskencryptionset.TimeoutsAttributes {
	return terra.ReferenceSingle[diskencryptionset.TimeoutsAttributes](des.ref.Append("timeouts"))
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
