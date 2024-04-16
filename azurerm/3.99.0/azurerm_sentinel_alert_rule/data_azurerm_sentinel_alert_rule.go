// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_alert_rule

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_sentinel_alert_rule.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asar *DataSource) DataSource() string {
	return "azurerm_sentinel_alert_rule"
}

// LocalName returns the local name for [DataSource].
func (asar *DataSource) LocalName() string {
	return asar.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asar *DataSource) Configuration() interface{} {
	return asar.Args
}

// Attributes returns the attributes for [DataSource].
func (asar *DataSource) Attributes() dataAzurermSentinelAlertRuleAttributes {
	return dataAzurermSentinelAlertRuleAttributes{ref: terra.ReferenceDataSource(asar)}
}

// DataArgs contains the configurations for azurerm_sentinel_alert_rule.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermSentinelAlertRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule.
func (asar dataAzurermSentinelAlertRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asar.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule.
func (asar dataAzurermSentinelAlertRuleAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(asar.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule.
func (asar dataAzurermSentinelAlertRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asar.ref.Append("name"))
}

func (asar dataAzurermSentinelAlertRuleAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](asar.ref.Append("timeouts"))
}
