// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cosmosdbcassandracluster "github.com/golingon/terraproviders/azurerm/3.58.0/cosmosdbcassandracluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCosmosdbCassandraCluster creates a new instance of [CosmosdbCassandraCluster].
func NewCosmosdbCassandraCluster(name string, args CosmosdbCassandraClusterArgs) *CosmosdbCassandraCluster {
	return &CosmosdbCassandraCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CosmosdbCassandraCluster)(nil)

// CosmosdbCassandraCluster represents the Terraform resource azurerm_cosmosdb_cassandra_cluster.
type CosmosdbCassandraCluster struct {
	Name      string
	Args      CosmosdbCassandraClusterArgs
	state     *cosmosdbCassandraClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) Type() string {
	return "azurerm_cosmosdb_cassandra_cluster"
}

// LocalName returns the local name for [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) LocalName() string {
	return ccc.Name
}

// Configuration returns the configuration (args) for [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) Configuration() interface{} {
	return ccc.Args
}

// DependOn is used for other resources to depend on [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ccc)
}

// Dependencies returns the list of resources [CosmosdbCassandraCluster] depends_on.
func (ccc *CosmosdbCassandraCluster) Dependencies() terra.Dependencies {
	return ccc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) LifecycleManagement() *terra.Lifecycle {
	return ccc.Lifecycle
}

// Attributes returns the attributes for [CosmosdbCassandraCluster].
func (ccc *CosmosdbCassandraCluster) Attributes() cosmosdbCassandraClusterAttributes {
	return cosmosdbCassandraClusterAttributes{ref: terra.ReferenceResource(ccc)}
}

// ImportState imports the given attribute values into [CosmosdbCassandraCluster]'s state.
func (ccc *CosmosdbCassandraCluster) ImportState(av io.Reader) error {
	ccc.state = &cosmosdbCassandraClusterState{}
	if err := json.NewDecoder(av).Decode(ccc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccc.Type(), ccc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CosmosdbCassandraCluster] has state.
func (ccc *CosmosdbCassandraCluster) State() (*cosmosdbCassandraClusterState, bool) {
	return ccc.state, ccc.state != nil
}

// StateMust returns the state for [CosmosdbCassandraCluster]. Panics if the state is nil.
func (ccc *CosmosdbCassandraCluster) StateMust() *cosmosdbCassandraClusterState {
	if ccc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccc.Type(), ccc.LocalName()))
	}
	return ccc.state
}

// CosmosdbCassandraClusterArgs contains the configurations for azurerm_cosmosdb_cassandra_cluster.
type CosmosdbCassandraClusterArgs struct {
	// AuthenticationMethod: string, optional
	AuthenticationMethod terra.StringValue `hcl:"authentication_method,attr"`
	// ClientCertificatePems: list of string, optional
	ClientCertificatePems terra.ListValue[terra.StringValue] `hcl:"client_certificate_pems,attr"`
	// DefaultAdminPassword: string, required
	DefaultAdminPassword terra.StringValue `hcl:"default_admin_password,attr" validate:"required"`
	// DelegatedManagementSubnetId: string, required
	DelegatedManagementSubnetId terra.StringValue `hcl:"delegated_management_subnet_id,attr" validate:"required"`
	// ExternalGossipCertificatePems: list of string, optional
	ExternalGossipCertificatePems terra.ListValue[terra.StringValue] `hcl:"external_gossip_certificate_pems,attr"`
	// ExternalSeedNodeIpAddresses: list of string, optional
	ExternalSeedNodeIpAddresses terra.ListValue[terra.StringValue] `hcl:"external_seed_node_ip_addresses,attr"`
	// HoursBetweenBackups: number, optional
	HoursBetweenBackups terra.NumberValue `hcl:"hours_between_backups,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RepairEnabled: bool, optional
	RepairEnabled terra.BoolValue `hcl:"repair_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Identity: optional
	Identity *cosmosdbcassandracluster.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *cosmosdbcassandracluster.Timeouts `hcl:"timeouts,block"`
}
type cosmosdbCassandraClusterAttributes struct {
	ref terra.Reference
}

// AuthenticationMethod returns a reference to field authentication_method of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) AuthenticationMethod() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("authentication_method"))
}

// ClientCertificatePems returns a reference to field client_certificate_pems of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) ClientCertificatePems() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ccc.ref.Append("client_certificate_pems"))
}

// DefaultAdminPassword returns a reference to field default_admin_password of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) DefaultAdminPassword() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("default_admin_password"))
}

// DelegatedManagementSubnetId returns a reference to field delegated_management_subnet_id of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) DelegatedManagementSubnetId() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("delegated_management_subnet_id"))
}

// ExternalGossipCertificatePems returns a reference to field external_gossip_certificate_pems of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) ExternalGossipCertificatePems() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ccc.ref.Append("external_gossip_certificate_pems"))
}

// ExternalSeedNodeIpAddresses returns a reference to field external_seed_node_ip_addresses of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) ExternalSeedNodeIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ccc.ref.Append("external_seed_node_ip_addresses"))
}

// HoursBetweenBackups returns a reference to field hours_between_backups of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) HoursBetweenBackups() terra.NumberValue {
	return terra.ReferenceAsNumber(ccc.ref.Append("hours_between_backups"))
}

// Id returns a reference to field id of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("name"))
}

// RepairEnabled returns a reference to field repair_enabled of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) RepairEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ccc.ref.Append("repair_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccc.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_cosmosdb_cassandra_cluster.
func (ccc cosmosdbCassandraClusterAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("version"))
}

func (ccc cosmosdbCassandraClusterAttributes) Identity() terra.ListValue[cosmosdbcassandracluster.IdentityAttributes] {
	return terra.ReferenceAsList[cosmosdbcassandracluster.IdentityAttributes](ccc.ref.Append("identity"))
}

func (ccc cosmosdbCassandraClusterAttributes) Timeouts() cosmosdbcassandracluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cosmosdbcassandracluster.TimeoutsAttributes](ccc.ref.Append("timeouts"))
}

type cosmosdbCassandraClusterState struct {
	AuthenticationMethod          string                                   `json:"authentication_method"`
	ClientCertificatePems         []string                                 `json:"client_certificate_pems"`
	DefaultAdminPassword          string                                   `json:"default_admin_password"`
	DelegatedManagementSubnetId   string                                   `json:"delegated_management_subnet_id"`
	ExternalGossipCertificatePems []string                                 `json:"external_gossip_certificate_pems"`
	ExternalSeedNodeIpAddresses   []string                                 `json:"external_seed_node_ip_addresses"`
	HoursBetweenBackups           float64                                  `json:"hours_between_backups"`
	Id                            string                                   `json:"id"`
	Location                      string                                   `json:"location"`
	Name                          string                                   `json:"name"`
	RepairEnabled                 bool                                     `json:"repair_enabled"`
	ResourceGroupName             string                                   `json:"resource_group_name"`
	Tags                          map[string]string                        `json:"tags"`
	Version                       string                                   `json:"version"`
	Identity                      []cosmosdbcassandracluster.IdentityState `json:"identity"`
	Timeouts                      *cosmosdbcassandracluster.TimeoutsState  `json:"timeouts"`
}
