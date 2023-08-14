// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbcassandradatacenter "github.com/golingon/terraproviders/azurerm/3.69.0/cosmosdbcassandradatacenter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbCassandraDatacenter creates a new instance of [CosmosdbCassandraDatacenter].
func NewCosmosdbCassandraDatacenter(name string, args CosmosdbCassandraDatacenterArgs) *CosmosdbCassandraDatacenter {
	return &CosmosdbCassandraDatacenter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbCassandraDatacenter)(nil)

// CosmosdbCassandraDatacenter represents the Terraform resource azurerm_cosmosdb_cassandra_datacenter.
type CosmosdbCassandraDatacenter struct {
	Name      string
	Args      CosmosdbCassandraDatacenterArgs
	state     *cosmosdbCassandraDatacenterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) Type() string {
	return "azurerm_cosmosdb_cassandra_datacenter"
}

// LocalName returns the local name for [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) LocalName() string {
	return ccd.Name
}

// Configuration returns the configuration (args) for [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) Configuration() interface{} {
	return ccd.Args
}

// DependOn is used for other resources to depend on [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) DependOn() terra.Reference {
	return terra.ReferenceResource(ccd)
}

// Dependencies returns the list of resources [CosmosdbCassandraDatacenter] depends_on.
func (ccd *CosmosdbCassandraDatacenter) Dependencies() terra.Dependencies {
	return ccd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) LifecycleManagement() *terra.Lifecycle {
	return ccd.Lifecycle
}

// Attributes returns the attributes for [CosmosdbCassandraDatacenter].
func (ccd *CosmosdbCassandraDatacenter) Attributes() cosmosdbCassandraDatacenterAttributes {
	return cosmosdbCassandraDatacenterAttributes{ref: terra.ReferenceResource(ccd)}
}

// ImportState imports the given attribute values into [CosmosdbCassandraDatacenter]'s state.
func (ccd *CosmosdbCassandraDatacenter) ImportState(av io.Reader) error {
	ccd.state = &cosmosdbCassandraDatacenterState{}
	if err := json.NewDecoder(av).Decode(ccd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccd.Type(), ccd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbCassandraDatacenter] has state.
func (ccd *CosmosdbCassandraDatacenter) State() (*cosmosdbCassandraDatacenterState, bool) {
	return ccd.state, ccd.state != nil
}

// StateMust returns the state for [CosmosdbCassandraDatacenter]. Panics if the state is nil.
func (ccd *CosmosdbCassandraDatacenter) StateMust() *cosmosdbCassandraDatacenterState {
	if ccd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccd.Type(), ccd.LocalName()))
	}
	return ccd.state
}

// CosmosdbCassandraDatacenterArgs contains the configurations for azurerm_cosmosdb_cassandra_datacenter.
type CosmosdbCassandraDatacenterArgs struct {
	// AvailabilityZonesEnabled: bool, optional
	AvailabilityZonesEnabled terra.BoolValue `hcl:"availability_zones_enabled,attr"`
	// BackupStorageCustomerKeyUri: string, optional
	BackupStorageCustomerKeyUri terra.StringValue `hcl:"backup_storage_customer_key_uri,attr"`
	// Base64EncodedYamlFragment: string, optional
	Base64EncodedYamlFragment terra.StringValue `hcl:"base64_encoded_yaml_fragment,attr"`
	// CassandraClusterId: string, required
	CassandraClusterId terra.StringValue `hcl:"cassandra_cluster_id,attr" validate:"required"`
	// DelegatedManagementSubnetId: string, required
	DelegatedManagementSubnetId terra.StringValue `hcl:"delegated_management_subnet_id,attr" validate:"required"`
	// DiskCount: number, optional
	DiskCount terra.NumberValue `hcl:"disk_count,attr"`
	// DiskSku: string, optional
	DiskSku terra.StringValue `hcl:"disk_sku,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedDiskCustomerKeyUri: string, optional
	ManagedDiskCustomerKeyUri terra.StringValue `hcl:"managed_disk_customer_key_uri,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeCount: number, optional
	NodeCount terra.NumberValue `hcl:"node_count,attr"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// Timeouts: optional
	Timeouts *cosmosdbcassandradatacenter.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbCassandraDatacenterAttributes struct {
	ref terra.Reference
}

// AvailabilityZonesEnabled returns a reference to field availability_zones_enabled of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) AvailabilityZonesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ccd.ref.Append("availability_zones_enabled"))
}

// BackupStorageCustomerKeyUri returns a reference to field backup_storage_customer_key_uri of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) BackupStorageCustomerKeyUri() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("backup_storage_customer_key_uri"))
}

// Base64EncodedYamlFragment returns a reference to field base64_encoded_yaml_fragment of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) Base64EncodedYamlFragment() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("base64_encoded_yaml_fragment"))
}

// CassandraClusterId returns a reference to field cassandra_cluster_id of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) CassandraClusterId() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("cassandra_cluster_id"))
}

// DelegatedManagementSubnetId returns a reference to field delegated_management_subnet_id of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) DelegatedManagementSubnetId() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("delegated_management_subnet_id"))
}

// DiskCount returns a reference to field disk_count of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) DiskCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ccd.ref.Append("disk_count"))
}

// DiskSku returns a reference to field disk_sku of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) DiskSku() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("disk_sku"))
}

// Id returns a reference to field id of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("location"))
}

// ManagedDiskCustomerKeyUri returns a reference to field managed_disk_customer_key_uri of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) ManagedDiskCustomerKeyUri() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("managed_disk_customer_key_uri"))
}

// Name returns a reference to field name of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("name"))
}

// NodeCount returns a reference to field node_count of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ccd.ref.Append("node_count"))
}

// SkuName returns a reference to field sku_name of azurerm_cosmosdb_cassandra_datacenter.
func (ccd cosmosdbCassandraDatacenterAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ccd.ref.Append("sku_name"))
}

func (ccd cosmosdbCassandraDatacenterAttributes) Timeouts() cosmosdbcassandradatacenter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbcassandradatacenter.TimeoutsAttributes](ccd.ref.Append("timeouts"))
}

type cosmosdbCassandraDatacenterState struct {
	AvailabilityZonesEnabled    bool                                       `json:"availability_zones_enabled"`
	BackupStorageCustomerKeyUri string                                     `json:"backup_storage_customer_key_uri"`
	Base64EncodedYamlFragment   string                                     `json:"base64_encoded_yaml_fragment"`
	CassandraClusterId          string                                     `json:"cassandra_cluster_id"`
	DelegatedManagementSubnetId string                                     `json:"delegated_management_subnet_id"`
	DiskCount                   float64                                    `json:"disk_count"`
	DiskSku                     string                                     `json:"disk_sku"`
	Id                          string                                     `json:"id"`
	Location                    string                                     `json:"location"`
	ManagedDiskCustomerKeyUri   string                                     `json:"managed_disk_customer_key_uri"`
	Name                        string                                     `json:"name"`
	NodeCount                   float64                                    `json:"node_count"`
	SkuName                     string                                     `json:"sku_name"`
	Timeouts                    *cosmosdbcassandradatacenter.TimeoutsState `json:"timeouts"`
}
