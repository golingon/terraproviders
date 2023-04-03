// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappintegrationaccountcertificate "github.com/golingon/terraproviders/azurerm/3.49.0/logicappintegrationaccountcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppIntegrationAccountCertificate creates a new instance of [LogicAppIntegrationAccountCertificate].
func NewLogicAppIntegrationAccountCertificate(name string, args LogicAppIntegrationAccountCertificateArgs) *LogicAppIntegrationAccountCertificate {
	return &LogicAppIntegrationAccountCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppIntegrationAccountCertificate)(nil)

// LogicAppIntegrationAccountCertificate represents the Terraform resource azurerm_logic_app_integration_account_certificate.
type LogicAppIntegrationAccountCertificate struct {
	Name      string
	Args      LogicAppIntegrationAccountCertificateArgs
	state     *logicAppIntegrationAccountCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) Type() string {
	return "azurerm_logic_app_integration_account_certificate"
}

// LocalName returns the local name for [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) LocalName() string {
	return laiac.Name
}

// Configuration returns the configuration (args) for [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) Configuration() interface{} {
	return laiac.Args
}

// DependOn is used for other resources to depend on [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(laiac)
}

// Dependencies returns the list of resources [LogicAppIntegrationAccountCertificate] depends_on.
func (laiac *LogicAppIntegrationAccountCertificate) Dependencies() terra.Dependencies {
	return laiac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) LifecycleManagement() *terra.Lifecycle {
	return laiac.Lifecycle
}

// Attributes returns the attributes for [LogicAppIntegrationAccountCertificate].
func (laiac *LogicAppIntegrationAccountCertificate) Attributes() logicAppIntegrationAccountCertificateAttributes {
	return logicAppIntegrationAccountCertificateAttributes{ref: terra.ReferenceResource(laiac)}
}

// ImportState imports the given attribute values into [LogicAppIntegrationAccountCertificate]'s state.
func (laiac *LogicAppIntegrationAccountCertificate) ImportState(av io.Reader) error {
	laiac.state = &logicAppIntegrationAccountCertificateState{}
	if err := json.NewDecoder(av).Decode(laiac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laiac.Type(), laiac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppIntegrationAccountCertificate] has state.
func (laiac *LogicAppIntegrationAccountCertificate) State() (*logicAppIntegrationAccountCertificateState, bool) {
	return laiac.state, laiac.state != nil
}

// StateMust returns the state for [LogicAppIntegrationAccountCertificate]. Panics if the state is nil.
func (laiac *LogicAppIntegrationAccountCertificate) StateMust() *logicAppIntegrationAccountCertificateState {
	if laiac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laiac.Type(), laiac.LocalName()))
	}
	return laiac.state
}

// LogicAppIntegrationAccountCertificateArgs contains the configurations for azurerm_logic_app_integration_account_certificate.
type LogicAppIntegrationAccountCertificateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationAccountName: string, required
	IntegrationAccountName terra.StringValue `hcl:"integration_account_name,attr" validate:"required"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicCertificate: string, optional
	PublicCertificate terra.StringValue `hcl:"public_certificate,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// KeyVaultKey: optional
	KeyVaultKey *logicappintegrationaccountcertificate.KeyVaultKey `hcl:"key_vault_key,block"`
	// Timeouts: optional
	Timeouts *logicappintegrationaccountcertificate.Timeouts `hcl:"timeouts,block"`
}
type logicAppIntegrationAccountCertificateAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("id"))
}

// IntegrationAccountName returns a reference to field integration_account_name of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("integration_account_name"))
}

// Metadata returns a reference to field metadata of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("name"))
}

// PublicCertificate returns a reference to field public_certificate of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) PublicCertificate() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("public_certificate"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account_certificate.
func (laiac logicAppIntegrationAccountCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laiac.ref.Append("resource_group_name"))
}

func (laiac logicAppIntegrationAccountCertificateAttributes) KeyVaultKey() terra.ListValue[logicappintegrationaccountcertificate.KeyVaultKeyAttributes] {
	return terra.ReferenceAsList[logicappintegrationaccountcertificate.KeyVaultKeyAttributes](laiac.ref.Append("key_vault_key"))
}

func (laiac logicAppIntegrationAccountCertificateAttributes) Timeouts() logicappintegrationaccountcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappintegrationaccountcertificate.TimeoutsAttributes](laiac.ref.Append("timeouts"))
}

type logicAppIntegrationAccountCertificateState struct {
	Id                     string                                                   `json:"id"`
	IntegrationAccountName string                                                   `json:"integration_account_name"`
	Metadata               string                                                   `json:"metadata"`
	Name                   string                                                   `json:"name"`
	PublicCertificate      string                                                   `json:"public_certificate"`
	ResourceGroupName      string                                                   `json:"resource_group_name"`
	KeyVaultKey            []logicappintegrationaccountcertificate.KeyVaultKeyState `json:"key_vault_key"`
	Timeouts               *logicappintegrationaccountcertificate.TimeoutsState     `json:"timeouts"`
}
