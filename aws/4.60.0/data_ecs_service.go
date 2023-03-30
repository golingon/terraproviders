// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataEcsService(name string, args DataEcsServiceArgs) *DataEcsService {
	return &DataEcsService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcsService)(nil)

type DataEcsService struct {
	Name string
	Args DataEcsServiceArgs
}

func (es *DataEcsService) DataSource() string {
	return "aws_ecs_service"
}

func (es *DataEcsService) LocalName() string {
	return es.Name
}

func (es *DataEcsService) Configuration() interface{} {
	return es.Args
}

func (es *DataEcsService) Attributes() dataEcsServiceAttributes {
	return dataEcsServiceAttributes{ref: terra.ReferenceDataResource(es)}
}

type DataEcsServiceArgs struct {
	// ClusterArn: string, required
	ClusterArn terra.StringValue `hcl:"cluster_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataEcsServiceAttributes struct {
	ref terra.Reference
}

func (es dataEcsServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("arn"))
}

func (es dataEcsServiceAttributes) ClusterArn() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("cluster_arn"))
}

func (es dataEcsServiceAttributes) DesiredCount() terra.NumberValue {
	return terra.ReferenceNumber(es.ref.Append("desired_count"))
}

func (es dataEcsServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("id"))
}

func (es dataEcsServiceAttributes) LaunchType() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("launch_type"))
}

func (es dataEcsServiceAttributes) SchedulingStrategy() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("scheduling_strategy"))
}

func (es dataEcsServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("service_name"))
}

func (es dataEcsServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](es.ref.Append("tags"))
}

func (es dataEcsServiceAttributes) TaskDefinition() terra.StringValue {
	return terra.ReferenceString(es.ref.Append("task_definition"))
}
