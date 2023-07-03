// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	marketplaceagreement "github.com/golingon/terraproviders/azurerm/3.63.0/marketplaceagreement"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMarketplaceAgreement creates a new instance of [MarketplaceAgreement].
func NewMarketplaceAgreement(name string, args MarketplaceAgreementArgs) *MarketplaceAgreement {
	return &MarketplaceAgreement{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MarketplaceAgreement)(nil)

// MarketplaceAgreement represents the Terraform resource azurerm_marketplace_agreement.
type MarketplaceAgreement struct {
	Name      string
	Args      MarketplaceAgreementArgs
	state     *marketplaceAgreementState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MarketplaceAgreement].
func (ma *MarketplaceAgreement) Type() string {
	return "azurerm_marketplace_agreement"
}

// LocalName returns the local name for [MarketplaceAgreement].
func (ma *MarketplaceAgreement) LocalName() string {
	return ma.Name
}

// Configuration returns the configuration (args) for [MarketplaceAgreement].
func (ma *MarketplaceAgreement) Configuration() interface{} {
	return ma.Args
}

// DependOn is used for other resources to depend on [MarketplaceAgreement].
func (ma *MarketplaceAgreement) DependOn() terra.Reference {
	return terra.ReferenceResource(ma)
}

// Dependencies returns the list of resources [MarketplaceAgreement] depends_on.
func (ma *MarketplaceAgreement) Dependencies() terra.Dependencies {
	return ma.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MarketplaceAgreement].
func (ma *MarketplaceAgreement) LifecycleManagement() *terra.Lifecycle {
	return ma.Lifecycle
}

// Attributes returns the attributes for [MarketplaceAgreement].
func (ma *MarketplaceAgreement) Attributes() marketplaceAgreementAttributes {
	return marketplaceAgreementAttributes{ref: terra.ReferenceResource(ma)}
}

// ImportState imports the given attribute values into [MarketplaceAgreement]'s state.
func (ma *MarketplaceAgreement) ImportState(av io.Reader) error {
	ma.state = &marketplaceAgreementState{}
	if err := json.NewDecoder(av).Decode(ma.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ma.Type(), ma.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MarketplaceAgreement] has state.
func (ma *MarketplaceAgreement) State() (*marketplaceAgreementState, bool) {
	return ma.state, ma.state != nil
}

// StateMust returns the state for [MarketplaceAgreement]. Panics if the state is nil.
func (ma *MarketplaceAgreement) StateMust() *marketplaceAgreementState {
	if ma.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ma.Type(), ma.LocalName()))
	}
	return ma.state
}

// MarketplaceAgreementArgs contains the configurations for azurerm_marketplace_agreement.
type MarketplaceAgreementArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Offer: string, required
	Offer terra.StringValue `hcl:"offer,attr" validate:"required"`
	// Plan: string, required
	Plan terra.StringValue `hcl:"plan,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *marketplaceagreement.Timeouts `hcl:"timeouts,block"`
}
type marketplaceAgreementAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("id"))
}

// LicenseTextLink returns a reference to field license_text_link of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) LicenseTextLink() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("license_text_link"))
}

// Offer returns a reference to field offer of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) Offer() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("offer"))
}

// Plan returns a reference to field plan of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) Plan() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("plan"))
}

// PrivacyPolicyLink returns a reference to field privacy_policy_link of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) PrivacyPolicyLink() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("privacy_policy_link"))
}

// Publisher returns a reference to field publisher of azurerm_marketplace_agreement.
func (ma marketplaceAgreementAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("publisher"))
}

func (ma marketplaceAgreementAttributes) Timeouts() marketplaceagreement.TimeoutsAttributes {
	return terra.ReferenceAsSingle[marketplaceagreement.TimeoutsAttributes](ma.ref.Append("timeouts"))
}

type marketplaceAgreementState struct {
	Id                string                              `json:"id"`
	LicenseTextLink   string                              `json:"license_text_link"`
	Offer             string                              `json:"offer"`
	Plan              string                              `json:"plan"`
	PrivacyPolicyLink string                              `json:"privacy_policy_link"`
	Publisher         string                              `json:"publisher"`
	Timeouts          *marketplaceagreement.TimeoutsState `json:"timeouts"`
}
