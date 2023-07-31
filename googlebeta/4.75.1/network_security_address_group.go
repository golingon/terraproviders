// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networksecurityaddressgroup "github.com/golingon/terraproviders/googlebeta/4.75.1/networksecurityaddressgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityAddressGroup creates a new instance of [NetworkSecurityAddressGroup].
func NewNetworkSecurityAddressGroup(name string, args NetworkSecurityAddressGroupArgs) *NetworkSecurityAddressGroup {
	return &NetworkSecurityAddressGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityAddressGroup)(nil)

// NetworkSecurityAddressGroup represents the Terraform resource google_network_security_address_group.
type NetworkSecurityAddressGroup struct {
	Name      string
	Args      NetworkSecurityAddressGroupArgs
	state     *networkSecurityAddressGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) Type() string {
	return "google_network_security_address_group"
}

// LocalName returns the local name for [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) LocalName() string {
	return nsag.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) Configuration() interface{} {
	return nsag.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(nsag)
}

// Dependencies returns the list of resources [NetworkSecurityAddressGroup] depends_on.
func (nsag *NetworkSecurityAddressGroup) Dependencies() terra.Dependencies {
	return nsag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) LifecycleManagement() *terra.Lifecycle {
	return nsag.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityAddressGroup].
func (nsag *NetworkSecurityAddressGroup) Attributes() networkSecurityAddressGroupAttributes {
	return networkSecurityAddressGroupAttributes{ref: terra.ReferenceResource(nsag)}
}

// ImportState imports the given attribute values into [NetworkSecurityAddressGroup]'s state.
func (nsag *NetworkSecurityAddressGroup) ImportState(av io.Reader) error {
	nsag.state = &networkSecurityAddressGroupState{}
	if err := json.NewDecoder(av).Decode(nsag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsag.Type(), nsag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityAddressGroup] has state.
func (nsag *NetworkSecurityAddressGroup) State() (*networkSecurityAddressGroupState, bool) {
	return nsag.state, nsag.state != nil
}

// StateMust returns the state for [NetworkSecurityAddressGroup]. Panics if the state is nil.
func (nsag *NetworkSecurityAddressGroup) StateMust() *networkSecurityAddressGroupState {
	if nsag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsag.Type(), nsag.LocalName()))
	}
	return nsag.state
}

// NetworkSecurityAddressGroupArgs contains the configurations for google_network_security_address_group.
type NetworkSecurityAddressGroupArgs struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Items: list of string, optional
	Items terra.ListValue[terra.StringValue] `hcl:"items,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networksecurityaddressgroup.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityAddressGroupAttributes struct {
	ref terra.Reference
}

// Capacity returns a reference to field capacity of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(nsag.ref.Append("capacity"))
}

// CreateTime returns a reference to field create_time of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("description"))
}

// Id returns a reference to field id of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("id"))
}

// Items returns a reference to field items of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Items() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsag.ref.Append("items"))
}

// Labels returns a reference to field labels of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsag.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("name"))
}

// Parent returns a reference to field parent of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("parent"))
}

// Type returns a reference to field type of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_network_security_address_group.
func (nsag networkSecurityAddressGroupAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsag.ref.Append("update_time"))
}

func (nsag networkSecurityAddressGroupAttributes) Timeouts() networksecurityaddressgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecurityaddressgroup.TimeoutsAttributes](nsag.ref.Append("timeouts"))
}

type networkSecurityAddressGroupState struct {
	Capacity    float64                                    `json:"capacity"`
	CreateTime  string                                     `json:"create_time"`
	Description string                                     `json:"description"`
	Id          string                                     `json:"id"`
	Items       []string                                   `json:"items"`
	Labels      map[string]string                          `json:"labels"`
	Location    string                                     `json:"location"`
	Name        string                                     `json:"name"`
	Parent      string                                     `json:"parent"`
	Type        string                                     `json:"type"`
	UpdateTime  string                                     `json:"update_time"`
	Timeouts    *networksecurityaddressgroup.TimeoutsState `json:"timeouts"`
}
