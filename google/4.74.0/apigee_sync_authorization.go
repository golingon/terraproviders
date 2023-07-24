// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeesyncauthorization "github.com/golingon/terraproviders/google/4.74.0/apigeesyncauthorization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeSyncAuthorization creates a new instance of [ApigeeSyncAuthorization].
func NewApigeeSyncAuthorization(name string, args ApigeeSyncAuthorizationArgs) *ApigeeSyncAuthorization {
	return &ApigeeSyncAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeSyncAuthorization)(nil)

// ApigeeSyncAuthorization represents the Terraform resource google_apigee_sync_authorization.
type ApigeeSyncAuthorization struct {
	Name      string
	Args      ApigeeSyncAuthorizationArgs
	state     *apigeeSyncAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) Type() string {
	return "google_apigee_sync_authorization"
}

// LocalName returns the local name for [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) LocalName() string {
	return asa.Name
}

// Configuration returns the configuration (args) for [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) Configuration() interface{} {
	return asa.Args
}

// DependOn is used for other resources to depend on [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(asa)
}

// Dependencies returns the list of resources [ApigeeSyncAuthorization] depends_on.
func (asa *ApigeeSyncAuthorization) Dependencies() terra.Dependencies {
	return asa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) LifecycleManagement() *terra.Lifecycle {
	return asa.Lifecycle
}

// Attributes returns the attributes for [ApigeeSyncAuthorization].
func (asa *ApigeeSyncAuthorization) Attributes() apigeeSyncAuthorizationAttributes {
	return apigeeSyncAuthorizationAttributes{ref: terra.ReferenceResource(asa)}
}

// ImportState imports the given attribute values into [ApigeeSyncAuthorization]'s state.
func (asa *ApigeeSyncAuthorization) ImportState(av io.Reader) error {
	asa.state = &apigeeSyncAuthorizationState{}
	if err := json.NewDecoder(av).Decode(asa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asa.Type(), asa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeSyncAuthorization] has state.
func (asa *ApigeeSyncAuthorization) State() (*apigeeSyncAuthorizationState, bool) {
	return asa.state, asa.state != nil
}

// StateMust returns the state for [ApigeeSyncAuthorization]. Panics if the state is nil.
func (asa *ApigeeSyncAuthorization) StateMust() *apigeeSyncAuthorizationState {
	if asa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asa.Type(), asa.LocalName()))
	}
	return asa.state
}

// ApigeeSyncAuthorizationArgs contains the configurations for google_apigee_sync_authorization.
type ApigeeSyncAuthorizationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identities: list of string, required
	Identities terra.ListValue[terra.StringValue] `hcl:"identities,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeesyncauthorization.Timeouts `hcl:"timeouts,block"`
}
type apigeeSyncAuthorizationAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_apigee_sync_authorization.
func (asa apigeeSyncAuthorizationAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("etag"))
}

// Id returns a reference to field id of google_apigee_sync_authorization.
func (asa apigeeSyncAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("id"))
}

// Identities returns a reference to field identities of google_apigee_sync_authorization.
func (asa apigeeSyncAuthorizationAttributes) Identities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asa.ref.Append("identities"))
}

// Name returns a reference to field name of google_apigee_sync_authorization.
func (asa apigeeSyncAuthorizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("name"))
}

func (asa apigeeSyncAuthorizationAttributes) Timeouts() apigeesyncauthorization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeesyncauthorization.TimeoutsAttributes](asa.ref.Append("timeouts"))
}

type apigeeSyncAuthorizationState struct {
	Etag       string                                 `json:"etag"`
	Id         string                                 `json:"id"`
	Identities []string                               `json:"identities"`
	Name       string                                 `json:"name"`
	Timeouts   *apigeesyncauthorization.TimeoutsState `json:"timeouts"`
}
