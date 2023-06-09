// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeAccessLogSubscription creates a new instance of [VpclatticeAccessLogSubscription].
func NewVpclatticeAccessLogSubscription(name string, args VpclatticeAccessLogSubscriptionArgs) *VpclatticeAccessLogSubscription {
	return &VpclatticeAccessLogSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeAccessLogSubscription)(nil)

// VpclatticeAccessLogSubscription represents the Terraform resource aws_vpclattice_access_log_subscription.
type VpclatticeAccessLogSubscription struct {
	Name      string
	Args      VpclatticeAccessLogSubscriptionArgs
	state     *vpclatticeAccessLogSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) Type() string {
	return "aws_vpclattice_access_log_subscription"
}

// LocalName returns the local name for [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) LocalName() string {
	return vals.Name
}

// Configuration returns the configuration (args) for [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) Configuration() interface{} {
	return vals.Args
}

// DependOn is used for other resources to depend on [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(vals)
}

// Dependencies returns the list of resources [VpclatticeAccessLogSubscription] depends_on.
func (vals *VpclatticeAccessLogSubscription) Dependencies() terra.Dependencies {
	return vals.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) LifecycleManagement() *terra.Lifecycle {
	return vals.Lifecycle
}

// Attributes returns the attributes for [VpclatticeAccessLogSubscription].
func (vals *VpclatticeAccessLogSubscription) Attributes() vpclatticeAccessLogSubscriptionAttributes {
	return vpclatticeAccessLogSubscriptionAttributes{ref: terra.ReferenceResource(vals)}
}

// ImportState imports the given attribute values into [VpclatticeAccessLogSubscription]'s state.
func (vals *VpclatticeAccessLogSubscription) ImportState(av io.Reader) error {
	vals.state = &vpclatticeAccessLogSubscriptionState{}
	if err := json.NewDecoder(av).Decode(vals.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vals.Type(), vals.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeAccessLogSubscription] has state.
func (vals *VpclatticeAccessLogSubscription) State() (*vpclatticeAccessLogSubscriptionState, bool) {
	return vals.state, vals.state != nil
}

// StateMust returns the state for [VpclatticeAccessLogSubscription]. Panics if the state is nil.
func (vals *VpclatticeAccessLogSubscription) StateMust() *vpclatticeAccessLogSubscriptionState {
	if vals.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vals.Type(), vals.LocalName()))
	}
	return vals.state
}

// VpclatticeAccessLogSubscriptionArgs contains the configurations for aws_vpclattice_access_log_subscription.
type VpclatticeAccessLogSubscriptionArgs struct {
	// DestinationArn: string, required
	DestinationArn terra.StringValue `hcl:"destination_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceIdentifier: string, required
	ResourceIdentifier terra.StringValue `hcl:"resource_identifier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type vpclatticeAccessLogSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vals.ref.Append("arn"))
}

// DestinationArn returns a reference to field destination_arn of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) DestinationArn() terra.StringValue {
	return terra.ReferenceAsString(vals.ref.Append("destination_arn"))
}

// Id returns a reference to field id of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vals.ref.Append("id"))
}

// ResourceArn returns a reference to field resource_arn of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(vals.ref.Append("resource_arn"))
}

// ResourceIdentifier returns a reference to field resource_identifier of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) ResourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vals.ref.Append("resource_identifier"))
}

// Tags returns a reference to field tags of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vals.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpclattice_access_log_subscription.
func (vals vpclatticeAccessLogSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vals.ref.Append("tags_all"))
}

type vpclatticeAccessLogSubscriptionState struct {
	Arn                string            `json:"arn"`
	DestinationArn     string            `json:"destination_arn"`
	Id                 string            `json:"id"`
	ResourceArn        string            `json:"resource_arn"`
	ResourceIdentifier string            `json:"resource_identifier"`
	Tags               map[string]string `json:"tags"`
	TagsAll            map[string]string `json:"tags_all"`
}
