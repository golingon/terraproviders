// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelmetadata "github.com/golingon/terraproviders/azurerm/3.66.0/sentinelmetadata"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelMetadata creates a new instance of [SentinelMetadata].
func NewSentinelMetadata(name string, args SentinelMetadataArgs) *SentinelMetadata {
	return &SentinelMetadata{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelMetadata)(nil)

// SentinelMetadata represents the Terraform resource azurerm_sentinel_metadata.
type SentinelMetadata struct {
	Name      string
	Args      SentinelMetadataArgs
	state     *sentinelMetadataState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelMetadata].
func (sm *SentinelMetadata) Type() string {
	return "azurerm_sentinel_metadata"
}

// LocalName returns the local name for [SentinelMetadata].
func (sm *SentinelMetadata) LocalName() string {
	return sm.Name
}

// Configuration returns the configuration (args) for [SentinelMetadata].
func (sm *SentinelMetadata) Configuration() interface{} {
	return sm.Args
}

// DependOn is used for other resources to depend on [SentinelMetadata].
func (sm *SentinelMetadata) DependOn() terra.Reference {
	return terra.ReferenceResource(sm)
}

// Dependencies returns the list of resources [SentinelMetadata] depends_on.
func (sm *SentinelMetadata) Dependencies() terra.Dependencies {
	return sm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelMetadata].
func (sm *SentinelMetadata) LifecycleManagement() *terra.Lifecycle {
	return sm.Lifecycle
}

// Attributes returns the attributes for [SentinelMetadata].
func (sm *SentinelMetadata) Attributes() sentinelMetadataAttributes {
	return sentinelMetadataAttributes{ref: terra.ReferenceResource(sm)}
}

// ImportState imports the given attribute values into [SentinelMetadata]'s state.
func (sm *SentinelMetadata) ImportState(av io.Reader) error {
	sm.state = &sentinelMetadataState{}
	if err := json.NewDecoder(av).Decode(sm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sm.Type(), sm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelMetadata] has state.
func (sm *SentinelMetadata) State() (*sentinelMetadataState, bool) {
	return sm.state, sm.state != nil
}

// StateMust returns the state for [SentinelMetadata]. Panics if the state is nil.
func (sm *SentinelMetadata) StateMust() *sentinelMetadataState {
	if sm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sm.Type(), sm.LocalName()))
	}
	return sm.state
}

