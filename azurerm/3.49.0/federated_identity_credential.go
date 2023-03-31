// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	federatedidentitycredential "github.com/golingon/terraproviders/azurerm/3.49.0/federatedidentitycredential"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewFederatedIdentityCredential(name string, args FederatedIdentityCredentialArgs) *FederatedIdentityCredential {
	return &FederatedIdentityCredential{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FederatedIdentityCredential)(nil)

type FederatedIdentityCredential struct {
	Name  string
	Args  FederatedIdentityCredentialArgs
	state *federatedIdentityCredentialState
}

func (fic *FederatedIdentityCredential) Type() string {
	return "azurerm_federated_identity_credential"
}

func (fic *FederatedIdentityCredential) LocalName() string {
	return fic.Name
}

func (fic *FederatedIdentityCredential) Configuration() interface{} {
	return fic.Args
}

func (fic *FederatedIdentityCredential) Attributes() federatedIdentityCredentialAttributes {
	return federatedIdentityCredentialAttributes{ref: terra.ReferenceResource(fic)}
}

func (fic *FederatedIdentityCredential) ImportState(av io.Reader) error {
	fic.state = &federatedIdentityCredentialState{}
	if err := json.NewDecoder(av).Decode(fic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fic.Type(), fic.LocalName(), err)
	}
	return nil
}

func (fic *FederatedIdentityCredential) State() (*federatedIdentityCredentialState, bool) {
	return fic.state, fic.state != nil
}

func (fic *FederatedIdentityCredential) StateMust() *federatedIdentityCredentialState {
	if fic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fic.Type(), fic.LocalName()))
	}
	return fic.state
}

func (fic *FederatedIdentityCredential) DependOn() terra.Reference {
	return terra.ReferenceResource(fic)
}

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
	// DependsOn contains resources that FederatedIdentityCredential depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type federatedIdentityCredentialAttributes struct {
	ref terra.Reference
}

func (fic federatedIdentityCredentialAttributes) Audience() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](fic.ref.Append("audience"))
}

func (fic federatedIdentityCredentialAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("id"))
}

func (fic federatedIdentityCredentialAttributes) Issuer() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("issuer"))
}

func (fic federatedIdentityCredentialAttributes) Name() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("name"))
}

func (fic federatedIdentityCredentialAttributes) ParentId() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("parent_id"))
}

func (fic federatedIdentityCredentialAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("resource_group_name"))
}

func (fic federatedIdentityCredentialAttributes) Subject() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("subject"))
}

func (fic federatedIdentityCredentialAttributes) Timeouts() federatedidentitycredential.TimeoutsAttributes {
	return terra.ReferenceSingle[federatedidentitycredential.TimeoutsAttributes](fic.ref.Append("timeouts"))
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
