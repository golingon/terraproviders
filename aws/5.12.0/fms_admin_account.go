// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fmsadminaccount "github.com/golingon/terraproviders/aws/5.12.0/fmsadminaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFmsAdminAccount creates a new instance of [FmsAdminAccount].
func NewFmsAdminAccount(name string, args FmsAdminAccountArgs) *FmsAdminAccount {
	return &FmsAdminAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FmsAdminAccount)(nil)

// FmsAdminAccount represents the Terraform resource aws_fms_admin_account.
type FmsAdminAccount struct {
	Name      string
	Args      FmsAdminAccountArgs
	state     *fmsAdminAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FmsAdminAccount].
func (faa *FmsAdminAccount) Type() string {
	return "aws_fms_admin_account"
}

// LocalName returns the local name for [FmsAdminAccount].
func (faa *FmsAdminAccount) LocalName() string {
	return faa.Name
}

// Configuration returns the configuration (args) for [FmsAdminAccount].
func (faa *FmsAdminAccount) Configuration() interface{} {
	return faa.Args
}

// DependOn is used for other resources to depend on [FmsAdminAccount].
func (faa *FmsAdminAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(faa)
}

// Dependencies returns the list of resources [FmsAdminAccount] depends_on.
func (faa *FmsAdminAccount) Dependencies() terra.Dependencies {
	return faa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FmsAdminAccount].
func (faa *FmsAdminAccount) LifecycleManagement() *terra.Lifecycle {
	return faa.Lifecycle
}

// Attributes returns the attributes for [FmsAdminAccount].
func (faa *FmsAdminAccount) Attributes() fmsAdminAccountAttributes {
	return fmsAdminAccountAttributes{ref: terra.ReferenceResource(faa)}
}

// ImportState imports the given attribute values into [FmsAdminAccount]'s state.
func (faa *FmsAdminAccount) ImportState(av io.Reader) error {
	faa.state = &fmsAdminAccountState{}
	if err := json.NewDecoder(av).Decode(faa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", faa.Type(), faa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FmsAdminAccount] has state.
func (faa *FmsAdminAccount) State() (*fmsAdminAccountState, bool) {
	return faa.state, faa.state != nil
}

// StateMust returns the state for [FmsAdminAccount]. Panics if the state is nil.
func (faa *FmsAdminAccount) StateMust() *fmsAdminAccountState {
	if faa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", faa.Type(), faa.LocalName()))
	}
	return faa.state
}

// FmsAdminAccountArgs contains the configurations for aws_fms_admin_account.
type FmsAdminAccountArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *fmsadminaccount.Timeouts `hcl:"timeouts,block"`
}
type fmsAdminAccountAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_fms_admin_account.
func (faa fmsAdminAccountAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("account_id"))
}

// Id returns a reference to field id of aws_fms_admin_account.
func (faa fmsAdminAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("id"))
}

func (faa fmsAdminAccountAttributes) Timeouts() fmsadminaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fmsadminaccount.TimeoutsAttributes](faa.ref.Append("timeouts"))
}

type fmsAdminAccountState struct {
	AccountId string                         `json:"account_id"`
	Id        string                         `json:"id"`
	Timeouts  *fmsadminaccount.TimeoutsState `json:"timeouts"`
}
