// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	publicip "github.com/golingon/terraproviders/azurestack/1.0.0/publicip"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPublicIp creates a new instance of [PublicIp].
func NewPublicIp(name string, args PublicIpArgs) *PublicIp {
	return &PublicIp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PublicIp)(nil)

// PublicIp represents the Terraform resource azurestack_public_ip.
type PublicIp struct {
	Name      string
	Args      PublicIpArgs
	state     *publicIpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PublicIp].
func (pi *PublicIp) Type() string {
	return "azurestack_public_ip"
}

// LocalName returns the local name for [PublicIp].
func (pi *PublicIp) LocalName() string {
	return pi.Name
}

// Configuration returns the configuration (args) for [PublicIp].
func (pi *PublicIp) Configuration() interface{} {
	return pi.Args
}

// DependOn is used for other resources to depend on [PublicIp].
func (pi *PublicIp) DependOn() terra.Reference {
	return terra.ReferenceResource(pi)
}

// Dependencies returns the list of resources [PublicIp] depends_on.
func (pi *PublicIp) Dependencies() terra.Dependencies {
	return pi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PublicIp].
func (pi *PublicIp) LifecycleManagement() *terra.Lifecycle {
	return pi.Lifecycle
}

// Attributes returns the attributes for [PublicIp].
func (pi *PublicIp) Attributes() publicIpAttributes {
	return publicIpAttributes{ref: terra.ReferenceResource(pi)}
}

// ImportState imports the given attribute values into [PublicIp]'s state.
func (pi *PublicIp) ImportState(av io.Reader) error {
	pi.state = &publicIpState{}
	if err := json.NewDecoder(av).Decode(pi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pi.Type(), pi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PublicIp] has state.
func (pi *PublicIp) State() (*publicIpState, bool) {
	return pi.state, pi.state != nil
}

// StateMust returns the state for [PublicIp]. Panics if the state is nil.
func (pi *PublicIp) StateMust() *publicIpState {
	if pi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pi.Type(), pi.LocalName()))
	}
	return pi.state
}

// PublicIpArgs contains the configurations for azurestack_public_ip.
type PublicIpArgs struct {
	// AllocationMethod: string, optional
	AllocationMethod terra.StringValue `hcl:"allocation_method,attr"`
	// DomainNameLabel: string, optional
	DomainNameLabel terra.StringValue `hcl:"domain_name_label,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleTimeoutInMinutes: number, optional
	IdleTimeoutInMinutes terra.NumberValue `hcl:"idle_timeout_in_minutes,attr"`
	// IpVersion: string, optional
	IpVersion terra.StringValue `hcl:"ip_version,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicIpAddressAllocation: string, optional
	PublicIpAddressAllocation terra.StringValue `hcl:"public_ip_address_allocation,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ReverseFqdn: string, optional
	ReverseFqdn terra.StringValue `hcl:"reverse_fqdn,attr"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *publicip.Timeouts `hcl:"timeouts,block"`
}
type publicIpAttributes struct {
	ref terra.Reference
}

// AllocationMethod returns a reference to field allocation_method of azurestack_public_ip.
func (pi publicIpAttributes) AllocationMethod() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("allocation_method"))
}

// DomainNameLabel returns a reference to field domain_name_label of azurestack_public_ip.
func (pi publicIpAttributes) DomainNameLabel() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("domain_name_label"))
}

// Fqdn returns a reference to field fqdn of azurestack_public_ip.
func (pi publicIpAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurestack_public_ip.
func (pi publicIpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("id"))
}

// IdleTimeoutInMinutes returns a reference to field idle_timeout_in_minutes of azurestack_public_ip.
func (pi publicIpAttributes) IdleTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(pi.ref.Append("idle_timeout_in_minutes"))
}

// IpAddress returns a reference to field ip_address of azurestack_public_ip.
func (pi publicIpAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("ip_address"))
}

// IpVersion returns a reference to field ip_version of azurestack_public_ip.
func (pi publicIpAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("ip_version"))
}

// Location returns a reference to field location of azurestack_public_ip.
func (pi publicIpAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_public_ip.
func (pi publicIpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("name"))
}

// PublicIpAddressAllocation returns a reference to field public_ip_address_allocation of azurestack_public_ip.
func (pi publicIpAttributes) PublicIpAddressAllocation() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("public_ip_address_allocation"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_public_ip.
func (pi publicIpAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("resource_group_name"))
}

// ReverseFqdn returns a reference to field reverse_fqdn of azurestack_public_ip.
func (pi publicIpAttributes) ReverseFqdn() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("reverse_fqdn"))
}

// Sku returns a reference to field sku of azurestack_public_ip.
func (pi publicIpAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurestack_public_ip.
func (pi publicIpAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pi.ref.Append("tags"))
}

func (pi publicIpAttributes) Timeouts() publicip.TimeoutsAttributes {
	return terra.ReferenceAsSingle[publicip.TimeoutsAttributes](pi.ref.Append("timeouts"))
}

type publicIpState struct {
	AllocationMethod          string                  `json:"allocation_method"`
	DomainNameLabel           string                  `json:"domain_name_label"`
	Fqdn                      string                  `json:"fqdn"`
	Id                        string                  `json:"id"`
	IdleTimeoutInMinutes      float64                 `json:"idle_timeout_in_minutes"`
	IpAddress                 string                  `json:"ip_address"`
	IpVersion                 string                  `json:"ip_version"`
	Location                  string                  `json:"location"`
	Name                      string                  `json:"name"`
	PublicIpAddressAllocation string                  `json:"public_ip_address_allocation"`
	ResourceGroupName         string                  `json:"resource_group_name"`
	ReverseFqdn               string                  `json:"reverse_fqdn"`
	Sku                       string                  `json:"sku"`
	Tags                      map[string]string       `json:"tags"`
	Timeouts                  *publicip.TimeoutsState `json:"timeouts"`
}
