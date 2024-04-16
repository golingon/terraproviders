// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_logic_app_integration_account_agreement

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

// Resource represents the Terraform resource azurerm_logic_app_integration_account_agreement.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermLogicAppIntegrationAccountAgreementState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alaiaa *Resource) Type() string {
	return "azurerm_logic_app_integration_account_agreement"
}

// LocalName returns the local name for [Resource].
func (alaiaa *Resource) LocalName() string {
	return alaiaa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alaiaa *Resource) Configuration() interface{} {
	return alaiaa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alaiaa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alaiaa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alaiaa *Resource) Dependencies() terra.Dependencies {
	return alaiaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alaiaa *Resource) LifecycleManagement() *terra.Lifecycle {
	return alaiaa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alaiaa *Resource) Attributes() azurermLogicAppIntegrationAccountAgreementAttributes {
	return azurermLogicAppIntegrationAccountAgreementAttributes{ref: terra.ReferenceResource(alaiaa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alaiaa *Resource) ImportState(state io.Reader) error {
	alaiaa.state = &azurermLogicAppIntegrationAccountAgreementState{}
	if err := json.NewDecoder(state).Decode(alaiaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alaiaa.Type(), alaiaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alaiaa *Resource) State() (*azurermLogicAppIntegrationAccountAgreementState, bool) {
	return alaiaa.state, alaiaa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alaiaa *Resource) StateMust() *azurermLogicAppIntegrationAccountAgreementState {
	if alaiaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alaiaa.Type(), alaiaa.LocalName()))
	}
	return alaiaa.state
}

// Args contains the configurations for azurerm_logic_app_integration_account_agreement.
type Args struct {
	// AgreementType: string, required
	AgreementType terra.StringValue `hcl:"agreement_type,attr" validate:"required"`
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// GuestPartnerName: string, required
	GuestPartnerName terra.StringValue `hcl:"guest_partner_name,attr" validate:"required"`
	// HostPartnerName: string, required
	HostPartnerName terra.StringValue `hcl:"host_partner_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationAccountName: string, required
	IntegrationAccountName terra.StringValue `hcl:"integration_account_name,attr" validate:"required"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// GuestIdentity: required
	GuestIdentity *GuestIdentity `hcl:"guest_identity,block" validate:"required"`
	// HostIdentity: required
	HostIdentity *HostIdentity `hcl:"host_identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermLogicAppIntegrationAccountAgreementAttributes struct {
	ref terra.Reference
}

// AgreementType returns a reference to field agreement_type of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) AgreementType() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("agreement_type"))
}

// Content returns a reference to field content of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("content"))
}

// GuestPartnerName returns a reference to field guest_partner_name of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) GuestPartnerName() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("guest_partner_name"))
}

// HostPartnerName returns a reference to field host_partner_name of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) HostPartnerName() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("host_partner_name"))
}

// Id returns a reference to field id of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("id"))
}

// IntegrationAccountName returns a reference to field integration_account_name of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("integration_account_name"))
}

// Metadata returns a reference to field metadata of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alaiaa.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account_agreement.
func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(alaiaa.ref.Append("resource_group_name"))
}

func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) GuestIdentity() terra.ListValue[GuestIdentityAttributes] {
	return terra.ReferenceAsList[GuestIdentityAttributes](alaiaa.ref.Append("guest_identity"))
}

func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) HostIdentity() terra.ListValue[HostIdentityAttributes] {
	return terra.ReferenceAsList[HostIdentityAttributes](alaiaa.ref.Append("host_identity"))
}

func (alaiaa azurermLogicAppIntegrationAccountAgreementAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alaiaa.ref.Append("timeouts"))
}

type azurermLogicAppIntegrationAccountAgreementState struct {
	AgreementType          string               `json:"agreement_type"`
	Content                string               `json:"content"`
	GuestPartnerName       string               `json:"guest_partner_name"`
	HostPartnerName        string               `json:"host_partner_name"`
	Id                     string               `json:"id"`
	IntegrationAccountName string               `json:"integration_account_name"`
	Metadata               map[string]string    `json:"metadata"`
	Name                   string               `json:"name"`
	ResourceGroupName      string               `json:"resource_group_name"`
	GuestIdentity          []GuestIdentityState `json:"guest_identity"`
	HostIdentity           []HostIdentityState  `json:"host_identity"`
	Timeouts               *TimeoutsState       `json:"timeouts"`
}
