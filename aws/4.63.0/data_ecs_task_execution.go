// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataecstaskexecution "github.com/golingon/terraproviders/aws/4.63.0/dataecstaskexecution"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEcsTaskExecution creates a new instance of [DataEcsTaskExecution].
func NewDataEcsTaskExecution(name string, args DataEcsTaskExecutionArgs) *DataEcsTaskExecution {
	return &DataEcsTaskExecution{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEcsTaskExecution)(nil)

// DataEcsTaskExecution represents the Terraform data resource aws_ecs_task_execution.
type DataEcsTaskExecution struct {
	Name string
	Args DataEcsTaskExecutionArgs
}

// DataSource returns the Terraform object type for [DataEcsTaskExecution].
func (ete *DataEcsTaskExecution) DataSource() string {
	return "aws_ecs_task_execution"
}

// LocalName returns the local name for [DataEcsTaskExecution].
func (ete *DataEcsTaskExecution) LocalName() string {
	return ete.Name
}

// Configuration returns the configuration (args) for [DataEcsTaskExecution].
func (ete *DataEcsTaskExecution) Configuration() interface{} {
	return ete.Args
}

// Attributes returns the attributes for [DataEcsTaskExecution].
func (ete *DataEcsTaskExecution) Attributes() dataEcsTaskExecutionAttributes {
	return dataEcsTaskExecutionAttributes{ref: terra.ReferenceDataResource(ete)}
}

// DataEcsTaskExecutionArgs contains the configurations for aws_ecs_task_execution.
type DataEcsTaskExecutionArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// DesiredCount: number, optional
	DesiredCount terra.NumberValue `hcl:"desired_count,attr"`
	// EnableEcsManagedTags: bool, optional
	EnableEcsManagedTags terra.BoolValue `hcl:"enable_ecs_managed_tags,attr"`
	// EnableExecuteCommand: bool, optional
	EnableExecuteCommand terra.BoolValue `hcl:"enable_execute_command,attr"`
	// Group: string, optional
	Group terra.StringValue `hcl:"group,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LaunchType: string, optional
	LaunchType terra.StringValue `hcl:"launch_type,attr"`
	// PlatformVersion: string, optional
	PlatformVersion terra.StringValue `hcl:"platform_version,attr"`
	// PropagateTags: string, optional
	PropagateTags terra.StringValue `hcl:"propagate_tags,attr"`
	// ReferenceId: string, optional
	ReferenceId terra.StringValue `hcl:"reference_id,attr"`
	// StartedBy: string, optional
	StartedBy terra.StringValue `hcl:"started_by,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TaskDefinition: string, required
	TaskDefinition terra.StringValue `hcl:"task_definition,attr" validate:"required"`
	// CapacityProviderStrategy: min=0
	CapacityProviderStrategy []dataecstaskexecution.CapacityProviderStrategy `hcl:"capacity_provider_strategy,block" validate:"min=0"`
	// NetworkConfiguration: optional
	NetworkConfiguration *dataecstaskexecution.NetworkConfiguration `hcl:"network_configuration,block"`
	// Overrides: optional
	Overrides *dataecstaskexecution.Overrides `hcl:"overrides,block"`
	// PlacementConstraints: min=0,max=10
	PlacementConstraints []dataecstaskexecution.PlacementConstraints `hcl:"placement_constraints,block" validate:"min=0,max=10"`
	// PlacementStrategy: min=0,max=5
	PlacementStrategy []dataecstaskexecution.PlacementStrategy `hcl:"placement_strategy,block" validate:"min=0,max=5"`
}
type dataEcsTaskExecutionAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("cluster"))
}

// DesiredCount returns a reference to field desired_count of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) DesiredCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ete.ref.Append("desired_count"))
}

// EnableEcsManagedTags returns a reference to field enable_ecs_managed_tags of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) EnableEcsManagedTags() terra.BoolValue {
	return terra.ReferenceAsBool(ete.ref.Append("enable_ecs_managed_tags"))
}

// EnableExecuteCommand returns a reference to field enable_execute_command of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) EnableExecuteCommand() terra.BoolValue {
	return terra.ReferenceAsBool(ete.ref.Append("enable_execute_command"))
}

// Group returns a reference to field group of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("group"))
}

// Id returns a reference to field id of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("id"))
}

// LaunchType returns a reference to field launch_type of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) LaunchType() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("launch_type"))
}

// PlatformVersion returns a reference to field platform_version of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("platform_version"))
}

// PropagateTags returns a reference to field propagate_tags of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) PropagateTags() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("propagate_tags"))
}

// ReferenceId returns a reference to field reference_id of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) ReferenceId() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("reference_id"))
}

// StartedBy returns a reference to field started_by of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) StartedBy() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("started_by"))
}

// Tags returns a reference to field tags of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ete.ref.Append("tags"))
}

// TaskArns returns a reference to field task_arns of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) TaskArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ete.ref.Append("task_arns"))
}

// TaskDefinition returns a reference to field task_definition of aws_ecs_task_execution.
func (ete dataEcsTaskExecutionAttributes) TaskDefinition() terra.StringValue {
	return terra.ReferenceAsString(ete.ref.Append("task_definition"))
}

func (ete dataEcsTaskExecutionAttributes) CapacityProviderStrategy() terra.SetValue[dataecstaskexecution.CapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[dataecstaskexecution.CapacityProviderStrategyAttributes](ete.ref.Append("capacity_provider_strategy"))
}

func (ete dataEcsTaskExecutionAttributes) NetworkConfiguration() terra.ListValue[dataecstaskexecution.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[dataecstaskexecution.NetworkConfigurationAttributes](ete.ref.Append("network_configuration"))
}

func (ete dataEcsTaskExecutionAttributes) Overrides() terra.ListValue[dataecstaskexecution.OverridesAttributes] {
	return terra.ReferenceAsList[dataecstaskexecution.OverridesAttributes](ete.ref.Append("overrides"))
}

func (ete dataEcsTaskExecutionAttributes) PlacementConstraints() terra.SetValue[dataecstaskexecution.PlacementConstraintsAttributes] {
	return terra.ReferenceAsSet[dataecstaskexecution.PlacementConstraintsAttributes](ete.ref.Append("placement_constraints"))
}

func (ete dataEcsTaskExecutionAttributes) PlacementStrategy() terra.ListValue[dataecstaskexecution.PlacementStrategyAttributes] {
	return terra.ReferenceAsList[dataecstaskexecution.PlacementStrategyAttributes](ete.ref.Append("placement_strategy"))
}
