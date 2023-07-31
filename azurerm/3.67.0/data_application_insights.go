// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapplicationinsights "github.com/golingon/terraproviders/azurerm/3.67.0/dataapplicationinsights"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApplicationInsights creates a new instance of [DataApplicationInsights].
func NewDataApplicationInsights(name string, args DataApplicationInsightsArgs) *DataApplicationInsights {
	return &DataApplicationInsights{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApplicationInsights)(nil)

// DataApplicationInsights represents the Terraform data resource azurerm_application_insights.
type DataApplicationInsights struct {
	Name string
	Args DataApplicationInsightsArgs
}

// DataSource returns the Terraform object type for [DataApplicationInsights].
func (ai *DataApplicationInsights) DataSource() string {
	return "azurerm_application_insights"
}

// LocalName returns the local name for [DataApplicationInsights].
func (ai *DataApplicationInsights) LocalName() string {
	return ai.Name
}

// Configuration returns the configuration (args) for [DataApplicationInsights].
func (ai *DataApplicationInsights) Configuration() interface{} {
	return ai.Args
}

// Attributes returns the attributes for [DataApplicationInsights].
func (ai *DataApplicationInsights) Attributes() dataApplicationInsightsAttributes {
	return dataApplicationInsightsAttributes{ref: terra.ReferenceDataResource(ai)}
}

// DataApplicationInsightsArgs contains the configurations for azurerm_application_insights.
type DataApplicationInsightsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataapplicationinsights.Timeouts `hcl:"timeouts,block"`
}
type dataApplicationInsightsAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("app_id"))
}

// ApplicationType returns a reference to field application_type of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) ApplicationType() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("application_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("connection_string"))
}

// Id returns a reference to field id of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("id"))
}

// InstrumentationKey returns a reference to field instrumentation_key of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) InstrumentationKey() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("instrumentation_key"))
}

// Location returns a reference to field location of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("resource_group_name"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ai.ref.Append("retention_in_days"))
}

// Tags returns a reference to field tags of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ai.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_application_insights.
func (ai dataApplicationInsightsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("workspace_id"))
}

func (ai dataApplicationInsightsAttributes) Timeouts() dataapplicationinsights.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapplicationinsights.TimeoutsAttributes](ai.ref.Append("timeouts"))
}
