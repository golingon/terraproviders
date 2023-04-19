// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	applicationfederatedidentitycredential "github.com/golingon/terraproviders/azuread/2.37.1/applicationfederatedidentitycredential"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationFederatedIdentityCredential creates a new instance of [ApplicationFederatedIdentityCredential].
func NewApplicationFederatedIdentityCredential(name string, args ApplicationFederatedIdentityCredentialArgs) *ApplicationFederatedIdentityCredential {
	return &ApplicationFederatedIdentityCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationFederatedIdentityCredential)(nil)

// ApplicationFederatedIdentityCredential represents the Terraform resource azuread_application_federated_identity_credential.
type ApplicationFederatedIdentityCredential struct {
	Name      string
	Args      ApplicationFederatedIdentityCredentialArgs
	state     *applicationFederatedIdentityCredentialState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) Type() string {
	return "azuread_application_federated_identity_credential"
}

// LocalName returns the local name for [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) LocalName() string {
	return afic.Name
}

// Configuration returns the configuration (args) for [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) Configuration() interface{} {
	return afic.Args
}

// DependOn is used for other resources to depend on [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(afic)
}

// Dependencies returns the list of resources [ApplicationFederatedIdentityCredential] depends_on.
func (afic *ApplicationFederatedIdentityCredential) Dependencies() terra.Dependencies {
	return afic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) LifecycleManagement() *terra.Lifecycle {
	return afic.Lifecycle
}

// Attributes returns the attributes for [ApplicationFederatedIdentityCredential].
func (afic *ApplicationFederatedIdentityCredential) Attributes() applicationFederatedIdentityCredentialAttributes {
	return applicationFederatedIdentityCredentialAttributes{ref: terra.ReferenceResource(afic)}
}

// ImportState imports the given attribute values into [ApplicationFederatedIdentityCredential]'s state.
func (afic *ApplicationFederatedIdentityCredential) ImportState(av io.Reader) error {
	afic.state = &applicationFederatedIdentityCredentialState{}
	if err := json.NewDecoder(av).Decode(afic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afic.Type(), afic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationFederatedIdentityCredential] has state.
func (afic *ApplicationFederatedIdentityCredential) State() (*applicationFederatedIdentityCredentialState, bool) {
	return afic.state, afic.state != nil
}

// StateMust returns the state for [ApplicationFederatedIdentityCredential]. Panics if the state is nil.
func (afic *ApplicationFederatedIdentityCredential) StateMust() *applicationFederatedIdentityCredentialState {
	if afic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afic.Type(), afic.LocalName()))
	}
	return afic.state
}

// ApplicationFederatedIdentityCredentialArgs contains the configurations for azuread_application_federated_identity_credential.
type ApplicationFederatedIdentityCredentialArgs struct {
	// ApplicationObjectId: string, required
	ApplicationObjectId terra.StringValue `hcl:"application_object_id,attr" validate:"required"`
	// Audiences: list of string, required
	Audiences terra.ListValue[terra.StringValue] `hcl:"audiences,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *applicationfederatedidentitycredential.Timeouts `hcl:"timeouts,block"`
}
type applicationFederatedIdentityCredentialAttributes struct {
	ref terra.Reference
}

// ApplicationObjectId returns a reference to field application_object_id of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) ApplicationObjectId() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("application_object_id"))
}

// Audiences returns a reference to field audiences of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) Audiences() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](afic.ref.Append("audiences"))
}

// CredentialId returns a reference to field credential_id of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) CredentialId() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("credential_id"))
}

// Description returns a reference to field description of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("id"))
}

// Issuer returns a reference to field issuer of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("issuer"))
}

// Subject returns a reference to field subject of azuread_application_federated_identity_credential.
func (afic applicationFederatedIdentityCredentialAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(afic.ref.Append("subject"))
}

func (afic applicationFederatedIdentityCredentialAttributes) Timeouts() applicationfederatedidentitycredential.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationfederatedidentitycredential.TimeoutsAttributes](afic.ref.Append("timeouts"))
}

type applicationFederatedIdentityCredentialState struct {
	ApplicationObjectId string                                                `json:"application_object_id"`
	Audiences           []string                                              `json:"audiences"`
	CredentialId        string                                                `json:"credential_id"`
	Description         string                                                `json:"description"`
	DisplayName         string                                                `json:"display_name"`
	Id                  string                                                `json:"id"`
	Issuer              string                                                `json:"issuer"`
	Subject             string                                                `json:"subject"`
	Timeouts            *applicationfederatedidentitycredential.TimeoutsState `json:"timeouts"`
}
