// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	paloaltolocalrulestackoutboundtrustcertificateassociation "github.com/golingon/terraproviders/azurerm/3.69.0/paloaltolocalrulestackoutboundtrustcertificateassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPaloAltoLocalRulestackOutboundTrustCertificateAssociation creates a new instance of [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func NewPaloAltoLocalRulestackOutboundTrustCertificateAssociation(name string, args PaloAltoLocalRulestackOutboundTrustCertificateAssociationArgs) *PaloAltoLocalRulestackOutboundTrustCertificateAssociation {
	return &PaloAltoLocalRulestackOutboundTrustCertificateAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PaloAltoLocalRulestackOutboundTrustCertificateAssociation)(nil)

// PaloAltoLocalRulestackOutboundTrustCertificateAssociation represents the Terraform resource azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
type PaloAltoLocalRulestackOutboundTrustCertificateAssociation struct {
	Name      string
	Args      PaloAltoLocalRulestackOutboundTrustCertificateAssociationArgs
	state     *paloAltoLocalRulestackOutboundTrustCertificateAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) Type() string {
	return "azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association"
}

// LocalName returns the local name for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) LocalName() string {
	return palrotca.Name
}

// Configuration returns the configuration (args) for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) Configuration() interface{} {
	return palrotca.Args
}

// DependOn is used for other resources to depend on [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(palrotca)
}

// Dependencies returns the list of resources [PaloAltoLocalRulestackOutboundTrustCertificateAssociation] depends_on.
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) Dependencies() terra.Dependencies {
	return palrotca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) LifecycleManagement() *terra.Lifecycle {
	return palrotca.Lifecycle
}

// Attributes returns the attributes for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation].
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) Attributes() paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes {
	return paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes{ref: terra.ReferenceResource(palrotca)}
}

// ImportState imports the given attribute values into [PaloAltoLocalRulestackOutboundTrustCertificateAssociation]'s state.
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) ImportState(av io.Reader) error {
	palrotca.state = &paloAltoLocalRulestackOutboundTrustCertificateAssociationState{}
	if err := json.NewDecoder(av).Decode(palrotca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", palrotca.Type(), palrotca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PaloAltoLocalRulestackOutboundTrustCertificateAssociation] has state.
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) State() (*paloAltoLocalRulestackOutboundTrustCertificateAssociationState, bool) {
	return palrotca.state, palrotca.state != nil
}

// StateMust returns the state for [PaloAltoLocalRulestackOutboundTrustCertificateAssociation]. Panics if the state is nil.
func (palrotca *PaloAltoLocalRulestackOutboundTrustCertificateAssociation) StateMust() *paloAltoLocalRulestackOutboundTrustCertificateAssociationState {
	if palrotca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", palrotca.Type(), palrotca.LocalName()))
	}
	return palrotca.state
}

// PaloAltoLocalRulestackOutboundTrustCertificateAssociationArgs contains the configurations for azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
type PaloAltoLocalRulestackOutboundTrustCertificateAssociationArgs struct {
	// CertificateId: string, required
	CertificateId terra.StringValue `hcl:"certificate_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *paloaltolocalrulestackoutboundtrustcertificateassociation.Timeouts `hcl:"timeouts,block"`
}
type paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes struct {
	ref terra.Reference
}

// CertificateId returns a reference to field certificate_id of azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
func (palrotca paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(palrotca.ref.Append("certificate_id"))
}

// Id returns a reference to field id of azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
func (palrotca paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(palrotca.ref.Append("id"))
}

func (palrotca paloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) Timeouts() paloaltolocalrulestackoutboundtrustcertificateassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[paloaltolocalrulestackoutboundtrustcertificateassociation.TimeoutsAttributes](palrotca.ref.Append("timeouts"))
}

type paloAltoLocalRulestackOutboundTrustCertificateAssociationState struct {
	CertificateId string                                                                   `json:"certificate_id"`
	Id            string                                                                   `json:"id"`
	Timeouts      *paloaltolocalrulestackoutboundtrustcertificateassociation.TimeoutsState `json:"timeouts"`
}