// SentinelMetadataArgs contains the configurations for azurerm_sentinel_metadata.
type SentinelMetadataArgs struct {
	// ContentId: string, required
	ContentId terra.StringValue `hcl:"content_id,attr" validate:"required"`
	// ContentSchemaVersion: string, optional
	ContentSchemaVersion terra.StringValue `hcl:"content_schema_version,attr"`
	// CustomVersion: string, optional
	CustomVersion terra.StringValue `hcl:"custom_version,attr"`
	// Dependency: string, optional
	Dependency terra.StringValue `hcl:"dependency,attr"`
	// FirstPublishDate: string, optional
	FirstPublishDate terra.StringValue `hcl:"first_publish_date,attr"`
	// IconId: string, optional
	IconId terra.StringValue `hcl:"icon_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// LastPublishDate: string, optional
	LastPublishDate terra.StringValue `hcl:"last_publish_date,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// PreviewImages: list of string, optional
	PreviewImages terra.ListValue[terra.StringValue] `hcl:"preview_images,attr"`
	// PreviewImagesDark: list of string, optional
	PreviewImagesDark terra.ListValue[terra.StringValue] `hcl:"preview_images_dark,attr"`
	// Providers: list of string, optional
	Providers terra.ListValue[terra.StringValue] `hcl:"providers,attr"`
	// ThreatAnalysisTactics: list of string, optional
	ThreatAnalysisTactics terra.ListValue[terra.StringValue] `hcl:"threat_analysis_tactics,attr"`
	// ThreatAnalysisTechniques: list of string, optional
	ThreatAnalysisTechniques terra.ListValue[terra.StringValue] `hcl:"threat_analysis_techniques,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Author: optional
	Author *sentinelmetadata.Author `hcl:"author,block"`
	// Category: optional
	Category *sentinelmetadata.Category `hcl:"category,block"`
	// Source: optional
	Source *sentinelmetadata.Source `hcl:"source,block"`
	// Support: optional
	Support *sentinelmetadata.Support `hcl:"support,block"`
	// Timeouts: optional
	Timeouts *sentinelmetadata.Timeouts `hcl:"timeouts,block"`
}
type sentinelMetadataAttributes struct {
	ref terra.Reference
}

// ContentId returns a reference to field content_id of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) ContentId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("content_id"))
}

// ContentSchemaVersion returns a reference to field content_schema_version of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) ContentSchemaVersion() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("content_schema_version"))
}

// CustomVersion returns a reference to field custom_version of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) CustomVersion() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("custom_version"))
}

// Dependency returns a reference to field dependency of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Dependency() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("dependency"))
}

// FirstPublishDate returns a reference to field first_publish_date of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) FirstPublishDate() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("first_publish_date"))
}

// IconId returns a reference to field icon_id of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) IconId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("icon_id"))
}

// Id returns a reference to field id of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("kind"))
}

// LastPublishDate returns a reference to field last_publish_date of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) LastPublishDate() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("last_publish_date"))
}

// Name returns a reference to field name of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("name"))
}

// ParentId returns a reference to field parent_id of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("parent_id"))
}

// PreviewImages returns a reference to field preview_images of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) PreviewImages() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sm.ref.Append("preview_images"))
}

// PreviewImagesDark returns a reference to field preview_images_dark of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) PreviewImagesDark() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sm.ref.Append("preview_images_dark"))
}

// Providers returns a reference to field providers of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Providers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sm.ref.Append("providers"))
}

// ThreatAnalysisTactics returns a reference to field threat_analysis_tactics of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) ThreatAnalysisTactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sm.ref.Append("threat_analysis_tactics"))
}

// ThreatAnalysisTechniques returns a reference to field threat_analysis_techniques of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) ThreatAnalysisTechniques() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sm.ref.Append("threat_analysis_techniques"))
}

// Version returns a reference to field version of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("version"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_sentinel_metadata.
func (sm sentinelMetadataAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("workspace_id"))
}

func (sm sentinelMetadataAttributes) Author() terra.ListValue[sentinelmetadata.AuthorAttributes] {
	return terra.ReferenceAsList[sentinelmetadata.AuthorAttributes](sm.ref.Append("author"))
}

func (sm sentinelMetadataAttributes) Category() terra.ListValue[sentinelmetadata.CategoryAttributes] {
	return terra.ReferenceAsList[sentinelmetadata.CategoryAttributes](sm.ref.Append("category"))
}

func (sm sentinelMetadataAttributes) Source() terra.ListValue[sentinelmetadata.SourceAttributes] {
	return terra.ReferenceAsList[sentinelmetadata.SourceAttributes](sm.ref.Append("source"))
}

func (sm sentinelMetadataAttributes) Support() terra.ListValue[sentinelmetadata.SupportAttributes] {
	return terra.ReferenceAsList[sentinelmetadata.SupportAttributes](sm.ref.Append("support"))
}

func (sm sentinelMetadataAttributes) Timeouts() sentinelmetadata.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelmetadata.TimeoutsAttributes](sm.ref.Append("timeouts"))
}

type sentinelMetadataState struct {
	ContentId                string                           `json:"content_id"`
	ContentSchemaVersion     string                           `json:"content_schema_version"`
	CustomVersion            string                           `json:"custom_version"`
	Dependency               string                           `json:"dependency"`
	FirstPublishDate         string                           `json:"first_publish_date"`
	IconId                   string                           `json:"icon_id"`
	Id                       string                           `json:"id"`
	Kind                     string                           `json:"kind"`
	LastPublishDate          string                           `json:"last_publish_date"`
	Name                     string                           `json:"name"`
	ParentId                 string                           `json:"parent_id"`
	PreviewImages            []string                         `json:"preview_images"`
	PreviewImagesDark        []string                         `json:"preview_images_dark"`
	Providers                []string                         `json:"providers"`
	ThreatAnalysisTactics    []string                         `json:"threat_analysis_tactics"`
	ThreatAnalysisTechniques []string                         `json:"threat_analysis_techniques"`
	Version                  string                           `json:"version"`
	WorkspaceId              string                           `json:"workspace_id"`
	Author                   []sentinelmetadata.AuthorState   `json:"author"`
	Category                 []sentinelmetadata.CategoryState `json:"category"`
	Source                   []sentinelmetadata.SourceState   `json:"source"`
	Support                  []sentinelmetadata.SupportState  `json:"support"`
	Timeouts                 *sentinelmetadata.TimeoutsState  `json:"timeouts"`
}
