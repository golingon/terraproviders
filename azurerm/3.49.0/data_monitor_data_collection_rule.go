// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamonitordatacollectionrule "github.com/golingon/terraproviders/azurerm/3.49.0/datamonitordatacollectionrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitorDataCollectionRule creates a new instance of [DataMonitorDataCollectionRule].
func NewDataMonitorDataCollectionRule(name string, args DataMonitorDataCollectionRuleArgs) *DataMonitorDataCollectionRule {
	return &DataMonitorDataCollectionRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorDataCollectionRule)(nil)

// DataMonitorDataCollectionRule represents the Terraform data resource azurerm_monitor_data_collection_rule.
type DataMonitorDataCollectionRule struct {
	Name string
	Args DataMonitorDataCollectionRuleArgs
}

// DataSource returns the Terraform object type for [DataMonitorDataCollectionRule].
func (mdcr *DataMonitorDataCollectionRule) DataSource() string {
	return "azurerm_monitor_data_collection_rule"
}

// LocalName returns the local name for [DataMonitorDataCollectionRule].
func (mdcr *DataMonitorDataCollectionRule) LocalName() string {
	return mdcr.Name
}

// Configuration returns the configuration (args) for [DataMonitorDataCollectionRule].
func (mdcr *DataMonitorDataCollectionRule) Configuration() interface{} {
	return mdcr.Args
}

// Attributes returns the attributes for [DataMonitorDataCollectionRule].
func (mdcr *DataMonitorDataCollectionRule) Attributes() dataMonitorDataCollectionRuleAttributes {
	return dataMonitorDataCollectionRuleAttributes{ref: terra.ReferenceDataResource(mdcr)}
}

// DataMonitorDataCollectionRuleArgs contains the configurations for azurerm_monitor_data_collection_rule.
type DataMonitorDataCollectionRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// DataFlow: min=0
	DataFlow []datamonitordatacollectionrule.DataFlow `hcl:"data_flow,block" validate:"min=0"`
	// DataSources: min=0
	DataSources []datamonitordatacollectionrule.DataSources `hcl:"data_sources,block" validate:"min=0"`
	// Destinations: min=0
	Destinations []datamonitordatacollectionrule.Destinations `hcl:"destinations,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamonitordatacollectionrule.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorDataCollectionRuleAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mdcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_data_collection_rule.
func (mdcr dataMonitorDataCollectionRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mdcr.ref.Append("tags"))
}

func (mdcr dataMonitorDataCollectionRuleAttributes) DataFlow() terra.ListValue[datamonitordatacollectionrule.DataFlowAttributes] {
	return terra.ReferenceAsList[datamonitordatacollectionrule.DataFlowAttributes](mdcr.ref.Append("data_flow"))
}

func (mdcr dataMonitorDataCollectionRuleAttributes) DataSources() terra.ListValue[datamonitordatacollectionrule.DataSourcesAttributes] {
	return terra.ReferenceAsList[datamonitordatacollectionrule.DataSourcesAttributes](mdcr.ref.Append("data_sources"))
}

func (mdcr dataMonitorDataCollectionRuleAttributes) Destinations() terra.ListValue[datamonitordatacollectionrule.DestinationsAttributes] {
	return terra.ReferenceAsList[datamonitordatacollectionrule.DestinationsAttributes](mdcr.ref.Append("destinations"))
}

func (mdcr dataMonitorDataCollectionRuleAttributes) Timeouts() datamonitordatacollectionrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitordatacollectionrule.TimeoutsAttributes](mdcr.ref.Append("timeouts"))
}
