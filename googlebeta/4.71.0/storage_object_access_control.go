// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storageobjectaccesscontrol "github.com/golingon/terraproviders/googlebeta/4.71.0/storageobjectaccesscontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageObjectAccessControl creates a new instance of [StorageObjectAccessControl].
func NewStorageObjectAccessControl(name string, args StorageObjectAccessControlArgs) *StorageObjectAccessControl {
	return &StorageObjectAccessControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageObjectAccessControl)(nil)

// StorageObjectAccessControl represents the Terraform resource google_storage_object_access_control.
type StorageObjectAccessControl struct {
	Name      string
	Args      StorageObjectAccessControlArgs
	state     *storageObjectAccessControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) Type() string {
	return "google_storage_object_access_control"
}

// LocalName returns the local name for [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) LocalName() string {
	return soac.Name
}

// Configuration returns the configuration (args) for [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) Configuration() interface{} {
	return soac.Args
}

// DependOn is used for other resources to depend on [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) DependOn() terra.Reference {
	return terra.ReferenceResource(soac)
}

// Dependencies returns the list of resources [StorageObjectAccessControl] depends_on.
func (soac *StorageObjectAccessControl) Dependencies() terra.Dependencies {
	return soac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) LifecycleManagement() *terra.Lifecycle {
	return soac.Lifecycle
}

// Attributes returns the attributes for [StorageObjectAccessControl].
func (soac *StorageObjectAccessControl) Attributes() storageObjectAccessControlAttributes {
	return storageObjectAccessControlAttributes{ref: terra.ReferenceResource(soac)}
}

// ImportState imports the given attribute values into [StorageObjectAccessControl]'s state.
func (soac *StorageObjectAccessControl) ImportState(av io.Reader) error {
	soac.state = &storageObjectAccessControlState{}
	if err := json.NewDecoder(av).Decode(soac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", soac.Type(), soac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageObjectAccessControl] has state.
func (soac *StorageObjectAccessControl) State() (*storageObjectAccessControlState, bool) {
	return soac.state, soac.state != nil
}

// StateMust returns the state for [StorageObjectAccessControl]. Panics if the state is nil.
func (soac *StorageObjectAccessControl) StateMust() *storageObjectAccessControlState {
	if soac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", soac.Type(), soac.LocalName()))
	}
	return soac.state
}

// StorageObjectAccessControlArgs contains the configurations for google_storage_object_access_control.
type StorageObjectAccessControlArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Entity: string, required
	Entity terra.StringValue `hcl:"entity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Object: string, required
	Object terra.StringValue `hcl:"object,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ProjectTeam: min=0
	ProjectTeam []storageobjectaccesscontrol.ProjectTeam `hcl:"project_team,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storageobjectaccesscontrol.Timeouts `hcl:"timeouts,block"`
}
type storageObjectAccessControlAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("bucket"))
}

// Domain returns a reference to field domain of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("domain"))
}

// Email returns a reference to field email of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("email"))
}

// Entity returns a reference to field entity of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Entity() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("entity"))
}

// EntityId returns a reference to field entity_id of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) EntityId() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("entity_id"))
}

// Generation returns a reference to field generation of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Generation() terra.NumberValue {
	return terra.ReferenceAsNumber(soac.ref.Append("generation"))
}

// Id returns a reference to field id of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("id"))
}

// Object returns a reference to field object of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Object() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("object"))
}

// Role returns a reference to field role of google_storage_object_access_control.
func (soac storageObjectAccessControlAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(soac.ref.Append("role"))
}

func (soac storageObjectAccessControlAttributes) ProjectTeam() terra.ListValue[storageobjectaccesscontrol.ProjectTeamAttributes] {
	return terra.ReferenceAsList[storageobjectaccesscontrol.ProjectTeamAttributes](soac.ref.Append("project_team"))
}

func (soac storageObjectAccessControlAttributes) Timeouts() storageobjectaccesscontrol.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageobjectaccesscontrol.TimeoutsAttributes](soac.ref.Append("timeouts"))
}

type storageObjectAccessControlState struct {
	Bucket      string                                        `json:"bucket"`
	Domain      string                                        `json:"domain"`
	Email       string                                        `json:"email"`
	Entity      string                                        `json:"entity"`
	EntityId    string                                        `json:"entity_id"`
	Generation  float64                                       `json:"generation"`
	Id          string                                        `json:"id"`
	Object      string                                        `json:"object"`
	Role        string                                        `json:"role"`
	ProjectTeam []storageobjectaccesscontrol.ProjectTeamState `json:"project_team"`
	Timeouts    *storageobjectaccesscontrol.TimeoutsState     `json:"timeouts"`
}
