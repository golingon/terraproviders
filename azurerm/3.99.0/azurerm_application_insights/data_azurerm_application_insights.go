// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_application_insights

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_application_insights.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aai *DataSource) DataSource() string {
	return "azurerm_application_insights"
}

// LocalName returns the local name for [DataSource].
func (aai *DataSource) LocalName() string {
	return aai.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aai *DataSource) Configuration() interface{} {
	return aai.Args
}

// Attributes returns the attributes for [DataSource].
func (aai *DataSource) Attributes() dataAzurermApplicationInsightsAttributes {
	return dataAzurermApplicationInsightsAttributes{ref: terra.ReferenceDataSource(aai)}
}

// DataArgs contains the configurations for azurerm_application_insights.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermApplicationInsightsAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("app_id"))
}

// ApplicationType returns a reference to field application_type of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) ApplicationType() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("application_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("connection_string"))
}

// Id returns a reference to field id of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("id"))
}

// InstrumentationKey returns a reference to field instrumentation_key of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) InstrumentationKey() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("instrumentation_key"))
}

// Location returns a reference to field location of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("resource_group_name"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(aai.ref.Append("retention_in_days"))
}

// Tags returns a reference to field tags of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aai.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_application_insights.
func (aai dataAzurermApplicationInsightsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("workspace_id"))
}

func (aai dataAzurermApplicationInsightsAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aai.ref.Append("timeouts"))
}
