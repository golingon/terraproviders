// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	eip "github.com/golingon/terraproviders/aws/5.12.0/eip"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEip creates a new instance of [Eip].
func NewEip(name string, args EipArgs) *Eip {
	return &Eip{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Eip)(nil)

// Eip represents the Terraform resource aws_eip.
type Eip struct {
	Name      string
	Args      EipArgs
	state     *eipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Eip].
func (e *Eip) Type() string {
	return "aws_eip"
}

// LocalName returns the local name for [Eip].
func (e *Eip) LocalName() string {
	return e.Name
}

// Configuration returns the configuration (args) for [Eip].
func (e *Eip) Configuration() interface{} {
	return e.Args
}

// DependOn is used for other resources to depend on [Eip].
func (e *Eip) DependOn() terra.Reference {
	return terra.ReferenceResource(e)
}

// Dependencies returns the list of resources [Eip] depends_on.
func (e *Eip) Dependencies() terra.Dependencies {
	return e.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Eip].
func (e *Eip) LifecycleManagement() *terra.Lifecycle {
	return e.Lifecycle
}

// Attributes returns the attributes for [Eip].
func (e *Eip) Attributes() eipAttributes {
	return eipAttributes{ref: terra.ReferenceResource(e)}
}

// ImportState imports the given attribute values into [Eip]'s state.
func (e *Eip) ImportState(av io.Reader) error {
	e.state = &eipState{}
	if err := json.NewDecoder(av).Decode(e.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", e.Type(), e.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Eip] has state.
func (e *Eip) State() (*eipState, bool) {
	return e.state, e.state != nil
}

// StateMust returns the state for [Eip]. Panics if the state is nil.
func (e *Eip) StateMust() *eipState {
	if e.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", e.Type(), e.LocalName()))
	}
	return e.state
}

// EipArgs contains the configurations for aws_eip.
type EipArgs struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// AssociateWithPrivateIp: string, optional
	AssociateWithPrivateIp terra.StringValue `hcl:"associate_with_private_ip,attr"`
	// CustomerOwnedIpv4Pool: string, optional
	CustomerOwnedIpv4Pool terra.StringValue `hcl:"customer_owned_ipv4_pool,attr"`
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// NetworkBorderGroup: string, optional
	NetworkBorderGroup terra.StringValue `hcl:"network_border_group,attr"`
	// NetworkInterface: string, optional
	NetworkInterface terra.StringValue `hcl:"network_interface,attr"`
	// PublicIpv4Pool: string, optional
	PublicIpv4Pool terra.StringValue `hcl:"public_ipv4_pool,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Vpc: bool, optional
	Vpc terra.BoolValue `hcl:"vpc,attr"`
	// Timeouts: optional
	Timeouts *eip.Timeouts `hcl:"timeouts,block"`
}
type eipAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of aws_eip.
func (e eipAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

// AllocationId returns a reference to field allocation_id of aws_eip.
func (e eipAttributes) AllocationId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("allocation_id"))
}

// AssociateWithPrivateIp returns a reference to field associate_with_private_ip of aws_eip.
func (e eipAttributes) AssociateWithPrivateIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("associate_with_private_ip"))
}

// AssociationId returns a reference to field association_id of aws_eip.
func (e eipAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("association_id"))
}

// CarrierIp returns a reference to field carrier_ip of aws_eip.
func (e eipAttributes) CarrierIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("carrier_ip"))
}

// CustomerOwnedIp returns a reference to field customer_owned_ip of aws_eip.
func (e eipAttributes) CustomerOwnedIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("customer_owned_ip"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_eip.
func (e eipAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("customer_owned_ipv4_pool"))
}

// Domain returns a reference to field domain of aws_eip.
func (e eipAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("domain"))
}

// Id returns a reference to field id of aws_eip.
func (e eipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("id"))
}

// Instance returns a reference to field instance of aws_eip.
func (e eipAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("instance"))
}

// NetworkBorderGroup returns a reference to field network_border_group of aws_eip.
func (e eipAttributes) NetworkBorderGroup() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("network_border_group"))
}

// NetworkInterface returns a reference to field network_interface of aws_eip.
func (e eipAttributes) NetworkInterface() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("network_interface"))
}

// PrivateDns returns a reference to field private_dns of aws_eip.
func (e eipAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("private_dns"))
}

// PrivateIp returns a reference to field private_ip of aws_eip.
func (e eipAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("private_ip"))
}

// PublicDns returns a reference to field public_dns of aws_eip.
func (e eipAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_dns"))
}

// PublicIp returns a reference to field public_ip of aws_eip.
func (e eipAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_ip"))
}

// PublicIpv4Pool returns a reference to field public_ipv4_pool of aws_eip.
func (e eipAttributes) PublicIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_ipv4_pool"))
}

// Tags returns a reference to field tags of aws_eip.
func (e eipAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eip.
func (e eipAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags_all"))
}

// Vpc returns a reference to field vpc of aws_eip.
func (e eipAttributes) Vpc() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("vpc"))
}

func (e eipAttributes) Timeouts() eip.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eip.TimeoutsAttributes](e.ref.Append("timeouts"))
}

type eipState struct {
	Address                string             `json:"address"`
	AllocationId           string             `json:"allocation_id"`
	AssociateWithPrivateIp string             `json:"associate_with_private_ip"`
	AssociationId          string             `json:"association_id"`
	CarrierIp              string             `json:"carrier_ip"`
	CustomerOwnedIp        string             `json:"customer_owned_ip"`
	CustomerOwnedIpv4Pool  string             `json:"customer_owned_ipv4_pool"`
	Domain                 string             `json:"domain"`
	Id                     string             `json:"id"`
	Instance               string             `json:"instance"`
	NetworkBorderGroup     string             `json:"network_border_group"`
	NetworkInterface       string             `json:"network_interface"`
	PrivateDns             string             `json:"private_dns"`
	PrivateIp              string             `json:"private_ip"`
	PublicDns              string             `json:"public_dns"`
	PublicIp               string             `json:"public_ip"`
	PublicIpv4Pool         string             `json:"public_ipv4_pool"`
	Tags                   map[string]string  `json:"tags"`
	TagsAll                map[string]string  `json:"tags_all"`
	Vpc                    bool               `json:"vpc"`
	Timeouts               *eip.TimeoutsState `json:"timeouts"`
}
