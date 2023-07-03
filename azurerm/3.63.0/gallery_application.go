// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	galleryapplication "github.com/golingon/terraproviders/azurerm/3.63.0/galleryapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGalleryApplication creates a new instance of [GalleryApplication].
func NewGalleryApplication(name string, args GalleryApplicationArgs) *GalleryApplication {
	return &GalleryApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GalleryApplication)(nil)

// GalleryApplication represents the Terraform resource azurerm_gallery_application.
type GalleryApplication struct {
	Name      string
	Args      GalleryApplicationArgs
	state     *galleryApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GalleryApplication].
func (ga *GalleryApplication) Type() string {
	return "azurerm_gallery_application"
}

// LocalName returns the local name for [GalleryApplication].
func (ga *GalleryApplication) LocalName() string {
	return ga.Name
}

// Configuration returns the configuration (args) for [GalleryApplication].
func (ga *GalleryApplication) Configuration() interface{} {
	return ga.Args
}

// DependOn is used for other resources to depend on [GalleryApplication].
func (ga *GalleryApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(ga)
}

// Dependencies returns the list of resources [GalleryApplication] depends_on.
func (ga *GalleryApplication) Dependencies() terra.Dependencies {
	return ga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GalleryApplication].
func (ga *GalleryApplication) LifecycleManagement() *terra.Lifecycle {
	return ga.Lifecycle
}

// Attributes returns the attributes for [GalleryApplication].
func (ga *GalleryApplication) Attributes() galleryApplicationAttributes {
	return galleryApplicationAttributes{ref: terra.ReferenceResource(ga)}
}

// ImportState imports the given attribute values into [GalleryApplication]'s state.
func (ga *GalleryApplication) ImportState(av io.Reader) error {
	ga.state = &galleryApplicationState{}
	if err := json.NewDecoder(av).Decode(ga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ga.Type(), ga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GalleryApplication] has state.
func (ga *GalleryApplication) State() (*galleryApplicationState, bool) {
	return ga.state, ga.state != nil
}

// StateMust returns the state for [GalleryApplication]. Panics if the state is nil.
func (ga *GalleryApplication) StateMust() *galleryApplicationState {
	if ga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ga.Type(), ga.LocalName()))
	}
	return ga.state
}

// GalleryApplicationArgs contains the configurations for azurerm_gallery_application.
type GalleryApplicationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EndOfLifeDate: string, optional
	EndOfLifeDate terra.StringValue `hcl:"end_of_life_date,attr"`
	// Eula: string, optional
	Eula terra.StringValue `hcl:"eula,attr"`
	// GalleryId: string, required
	GalleryId terra.StringValue `hcl:"gallery_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivacyStatementUri: string, optional
	PrivacyStatementUri terra.StringValue `hcl:"privacy_statement_uri,attr"`
	// ReleaseNoteUri: string, optional
	ReleaseNoteUri terra.StringValue `hcl:"release_note_uri,attr"`
	// SupportedOsType: string, required
	SupportedOsType terra.StringValue `hcl:"supported_os_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *galleryapplication.Timeouts `hcl:"timeouts,block"`
}
type galleryApplicationAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("description"))
}

// EndOfLifeDate returns a reference to field end_of_life_date of azurerm_gallery_application.
func (ga galleryApplicationAttributes) EndOfLifeDate() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("end_of_life_date"))
}

// Eula returns a reference to field eula of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Eula() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("eula"))
}

// GalleryId returns a reference to field gallery_id of azurerm_gallery_application.
func (ga galleryApplicationAttributes) GalleryId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("gallery_id"))
}

// Id returns a reference to field id of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("name"))
}

// PrivacyStatementUri returns a reference to field privacy_statement_uri of azurerm_gallery_application.
func (ga galleryApplicationAttributes) PrivacyStatementUri() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("privacy_statement_uri"))
}

// ReleaseNoteUri returns a reference to field release_note_uri of azurerm_gallery_application.
func (ga galleryApplicationAttributes) ReleaseNoteUri() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("release_note_uri"))
}

// SupportedOsType returns a reference to field supported_os_type of azurerm_gallery_application.
func (ga galleryApplicationAttributes) SupportedOsType() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("supported_os_type"))
}

// Tags returns a reference to field tags of azurerm_gallery_application.
func (ga galleryApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags"))
}

func (ga galleryApplicationAttributes) Timeouts() galleryapplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[galleryapplication.TimeoutsAttributes](ga.ref.Append("timeouts"))
}

type galleryApplicationState struct {
	Description         string                            `json:"description"`
	EndOfLifeDate       string                            `json:"end_of_life_date"`
	Eula                string                            `json:"eula"`
	GalleryId           string                            `json:"gallery_id"`
	Id                  string                            `json:"id"`
	Location            string                            `json:"location"`
	Name                string                            `json:"name"`
	PrivacyStatementUri string                            `json:"privacy_statement_uri"`
	ReleaseNoteUri      string                            `json:"release_note_uri"`
	SupportedOsType     string                            `json:"supported_os_type"`
	Tags                map[string]string                 `json:"tags"`
	Timeouts            *galleryapplication.TimeoutsState `json:"timeouts"`
}
