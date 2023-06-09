// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dbparametergroup "github.com/golingon/terraproviders/aws/5.6.2/dbparametergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbParameterGroup creates a new instance of [DbParameterGroup].
func NewDbParameterGroup(name string, args DbParameterGroupArgs) *DbParameterGroup {
	return &DbParameterGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbParameterGroup)(nil)

// DbParameterGroup represents the Terraform resource aws_db_parameter_group.
type DbParameterGroup struct {
	Name      string
	Args      DbParameterGroupArgs
	state     *dbParameterGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbParameterGroup].
func (dpg *DbParameterGroup) Type() string {
	return "aws_db_parameter_group"
}

// LocalName returns the local name for [DbParameterGroup].
func (dpg *DbParameterGroup) LocalName() string {
	return dpg.Name
}

// Configuration returns the configuration (args) for [DbParameterGroup].
func (dpg *DbParameterGroup) Configuration() interface{} {
	return dpg.Args
}

// DependOn is used for other resources to depend on [DbParameterGroup].
func (dpg *DbParameterGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dpg)
}

// Dependencies returns the list of resources [DbParameterGroup] depends_on.
func (dpg *DbParameterGroup) Dependencies() terra.Dependencies {
	return dpg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbParameterGroup].
func (dpg *DbParameterGroup) LifecycleManagement() *terra.Lifecycle {
	return dpg.Lifecycle
}

// Attributes returns the attributes for [DbParameterGroup].
func (dpg *DbParameterGroup) Attributes() dbParameterGroupAttributes {
	return dbParameterGroupAttributes{ref: terra.ReferenceResource(dpg)}
}

// ImportState imports the given attribute values into [DbParameterGroup]'s state.
func (dpg *DbParameterGroup) ImportState(av io.Reader) error {
	dpg.state = &dbParameterGroupState{}
	if err := json.NewDecoder(av).Decode(dpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpg.Type(), dpg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbParameterGroup] has state.
func (dpg *DbParameterGroup) State() (*dbParameterGroupState, bool) {
	return dpg.state, dpg.state != nil
}

// StateMust returns the state for [DbParameterGroup]. Panics if the state is nil.
func (dpg *DbParameterGroup) StateMust() *dbParameterGroupState {
	if dpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpg.Type(), dpg.LocalName()))
	}
	return dpg.state
}

// DbParameterGroupArgs contains the configurations for aws_db_parameter_group.
type DbParameterGroupArgs struct {
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
	Parameter []dbparametergroup.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type dbParameterGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("description"))
}

// Family returns a reference to field family of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("family"))
}

// Id returns a reference to field id of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(dpg.ref.Append("name_prefix"))
}

// Tags returns a reference to field tags of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_parameter_group.
func (dpg dbParameterGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpg.ref.Append("tags_all"))
}

func (dpg dbParameterGroupAttributes) Parameter() terra.SetValue[dbparametergroup.ParameterAttributes] {
	return terra.ReferenceAsSet[dbparametergroup.ParameterAttributes](dpg.ref.Append("parameter"))
}

type dbParameterGroupState struct {
	Arn         string                            `json:"arn"`
	Description string                            `json:"description"`
	Family      string                            `json:"family"`
	Id          string                            `json:"id"`
	Name        string                            `json:"name"`
	NamePrefix  string                            `json:"name_prefix"`
	Tags        map[string]string                 `json:"tags"`
	TagsAll     map[string]string                 `json:"tags_all"`
	Parameter   []dbparametergroup.ParameterState `json:"parameter"`
}
