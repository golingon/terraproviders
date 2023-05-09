// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedtransitvirtualinterface "github.com/golingon/terraproviders/aws/4.66.1/dxhostedtransitvirtualinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxHostedTransitVirtualInterface creates a new instance of [DxHostedTransitVirtualInterface].
func NewDxHostedTransitVirtualInterface(name string, args DxHostedTransitVirtualInterfaceArgs) *DxHostedTransitVirtualInterface {
	return &DxHostedTransitVirtualInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedTransitVirtualInterface)(nil)

// DxHostedTransitVirtualInterface represents the Terraform resource aws_dx_hosted_transit_virtual_interface.
type DxHostedTransitVirtualInterface struct {
	Name      string
	Args      DxHostedTransitVirtualInterfaceArgs
	state     *dxHostedTransitVirtualInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) Type() string {
	return "aws_dx_hosted_transit_virtual_interface"
}

// LocalName returns the local name for [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) LocalName() string {
	return dhtvi.Name
}

// Configuration returns the configuration (args) for [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) Configuration() interface{} {
	return dhtvi.Args
}

// DependOn is used for other resources to depend on [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(dhtvi)
}

// Dependencies returns the list of resources [DxHostedTransitVirtualInterface] depends_on.
func (dhtvi *DxHostedTransitVirtualInterface) Dependencies() terra.Dependencies {
	return dhtvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) LifecycleManagement() *terra.Lifecycle {
	return dhtvi.Lifecycle
}

// Attributes returns the attributes for [DxHostedTransitVirtualInterface].
func (dhtvi *DxHostedTransitVirtualInterface) Attributes() dxHostedTransitVirtualInterfaceAttributes {
	return dxHostedTransitVirtualInterfaceAttributes{ref: terra.ReferenceResource(dhtvi)}
}

// ImportState imports the given attribute values into [DxHostedTransitVirtualInterface]'s state.
func (dhtvi *DxHostedTransitVirtualInterface) ImportState(av io.Reader) error {
	dhtvi.state = &dxHostedTransitVirtualInterfaceState{}
	if err := json.NewDecoder(av).Decode(dhtvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhtvi.Type(), dhtvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxHostedTransitVirtualInterface] has state.
func (dhtvi *DxHostedTransitVirtualInterface) State() (*dxHostedTransitVirtualInterfaceState, bool) {
	return dhtvi.state, dhtvi.state != nil
}

// StateMust returns the state for [DxHostedTransitVirtualInterface]. Panics if the state is nil.
func (dhtvi *DxHostedTransitVirtualInterface) StateMust() *dxHostedTransitVirtualInterfaceState {
	if dhtvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhtvi.Type(), dhtvi.LocalName()))
	}
	return dhtvi.state
}

// DxHostedTransitVirtualInterfaceArgs contains the configurations for aws_dx_hosted_transit_virtual_interface.
type DxHostedTransitVirtualInterfaceArgs struct {
	// AddressFamily: string, required
	AddressFamily terra.StringValue `hcl:"address_family,attr" validate:"required"`
	// AmazonAddress: string, optional
	AmazonAddress terra.StringValue `hcl:"amazon_address,attr"`
	// BgpAsn: number, required
	BgpAsn terra.NumberValue `hcl:"bgp_asn,attr" validate:"required"`
	// BgpAuthKey: string, optional
	BgpAuthKey terra.StringValue `hcl:"bgp_auth_key,attr"`
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// CustomerAddress: string, optional
	CustomerAddress terra.StringValue `hcl:"customer_address,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Mtu: number, optional
	Mtu terra.NumberValue `hcl:"mtu,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OwnerAccountId: string, required
	OwnerAccountId terra.StringValue `hcl:"owner_account_id,attr" validate:"required"`
	// Vlan: number, required
	Vlan terra.NumberValue `hcl:"vlan,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxhostedtransitvirtualinterface.Timeouts `hcl:"timeouts,block"`
}
type dxHostedTransitVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("amazon_address"))
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(dhtvi.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("bgp_auth_key"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("connection_id"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("id"))
}

// JumboFrameCapable returns a reference to field jumbo_frame_capable of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) JumboFrameCapable() terra.BoolValue {
	return terra.ReferenceAsBool(dhtvi.ref.Append("jumbo_frame_capable"))
}

// Mtu returns a reference to field mtu of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Mtu() terra.NumberValue {
	return terra.ReferenceAsNumber(dhtvi.ref.Append("mtu"))
}

// Name returns a reference to field name of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dhtvi.ref.Append("owner_account_id"))
}

// Vlan returns a reference to field vlan of aws_dx_hosted_transit_virtual_interface.
func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(dhtvi.ref.Append("vlan"))
}

func (dhtvi dxHostedTransitVirtualInterfaceAttributes) Timeouts() dxhostedtransitvirtualinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxhostedtransitvirtualinterface.TimeoutsAttributes](dhtvi.ref.Append("timeouts"))
}

type dxHostedTransitVirtualInterfaceState struct {
	AddressFamily     string                                         `json:"address_family"`
	AmazonAddress     string                                         `json:"amazon_address"`
	AmazonSideAsn     string                                         `json:"amazon_side_asn"`
	Arn               string                                         `json:"arn"`
	AwsDevice         string                                         `json:"aws_device"`
	BgpAsn            float64                                        `json:"bgp_asn"`
	BgpAuthKey        string                                         `json:"bgp_auth_key"`
	ConnectionId      string                                         `json:"connection_id"`
	CustomerAddress   string                                         `json:"customer_address"`
	Id                string                                         `json:"id"`
	JumboFrameCapable bool                                           `json:"jumbo_frame_capable"`
	Mtu               float64                                        `json:"mtu"`
	Name              string                                         `json:"name"`
	OwnerAccountId    string                                         `json:"owner_account_id"`
	Vlan              float64                                        `json:"vlan"`
	Timeouts          *dxhostedtransitvirtualinterface.TimeoutsState `json:"timeouts"`
}
