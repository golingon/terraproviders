// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	servicedirectoryendpoint "github.com/golingon/terraproviders/googlebeta/4.71.0/servicedirectoryendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryEndpoint creates a new instance of [ServiceDirectoryEndpoint].
func NewServiceDirectoryEndpoint(name string, args ServiceDirectoryEndpointArgs) *ServiceDirectoryEndpoint {
	return &ServiceDirectoryEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryEndpoint)(nil)

// ServiceDirectoryEndpoint represents the Terraform resource google_service_directory_endpoint.
type ServiceDirectoryEndpoint struct {
	Name      string
	Args      ServiceDirectoryEndpointArgs
	state     *serviceDirectoryEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) Type() string {
	return "google_service_directory_endpoint"
}

// LocalName returns the local name for [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) LocalName() string {
	return sde.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) Configuration() interface{} {
	return sde.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(sde)
}

// Dependencies returns the list of resources [ServiceDirectoryEndpoint] depends_on.
func (sde *ServiceDirectoryEndpoint) Dependencies() terra.Dependencies {
	return sde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) LifecycleManagement() *terra.Lifecycle {
	return sde.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryEndpoint].
func (sde *ServiceDirectoryEndpoint) Attributes() serviceDirectoryEndpointAttributes {
	return serviceDirectoryEndpointAttributes{ref: terra.ReferenceResource(sde)}
}

// ImportState imports the given attribute values into [ServiceDirectoryEndpoint]'s state.
func (sde *ServiceDirectoryEndpoint) ImportState(av io.Reader) error {
	sde.state = &serviceDirectoryEndpointState{}
	if err := json.NewDecoder(av).Decode(sde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sde.Type(), sde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryEndpoint] has state.
func (sde *ServiceDirectoryEndpoint) State() (*serviceDirectoryEndpointState, bool) {
	return sde.state, sde.state != nil
}

// StateMust returns the state for [ServiceDirectoryEndpoint]. Panics if the state is nil.
func (sde *ServiceDirectoryEndpoint) StateMust() *serviceDirectoryEndpointState {
	if sde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sde.Type(), sde.LocalName()))
	}
	return sde.state
}

// ServiceDirectoryEndpointArgs contains the configurations for google_service_directory_endpoint.
type ServiceDirectoryEndpointArgs struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// EndpointId: string, required
	EndpointId terra.StringValue `hcl:"endpoint_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicedirectoryendpoint.Timeouts `hcl:"timeouts,block"`
}
type serviceDirectoryEndpointAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("address"))
}

// EndpointId returns a reference to field endpoint_id of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("endpoint_id"))
}

// Id returns a reference to field id of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("id"))
}

// Metadata returns a reference to field metadata of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sde.ref.Append("metadata"))
}

// Name returns a reference to field name of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("name"))
}

// Network returns a reference to field network of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("network"))
}

// Port returns a reference to field port of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(sde.ref.Append("port"))
}

// Service returns a reference to field service of google_service_directory_endpoint.
func (sde serviceDirectoryEndpointAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(sde.ref.Append("service"))
}

func (sde serviceDirectoryEndpointAttributes) Timeouts() servicedirectoryendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicedirectoryendpoint.TimeoutsAttributes](sde.ref.Append("timeouts"))
}

type serviceDirectoryEndpointState struct {
	Address    string                                  `json:"address"`
	EndpointId string                                  `json:"endpoint_id"`
	Id         string                                  `json:"id"`
	Metadata   map[string]string                       `json:"metadata"`
	Name       string                                  `json:"name"`
	Network    string                                  `json:"network"`
	Port       float64                                 `json:"port"`
	Service    string                                  `json:"service"`
	Timeouts   *servicedirectoryendpoint.TimeoutsState `json:"timeouts"`
}
