// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	springcloudstorage "github.com/golingon/terraproviders/azurerm/3.98.0/springcloudstorage"
	"io"
)

// NewSpringCloudStorage creates a new instance of [SpringCloudStorage].
func NewSpringCloudStorage(name string, args SpringCloudStorageArgs) *SpringCloudStorage {
	return &SpringCloudStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudStorage)(nil)

// SpringCloudStorage represents the Terraform resource azurerm_spring_cloud_storage.
type SpringCloudStorage struct {
	Name      string
	Args      SpringCloudStorageArgs
	state     *springCloudStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudStorage].
func (scs *SpringCloudStorage) Type() string {
	return "azurerm_spring_cloud_storage"
}

// LocalName returns the local name for [SpringCloudStorage].
func (scs *SpringCloudStorage) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [SpringCloudStorage].
func (scs *SpringCloudStorage) Configuration() interface{} {
	return scs.Args
}

// DependOn is used for other resources to depend on [SpringCloudStorage].
func (scs *SpringCloudStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(scs)
}

// Dependencies returns the list of resources [SpringCloudStorage] depends_on.
func (scs *SpringCloudStorage) Dependencies() terra.Dependencies {
	return scs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudStorage].
func (scs *SpringCloudStorage) LifecycleManagement() *terra.Lifecycle {
	return scs.Lifecycle
}

// Attributes returns the attributes for [SpringCloudStorage].
func (scs *SpringCloudStorage) Attributes() springCloudStorageAttributes {
	return springCloudStorageAttributes{ref: terra.ReferenceResource(scs)}
}

// ImportState imports the given attribute values into [SpringCloudStorage]'s state.
func (scs *SpringCloudStorage) ImportState(av io.Reader) error {
	scs.state = &springCloudStorageState{}
	if err := json.NewDecoder(av).Decode(scs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scs.Type(), scs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudStorage] has state.
func (scs *SpringCloudStorage) State() (*springCloudStorageState, bool) {
	return scs.state, scs.state != nil
}

// StateMust returns the state for [SpringCloudStorage]. Panics if the state is nil.
func (scs *SpringCloudStorage) StateMust() *springCloudStorageState {
	if scs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scs.Type(), scs.LocalName()))
	}
	return scs.state
}

// SpringCloudStorageArgs contains the configurations for azurerm_spring_cloud_storage.
type SpringCloudStorageArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// StorageAccountKey: string, required
	StorageAccountKey terra.StringValue `hcl:"storage_account_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudstorage.Timeouts `hcl:"timeouts,block"`
}
type springCloudStorageAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_storage.
func (scs springCloudStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_storage.
func (scs springCloudStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_storage.
func (scs springCloudStorageAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("spring_cloud_service_id"))
}

// StorageAccountKey returns a reference to field storage_account_key of azurerm_spring_cloud_storage.
func (scs springCloudStorageAttributes) StorageAccountKey() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("storage_account_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_spring_cloud_storage.
func (scs springCloudStorageAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("storage_account_name"))
}

func (scs springCloudStorageAttributes) Timeouts() springcloudstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudstorage.TimeoutsAttributes](scs.ref.Append("timeouts"))
}

type springCloudStorageState struct {
	Id                   string                            `json:"id"`
	Name                 string                            `json:"name"`
	SpringCloudServiceId string                            `json:"spring_cloud_service_id"`
	StorageAccountKey    string                            `json:"storage_account_key"`
	StorageAccountName   string                            `json:"storage_account_name"`
	Timeouts             *springcloudstorage.TimeoutsState `json:"timeouts"`
}
