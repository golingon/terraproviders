// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedpublicvirtualinterface "github.com/golingon/terraproviders/aws/5.8.0/dxhostedpublicvirtualinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxHostedPublicVirtualInterface creates a new instance of [DxHostedPublicVirtualInterface].
func NewDxHostedPublicVirtualInterface(name string, args DxHostedPublicVirtualInterfaceArgs) *DxHostedPublicVirtualInterface {
	return &DxHostedPublicVirtualInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedPublicVirtualInterface)(nil)

// DxHostedPublicVirtualInterface represents the Terraform resource aws_dx_hosted_public_virtual_interface.
type DxHostedPublicVirtualInterface struct {
	Name      string
	Args      DxHostedPublicVirtualInterfaceArgs
	state     *dxHostedPublicVirtualInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) Type() string {
	return "aws_dx_hosted_public_virtual_interface"
}

// LocalName returns the local name for [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) LocalName() string {
	return dhpvi.Name
}

// Configuration returns the configuration (args) for [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) Configuration() interface{} {
	return dhpvi.Args
}

// DependOn is used for other resources to depend on [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(dhpvi)
}

// Dependencies returns the list of resources [DxHostedPublicVirtualInterface] depends_on.
func (dhpvi *DxHostedPublicVirtualInterface) Dependencies() terra.Dependencies {
	return dhpvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) LifecycleManagement() *terra.Lifecycle {
	return dhpvi.Lifecycle
}

// Attributes returns the attributes for [DxHostedPublicVirtualInterface].
func (dhpvi *DxHostedPublicVirtualInterface) Attributes() dxHostedPublicVirtualInterfaceAttributes {
	return dxHostedPublicVirtualInterfaceAttributes{ref: terra.ReferenceResource(dhpvi)}
}

// ImportState imports the given attribute values into [DxHostedPublicVirtualInterface]'s state.
func (dhpvi *DxHostedPublicVirtualInterface) ImportState(av io.Reader) error {
	dhpvi.state = &dxHostedPublicVirtualInterfaceState{}
	if err := json.NewDecoder(av).Decode(dhpvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhpvi.Type(), dhpvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxHostedPublicVirtualInterface] has state.
func (dhpvi *DxHostedPublicVirtualInterface) State() (*dxHostedPublicVirtualInterfaceState, bool) {
	return dhpvi.state, dhpvi.state != nil
}

// StateMust returns the state for [DxHostedPublicVirtualInterface]. Panics if the state is nil.
func (dhpvi *DxHostedPublicVirtualInterface) StateMust() *dxHostedPublicVirtualInterfaceState {
	if dhpvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhpvi.Type(), dhpvi.LocalName()))
	}
	return dhpvi.state
}

// DxHostedPublicVirtualInterfaceArgs contains the configurations for aws_dx_hosted_public_virtual_interface.
type DxHostedPublicVirtualInterfaceArgs struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OwnerAccountId: string, required
	OwnerAccountId terra.StringValue `hcl:"owner_account_id,attr" validate:"required"`
	// RouteFilterPrefixes: set of string, required
	RouteFilterPrefixes terra.SetValue[terra.StringValue] `hcl:"route_filter_prefixes,attr" validate:"required"`
	// Vlan: number, required
	Vlan terra.NumberValue `hcl:"vlan,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxhostedpublicvirtualinterface.Timeouts `hcl:"timeouts,block"`
}
type dxHostedPublicVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("amazon_address"))
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(dhpvi.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("bgp_auth_key"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("connection_id"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("id"))
}

// Name returns a reference to field name of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dhpvi.ref.Append("owner_account_id"))
}

// RouteFilterPrefixes returns a reference to field route_filter_prefixes of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) RouteFilterPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dhpvi.ref.Append("route_filter_prefixes"))
}

// Vlan returns a reference to field vlan of aws_dx_hosted_public_virtual_interface.
func (dhpvi dxHostedPublicVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(dhpvi.ref.Append("vlan"))
}

func (dhpvi dxHostedPublicVirtualInterfaceAttributes) Timeouts() dxhostedpublicvirtualinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxhostedpublicvirtualinterface.TimeoutsAttributes](dhpvi.ref.Append("timeouts"))
}

type dxHostedPublicVirtualInterfaceState struct {
	AddressFamily       string                                        `json:"address_family"`
	AmazonAddress       string                                        `json:"amazon_address"`
	AmazonSideAsn       string                                        `json:"amazon_side_asn"`
	Arn                 string                                        `json:"arn"`
	AwsDevice           string                                        `json:"aws_device"`
	BgpAsn              float64                                       `json:"bgp_asn"`
	BgpAuthKey          string                                        `json:"bgp_auth_key"`
	ConnectionId        string                                        `json:"connection_id"`
	CustomerAddress     string                                        `json:"customer_address"`
	Id                  string                                        `json:"id"`
	Name                string                                        `json:"name"`
	OwnerAccountId      string                                        `json:"owner_account_id"`
	RouteFilterPrefixes []string                                      `json:"route_filter_prefixes"`
	Vlan                float64                                       `json:"vlan"`
	Timeouts            *dxhostedpublicvirtualinterface.TimeoutsState `json:"timeouts"`
}
