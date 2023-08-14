// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	publiccaexternalaccountkey "github.com/golingon/terraproviders/googlebeta/4.74.0/publiccaexternalaccountkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPublicCaExternalAccountKey creates a new instance of [PublicCaExternalAccountKey].
func NewPublicCaExternalAccountKey(name string, args PublicCaExternalAccountKeyArgs) *PublicCaExternalAccountKey {
	return &PublicCaExternalAccountKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PublicCaExternalAccountKey)(nil)

// PublicCaExternalAccountKey represents the Terraform resource google_public_ca_external_account_key.
type PublicCaExternalAccountKey struct {
	Name      string
	Args      PublicCaExternalAccountKeyArgs
	state     *publicCaExternalAccountKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) Type() string {
	return "google_public_ca_external_account_key"
}

// LocalName returns the local name for [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) LocalName() string {
	return pceak.Name
}

// Configuration returns the configuration (args) for [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) Configuration() interface{} {
	return pceak.Args
}

// DependOn is used for other resources to depend on [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) DependOn() terra.Reference {
	return terra.ReferenceResource(pceak)
}

// Dependencies returns the list of resources [PublicCaExternalAccountKey] depends_on.
func (pceak *PublicCaExternalAccountKey) Dependencies() terra.Dependencies {
	return pceak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) LifecycleManagement() *terra.Lifecycle {
	return pceak.Lifecycle
}

// Attributes returns the attributes for [PublicCaExternalAccountKey].
func (pceak *PublicCaExternalAccountKey) Attributes() publicCaExternalAccountKeyAttributes {
	return publicCaExternalAccountKeyAttributes{ref: terra.ReferenceResource(pceak)}
}

// ImportState imports the given attribute values into [PublicCaExternalAccountKey]'s state.
func (pceak *PublicCaExternalAccountKey) ImportState(av io.Reader) error {
	pceak.state = &publicCaExternalAccountKeyState{}
	if err := json.NewDecoder(av).Decode(pceak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pceak.Type(), pceak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PublicCaExternalAccountKey] has state.
func (pceak *PublicCaExternalAccountKey) State() (*publicCaExternalAccountKeyState, bool) {
	return pceak.state, pceak.state != nil
}

// StateMust returns the state for [PublicCaExternalAccountKey]. Panics if the state is nil.
func (pceak *PublicCaExternalAccountKey) StateMust() *publicCaExternalAccountKeyState {
	if pceak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pceak.Type(), pceak.LocalName()))
	}
	return pceak.state
}

// PublicCaExternalAccountKeyArgs contains the configurations for google_public_ca_external_account_key.
type PublicCaExternalAccountKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *publiccaexternalaccountkey.Timeouts `hcl:"timeouts,block"`
}
type publicCaExternalAccountKeyAttributes struct {
	ref terra.Reference
}

// B64MacKey returns a reference to field b64_mac_key of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) B64MacKey() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("b64_mac_key"))
}

// Id returns a reference to field id of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("id"))
}

// KeyId returns a reference to field key_id of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("key_id"))
}

// Location returns a reference to field location of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("location"))
}

// Name returns a reference to field name of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("name"))
}

// Project returns a reference to field project of google_public_ca_external_account_key.
func (pceak publicCaExternalAccountKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pceak.ref.Append("project"))
}

func (pceak publicCaExternalAccountKeyAttributes) Timeouts() publiccaexternalaccountkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[publiccaexternalaccountkey.TimeoutsAttributes](pceak.ref.Append("timeouts"))
}

type publicCaExternalAccountKeyState struct {
	B64MacKey string                                    `json:"b64_mac_key"`
	Id        string                                    `json:"id"`
	KeyId     string                                    `json:"key_id"`
	Location  string                                    `json:"location"`
	Name      string                                    `json:"name"`
	Project   string                                    `json:"project"`
	Timeouts  *publiccaexternalaccountkey.TimeoutsState `json:"timeouts"`
}