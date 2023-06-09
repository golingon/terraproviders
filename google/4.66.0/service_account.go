// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	serviceaccount "github.com/golingon/terraproviders/google/4.66.0/serviceaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceAccount creates a new instance of [ServiceAccount].
func NewServiceAccount(name string, args ServiceAccountArgs) *ServiceAccount {
	return &ServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceAccount)(nil)

// ServiceAccount represents the Terraform resource google_service_account.
type ServiceAccount struct {
	Name      string
	Args      ServiceAccountArgs
	state     *serviceAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceAccount].
func (sa *ServiceAccount) Type() string {
	return "google_service_account"
}

// LocalName returns the local name for [ServiceAccount].
func (sa *ServiceAccount) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [ServiceAccount].
func (sa *ServiceAccount) Configuration() interface{} {
	return sa.Args
}

// DependOn is used for other resources to depend on [ServiceAccount].
func (sa *ServiceAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(sa)
}

// Dependencies returns the list of resources [ServiceAccount] depends_on.
func (sa *ServiceAccount) Dependencies() terra.Dependencies {
	return sa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceAccount].
func (sa *ServiceAccount) LifecycleManagement() *terra.Lifecycle {
	return sa.Lifecycle
}

// Attributes returns the attributes for [ServiceAccount].
func (sa *ServiceAccount) Attributes() serviceAccountAttributes {
	return serviceAccountAttributes{ref: terra.ReferenceResource(sa)}
}

// ImportState imports the given attribute values into [ServiceAccount]'s state.
func (sa *ServiceAccount) ImportState(av io.Reader) error {
	sa.state = &serviceAccountState{}
	if err := json.NewDecoder(av).Decode(sa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sa.Type(), sa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceAccount] has state.
func (sa *ServiceAccount) State() (*serviceAccountState, bool) {
	return sa.state, sa.state != nil
}

// StateMust returns the state for [ServiceAccount]. Panics if the state is nil.
func (sa *ServiceAccount) StateMust() *serviceAccountState {
	if sa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sa.Type(), sa.LocalName()))
	}
	return sa.state
}

// ServiceAccountArgs contains the configurations for google_service_account.
type ServiceAccountArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *serviceaccount.Timeouts `hcl:"timeouts,block"`
}
type serviceAccountAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of google_service_account.
func (sa serviceAccountAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("account_id"))
}

// Description returns a reference to field description of google_service_account.
func (sa serviceAccountAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_service_account.
func (sa serviceAccountAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_service_account.
func (sa serviceAccountAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("display_name"))
}

// Email returns a reference to field email of google_service_account.
func (sa serviceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("email"))
}

// Id returns a reference to field id of google_service_account.
func (sa serviceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Member returns a reference to field member of google_service_account.
func (sa serviceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("member"))
}

// Name returns a reference to field name of google_service_account.
func (sa serviceAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

// Project returns a reference to field project of google_service_account.
func (sa serviceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("project"))
}

// UniqueId returns a reference to field unique_id of google_service_account.
func (sa serviceAccountAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("unique_id"))
}

func (sa serviceAccountAttributes) Timeouts() serviceaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceaccount.TimeoutsAttributes](sa.ref.Append("timeouts"))
}

type serviceAccountState struct {
	AccountId   string                        `json:"account_id"`
	Description string                        `json:"description"`
	Disabled    bool                          `json:"disabled"`
	DisplayName string                        `json:"display_name"`
	Email       string                        `json:"email"`
	Id          string                        `json:"id"`
	Member      string                        `json:"member"`
	Name        string                        `json:"name"`
	Project     string                        `json:"project"`
	UniqueId    string                        `json:"unique_id"`
	Timeouts    *serviceaccount.TimeoutsState `json:"timeouts"`
}
