// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicesmesh "github.com/golingon/terraproviders/googlebeta/4.72.1/networkservicesmesh"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesMesh creates a new instance of [NetworkServicesMesh].
func NewNetworkServicesMesh(name string, args NetworkServicesMeshArgs) *NetworkServicesMesh {
	return &NetworkServicesMesh{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesMesh)(nil)

// NetworkServicesMesh represents the Terraform resource google_network_services_mesh.
type NetworkServicesMesh struct {
	Name      string
	Args      NetworkServicesMeshArgs
	state     *networkServicesMeshState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) Type() string {
	return "google_network_services_mesh"
}

// LocalName returns the local name for [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) LocalName() string {
	return nsm.Name
}

// Configuration returns the configuration (args) for [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) Configuration() interface{} {
	return nsm.Args
}

// DependOn is used for other resources to depend on [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) DependOn() terra.Reference {
	return terra.ReferenceResource(nsm)
}

// Dependencies returns the list of resources [NetworkServicesMesh] depends_on.
func (nsm *NetworkServicesMesh) Dependencies() terra.Dependencies {
	return nsm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) LifecycleManagement() *terra.Lifecycle {
	return nsm.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesMesh].
func (nsm *NetworkServicesMesh) Attributes() networkServicesMeshAttributes {
	return networkServicesMeshAttributes{ref: terra.ReferenceResource(nsm)}
}

// ImportState imports the given attribute values into [NetworkServicesMesh]'s state.
func (nsm *NetworkServicesMesh) ImportState(av io.Reader) error {
	nsm.state = &networkServicesMeshState{}
	if err := json.NewDecoder(av).Decode(nsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsm.Type(), nsm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesMesh] has state.
func (nsm *NetworkServicesMesh) State() (*networkServicesMeshState, bool) {
	return nsm.state, nsm.state != nil
}

// StateMust returns the state for [NetworkServicesMesh]. Panics if the state is nil.
func (nsm *NetworkServicesMesh) StateMust() *networkServicesMeshState {
	if nsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsm.Type(), nsm.LocalName()))
	}
	return nsm.state
}

// NetworkServicesMeshArgs contains the configurations for google_network_services_mesh.
type NetworkServicesMeshArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InterceptionPort: number, optional
	InterceptionPort terra.NumberValue `hcl:"interception_port,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *networkservicesmesh.Timeouts `hcl:"timeouts,block"`
}
type networkServicesMeshAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("description"))
}

// Id returns a reference to field id of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("id"))
}

// InterceptionPort returns a reference to field interception_port of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) InterceptionPort() terra.NumberValue {
	return terra.ReferenceAsNumber(nsm.ref.Append("interception_port"))
}

// Labels returns a reference to field labels of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsm.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("self_link"))
}

// UpdateTime returns a reference to field update_time of google_network_services_mesh.
func (nsm networkServicesMeshAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsm.ref.Append("update_time"))
}

func (nsm networkServicesMeshAttributes) Timeouts() networkservicesmesh.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesmesh.TimeoutsAttributes](nsm.ref.Append("timeouts"))
}

type networkServicesMeshState struct {
	CreateTime       string                             `json:"create_time"`
	Description      string                             `json:"description"`
	Id               string                             `json:"id"`
	InterceptionPort float64                            `json:"interception_port"`
	Labels           map[string]string                  `json:"labels"`
	Name             string                             `json:"name"`
	Project          string                             `json:"project"`
	SelfLink         string                             `json:"self_link"`
	UpdateTime       string                             `json:"update_time"`
	Timeouts         *networkservicesmesh.TimeoutsState `json:"timeouts"`
}
