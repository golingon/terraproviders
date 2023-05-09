// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	tagslocationtagbinding "github.com/golingon/terraproviders/googlebeta/4.64.0/tagslocationtagbinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTagsLocationTagBinding creates a new instance of [TagsLocationTagBinding].
func NewTagsLocationTagBinding(name string, args TagsLocationTagBindingArgs) *TagsLocationTagBinding {
	return &TagsLocationTagBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TagsLocationTagBinding)(nil)

// TagsLocationTagBinding represents the Terraform resource google_tags_location_tag_binding.
type TagsLocationTagBinding struct {
	Name      string
	Args      TagsLocationTagBindingArgs
	state     *tagsLocationTagBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) Type() string {
	return "google_tags_location_tag_binding"
}

// LocalName returns the local name for [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) LocalName() string {
	return tltb.Name
}

// Configuration returns the configuration (args) for [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) Configuration() interface{} {
	return tltb.Args
}

// DependOn is used for other resources to depend on [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(tltb)
}

// Dependencies returns the list of resources [TagsLocationTagBinding] depends_on.
func (tltb *TagsLocationTagBinding) Dependencies() terra.Dependencies {
	return tltb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) LifecycleManagement() *terra.Lifecycle {
	return tltb.Lifecycle
}

// Attributes returns the attributes for [TagsLocationTagBinding].
func (tltb *TagsLocationTagBinding) Attributes() tagsLocationTagBindingAttributes {
	return tagsLocationTagBindingAttributes{ref: terra.ReferenceResource(tltb)}
}

// ImportState imports the given attribute values into [TagsLocationTagBinding]'s state.
func (tltb *TagsLocationTagBinding) ImportState(av io.Reader) error {
	tltb.state = &tagsLocationTagBindingState{}
	if err := json.NewDecoder(av).Decode(tltb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tltb.Type(), tltb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TagsLocationTagBinding] has state.
func (tltb *TagsLocationTagBinding) State() (*tagsLocationTagBindingState, bool) {
	return tltb.state, tltb.state != nil
}

// StateMust returns the state for [TagsLocationTagBinding]. Panics if the state is nil.
func (tltb *TagsLocationTagBinding) StateMust() *tagsLocationTagBindingState {
	if tltb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tltb.Type(), tltb.LocalName()))
	}
	return tltb.state
}

// TagsLocationTagBindingArgs contains the configurations for google_tags_location_tag_binding.
type TagsLocationTagBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *tagslocationtagbinding.Timeouts `hcl:"timeouts,block"`
}
type tagsLocationTagBindingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_tags_location_tag_binding.
func (tltb tagsLocationTagBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tltb.ref.Append("id"))
}

// Location returns a reference to field location of google_tags_location_tag_binding.
func (tltb tagsLocationTagBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(tltb.ref.Append("location"))
}

// Name returns a reference to field name of google_tags_location_tag_binding.
func (tltb tagsLocationTagBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tltb.ref.Append("name"))
}

// Parent returns a reference to field parent of google_tags_location_tag_binding.
func (tltb tagsLocationTagBindingAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(tltb.ref.Append("parent"))
}

// TagValue returns a reference to field tag_value of google_tags_location_tag_binding.
func (tltb tagsLocationTagBindingAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(tltb.ref.Append("tag_value"))
}

func (tltb tagsLocationTagBindingAttributes) Timeouts() tagslocationtagbinding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[tagslocationtagbinding.TimeoutsAttributes](tltb.ref.Append("timeouts"))
}

type tagsLocationTagBindingState struct {
	Id       string                                `json:"id"`
	Location string                                `json:"location"`
	Name     string                                `json:"name"`
	Parent   string                                `json:"parent"`
	TagValue string                                `json:"tag_value"`
	Timeouts *tagslocationtagbinding.TimeoutsState `json:"timeouts"`
}
