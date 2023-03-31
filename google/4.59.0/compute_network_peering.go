// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenetworkpeering "github.com/golingon/terraproviders/google/4.59.0/computenetworkpeering"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeNetworkPeering(name string, args ComputeNetworkPeeringArgs) *ComputeNetworkPeering {
	return &ComputeNetworkPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkPeering)(nil)

type ComputeNetworkPeering struct {
	Name  string
	Args  ComputeNetworkPeeringArgs
	state *computeNetworkPeeringState
}

func (cnp *ComputeNetworkPeering) Type() string {
	return "google_compute_network_peering"
}

func (cnp *ComputeNetworkPeering) LocalName() string {
	return cnp.Name
}

func (cnp *ComputeNetworkPeering) Configuration() interface{} {
	return cnp.Args
}

func (cnp *ComputeNetworkPeering) Attributes() computeNetworkPeeringAttributes {
	return computeNetworkPeeringAttributes{ref: terra.ReferenceResource(cnp)}
}

func (cnp *ComputeNetworkPeering) ImportState(av io.Reader) error {
	cnp.state = &computeNetworkPeeringState{}
	if err := json.NewDecoder(av).Decode(cnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnp.Type(), cnp.LocalName(), err)
	}
	return nil
}

func (cnp *ComputeNetworkPeering) State() (*computeNetworkPeeringState, bool) {
	return cnp.state, cnp.state != nil
}

func (cnp *ComputeNetworkPeering) StateMust() *computeNetworkPeeringState {
	if cnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnp.Type(), cnp.LocalName()))
	}
	return cnp.state
}

func (cnp *ComputeNetworkPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(cnp)
}

type ComputeNetworkPeeringArgs struct {
	// ExportCustomRoutes: bool, optional
	ExportCustomRoutes terra.BoolValue `hcl:"export_custom_routes,attr"`
	// ExportSubnetRoutesWithPublicIp: bool, optional
	ExportSubnetRoutesWithPublicIp terra.BoolValue `hcl:"export_subnet_routes_with_public_ip,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportCustomRoutes: bool, optional
	ImportCustomRoutes terra.BoolValue `hcl:"import_custom_routes,attr"`
	// ImportSubnetRoutesWithPublicIp: bool, optional
	ImportSubnetRoutesWithPublicIp terra.BoolValue `hcl:"import_subnet_routes_with_public_ip,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// PeerNetwork: string, required
	PeerNetwork terra.StringValue `hcl:"peer_network,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computenetworkpeering.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeNetworkPeering depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeNetworkPeeringAttributes struct {
	ref terra.Reference
}

func (cnp computeNetworkPeeringAttributes) ExportCustomRoutes() terra.BoolValue {
	return terra.ReferenceBool(cnp.ref.Append("export_custom_routes"))
}

func (cnp computeNetworkPeeringAttributes) ExportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceBool(cnp.ref.Append("export_subnet_routes_with_public_ip"))
}

func (cnp computeNetworkPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("id"))
}

func (cnp computeNetworkPeeringAttributes) ImportCustomRoutes() terra.BoolValue {
	return terra.ReferenceBool(cnp.ref.Append("import_custom_routes"))
}

func (cnp computeNetworkPeeringAttributes) ImportSubnetRoutesWithPublicIp() terra.BoolValue {
	return terra.ReferenceBool(cnp.ref.Append("import_subnet_routes_with_public_ip"))
}

func (cnp computeNetworkPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("name"))
}

func (cnp computeNetworkPeeringAttributes) Network() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("network"))
}

func (cnp computeNetworkPeeringAttributes) PeerNetwork() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("peer_network"))
}

func (cnp computeNetworkPeeringAttributes) State() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("state"))
}

func (cnp computeNetworkPeeringAttributes) StateDetails() terra.StringValue {
	return terra.ReferenceString(cnp.ref.Append("state_details"))
}

func (cnp computeNetworkPeeringAttributes) Timeouts() computenetworkpeering.TimeoutsAttributes {
	return terra.ReferenceSingle[computenetworkpeering.TimeoutsAttributes](cnp.ref.Append("timeouts"))
}

type computeNetworkPeeringState struct {
	ExportCustomRoutes             bool                                 `json:"export_custom_routes"`
	ExportSubnetRoutesWithPublicIp bool                                 `json:"export_subnet_routes_with_public_ip"`
	Id                             string                               `json:"id"`
	ImportCustomRoutes             bool                                 `json:"import_custom_routes"`
	ImportSubnetRoutesWithPublicIp bool                                 `json:"import_subnet_routes_with_public_ip"`
	Name                           string                               `json:"name"`
	Network                        string                               `json:"network"`
	PeerNetwork                    string                               `json:"peer_network"`
	State                          string                               `json:"state"`
	StateDetails                   string                               `json:"state_details"`
	Timeouts                       *computenetworkpeering.TimeoutsState `json:"timeouts"`
}
