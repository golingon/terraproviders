// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sqldatabaseinstance "github.com/golingon/terraproviders/google/4.59.0/sqldatabaseinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlDatabaseInstance creates a new instance of [SqlDatabaseInstance].
func NewSqlDatabaseInstance(name string, args SqlDatabaseInstanceArgs) *SqlDatabaseInstance {
	return &SqlDatabaseInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlDatabaseInstance)(nil)

// SqlDatabaseInstance represents the Terraform resource google_sql_database_instance.
type SqlDatabaseInstance struct {
	Name      string
	Args      SqlDatabaseInstanceArgs
	state     *sqlDatabaseInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) Type() string {
	return "google_sql_database_instance"
}

// LocalName returns the local name for [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) LocalName() string {
	return sdi.Name
}

// Configuration returns the configuration (args) for [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) Configuration() interface{} {
	return sdi.Args
}

// DependOn is used for other resources to depend on [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(sdi)
}

// Dependencies returns the list of resources [SqlDatabaseInstance] depends_on.
func (sdi *SqlDatabaseInstance) Dependencies() terra.Dependencies {
	return sdi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) LifecycleManagement() *terra.Lifecycle {
	return sdi.Lifecycle
}

// Attributes returns the attributes for [SqlDatabaseInstance].
func (sdi *SqlDatabaseInstance) Attributes() sqlDatabaseInstanceAttributes {
	return sqlDatabaseInstanceAttributes{ref: terra.ReferenceResource(sdi)}
}

// ImportState imports the given attribute values into [SqlDatabaseInstance]'s state.
func (sdi *SqlDatabaseInstance) ImportState(av io.Reader) error {
	sdi.state = &sqlDatabaseInstanceState{}
	if err := json.NewDecoder(av).Decode(sdi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdi.Type(), sdi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlDatabaseInstance] has state.
func (sdi *SqlDatabaseInstance) State() (*sqlDatabaseInstanceState, bool) {
	return sdi.state, sdi.state != nil
}

// StateMust returns the state for [SqlDatabaseInstance]. Panics if the state is nil.
func (sdi *SqlDatabaseInstance) StateMust() *sqlDatabaseInstanceState {
	if sdi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdi.Type(), sdi.LocalName()))
	}
	return sdi.state
}

