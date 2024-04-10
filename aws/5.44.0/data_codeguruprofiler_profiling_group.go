// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datacodeguruprofilerprofilinggroup "github.com/golingon/terraproviders/aws/5.44.0/datacodeguruprofilerprofilinggroup"
)

// NewDataCodeguruprofilerProfilingGroup creates a new instance of [DataCodeguruprofilerProfilingGroup].
func NewDataCodeguruprofilerProfilingGroup(name string, args DataCodeguruprofilerProfilingGroupArgs) *DataCodeguruprofilerProfilingGroup {
	return &DataCodeguruprofilerProfilingGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCodeguruprofilerProfilingGroup)(nil)

// DataCodeguruprofilerProfilingGroup represents the Terraform data resource aws_codeguruprofiler_profiling_group.
type DataCodeguruprofilerProfilingGroup struct {
	Name string
	Args DataCodeguruprofilerProfilingGroupArgs
}

// DataSource returns the Terraform object type for [DataCodeguruprofilerProfilingGroup].
func (cpg *DataCodeguruprofilerProfilingGroup) DataSource() string {
	return "aws_codeguruprofiler_profiling_group"
}

// LocalName returns the local name for [DataCodeguruprofilerProfilingGroup].
func (cpg *DataCodeguruprofilerProfilingGroup) LocalName() string {
	return cpg.Name
}

// Configuration returns the configuration (args) for [DataCodeguruprofilerProfilingGroup].
func (cpg *DataCodeguruprofilerProfilingGroup) Configuration() interface{} {
	return cpg.Args
}

// Attributes returns the attributes for [DataCodeguruprofilerProfilingGroup].
func (cpg *DataCodeguruprofilerProfilingGroup) Attributes() dataCodeguruprofilerProfilingGroupAttributes {
	return dataCodeguruprofilerProfilingGroupAttributes{ref: terra.ReferenceDataResource(cpg)}
}

// DataCodeguruprofilerProfilingGroupArgs contains the configurations for aws_codeguruprofiler_profiling_group.
type DataCodeguruprofilerProfilingGroupArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// AgentOrchestrationConfig: min=0
	AgentOrchestrationConfig []datacodeguruprofilerprofilinggroup.AgentOrchestrationConfig `hcl:"agent_orchestration_config,block" validate:"min=0"`
	// ProfilingStatus: min=0
	ProfilingStatus []datacodeguruprofilerprofilinggroup.ProfilingStatus `hcl:"profiling_status,block" validate:"min=0"`
}
type dataCodeguruprofilerProfilingGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("arn"))
}

// ComputePlatform returns a reference to field compute_platform of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) ComputePlatform() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("compute_platform"))
}

// CreatedAt returns a reference to field created_at of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cpg.ref.Append("tags"))
}

// UpdatedAt returns a reference to field updated_at of aws_codeguruprofiler_profiling_group.
func (cpg dataCodeguruprofilerProfilingGroupAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(cpg.ref.Append("updated_at"))
}

func (cpg dataCodeguruprofilerProfilingGroupAttributes) AgentOrchestrationConfig() terra.ListValue[datacodeguruprofilerprofilinggroup.AgentOrchestrationConfigAttributes] {
	return terra.ReferenceAsList[datacodeguruprofilerprofilinggroup.AgentOrchestrationConfigAttributes](cpg.ref.Append("agent_orchestration_config"))
}

func (cpg dataCodeguruprofilerProfilingGroupAttributes) ProfilingStatus() terra.ListValue[datacodeguruprofilerprofilinggroup.ProfilingStatusAttributes] {
	return terra.ReferenceAsList[datacodeguruprofilerprofilinggroup.ProfilingStatusAttributes](cpg.ref.Append("profiling_status"))
}
