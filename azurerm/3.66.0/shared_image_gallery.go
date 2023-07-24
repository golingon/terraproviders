// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sharedimagegallery "github.com/golingon/terraproviders/azurerm/3.66.0/sharedimagegallery"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSharedImageGallery creates a new instance of [SharedImageGallery].
func NewSharedImageGallery(name string, args SharedImageGalleryArgs) *SharedImageGallery {
	return &SharedImageGallery{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SharedImageGallery)(nil)

// SharedImageGallery represents the Terraform resource azurerm_shared_image_gallery.
type SharedImageGallery struct {
	Name      string
	Args      SharedImageGalleryArgs
	state     *sharedImageGalleryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SharedImageGallery].
func (sig *SharedImageGallery) Type() string {
	return "azurerm_shared_image_gallery"
}

// LocalName returns the local name for [SharedImageGallery].
func (sig *SharedImageGallery) LocalName() string {
	return sig.Name
}

// Configuration returns the configuration (args) for [SharedImageGallery].
func (sig *SharedImageGallery) Configuration() interface{} {
	return sig.Args
}

// DependOn is used for other resources to depend on [SharedImageGallery].
func (sig *SharedImageGallery) DependOn() terra.Reference {
	return terra.ReferenceResource(sig)
}

// Dependencies returns the list of resources [SharedImageGallery] depends_on.
func (sig *SharedImageGallery) Dependencies() terra.Dependencies {
	return sig.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SharedImageGallery].
func (sig *SharedImageGallery) LifecycleManagement() *terra.Lifecycle {
	return sig.Lifecycle
}

// Attributes returns the attributes for [SharedImageGallery].
func (sig *SharedImageGallery) Attributes() sharedImageGalleryAttributes {
	return sharedImageGalleryAttributes{ref: terra.ReferenceResource(sig)}
}

// ImportState imports the given attribute values into [SharedImageGallery]'s state.
func (sig *SharedImageGallery) ImportState(av io.Reader) error {
	sig.state = &sharedImageGalleryState{}
	if err := json.NewDecoder(av).Decode(sig.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sig.Type(), sig.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SharedImageGallery] has state.
func (sig *SharedImageGallery) State() (*sharedImageGalleryState, bool) {
	return sig.state, sig.state != nil
}

// StateMust returns the state for [SharedImageGallery]. Panics if the state is nil.
func (sig *SharedImageGallery) StateMust() *sharedImageGalleryState {
	if sig.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sig.Type(), sig.LocalName()))
	}
	return sig.state
}

// SharedImageGalleryArgs contains the configurations for azurerm_shared_image_gallery.
type SharedImageGalleryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *sharedimagegallery.Timeouts `hcl:"timeouts,block"`
}
type sharedImageGalleryAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sig.ref.Append("tags"))
}

// UniqueName returns a reference to field unique_name of azurerm_shared_image_gallery.
func (sig sharedImageGalleryAttributes) UniqueName() terra.StringValue {
	return terra.ReferenceAsString(sig.ref.Append("unique_name"))
}

func (sig sharedImageGalleryAttributes) Timeouts() sharedimagegallery.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sharedimagegallery.TimeoutsAttributes](sig.ref.Append("timeouts"))
}

type sharedImageGalleryState struct {
	Description       string                            `json:"description"`
	Id                string                            `json:"id"`
	Location          string                            `json:"location"`
	Name              string                            `json:"name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	Tags              map[string]string                 `json:"tags"`
	UniqueName        string                            `json:"unique_name"`
	Timeouts          *sharedimagegallery.TimeoutsState `json:"timeouts"`
}
