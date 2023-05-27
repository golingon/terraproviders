// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	storagehmackey "github.com/golingon/terraproviders/google/4.66.0/storagehmackey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageHmacKey creates a new instance of [StorageHmacKey].
func NewStorageHmacKey(name string, args StorageHmacKeyArgs) *StorageHmacKey {
	return &StorageHmacKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageHmacKey)(nil)

// StorageHmacKey represents the Terraform resource google_storage_hmac_key.
type StorageHmacKey struct {
	Name      string
	Args      StorageHmacKeyArgs
	state     *storageHmacKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageHmacKey].
func (shk *StorageHmacKey) Type() string {
	return "google_storage_hmac_key"
}

// LocalName returns the local name for [StorageHmacKey].
func (shk *StorageHmacKey) LocalName() string {
	return shk.Name
}

// Configuration returns the configuration (args) for [StorageHmacKey].
func (shk *StorageHmacKey) Configuration() interface{} {
	return shk.Args
}

// DependOn is used for other resources to depend on [StorageHmacKey].
func (shk *StorageHmacKey) DependOn() terra.Reference {
	return terra.ReferenceResource(shk)
}

// Dependencies returns the list of resources [StorageHmacKey] depends_on.
func (shk *StorageHmacKey) Dependencies() terra.Dependencies {
	return shk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageHmacKey].
func (shk *StorageHmacKey) LifecycleManagement() *terra.Lifecycle {
	return shk.Lifecycle
}

// Attributes returns the attributes for [StorageHmacKey].
func (shk *StorageHmacKey) Attributes() storageHmacKeyAttributes {
	return storageHmacKeyAttributes{ref: terra.ReferenceResource(shk)}
}

// ImportState imports the given attribute values into [StorageHmacKey]'s state.
func (shk *StorageHmacKey) ImportState(av io.Reader) error {
	shk.state = &storageHmacKeyState{}
	if err := json.NewDecoder(av).Decode(shk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", shk.Type(), shk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageHmacKey] has state.
func (shk *StorageHmacKey) State() (*storageHmacKeyState, bool) {
	return shk.state, shk.state != nil
}

// StateMust returns the state for [StorageHmacKey]. Panics if the state is nil.
func (shk *StorageHmacKey) StateMust() *storageHmacKeyState {
	if shk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", shk.Type(), shk.LocalName()))
	}
	return shk.state
}

// StorageHmacKeyArgs contains the configurations for google_storage_hmac_key.
type StorageHmacKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceAccountEmail: string, required
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr" validate:"required"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Timeouts: optional
	Timeouts *storagehmackey.Timeouts `hcl:"timeouts,block"`
}
type storageHmacKeyAttributes struct {
	ref terra.Reference
}

// AccessId returns a reference to field access_id of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) AccessId() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("access_id"))
}

// Id returns a reference to field id of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("id"))
}

// Project returns a reference to field project of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("project"))
}

// Secret returns a reference to field secret of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("secret"))
}

// ServiceAccountEmail returns a reference to field service_account_email of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("service_account_email"))
}

// State returns a reference to field state of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("state"))
}

// TimeCreated returns a reference to field time_created of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("time_created"))
}

// Updated returns a reference to field updated of google_storage_hmac_key.
func (shk storageHmacKeyAttributes) Updated() terra.StringValue {
	return terra.ReferenceAsString(shk.ref.Append("updated"))
}

func (shk storageHmacKeyAttributes) Timeouts() storagehmackey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagehmackey.TimeoutsAttributes](shk.ref.Append("timeouts"))
}

type storageHmacKeyState struct {
	AccessId            string                        `json:"access_id"`
	Id                  string                        `json:"id"`
	Project             string                        `json:"project"`
	Secret              string                        `json:"secret"`
	ServiceAccountEmail string                        `json:"service_account_email"`
	State               string                        `json:"state"`
	TimeCreated         string                        `json:"time_created"`
	Updated             string                        `json:"updated"`
	Timeouts            *storagehmackey.TimeoutsState `json:"timeouts"`
}
