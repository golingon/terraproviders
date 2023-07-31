// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailLb creates a new instance of [LightsailLb].
func NewLightsailLb(name string, args LightsailLbArgs) *LightsailLb {
	return &LightsailLb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailLb)(nil)

// LightsailLb represents the Terraform resource aws_lightsail_lb.
type LightsailLb struct {
	Name      string
	Args      LightsailLbArgs
	state     *lightsailLbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailLb].
func (ll *LightsailLb) Type() string {
	return "aws_lightsail_lb"
}

// LocalName returns the local name for [LightsailLb].
func (ll *LightsailLb) LocalName() string {
	return ll.Name
}

// Configuration returns the configuration (args) for [LightsailLb].
func (ll *LightsailLb) Configuration() interface{} {
	return ll.Args
}

// DependOn is used for other resources to depend on [LightsailLb].
func (ll *LightsailLb) DependOn() terra.Reference {
	return terra.ReferenceResource(ll)
}

// Dependencies returns the list of resources [LightsailLb] depends_on.
func (ll *LightsailLb) Dependencies() terra.Dependencies {
	return ll.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailLb].
func (ll *LightsailLb) LifecycleManagement() *terra.Lifecycle {
	return ll.Lifecycle
}

// Attributes returns the attributes for [LightsailLb].
func (ll *LightsailLb) Attributes() lightsailLbAttributes {
	return lightsailLbAttributes{ref: terra.ReferenceResource(ll)}
}

// ImportState imports the given attribute values into [LightsailLb]'s state.
func (ll *LightsailLb) ImportState(av io.Reader) error {
	ll.state = &lightsailLbState{}
	if err := json.NewDecoder(av).Decode(ll.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ll.Type(), ll.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailLb] has state.
func (ll *LightsailLb) State() (*lightsailLbState, bool) {
	return ll.state, ll.state != nil
}

// StateMust returns the state for [LightsailLb]. Panics if the state is nil.
func (ll *LightsailLb) StateMust() *lightsailLbState {
	if ll.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ll.Type(), ll.LocalName()))
	}
	return ll.state
}

// LightsailLbArgs contains the configurations for aws_lightsail_lb.
type LightsailLbArgs struct {
	// HealthCheckPath: string, optional
	HealthCheckPath terra.StringValue `hcl:"health_check_path,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstancePort: number, required
	InstancePort terra.NumberValue `hcl:"instance_port,attr" validate:"required"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type lightsailLbAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_lb.
func (ll lightsailLbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_lb.
func (ll lightsailLbAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("created_at"))
}

// DnsName returns a reference to field dns_name of aws_lightsail_lb.
func (ll lightsailLbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("dns_name"))
}

// HealthCheckPath returns a reference to field health_check_path of aws_lightsail_lb.
func (ll lightsailLbAttributes) HealthCheckPath() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("health_check_path"))
}

// Id returns a reference to field id of aws_lightsail_lb.
func (ll lightsailLbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("id"))
}

// InstancePort returns a reference to field instance_port of aws_lightsail_lb.
func (ll lightsailLbAttributes) InstancePort() terra.NumberValue {
	return terra.ReferenceAsNumber(ll.ref.Append("instance_port"))
}

// IpAddressType returns a reference to field ip_address_type of aws_lightsail_lb.
func (ll lightsailLbAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("ip_address_type"))
}

// Name returns a reference to field name of aws_lightsail_lb.
func (ll lightsailLbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("name"))
}

// Protocol returns a reference to field protocol of aws_lightsail_lb.
func (ll lightsailLbAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("protocol"))
}

// PublicPorts returns a reference to field public_ports of aws_lightsail_lb.
func (ll lightsailLbAttributes) PublicPorts() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](ll.ref.Append("public_ports"))
}

// SupportCode returns a reference to field support_code of aws_lightsail_lb.
func (ll lightsailLbAttributes) SupportCode() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("support_code"))
}

// Tags returns a reference to field tags of aws_lightsail_lb.
func (ll lightsailLbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ll.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_lb.
func (ll lightsailLbAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ll.ref.Append("tags_all"))
}

type lightsailLbState struct {
	Arn             string            `json:"arn"`
	CreatedAt       string            `json:"created_at"`
	DnsName         string            `json:"dns_name"`
	HealthCheckPath string            `json:"health_check_path"`
	Id              string            `json:"id"`
	InstancePort    float64           `json:"instance_port"`
	IpAddressType   string            `json:"ip_address_type"`
	Name            string            `json:"name"`
	Protocol        string            `json:"protocol"`
	PublicPorts     []float64         `json:"public_ports"`
	SupportCode     string            `json:"support_code"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}
