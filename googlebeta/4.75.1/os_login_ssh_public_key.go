// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	osloginsshpublickey "github.com/golingon/terraproviders/googlebeta/4.75.1/osloginsshpublickey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOsLoginSshPublicKey creates a new instance of [OsLoginSshPublicKey].
func NewOsLoginSshPublicKey(name string, args OsLoginSshPublicKeyArgs) *OsLoginSshPublicKey {
	return &OsLoginSshPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OsLoginSshPublicKey)(nil)

// OsLoginSshPublicKey represents the Terraform resource google_os_login_ssh_public_key.
type OsLoginSshPublicKey struct {
	Name      string
	Args      OsLoginSshPublicKeyArgs
	state     *osLoginSshPublicKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) Type() string {
	return "google_os_login_ssh_public_key"
}

// LocalName returns the local name for [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) LocalName() string {
	return olspk.Name
}

// Configuration returns the configuration (args) for [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) Configuration() interface{} {
	return olspk.Args
}

// DependOn is used for other resources to depend on [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) DependOn() terra.Reference {
	return terra.ReferenceResource(olspk)
}

// Dependencies returns the list of resources [OsLoginSshPublicKey] depends_on.
func (olspk *OsLoginSshPublicKey) Dependencies() terra.Dependencies {
	return olspk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) LifecycleManagement() *terra.Lifecycle {
	return olspk.Lifecycle
}

// Attributes returns the attributes for [OsLoginSshPublicKey].
func (olspk *OsLoginSshPublicKey) Attributes() osLoginSshPublicKeyAttributes {
	return osLoginSshPublicKeyAttributes{ref: terra.ReferenceResource(olspk)}
}

// ImportState imports the given attribute values into [OsLoginSshPublicKey]'s state.
func (olspk *OsLoginSshPublicKey) ImportState(av io.Reader) error {
	olspk.state = &osLoginSshPublicKeyState{}
	if err := json.NewDecoder(av).Decode(olspk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", olspk.Type(), olspk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OsLoginSshPublicKey] has state.
func (olspk *OsLoginSshPublicKey) State() (*osLoginSshPublicKeyState, bool) {
	return olspk.state, olspk.state != nil
}

// StateMust returns the state for [OsLoginSshPublicKey]. Panics if the state is nil.
func (olspk *OsLoginSshPublicKey) StateMust() *osLoginSshPublicKeyState {
	if olspk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", olspk.Type(), olspk.LocalName()))
	}
	return olspk.state
}

// OsLoginSshPublicKeyArgs contains the configurations for google_os_login_ssh_public_key.
type OsLoginSshPublicKeyArgs struct {
	// ExpirationTimeUsec: string, optional
	ExpirationTimeUsec terra.StringValue `hcl:"expiration_time_usec,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *osloginsshpublickey.Timeouts `hcl:"timeouts,block"`
}
type osLoginSshPublicKeyAttributes struct {
	ref terra.Reference
}

// ExpirationTimeUsec returns a reference to field expiration_time_usec of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) ExpirationTimeUsec() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("expiration_time_usec"))
}

// Fingerprint returns a reference to field fingerprint of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("id"))
}

// Key returns a reference to field key of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("key"))
}

// Project returns a reference to field project of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("project"))
}

// User returns a reference to field user of google_os_login_ssh_public_key.
func (olspk osLoginSshPublicKeyAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(olspk.ref.Append("user"))
}

func (olspk osLoginSshPublicKeyAttributes) Timeouts() osloginsshpublickey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[osloginsshpublickey.TimeoutsAttributes](olspk.ref.Append("timeouts"))
}

type osLoginSshPublicKeyState struct {
	ExpirationTimeUsec string                             `json:"expiration_time_usec"`
	Fingerprint        string                             `json:"fingerprint"`
	Id                 string                             `json:"id"`
	Key                string                             `json:"key"`
	Project            string                             `json:"project"`
	User               string                             `json:"user"`
	Timeouts           *osloginsshpublickey.TimeoutsState `json:"timeouts"`
}
