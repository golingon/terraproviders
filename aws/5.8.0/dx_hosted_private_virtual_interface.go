// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedprivatevirtualinterface "github.com/golingon/terraproviders/aws/5.8.0/dxhostedprivatevirtualinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxHostedPrivateVirtualInterface creates a new instance of [DxHostedPrivateVirtualInterface].
func NewDxHostedPrivateVirtualInterface(name string, args DxHostedPrivateVirtualInterfaceArgs) *DxHostedPrivateVirtualInterface {
	return &DxHostedPrivateVirtualInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedPrivateVirtualInterface)(nil)

// DxHostedPrivateVirtualInterface represents the Terraform resource aws_dx_hosted_private_virtual_interface.
type DxHostedPrivateVirtualInterface struct {
	Name      string
	Args      DxHostedPrivateVirtualInterfaceArgs
	state     *dxHostedPrivateVirtualInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) Type() string {
	return "aws_dx_hosted_private_virtual_interface"
}

// LocalName returns the local name for [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) LocalName() string {
	return dhpvi.Name
}

// Configuration returns the configuration (args) for [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) Configuration() interface{} {
	return dhpvi.Args
}

// DependOn is used for other resources to depend on [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(dhpvi)
}

// Dependencies returns the list of resources [DxHostedPrivateVirtualInterface] depends_on.
func (dhpvi *DxHostedPrivateVirtualInterface) Dependencies() terra.Dependencies {
	return dhpvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) LifecycleManagement() *terra.Lifecycle {
	return dhpvi.Lifecycle
}

// Attributes returns the attributes for [DxHostedPrivateVirtualInterface].
func (dhpvi *DxHostedPrivateVirtualInterface) Attributes() dxHostedPrivateVirtualInterfaceAttributes {
	return dxHostedPrivateVirtualInterfaceAttributes{ref: terra.ReferenceResource(dhpvi)}
}

// ImportState imports the given attribute values into [DxHostedPrivateVirtualInterface]'s state.
func (dhpvi *DxHostedPrivateVirtualInterface) ImportState(av io.Reader) error {
	dhpvi.state = &dxHostedPrivateVirtualInterfaceState{}
	if err := json.NewDecoder(av).Decode(dhpvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhpvi.Type(), dhpvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxHostedPrivateVirtualInterface] has state.
func (dhpvi *DxHostedPrivateVirtualInterface) State() (*dxHostedPrivateVirtualInterfaceState, bool) {
	return dhpvi.state, dhpvi.state != nil
}

// StateMust returns the state for [DxHostedPrivateVirtualInterface]. Panics if the state is nil.
func (dhpvi *DxHostedPrivateVirtualInterface) StateMust() *dxHostedPrivateVirtualInterfaceState {
	if dhpvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhpvi.Type(), dhpvi.LocalName()))
	}
	return dhpvi.state
}

// DxHostedPrivateVirtualInterfaceArgs contains the configurations for aws_dx_hosted_private_virtual_interface.
type DxHostedPrivateVirtualInterfaceArgs struct {
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
	Timeouts *dxhostedprivatevirtualinterface.Timeouts `hcl:"timeouts,block"`
}
type dxHostedPrivateVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("amazon_address"))
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(dhpvi.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("bgp_auth_key"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("connection_id"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("id"))
}

// JumboFrameCapable returns a reference to field jumbo_frame_capable of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) JumboFrameCapable() terra.BoolValue {
	return terra.ReferenceAsBool(dhpvi.ref.Append("jumbo_frame_capable"))
}

// Mtu returns a reference to field mtu of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Mtu() terra.NumberValue {
	return terra.ReferenceAsNumber(dhpvi.ref.Append("mtu"))
}

// Name returns a reference to field name of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("owner_account_id"))
}

// Vlan returns a reference to field vlan of aws_dx_hosted_private_virtual_interface.
func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(dhpvi.ref.Append("vlan"))
}

func (dhpvi dxHostedPrivateVirtualInterfaceAttributes) Timeouts() dxhostedprivatevirtualinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxhostedprivatevirtualinterface.TimeoutsAttributes](dhpvi.ref.Append("timeouts"))
}

type dxHostedPrivateVirtualInterfaceState struct {
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
	Timeouts          *dxhostedprivatevirtualinterface.TimeoutsState `json:"timeouts"`
}