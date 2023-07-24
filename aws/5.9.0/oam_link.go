// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	oamlink "github.com/golingon/terraproviders/aws/5.9.0/oamlink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOamLink creates a new instance of [OamLink].
func NewOamLink(name string, args OamLinkArgs) *OamLink {
	return &OamLink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OamLink)(nil)

// OamLink represents the Terraform resource aws_oam_link.
type OamLink struct {
	Name      string
	Args      OamLinkArgs
	state     *oamLinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OamLink].
func (ol *OamLink) Type() string {
	return "aws_oam_link"
}

// LocalName returns the local name for [OamLink].
func (ol *OamLink) LocalName() string {
	return ol.Name
}

// Configuration returns the configuration (args) for [OamLink].
func (ol *OamLink) Configuration() interface{} {
	return ol.Args
}

// DependOn is used for other resources to depend on [OamLink].
func (ol *OamLink) DependOn() terra.Reference {
	return terra.ReferenceResource(ol)
}

// Dependencies returns the list of resources [OamLink] depends_on.
func (ol *OamLink) Dependencies() terra.Dependencies {
	return ol.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OamLink].
func (ol *OamLink) LifecycleManagement() *terra.Lifecycle {
	return ol.Lifecycle
}

// Attributes returns the attributes for [OamLink].
func (ol *OamLink) Attributes() oamLinkAttributes {
	return oamLinkAttributes{ref: terra.ReferenceResource(ol)}
}

// ImportState imports the given attribute values into [OamLink]'s state.
func (ol *OamLink) ImportState(av io.Reader) error {
	ol.state = &oamLinkState{}
	if err := json.NewDecoder(av).Decode(ol.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ol.Type(), ol.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OamLink] has state.
func (ol *OamLink) State() (*oamLinkState, bool) {
	return ol.state, ol.state != nil
}

// StateMust returns the state for [OamLink]. Panics if the state is nil.
func (ol *OamLink) StateMust() *oamLinkState {
	if ol.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ol.Type(), ol.LocalName()))
	}
	return ol.state
}

// OamLinkArgs contains the configurations for aws_oam_link.
type OamLinkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabelTemplate: string, required
	LabelTemplate terra.StringValue `hcl:"label_template,attr" validate:"required"`
	// ResourceTypes: set of string, required
	ResourceTypes terra.SetValue[terra.StringValue] `hcl:"resource_types,attr" validate:"required"`
	// SinkIdentifier: string, required
	SinkIdentifier terra.StringValue `hcl:"sink_identifier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *oamlink.Timeouts `hcl:"timeouts,block"`
}
type oamLinkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_oam_link.
func (ol oamLinkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("arn"))
}

// Id returns a reference to field id of aws_oam_link.
func (ol oamLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("id"))
}

// Label returns a reference to field label of aws_oam_link.
func (ol oamLinkAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("label"))
}

// LabelTemplate returns a reference to field label_template of aws_oam_link.
func (ol oamLinkAttributes) LabelTemplate() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("label_template"))
}

// LinkId returns a reference to field link_id of aws_oam_link.
func (ol oamLinkAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("link_id"))
}

// ResourceTypes returns a reference to field resource_types of aws_oam_link.
func (ol oamLinkAttributes) ResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ol.ref.Append("resource_types"))
}

// SinkArn returns a reference to field sink_arn of aws_oam_link.
func (ol oamLinkAttributes) SinkArn() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("sink_arn"))
}

// SinkIdentifier returns a reference to field sink_identifier of aws_oam_link.
func (ol oamLinkAttributes) SinkIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("sink_identifier"))
}

// Tags returns a reference to field tags of aws_oam_link.
func (ol oamLinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ol.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_oam_link.
func (ol oamLinkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ol.ref.Append("tags_all"))
}

func (ol oamLinkAttributes) Timeouts() oamlink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[oamlink.TimeoutsAttributes](ol.ref.Append("timeouts"))
}

type oamLinkState struct {
	Arn            string                 `json:"arn"`
	Id             string                 `json:"id"`
	Label          string                 `json:"label"`
	LabelTemplate  string                 `json:"label_template"`
	LinkId         string                 `json:"link_id"`
	ResourceTypes  []string               `json:"resource_types"`
	SinkArn        string                 `json:"sink_arn"`
	SinkIdentifier string                 `json:"sink_identifier"`
	Tags           map[string]string      `json:"tags"`
	TagsAll        map[string]string      `json:"tags_all"`
	Timeouts       *oamlink.TimeoutsState `json:"timeouts"`
}
