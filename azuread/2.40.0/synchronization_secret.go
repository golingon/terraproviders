// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	synchronizationsecret "github.com/golingon/terraproviders/azuread/2.40.0/synchronizationsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynchronizationSecret creates a new instance of [SynchronizationSecret].
func NewSynchronizationSecret(name string, args SynchronizationSecretArgs) *SynchronizationSecret {
	return &SynchronizationSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynchronizationSecret)(nil)

// SynchronizationSecret represents the Terraform resource azuread_synchronization_secret.
type SynchronizationSecret struct {
	Name      string
	Args      SynchronizationSecretArgs
	state     *synchronizationSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynchronizationSecret].
func (ss *SynchronizationSecret) Type() string {
	return "azuread_synchronization_secret"
}

// LocalName returns the local name for [SynchronizationSecret].
func (ss *SynchronizationSecret) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SynchronizationSecret].
func (ss *SynchronizationSecret) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SynchronizationSecret].
func (ss *SynchronizationSecret) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SynchronizationSecret] depends_on.
func (ss *SynchronizationSecret) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynchronizationSecret].
func (ss *SynchronizationSecret) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SynchronizationSecret].
func (ss *SynchronizationSecret) Attributes() synchronizationSecretAttributes {
	return synchronizationSecretAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SynchronizationSecret]'s state.
func (ss *SynchronizationSecret) ImportState(av io.Reader) error {
	ss.state = &synchronizationSecretState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynchronizationSecret] has state.
func (ss *SynchronizationSecret) State() (*synchronizationSecretState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SynchronizationSecret]. Panics if the state is nil.
func (ss *SynchronizationSecret) StateMust() *synchronizationSecretState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SynchronizationSecretArgs contains the configurations for azuread_synchronization_secret.
type SynchronizationSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
	// Credential: min=0
	Credential []synchronizationsecret.Credential `hcl:"credential,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *synchronizationsecret.Timeouts `hcl:"timeouts,block"`
}
type synchronizationSecretAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azuread_synchronization_secret.
func (ss synchronizationSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azuread_synchronization_secret.
func (ss synchronizationSecretAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("service_principal_id"))
}

func (ss synchronizationSecretAttributes) Credential() terra.ListValue[synchronizationsecret.CredentialAttributes] {
	return terra.ReferenceAsList[synchronizationsecret.CredentialAttributes](ss.ref.Append("credential"))
}

func (ss synchronizationSecretAttributes) Timeouts() synchronizationsecret.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synchronizationsecret.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type synchronizationSecretState struct {
	Id                 string                                  `json:"id"`
	ServicePrincipalId string                                  `json:"service_principal_id"`
	Credential         []synchronizationsecret.CredentialState `json:"credential"`
	Timeouts           *synchronizationsecret.TimeoutsState    `json:"timeouts"`
}
