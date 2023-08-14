// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentineldataconnectorazuresecuritycenter "github.com/golingon/terraproviders/azurerm/3.69.0/sentineldataconnectorazuresecuritycenter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelDataConnectorAzureSecurityCenter creates a new instance of [SentinelDataConnectorAzureSecurityCenter].
func NewSentinelDataConnectorAzureSecurityCenter(name string, args SentinelDataConnectorAzureSecurityCenterArgs) *SentinelDataConnectorAzureSecurityCenter {
	return &SentinelDataConnectorAzureSecurityCenter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelDataConnectorAzureSecurityCenter)(nil)

// SentinelDataConnectorAzureSecurityCenter represents the Terraform resource azurerm_sentinel_data_connector_azure_security_center.
type SentinelDataConnectorAzureSecurityCenter struct {
	Name      string
	Args      SentinelDataConnectorAzureSecurityCenterArgs
	state     *sentinelDataConnectorAzureSecurityCenterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) Type() string {
	return "azurerm_sentinel_data_connector_azure_security_center"
}

// LocalName returns the local name for [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) LocalName() string {
	return sdcasc.Name
}

// Configuration returns the configuration (args) for [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) Configuration() interface{} {
	return sdcasc.Args
}

// DependOn is used for other resources to depend on [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) DependOn() terra.Reference {
	return terra.ReferenceResource(sdcasc)
}

// Dependencies returns the list of resources [SentinelDataConnectorAzureSecurityCenter] depends_on.
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) Dependencies() terra.Dependencies {
	return sdcasc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) LifecycleManagement() *terra.Lifecycle {
	return sdcasc.Lifecycle
}

// Attributes returns the attributes for [SentinelDataConnectorAzureSecurityCenter].
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) Attributes() sentinelDataConnectorAzureSecurityCenterAttributes {
	return sentinelDataConnectorAzureSecurityCenterAttributes{ref: terra.ReferenceResource(sdcasc)}
}

// ImportState imports the given attribute values into [SentinelDataConnectorAzureSecurityCenter]'s state.
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) ImportState(av io.Reader) error {
	sdcasc.state = &sentinelDataConnectorAzureSecurityCenterState{}
	if err := json.NewDecoder(av).Decode(sdcasc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdcasc.Type(), sdcasc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelDataConnectorAzureSecurityCenter] has state.
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) State() (*sentinelDataConnectorAzureSecurityCenterState, bool) {
	return sdcasc.state, sdcasc.state != nil
}

// StateMust returns the state for [SentinelDataConnectorAzureSecurityCenter]. Panics if the state is nil.
func (sdcasc *SentinelDataConnectorAzureSecurityCenter) StateMust() *sentinelDataConnectorAzureSecurityCenterState {
	if sdcasc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdcasc.Type(), sdcasc.LocalName()))
	}
	return sdcasc.state
}

// SentinelDataConnectorAzureSecurityCenterArgs contains the configurations for azurerm_sentinel_data_connector_azure_security_center.
type SentinelDataConnectorAzureSecurityCenterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// Timeouts: optional
	Timeouts *sentineldataconnectorazuresecuritycenter.Timeouts `hcl:"timeouts,block"`
}
type sentinelDataConnectorAzureSecurityCenterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_data_connector_azure_security_center.
func (sdcasc sentinelDataConnectorAzureSecurityCenterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdcasc.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_data_connector_azure_security_center.
func (sdcasc sentinelDataConnectorAzureSecurityCenterAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sdcasc.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_data_connector_azure_security_center.
func (sdcasc sentinelDataConnectorAzureSecurityCenterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdcasc.ref.Append("name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_sentinel_data_connector_azure_security_center.
func (sdcasc sentinelDataConnectorAzureSecurityCenterAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(sdcasc.ref.Append("subscription_id"))
}

func (sdcasc sentinelDataConnectorAzureSecurityCenterAttributes) Timeouts() sentineldataconnectorazuresecuritycenter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentineldataconnectorazuresecuritycenter.TimeoutsAttributes](sdcasc.ref.Append("timeouts"))
}

type sentinelDataConnectorAzureSecurityCenterState struct {
	Id                      string                                                  `json:"id"`
	LogAnalyticsWorkspaceId string                                                  `json:"log_analytics_workspace_id"`
	Name                    string                                                  `json:"name"`
	SubscriptionId          string                                                  `json:"subscription_id"`
	Timeouts                *sentineldataconnectorazuresecuritycenter.TimeoutsState `json:"timeouts"`
}
