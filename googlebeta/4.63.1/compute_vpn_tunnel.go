// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computevpntunnel "github.com/golingon/terraproviders/googlebeta/4.63.1/computevpntunnel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeVpnTunnel creates a new instance of [ComputeVpnTunnel].
func NewComputeVpnTunnel(name string, args ComputeVpnTunnelArgs) *ComputeVpnTunnel {
	return &ComputeVpnTunnel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeVpnTunnel)(nil)

// ComputeVpnTunnel represents the Terraform resource google_compute_vpn_tunnel.
type ComputeVpnTunnel struct {
	Name      string
	Args      ComputeVpnTunnelArgs
	state     *computeVpnTunnelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) Type() string {
	return "google_compute_vpn_tunnel"
}

// LocalName returns the local name for [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) LocalName() string {
	return cvt.Name
}

// Configuration returns the configuration (args) for [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) Configuration() interface{} {
	return cvt.Args
}

// DependOn is used for other resources to depend on [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) DependOn() terra.Reference {
	return terra.ReferenceResource(cvt)
}

// Dependencies returns the list of resources [ComputeVpnTunnel] depends_on.
func (cvt *ComputeVpnTunnel) Dependencies() terra.Dependencies {
	return cvt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) LifecycleManagement() *terra.Lifecycle {
	return cvt.Lifecycle
}

// Attributes returns the attributes for [ComputeVpnTunnel].
func (cvt *ComputeVpnTunnel) Attributes() computeVpnTunnelAttributes {
	return computeVpnTunnelAttributes{ref: terra.ReferenceResource(cvt)}
}

// ImportState imports the given attribute values into [ComputeVpnTunnel]'s state.
func (cvt *ComputeVpnTunnel) ImportState(av io.Reader) error {
	cvt.state = &computeVpnTunnelState{}
	if err := json.NewDecoder(av).Decode(cvt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvt.Type(), cvt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeVpnTunnel] has state.
func (cvt *ComputeVpnTunnel) State() (*computeVpnTunnelState, bool) {
	return cvt.state, cvt.state != nil
}

// StateMust returns the state for [ComputeVpnTunnel]. Panics if the state is nil.
func (cvt *ComputeVpnTunnel) StateMust() *computeVpnTunnelState {
	if cvt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvt.Type(), cvt.LocalName()))
	}
	return cvt.state
}

