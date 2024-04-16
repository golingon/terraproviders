// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_document_ai_warehouse_location

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

// Resource represents the Terraform resource google_document_ai_warehouse_location.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDocumentAiWarehouseLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdawl *Resource) Type() string {
	return "google_document_ai_warehouse_location"
}

// LocalName returns the local name for [Resource].
func (gdawl *Resource) LocalName() string {
	return gdawl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdawl *Resource) Configuration() interface{} {
	return gdawl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdawl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdawl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdawl *Resource) Dependencies() terra.Dependencies {
	return gdawl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdawl *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdawl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdawl *Resource) Attributes() googleDocumentAiWarehouseLocationAttributes {
	return googleDocumentAiWarehouseLocationAttributes{ref: terra.ReferenceResource(gdawl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdawl *Resource) ImportState(state io.Reader) error {
	gdawl.state = &googleDocumentAiWarehouseLocationState{}
	if err := json.NewDecoder(state).Decode(gdawl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdawl.Type(), gdawl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdawl *Resource) State() (*googleDocumentAiWarehouseLocationState, bool) {
	return gdawl.state, gdawl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdawl *Resource) StateMust() *googleDocumentAiWarehouseLocationState {
	if gdawl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdawl.Type(), gdawl.LocalName()))
	}
	return gdawl.state
}

// Args contains the configurations for google_document_ai_warehouse_location.
type Args struct {
	// AccessControlMode: string, required
	AccessControlMode terra.StringValue `hcl:"access_control_mode,attr" validate:"required"`
	// DatabaseType: string, required
	DatabaseType terra.StringValue `hcl:"database_type,attr" validate:"required"`
	// DocumentCreatorDefaultRole: string, optional
	DocumentCreatorDefaultRole terra.StringValue `hcl:"document_creator_default_role,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ProjectNumber: string, required
	ProjectNumber terra.StringValue `hcl:"project_number,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDocumentAiWarehouseLocationAttributes struct {
	ref terra.Reference
}

// AccessControlMode returns a reference to field access_control_mode of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) AccessControlMode() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("access_control_mode"))
}

// DatabaseType returns a reference to field database_type of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("database_type"))
}

// DocumentCreatorDefaultRole returns a reference to field document_creator_default_role of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) DocumentCreatorDefaultRole() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("document_creator_default_role"))
}

// Id returns a reference to field id of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("id"))
}

// KmsKey returns a reference to field kms_key of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("kms_key"))
}

// Location returns a reference to field location of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("location"))
}

// ProjectNumber returns a reference to field project_number of google_document_ai_warehouse_location.
func (gdawl googleDocumentAiWarehouseLocationAttributes) ProjectNumber() terra.StringValue {
	return terra.ReferenceAsString(gdawl.ref.Append("project_number"))
}

func (gdawl googleDocumentAiWarehouseLocationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdawl.ref.Append("timeouts"))
}

type googleDocumentAiWarehouseLocationState struct {
	AccessControlMode          string         `json:"access_control_mode"`
	DatabaseType               string         `json:"database_type"`
	DocumentCreatorDefaultRole string         `json:"document_creator_default_role"`
	Id                         string         `json:"id"`
	KmsKey                     string         `json:"kms_key"`
	Location                   string         `json:"location"`
	ProjectNumber              string         `json:"project_number"`
	Timeouts                   *TimeoutsState `json:"timeouts"`
}
