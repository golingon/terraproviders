// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53TrafficPolicyInstance creates a new instance of [Route53TrafficPolicyInstance].
func NewRoute53TrafficPolicyInstance(name string, args Route53TrafficPolicyInstanceArgs) *Route53TrafficPolicyInstance {
	return &Route53TrafficPolicyInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53TrafficPolicyInstance)(nil)

// Route53TrafficPolicyInstance represents the Terraform resource aws_route53_traffic_policy_instance.
type Route53TrafficPolicyInstance struct {
	Name      string
	Args      Route53TrafficPolicyInstanceArgs
	state     *route53TrafficPolicyInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) Type() string {
	return "aws_route53_traffic_policy_instance"
}

// LocalName returns the local name for [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) LocalName() string {
	return rtpi.Name
}

// Configuration returns the configuration (args) for [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) Configuration() interface{} {
	return rtpi.Args
}

// DependOn is used for other resources to depend on [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(rtpi)
}

// Dependencies returns the list of resources [Route53TrafficPolicyInstance] depends_on.
func (rtpi *Route53TrafficPolicyInstance) Dependencies() terra.Dependencies {
	return rtpi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) LifecycleManagement() *terra.Lifecycle {
	return rtpi.Lifecycle
}

// Attributes returns the attributes for [Route53TrafficPolicyInstance].
func (rtpi *Route53TrafficPolicyInstance) Attributes() route53TrafficPolicyInstanceAttributes {
	return route53TrafficPolicyInstanceAttributes{ref: terra.ReferenceResource(rtpi)}
}

// ImportState imports the given attribute values into [Route53TrafficPolicyInstance]'s state.
func (rtpi *Route53TrafficPolicyInstance) ImportState(av io.Reader) error {
	rtpi.state = &route53TrafficPolicyInstanceState{}
	if err := json.NewDecoder(av).Decode(rtpi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rtpi.Type(), rtpi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53TrafficPolicyInstance] has state.
func (rtpi *Route53TrafficPolicyInstance) State() (*route53TrafficPolicyInstanceState, bool) {
	return rtpi.state, rtpi.state != nil
}

// StateMust returns the state for [Route53TrafficPolicyInstance]. Panics if the state is nil.
func (rtpi *Route53TrafficPolicyInstance) StateMust() *route53TrafficPolicyInstanceState {
	if rtpi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rtpi.Type(), rtpi.LocalName()))
	}
	return rtpi.state
}

// Route53TrafficPolicyInstanceArgs contains the configurations for aws_route53_traffic_policy_instance.
type Route53TrafficPolicyInstanceArgs struct {
	// HostedZoneId: string, required
	HostedZoneId terra.StringValue `hcl:"hosted_zone_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TrafficPolicyId: string, required
	TrafficPolicyId terra.StringValue `hcl:"traffic_policy_id,attr" validate:"required"`
	// TrafficPolicyVersion: number, required
	TrafficPolicyVersion terra.NumberValue `hcl:"traffic_policy_version,attr" validate:"required"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
}
type route53TrafficPolicyInstanceAttributes struct {
	ref terra.Reference
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(rtpi.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rtpi.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rtpi.ref.Append("name"))
}

// TrafficPolicyId returns a reference to field traffic_policy_id of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) TrafficPolicyId() terra.StringValue {
	return terra.ReferenceAsString(rtpi.ref.Append("traffic_policy_id"))
}

// TrafficPolicyVersion returns a reference to field traffic_policy_version of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) TrafficPolicyVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(rtpi.ref.Append("traffic_policy_version"))
}

// Ttl returns a reference to field ttl of aws_route53_traffic_policy_instance.
func (rtpi route53TrafficPolicyInstanceAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(rtpi.ref.Append("ttl"))
}

type route53TrafficPolicyInstanceState struct {
	HostedZoneId         string  `json:"hosted_zone_id"`
	Id                   string  `json:"id"`
	Name                 string  `json:"name"`
	TrafficPolicyId      string  `json:"traffic_policy_id"`
	TrafficPolicyVersion float64 `json:"traffic_policy_version"`
	Ttl                  float64 `json:"ttl"`
}