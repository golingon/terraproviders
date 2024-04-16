// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_docdb_cluster_parameter_group

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_docdb_cluster_parameter_group.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDocdbClusterParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adcpg *Resource) Type() string {
	return "aws_docdb_cluster_parameter_group"
}

// LocalName returns the local name for [Resource].
func (adcpg *Resource) LocalName() string {
	return adcpg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adcpg *Resource) Configuration() interface{} {
	return adcpg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adcpg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adcpg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adcpg *Resource) Dependencies() terra.Dependencies {
	return adcpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adcpg *Resource) LifecycleManagement() *terra.Lifecycle {
	return adcpg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adcpg *Resource) Attributes() awsDocdbClusterParameterGroupAttributes {
	return awsDocdbClusterParameterGroupAttributes{ref: terra.ReferenceResource(adcpg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adcpg *Resource) ImportState(state io.Reader) error {
	adcpg.state = &awsDocdbClusterParameterGroupState{}
	if err := json.NewDecoder(state).Decode(adcpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adcpg.Type(), adcpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adcpg *Resource) State() (*awsDocdbClusterParameterGroupState, bool) {
	return adcpg.state, adcpg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adcpg *Resource) StateMust() *awsDocdbClusterParameterGroupState {
	if adcpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adcpg.Type(), adcpg.LocalName()))
	}
	return adcpg.state
}

// Args contains the configurations for aws_docdb_cluster_parameter_group.
type Args struct {
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
	Parameter []Parameter `hcl:"parameter,block" validate:"min=0"`
}

type awsDocdbClusterParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(adcpg.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adcpg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_docdb_cluster_parameter_group.
func (adcpg awsDocdbClusterParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adcpg.ref.Append("tags_all"))
}

func (adcpg awsDocdbClusterParameterGroupAttributes) Parameter() terra.SetValue[ParameterAttributes] {
	return terra.ReferenceAsSet[ParameterAttributes](adcpg.ref.Append("parameter"))
}

type awsDocdbClusterParameterGroupState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Family      string            `json:"family"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
	Parameter   []ParameterState  `json:"parameter"`
}
