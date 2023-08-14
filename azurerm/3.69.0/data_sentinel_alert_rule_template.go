// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasentinelalertruletemplate "github.com/golingon/terraproviders/azurerm/3.69.0/datasentinelalertruletemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSentinelAlertRuleTemplate creates a new instance of [DataSentinelAlertRuleTemplate].
func NewDataSentinelAlertRuleTemplate(name string, args DataSentinelAlertRuleTemplateArgs) *DataSentinelAlertRuleTemplate {
	return &DataSentinelAlertRuleTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSentinelAlertRuleTemplate)(nil)

// DataSentinelAlertRuleTemplate represents the Terraform data resource azurerm_sentinel_alert_rule_template.
type DataSentinelAlertRuleTemplate struct {
	Name string
	Args DataSentinelAlertRuleTemplateArgs
}

// DataSource returns the Terraform object type for [DataSentinelAlertRuleTemplate].
func (sart *DataSentinelAlertRuleTemplate) DataSource() string {
	return "azurerm_sentinel_alert_rule_template"
}

// LocalName returns the local name for [DataSentinelAlertRuleTemplate].
func (sart *DataSentinelAlertRuleTemplate) LocalName() string {
	return sart.Name
}

// Configuration returns the configuration (args) for [DataSentinelAlertRuleTemplate].
func (sart *DataSentinelAlertRuleTemplate) Configuration() interface{} {
	return sart.Args
}

// Attributes returns the attributes for [DataSentinelAlertRuleTemplate].
func (sart *DataSentinelAlertRuleTemplate) Attributes() dataSentinelAlertRuleTemplateAttributes {
	return dataSentinelAlertRuleTemplateAttributes{ref: terra.ReferenceDataResource(sart)}
}

// DataSentinelAlertRuleTemplateArgs contains the configurations for azurerm_sentinel_alert_rule_template.
type DataSentinelAlertRuleTemplateArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NrtTemplate: min=0
	NrtTemplate []datasentinelalertruletemplate.NrtTemplate `hcl:"nrt_template,block" validate:"min=0"`
	// ScheduledTemplate: min=0
	ScheduledTemplate []datasentinelalertruletemplate.ScheduledTemplate `hcl:"scheduled_template,block" validate:"min=0"`
	// SecurityIncidentTemplate: min=0
	SecurityIncidentTemplate []datasentinelalertruletemplate.SecurityIncidentTemplate `hcl:"security_incident_template,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasentinelalertruletemplate.Timeouts `hcl:"timeouts,block"`
}
type dataSentinelAlertRuleTemplateAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_template.
func (sart dataSentinelAlertRuleTemplateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sart.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_template.
func (sart dataSentinelAlertRuleTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sart.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_template.
func (sart dataSentinelAlertRuleTemplateAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sart.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_template.
func (sart dataSentinelAlertRuleTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sart.ref.Append("name"))
}

func (sart dataSentinelAlertRuleTemplateAttributes) NrtTemplate() terra.ListValue[datasentinelalertruletemplate.NrtTemplateAttributes] {
	return terra.ReferenceAsList[datasentinelalertruletemplate.NrtTemplateAttributes](sart.ref.Append("nrt_template"))
}

func (sart dataSentinelAlertRuleTemplateAttributes) ScheduledTemplate() terra.ListValue[datasentinelalertruletemplate.ScheduledTemplateAttributes] {
	return terra.ReferenceAsList[datasentinelalertruletemplate.ScheduledTemplateAttributes](sart.ref.Append("scheduled_template"))
}

func (sart dataSentinelAlertRuleTemplateAttributes) SecurityIncidentTemplate() terra.ListValue[datasentinelalertruletemplate.SecurityIncidentTemplateAttributes] {
	return terra.ReferenceAsList[datasentinelalertruletemplate.SecurityIncidentTemplateAttributes](sart.ref.Append("security_incident_template"))
}

func (sart dataSentinelAlertRuleTemplateAttributes) Timeouts() datasentinelalertruletemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasentinelalertruletemplate.TimeoutsAttributes](sart.ref.Append("timeouts"))
}
