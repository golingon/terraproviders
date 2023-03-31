// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerappenvironmentstorage "github.com/golingon/terraproviders/azurerm/3.49.0/containerappenvironmentstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewContainerAppEnvironmentStorage(name string, args ContainerAppEnvironmentStorageArgs) *ContainerAppEnvironmentStorage {
	return &ContainerAppEnvironmentStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAppEnvironmentStorage)(nil)

type ContainerAppEnvironmentStorage struct {
	Name  string
	Args  ContainerAppEnvironmentStorageArgs
	state *containerAppEnvironmentStorageState
}

func (caes *ContainerAppEnvironmentStorage) Type() string {
	return "azurerm_container_app_environment_storage"
}

func (caes *ContainerAppEnvironmentStorage) LocalName() string {
	return caes.Name
}

func (caes *ContainerAppEnvironmentStorage) Configuration() interface{} {
	return caes.Args
}

func (caes *ContainerAppEnvironmentStorage) Attributes() containerAppEnvironmentStorageAttributes {
	return containerAppEnvironmentStorageAttributes{ref: terra.ReferenceResource(caes)}
}

func (caes *ContainerAppEnvironmentStorage) ImportState(av io.Reader) error {
	caes.state = &containerAppEnvironmentStorageState{}
	if err := json.NewDecoder(av).Decode(caes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caes.Type(), caes.LocalName(), err)
	}
	return nil
}

func (caes *ContainerAppEnvironmentStorage) State() (*containerAppEnvironmentStorageState, bool) {
	return caes.state, caes.state != nil
}

func (caes *ContainerAppEnvironmentStorage) StateMust() *containerAppEnvironmentStorageState {
	if caes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caes.Type(), caes.LocalName()))
	}
	return caes.state
}

func (caes *ContainerAppEnvironmentStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(caes)
}

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
	// DependsOn contains resources that ContainerAppEnvironmentStorage depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type containerAppEnvironmentStorageAttributes struct {
	ref terra.Reference
}

func (caes containerAppEnvironmentStorageAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("access_key"))
}

func (caes containerAppEnvironmentStorageAttributes) AccessMode() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("access_mode"))
}

func (caes containerAppEnvironmentStorageAttributes) AccountName() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("account_name"))
}

func (caes containerAppEnvironmentStorageAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("container_app_environment_id"))
}

func (caes containerAppEnvironmentStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("id"))
}

func (caes containerAppEnvironmentStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("name"))
}

func (caes containerAppEnvironmentStorageAttributes) ShareName() terra.StringValue {
	return terra.ReferenceString(caes.ref.Append("share_name"))
}

func (caes containerAppEnvironmentStorageAttributes) Timeouts() containerappenvironmentstorage.TimeoutsAttributes {
	return terra.ReferenceSingle[containerappenvironmentstorage.TimeoutsAttributes](caes.ref.Append("timeouts"))
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
