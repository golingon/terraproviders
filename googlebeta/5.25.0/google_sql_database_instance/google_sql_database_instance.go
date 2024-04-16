// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_sql_database_instance

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_sql_database_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSqlDatabaseInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gsdi *Resource) Type() string {
	return "google_sql_database_instance"
}

// LocalName returns the local name for [Resource].
func (gsdi *Resource) LocalName() string {
	return gsdi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gsdi *Resource) Configuration() interface{} {
	return gsdi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gsdi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gsdi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gsdi *Resource) Dependencies() terra.Dependencies {
	return gsdi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gsdi *Resource) LifecycleManagement() *terra.Lifecycle {
	return gsdi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gsdi *Resource) Attributes() googleSqlDatabaseInstanceAttributes {
	return googleSqlDatabaseInstanceAttributes{ref: terra.ReferenceResource(gsdi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gsdi *Resource) ImportState(state io.Reader) error {
	gsdi.state = &googleSqlDatabaseInstanceState{}
	if err := json.NewDecoder(state).Decode(gsdi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsdi.Type(), gsdi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gsdi *Resource) State() (*googleSqlDatabaseInstanceState, bool) {
	return gsdi.state, gsdi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gsdi *Resource) StateMust() *googleSqlDatabaseInstanceState {
	if gsdi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsdi.Type(), gsdi.LocalName()))
	}
	return gsdi.state
}

// Args contains the configurations for google_sql_database_instance.
type Args struct {
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
	// Clone: optional
	Clone *Clone `hcl:"clone,block"`
	// ReplicaConfiguration: optional
	ReplicaConfiguration *ReplicaConfiguration `hcl:"replica_configuration,block"`
	// RestoreBackupContext: optional
	RestoreBackupContext *RestoreBackupContext `hcl:"restore_backup_context,block"`
	// Settings: optional
	Settings *Settings `hcl:"settings,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleSqlDatabaseInstanceAttributes struct {
	ref terra.Reference
}

// AvailableMaintenanceVersions returns a reference to field available_maintenance_versions of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) AvailableMaintenanceVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gsdi.ref.Append("available_maintenance_versions"))
}

// ConnectionName returns a reference to field connection_name of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) ConnectionName() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("connection_name"))
}

// DatabaseVersion returns a reference to field database_version of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) DatabaseVersion() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("database_version"))
}

// DeletionProtection returns a reference to field deletion_protection of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(gsdi.ref.Append("deletion_protection"))
}

// DnsName returns a reference to field dns_name of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("dns_name"))
}

// EncryptionKeyName returns a reference to field encryption_key_name of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) EncryptionKeyName() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("encryption_key_name"))
}

// FirstIpAddress returns a reference to field first_ip_address of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) FirstIpAddress() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("first_ip_address"))
}

// Id returns a reference to field id of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("instance_type"))
}

// MaintenanceVersion returns a reference to field maintenance_version of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) MaintenanceVersion() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("maintenance_version"))
}

// MasterInstanceName returns a reference to field master_instance_name of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) MasterInstanceName() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("master_instance_name"))
}

// Name returns a reference to field name of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("private_ip_address"))
}

// Project returns a reference to field project of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("project"))
}

// PscServiceAttachmentLink returns a reference to field psc_service_attachment_link of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) PscServiceAttachmentLink() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("psc_service_attachment_link"))
}

// PublicIpAddress returns a reference to field public_ip_address of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("public_ip_address"))
}

// Region returns a reference to field region of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("region"))
}

// RootPassword returns a reference to field root_password of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) RootPassword() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("root_password"))
}

// SelfLink returns a reference to field self_link of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("self_link"))
}

// ServiceAccountEmailAddress returns a reference to field service_account_email_address of google_sql_database_instance.
func (gsdi googleSqlDatabaseInstanceAttributes) ServiceAccountEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(gsdi.ref.Append("service_account_email_address"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) IpAddress() terra.ListValue[IpAddressAttributes] {
	return terra.ReferenceAsList[IpAddressAttributes](gsdi.ref.Append("ip_address"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) ServerCaCert() terra.ListValue[ServerCaCertAttributes] {
	return terra.ReferenceAsList[ServerCaCertAttributes](gsdi.ref.Append("server_ca_cert"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) Clone() terra.ListValue[CloneAttributes] {
	return terra.ReferenceAsList[CloneAttributes](gsdi.ref.Append("clone"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) ReplicaConfiguration() terra.ListValue[ReplicaConfigurationAttributes] {
	return terra.ReferenceAsList[ReplicaConfigurationAttributes](gsdi.ref.Append("replica_configuration"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) RestoreBackupContext() terra.ListValue[RestoreBackupContextAttributes] {
	return terra.ReferenceAsList[RestoreBackupContextAttributes](gsdi.ref.Append("restore_backup_context"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) Settings() terra.ListValue[SettingsAttributes] {
	return terra.ReferenceAsList[SettingsAttributes](gsdi.ref.Append("settings"))
}

func (gsdi googleSqlDatabaseInstanceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gsdi.ref.Append("timeouts"))
}

type googleSqlDatabaseInstanceState struct {
	AvailableMaintenanceVersions []string                    `json:"available_maintenance_versions"`
	ConnectionName               string                      `json:"connection_name"`
	DatabaseVersion              string                      `json:"database_version"`
	DeletionProtection           bool                        `json:"deletion_protection"`
	DnsName                      string                      `json:"dns_name"`
	EncryptionKeyName            string                      `json:"encryption_key_name"`
	FirstIpAddress               string                      `json:"first_ip_address"`
	Id                           string                      `json:"id"`
	InstanceType                 string                      `json:"instance_type"`
	MaintenanceVersion           string                      `json:"maintenance_version"`
	MasterInstanceName           string                      `json:"master_instance_name"`
	Name                         string                      `json:"name"`
	PrivateIpAddress             string                      `json:"private_ip_address"`
	Project                      string                      `json:"project"`
	PscServiceAttachmentLink     string                      `json:"psc_service_attachment_link"`
	PublicIpAddress              string                      `json:"public_ip_address"`
	Region                       string                      `json:"region"`
	RootPassword                 string                      `json:"root_password"`
	SelfLink                     string                      `json:"self_link"`
	ServiceAccountEmailAddress   string                      `json:"service_account_email_address"`
	IpAddress                    []IpAddressState            `json:"ip_address"`
	ServerCaCert                 []ServerCaCertState         `json:"server_ca_cert"`
	Clone                        []CloneState                `json:"clone"`
	ReplicaConfiguration         []ReplicaConfigurationState `json:"replica_configuration"`
	RestoreBackupContext         []RestoreBackupContextState `json:"restore_backup_context"`
	Settings                     []SettingsState             `json:"settings"`
	Timeouts                     *TimeoutsState              `json:"timeouts"`
}
