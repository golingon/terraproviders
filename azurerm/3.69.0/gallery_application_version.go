// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	galleryapplicationversion "github.com/golingon/terraproviders/azurerm/3.69.0/galleryapplicationversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGalleryApplicationVersion creates a new instance of [GalleryApplicationVersion].
func NewGalleryApplicationVersion(name string, args GalleryApplicationVersionArgs) *GalleryApplicationVersion {
	return &GalleryApplicationVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GalleryApplicationVersion)(nil)

// GalleryApplicationVersion represents the Terraform resource azurerm_gallery_application_version.
type GalleryApplicationVersion struct {
	Name      string
	Args      GalleryApplicationVersionArgs
	state     *galleryApplicationVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) Type() string {
	return "azurerm_gallery_application_version"
}

// LocalName returns the local name for [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) LocalName() string {
	return gav.Name
}

// Configuration returns the configuration (args) for [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) Configuration() interface{} {
	return gav.Args
}

// DependOn is used for other resources to depend on [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(gav)
}

// Dependencies returns the list of resources [GalleryApplicationVersion] depends_on.
func (gav *GalleryApplicationVersion) Dependencies() terra.Dependencies {
	return gav.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) LifecycleManagement() *terra.Lifecycle {
	return gav.Lifecycle
}

// Attributes returns the attributes for [GalleryApplicationVersion].
func (gav *GalleryApplicationVersion) Attributes() galleryApplicationVersionAttributes {
	return galleryApplicationVersionAttributes{ref: terra.ReferenceResource(gav)}
}

// ImportState imports the given attribute values into [GalleryApplicationVersion]'s state.
func (gav *GalleryApplicationVersion) ImportState(av io.Reader) error {
	gav.state = &galleryApplicationVersionState{}
	if err := json.NewDecoder(av).Decode(gav.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gav.Type(), gav.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GalleryApplicationVersion] has state.
func (gav *GalleryApplicationVersion) State() (*galleryApplicationVersionState, bool) {
	return gav.state, gav.state != nil
}

// StateMust returns the state for [GalleryApplicationVersion]. Panics if the state is nil.
func (gav *GalleryApplicationVersion) StateMust() *galleryApplicationVersionState {
	if gav.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gav.Type(), gav.LocalName()))
	}
	return gav.state
}

// GalleryApplicationVersionArgs contains the configurations for azurerm_gallery_application_version.
type GalleryApplicationVersionArgs struct {
	// EnableHealthCheck: bool, optional
	EnableHealthCheck terra.BoolValue `hcl:"enable_health_check,attr"`
	// EndOfLifeDate: string, optional
	EndOfLifeDate terra.StringValue `hcl:"end_of_life_date,attr"`
	// ExcludeFromLatest: bool, optional
	ExcludeFromLatest terra.BoolValue `hcl:"exclude_from_latest,attr"`
	// GalleryApplicationId: string, required
	GalleryApplicationId terra.StringValue `hcl:"gallery_application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ManageAction: required
	ManageAction *galleryapplicationversion.ManageAction `hcl:"manage_action,block" validate:"required"`
	// Source: required
	Source *galleryapplicationversion.Source `hcl:"source,block" validate:"required"`
	// TargetRegion: min=1
	TargetRegion []galleryapplicationversion.TargetRegion `hcl:"target_region,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *galleryapplicationversion.Timeouts `hcl:"timeouts,block"`
}
type galleryApplicationVersionAttributes struct {
	ref terra.Reference
}

// EnableHealthCheck returns a reference to field enable_health_check of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) EnableHealthCheck() terra.BoolValue {
	return terra.ReferenceAsBool(gav.ref.Append("enable_health_check"))
}

// EndOfLifeDate returns a reference to field end_of_life_date of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) EndOfLifeDate() terra.StringValue {
	return terra.ReferenceAsString(gav.ref.Append("end_of_life_date"))
}

// ExcludeFromLatest returns a reference to field exclude_from_latest of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) ExcludeFromLatest() terra.BoolValue {
	return terra.ReferenceAsBool(gav.ref.Append("exclude_from_latest"))
}

// GalleryApplicationId returns a reference to field gallery_application_id of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) GalleryApplicationId() terra.StringValue {
	return terra.ReferenceAsString(gav.ref.Append("gallery_application_id"))
}

// Id returns a reference to field id of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gav.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gav.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gav.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_gallery_application_version.
func (gav galleryApplicationVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gav.ref.Append("tags"))
}

func (gav galleryApplicationVersionAttributes) ManageAction() terra.ListValue[galleryapplicationversion.ManageActionAttributes] {
	return terra.ReferenceAsList[galleryapplicationversion.ManageActionAttributes](gav.ref.Append("manage_action"))
}

func (gav galleryApplicationVersionAttributes) Source() terra.ListValue[galleryapplicationversion.SourceAttributes] {
	return terra.ReferenceAsList[galleryapplicationversion.SourceAttributes](gav.ref.Append("source"))
}

func (gav galleryApplicationVersionAttributes) TargetRegion() terra.ListValue[galleryapplicationversion.TargetRegionAttributes] {
	return terra.ReferenceAsList[galleryapplicationversion.TargetRegionAttributes](gav.ref.Append("target_region"))
}

func (gav galleryApplicationVersionAttributes) Timeouts() galleryapplicationversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[galleryapplicationversion.TimeoutsAttributes](gav.ref.Append("timeouts"))
}

type galleryApplicationVersionState struct {
	EnableHealthCheck    bool                                          `json:"enable_health_check"`
	EndOfLifeDate        string                                        `json:"end_of_life_date"`
	ExcludeFromLatest    bool                                          `json:"exclude_from_latest"`
	GalleryApplicationId string                                        `json:"gallery_application_id"`
	Id                   string                                        `json:"id"`
	Location             string                                        `json:"location"`
	Name                 string                                        `json:"name"`
	Tags                 map[string]string                             `json:"tags"`
	ManageAction         []galleryapplicationversion.ManageActionState `json:"manage_action"`
	Source               []galleryapplicationversion.SourceState       `json:"source"`
	TargetRegion         []galleryapplicationversion.TargetRegionState `json:"target_region"`
	Timeouts             *galleryapplicationversion.TimeoutsState      `json:"timeouts"`
}