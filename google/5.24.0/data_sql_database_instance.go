// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datasqldatabaseinstance "github.com/golingon/terraproviders/google/5.24.0/datasqldatabaseinstance"
)

// NewDataSqlDatabaseInstance creates a new instance of [DataSqlDatabaseInstance].
func NewDataSqlDatabaseInstance(name string, args DataSqlDatabaseInstanceArgs) *DataSqlDatabaseInstance {
	return &DataSqlDatabaseInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSqlDatabaseInstance)(nil)

// DataSqlDatabaseInstance represents the Terraform data resource google_sql_database_instance.
type DataSqlDatabaseInstance struct {
	Name string
	Args DataSqlDatabaseInstanceArgs
}

// DataSource returns the Terraform object type for [DataSqlDatabaseInstance].
func (sdi *DataSqlDatabaseInstance) DataSource() string {
	return "google_sql_database_instance"
}

// LocalName returns the local name for [DataSqlDatabaseInstance].
func (sdi *DataSqlDatabaseInstance) LocalName() string {
	return sdi.Name
}

// Configuration returns the configuration (args) for [DataSqlDatabaseInstance].
func (sdi *DataSqlDatabaseInstance) Configuration() interface{} {
	return sdi.Args
}

// Attributes returns the attributes for [DataSqlDatabaseInstance].
func (sdi *DataSqlDatabaseInstance) Attributes() dataSqlDatabaseInstanceAttributes {
	return dataSqlDatabaseInstanceAttributes{ref: terra.ReferenceDataResource(sdi)}
}

// DataSqlDatabaseInstanceArgs contains the configurations for google_sql_database_instance.
type DataSqlDatabaseInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Clone: min=0
	Clone []datasqldatabaseinstance.Clone `hcl:"clone,block" validate:"min=0"`
	// IpAddress: min=0
	IpAddress []datasqldatabaseinstance.IpAddress `hcl:"ip_address,block" validate:"min=0"`
	// ReplicaConfiguration: min=0
	ReplicaConfiguration []datasqldatabaseinstance.ReplicaConfiguration `hcl:"replica_configuration,block" validate:"min=0"`
	// RestoreBackupContext: min=0
	RestoreBackupContext []datasqldatabaseinstance.RestoreBackupContext `hcl:"restore_backup_context,block" validate:"min=0"`
	// ServerCaCert: min=0
	ServerCaCert []datasqldatabaseinstance.ServerCaCert `hcl:"server_ca_cert,block" validate:"min=0"`
	// Settings: min=0
	Settings []datasqldatabaseinstance.Settings `hcl:"settings,block" validate:"min=0"`
}
type dataSqlDatabaseInstanceAttributes struct {
	ref terra.Reference
}

// AvailableMaintenanceVersions returns a reference to field available_maintenance_versions of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) AvailableMaintenanceVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdi.ref.Append("available_maintenance_versions"))
}

// ConnectionName returns a reference to field connection_name of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("connection_name"))
}

// DatabaseVersion returns a reference to field database_version of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("database_version"))
}

// DeletionProtection returns a reference to field deletion_protection of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(sdi.ref.Append("deletion_protection"))
}

// DnsName returns a reference to field dns_name of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("dns_name"))
}

// EncryptionKeyName returns a reference to field encryption_key_name of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) EncryptionKeyName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("encryption_key_name"))
}

// FirstIpAddress returns a reference to field first_ip_address of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) FirstIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("first_ip_address"))
}

// Id returns a reference to field id of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("instance_type"))
}

// MaintenanceVersion returns a reference to field maintenance_version of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) MaintenanceVersion() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("maintenance_version"))
}

// MasterInstanceName returns a reference to field master_instance_name of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) MasterInstanceName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("master_instance_name"))
}

// Name returns a reference to field name of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("private_ip_address"))
}

// Project returns a reference to field project of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("project"))
}

// PscServiceAttachmentLink returns a reference to field psc_service_attachment_link of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) PscServiceAttachmentLink() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("psc_service_attachment_link"))
}

// PublicIpAddress returns a reference to field public_ip_address of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("public_ip_address"))
}

// Region returns a reference to field region of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("region"))
}

// RootPassword returns a reference to field root_password of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) RootPassword() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("root_password"))
}

// SelfLink returns a reference to field self_link of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("self_link"))
}

// ServiceAccountEmailAddress returns a reference to field service_account_email_address of google_sql_database_instance.
func (sdi dataSqlDatabaseInstanceAttributes) ServiceAccountEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("service_account_email_address"))
}

func (sdi dataSqlDatabaseInstanceAttributes) Clone() terra.ListValue[datasqldatabaseinstance.CloneAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.CloneAttributes](sdi.ref.Append("clone"))
}

func (sdi dataSqlDatabaseInstanceAttributes) IpAddress() terra.ListValue[datasqldatabaseinstance.IpAddressAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.IpAddressAttributes](sdi.ref.Append("ip_address"))
}

func (sdi dataSqlDatabaseInstanceAttributes) ReplicaConfiguration() terra.ListValue[datasqldatabaseinstance.ReplicaConfigurationAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.ReplicaConfigurationAttributes](sdi.ref.Append("replica_configuration"))
}

func (sdi dataSqlDatabaseInstanceAttributes) RestoreBackupContext() terra.ListValue[datasqldatabaseinstance.RestoreBackupContextAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.RestoreBackupContextAttributes](sdi.ref.Append("restore_backup_context"))
}

func (sdi dataSqlDatabaseInstanceAttributes) ServerCaCert() terra.ListValue[datasqldatabaseinstance.ServerCaCertAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.ServerCaCertAttributes](sdi.ref.Append("server_ca_cert"))
}

func (sdi dataSqlDatabaseInstanceAttributes) Settings() terra.ListValue[datasqldatabaseinstance.SettingsAttributes] {
	return terra.ReferenceAsList[datasqldatabaseinstance.SettingsAttributes](sdi.ref.Append("settings"))
}
