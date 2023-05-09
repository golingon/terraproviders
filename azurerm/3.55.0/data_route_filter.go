// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataroutefilter "github.com/golingon/terraproviders/azurerm/3.55.0/dataroutefilter"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRouteFilter creates a new instance of [DataRouteFilter].
func NewDataRouteFilter(name string, args DataRouteFilterArgs) *DataRouteFilter {
	return &DataRouteFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRouteFilter)(nil)

// DataRouteFilter represents the Terraform data resource azurerm_route_filter.
type DataRouteFilter struct {
	Name string
	Args DataRouteFilterArgs
}

// DataSource returns the Terraform object type for [DataRouteFilter].
func (rf *DataRouteFilter) DataSource() string {
	return "azurerm_route_filter"
}

// LocalName returns the local name for [DataRouteFilter].
func (rf *DataRouteFilter) LocalName() string {
	return rf.Name
}

// Configuration returns the configuration (args) for [DataRouteFilter].
func (rf *DataRouteFilter) Configuration() interface{} {
	return rf.Args
}

// Attributes returns the attributes for [DataRouteFilter].
func (rf *DataRouteFilter) Attributes() dataRouteFilterAttributes {
	return dataRouteFilterAttributes{ref: terra.ReferenceDataResource(rf)}
}

// DataRouteFilterArgs contains the configurations for azurerm_route_filter.
type DataRouteFilterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Rule: min=0
	Rule []dataroutefilter.Rule `hcl:"rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataroutefilter.Timeouts `hcl:"timeouts,block"`
}
type dataRouteFilterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_route_filter.
func (rf dataRouteFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_route_filter.
func (rf dataRouteFilterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_route_filter.
func (rf dataRouteFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_route_filter.
func (rf dataRouteFilterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rf.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_route_filter.
func (rf dataRouteFilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rf.ref.Append("tags"))
}

func (rf dataRouteFilterAttributes) Rule() terra.ListValue[dataroutefilter.RuleAttributes] {
	return terra.ReferenceAsList[dataroutefilter.RuleAttributes](rf.ref.Append("rule"))
}

func (rf dataRouteFilterAttributes) Timeouts() dataroutefilter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataroutefilter.TimeoutsAttributes](rf.ref.Append("timeouts"))
}
