// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	docdbclusterparametergroup "github.com/golingon/terraproviders/aws/4.60.0/docdbclusterparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocdbClusterParameterGroup creates a new instance of [DocdbClusterParameterGroup].
func NewDocdbClusterParameterGroup(name string, args DocdbClusterParameterGroupArgs) *DocdbClusterParameterGroup {
	return &DocdbClusterParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocdbClusterParameterGroup)(nil)

// DocdbClusterParameterGroup represents the Terraform resource aws_docdb_cluster_parameter_group.
type DocdbClusterParameterGroup struct {
	Name      string
	Args      DocdbClusterParameterGroupArgs
	state     *docdbClusterParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) Type() string {
	return "aws_docdb_cluster_parameter_group"
}

// LocalName returns the local name for [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) LocalName() string {
	return dcpg.Name
}

// Configuration returns the configuration (args) for [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) Configuration() interface{} {
	return dcpg.Args
}

// DependOn is used for other resources to depend on [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dcpg)
}

// Dependencies returns the list of resources [DocdbClusterParameterGroup] depends_on.
func (dcpg *DocdbClusterParameterGroup) Dependencies() terra.Dependencies {
	return dcpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) LifecycleManagement() *terra.Lifecycle {
	return dcpg.Lifecycle
}

// Attributes returns the attributes for [DocdbClusterParameterGroup].
func (dcpg *DocdbClusterParameterGroup) Attributes() docdbClusterParameterGroupAttributes {
	return docdbClusterParameterGroupAttributes{ref: terra.ReferenceResource(dcpg)}
}

// ImportState imports the given attribute values into [DocdbClusterParameterGroup]'s state.
func (dcpg *DocdbClusterParameterGroup) ImportState(av io.Reader) error {
	dcpg.state = &docdbClusterParameterGroupState{}
	if err := json.NewDecoder(av).Decode(dcpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcpg.Type(), dcpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocdbClusterParameterGroup] has state.
func (dcpg *DocdbClusterParameterGroup) State() (*docdbClusterParameterGroupState, bool) {
	return dcpg.state, dcpg.state != nil
}

// StateMust returns the state for [DocdbClusterParameterGroup]. Panics if the state is nil.
func (dcpg *DocdbClusterParameterGroup) StateMust() *docdbClusterParameterGroupState {
	if dcpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcpg.Type(), dcpg.LocalName()))
	}
	return dcpg.state
}

// DocdbClusterParameterGroupArgs contains the configurations for aws_docdb_cluster_parameter_group.
type DocdbClusterParameterGroupArgs struct {
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
	Parameter []docdbclusterparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type docdbClusterParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(dcpg.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcpg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_docdb_cluster_parameter_group.
func (dcpg docdbClusterParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcpg.ref.Append("tags_all"))
}

func (dcpg docdbClusterParameterGroupAttributes) Parameter() terra.SetValue[docdbclusterparametergroup.ParameterAttributes] {
	return terra.ReferenceAsSet[docdbclusterparametergroup.ParameterAttributes](dcpg.ref.Append("parameter"))
}

type docdbClusterParameterGroupState struct {
	Arn         string                                      `json:"arn"`
	Description string                                      `json:"description"`
	Family      string                                      `json:"family"`
	Id          string                                      `json:"id"`
	Name        string                                      `json:"name"`
	NamePrefix  string                                      `json:"name_prefix"`
	Tags        map[string]string                           `json:"tags"`
	TagsAll     map[string]string                           `json:"tags_all"`
	Parameter   []docdbclusterparametergroup.ParameterState `json:"parameter"`
}
