// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamssqlelasticpool "github.com/golingon/terraproviders/azurerm/3.58.0/datamssqlelasticpool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMssqlElasticpool creates a new instance of [DataMssqlElasticpool].
func NewDataMssqlElasticpool(name string, args DataMssqlElasticpoolArgs) *DataMssqlElasticpool {
	return &DataMssqlElasticpool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMssqlElasticpool)(nil)

// DataMssqlElasticpool represents the Terraform data resource azurerm_mssql_elasticpool.
type DataMssqlElasticpool struct {
	Name string
	Args DataMssqlElasticpoolArgs
}

// DataSource returns the Terraform object type for [DataMssqlElasticpool].
func (me *DataMssqlElasticpool) DataSource() string {
	return "azurerm_mssql_elasticpool"
}

// LocalName returns the local name for [DataMssqlElasticpool].
func (me *DataMssqlElasticpool) LocalName() string {
	return me.Name
}

// Configuration returns the configuration (args) for [DataMssqlElasticpool].
func (me *DataMssqlElasticpool) Configuration() interface{} {
	return me.Args
}

// Attributes returns the attributes for [DataMssqlElasticpool].
func (me *DataMssqlElasticpool) Attributes() dataMssqlElasticpoolAttributes {
	return dataMssqlElasticpoolAttributes{ref: terra.ReferenceDataResource(me)}
}

// DataMssqlElasticpoolArgs contains the configurations for azurerm_mssql_elasticpool.
type DataMssqlElasticpoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Sku: min=0
	Sku []datamssqlelasticpool.Sku `hcl:"sku,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamssqlelasticpool.Timeouts `hcl:"timeouts,block"`
}
type dataMssqlElasticpoolAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("location"))
}

// MaxSizeBytes returns a reference to field max_size_bytes of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) MaxSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("max_size_bytes"))
}

// MaxSizeGb returns a reference to field max_size_gb of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) MaxSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("max_size_gb"))
}

// Name returns a reference to field name of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("name"))
}

// PerDbMaxCapacity returns a reference to field per_db_max_capacity of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) PerDbMaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("per_db_max_capacity"))
}

// PerDbMinCapacity returns a reference to field per_db_min_capacity of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) PerDbMinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("per_db_min_capacity"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("server_name"))
}

// Tags returns a reference to field tags of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](me.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_mssql_elasticpool.
func (me dataMssqlElasticpoolAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(me.ref.Append("zone_redundant"))
}

func (me dataMssqlElasticpoolAttributes) Sku() terra.ListValue[datamssqlelasticpool.SkuAttributes] {
	return terra.ReferenceAsList[datamssqlelasticpool.SkuAttributes](me.ref.Append("sku"))
}

func (me dataMssqlElasticpoolAttributes) Timeouts() datamssqlelasticpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamssqlelasticpool.TimeoutsAttributes](me.ref.Append("timeouts"))
}
