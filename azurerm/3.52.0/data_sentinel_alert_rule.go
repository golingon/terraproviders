// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasentinelalertrule "github.com/golingon/terraproviders/azurerm/3.52.0/datasentinelalertrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSentinelAlertRule creates a new instance of [DataSentinelAlertRule].
func NewDataSentinelAlertRule(name string, args DataSentinelAlertRuleArgs) *DataSentinelAlertRule {
	return &DataSentinelAlertRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSentinelAlertRule)(nil)

// DataSentinelAlertRule represents the Terraform data resource azurerm_sentinel_alert_rule.
type DataSentinelAlertRule struct {
	Name string
	Args DataSentinelAlertRuleArgs
}

// DataSource returns the Terraform object type for [DataSentinelAlertRule].
func (sar *DataSentinelAlertRule) DataSource() string {
	return "azurerm_sentinel_alert_rule"
}

// LocalName returns the local name for [DataSentinelAlertRule].
func (sar *DataSentinelAlertRule) LocalName() string {
	return sar.Name
}

// Configuration returns the configuration (args) for [DataSentinelAlertRule].
func (sar *DataSentinelAlertRule) Configuration() interface{} {
	return sar.Args
}

// Attributes returns the attributes for [DataSentinelAlertRule].
func (sar *DataSentinelAlertRule) Attributes() dataSentinelAlertRuleAttributes {
	return dataSentinelAlertRuleAttributes{ref: terra.ReferenceDataResource(sar)}
}

// DataSentinelAlertRuleArgs contains the configurations for azurerm_sentinel_alert_rule.
type DataSentinelAlertRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasentinelalertrule.Timeouts `hcl:"timeouts,block"`
}
type dataSentinelAlertRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule.
func (sar dataSentinelAlertRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule.
func (sar dataSentinelAlertRuleAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule.
func (sar dataSentinelAlertRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sar.ref.Append("name"))
}

func (sar dataSentinelAlertRuleAttributes) Timeouts() datasentinelalertrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasentinelalertrule.TimeoutsAttributes](sar.ref.Append("timeouts"))
}
