// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dataexchange_revision

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

// Resource represents the Terraform resource aws_dataexchange_revision.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDataexchangeRevisionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adr *Resource) Type() string {
	return "aws_dataexchange_revision"
}

// LocalName returns the local name for [Resource].
func (adr *Resource) LocalName() string {
	return adr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adr *Resource) Configuration() interface{} {
	return adr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adr *Resource) Dependencies() terra.Dependencies {
	return adr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adr *Resource) LifecycleManagement() *terra.Lifecycle {
	return adr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adr *Resource) Attributes() awsDataexchangeRevisionAttributes {
	return awsDataexchangeRevisionAttributes{ref: terra.ReferenceResource(adr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adr *Resource) ImportState(state io.Reader) error {
	adr.state = &awsDataexchangeRevisionState{}
	if err := json.NewDecoder(state).Decode(adr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adr.Type(), adr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adr *Resource) State() (*awsDataexchangeRevisionState, bool) {
	return adr.state, adr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adr *Resource) StateMust() *awsDataexchangeRevisionState {
	if adr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adr.Type(), adr.LocalName()))
	}
	return adr.state
}

// Args contains the configurations for aws_dataexchange_revision.
type Args struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DataSetId: string, required
	DataSetId terra.StringValue `hcl:"data_set_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsDataexchangeRevisionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("arn"))
}

// Comment returns a reference to field comment of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("comment"))
}

// DataSetId returns a reference to field data_set_id of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) DataSetId() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("data_set_id"))
}

// Id returns a reference to field id of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("id"))
}

// RevisionId returns a reference to field revision_id of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) RevisionId() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("revision_id"))
}

// Tags returns a reference to field tags of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dataexchange_revision.
func (adr awsDataexchangeRevisionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adr.ref.Append("tags_all"))
}

type awsDataexchangeRevisionState struct {
	Arn        string            `json:"arn"`
	Comment    string            `json:"comment"`
	DataSetId  string            `json:"data_set_id"`
	Id         string            `json:"id"`
	RevisionId string            `json:"revision_id"`
	Tags       map[string]string `json:"tags"`
	TagsAll    map[string]string `json:"tags_all"`
}
