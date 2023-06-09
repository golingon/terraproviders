// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasearchservice "github.com/golingon/terraproviders/azurerm/3.49.0/datasearchservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSearchService creates a new instance of [DataSearchService].
func NewDataSearchService(name string, args DataSearchServiceArgs) *DataSearchService {
	return &DataSearchService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSearchService)(nil)

// DataSearchService represents the Terraform data resource azurerm_search_service.
type DataSearchService struct {
	Name string
	Args DataSearchServiceArgs
}

// DataSource returns the Terraform object type for [DataSearchService].
func (ss *DataSearchService) DataSource() string {
	return "azurerm_search_service"
}

// LocalName returns the local name for [DataSearchService].
func (ss *DataSearchService) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataSearchService].
func (ss *DataSearchService) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataSearchService].
func (ss *DataSearchService) Attributes() dataSearchServiceAttributes {
	return dataSearchServiceAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataSearchServiceArgs contains the configurations for azurerm_search_service.
type DataSearchServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datasearchservice.Identity `hcl:"identity,block" validate:"min=0"`
	// QueryKeys: min=0
	QueryKeys []datasearchservice.QueryKeys `hcl:"query_keys,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasearchservice.Timeouts `hcl:"timeouts,block"`
}
type dataSearchServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_search_service.
func (ss dataSearchServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_search_service.
func (ss dataSearchServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// PartitionCount returns a reference to field partition_count of azurerm_search_service.
func (ss dataSearchServiceAttributes) PartitionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("partition_count"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_search_service.
func (ss dataSearchServiceAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("primary_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_search_service.
func (ss dataSearchServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("public_network_access_enabled"))
}

// ReplicaCount returns a reference to field replica_count of azurerm_search_service.
func (ss dataSearchServiceAttributes) ReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("replica_count"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_search_service.
func (ss dataSearchServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_group_name"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_search_service.
func (ss dataSearchServiceAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("secondary_key"))
}

func (ss dataSearchServiceAttributes) Identity() terra.ListValue[datasearchservice.IdentityAttributes] {
	return terra.ReferenceAsList[datasearchservice.IdentityAttributes](ss.ref.Append("identity"))
}

func (ss dataSearchServiceAttributes) QueryKeys() terra.ListValue[datasearchservice.QueryKeysAttributes] {
	return terra.ReferenceAsList[datasearchservice.QueryKeysAttributes](ss.ref.Append("query_keys"))
}

func (ss dataSearchServiceAttributes) Timeouts() datasearchservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasearchservice.TimeoutsAttributes](ss.ref.Append("timeouts"))
}
