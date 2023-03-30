// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataBatchComputeEnvironment(name string, args DataBatchComputeEnvironmentArgs) *DataBatchComputeEnvironment {
	return &DataBatchComputeEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBatchComputeEnvironment)(nil)

type DataBatchComputeEnvironment struct {
	Name string
	Args DataBatchComputeEnvironmentArgs
}

func (bce *DataBatchComputeEnvironment) DataSource() string {
	return "aws_batch_compute_environment"
}

func (bce *DataBatchComputeEnvironment) LocalName() string {
	return bce.Name
}

func (bce *DataBatchComputeEnvironment) Configuration() interface{} {
	return bce.Args
}

func (bce *DataBatchComputeEnvironment) Attributes() dataBatchComputeEnvironmentAttributes {
	return dataBatchComputeEnvironmentAttributes{ref: terra.ReferenceDataResource(bce)}
}

type DataBatchComputeEnvironmentArgs struct {
	// ComputeEnvironmentName: string, required
	ComputeEnvironmentName terra.StringValue `hcl:"compute_environment_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataBatchComputeEnvironmentAttributes struct {
	ref terra.Reference
}

func (bce dataBatchComputeEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("arn"))
}

func (bce dataBatchComputeEnvironmentAttributes) ComputeEnvironmentName() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("compute_environment_name"))
}

func (bce dataBatchComputeEnvironmentAttributes) EcsClusterArn() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("ecs_cluster_arn"))
}

func (bce dataBatchComputeEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("id"))
}

func (bce dataBatchComputeEnvironmentAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("service_role"))
}

func (bce dataBatchComputeEnvironmentAttributes) State() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("state"))
}

func (bce dataBatchComputeEnvironmentAttributes) Status() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("status"))
}

func (bce dataBatchComputeEnvironmentAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("status_reason"))
}

func (bce dataBatchComputeEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](bce.ref.Append("tags"))
}

func (bce dataBatchComputeEnvironmentAttributes) Type() terra.StringValue {
	return terra.ReferenceString(bce.ref.Append("type"))
}
