// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53ResolverDnssecConfig creates a new instance of [Route53ResolverDnssecConfig].
func NewRoute53ResolverDnssecConfig(name string, args Route53ResolverDnssecConfigArgs) *Route53ResolverDnssecConfig {
	return &Route53ResolverDnssecConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverDnssecConfig)(nil)

// Route53ResolverDnssecConfig represents the Terraform resource aws_route53_resolver_dnssec_config.
type Route53ResolverDnssecConfig struct {
	Name      string
	Args      Route53ResolverDnssecConfigArgs
	state     *route53ResolverDnssecConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) Type() string {
	return "aws_route53_resolver_dnssec_config"
}

// LocalName returns the local name for [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) LocalName() string {
	return rrdc.Name
}

// Configuration returns the configuration (args) for [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) Configuration() interface{} {
	return rrdc.Args
}

// DependOn is used for other resources to depend on [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(rrdc)
}

// Dependencies returns the list of resources [Route53ResolverDnssecConfig] depends_on.
func (rrdc *Route53ResolverDnssecConfig) Dependencies() terra.Dependencies {
	return rrdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) LifecycleManagement() *terra.Lifecycle {
	return rrdc.Lifecycle
}

// Attributes returns the attributes for [Route53ResolverDnssecConfig].
func (rrdc *Route53ResolverDnssecConfig) Attributes() route53ResolverDnssecConfigAttributes {
	return route53ResolverDnssecConfigAttributes{ref: terra.ReferenceResource(rrdc)}
}

// ImportState imports the given attribute values into [Route53ResolverDnssecConfig]'s state.
func (rrdc *Route53ResolverDnssecConfig) ImportState(av io.Reader) error {
	rrdc.state = &route53ResolverDnssecConfigState{}
	if err := json.NewDecoder(av).Decode(rrdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrdc.Type(), rrdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53ResolverDnssecConfig] has state.
func (rrdc *Route53ResolverDnssecConfig) State() (*route53ResolverDnssecConfigState, bool) {
	return rrdc.state, rrdc.state != nil
}

// StateMust returns the state for [Route53ResolverDnssecConfig]. Panics if the state is nil.
func (rrdc *Route53ResolverDnssecConfig) StateMust() *route53ResolverDnssecConfigState {
	if rrdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrdc.Type(), rrdc.LocalName()))
	}
	return rrdc.state
}

// Route53ResolverDnssecConfigArgs contains the configurations for aws_route53_resolver_dnssec_config.
type Route53ResolverDnssecConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
}
type route53ResolverDnssecConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_dnssec_config.
func (rrdc route53ResolverDnssecConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrdc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53_resolver_dnssec_config.
func (rrdc route53ResolverDnssecConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrdc.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_dnssec_config.
func (rrdc route53ResolverDnssecConfigAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rrdc.ref.Append("owner_id"))
}

// ResourceId returns a reference to field resource_id of aws_route53_resolver_dnssec_config.
func (rrdc route53ResolverDnssecConfigAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(rrdc.ref.Append("resource_id"))
}

// ValidationStatus returns a reference to field validation_status of aws_route53_resolver_dnssec_config.
func (rrdc route53ResolverDnssecConfigAttributes) ValidationStatus() terra.StringValue {
	return terra.ReferenceAsString(rrdc.ref.Append("validation_status"))
}

type route53ResolverDnssecConfigState struct {
	Arn              string `json:"arn"`
	Id               string `json:"id"`
	OwnerId          string `json:"owner_id"`
	ResourceId       string `json:"resource_id"`
	ValidationStatus string `json:"validation_status"`
}