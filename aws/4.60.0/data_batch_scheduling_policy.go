// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	databatchschedulingpolicy "github.com/golingon/terraproviders/aws/4.60.0/databatchschedulingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBatchSchedulingPolicy creates a new instance of [DataBatchSchedulingPolicy].
func NewDataBatchSchedulingPolicy(name string, args DataBatchSchedulingPolicyArgs) *DataBatchSchedulingPolicy {
	return &DataBatchSchedulingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchSchedulingPolicy)(nil)

// DataBatchSchedulingPolicy represents the Terraform data resource aws_batch_scheduling_policy.
type DataBatchSchedulingPolicy struct {
	Name string
	Args DataBatchSchedulingPolicyArgs
}

// DataSource returns the Terraform object type for [DataBatchSchedulingPolicy].
func (bsp *DataBatchSchedulingPolicy) DataSource() string {
	return "aws_batch_scheduling_policy"
}

// LocalName returns the local name for [DataBatchSchedulingPolicy].
func (bsp *DataBatchSchedulingPolicy) LocalName() string {
	return bsp.Name
}

// Configuration returns the configuration (args) for [DataBatchSchedulingPolicy].
func (bsp *DataBatchSchedulingPolicy) Configuration() interface{} {
	return bsp.Args
}

// Attributes returns the attributes for [DataBatchSchedulingPolicy].
func (bsp *DataBatchSchedulingPolicy) Attributes() dataBatchSchedulingPolicyAttributes {
	return dataBatchSchedulingPolicyAttributes{ref: terra.ReferenceDataResource(bsp)}
}

// DataBatchSchedulingPolicyArgs contains the configurations for aws_batch_scheduling_policy.
type DataBatchSchedulingPolicyArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// FairSharePolicy: min=0
	FairSharePolicy []databatchschedulingpolicy.FairSharePolicy `hcl:"fair_share_policy,block" validate:"min=0"`
}
type dataBatchSchedulingPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_batch_scheduling_policy.
func (bsp dataBatchSchedulingPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_batch_scheduling_policy.
func (bsp dataBatchSchedulingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("id"))
}

// Name returns a reference to field name of aws_batch_scheduling_policy.
func (bsp dataBatchSchedulingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_batch_scheduling_policy.
func (bsp dataBatchSchedulingPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bsp.ref.Append("tags"))
}

func (bsp dataBatchSchedulingPolicyAttributes) FairSharePolicy() terra.ListValue[databatchschedulingpolicy.FairSharePolicyAttributes] {
	return terra.ReferenceAsList[databatchschedulingpolicy.FairSharePolicyAttributes](bsp.ref.Append("fair_share_policy"))
}