// SqlDatabaseInstanceArgs contains the configurations for google_sql_database_instance.
type SqlDatabaseInstanceArgs struct {
	// DatabaseVersion: string, required
	DatabaseVersion terra.StringValue `hcl:"database_version,attr" validate:"required"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// EncryptionKeyName: string, optional
	EncryptionKeyName terra.StringValue `hcl:"encryption_key_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// MaintenanceVersion: string, optional
	MaintenanceVersion terra.StringValue `hcl:"maintenance_version,attr"`
	// MasterInstanceName: string, optional
	MasterInstanceName terra.StringValue `hcl:"master_instance_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RootPassword: string, optional
	RootPassword terra.StringValue `hcl:"root_password,attr"`
	// IpAddress: min=0
	IpAddress []sqldatabaseinstance.IpAddress `hcl:"ip_address,block" validate:"min=0"`
	// ServerCaCert: min=0
	ServerCaCert []sqldatabaseinstance.ServerCaCert `hcl:"server_ca_cert,block" validate:"min=0"`
	// Clone: optional
	Clone *sqldatabaseinstance.Clone `hcl:"clone,block"`
	// ReplicaConfiguration: optional
	ReplicaConfiguration *sqldatabaseinstance.ReplicaConfiguration `hcl:"replica_configuration,block"`
	// RestoreBackupContext: optional
	RestoreBackupContext *sqldatabaseinstance.RestoreBackupContext `hcl:"restore_backup_context,block"`
	// Settings: optional
	Settings *sqldatabaseinstance.Settings `hcl:"settings,block"`
	// Timeouts: optional
	Timeouts *sqldatabaseinstance.Timeouts `hcl:"timeouts,block"`
}
type sqlDatabaseInstanceAttributes struct {
	ref terra.Reference
}

// AvailableMaintenanceVersions returns a reference to field available_maintenance_versions of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) AvailableMaintenanceVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdi.ref.Append("available_maintenance_versions"))
}

// ConnectionName returns a reference to field connection_name of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("connection_name"))
}

// DatabaseVersion returns a reference to field database_version of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("database_version"))
}

// DeletionProtection returns a reference to field deletion_protection of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(sdi.ref.Append("deletion_protection"))
}

// EncryptionKeyName returns a reference to field encryption_key_name of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) EncryptionKeyName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("encryption_key_name"))
}

// FirstIpAddress returns a reference to field first_ip_address of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) FirstIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("first_ip_address"))
}

// Id returns a reference to field id of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("instance_type"))
}

// MaintenanceVersion returns a reference to field maintenance_version of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) MaintenanceVersion() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("maintenance_version"))
}

// MasterInstanceName returns a reference to field master_instance_name of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) MasterInstanceName() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("master_instance_name"))
}

// Name returns a reference to field name of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("private_ip_address"))
}

// Project returns a reference to field project of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("project"))
}

// PublicIpAddress returns a reference to field public_ip_address of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("public_ip_address"))
}

// Region returns a reference to field region of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("region"))
}

// RootPassword returns a reference to field root_password of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) RootPassword() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("root_password"))
}

// SelfLink returns a reference to field self_link of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("self_link"))
}

// ServiceAccountEmailAddress returns a reference to field service_account_email_address of google_sql_database_instance.
func (sdi sqlDatabaseInstanceAttributes) ServiceAccountEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("service_account_email_address"))
}

func (sdi sqlDatabaseInstanceAttributes) IpAddress() terra.ListValue[sqldatabaseinstance.IpAddressAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.IpAddressAttributes](sdi.ref.Append("ip_address"))
}

func (sdi sqlDatabaseInstanceAttributes) ServerCaCert() terra.ListValue[sqldatabaseinstance.ServerCaCertAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.ServerCaCertAttributes](sdi.ref.Append("server_ca_cert"))
}

func (sdi sqlDatabaseInstanceAttributes) Clone() terra.ListValue[sqldatabaseinstance.CloneAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.CloneAttributes](sdi.ref.Append("clone"))
}

func (sdi sqlDatabaseInstanceAttributes) ReplicaConfiguration() terra.ListValue[sqldatabaseinstance.ReplicaConfigurationAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.ReplicaConfigurationAttributes](sdi.ref.Append("replica_configuration"))
}

func (sdi sqlDatabaseInstanceAttributes) RestoreBackupContext() terra.ListValue[sqldatabaseinstance.RestoreBackupContextAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.RestoreBackupContextAttributes](sdi.ref.Append("restore_backup_context"))
}

func (sdi sqlDatabaseInstanceAttributes) Settings() terra.ListValue[sqldatabaseinstance.SettingsAttributes] {
	return terra.ReferenceAsList[sqldatabaseinstance.SettingsAttributes](sdi.ref.Append("settings"))
}

func (sdi sqlDatabaseInstanceAttributes) Timeouts() sqldatabaseinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqldatabaseinstance.TimeoutsAttributes](sdi.ref.Append("timeouts"))
}

type sqlDatabaseInstanceState struct {
	AvailableMaintenanceVersions []string                                        `json:"available_maintenance_versions"`
	ConnectionName               string                                          `json:"connection_name"`
	DatabaseVersion              string                                          `json:"database_version"`
	DeletionProtection           bool                                            `json:"deletion_protection"`
	EncryptionKeyName            string                                          `json:"encryption_key_name"`
	FirstIpAddress               string                                          `json:"first_ip_address"`
	Id                           string                                          `json:"id"`
	InstanceType                 string                                          `json:"instance_type"`
	MaintenanceVersion           string                                          `json:"maintenance_version"`
	MasterInstanceName           string                                          `json:"master_instance_name"`
	Name                         string                                          `json:"name"`
	PrivateIpAddress             string                                          `json:"private_ip_address"`
	Project                      string                                          `json:"project"`
	PublicIpAddress              string                                          `json:"public_ip_address"`
	Region                       string                                          `json:"region"`
	RootPassword                 string                                          `json:"root_password"`
	SelfLink                     string                                          `json:"self_link"`
	ServiceAccountEmailAddress   string                                          `json:"service_account_email_address"`
	IpAddress                    []sqldatabaseinstance.IpAddressState            `json:"ip_address"`
	ServerCaCert                 []sqldatabaseinstance.ServerCaCertState         `json:"server_ca_cert"`
	Clone                        []sqldatabaseinstance.CloneState                `json:"clone"`
	ReplicaConfiguration         []sqldatabaseinstance.ReplicaConfigurationState `json:"replica_configuration"`
	RestoreBackupContext         []sqldatabaseinstance.RestoreBackupContextState `json:"restore_backup_context"`
	Settings                     []sqldatabaseinstance.SettingsState             `json:"settings"`
	Timeouts                     *sqldatabaseinstance.TimeoutsState              `json:"timeouts"`
}
