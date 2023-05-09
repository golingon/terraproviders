// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	globalacceleratoraccelerator "github.com/golingon/terraproviders/aws/4.66.1/globalacceleratoraccelerator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlobalacceleratorAccelerator creates a new instance of [GlobalacceleratorAccelerator].
func NewGlobalacceleratorAccelerator(name string, args GlobalacceleratorAcceleratorArgs) *GlobalacceleratorAccelerator {
	return &GlobalacceleratorAccelerator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlobalacceleratorAccelerator)(nil)

// GlobalacceleratorAccelerator represents the Terraform resource aws_globalaccelerator_accelerator.
type GlobalacceleratorAccelerator struct {
	Name      string
	Args      GlobalacceleratorAcceleratorArgs
	state     *globalacceleratorAcceleratorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) Type() string {
	return "aws_globalaccelerator_accelerator"
}

// LocalName returns the local name for [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) LocalName() string {
	return ga.Name
}

// Configuration returns the configuration (args) for [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) Configuration() interface{} {
	return ga.Args
}

// DependOn is used for other resources to depend on [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) DependOn() terra.Reference {
	return terra.ReferenceResource(ga)
}

// Dependencies returns the list of resources [GlobalacceleratorAccelerator] depends_on.
func (ga *GlobalacceleratorAccelerator) Dependencies() terra.Dependencies {
	return ga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) LifecycleManagement() *terra.Lifecycle {
	return ga.Lifecycle
}

// Attributes returns the attributes for [GlobalacceleratorAccelerator].
func (ga *GlobalacceleratorAccelerator) Attributes() globalacceleratorAcceleratorAttributes {
	return globalacceleratorAcceleratorAttributes{ref: terra.ReferenceResource(ga)}
}

// ImportState imports the given attribute values into [GlobalacceleratorAccelerator]'s state.
func (ga *GlobalacceleratorAccelerator) ImportState(av io.Reader) error {
	ga.state = &globalacceleratorAcceleratorState{}
	if err := json.NewDecoder(av).Decode(ga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ga.Type(), ga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlobalacceleratorAccelerator] has state.
func (ga *GlobalacceleratorAccelerator) State() (*globalacceleratorAcceleratorState, bool) {
	return ga.state, ga.state != nil
}

// StateMust returns the state for [GlobalacceleratorAccelerator]. Panics if the state is nil.
func (ga *GlobalacceleratorAccelerator) StateMust() *globalacceleratorAcceleratorState {
	if ga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ga.Type(), ga.LocalName()))
	}
	return ga.state
}

// GlobalacceleratorAcceleratorArgs contains the configurations for aws_globalaccelerator_accelerator.
type GlobalacceleratorAcceleratorArgs struct {
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
	IpSets []globalacceleratoraccelerator.IpSets `hcl:"ip_sets,block" validate:"min=0"`
	// Attributes: optional
	Attributes *globalacceleratoraccelerator.Attributes `hcl:"attributes,block"`
	// Timeouts: optional
	Timeouts *globalacceleratoraccelerator.Timeouts `hcl:"timeouts,block"`
}
type globalacceleratorAcceleratorAttributes struct {
	ref terra.Reference
}

// DnsName returns a reference to field dns_name of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("dns_name"))
}

// DualStackDnsName returns a reference to field dual_stack_dns_name of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) DualStackDnsName() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("dual_stack_dns_name"))
}

// Enabled returns a reference to field enabled of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ga.ref.Append("enabled"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("ip_address_type"))
}

// IpAddresses returns a reference to field ip_addresses of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ga.ref.Append("ip_addresses"))
}

// Name returns a reference to field name of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_globalaccelerator_accelerator.
func (ga globalacceleratorAcceleratorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags_all"))
}

func (ga globalacceleratorAcceleratorAttributes) IpSets() terra.ListValue[globalacceleratoraccelerator.IpSetsAttributes] {
	return terra.ReferenceAsList[globalacceleratoraccelerator.IpSetsAttributes](ga.ref.Append("ip_sets"))
}

func (ga globalacceleratorAcceleratorAttributes) Attributes() terra.ListValue[globalacceleratoraccelerator.AttributesAttributes] {
	return terra.ReferenceAsList[globalacceleratoraccelerator.AttributesAttributes](ga.ref.Append("attributes"))
}

func (ga globalacceleratorAcceleratorAttributes) Timeouts() globalacceleratoraccelerator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[globalacceleratoraccelerator.TimeoutsAttributes](ga.ref.Append("timeouts"))
}

type globalacceleratorAcceleratorState struct {
	DnsName          string                                         `json:"dns_name"`
	DualStackDnsName string                                         `json:"dual_stack_dns_name"`
	Enabled          bool                                           `json:"enabled"`
	HostedZoneId     string                                         `json:"hosted_zone_id"`
	Id               string                                         `json:"id"`
	IpAddressType    string                                         `json:"ip_address_type"`
	IpAddresses      []string                                       `json:"ip_addresses"`
	Name             string                                         `json:"name"`
	Tags             map[string]string                              `json:"tags"`
	TagsAll          map[string]string                              `json:"tags_all"`
	IpSets           []globalacceleratoraccelerator.IpSetsState     `json:"ip_sets"`
	Attributes       []globalacceleratoraccelerator.AttributesState `json:"attributes"`
	Timeouts         *globalacceleratoraccelerator.TimeoutsState    `json:"timeouts"`
}
