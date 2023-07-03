// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	ipgroupcidr "github.com/golingon/terraproviders/azurerm/3.63.0/ipgroupcidr"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIpGroupCidr creates a new instance of [IpGroupCidr].
func NewIpGroupCidr(name string, args IpGroupCidrArgs) *IpGroupCidr {
	return &IpGroupCidr{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IpGroupCidr)(nil)

// IpGroupCidr represents the Terraform resource azurerm_ip_group_cidr.
type IpGroupCidr struct {
	Name      string
	Args      IpGroupCidrArgs
	state     *ipGroupCidrState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IpGroupCidr].
func (igc *IpGroupCidr) Type() string {
	return "azurerm_ip_group_cidr"
}

// LocalName returns the local name for [IpGroupCidr].
func (igc *IpGroupCidr) LocalName() string {
	return igc.Name
}

// Configuration returns the configuration (args) for [IpGroupCidr].
func (igc *IpGroupCidr) Configuration() interface{} {
	return igc.Args
}

// DependOn is used for other resources to depend on [IpGroupCidr].
func (igc *IpGroupCidr) DependOn() terra.Reference {
	return terra.ReferenceResource(igc)
}

// Dependencies returns the list of resources [IpGroupCidr] depends_on.
func (igc *IpGroupCidr) Dependencies() terra.Dependencies {
	return igc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IpGroupCidr].
func (igc *IpGroupCidr) LifecycleManagement() *terra.Lifecycle {
	return igc.Lifecycle
}

// Attributes returns the attributes for [IpGroupCidr].
func (igc *IpGroupCidr) Attributes() ipGroupCidrAttributes {
	return ipGroupCidrAttributes{ref: terra.ReferenceResource(igc)}
}

// ImportState imports the given attribute values into [IpGroupCidr]'s state.
func (igc *IpGroupCidr) ImportState(av io.Reader) error {
	igc.state = &ipGroupCidrState{}
	if err := json.NewDecoder(av).Decode(igc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", igc.Type(), igc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IpGroupCidr] has state.
func (igc *IpGroupCidr) State() (*ipGroupCidrState, bool) {
	return igc.state, igc.state != nil
}

// StateMust returns the state for [IpGroupCidr]. Panics if the state is nil.
func (igc *IpGroupCidr) StateMust() *ipGroupCidrState {
	if igc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", igc.Type(), igc.LocalName()))
	}
	return igc.state
}

// IpGroupCidrArgs contains the configurations for azurerm_ip_group_cidr.
type IpGroupCidrArgs struct {
	// Cidr: string, required
	Cidr terra.StringValue `hcl:"cidr,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpGroupId: string, required
	IpGroupId terra.StringValue `hcl:"ip_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ipgroupcidr.Timeouts `hcl:"timeouts,block"`
}
type ipGroupCidrAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of azurerm_ip_group_cidr.
func (igc ipGroupCidrAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(igc.ref.Append("cidr"))
}

// Id returns a reference to field id of azurerm_ip_group_cidr.
func (igc ipGroupCidrAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(igc.ref.Append("id"))
}

// IpGroupId returns a reference to field ip_group_id of azurerm_ip_group_cidr.
func (igc ipGroupCidrAttributes) IpGroupId() terra.StringValue {
	return terra.ReferenceAsString(igc.ref.Append("ip_group_id"))
}

func (igc ipGroupCidrAttributes) Timeouts() ipgroupcidr.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ipgroupcidr.TimeoutsAttributes](igc.ref.Append("timeouts"))
}

type ipGroupCidrState struct {
	Cidr      string                     `json:"cidr"`
	Id        string                     `json:"id"`
	IpGroupId string                     `json:"ip_group_id"`
	Timeouts  *ipgroupcidr.TimeoutsState `json:"timeouts"`
}
