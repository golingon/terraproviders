// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_globalaccelerator_custom_routing_accelerator

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

// Resource represents the Terraform resource aws_globalaccelerator_custom_routing_accelerator.
type Resource struct {
	Name      string
	Args      Args
	state     *awsGlobalacceleratorCustomRoutingAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (agcra *Resource) Type() string {
	return "aws_globalaccelerator_custom_routing_accelerator"
}

// LocalName returns the local name for [Resource].
func (agcra *Resource) LocalName() string {
	return agcra.Name
}

// Configuration returns the configuration (args) for [Resource].
func (agcra *Resource) Configuration() interface{} {
	return agcra.Args
}

// DependOn is used for other resources to depend on [Resource].
func (agcra *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(agcra)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (agcra *Resource) Dependencies() terra.Dependencies {
	return agcra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (agcra *Resource) LifecycleManagement() *terra.Lifecycle {
	return agcra.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (agcra *Resource) Attributes() awsGlobalacceleratorCustomRoutingAcceleratorAttributes {
	return awsGlobalacceleratorCustomRoutingAcceleratorAttributes{ref: terra.ReferenceResource(agcra)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (agcra *Resource) ImportState(state io.Reader) error {
	agcra.state = &awsGlobalacceleratorCustomRoutingAcceleratorState{}
	if err := json.NewDecoder(state).Decode(agcra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agcra.Type(), agcra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (agcra *Resource) State() (*awsGlobalacceleratorCustomRoutingAcceleratorState, bool) {
	return agcra.state, agcra.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (agcra *Resource) StateMust() *awsGlobalacceleratorCustomRoutingAcceleratorState {
	if agcra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agcra.Type(), agcra.LocalName()))
	}
	return agcra.state
}

// Args contains the configurations for aws_globalaccelerator_custom_routing_accelerator.
type Args struct {
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
	// Attributes: optional
	Attributes *Attributes `hcl:"attributes,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsGlobalacceleratorCustomRoutingAcceleratorAttributes struct {
	ref terra.Reference
}

// DnsName returns a reference to field dns_name of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(agcra.ref.Append("dns_name"))
}

// Enabled returns a reference to field enabled of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(agcra.ref.Append("enabled"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(agcra.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agcra.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(agcra.ref.Append("ip_address_type"))
}

// IpAddresses returns a reference to field ip_addresses of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](agcra.ref.Append("ip_addresses"))
}

// Name returns a reference to field name of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agcra.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agcra.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_globalaccelerator_custom_routing_accelerator.
func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agcra.ref.Append("tags_all"))
}

func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) IpSets() terra.ListValue[IpSetsAttributes] {
	return terra.ReferenceAsList[IpSetsAttributes](agcra.ref.Append("ip_sets"))
}

func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Attributes() terra.ListValue[AttributesAttributes] {
	return terra.ReferenceAsList[AttributesAttributes](agcra.ref.Append("attributes"))
}

func (agcra awsGlobalacceleratorCustomRoutingAcceleratorAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](agcra.ref.Append("timeouts"))
}

type awsGlobalacceleratorCustomRoutingAcceleratorState struct {
	DnsName       string            `json:"dns_name"`
	Enabled       bool              `json:"enabled"`
	HostedZoneId  string            `json:"hosted_zone_id"`
	Id            string            `json:"id"`
	IpAddressType string            `json:"ip_address_type"`
	IpAddresses   []string          `json:"ip_addresses"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	TagsAll       map[string]string `json:"tags_all"`
	IpSets        []IpSetsState     `json:"ip_sets"`
	Attributes    []AttributesState `json:"attributes"`
	Timeouts      *TimeoutsState    `json:"timeouts"`
}
