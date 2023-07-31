// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	federatedidentitycredential "github.com/golingon/terraproviders/azurerm/3.67.0/federatedidentitycredential"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFederatedIdentityCredential creates a new instance of [FederatedIdentityCredential].
func NewFederatedIdentityCredential(name string, args FederatedIdentityCredentialArgs) *FederatedIdentityCredential {
	return &FederatedIdentityCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FederatedIdentityCredential)(nil)

// FederatedIdentityCredential represents the Terraform resource azurerm_federated_identity_credential.
type FederatedIdentityCredential struct {
	Name      string
	Args      FederatedIdentityCredentialArgs
	state     *federatedIdentityCredentialState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) Type() string {
	return "azurerm_federated_identity_credential"
}

// LocalName returns the local name for [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) LocalName() string {
	return fic.Name
}

// Configuration returns the configuration (args) for [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) Configuration() interface{} {
	return fic.Args
}

// DependOn is used for other resources to depend on [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(fic)
}

// Dependencies returns the list of resources [FederatedIdentityCredential] depends_on.
func (fic *FederatedIdentityCredential) Dependencies() terra.Dependencies {
	return fic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) LifecycleManagement() *terra.Lifecycle {
	return fic.Lifecycle
}

// Attributes returns the attributes for [FederatedIdentityCredential].
func (fic *FederatedIdentityCredential) Attributes() federatedIdentityCredentialAttributes {
	return federatedIdentityCredentialAttributes{ref: terra.ReferenceResource(fic)}
}

// ImportState imports the given attribute values into [FederatedIdentityCredential]'s state.
func (fic *FederatedIdentityCredential) ImportState(av io.Reader) error {
	fic.state = &federatedIdentityCredentialState{}
	if err := json.NewDecoder(av).Decode(fic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fic.Type(), fic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FederatedIdentityCredential] has state.
func (fic *FederatedIdentityCredential) State() (*federatedIdentityCredentialState, bool) {
	return fic.state, fic.state != nil
}

// StateMust returns the state for [FederatedIdentityCredential]. Panics if the state is nil.
func (fic *FederatedIdentityCredential) StateMust() *federatedIdentityCredentialState {
	if fic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fic.Type(), fic.LocalName()))
	}
	return fic.state
}

// FederatedIdentityCredentialArgs contains the configurations for azurerm_federated_identity_credential.
type FederatedIdentityCredentialArgs struct {
	// Audience: list of string, required
	Audience terra.ListValue[terra.StringValue] `hcl:"audience,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *federatedidentitycredential.Timeouts `hcl:"timeouts,block"`
}
type federatedIdentityCredentialAttributes struct {
	ref terra.Reference
}

// Audience returns a reference to field audience of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) Audience() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fic.ref.Append("audience"))
}

// Id returns a reference to field id of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("id"))
}

// Issuer returns a reference to field issuer of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("issuer"))
}

// Name returns a reference to field name of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("name"))
}

// ParentId returns a reference to field parent_id of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("parent_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("resource_group_name"))
}

// Subject returns a reference to field subject of azurerm_federated_identity_credential.
func (fic federatedIdentityCredentialAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("subject"))
}

func (fic federatedIdentityCredentialAttributes) Timeouts() federatedidentitycredential.TimeoutsAttributes {
	return terra.ReferenceAsSingle[federatedidentitycredential.TimeoutsAttributes](fic.ref.Append("timeouts"))
}

type federatedIdentityCredentialState struct {
	Audience          []string                                   `json:"audience"`
	Id                string                                     `json:"id"`
	Issuer            string                                     `json:"issuer"`
	Name              string                                     `json:"name"`
	ParentId          string                                     `json:"parent_id"`
	ResourceGroupName string                                     `json:"resource_group_name"`
	Subject           string                                     `json:"subject"`
	Timeouts          *federatedidentitycredential.TimeoutsState `json:"timeouts"`
}
