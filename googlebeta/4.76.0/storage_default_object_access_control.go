// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagedefaultobjectaccesscontrol "github.com/golingon/terraproviders/googlebeta/4.76.0/storagedefaultobjectaccesscontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageDefaultObjectAccessControl creates a new instance of [StorageDefaultObjectAccessControl].
func NewStorageDefaultObjectAccessControl(name string, args StorageDefaultObjectAccessControlArgs) *StorageDefaultObjectAccessControl {
	return &StorageDefaultObjectAccessControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageDefaultObjectAccessControl)(nil)

// StorageDefaultObjectAccessControl represents the Terraform resource google_storage_default_object_access_control.
type StorageDefaultObjectAccessControl struct {
	Name      string
	Args      StorageDefaultObjectAccessControlArgs
	state     *storageDefaultObjectAccessControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) Type() string {
	return "google_storage_default_object_access_control"
}

// LocalName returns the local name for [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) LocalName() string {
	return sdoac.Name
}

// Configuration returns the configuration (args) for [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) Configuration() interface{} {
	return sdoac.Args
}

// DependOn is used for other resources to depend on [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) DependOn() terra.Reference {
	return terra.ReferenceResource(sdoac)
}

// Dependencies returns the list of resources [StorageDefaultObjectAccessControl] depends_on.
func (sdoac *StorageDefaultObjectAccessControl) Dependencies() terra.Dependencies {
	return sdoac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) LifecycleManagement() *terra.Lifecycle {
	return sdoac.Lifecycle
}

// Attributes returns the attributes for [StorageDefaultObjectAccessControl].
func (sdoac *StorageDefaultObjectAccessControl) Attributes() storageDefaultObjectAccessControlAttributes {
	return storageDefaultObjectAccessControlAttributes{ref: terra.ReferenceResource(sdoac)}
}

// ImportState imports the given attribute values into [StorageDefaultObjectAccessControl]'s state.
func (sdoac *StorageDefaultObjectAccessControl) ImportState(av io.Reader) error {
	sdoac.state = &storageDefaultObjectAccessControlState{}
	if err := json.NewDecoder(av).Decode(sdoac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdoac.Type(), sdoac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageDefaultObjectAccessControl] has state.
func (sdoac *StorageDefaultObjectAccessControl) State() (*storageDefaultObjectAccessControlState, bool) {
	return sdoac.state, sdoac.state != nil
}

// StateMust returns the state for [StorageDefaultObjectAccessControl]. Panics if the state is nil.
func (sdoac *StorageDefaultObjectAccessControl) StateMust() *storageDefaultObjectAccessControlState {
	if sdoac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdoac.Type(), sdoac.LocalName()))
	}
	return sdoac.state
}

// StorageDefaultObjectAccessControlArgs contains the configurations for google_storage_default_object_access_control.
type StorageDefaultObjectAccessControlArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Entity: string, required
	Entity terra.StringValue `hcl:"entity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Object: string, optional
	Object terra.StringValue `hcl:"object,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ProjectTeam: min=0
	ProjectTeam []storagedefaultobjectaccesscontrol.ProjectTeam `hcl:"project_team,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storagedefaultobjectaccesscontrol.Timeouts `hcl:"timeouts,block"`
}
type storageDefaultObjectAccessControlAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("bucket"))
}

// Domain returns a reference to field domain of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("domain"))
}

// Email returns a reference to field email of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("email"))
}

// Entity returns a reference to field entity of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Entity() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("entity"))
}

// EntityId returns a reference to field entity_id of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) EntityId() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("entity_id"))
}

// Generation returns a reference to field generation of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Generation() terra.NumberValue {
	return terra.ReferenceAsNumber(sdoac.ref.Append("generation"))
}

// Id returns a reference to field id of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("id"))
}

// Object returns a reference to field object of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Object() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("object"))
}

// Role returns a reference to field role of google_storage_default_object_access_control.
func (sdoac storageDefaultObjectAccessControlAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sdoac.ref.Append("role"))
}

func (sdoac storageDefaultObjectAccessControlAttributes) ProjectTeam() terra.ListValue[storagedefaultobjectaccesscontrol.ProjectTeamAttributes] {
	return terra.ReferenceAsList[storagedefaultobjectaccesscontrol.ProjectTeamAttributes](sdoac.ref.Append("project_team"))
}

func (sdoac storageDefaultObjectAccessControlAttributes) Timeouts() storagedefaultobjectaccesscontrol.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagedefaultobjectaccesscontrol.TimeoutsAttributes](sdoac.ref.Append("timeouts"))
}

type storageDefaultObjectAccessControlState struct {
	Bucket      string                                               `json:"bucket"`
	Domain      string                                               `json:"domain"`
	Email       string                                               `json:"email"`
	Entity      string                                               `json:"entity"`
	EntityId    string                                               `json:"entity_id"`
	Generation  float64                                              `json:"generation"`
	Id          string                                               `json:"id"`
	Object      string                                               `json:"object"`
	Role        string                                               `json:"role"`
	ProjectTeam []storagedefaultobjectaccesscontrol.ProjectTeamState `json:"project_team"`
	Timeouts    *storagedefaultobjectaccesscontrol.TimeoutsState     `json:"timeouts"`
}
