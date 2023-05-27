// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxpublicvirtualinterface "github.com/golingon/terraproviders/aws/5.0.1/dxpublicvirtualinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxPublicVirtualInterface creates a new instance of [DxPublicVirtualInterface].
func NewDxPublicVirtualInterface(name string, args DxPublicVirtualInterfaceArgs) *DxPublicVirtualInterface {
	return &DxPublicVirtualInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxPublicVirtualInterface)(nil)

// DxPublicVirtualInterface represents the Terraform resource aws_dx_public_virtual_interface.
type DxPublicVirtualInterface struct {
	Name      string
	Args      DxPublicVirtualInterfaceArgs
	state     *dxPublicVirtualInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) Type() string {
	return "aws_dx_public_virtual_interface"
}

// LocalName returns the local name for [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) LocalName() string {
	return dpvi.Name
}

// Configuration returns the configuration (args) for [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) Configuration() interface{} {
	return dpvi.Args
}

// DependOn is used for other resources to depend on [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(dpvi)
}

// Dependencies returns the list of resources [DxPublicVirtualInterface] depends_on.
func (dpvi *DxPublicVirtualInterface) Dependencies() terra.Dependencies {
	return dpvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) LifecycleManagement() *terra.Lifecycle {
	return dpvi.Lifecycle
}

// Attributes returns the attributes for [DxPublicVirtualInterface].
func (dpvi *DxPublicVirtualInterface) Attributes() dxPublicVirtualInterfaceAttributes {
	return dxPublicVirtualInterfaceAttributes{ref: terra.ReferenceResource(dpvi)}
}

// ImportState imports the given attribute values into [DxPublicVirtualInterface]'s state.
func (dpvi *DxPublicVirtualInterface) ImportState(av io.Reader) error {
	dpvi.state = &dxPublicVirtualInterfaceState{}
	if err := json.NewDecoder(av).Decode(dpvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpvi.Type(), dpvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxPublicVirtualInterface] has state.
func (dpvi *DxPublicVirtualInterface) State() (*dxPublicVirtualInterfaceState, bool) {
	return dpvi.state, dpvi.state != nil
}

// StateMust returns the state for [DxPublicVirtualInterface]. Panics if the state is nil.
func (dpvi *DxPublicVirtualInterface) StateMust() *dxPublicVirtualInterfaceState {
	if dpvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpvi.Type(), dpvi.LocalName()))
	}
	return dpvi.state
}

// DxPublicVirtualInterfaceArgs contains the configurations for aws_dx_public_virtual_interface.
type DxPublicVirtualInterfaceArgs struct {
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
	// RouteFilterPrefixes: set of string, required
	RouteFilterPrefixes terra.SetValue[terra.StringValue] `hcl:"route_filter_prefixes,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Vlan: number, required
	Vlan terra.NumberValue `hcl:"vlan,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxpublicvirtualinterface.Timeouts `hcl:"timeouts,block"`
}
type dxPublicVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("amazon_address"))
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(dpvi.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("bgp_auth_key"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("connection_id"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("id"))
}

// Name returns a reference to field name of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpvi.ref.Append("name"))
}

// RouteFilterPrefixes returns a reference to field route_filter_prefixes of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) RouteFilterPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dpvi.ref.Append("route_filter_prefixes"))
}

// Tags returns a reference to field tags of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpvi.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpvi.ref.Append("tags_all"))
}

// Vlan returns a reference to field vlan of aws_dx_public_virtual_interface.
func (dpvi dxPublicVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(dpvi.ref.Append("vlan"))
}

func (dpvi dxPublicVirtualInterfaceAttributes) Timeouts() dxpublicvirtualinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxpublicvirtualinterface.TimeoutsAttributes](dpvi.ref.Append("timeouts"))
}

type dxPublicVirtualInterfaceState struct {
	AddressFamily       string                                  `json:"address_family"`
	AmazonAddress       string                                  `json:"amazon_address"`
	AmazonSideAsn       string                                  `json:"amazon_side_asn"`
	Arn                 string                                  `json:"arn"`
	AwsDevice           string                                  `json:"aws_device"`
	BgpAsn              float64                                 `json:"bgp_asn"`
	BgpAuthKey          string                                  `json:"bgp_auth_key"`
	ConnectionId        string                                  `json:"connection_id"`
	CustomerAddress     string                                  `json:"customer_address"`
	Id                  string                                  `json:"id"`
	Name                string                                  `json:"name"`
	RouteFilterPrefixes []string                                `json:"route_filter_prefixes"`
	Tags                map[string]string                       `json:"tags"`
	TagsAll             map[string]string                       `json:"tags_all"`
	Vlan                float64                                 `json:"vlan"`
	Timeouts            *dxpublicvirtualinterface.TimeoutsState `json:"timeouts"`
}
