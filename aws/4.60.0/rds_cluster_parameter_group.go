// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rdsclusterparametergroup "github.com/golingon/terraproviders/aws/4.60.0/rdsclusterparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRdsClusterParameterGroup creates a new instance of [RdsClusterParameterGroup].
func NewRdsClusterParameterGroup(name string, args RdsClusterParameterGroupArgs) *RdsClusterParameterGroup {
	return &RdsClusterParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsClusterParameterGroup)(nil)

// RdsClusterParameterGroup represents the Terraform resource aws_rds_cluster_parameter_group.
type RdsClusterParameterGroup struct {
	Name      string
	Args      RdsClusterParameterGroupArgs
	state     *rdsClusterParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) Type() string {
	return "aws_rds_cluster_parameter_group"
}

// LocalName returns the local name for [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) LocalName() string {
	return rcpg.Name
}

// Configuration returns the configuration (args) for [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) Configuration() interface{} {
	return rcpg.Args
}

// DependOn is used for other resources to depend on [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(rcpg)
}

// Dependencies returns the list of resources [RdsClusterParameterGroup] depends_on.
func (rcpg *RdsClusterParameterGroup) Dependencies() terra.Dependencies {
	return rcpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) LifecycleManagement() *terra.Lifecycle {
	return rcpg.Lifecycle
}

// Attributes returns the attributes for [RdsClusterParameterGroup].
func (rcpg *RdsClusterParameterGroup) Attributes() rdsClusterParameterGroupAttributes {
	return rdsClusterParameterGroupAttributes{ref: terra.ReferenceResource(rcpg)}
}

// ImportState imports the given attribute values into [RdsClusterParameterGroup]'s state.
func (rcpg *RdsClusterParameterGroup) ImportState(av io.Reader) error {
	rcpg.state = &rdsClusterParameterGroupState{}
	if err := json.NewDecoder(av).Decode(rcpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcpg.Type(), rcpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RdsClusterParameterGroup] has state.
func (rcpg *RdsClusterParameterGroup) State() (*rdsClusterParameterGroupState, bool) {
	return rcpg.state, rcpg.state != nil
}

// StateMust returns the state for [RdsClusterParameterGroup]. Panics if the state is nil.
func (rcpg *RdsClusterParameterGroup) StateMust() *rdsClusterParameterGroupState {
	if rcpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcpg.Type(), rcpg.LocalName()))
	}
	return rcpg.state
}

// RdsClusterParameterGroupArgs contains the configurations for aws_rds_cluster_parameter_group.
type RdsClusterParameterGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Family: string, required
	Family terra.StringValue `hcl:"family,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Parameter: min=0
	Parameter []rdsclusterparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type rdsClusterParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(rcpg.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rcpg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rds_cluster_parameter_group.
func (rcpg rdsClusterParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rcpg.ref.Append("tags_all"))
}

func (rcpg rdsClusterParameterGroupAttributes) Parameter() terra.SetValue[rdsclusterparametergroup.ParameterAttributes] {
	return terra.ReferenceAsSet[rdsclusterparametergroup.ParameterAttributes](rcpg.ref.Append("parameter"))
}

type rdsClusterParameterGroupState struct {
	Arn         string                                    `json:"arn"`
	Description string                                    `json:"description"`
	Family      string                                    `json:"family"`
	Id          string                                    `json:"id"`
	Name        string                                    `json:"name"`
	NamePrefix  string                                    `json:"name_prefix"`
	Tags        map[string]string                         `json:"tags"`
	TagsAll     map[string]string                         `json:"tags_all"`
	Parameter   []rdsclusterparametergroup.ParameterState `json:"parameter"`
}
