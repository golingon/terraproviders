// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dboptiongroup "github.com/golingon/terraproviders/aws/4.66.1/dboptiongroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbOptionGroup creates a new instance of [DbOptionGroup].
func NewDbOptionGroup(name string, args DbOptionGroupArgs) *DbOptionGroup {
	return &DbOptionGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbOptionGroup)(nil)

// DbOptionGroup represents the Terraform resource aws_db_option_group.
type DbOptionGroup struct {
	Name      string
	Args      DbOptionGroupArgs
	state     *dbOptionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbOptionGroup].
func (dog *DbOptionGroup) Type() string {
	return "aws_db_option_group"
}

// LocalName returns the local name for [DbOptionGroup].
func (dog *DbOptionGroup) LocalName() string {
	return dog.Name
}

// Configuration returns the configuration (args) for [DbOptionGroup].
func (dog *DbOptionGroup) Configuration() interface{} {
	return dog.Args
}

// DependOn is used for other resources to depend on [DbOptionGroup].
func (dog *DbOptionGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dog)
}

// Dependencies returns the list of resources [DbOptionGroup] depends_on.
func (dog *DbOptionGroup) Dependencies() terra.Dependencies {
	return dog.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbOptionGroup].
func (dog *DbOptionGroup) LifecycleManagement() *terra.Lifecycle {
	return dog.Lifecycle
}

// Attributes returns the attributes for [DbOptionGroup].
func (dog *DbOptionGroup) Attributes() dbOptionGroupAttributes {
	return dbOptionGroupAttributes{ref: terra.ReferenceResource(dog)}
}

// ImportState imports the given attribute values into [DbOptionGroup]'s state.
func (dog *DbOptionGroup) ImportState(av io.Reader) error {
	dog.state = &dbOptionGroupState{}
	if err := json.NewDecoder(av).Decode(dog.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dog.Type(), dog.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbOptionGroup] has state.
func (dog *DbOptionGroup) State() (*dbOptionGroupState, bool) {
	return dog.state, dog.state != nil
}

// StateMust returns the state for [DbOptionGroup]. Panics if the state is nil.
func (dog *DbOptionGroup) StateMust() *dbOptionGroupState {
	if dog.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dog.Type(), dog.LocalName()))
	}
	return dog.state
}

// DbOptionGroupArgs contains the configurations for aws_db_option_group.
type DbOptionGroupArgs struct {
	// EngineName: string, required
	EngineName terra.StringValue `hcl:"engine_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MajorEngineVersion: string, required
	MajorEngineVersion terra.StringValue `hcl:"major_engine_version,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// OptionGroupDescription: string, optional
	OptionGroupDescription terra.StringValue `hcl:"option_group_description,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Option: min=0
	Option []dboptiongroup.Option `hcl:"option,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dboptiongroup.Timeouts `hcl:"timeouts,block"`
}
type dbOptionGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_option_group.
func (dog dbOptionGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("arn"))
}

// EngineName returns a reference to field engine_name of aws_db_option_group.
func (dog dbOptionGroupAttributes) EngineName() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("engine_name"))
}

// Id returns a reference to field id of aws_db_option_group.
func (dog dbOptionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("id"))
}

// MajorEngineVersion returns a reference to field major_engine_version of aws_db_option_group.
func (dog dbOptionGroupAttributes) MajorEngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("major_engine_version"))
}

// Name returns a reference to field name of aws_db_option_group.
func (dog dbOptionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_db_option_group.
func (dog dbOptionGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("name_prefix"))
}

// OptionGroupDescription returns a reference to field option_group_description of aws_db_option_group.
func (dog dbOptionGroupAttributes) OptionGroupDescription() terra.StringValue {
	return terra.ReferenceAsString(dog.ref.Append("option_group_description"))
}

// Tags returns a reference to field tags of aws_db_option_group.
func (dog dbOptionGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dog.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_option_group.
func (dog dbOptionGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dog.ref.Append("tags_all"))
}

func (dog dbOptionGroupAttributes) Option() terra.SetValue[dboptiongroup.OptionAttributes] {
	return terra.ReferenceAsSet[dboptiongroup.OptionAttributes](dog.ref.Append("option"))
}

func (dog dbOptionGroupAttributes) Timeouts() dboptiongroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dboptiongroup.TimeoutsAttributes](dog.ref.Append("timeouts"))
}

type dbOptionGroupState struct {
	Arn                    string                       `json:"arn"`
	EngineName             string                       `json:"engine_name"`
	Id                     string                       `json:"id"`
	MajorEngineVersion     string                       `json:"major_engine_version"`
	Name                   string                       `json:"name"`
	NamePrefix             string                       `json:"name_prefix"`
	OptionGroupDescription string                       `json:"option_group_description"`
	Tags                   map[string]string            `json:"tags"`
	TagsAll                map[string]string            `json:"tags_all"`
	Option                 []dboptiongroup.OptionState  `json:"option"`
	Timeouts               *dboptiongroup.TimeoutsState `json:"timeouts"`
}
