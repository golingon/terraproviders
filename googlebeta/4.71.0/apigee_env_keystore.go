// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeeenvkeystore "github.com/golingon/terraproviders/googlebeta/4.71.0/apigeeenvkeystore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvKeystore creates a new instance of [ApigeeEnvKeystore].
func NewApigeeEnvKeystore(name string, args ApigeeEnvKeystoreArgs) *ApigeeEnvKeystore {
	return &ApigeeEnvKeystore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvKeystore)(nil)

// ApigeeEnvKeystore represents the Terraform resource google_apigee_env_keystore.
type ApigeeEnvKeystore struct {
	Name      string
	Args      ApigeeEnvKeystoreArgs
	state     *apigeeEnvKeystoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) Type() string {
	return "google_apigee_env_keystore"
}

// LocalName returns the local name for [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) LocalName() string {
	return aek.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) Configuration() interface{} {
	return aek.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) DependOn() terra.Reference {
	return terra.ReferenceResource(aek)
}

// Dependencies returns the list of resources [ApigeeEnvKeystore] depends_on.
func (aek *ApigeeEnvKeystore) Dependencies() terra.Dependencies {
	return aek.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) LifecycleManagement() *terra.Lifecycle {
	return aek.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvKeystore].
func (aek *ApigeeEnvKeystore) Attributes() apigeeEnvKeystoreAttributes {
	return apigeeEnvKeystoreAttributes{ref: terra.ReferenceResource(aek)}
}

// ImportState imports the given attribute values into [ApigeeEnvKeystore]'s state.
func (aek *ApigeeEnvKeystore) ImportState(av io.Reader) error {
	aek.state = &apigeeEnvKeystoreState{}
	if err := json.NewDecoder(av).Decode(aek.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aek.Type(), aek.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvKeystore] has state.
func (aek *ApigeeEnvKeystore) State() (*apigeeEnvKeystoreState, bool) {
	return aek.state, aek.state != nil
}

// StateMust returns the state for [ApigeeEnvKeystore]. Panics if the state is nil.
func (aek *ApigeeEnvKeystore) StateMust() *apigeeEnvKeystoreState {
	if aek.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aek.Type(), aek.LocalName()))
	}
	return aek.state
}

// ApigeeEnvKeystoreArgs contains the configurations for google_apigee_env_keystore.
type ApigeeEnvKeystoreArgs struct {
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Timeouts: optional
	Timeouts *apigeeenvkeystore.Timeouts `hcl:"timeouts,block"`
}
type apigeeEnvKeystoreAttributes struct {
	ref terra.Reference
}

// Aliases returns a reference to field aliases of google_apigee_env_keystore.
func (aek apigeeEnvKeystoreAttributes) Aliases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aek.ref.Append("aliases"))
}

// EnvId returns a reference to field env_id of google_apigee_env_keystore.
func (aek apigeeEnvKeystoreAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(aek.ref.Append("env_id"))
}

// Id returns a reference to field id of google_apigee_env_keystore.
func (aek apigeeEnvKeystoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aek.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_env_keystore.
func (aek apigeeEnvKeystoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aek.ref.Append("name"))
}

func (aek apigeeEnvKeystoreAttributes) Timeouts() apigeeenvkeystore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeenvkeystore.TimeoutsAttributes](aek.ref.Append("timeouts"))
}

type apigeeEnvKeystoreState struct {
	Aliases  []string                         `json:"aliases"`
	EnvId    string                           `json:"env_id"`
	Id       string                           `json:"id"`
	Name     string                           `json:"name"`
	Timeouts *apigeeenvkeystore.TimeoutsState `json:"timeouts"`
}
