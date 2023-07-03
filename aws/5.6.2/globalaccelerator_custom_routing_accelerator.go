// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	globalacceleratorcustomroutingaccelerator "github.com/golingon/terraproviders/aws/5.6.2/globalacceleratorcustomroutingaccelerator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlobalacceleratorCustomRoutingAccelerator creates a new instance of [GlobalacceleratorCustomRoutingAccelerator].
func NewGlobalacceleratorCustomRoutingAccelerator(name string, args GlobalacceleratorCustomRoutingAcceleratorArgs) *GlobalacceleratorCustomRoutingAccelerator {
	return &GlobalacceleratorCustomRoutingAccelerator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlobalacceleratorCustomRoutingAccelerator)(nil)

// GlobalacceleratorCustomRoutingAccelerator represents the Terraform resource aws_globalaccelerator_custom_routing_accelerator.
type GlobalacceleratorCustomRoutingAccelerator struct {
	Name      string
	Args      GlobalacceleratorCustomRoutingAcceleratorArgs
	state     *globalacceleratorCustomRoutingAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) Type() string {
	return "aws_globalaccelerator_custom_routing_accelerator"
}

// LocalName returns the local name for [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) LocalName() string {
	return gcra.Name
}

// Configuration returns the configuration (args) for [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) Configuration() interface{} {
	return gcra.Args
}

// DependOn is used for other resources to depend on [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) DependOn() terra.Reference {
	return terra.ReferenceResource(gcra)
}

// Dependencies returns the list of resources [GlobalacceleratorCustomRoutingAccelerator] depends_on.
func (gcra *GlobalacceleratorCustomRoutingAccelerator) Dependencies() terra.Dependencies {
	return gcra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) LifecycleManagement() *terra.Lifecycle {
	return gcra.Lifecycle
}

// Attributes returns the attributes for [GlobalacceleratorCustomRoutingAccelerator].
func (gcra *GlobalacceleratorCustomRoutingAccelerator) Attributes() globalacceleratorCustomRoutingAcceleratorAttributes {
	return globalacceleratorCustomRoutingAcceleratorAttributes{ref: terra.ReferenceResource(gcra)}
}

// ImportState imports the given attribute values into [GlobalacceleratorCustomRoutingAccelerator]'s state.
func (gcra *GlobalacceleratorCustomRoutingAccelerator) ImportState(av io.Reader) error {
	gcra.state = &globalacceleratorCustomRoutingAcceleratorState{}
	if err := json.NewDecoder(av).Decode(gcra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcra.Type(), gcra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlobalacceleratorCustomRoutingAccelerator] has state.
func (gcra *GlobalacceleratorCustomRoutingAccelerator) State() (*globalacceleratorCustomRoutingAcceleratorState, bool) {
	return gcra.state, gcra.state != nil
}

// StateMust returns the state for [GlobalacceleratorCustomRoutingAccelerator]. Panics if the state is nil.
func (gcra *GlobalacceleratorCustomRoutingAccelerator) StateMust() *globalacceleratorCustomRoutingAcceleratorState {
	if gcra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcra.Type(), gcra.LocalName()))
	}
	return gcra.state
}

// GlobalacceleratorCustomRoutingAcceleratorArgs contains the configurations for aws_globalaccelerator_custom_routing_accelerator.
type GlobalacceleratorCustomRoutingAcceleratorArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// IpAddresses: list of string, optional
	IpAddresses terra.ListValue[terra.StringValue] `hcl:"ip_addresses,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// IpSets: min=0
	IpSets []globalacceleratorcustomroutingaccelerator.IpSets `hcl:"ip_sets,block" validate:"min=0"`
	// Attributes: optional
	Attributes *globalacceleratorcustomroutingaccelerator.Attributes `hcl:"attributes,block"`
	// Timeouts: optional
	Timeouts *globalacceleratorcustomroutingaccelerator.Timeouts `hcl:"timeouts,block"`
}
type globalacceleratorCustomRoutingAcceleratorAttributes struct {
	ref terra.Reference
}

// DnsName returns a reference to field dns_name of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(gcra.ref.Append("dns_name"))
}

// Enabled returns a reference to field enabled of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(gcra.ref.Append("enabled"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(gcra.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcra.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(gcra.ref.Append("ip_address_type"))
}

// IpAddresses returns a reference to field ip_addresses of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcra.ref.Append("ip_addresses"))
}

// Name returns a reference to field name of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcra.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcra.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_globalaccelerator_custom_routing_accelerator.
func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcra.ref.Append("tags_all"))
}

func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) IpSets() terra.ListValue[globalacceleratorcustomroutingaccelerator.IpSetsAttributes] {
	return terra.ReferenceAsList[globalacceleratorcustomroutingaccelerator.IpSetsAttributes](gcra.ref.Append("ip_sets"))
}

func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Attributes() terra.ListValue[globalacceleratorcustomroutingaccelerator.AttributesAttributes] {
	return terra.ReferenceAsList[globalacceleratorcustomroutingaccelerator.AttributesAttributes](gcra.ref.Append("attributes"))
}

func (gcra globalacceleratorCustomRoutingAcceleratorAttributes) Timeouts() globalacceleratorcustomroutingaccelerator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[globalacceleratorcustomroutingaccelerator.TimeoutsAttributes](gcra.ref.Append("timeouts"))
}

type globalacceleratorCustomRoutingAcceleratorState struct {
	DnsName       string                                                      `json:"dns_name"`
	Enabled       bool                                                        `json:"enabled"`
	HostedZoneId  string                                                      `json:"hosted_zone_id"`
	Id            string                                                      `json:"id"`
	IpAddressType string                                                      `json:"ip_address_type"`
	IpAddresses   []string                                                    `json:"ip_addresses"`
	Name          string                                                      `json:"name"`
	Tags          map[string]string                                           `json:"tags"`
	TagsAll       map[string]string                                           `json:"tags_all"`
	IpSets        []globalacceleratorcustomroutingaccelerator.IpSetsState     `json:"ip_sets"`
	Attributes    []globalacceleratorcustomroutingaccelerator.AttributesState `json:"attributes"`
	Timeouts      *globalacceleratorcustomroutingaccelerator.TimeoutsState    `json:"timeouts"`
}
