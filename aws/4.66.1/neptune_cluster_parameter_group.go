// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptuneclusterparametergroup "github.com/golingon/terraproviders/aws/4.66.1/neptuneclusterparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNeptuneClusterParameterGroup creates a new instance of [NeptuneClusterParameterGroup].
func NewNeptuneClusterParameterGroup(name string, args NeptuneClusterParameterGroupArgs) *NeptuneClusterParameterGroup {
	return &NeptuneClusterParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneClusterParameterGroup)(nil)

// NeptuneClusterParameterGroup represents the Terraform resource aws_neptune_cluster_parameter_group.
type NeptuneClusterParameterGroup struct {
	Name      string
	Args      NeptuneClusterParameterGroupArgs
	state     *neptuneClusterParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) Type() string {
	return "aws_neptune_cluster_parameter_group"
}

// LocalName returns the local name for [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) LocalName() string {
	return ncpg.Name
}

// Configuration returns the configuration (args) for [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) Configuration() interface{} {
	return ncpg.Args
}

// DependOn is used for other resources to depend on [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ncpg)
}

// Dependencies returns the list of resources [NeptuneClusterParameterGroup] depends_on.
func (ncpg *NeptuneClusterParameterGroup) Dependencies() terra.Dependencies {
	return ncpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) LifecycleManagement() *terra.Lifecycle {
	return ncpg.Lifecycle
}

// Attributes returns the attributes for [NeptuneClusterParameterGroup].
func (ncpg *NeptuneClusterParameterGroup) Attributes() neptuneClusterParameterGroupAttributes {
	return neptuneClusterParameterGroupAttributes{ref: terra.ReferenceResource(ncpg)}
}

// ImportState imports the given attribute values into [NeptuneClusterParameterGroup]'s state.
func (ncpg *NeptuneClusterParameterGroup) ImportState(av io.Reader) error {
	ncpg.state = &neptuneClusterParameterGroupState{}
	if err := json.NewDecoder(av).Decode(ncpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncpg.Type(), ncpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneClusterParameterGroup] has state.
func (ncpg *NeptuneClusterParameterGroup) State() (*neptuneClusterParameterGroupState, bool) {
	return ncpg.state, ncpg.state != nil
}

// StateMust returns the state for [NeptuneClusterParameterGroup]. Panics if the state is nil.
func (ncpg *NeptuneClusterParameterGroup) StateMust() *neptuneClusterParameterGroupState {
	if ncpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncpg.Type(), ncpg.LocalName()))
	}
	return ncpg.state
}

// NeptuneClusterParameterGroupArgs contains the configurations for aws_neptune_cluster_parameter_group.
type NeptuneClusterParameterGroupArgs struct {
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
	Parameter []neptuneclusterparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type neptuneClusterParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ncpg.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncpg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_cluster_parameter_group.
func (ncpg neptuneClusterParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncpg.ref.Append("tags_all"))
}

func (ncpg neptuneClusterParameterGroupAttributes) Parameter() terra.SetValue[neptuneclusterparametergroup.ParameterAttributes] {
	return terra.ReferenceAsSet[neptuneclusterparametergroup.ParameterAttributes](ncpg.ref.Append("parameter"))
}

type neptuneClusterParameterGroupState struct {
	Arn         string                                        `json:"arn"`
	Description string                                        `json:"description"`
	Family      string                                        `json:"family"`
	Id          string                                        `json:"id"`
	Name        string                                        `json:"name"`
	NamePrefix  string                                        `json:"name_prefix"`
	Tags        map[string]string                             `json:"tags"`
	TagsAll     map[string]string                             `json:"tags_all"`
	Parameter   []neptuneclusterparametergroup.ParameterState `json:"parameter"`
}