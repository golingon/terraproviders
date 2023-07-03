// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vmwareenginenetwork "github.com/golingon/terraproviders/googlebeta/4.71.0/vmwareenginenetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareengineNetwork creates a new instance of [VmwareengineNetwork].
func NewVmwareengineNetwork(name string, args VmwareengineNetworkArgs) *VmwareengineNetwork {
	return &VmwareengineNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareengineNetwork)(nil)

// VmwareengineNetwork represents the Terraform resource google_vmwareengine_network.
type VmwareengineNetwork struct {
	Name      string
	Args      VmwareengineNetworkArgs
	state     *vmwareengineNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareengineNetwork].
func (vn *VmwareengineNetwork) Type() string {
	return "google_vmwareengine_network"
}

// LocalName returns the local name for [VmwareengineNetwork].
func (vn *VmwareengineNetwork) LocalName() string {
	return vn.Name
}

// Configuration returns the configuration (args) for [VmwareengineNetwork].
func (vn *VmwareengineNetwork) Configuration() interface{} {
	return vn.Args
}

// DependOn is used for other resources to depend on [VmwareengineNetwork].
func (vn *VmwareengineNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(vn)
}

// Dependencies returns the list of resources [VmwareengineNetwork] depends_on.
func (vn *VmwareengineNetwork) Dependencies() terra.Dependencies {
	return vn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareengineNetwork].
func (vn *VmwareengineNetwork) LifecycleManagement() *terra.Lifecycle {
	return vn.Lifecycle
}

// Attributes returns the attributes for [VmwareengineNetwork].
func (vn *VmwareengineNetwork) Attributes() vmwareengineNetworkAttributes {
	return vmwareengineNetworkAttributes{ref: terra.ReferenceResource(vn)}
}

// ImportState imports the given attribute values into [VmwareengineNetwork]'s state.
func (vn *VmwareengineNetwork) ImportState(av io.Reader) error {
	vn.state = &vmwareengineNetworkState{}
	if err := json.NewDecoder(av).Decode(vn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vn.Type(), vn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareengineNetwork] has state.
func (vn *VmwareengineNetwork) State() (*vmwareengineNetworkState, bool) {
	return vn.state, vn.state != nil
}

// StateMust returns the state for [VmwareengineNetwork]. Panics if the state is nil.
func (vn *VmwareengineNetwork) StateMust() *vmwareengineNetworkState {
	if vn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vn.Type(), vn.LocalName()))
	}
	return vn.state
}

// VmwareengineNetworkArgs contains the configurations for google_vmwareengine_network.
type VmwareengineNetworkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// VpcNetworks: min=0
	VpcNetworks []vmwareenginenetwork.VpcNetworks `hcl:"vpc_networks,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vmwareenginenetwork.Timeouts `hcl:"timeouts,block"`
}
type vmwareengineNetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("description"))
}

// Id returns a reference to field id of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("id"))
}

// Location returns a reference to field location of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("location"))
}

// Name returns a reference to field name of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("name"))
}

// Project returns a reference to field project of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("project"))
}

// State returns a reference to field state of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("state"))
}

// Type returns a reference to field type of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("type"))
}

// Uid returns a reference to field uid of google_vmwareengine_network.
func (vn vmwareengineNetworkAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("uid"))
}

func (vn vmwareengineNetworkAttributes) VpcNetworks() terra.ListValue[vmwareenginenetwork.VpcNetworksAttributes] {
	return terra.ReferenceAsList[vmwareenginenetwork.VpcNetworksAttributes](vn.ref.Append("vpc_networks"))
}

func (vn vmwareengineNetworkAttributes) Timeouts() vmwareenginenetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareenginenetwork.TimeoutsAttributes](vn.ref.Append("timeouts"))
}

type vmwareengineNetworkState struct {
	Description string                                 `json:"description"`
	Id          string                                 `json:"id"`
	Location    string                                 `json:"location"`
	Name        string                                 `json:"name"`
	Project     string                                 `json:"project"`
	State       string                                 `json:"state"`
	Type        string                                 `json:"type"`
	Uid         string                                 `json:"uid"`
	VpcNetworks []vmwareenginenetwork.VpcNetworksState `json:"vpc_networks"`
	Timeouts    *vmwareenginenetwork.TimeoutsState     `json:"timeouts"`
}
