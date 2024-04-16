// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dx_public_virtual_interface

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_dx_public_virtual_interface.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDxPublicVirtualInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adpvi *Resource) Type() string {
	return "aws_dx_public_virtual_interface"
}

// LocalName returns the local name for [Resource].
func (adpvi *Resource) LocalName() string {
	return adpvi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adpvi *Resource) Configuration() interface{} {
	return adpvi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adpvi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adpvi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adpvi *Resource) Dependencies() terra.Dependencies {
	return adpvi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adpvi *Resource) LifecycleManagement() *terra.Lifecycle {
	return adpvi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adpvi *Resource) Attributes() awsDxPublicVirtualInterfaceAttributes {
	return awsDxPublicVirtualInterfaceAttributes{ref: terra.ReferenceResource(adpvi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adpvi *Resource) ImportState(state io.Reader) error {
	adpvi.state = &awsDxPublicVirtualInterfaceState{}
	if err := json.NewDecoder(state).Decode(adpvi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adpvi.Type(), adpvi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adpvi *Resource) State() (*awsDxPublicVirtualInterfaceState, bool) {
	return adpvi.state, adpvi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adpvi *Resource) StateMust() *awsDxPublicVirtualInterfaceState {
	if adpvi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adpvi.Type(), adpvi.LocalName()))
	}
	return adpvi.state
}

// Args contains the configurations for aws_dx_public_virtual_interface.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDxPublicVirtualInterfaceAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("amazon_address"))
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) AmazonSideAsn() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(adpvi.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("bgp_auth_key"))
}

// ConnectionId returns a reference to field connection_id of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("connection_id"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("id"))
}

// Name returns a reference to field name of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adpvi.ref.Append("name"))
}

// RouteFilterPrefixes returns a reference to field route_filter_prefixes of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) RouteFilterPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adpvi.ref.Append("route_filter_prefixes"))
}

// Tags returns a reference to field tags of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adpvi.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adpvi.ref.Append("tags_all"))
}

// Vlan returns a reference to field vlan of aws_dx_public_virtual_interface.
func (adpvi awsDxPublicVirtualInterfaceAttributes) Vlan() terra.NumberValue {
	return terra.ReferenceAsNumber(adpvi.ref.Append("vlan"))
}

func (adpvi awsDxPublicVirtualInterfaceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adpvi.ref.Append("timeouts"))
}

type awsDxPublicVirtualInterfaceState struct {
	AddressFamily       string            `json:"address_family"`
	AmazonAddress       string            `json:"amazon_address"`
	AmazonSideAsn       string            `json:"amazon_side_asn"`
	Arn                 string            `json:"arn"`
	AwsDevice           string            `json:"aws_device"`
	BgpAsn              float64           `json:"bgp_asn"`
	BgpAuthKey          string            `json:"bgp_auth_key"`
	ConnectionId        string            `json:"connection_id"`
	CustomerAddress     string            `json:"customer_address"`
	Id                  string            `json:"id"`
	Name                string            `json:"name"`
	RouteFilterPrefixes []string          `json:"route_filter_prefixes"`
	Tags                map[string]string `json:"tags"`
	TagsAll             map[string]string `json:"tags_all"`
	Vlan                float64           `json:"vlan"`
	Timeouts            *TimeoutsState    `json:"timeouts"`
}
