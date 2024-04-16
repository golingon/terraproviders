// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_public_ca_external_account_key

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

// Resource represents the Terraform resource google_public_ca_external_account_key.
type Resource struct {
	Name      string
	Args      Args
	state     *googlePublicCaExternalAccountKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gpceak *Resource) Type() string {
	return "google_public_ca_external_account_key"
}

// LocalName returns the local name for [Resource].
func (gpceak *Resource) LocalName() string {
	return gpceak.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gpceak *Resource) Configuration() interface{} {
	return gpceak.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gpceak *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gpceak)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gpceak *Resource) Dependencies() terra.Dependencies {
	return gpceak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gpceak *Resource) LifecycleManagement() *terra.Lifecycle {
	return gpceak.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gpceak *Resource) Attributes() googlePublicCaExternalAccountKeyAttributes {
	return googlePublicCaExternalAccountKeyAttributes{ref: terra.ReferenceResource(gpceak)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gpceak *Resource) ImportState(state io.Reader) error {
	gpceak.state = &googlePublicCaExternalAccountKeyState{}
	if err := json.NewDecoder(state).Decode(gpceak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gpceak.Type(), gpceak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gpceak *Resource) State() (*googlePublicCaExternalAccountKeyState, bool) {
	return gpceak.state, gpceak.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gpceak *Resource) StateMust() *googlePublicCaExternalAccountKeyState {
	if gpceak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gpceak.Type(), gpceak.LocalName()))
	}
	return gpceak.state
}

// Args contains the configurations for google_public_ca_external_account_key.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googlePublicCaExternalAccountKeyAttributes struct {
	ref terra.Reference
}

// B64MacKey returns a reference to field b64_mac_key of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) B64MacKey() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("b64_mac_key"))
}

// Id returns a reference to field id of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("id"))
}

// KeyId returns a reference to field key_id of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("key_id"))
}

// Location returns a reference to field location of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("location"))
}

// Name returns a reference to field name of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("name"))
}

// Project returns a reference to field project of google_public_ca_external_account_key.
func (gpceak googlePublicCaExternalAccountKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gpceak.ref.Append("project"))
}

func (gpceak googlePublicCaExternalAccountKeyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gpceak.ref.Append("timeouts"))
}

type googlePublicCaExternalAccountKeyState struct {
	B64MacKey string         `json:"b64_mac_key"`
	Id        string         `json:"id"`
	KeyId     string         `json:"key_id"`
	Location  string         `json:"location"`
	Name      string         `json:"name"`
	Project   string         `json:"project"`
	Timeouts  *TimeoutsState `json:"timeouts"`
}
