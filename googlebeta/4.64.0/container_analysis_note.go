// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	containeranalysisnote "github.com/golingon/terraproviders/googlebeta/4.64.0/containeranalysisnote"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAnalysisNote creates a new instance of [ContainerAnalysisNote].
func NewContainerAnalysisNote(name string, args ContainerAnalysisNoteArgs) *ContainerAnalysisNote {
	return &ContainerAnalysisNote{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAnalysisNote)(nil)

// ContainerAnalysisNote represents the Terraform resource google_container_analysis_note.
type ContainerAnalysisNote struct {
	Name      string
	Args      ContainerAnalysisNoteArgs
	state     *containerAnalysisNoteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) Type() string {
	return "google_container_analysis_note"
}

// LocalName returns the local name for [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) LocalName() string {
	return can.Name
}

// Configuration returns the configuration (args) for [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) Configuration() interface{} {
	return can.Args
}

// DependOn is used for other resources to depend on [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) DependOn() terra.Reference {
	return terra.ReferenceResource(can)
}

// Dependencies returns the list of resources [ContainerAnalysisNote] depends_on.
func (can *ContainerAnalysisNote) Dependencies() terra.Dependencies {
	return can.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) LifecycleManagement() *terra.Lifecycle {
	return can.Lifecycle
}

// Attributes returns the attributes for [ContainerAnalysisNote].
func (can *ContainerAnalysisNote) Attributes() containerAnalysisNoteAttributes {
	return containerAnalysisNoteAttributes{ref: terra.ReferenceResource(can)}
}

// ImportState imports the given attribute values into [ContainerAnalysisNote]'s state.
func (can *ContainerAnalysisNote) ImportState(av io.Reader) error {
	can.state = &containerAnalysisNoteState{}
	if err := json.NewDecoder(av).Decode(can.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", can.Type(), can.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAnalysisNote] has state.
func (can *ContainerAnalysisNote) State() (*containerAnalysisNoteState, bool) {
	return can.state, can.state != nil
}

// StateMust returns the state for [ContainerAnalysisNote]. Panics if the state is nil.
func (can *ContainerAnalysisNote) StateMust() *containerAnalysisNoteState {
	if can.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", can.Type(), can.LocalName()))
	}
	return can.state
}

// ContainerAnalysisNoteArgs contains the configurations for google_container_analysis_note.
type ContainerAnalysisNoteArgs struct {
	// ExpirationTime: string, optional
	ExpirationTime terra.StringValue `hcl:"expiration_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LongDescription: string, optional
	LongDescription terra.StringValue `hcl:"long_description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RelatedNoteNames: set of string, optional
	RelatedNoteNames terra.SetValue[terra.StringValue] `hcl:"related_note_names,attr"`
	// ShortDescription: string, optional
	ShortDescription terra.StringValue `hcl:"short_description,attr"`
	// AttestationAuthority: required
	AttestationAuthority *containeranalysisnote.AttestationAuthority `hcl:"attestation_authority,block" validate:"required"`
	// RelatedUrl: min=0
	RelatedUrl []containeranalysisnote.RelatedUrl `hcl:"related_url,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *containeranalysisnote.Timeouts `hcl:"timeouts,block"`
}
type containerAnalysisNoteAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("create_time"))
}

// ExpirationTime returns a reference to field expiration_time of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("expiration_time"))
}

// Id returns a reference to field id of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("id"))
}

// Kind returns a reference to field kind of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("kind"))
}

// LongDescription returns a reference to field long_description of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) LongDescription() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("long_description"))
}

// Name returns a reference to field name of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("name"))
}

// Project returns a reference to field project of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("project"))
}

// RelatedNoteNames returns a reference to field related_note_names of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) RelatedNoteNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](can.ref.Append("related_note_names"))
}

// ShortDescription returns a reference to field short_description of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) ShortDescription() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("short_description"))
}

// UpdateTime returns a reference to field update_time of google_container_analysis_note.
func (can containerAnalysisNoteAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(can.ref.Append("update_time"))
}

func (can containerAnalysisNoteAttributes) AttestationAuthority() terra.ListValue[containeranalysisnote.AttestationAuthorityAttributes] {
	return terra.ReferenceAsList[containeranalysisnote.AttestationAuthorityAttributes](can.ref.Append("attestation_authority"))
}

func (can containerAnalysisNoteAttributes) RelatedUrl() terra.SetValue[containeranalysisnote.RelatedUrlAttributes] {
	return terra.ReferenceAsSet[containeranalysisnote.RelatedUrlAttributes](can.ref.Append("related_url"))
}

func (can containerAnalysisNoteAttributes) Timeouts() containeranalysisnote.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containeranalysisnote.TimeoutsAttributes](can.ref.Append("timeouts"))
}

type containerAnalysisNoteState struct {
	CreateTime           string                                            `json:"create_time"`
	ExpirationTime       string                                            `json:"expiration_time"`
	Id                   string                                            `json:"id"`
	Kind                 string                                            `json:"kind"`
	LongDescription      string                                            `json:"long_description"`
	Name                 string                                            `json:"name"`
	Project              string                                            `json:"project"`
	RelatedNoteNames     []string                                          `json:"related_note_names"`
	ShortDescription     string                                            `json:"short_description"`
	UpdateTime           string                                            `json:"update_time"`
	AttestationAuthority []containeranalysisnote.AttestationAuthorityState `json:"attestation_authority"`
	RelatedUrl           []containeranalysisnote.RelatedUrlState           `json:"related_url"`
	Timeouts             *containeranalysisnote.TimeoutsState              `json:"timeouts"`
}
