// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageshare "github.com/golingon/terraproviders/azurerm/3.66.0/storageshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageShare creates a new instance of [StorageShare].
func NewStorageShare(name string, args StorageShareArgs) *StorageShare {
	return &StorageShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageShare)(nil)

// StorageShare represents the Terraform resource azurerm_storage_share.
type StorageShare struct {
	Name      string
	Args      StorageShareArgs
	state     *storageShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageShare].
func (ss *StorageShare) Type() string {
	return "azurerm_storage_share"
}

// LocalName returns the local name for [StorageShare].
func (ss *StorageShare) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [StorageShare].
func (ss *StorageShare) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [StorageShare].
func (ss *StorageShare) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [StorageShare] depends_on.
func (ss *StorageShare) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageShare].
func (ss *StorageShare) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [StorageShare].
func (ss *StorageShare) Attributes() storageShareAttributes {
	return storageShareAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [StorageShare]'s state.
func (ss *StorageShare) ImportState(av io.Reader) error {
	ss.state = &storageShareState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageShare] has state.
func (ss *StorageShare) State() (*storageShareState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [StorageShare]. Panics if the state is nil.
func (ss *StorageShare) StateMust() *storageShareState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// StorageShareArgs contains the configurations for azurerm_storage_share.
type StorageShareArgs struct {
	// AccessTier: string, optional
	AccessTier terra.StringValue `hcl:"access_tier,attr"`
	// EnabledProtocol: string, optional
	EnabledProtocol terra.StringValue `hcl:"enabled_protocol,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Quota: number, required
	Quota terra.NumberValue `hcl:"quota,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Acl: min=0
	Acl []storageshare.Acl `hcl:"acl,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storageshare.Timeouts `hcl:"timeouts,block"`
}
type storageShareAttributes struct {
	ref terra.Reference
}

// AccessTier returns a reference to field access_tier of azurerm_storage_share.
func (ss storageShareAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("access_tier"))
}

// EnabledProtocol returns a reference to field enabled_protocol of azurerm_storage_share.
func (ss storageShareAttributes) EnabledProtocol() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("enabled_protocol"))
}

// Id returns a reference to field id of azurerm_storage_share.
func (ss storageShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_share.
func (ss storageShareAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_share.
func (ss storageShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// Quota returns a reference to field quota of azurerm_storage_share.
func (ss storageShareAttributes) Quota() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("quota"))
}

// ResourceManagerId returns a reference to field resource_manager_id of azurerm_storage_share.
func (ss storageShareAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_manager_id"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_share.
func (ss storageShareAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("storage_account_name"))
}

// Url returns a reference to field url of azurerm_storage_share.
func (ss storageShareAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("url"))
}

func (ss storageShareAttributes) Acl() terra.SetValue[storageshare.AclAttributes] {
	return terra.ReferenceAsSet[storageshare.AclAttributes](ss.ref.Append("acl"))
}

func (ss storageShareAttributes) Timeouts() storageshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageshare.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type storageShareState struct {
	AccessTier         string                      `json:"access_tier"`
	EnabledProtocol    string                      `json:"enabled_protocol"`
	Id                 string                      `json:"id"`
	Metadata           map[string]string           `json:"metadata"`
	Name               string                      `json:"name"`
	Quota              float64                     `json:"quota"`
	ResourceManagerId  string                      `json:"resource_manager_id"`
	StorageAccountName string                      `json:"storage_account_name"`
	Url                string                      `json:"url"`
	Acl                []storageshare.AclState     `json:"acl"`
	Timeouts           *storageshare.TimeoutsState `json:"timeouts"`
}
