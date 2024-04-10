// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRoute53HostedZoneDnssec creates a new instance of [Route53HostedZoneDnssec].
func NewRoute53HostedZoneDnssec(name string, args Route53HostedZoneDnssecArgs) *Route53HostedZoneDnssec {
	return &Route53HostedZoneDnssec{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53HostedZoneDnssec)(nil)

// Route53HostedZoneDnssec represents the Terraform resource aws_route53_hosted_zone_dnssec.
type Route53HostedZoneDnssec struct {
	Name      string
	Args      Route53HostedZoneDnssecArgs
	state     *route53HostedZoneDnssecState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) Type() string {
	return "aws_route53_hosted_zone_dnssec"
}

// LocalName returns the local name for [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) LocalName() string {
	return rhzd.Name
}

// Configuration returns the configuration (args) for [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) Configuration() interface{} {
	return rhzd.Args
}

// DependOn is used for other resources to depend on [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) DependOn() terra.Reference {
	return terra.ReferenceResource(rhzd)
}

// Dependencies returns the list of resources [Route53HostedZoneDnssec] depends_on.
func (rhzd *Route53HostedZoneDnssec) Dependencies() terra.Dependencies {
	return rhzd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) LifecycleManagement() *terra.Lifecycle {
	return rhzd.Lifecycle
}

// Attributes returns the attributes for [Route53HostedZoneDnssec].
func (rhzd *Route53HostedZoneDnssec) Attributes() route53HostedZoneDnssecAttributes {
	return route53HostedZoneDnssecAttributes{ref: terra.ReferenceResource(rhzd)}
}

// ImportState imports the given attribute values into [Route53HostedZoneDnssec]'s state.
func (rhzd *Route53HostedZoneDnssec) ImportState(av io.Reader) error {
	rhzd.state = &route53HostedZoneDnssecState{}
	if err := json.NewDecoder(av).Decode(rhzd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rhzd.Type(), rhzd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53HostedZoneDnssec] has state.
func (rhzd *Route53HostedZoneDnssec) State() (*route53HostedZoneDnssecState, bool) {
	return rhzd.state, rhzd.state != nil
}

// StateMust returns the state for [Route53HostedZoneDnssec]. Panics if the state is nil.
func (rhzd *Route53HostedZoneDnssec) StateMust() *route53HostedZoneDnssecState {
	if rhzd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rhzd.Type(), rhzd.LocalName()))
	}
	return rhzd.state
}

// Route53HostedZoneDnssecArgs contains the configurations for aws_route53_hosted_zone_dnssec.
type Route53HostedZoneDnssecArgs struct {
	// HostedZoneId: string, required
	HostedZoneId terra.StringValue `hcl:"hosted_zone_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SigningStatus: string, optional
	SigningStatus terra.StringValue `hcl:"signing_status,attr"`
}
type route53HostedZoneDnssecAttributes struct {
	ref terra.Reference
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_route53_hosted_zone_dnssec.
func (rhzd route53HostedZoneDnssecAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(rhzd.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_route53_hosted_zone_dnssec.
func (rhzd route53HostedZoneDnssecAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rhzd.ref.Append("id"))
}

// SigningStatus returns a reference to field signing_status of aws_route53_hosted_zone_dnssec.
func (rhzd route53HostedZoneDnssecAttributes) SigningStatus() terra.StringValue {
	return terra.ReferenceAsString(rhzd.ref.Append("signing_status"))
}

type route53HostedZoneDnssecState struct {
	HostedZoneId  string `json:"hosted_zone_id"`
	Id            string `json:"id"`
	SigningStatus string `json:"signing_status"`
}
