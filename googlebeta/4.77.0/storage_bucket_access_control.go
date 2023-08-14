// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagebucketaccesscontrol "github.com/golingon/terraproviders/googlebeta/4.77.0/storagebucketaccesscontrol"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBucketAccessControl creates a new instance of [StorageBucketAccessControl].
func NewStorageBucketAccessControl(name string, args StorageBucketAccessControlArgs) *StorageBucketAccessControl {
	return &StorageBucketAccessControl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBucketAccessControl)(nil)

// StorageBucketAccessControl represents the Terraform resource google_storage_bucket_access_control.
type StorageBucketAccessControl struct {
	Name      string
	Args      StorageBucketAccessControlArgs
	state     *storageBucketAccessControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) Type() string {
	return "google_storage_bucket_access_control"
}

// LocalName returns the local name for [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) LocalName() string {
	return sbac.Name
}

// Configuration returns the configuration (args) for [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) Configuration() interface{} {
	return sbac.Args
}

// DependOn is used for other resources to depend on [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) DependOn() terra.Reference {
	return terra.ReferenceResource(sbac)
}

// Dependencies returns the list of resources [StorageBucketAccessControl] depends_on.
func (sbac *StorageBucketAccessControl) Dependencies() terra.Dependencies {
	return sbac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) LifecycleManagement() *terra.Lifecycle {
	return sbac.Lifecycle
}

// Attributes returns the attributes for [StorageBucketAccessControl].
func (sbac *StorageBucketAccessControl) Attributes() storageBucketAccessControlAttributes {
	return storageBucketAccessControlAttributes{ref: terra.ReferenceResource(sbac)}
}

// ImportState imports the given attribute values into [StorageBucketAccessControl]'s state.
func (sbac *StorageBucketAccessControl) ImportState(av io.Reader) error {
	sbac.state = &storageBucketAccessControlState{}
	if err := json.NewDecoder(av).Decode(sbac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbac.Type(), sbac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBucketAccessControl] has state.
func (sbac *StorageBucketAccessControl) State() (*storageBucketAccessControlState, bool) {
	return sbac.state, sbac.state != nil
}

// StateMust returns the state for [StorageBucketAccessControl]. Panics if the state is nil.
func (sbac *StorageBucketAccessControl) StateMust() *storageBucketAccessControlState {
	if sbac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbac.Type(), sbac.LocalName()))
	}
	return sbac.state
}

// StorageBucketAccessControlArgs contains the configurations for google_storage_bucket_access_control.
type StorageBucketAccessControlArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Entity: string, required
	Entity terra.StringValue `hcl:"entity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Role: string, optional
	Role terra.StringValue `hcl:"role,attr"`
	// Timeouts: optional
	Timeouts *storagebucketaccesscontrol.Timeouts `hcl:"timeouts,block"`
}
type storageBucketAccessControlAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("bucket"))
}

// Domain returns a reference to field domain of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("domain"))
}

// Email returns a reference to field email of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("email"))
}

// Entity returns a reference to field entity of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Entity() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("entity"))
}

// Id returns a reference to field id of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("id"))
}

// Role returns a reference to field role of google_storage_bucket_access_control.
func (sbac storageBucketAccessControlAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sbac.ref.Append("role"))
}

func (sbac storageBucketAccessControlAttributes) Timeouts() storagebucketaccesscontrol.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagebucketaccesscontrol.TimeoutsAttributes](sbac.ref.Append("timeouts"))
}

type storageBucketAccessControlState struct {
	Bucket   string                                    `json:"bucket"`
	Domain   string                                    `json:"domain"`
	Email    string                                    `json:"email"`
	Entity   string                                    `json:"entity"`
	Id       string                                    `json:"id"`
	Role     string                                    `json:"role"`
	Timeouts *storagebucketaccesscontrol.TimeoutsState `json:"timeouts"`
}
