// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataEcsContainerDefinition creates a new instance of [DataEcsContainerDefinition].
func NewDataEcsContainerDefinition(name string, args DataEcsContainerDefinitionArgs) *DataEcsContainerDefinition {
	return &DataEcsContainerDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcsContainerDefinition)(nil)

// DataEcsContainerDefinition represents the Terraform data resource aws_ecs_container_definition.
type DataEcsContainerDefinition struct {
	Name string
	Args DataEcsContainerDefinitionArgs
}

// DataSource returns the Terraform object type for [DataEcsContainerDefinition].
func (ecd *DataEcsContainerDefinition) DataSource() string {
	return "aws_ecs_container_definition"
}

// LocalName returns the local name for [DataEcsContainerDefinition].
func (ecd *DataEcsContainerDefinition) LocalName() string {
	return ecd.Name
}

// Configuration returns the configuration (args) for [DataEcsContainerDefinition].
func (ecd *DataEcsContainerDefinition) Configuration() interface{} {
	return ecd.Args
}

// Attributes returns the attributes for [DataEcsContainerDefinition].
func (ecd *DataEcsContainerDefinition) Attributes() dataEcsContainerDefinitionAttributes {
	return dataEcsContainerDefinitionAttributes{ref: terra.ReferenceDataResource(ecd)}
}

// DataEcsContainerDefinitionArgs contains the configurations for aws_ecs_container_definition.
type DataEcsContainerDefinitionArgs struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TaskDefinition: string, required
	TaskDefinition terra.StringValue `hcl:"task_definition,attr" validate:"required"`
}
type dataEcsContainerDefinitionAttributes struct {
	ref terra.Reference
}

// ContainerName returns a reference to field container_name of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(ecd.ref.Append("container_name"))
}

// Cpu returns a reference to field cpu of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) Cpu() terra.NumberValue {
	return terra.ReferenceAsNumber(ecd.ref.Append("cpu"))
}

// DisableNetworking returns a reference to field disable_networking of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) DisableNetworking() terra.BoolValue {
	return terra.ReferenceAsBool(ecd.ref.Append("disable_networking"))
}

// DockerLabels returns a reference to field docker_labels of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) DockerLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecd.ref.Append("docker_labels"))
}

// Environment returns a reference to field environment of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) Environment() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecd.ref.Append("environment"))
}

// Id returns a reference to field id of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecd.ref.Append("id"))
}

// Image returns a reference to field image of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ecd.ref.Append("image"))
}

// ImageDigest returns a reference to field image_digest of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) ImageDigest() terra.StringValue {
	return terra.ReferenceAsString(ecd.ref.Append("image_digest"))
}

// Memory returns a reference to field memory of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) Memory() terra.NumberValue {
	return terra.ReferenceAsNumber(ecd.ref.Append("memory"))
}

// MemoryReservation returns a reference to field memory_reservation of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) MemoryReservation() terra.NumberValue {
	return terra.ReferenceAsNumber(ecd.ref.Append("memory_reservation"))
}

// TaskDefinition returns a reference to field task_definition of aws_ecs_container_definition.
func (ecd dataEcsContainerDefinitionAttributes) TaskDefinition() terra.StringValue {
	return terra.ReferenceAsString(ecd.ref.Append("task_definition"))
}
