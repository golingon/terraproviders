// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewMediaStoreContainer creates a new instance of [MediaStoreContainer].
func NewMediaStoreContainer(name string, args MediaStoreContainerArgs) *MediaStoreContainer {
	return &MediaStoreContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaStoreContainer)(nil)

// MediaStoreContainer represents the Terraform resource aws_media_store_container.
type MediaStoreContainer struct {
	Name      string
	Args      MediaStoreContainerArgs
	state     *mediaStoreContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaStoreContainer].
func (msc *MediaStoreContainer) Type() string {
	return "aws_media_store_container"
}

// LocalName returns the local name for [MediaStoreContainer].
func (msc *MediaStoreContainer) LocalName() string {
	return msc.Name
}

// Configuration returns the configuration (args) for [MediaStoreContainer].
func (msc *MediaStoreContainer) Configuration() interface{} {
	return msc.Args
}

// DependOn is used for other resources to depend on [MediaStoreContainer].
func (msc *MediaStoreContainer) DependOn() terra.Reference {
	return terra.ReferenceResource(msc)
}

// Dependencies returns the list of resources [MediaStoreContainer] depends_on.
func (msc *MediaStoreContainer) Dependencies() terra.Dependencies {
	return msc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaStoreContainer].
func (msc *MediaStoreContainer) LifecycleManagement() *terra.Lifecycle {
	return msc.Lifecycle
}

// Attributes returns the attributes for [MediaStoreContainer].
func (msc *MediaStoreContainer) Attributes() mediaStoreContainerAttributes {
	return mediaStoreContainerAttributes{ref: terra.ReferenceResource(msc)}
}

// ImportState imports the given attribute values into [MediaStoreContainer]'s state.
func (msc *MediaStoreContainer) ImportState(av io.Reader) error {
	msc.state = &mediaStoreContainerState{}
	if err := json.NewDecoder(av).Decode(msc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msc.Type(), msc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaStoreContainer] has state.
func (msc *MediaStoreContainer) State() (*mediaStoreContainerState, bool) {
	return msc.state, msc.state != nil
}

// StateMust returns the state for [MediaStoreContainer]. Panics if the state is nil.
func (msc *MediaStoreContainer) StateMust() *mediaStoreContainerState {
	if msc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msc.Type(), msc.LocalName()))
	}
	return msc.state
}

// MediaStoreContainerArgs contains the configurations for aws_media_store_container.
type MediaStoreContainerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type mediaStoreContainerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_media_store_container.
func (msc mediaStoreContainerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("arn"))
}

// Endpoint returns a reference to field endpoint of aws_media_store_container.
func (msc mediaStoreContainerAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_media_store_container.
func (msc mediaStoreContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("id"))
}

// Name returns a reference to field name of aws_media_store_container.
func (msc mediaStoreContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_media_store_container.
func (msc mediaStoreContainerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_media_store_container.
func (msc mediaStoreContainerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msc.ref.Append("tags_all"))
}

type mediaStoreContainerState struct {
	Arn      string            `json:"arn"`
	Endpoint string            `json:"endpoint"`
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Tags     map[string]string `json:"tags"`
	TagsAll  map[string]string `json:"tags_all"`
}
