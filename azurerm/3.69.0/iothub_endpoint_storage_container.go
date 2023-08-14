// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubendpointstoragecontainer "github.com/golingon/terraproviders/azurerm/3.69.0/iothubendpointstoragecontainer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubEndpointStorageContainer creates a new instance of [IothubEndpointStorageContainer].
func NewIothubEndpointStorageContainer(name string, args IothubEndpointStorageContainerArgs) *IothubEndpointStorageContainer {
	return &IothubEndpointStorageContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubEndpointStorageContainer)(nil)

// IothubEndpointStorageContainer represents the Terraform resource azurerm_iothub_endpoint_storage_container.
type IothubEndpointStorageContainer struct {
	Name      string
	Args      IothubEndpointStorageContainerArgs
	state     *iothubEndpointStorageContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) Type() string {
	return "azurerm_iothub_endpoint_storage_container"
}

// LocalName returns the local name for [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) LocalName() string {
	return iesc.Name
}

// Configuration returns the configuration (args) for [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) Configuration() interface{} {
	return iesc.Args
}

// DependOn is used for other resources to depend on [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) DependOn() terra.Reference {
	return terra.ReferenceResource(iesc)
}

// Dependencies returns the list of resources [IothubEndpointStorageContainer] depends_on.
func (iesc *IothubEndpointStorageContainer) Dependencies() terra.Dependencies {
	return iesc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) LifecycleManagement() *terra.Lifecycle {
	return iesc.Lifecycle
}

// Attributes returns the attributes for [IothubEndpointStorageContainer].
func (iesc *IothubEndpointStorageContainer) Attributes() iothubEndpointStorageContainerAttributes {
	return iothubEndpointStorageContainerAttributes{ref: terra.ReferenceResource(iesc)}
}

// ImportState imports the given attribute values into [IothubEndpointStorageContainer]'s state.
func (iesc *IothubEndpointStorageContainer) ImportState(av io.Reader) error {
	iesc.state = &iothubEndpointStorageContainerState{}
	if err := json.NewDecoder(av).Decode(iesc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iesc.Type(), iesc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubEndpointStorageContainer] has state.
func (iesc *IothubEndpointStorageContainer) State() (*iothubEndpointStorageContainerState, bool) {
	return iesc.state, iesc.state != nil
}

// StateMust returns the state for [IothubEndpointStorageContainer]. Panics if the state is nil.
func (iesc *IothubEndpointStorageContainer) StateMust() *iothubEndpointStorageContainerState {
	if iesc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iesc.Type(), iesc.LocalName()))
	}
	return iesc.state
}

// IothubEndpointStorageContainerArgs contains the configurations for azurerm_iothub_endpoint_storage_container.
type IothubEndpointStorageContainerArgs struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// BatchFrequencyInSeconds: number, optional
	BatchFrequencyInSeconds terra.NumberValue `hcl:"batch_frequency_in_seconds,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// EndpointUri: string, optional
	EndpointUri terra.StringValue `hcl:"endpoint_uri,attr"`
	// FileNameFormat: string, optional
	FileNameFormat terra.StringValue `hcl:"file_name_format,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// MaxChunkSizeInBytes: number, optional
	MaxChunkSizeInBytes terra.NumberValue `hcl:"max_chunk_size_in_bytes,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iothubendpointstoragecontainer.Timeouts `hcl:"timeouts,block"`
}
type iothubEndpointStorageContainerAttributes struct {
	ref terra.Reference
}

// AuthenticationType returns a reference to field authentication_type of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("authentication_type"))
}

// BatchFrequencyInSeconds returns a reference to field batch_frequency_in_seconds of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) BatchFrequencyInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(iesc.ref.Append("batch_frequency_in_seconds"))
}

// ConnectionString returns a reference to field connection_string of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("connection_string"))
}

// ContainerName returns a reference to field container_name of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("container_name"))
}

// Encoding returns a reference to field encoding of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("encoding"))
}

// EndpointUri returns a reference to field endpoint_uri of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("endpoint_uri"))
}

// FileNameFormat returns a reference to field file_name_format of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) FileNameFormat() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("file_name_format"))
}

// Id returns a reference to field id of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("id"))
}

// IdentityId returns a reference to field identity_id of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("identity_id"))
}

// IothubId returns a reference to field iothub_id of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("iothub_id"))
}

// MaxChunkSizeInBytes returns a reference to field max_chunk_size_in_bytes of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) MaxChunkSizeInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(iesc.ref.Append("max_chunk_size_in_bytes"))
}

// Name returns a reference to field name of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_endpoint_storage_container.
func (iesc iothubEndpointStorageContainerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(iesc.ref.Append("resource_group_name"))
}

func (iesc iothubEndpointStorageContainerAttributes) Timeouts() iothubendpointstoragecontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubendpointstoragecontainer.TimeoutsAttributes](iesc.ref.Append("timeouts"))
}

type iothubEndpointStorageContainerState struct {
	AuthenticationType      string                                        `json:"authentication_type"`
	BatchFrequencyInSeconds float64                                       `json:"batch_frequency_in_seconds"`
	ConnectionString        string                                        `json:"connection_string"`
	ContainerName           string                                        `json:"container_name"`
	Encoding                string                                        `json:"encoding"`
	EndpointUri             string                                        `json:"endpoint_uri"`
	FileNameFormat          string                                        `json:"file_name_format"`
	Id                      string                                        `json:"id"`
	IdentityId              string                                        `json:"identity_id"`
	IothubId                string                                        `json:"iothub_id"`
	MaxChunkSizeInBytes     float64                                       `json:"max_chunk_size_in_bytes"`
	Name                    string                                        `json:"name"`
	ResourceGroupName       string                                        `json:"resource_group_name"`
	Timeouts                *iothubendpointstoragecontainer.TimeoutsState `json:"timeouts"`
}
