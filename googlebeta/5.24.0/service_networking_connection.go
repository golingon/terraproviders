// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	servicenetworkingconnection "github.com/golingon/terraproviders/googlebeta/5.24.0/servicenetworkingconnection"
	"io"
)

// NewServiceNetworkingConnection creates a new instance of [ServiceNetworkingConnection].
func NewServiceNetworkingConnection(name string, args ServiceNetworkingConnectionArgs) *ServiceNetworkingConnection {
	return &ServiceNetworkingConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceNetworkingConnection)(nil)

// ServiceNetworkingConnection represents the Terraform resource google_service_networking_connection.
type ServiceNetworkingConnection struct {
	Name      string
	Args      ServiceNetworkingConnectionArgs
	state     *serviceNetworkingConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) Type() string {
	return "google_service_networking_connection"
}

// LocalName returns the local name for [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) LocalName() string {
	return snc.Name
}

// Configuration returns the configuration (args) for [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) Configuration() interface{} {
	return snc.Args
}

// DependOn is used for other resources to depend on [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(snc)
}

// Dependencies returns the list of resources [ServiceNetworkingConnection] depends_on.
func (snc *ServiceNetworkingConnection) Dependencies() terra.Dependencies {
	return snc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) LifecycleManagement() *terra.Lifecycle {
	return snc.Lifecycle
}

// Attributes returns the attributes for [ServiceNetworkingConnection].
func (snc *ServiceNetworkingConnection) Attributes() serviceNetworkingConnectionAttributes {
	return serviceNetworkingConnectionAttributes{ref: terra.ReferenceResource(snc)}
}

// ImportState imports the given attribute values into [ServiceNetworkingConnection]'s state.
func (snc *ServiceNetworkingConnection) ImportState(av io.Reader) error {
	snc.state = &serviceNetworkingConnectionState{}
	if err := json.NewDecoder(av).Decode(snc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snc.Type(), snc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceNetworkingConnection] has state.
func (snc *ServiceNetworkingConnection) State() (*serviceNetworkingConnectionState, bool) {
	return snc.state, snc.state != nil
}

// StateMust returns the state for [ServiceNetworkingConnection]. Panics if the state is nil.
func (snc *ServiceNetworkingConnection) StateMust() *serviceNetworkingConnectionState {
	if snc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snc.Type(), snc.LocalName()))
	}
	return snc.state
}

// ServiceNetworkingConnectionArgs contains the configurations for google_service_networking_connection.
type ServiceNetworkingConnectionArgs struct {
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// ReservedPeeringRanges: list of string, required
	ReservedPeeringRanges terra.ListValue[terra.StringValue] `hcl:"reserved_peering_ranges,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *servicenetworkingconnection.Timeouts `hcl:"timeouts,block"`
}
type serviceNetworkingConnectionAttributes struct {
	ref terra.Reference
}

// DeletionPolicy returns a reference to field deletion_policy of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("deletion_policy"))
}

// Id returns a reference to field id of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("id"))
}

// Network returns a reference to field network of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("network"))
}

// Peering returns a reference to field peering of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) Peering() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("peering"))
}

// ReservedPeeringRanges returns a reference to field reserved_peering_ranges of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) ReservedPeeringRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](snc.ref.Append("reserved_peering_ranges"))
}

// Service returns a reference to field service of google_service_networking_connection.
func (snc serviceNetworkingConnectionAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("service"))
}

func (snc serviceNetworkingConnectionAttributes) Timeouts() servicenetworkingconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicenetworkingconnection.TimeoutsAttributes](snc.ref.Append("timeouts"))
}

type serviceNetworkingConnectionState struct {
	DeletionPolicy        string                                     `json:"deletion_policy"`
	Id                    string                                     `json:"id"`
	Network               string                                     `json:"network"`
	Peering               string                                     `json:"peering"`
	ReservedPeeringRanges []string                                   `json:"reserved_peering_ranges"`
	Service               string                                     `json:"service"`
	Timeouts              *servicenetworkingconnection.TimeoutsState `json:"timeouts"`
}
