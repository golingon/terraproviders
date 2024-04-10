// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datamssqldatabase "github.com/golingon/terraproviders/azurerm/3.98.0/datamssqldatabase"
)

// NewDataMssqlDatabase creates a new instance of [DataMssqlDatabase].
func NewDataMssqlDatabase(name string, args DataMssqlDatabaseArgs) *DataMssqlDatabase {
	return &DataMssqlDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMssqlDatabase)(nil)

// DataMssqlDatabase represents the Terraform data resource azurerm_mssql_database.
type DataMssqlDatabase struct {
	Name string
	Args DataMssqlDatabaseArgs
}

// DataSource returns the Terraform object type for [DataMssqlDatabase].
func (md *DataMssqlDatabase) DataSource() string {
	return "azurerm_mssql_database"
}

// LocalName returns the local name for [DataMssqlDatabase].
func (md *DataMssqlDatabase) LocalName() string {
	return md.Name
}

// Configuration returns the configuration (args) for [DataMssqlDatabase].
func (md *DataMssqlDatabase) Configuration() interface{} {
	return md.Args
}

// Attributes returns the attributes for [DataMssqlDatabase].
func (md *DataMssqlDatabase) Attributes() dataMssqlDatabaseAttributes {
	return dataMssqlDatabaseAttributes{ref: terra.ReferenceDataResource(md)}
}

// DataMssqlDatabaseArgs contains the configurations for azurerm_mssql_database.
type DataMssqlDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Identity: min=0
	Identity []datamssqldatabase.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamssqldatabase.Timeouts `hcl:"timeouts,block"`
}
type dataMssqlDatabaseAttributes struct {
	ref terra.Reference
}

// Collation returns a reference to field collation of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) Collation() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("collation"))
}

// ElasticPoolId returns a reference to field elastic_pool_id of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) ElasticPoolId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("elastic_pool_id"))
}

// EnclaveType returns a reference to field enclave_type of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) EnclaveType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("enclave_type"))
}

// Id returns a reference to field id of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("license_type"))
}

// MaxSizeGb returns a reference to field max_size_gb of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) MaxSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("max_size_gb"))
}

// Name returns a reference to field name of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("name"))
}

// ReadReplicaCount returns a reference to field read_replica_count of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) ReadReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("read_replica_count"))
}

// ReadScale returns a reference to field read_scale of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) ReadScale() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("read_scale"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("server_id"))
}

// SkuName returns a reference to field sku_name of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("sku_name"))
}

// StorageAccountType returns a reference to field storage_account_type of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("storage_account_type"))
}

// Tags returns a reference to field tags of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](md.ref.Append("tags"))
}

// TransparentDataEncryptionEnabled returns a reference to field transparent_data_encryption_enabled of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) TransparentDataEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("transparent_data_encryption_enabled"))
}

// TransparentDataEncryptionKeyAutomaticRotationEnabled returns a reference to field transparent_data_encryption_key_automatic_rotation_enabled of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) TransparentDataEncryptionKeyAutomaticRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("transparent_data_encryption_key_automatic_rotation_enabled"))
}

// TransparentDataEncryptionKeyVaultKeyId returns a reference to field transparent_data_encryption_key_vault_key_id of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) TransparentDataEncryptionKeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("transparent_data_encryption_key_vault_key_id"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_mssql_database.
func (md dataMssqlDatabaseAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(md.ref.Append("zone_redundant"))
}

func (md dataMssqlDatabaseAttributes) Identity() terra.ListValue[datamssqldatabase.IdentityAttributes] {
	return terra.ReferenceAsList[datamssqldatabase.IdentityAttributes](md.ref.Append("identity"))
}

func (md dataMssqlDatabaseAttributes) Timeouts() datamssqldatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamssqldatabase.TimeoutsAttributes](md.ref.Append("timeouts"))
}
