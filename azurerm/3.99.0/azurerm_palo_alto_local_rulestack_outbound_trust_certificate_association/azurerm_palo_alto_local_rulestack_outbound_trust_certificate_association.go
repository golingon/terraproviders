// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association

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

// Resource represents the Terraform resource azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apalrotca *Resource) Type() string {
	return "azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association"
}

// LocalName returns the local name for [Resource].
func (apalrotca *Resource) LocalName() string {
	return apalrotca.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apalrotca *Resource) Configuration() interface{} {
	return apalrotca.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apalrotca *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apalrotca)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apalrotca *Resource) Dependencies() terra.Dependencies {
	return apalrotca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apalrotca *Resource) LifecycleManagement() *terra.Lifecycle {
	return apalrotca.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apalrotca *Resource) Attributes() azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes {
	return azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes{ref: terra.ReferenceResource(apalrotca)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apalrotca *Resource) ImportState(state io.Reader) error {
	apalrotca.state = &azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationState{}
	if err := json.NewDecoder(state).Decode(apalrotca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apalrotca.Type(), apalrotca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apalrotca *Resource) State() (*azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationState, bool) {
	return apalrotca.state, apalrotca.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apalrotca *Resource) StateMust() *azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationState {
	if apalrotca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apalrotca.Type(), apalrotca.LocalName()))
	}
	return apalrotca.state
}

// Args contains the configurations for azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
type Args struct {
	// CertificateId: string, required
	CertificateId terra.StringValue `hcl:"certificate_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes struct {
	ref terra.Reference
}

// CertificateId returns a reference to field certificate_id of azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
func (apalrotca azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(apalrotca.ref.Append("certificate_id"))
}

// Id returns a reference to field id of azurerm_palo_alto_local_rulestack_outbound_trust_certificate_association.
func (apalrotca azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apalrotca.ref.Append("id"))
}

func (apalrotca azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apalrotca.ref.Append("timeouts"))
}

type azurermPaloAltoLocalRulestackOutboundTrustCertificateAssociationState struct {
	CertificateId string         `json:"certificate_id"`
	Id            string         `json:"id"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