// ComputeVpnTunnelArgs contains the configurations for google_compute_vpn_tunnel.
type ComputeVpnTunnelArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IkeVersion: number, optional
	IkeVersion terra.NumberValue `hcl:"ike_version,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LocalTrafficSelector: set of string, optional
	LocalTrafficSelector terra.SetValue[terra.StringValue] `hcl:"local_traffic_selector,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerExternalGateway: string, optional
	PeerExternalGateway terra.StringValue `hcl:"peer_external_gateway,attr"`
	// PeerExternalGatewayInterface: number, optional
	PeerExternalGatewayInterface terra.NumberValue `hcl:"peer_external_gateway_interface,attr"`
	// PeerGcpGateway: string, optional
	PeerGcpGateway terra.StringValue `hcl:"peer_gcp_gateway,attr"`
	// PeerIp: string, optional
	PeerIp terra.StringValue `hcl:"peer_ip,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RemoteTrafficSelector: set of string, optional
	RemoteTrafficSelector terra.SetValue[terra.StringValue] `hcl:"remote_traffic_selector,attr"`
	// Router: string, optional
	Router terra.StringValue `hcl:"router,attr"`
	// SharedSecret: string, required
	SharedSecret terra.StringValue `hcl:"shared_secret,attr" validate:"required"`
	// TargetVpnGateway: string, optional
	TargetVpnGateway terra.StringValue `hcl:"target_vpn_gateway,attr"`
	// VpnGateway: string, optional
	VpnGateway terra.StringValue `hcl:"vpn_gateway,attr"`
	// VpnGatewayInterface: number, optional
	VpnGatewayInterface terra.NumberValue `hcl:"vpn_gateway_interface,attr"`
	// Timeouts: optional
	Timeouts *computevpntunnel.Timeouts `hcl:"timeouts,block"`
}
type computeVpnTunnelAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("description"))
}

// DetailedStatus returns a reference to field detailed_status of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) DetailedStatus() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("detailed_status"))
}

// Id returns a reference to field id of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("id"))
}

// IkeVersion returns a reference to field ike_version of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) IkeVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(cvt.ref.Append("ike_version"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cvt.ref.Append("labels"))
}

// LocalTrafficSelector returns a reference to field local_traffic_selector of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) LocalTrafficSelector() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cvt.ref.Append("local_traffic_selector"))
}

// Name returns a reference to field name of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("name"))
}

// PeerExternalGateway returns a reference to field peer_external_gateway of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) PeerExternalGateway() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("peer_external_gateway"))
}

// PeerExternalGatewayInterface returns a reference to field peer_external_gateway_interface of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) PeerExternalGatewayInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(cvt.ref.Append("peer_external_gateway_interface"))
}

// PeerGcpGateway returns a reference to field peer_gcp_gateway of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) PeerGcpGateway() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("peer_gcp_gateway"))
}

// PeerIp returns a reference to field peer_ip of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) PeerIp() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("peer_ip"))
}

// Project returns a reference to field project of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("region"))
}

// RemoteTrafficSelector returns a reference to field remote_traffic_selector of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) RemoteTrafficSelector() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cvt.ref.Append("remote_traffic_selector"))
}

// Router returns a reference to field router of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("router"))
}

// SelfLink returns a reference to field self_link of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("self_link"))
}

// SharedSecret returns a reference to field shared_secret of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) SharedSecret() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("shared_secret"))
}

// SharedSecretHash returns a reference to field shared_secret_hash of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) SharedSecretHash() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("shared_secret_hash"))
}

// TargetVpnGateway returns a reference to field target_vpn_gateway of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) TargetVpnGateway() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("target_vpn_gateway"))
}

// TunnelId returns a reference to field tunnel_id of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) TunnelId() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("tunnel_id"))
}

// VpnGateway returns a reference to field vpn_gateway of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) VpnGateway() terra.StringValue {
	return terra.ReferenceAsString(cvt.ref.Append("vpn_gateway"))
}

// VpnGatewayInterface returns a reference to field vpn_gateway_interface of google_compute_vpn_tunnel.
func (cvt computeVpnTunnelAttributes) VpnGatewayInterface() terra.NumberValue {
	return terra.ReferenceAsNumber(cvt.ref.Append("vpn_gateway_interface"))
}

func (cvt computeVpnTunnelAttributes) Timeouts() computevpntunnel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computevpntunnel.TimeoutsAttributes](cvt.ref.Append("timeouts"))
}

type computeVpnTunnelState struct {
	CreationTimestamp            string                          `json:"creation_timestamp"`
	Description                  string                          `json:"description"`
	DetailedStatus               string                          `json:"detailed_status"`
	Id                           string                          `json:"id"`
	IkeVersion                   float64                         `json:"ike_version"`
	LabelFingerprint             string                          `json:"label_fingerprint"`
	Labels                       map[string]string               `json:"labels"`
	LocalTrafficSelector         []string                        `json:"local_traffic_selector"`
	Name                         string                          `json:"name"`
	PeerExternalGateway          string                          `json:"peer_external_gateway"`
	PeerExternalGatewayInterface float64                         `json:"peer_external_gateway_interface"`
	PeerGcpGateway               string                          `json:"peer_gcp_gateway"`
	PeerIp                       string                          `json:"peer_ip"`
	Project                      string                          `json:"project"`
	Region                       string                          `json:"region"`
	RemoteTrafficSelector        []string                        `json:"remote_traffic_selector"`
	Router                       string                          `json:"router"`
	SelfLink                     string                          `json:"self_link"`
	SharedSecret                 string                          `json:"shared_secret"`
	SharedSecretHash             string                          `json:"shared_secret_hash"`
	TargetVpnGateway             string                          `json:"target_vpn_gateway"`
	TunnelId                     string                          `json:"tunnel_id"`
	VpnGateway                   string                          `json:"vpn_gateway"`
	VpnGatewayInterface          float64                         `json:"vpn_gateway_interface"`
	Timeouts                     *computevpntunnel.TimeoutsState `json:"timeouts"`
}
