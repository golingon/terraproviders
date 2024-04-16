// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_project_default_service_accounts

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

// Resource represents the Terraform resource google_project_default_service_accounts.
type Resource struct {
	Name      string
	Args      Args
	state     *googleProjectDefaultServiceAccountsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gpdsa *Resource) Type() string {
	return "google_project_default_service_accounts"
}

// LocalName returns the local name for [Resource].
func (gpdsa *Resource) LocalName() string {
	return gpdsa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gpdsa *Resource) Configuration() interface{} {
	return gpdsa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gpdsa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gpdsa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gpdsa *Resource) Dependencies() terra.Dependencies {
	return gpdsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gpdsa *Resource) LifecycleManagement() *terra.Lifecycle {
	return gpdsa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gpdsa *Resource) Attributes() googleProjectDefaultServiceAccountsAttributes {
	return googleProjectDefaultServiceAccountsAttributes{ref: terra.ReferenceResource(gpdsa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gpdsa *Resource) ImportState(state io.Reader) error {
	gpdsa.state = &googleProjectDefaultServiceAccountsState{}
	if err := json.NewDecoder(state).Decode(gpdsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gpdsa.Type(), gpdsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gpdsa *Resource) State() (*googleProjectDefaultServiceAccountsState, bool) {
	return gpdsa.state, gpdsa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gpdsa *Resource) StateMust() *googleProjectDefaultServiceAccountsState {
	if gpdsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gpdsa.Type(), gpdsa.LocalName()))
	}
	return gpdsa.state
}

// Args contains the configurations for google_project_default_service_accounts.
type Args struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// RestorePolicy: string, optional
	RestorePolicy terra.StringValue `hcl:"restore_policy,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleProjectDefaultServiceAccountsAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_project_default_service_accounts.
func (gpdsa googleProjectDefaultServiceAccountsAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(gpdsa.ref.Append("action"))
}

// Id returns a reference to field id of google_project_default_service_accounts.
func (gpdsa googleProjectDefaultServiceAccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpdsa.ref.Append("id"))
}

// Project returns a reference to field project of google_project_default_service_accounts.
func (gpdsa googleProjectDefaultServiceAccountsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gpdsa.ref.Append("project"))
}

// RestorePolicy returns a reference to field restore_policy of google_project_default_service_accounts.
func (gpdsa googleProjectDefaultServiceAccountsAttributes) RestorePolicy() terra.StringValue {
	return terra.ReferenceAsString(gpdsa.ref.Append("restore_policy"))
}

// ServiceAccounts returns a reference to field service_accounts of google_project_default_service_accounts.
func (gpdsa googleProjectDefaultServiceAccountsAttributes) ServiceAccounts() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gpdsa.ref.Append("service_accounts"))
}

func (gpdsa googleProjectDefaultServiceAccountsAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gpdsa.ref.Append("timeouts"))
}

type googleProjectDefaultServiceAccountsState struct {
	Action          string            `json:"action"`
	Id              string            `json:"id"`
	Project         string            `json:"project"`
	RestorePolicy   string            `json:"restore_policy"`
	ServiceAccounts map[string]string `json:"service_accounts"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
