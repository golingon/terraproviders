// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apikeyskey "github.com/golingon/terraproviders/googlebeta/4.74.0/apikeyskey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApikeysKey creates a new instance of [ApikeysKey].
func NewApikeysKey(name string, args ApikeysKeyArgs) *ApikeysKey {
	return &ApikeysKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApikeysKey)(nil)

// ApikeysKey represents the Terraform resource google_apikeys_key.
type ApikeysKey struct {
	Name      string
	Args      ApikeysKeyArgs
	state     *apikeysKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApikeysKey].
func (ak *ApikeysKey) Type() string {
	return "google_apikeys_key"
}

// LocalName returns the local name for [ApikeysKey].
func (ak *ApikeysKey) LocalName() string {
	return ak.Name
}

// Configuration returns the configuration (args) for [ApikeysKey].
func (ak *ApikeysKey) Configuration() interface{} {
	return ak.Args
}

// DependOn is used for other resources to depend on [ApikeysKey].
func (ak *ApikeysKey) DependOn() terra.Reference {
	return terra.ReferenceResource(ak)
}

// Dependencies returns the list of resources [ApikeysKey] depends_on.
func (ak *ApikeysKey) Dependencies() terra.Dependencies {
	return ak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApikeysKey].
func (ak *ApikeysKey) LifecycleManagement() *terra.Lifecycle {
	return ak.Lifecycle
}

// Attributes returns the attributes for [ApikeysKey].
func (ak *ApikeysKey) Attributes() apikeysKeyAttributes {
	return apikeysKeyAttributes{ref: terra.ReferenceResource(ak)}
}

// ImportState imports the given attribute values into [ApikeysKey]'s state.
func (ak *ApikeysKey) ImportState(av io.Reader) error {
	ak.state = &apikeysKeyState{}
	if err := json.NewDecoder(av).Decode(ak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ak.Type(), ak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApikeysKey] has state.
func (ak *ApikeysKey) State() (*apikeysKeyState, bool) {
	return ak.state, ak.state != nil
}

// StateMust returns the state for [ApikeysKey]. Panics if the state is nil.
func (ak *ApikeysKey) StateMust() *apikeysKeyState {
	if ak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ak.Type(), ak.LocalName()))
	}
	return ak.state
}

// ApikeysKeyArgs contains the configurations for google_apikeys_key.
type ApikeysKeyArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Restrictions: optional
	Restrictions *apikeyskey.Restrictions `hcl:"restrictions,block"`
	// Timeouts: optional
	Timeouts *apikeyskey.Timeouts `hcl:"timeouts,block"`
}
type apikeysKeyAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_apikeys_key.
func (ak apikeysKeyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("display_name"))
}

// Id returns a reference to field id of google_apikeys_key.
func (ak apikeysKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("id"))
}

// KeyString returns a reference to field key_string of google_apikeys_key.
func (ak apikeysKeyAttributes) KeyString() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("key_string"))
}

// Name returns a reference to field name of google_apikeys_key.
func (ak apikeysKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("name"))
}

// Project returns a reference to field project of google_apikeys_key.
func (ak apikeysKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("project"))
}

// Uid returns a reference to field uid of google_apikeys_key.
func (ak apikeysKeyAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ak.ref.Append("uid"))
}

func (ak apikeysKeyAttributes) Restrictions() terra.ListValue[apikeyskey.RestrictionsAttributes] {
	return terra.ReferenceAsList[apikeyskey.RestrictionsAttributes](ak.ref.Append("restrictions"))
}

func (ak apikeysKeyAttributes) Timeouts() apikeyskey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apikeyskey.TimeoutsAttributes](ak.ref.Append("timeouts"))
}

type apikeysKeyState struct {
	DisplayName  string                         `json:"display_name"`
	Id           string                         `json:"id"`
	KeyString    string                         `json:"key_string"`
	Name         string                         `json:"name"`
	Project      string                         `json:"project"`
	Uid          string                         `json:"uid"`
	Restrictions []apikeyskey.RestrictionsState `json:"restrictions"`
	Timeouts     *apikeyskey.TimeoutsState      `json:"timeouts"`
}
