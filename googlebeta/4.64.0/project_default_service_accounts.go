// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	projectdefaultserviceaccounts "github.com/golingon/terraproviders/googlebeta/4.64.0/projectdefaultserviceaccounts"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectDefaultServiceAccounts creates a new instance of [ProjectDefaultServiceAccounts].
func NewProjectDefaultServiceAccounts(name string, args ProjectDefaultServiceAccountsArgs) *ProjectDefaultServiceAccounts {
	return &ProjectDefaultServiceAccounts{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectDefaultServiceAccounts)(nil)

// ProjectDefaultServiceAccounts represents the Terraform resource google_project_default_service_accounts.
type ProjectDefaultServiceAccounts struct {
	Name      string
	Args      ProjectDefaultServiceAccountsArgs
	state     *projectDefaultServiceAccountsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) Type() string {
	return "google_project_default_service_accounts"
}

// LocalName returns the local name for [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) LocalName() string {
	return pdsa.Name
}

// Configuration returns the configuration (args) for [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) Configuration() interface{} {
	return pdsa.Args
}

// DependOn is used for other resources to depend on [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) DependOn() terra.Reference {
	return terra.ReferenceResource(pdsa)
}

// Dependencies returns the list of resources [ProjectDefaultServiceAccounts] depends_on.
func (pdsa *ProjectDefaultServiceAccounts) Dependencies() terra.Dependencies {
	return pdsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) LifecycleManagement() *terra.Lifecycle {
	return pdsa.Lifecycle
}

// Attributes returns the attributes for [ProjectDefaultServiceAccounts].
func (pdsa *ProjectDefaultServiceAccounts) Attributes() projectDefaultServiceAccountsAttributes {
	return projectDefaultServiceAccountsAttributes{ref: terra.ReferenceResource(pdsa)}
}

// ImportState imports the given attribute values into [ProjectDefaultServiceAccounts]'s state.
func (pdsa *ProjectDefaultServiceAccounts) ImportState(av io.Reader) error {
	pdsa.state = &projectDefaultServiceAccountsState{}
	if err := json.NewDecoder(av).Decode(pdsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdsa.Type(), pdsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectDefaultServiceAccounts] has state.
func (pdsa *ProjectDefaultServiceAccounts) State() (*projectDefaultServiceAccountsState, bool) {
	return pdsa.state, pdsa.state != nil
}

// StateMust returns the state for [ProjectDefaultServiceAccounts]. Panics if the state is nil.
func (pdsa *ProjectDefaultServiceAccounts) StateMust() *projectDefaultServiceAccountsState {
	if pdsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdsa.Type(), pdsa.LocalName()))
	}
	return pdsa.state
}

// ProjectDefaultServiceAccountsArgs contains the configurations for google_project_default_service_accounts.
type ProjectDefaultServiceAccountsArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// RestorePolicy: string, optional
	RestorePolicy terra.StringValue `hcl:"restore_policy,attr"`
	// Timeouts: optional
	Timeouts *projectdefaultserviceaccounts.Timeouts `hcl:"timeouts,block"`
}
type projectDefaultServiceAccountsAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_project_default_service_accounts.
func (pdsa projectDefaultServiceAccountsAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(pdsa.ref.Append("action"))
}

// Id returns a reference to field id of google_project_default_service_accounts.
func (pdsa projectDefaultServiceAccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdsa.ref.Append("id"))
}

// Project returns a reference to field project of google_project_default_service_accounts.
func (pdsa projectDefaultServiceAccountsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pdsa.ref.Append("project"))
}

// RestorePolicy returns a reference to field restore_policy of google_project_default_service_accounts.
func (pdsa projectDefaultServiceAccountsAttributes) RestorePolicy() terra.StringValue {
	return terra.ReferenceAsString(pdsa.ref.Append("restore_policy"))
}

// ServiceAccounts returns a reference to field service_accounts of google_project_default_service_accounts.
func (pdsa projectDefaultServiceAccountsAttributes) ServiceAccounts() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdsa.ref.Append("service_accounts"))
}

func (pdsa projectDefaultServiceAccountsAttributes) Timeouts() projectdefaultserviceaccounts.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectdefaultserviceaccounts.TimeoutsAttributes](pdsa.ref.Append("timeouts"))
}

type projectDefaultServiceAccountsState struct {
	Action          string                                       `json:"action"`
	Id              string                                       `json:"id"`
	Project         string                                       `json:"project"`
	RestorePolicy   string                                       `json:"restore_policy"`
	ServiceAccounts map[string]string                            `json:"service_accounts"`
	Timeouts        *projectdefaultserviceaccounts.TimeoutsState `json:"timeouts"`
}
