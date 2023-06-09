// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53recoveryreadinessreadinesscheck "github.com/golingon/terraproviders/aws/4.63.0/route53recoveryreadinessreadinesscheck"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53RecoveryreadinessReadinessCheck creates a new instance of [Route53RecoveryreadinessReadinessCheck].
func NewRoute53RecoveryreadinessReadinessCheck(name string, args Route53RecoveryreadinessReadinessCheckArgs) *Route53RecoveryreadinessReadinessCheck {
	return &Route53RecoveryreadinessReadinessCheck{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53RecoveryreadinessReadinessCheck)(nil)

// Route53RecoveryreadinessReadinessCheck represents the Terraform resource aws_route53recoveryreadiness_readiness_check.
type Route53RecoveryreadinessReadinessCheck struct {
	Name      string
	Args      Route53RecoveryreadinessReadinessCheckArgs
	state     *route53RecoveryreadinessReadinessCheckState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) Type() string {
	return "aws_route53recoveryreadiness_readiness_check"
}

// LocalName returns the local name for [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) LocalName() string {
	return rrc.Name
}

// Configuration returns the configuration (args) for [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) Configuration() interface{} {
	return rrc.Args
}

// DependOn is used for other resources to depend on [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) DependOn() terra.Reference {
	return terra.ReferenceResource(rrc)
}

// Dependencies returns the list of resources [Route53RecoveryreadinessReadinessCheck] depends_on.
func (rrc *Route53RecoveryreadinessReadinessCheck) Dependencies() terra.Dependencies {
	return rrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) LifecycleManagement() *terra.Lifecycle {
	return rrc.Lifecycle
}

// Attributes returns the attributes for [Route53RecoveryreadinessReadinessCheck].
func (rrc *Route53RecoveryreadinessReadinessCheck) Attributes() route53RecoveryreadinessReadinessCheckAttributes {
	return route53RecoveryreadinessReadinessCheckAttributes{ref: terra.ReferenceResource(rrc)}
}

// ImportState imports the given attribute values into [Route53RecoveryreadinessReadinessCheck]'s state.
func (rrc *Route53RecoveryreadinessReadinessCheck) ImportState(av io.Reader) error {
	rrc.state = &route53RecoveryreadinessReadinessCheckState{}
	if err := json.NewDecoder(av).Decode(rrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrc.Type(), rrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53RecoveryreadinessReadinessCheck] has state.
func (rrc *Route53RecoveryreadinessReadinessCheck) State() (*route53RecoveryreadinessReadinessCheckState, bool) {
	return rrc.state, rrc.state != nil
}

// StateMust returns the state for [Route53RecoveryreadinessReadinessCheck]. Panics if the state is nil.
func (rrc *Route53RecoveryreadinessReadinessCheck) StateMust() *route53RecoveryreadinessReadinessCheckState {
	if rrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrc.Type(), rrc.LocalName()))
	}
	return rrc.state
}

// Route53RecoveryreadinessReadinessCheckArgs contains the configurations for aws_route53recoveryreadiness_readiness_check.
type Route53RecoveryreadinessReadinessCheckArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReadinessCheckName: string, required
	ReadinessCheckName terra.StringValue `hcl:"readiness_check_name,attr" validate:"required"`
	// ResourceSetName: string, required
	ResourceSetName terra.StringValue `hcl:"resource_set_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *route53recoveryreadinessreadinesscheck.Timeouts `hcl:"timeouts,block"`
}
type route53RecoveryreadinessReadinessCheckAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrc.ref.Append("id"))
}

// ReadinessCheckName returns a reference to field readiness_check_name of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) ReadinessCheckName() terra.StringValue {
	return terra.ReferenceAsString(rrc.ref.Append("readiness_check_name"))
}

// ResourceSetName returns a reference to field resource_set_name of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) ResourceSetName() terra.StringValue {
	return terra.ReferenceAsString(rrc.ref.Append("resource_set_name"))
}

// Tags returns a reference to field tags of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53recoveryreadiness_readiness_check.
func (rrc route53RecoveryreadinessReadinessCheckAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrc.ref.Append("tags_all"))
}

func (rrc route53RecoveryreadinessReadinessCheckAttributes) Timeouts() route53recoveryreadinessreadinesscheck.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53recoveryreadinessreadinesscheck.TimeoutsAttributes](rrc.ref.Append("timeouts"))
}

type route53RecoveryreadinessReadinessCheckState struct {
	Arn                string                                                `json:"arn"`
	Id                 string                                                `json:"id"`
	ReadinessCheckName string                                                `json:"readiness_check_name"`
	ResourceSetName    string                                                `json:"resource_set_name"`
	Tags               map[string]string                                     `json:"tags"`
	TagsAll            map[string]string                                     `json:"tags_all"`
	Timeouts           *route53recoveryreadinessreadinesscheck.TimeoutsState `json:"timeouts"`
}
