// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerappenvironmentstorage "github.com/golingon/terraproviders/azurerm/3.52.0/containerappenvironmentstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAppEnvironmentStorage creates a new instance of [ContainerAppEnvironmentStorage].
func NewContainerAppEnvironmentStorage(name string, args ContainerAppEnvironmentStorageArgs) *ContainerAppEnvironmentStorage {
	return &ContainerAppEnvironmentStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAppEnvironmentStorage)(nil)

// ContainerAppEnvironmentStorage represents the Terraform resource azurerm_container_app_environment_storage.
type ContainerAppEnvironmentStorage struct {
	Name      string
	Args      ContainerAppEnvironmentStorageArgs
	state     *containerAppEnvironmentStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) Type() string {
	return "azurerm_container_app_environment_storage"
}

// LocalName returns the local name for [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) LocalName() string {
	return caes.Name
}

// Configuration returns the configuration (args) for [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) Configuration() interface{} {
	return caes.Args
}

// DependOn is used for other resources to depend on [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(caes)
}

// Dependencies returns the list of resources [ContainerAppEnvironmentStorage] depends_on.
func (caes *ContainerAppEnvironmentStorage) Dependencies() terra.Dependencies {
	return caes.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) LifecycleManagement() *terra.Lifecycle {
	return caes.Lifecycle
}

// Attributes returns the attributes for [ContainerAppEnvironmentStorage].
func (caes *ContainerAppEnvironmentStorage) Attributes() containerAppEnvironmentStorageAttributes {
	return containerAppEnvironmentStorageAttributes{ref: terra.ReferenceResource(caes)}
}

// ImportState imports the given attribute values into [ContainerAppEnvironmentStorage]'s state.
func (caes *ContainerAppEnvironmentStorage) ImportState(av io.Reader) error {
	caes.state = &containerAppEnvironmentStorageState{}
	if err := json.NewDecoder(av).Decode(caes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caes.Type(), caes.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAppEnvironmentStorage] has state.
func (caes *ContainerAppEnvironmentStorage) State() (*containerAppEnvironmentStorageState, bool) {
	return caes.state, caes.state != nil
}

// StateMust returns the state for [ContainerAppEnvironmentStorage]. Panics if the state is nil.
func (caes *ContainerAppEnvironmentStorage) StateMust() *containerAppEnvironmentStorageState {
	if caes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caes.Type(), caes.LocalName()))
	}
	return caes.state
}

// ContainerAppEnvironmentStorageArgs contains the configurations for azurerm_container_app_environment_storage.
type ContainerAppEnvironmentStorageArgs struct {
	// AccessKey: string, required
	AccessKey terra.StringValue `hcl:"access_key,attr" validate:"required"`
	// AccessMode: string, required
	AccessMode terra.StringValue `hcl:"access_mode,attr" validate:"required"`
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// ContainerAppEnvironmentId: string, required
	ContainerAppEnvironmentId terra.StringValue `hcl:"container_app_environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareName: string, required
	ShareName terra.StringValue `hcl:"share_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *containerappenvironmentstorage.Timeouts `hcl:"timeouts,block"`
}
type containerAppEnvironmentStorageAttributes struct {
	ref terra.Reference
}

// AccessKey returns a reference to field access_key of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("access_key"))
}

// AccessMode returns a reference to field access_mode of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) AccessMode() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("access_mode"))
}

// AccountName returns a reference to field account_name of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("account_name"))
}

// ContainerAppEnvironmentId returns a reference to field container_app_environment_id of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("container_app_environment_id"))
}

// Id returns a reference to field id of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("name"))
}

// ShareName returns a reference to field share_name of azurerm_container_app_environment_storage.
func (caes containerAppEnvironmentStorageAttributes) ShareName() terra.StringValue {
	return terra.ReferenceAsString(caes.ref.Append("share_name"))
}

func (caes containerAppEnvironmentStorageAttributes) Timeouts() containerappenvironmentstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerappenvironmentstorage.TimeoutsAttributes](caes.ref.Append("timeouts"))
}

type containerAppEnvironmentStorageState struct {
	AccessKey                 string                                        `json:"access_key"`
	AccessMode                string                                        `json:"access_mode"`
	AccountName               string                                        `json:"account_name"`
	ContainerAppEnvironmentId string                                        `json:"container_app_environment_id"`
	Id                        string                                        `json:"id"`
	Name                      string                                        `json:"name"`
	ShareName                 string                                        `json:"share_name"`
	Timeouts                  *containerappenvironmentstorage.TimeoutsState `json:"timeouts"`
}
