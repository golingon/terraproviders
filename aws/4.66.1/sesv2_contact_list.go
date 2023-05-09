// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesv2contactlist "github.com/golingon/terraproviders/aws/4.66.1/sesv2contactlist"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesv2ContactList creates a new instance of [Sesv2ContactList].
func NewSesv2ContactList(name string, args Sesv2ContactListArgs) *Sesv2ContactList {
	return &Sesv2ContactList{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sesv2ContactList)(nil)

// Sesv2ContactList represents the Terraform resource aws_sesv2_contact_list.
type Sesv2ContactList struct {
	Name      string
	Args      Sesv2ContactListArgs
	state     *sesv2ContactListState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Sesv2ContactList].
func (scl *Sesv2ContactList) Type() string {
	return "aws_sesv2_contact_list"
}

// LocalName returns the local name for [Sesv2ContactList].
func (scl *Sesv2ContactList) LocalName() string {
	return scl.Name
}

// Configuration returns the configuration (args) for [Sesv2ContactList].
func (scl *Sesv2ContactList) Configuration() interface{} {
	return scl.Args
}

// DependOn is used for other resources to depend on [Sesv2ContactList].
func (scl *Sesv2ContactList) DependOn() terra.Reference {
	return terra.ReferenceResource(scl)
}

// Dependencies returns the list of resources [Sesv2ContactList] depends_on.
func (scl *Sesv2ContactList) Dependencies() terra.Dependencies {
	return scl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Sesv2ContactList].
func (scl *Sesv2ContactList) LifecycleManagement() *terra.Lifecycle {
	return scl.Lifecycle
}

// Attributes returns the attributes for [Sesv2ContactList].
func (scl *Sesv2ContactList) Attributes() sesv2ContactListAttributes {
	return sesv2ContactListAttributes{ref: terra.ReferenceResource(scl)}
}

// ImportState imports the given attribute values into [Sesv2ContactList]'s state.
func (scl *Sesv2ContactList) ImportState(av io.Reader) error {
	scl.state = &sesv2ContactListState{}
	if err := json.NewDecoder(av).Decode(scl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scl.Type(), scl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Sesv2ContactList] has state.
func (scl *Sesv2ContactList) State() (*sesv2ContactListState, bool) {
	return scl.state, scl.state != nil
}

// StateMust returns the state for [Sesv2ContactList]. Panics if the state is nil.
func (scl *Sesv2ContactList) StateMust() *sesv2ContactListState {
	if scl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scl.Type(), scl.LocalName()))
	}
	return scl.state
}

// Sesv2ContactListArgs contains the configurations for aws_sesv2_contact_list.
type Sesv2ContactListArgs struct {
	// ContactListName: string, required
	ContactListName terra.StringValue `hcl:"contact_list_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Topic: min=0
	Topic []sesv2contactlist.Topic `hcl:"topic,block" validate:"min=0"`
}
type sesv2ContactListAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("arn"))
}

// ContactListName returns a reference to field contact_list_name of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) ContactListName() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("contact_list_name"))
}

// CreatedTimestamp returns a reference to field created_timestamp of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) CreatedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("created_timestamp"))
}

// Description returns a reference to field description of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("description"))
}

// Id returns a reference to field id of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("id"))
}

// LastUpdatedTimestamp returns a reference to field last_updated_timestamp of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) LastUpdatedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(scl.ref.Append("last_updated_timestamp"))
}

// Tags returns a reference to field tags of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sesv2_contact_list.
func (scl sesv2ContactListAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scl.ref.Append("tags_all"))
}

func (scl sesv2ContactListAttributes) Topic() terra.SetValue[sesv2contactlist.TopicAttributes] {
	return terra.ReferenceAsSet[sesv2contactlist.TopicAttributes](scl.ref.Append("topic"))
}

type sesv2ContactListState struct {
	Arn                  string                        `json:"arn"`
	ContactListName      string                        `json:"contact_list_name"`
	CreatedTimestamp     string                        `json:"created_timestamp"`
	Description          string                        `json:"description"`
	Id                   string                        `json:"id"`
	LastUpdatedTimestamp string                        `json:"last_updated_timestamp"`
	Tags                 map[string]string             `json:"tags"`
	TagsAll              map[string]string             `json:"tags_all"`
	Topic                []sesv2contactlist.TopicState `json:"topic"`
}
