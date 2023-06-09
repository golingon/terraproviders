// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcNetworkPerformanceMetricSubscription creates a new instance of [VpcNetworkPerformanceMetricSubscription].
func NewVpcNetworkPerformanceMetricSubscription(name string, args VpcNetworkPerformanceMetricSubscriptionArgs) *VpcNetworkPerformanceMetricSubscription {
	return &VpcNetworkPerformanceMetricSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcNetworkPerformanceMetricSubscription)(nil)

// VpcNetworkPerformanceMetricSubscription represents the Terraform resource aws_vpc_network_performance_metric_subscription.
type VpcNetworkPerformanceMetricSubscription struct {
	Name      string
	Args      VpcNetworkPerformanceMetricSubscriptionArgs
	state     *vpcNetworkPerformanceMetricSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) Type() string {
	return "aws_vpc_network_performance_metric_subscription"
}

// LocalName returns the local name for [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) LocalName() string {
	return vnpms.Name
}

// Configuration returns the configuration (args) for [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) Configuration() interface{} {
	return vnpms.Args
}

// DependOn is used for other resources to depend on [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(vnpms)
}

// Dependencies returns the list of resources [VpcNetworkPerformanceMetricSubscription] depends_on.
func (vnpms *VpcNetworkPerformanceMetricSubscription) Dependencies() terra.Dependencies {
	return vnpms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) LifecycleManagement() *terra.Lifecycle {
	return vnpms.Lifecycle
}

// Attributes returns the attributes for [VpcNetworkPerformanceMetricSubscription].
func (vnpms *VpcNetworkPerformanceMetricSubscription) Attributes() vpcNetworkPerformanceMetricSubscriptionAttributes {
	return vpcNetworkPerformanceMetricSubscriptionAttributes{ref: terra.ReferenceResource(vnpms)}
}

// ImportState imports the given attribute values into [VpcNetworkPerformanceMetricSubscription]'s state.
func (vnpms *VpcNetworkPerformanceMetricSubscription) ImportState(av io.Reader) error {
	vnpms.state = &vpcNetworkPerformanceMetricSubscriptionState{}
	if err := json.NewDecoder(av).Decode(vnpms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnpms.Type(), vnpms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcNetworkPerformanceMetricSubscription] has state.
func (vnpms *VpcNetworkPerformanceMetricSubscription) State() (*vpcNetworkPerformanceMetricSubscriptionState, bool) {
	return vnpms.state, vnpms.state != nil
}

// StateMust returns the state for [VpcNetworkPerformanceMetricSubscription]. Panics if the state is nil.
func (vnpms *VpcNetworkPerformanceMetricSubscription) StateMust() *vpcNetworkPerformanceMetricSubscriptionState {
	if vnpms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnpms.Type(), vnpms.LocalName()))
	}
	return vnpms.state
}

// VpcNetworkPerformanceMetricSubscriptionArgs contains the configurations for aws_vpc_network_performance_metric_subscription.
type VpcNetworkPerformanceMetricSubscriptionArgs struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metric: string, optional
	Metric terra.StringValue `hcl:"metric,attr"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Statistic: string, optional
	Statistic terra.StringValue `hcl:"statistic,attr"`
}
type vpcNetworkPerformanceMetricSubscriptionAttributes struct {
	ref terra.Reference
}

// Destination returns a reference to field destination of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("destination"))
}

// Id returns a reference to field id of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("id"))
}

// Metric returns a reference to field metric of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Metric() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("metric"))
}

// Period returns a reference to field period of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Period() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("period"))
}

// Source returns a reference to field source of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("source"))
}

// Statistic returns a reference to field statistic of aws_vpc_network_performance_metric_subscription.
func (vnpms vpcNetworkPerformanceMetricSubscriptionAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(vnpms.ref.Append("statistic"))
}

type vpcNetworkPerformanceMetricSubscriptionState struct {
	Destination string `json:"destination"`
	Id          string `json:"id"`
	Metric      string `json:"metric"`
	Period      string `json:"period"`
	Source      string `json:"source"`
	Statistic   string `json:"statistic"`
}
