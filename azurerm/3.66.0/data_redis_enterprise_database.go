// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataredisenterprisedatabase "github.com/golingon/terraproviders/azurerm/3.66.0/dataredisenterprisedatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRedisEnterpriseDatabase creates a new instance of [DataRedisEnterpriseDatabase].
func NewDataRedisEnterpriseDatabase(name string, args DataRedisEnterpriseDatabaseArgs) *DataRedisEnterpriseDatabase {
	return &DataRedisEnterpriseDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRedisEnterpriseDatabase)(nil)

// DataRedisEnterpriseDatabase represents the Terraform data resource azurerm_redis_enterprise_database.
type DataRedisEnterpriseDatabase struct {
	Name string
	Args DataRedisEnterpriseDatabaseArgs
}

// DataSource returns the Terraform object type for [DataRedisEnterpriseDatabase].
func (red *DataRedisEnterpriseDatabase) DataSource() string {
	return "azurerm_redis_enterprise_database"
}

// LocalName returns the local name for [DataRedisEnterpriseDatabase].
func (red *DataRedisEnterpriseDatabase) LocalName() string {
	return red.Name
}

// Configuration returns the configuration (args) for [DataRedisEnterpriseDatabase].
func (red *DataRedisEnterpriseDatabase) Configuration() interface{} {
	return red.Args
}

// Attributes returns the attributes for [DataRedisEnterpriseDatabase].
func (red *DataRedisEnterpriseDatabase) Attributes() dataRedisEnterpriseDatabaseAttributes {
	return dataRedisEnterpriseDatabaseAttributes{ref: terra.ReferenceDataResource(red)}
}

// DataRedisEnterpriseDatabaseArgs contains the configurations for azurerm_redis_enterprise_database.
type DataRedisEnterpriseDatabaseArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Timeouts: optional
	Timeouts *dataredisenterprisedatabase.Timeouts `hcl:"timeouts,block"`
}
type dataRedisEnterpriseDatabaseAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("id"))
}

// LinkedDatabaseGroupNickname returns a reference to field linked_database_group_nickname of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) LinkedDatabaseGroupNickname() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("linked_database_group_nickname"))
}

// LinkedDatabaseId returns a reference to field linked_database_id of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) LinkedDatabaseId() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](red.ref.Append("linked_database_id"))
}

// Name returns a reference to field name of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_redis_enterprise_database.
func (red dataRedisEnterpriseDatabaseAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(red.ref.Append("secondary_access_key"))
}

func (red dataRedisEnterpriseDatabaseAttributes) Timeouts() dataredisenterprisedatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataredisenterprisedatabase.TimeoutsAttributes](red.ref.Append("timeouts"))
}
