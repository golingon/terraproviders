// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsightssmartdetectionrule "github.com/golingon/terraproviders/azurerm/3.58.0/applicationinsightssmartdetectionrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsightsSmartDetectionRule creates a new instance of [ApplicationInsightsSmartDetectionRule].
func NewApplicationInsightsSmartDetectionRule(name string, args ApplicationInsightsSmartDetectionRuleArgs) *ApplicationInsightsSmartDetectionRule {
	return &ApplicationInsightsSmartDetectionRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsightsSmartDetectionRule)(nil)

// ApplicationInsightsSmartDetectionRule represents the Terraform resource azurerm_application_insights_smart_detection_rule.
type ApplicationInsightsSmartDetectionRule struct {
	Name      string
	Args      ApplicationInsightsSmartDetectionRuleArgs
	state     *applicationInsightsSmartDetectionRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) Type() string {
	return "azurerm_application_insights_smart_detection_rule"
}

// LocalName returns the local name for [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) LocalName() string {
	return aisdr.Name
}

// Configuration returns the configuration (args) for [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) Configuration() interface{} {
	return aisdr.Args
}

// DependOn is used for other resources to depend on [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) DependOn() terra.Reference {
	return terra.ReferenceResource(aisdr)
}

// Dependencies returns the list of resources [ApplicationInsightsSmartDetectionRule] depends_on.
func (aisdr *ApplicationInsightsSmartDetectionRule) Dependencies() terra.Dependencies {
	return aisdr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) LifecycleManagement() *terra.Lifecycle {
	return aisdr.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsightsSmartDetectionRule].
func (aisdr *ApplicationInsightsSmartDetectionRule) Attributes() applicationInsightsSmartDetectionRuleAttributes {
	return applicationInsightsSmartDetectionRuleAttributes{ref: terra.ReferenceResource(aisdr)}
}

// ImportState imports the given attribute values into [ApplicationInsightsSmartDetectionRule]'s state.
func (aisdr *ApplicationInsightsSmartDetectionRule) ImportState(av io.Reader) error {
	aisdr.state = &applicationInsightsSmartDetectionRuleState{}
	if err := json.NewDecoder(av).Decode(aisdr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aisdr.Type(), aisdr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsightsSmartDetectionRule] has state.
func (aisdr *ApplicationInsightsSmartDetectionRule) State() (*applicationInsightsSmartDetectionRuleState, bool) {
	return aisdr.state, aisdr.state != nil
}

// StateMust returns the state for [ApplicationInsightsSmartDetectionRule]. Panics if the state is nil.
func (aisdr *ApplicationInsightsSmartDetectionRule) StateMust() *applicationInsightsSmartDetectionRuleState {
	if aisdr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aisdr.Type(), aisdr.LocalName()))
	}
	return aisdr.state
}

// ApplicationInsightsSmartDetectionRuleArgs contains the configurations for azurerm_application_insights_smart_detection_rule.
type ApplicationInsightsSmartDetectionRuleArgs struct {
	// AdditionalEmailRecipients: set of string, optional
	AdditionalEmailRecipients terra.SetValue[terra.StringValue] `hcl:"additional_email_recipients,attr"`
	// ApplicationInsightsId: string, required
	ApplicationInsightsId terra.StringValue `hcl:"application_insights_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SendEmailsToSubscriptionOwners: bool, optional
	SendEmailsToSubscriptionOwners terra.BoolValue `hcl:"send_emails_to_subscription_owners,attr"`
	// Timeouts: optional
	Timeouts *applicationinsightssmartdetectionrule.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsSmartDetectionRuleAttributes struct {
	ref terra.Reference
}

// AdditionalEmailRecipients returns a reference to field additional_email_recipients of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) AdditionalEmailRecipients() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aisdr.ref.Append("additional_email_recipients"))
}

// ApplicationInsightsId returns a reference to field application_insights_id of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceAsString(aisdr.ref.Append("application_insights_id"))
}

// Enabled returns a reference to field enabled of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aisdr.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aisdr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aisdr.ref.Append("name"))
}

// SendEmailsToSubscriptionOwners returns a reference to field send_emails_to_subscription_owners of azurerm_application_insights_smart_detection_rule.
func (aisdr applicationInsightsSmartDetectionRuleAttributes) SendEmailsToSubscriptionOwners() terra.BoolValue {
	return terra.ReferenceAsBool(aisdr.ref.Append("send_emails_to_subscription_owners"))
}

func (aisdr applicationInsightsSmartDetectionRuleAttributes) Timeouts() applicationinsightssmartdetectionrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsightssmartdetectionrule.TimeoutsAttributes](aisdr.ref.Append("timeouts"))
}

type applicationInsightsSmartDetectionRuleState struct {
	AdditionalEmailRecipients      []string                                             `json:"additional_email_recipients"`
	ApplicationInsightsId          string                                               `json:"application_insights_id"`
	Enabled                        bool                                                 `json:"enabled"`
	Id                             string                                               `json:"id"`
	Name                           string                                               `json:"name"`
	SendEmailsToSubscriptionOwners bool                                                 `json:"send_emails_to_subscription_owners"`
	Timeouts                       *applicationinsightssmartdetectionrule.TimeoutsState `json:"timeouts"`
}
