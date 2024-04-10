// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datamysqlserver "github.com/golingon/terraproviders/azurerm/3.98.0/datamysqlserver"
)

// NewDataMysqlServer creates a new instance of [DataMysqlServer].
func NewDataMysqlServer(name string, args DataMysqlServerArgs) *DataMysqlServer {
	return &DataMysqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMysqlServer)(nil)

// DataMysqlServer represents the Terraform data resource azurerm_mysql_server.
type DataMysqlServer struct {
	Name string
	Args DataMysqlServerArgs
}

// DataSource returns the Terraform object type for [DataMysqlServer].
func (ms *DataMysqlServer) DataSource() string {
	return "azurerm_mysql_server"
}

// LocalName returns the local name for [DataMysqlServer].
func (ms *DataMysqlServer) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [DataMysqlServer].
func (ms *DataMysqlServer) Configuration() interface{} {
	return ms.Args
}

// Attributes returns the attributes for [DataMysqlServer].
func (ms *DataMysqlServer) Attributes() dataMysqlServerAttributes {
	return dataMysqlServerAttributes{ref: terra.ReferenceDataResource(ms)}
}

// DataMysqlServerArgs contains the configurations for azurerm_mysql_server.
type DataMysqlServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datamysqlserver.Identity `hcl:"identity,block" validate:"min=0"`
	// ThreatDetectionPolicy: min=0
	ThreatDetectionPolicy []datamysqlserver.ThreatDetectionPolicy `hcl:"threat_detection_policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamysqlserver.Timeouts `hcl:"timeouts,block"`
}
type dataMysqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login"))
}

// AutoGrowEnabled returns a reference to field auto_grow_enabled of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) AutoGrowEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("auto_grow_enabled"))
}

// BackupRetentionDays returns a reference to field backup_retention_days of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) BackupRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("backup_retention_days"))
}

// Fqdn returns a reference to field fqdn of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("fqdn"))
}

// GeoRedundantBackupEnabled returns a reference to field geo_redundant_backup_enabled of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) GeoRedundantBackupEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("geo_redundant_backup_enabled"))
}

// Id returns a reference to field id of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("infrastructure_encryption_enabled"))
}

// Location returns a reference to field location of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("resource_group_name"))
}

// RestorePointInTime returns a reference to field restore_point_in_time of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) RestorePointInTime() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("restore_point_in_time"))
}

// SkuName returns a reference to field sku_name of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("sku_name"))
}

// SslEnforcementEnabled returns a reference to field ssl_enforcement_enabled of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) SslEnforcementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("ssl_enforcement_enabled"))
}

// SslMinimalTlsVersionEnforced returns a reference to field ssl_minimal_tls_version_enforced of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) SslMinimalTlsVersionEnforced() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("ssl_minimal_tls_version_enforced"))
}

// StorageMb returns a reference to field storage_mb of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) StorageMb() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("storage_mb"))
}

// Tags returns a reference to field tags of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_mysql_server.
func (ms dataMysqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("version"))
}

func (ms dataMysqlServerAttributes) Identity() terra.ListValue[datamysqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[datamysqlserver.IdentityAttributes](ms.ref.Append("identity"))
}

func (ms dataMysqlServerAttributes) ThreatDetectionPolicy() terra.ListValue[datamysqlserver.ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[datamysqlserver.ThreatDetectionPolicyAttributes](ms.ref.Append("threat_detection_policy"))
}

func (ms dataMysqlServerAttributes) Timeouts() datamysqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamysqlserver.TimeoutsAttributes](ms.ref.Append("timeouts"))
}
