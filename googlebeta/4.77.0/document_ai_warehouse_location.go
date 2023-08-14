// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	documentaiwarehouselocation "github.com/golingon/terraproviders/googlebeta/4.77.0/documentaiwarehouselocation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocumentAiWarehouseLocation creates a new instance of [DocumentAiWarehouseLocation].
func NewDocumentAiWarehouseLocation(name string, args DocumentAiWarehouseLocationArgs) *DocumentAiWarehouseLocation {
	return &DocumentAiWarehouseLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocumentAiWarehouseLocation)(nil)

// DocumentAiWarehouseLocation represents the Terraform resource google_document_ai_warehouse_location.
type DocumentAiWarehouseLocation struct {
	Name      string
	Args      DocumentAiWarehouseLocationArgs
	state     *documentAiWarehouseLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) Type() string {
	return "google_document_ai_warehouse_location"
}

// LocalName returns the local name for [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) LocalName() string {
	return dawl.Name
}

// Configuration returns the configuration (args) for [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) Configuration() interface{} {
	return dawl.Args
}

// DependOn is used for other resources to depend on [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) DependOn() terra.Reference {
	return terra.ReferenceResource(dawl)
}

// Dependencies returns the list of resources [DocumentAiWarehouseLocation] depends_on.
func (dawl *DocumentAiWarehouseLocation) Dependencies() terra.Dependencies {
	return dawl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) LifecycleManagement() *terra.Lifecycle {
	return dawl.Lifecycle
}

// Attributes returns the attributes for [DocumentAiWarehouseLocation].
func (dawl *DocumentAiWarehouseLocation) Attributes() documentAiWarehouseLocationAttributes {
	return documentAiWarehouseLocationAttributes{ref: terra.ReferenceResource(dawl)}
}

// ImportState imports the given attribute values into [DocumentAiWarehouseLocation]'s state.
func (dawl *DocumentAiWarehouseLocation) ImportState(av io.Reader) error {
	dawl.state = &documentAiWarehouseLocationState{}
	if err := json.NewDecoder(av).Decode(dawl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dawl.Type(), dawl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocumentAiWarehouseLocation] has state.
func (dawl *DocumentAiWarehouseLocation) State() (*documentAiWarehouseLocationState, bool) {
	return dawl.state, dawl.state != nil
}

// StateMust returns the state for [DocumentAiWarehouseLocation]. Panics if the state is nil.
func (dawl *DocumentAiWarehouseLocation) StateMust() *documentAiWarehouseLocationState {
	if dawl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dawl.Type(), dawl.LocalName()))
	}
	return dawl.state
}

// DocumentAiWarehouseLocationArgs contains the configurations for google_document_ai_warehouse_location.
type DocumentAiWarehouseLocationArgs struct {
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
	Timeouts *documentaiwarehouselocation.Timeouts `hcl:"timeouts,block"`
}
type documentAiWarehouseLocationAttributes struct {
	ref terra.Reference
}

// AccessControlMode returns a reference to field access_control_mode of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) AccessControlMode() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("access_control_mode"))
}

// DatabaseType returns a reference to field database_type of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("database_type"))
}

// DocumentCreatorDefaultRole returns a reference to field document_creator_default_role of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) DocumentCreatorDefaultRole() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("document_creator_default_role"))
}

// Id returns a reference to field id of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("id"))
}

// KmsKey returns a reference to field kms_key of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("kms_key"))
}

// Location returns a reference to field location of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("location"))
}

// ProjectNumber returns a reference to field project_number of google_document_ai_warehouse_location.
func (dawl documentAiWarehouseLocationAttributes) ProjectNumber() terra.StringValue {
	return terra.ReferenceAsString(dawl.ref.Append("project_number"))
}

func (dawl documentAiWarehouseLocationAttributes) Timeouts() documentaiwarehouselocation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[documentaiwarehouselocation.TimeoutsAttributes](dawl.ref.Append("timeouts"))
}

type documentAiWarehouseLocationState struct {
	AccessControlMode          string                                     `json:"access_control_mode"`
	DatabaseType               string                                     `json:"database_type"`
	DocumentCreatorDefaultRole string                                     `json:"document_creator_default_role"`
	Id                         string                                     `json:"id"`
	KmsKey                     string                                     `json:"kms_key"`
	Location                   string                                     `json:"location"`
	ProjectNumber              string                                     `json:"project_number"`
	Timeouts                   *documentaiwarehouselocation.TimeoutsState `json:"timeouts"`
}
