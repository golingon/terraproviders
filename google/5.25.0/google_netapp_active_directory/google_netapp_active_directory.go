// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_netapp_active_directory

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

// Resource represents the Terraform resource google_netapp_active_directory.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNetappActiveDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnad *Resource) Type() string {
	return "google_netapp_active_directory"
}

// LocalName returns the local name for [Resource].
func (gnad *Resource) LocalName() string {
	return gnad.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnad *Resource) Configuration() interface{} {
	return gnad.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnad *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnad)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnad *Resource) Dependencies() terra.Dependencies {
	return gnad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnad *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnad.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnad *Resource) Attributes() googleNetappActiveDirectoryAttributes {
	return googleNetappActiveDirectoryAttributes{ref: terra.ReferenceResource(gnad)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnad *Resource) ImportState(state io.Reader) error {
	gnad.state = &googleNetappActiveDirectoryState{}
	if err := json.NewDecoder(state).Decode(gnad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnad.Type(), gnad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnad *Resource) State() (*googleNetappActiveDirectoryState, bool) {
	return gnad.state, gnad.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnad *Resource) StateMust() *googleNetappActiveDirectoryState {
	if gnad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnad.Type(), gnad.LocalName()))
	}
	return gnad.state
}

// Args contains the configurations for google_netapp_active_directory.
type Args struct {
	// AesEncryption: bool, optional
	AesEncryption terra.BoolValue `hcl:"aes_encryption,attr"`
	// BackupOperators: list of string, optional
	BackupOperators terra.ListValue[terra.StringValue] `hcl:"backup_operators,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Dns: string, required
	Dns terra.StringValue `hcl:"dns,attr" validate:"required"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// EncryptDcConnections: bool, optional
	EncryptDcConnections terra.BoolValue `hcl:"encrypt_dc_connections,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KdcHostname: string, optional
	KdcHostname terra.StringValue `hcl:"kdc_hostname,attr"`
	// KdcIp: string, optional
	KdcIp terra.StringValue `hcl:"kdc_ip,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LdapSigning: bool, optional
	LdapSigning terra.BoolValue `hcl:"ldap_signing,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetBiosPrefix: string, required
	NetBiosPrefix terra.StringValue `hcl:"net_bios_prefix,attr" validate:"required"`
	// NfsUsersWithLdap: bool, optional
	NfsUsersWithLdap terra.BoolValue `hcl:"nfs_users_with_ldap,attr"`
	// OrganizationalUnit: string, optional
	OrganizationalUnit terra.StringValue `hcl:"organizational_unit,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecurityOperators: list of string, optional
	SecurityOperators terra.ListValue[terra.StringValue] `hcl:"security_operators,attr"`
	// Site: string, optional
	Site terra.StringValue `hcl:"site,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleNetappActiveDirectoryAttributes struct {
	ref terra.Reference
}

// AesEncryption returns a reference to field aes_encryption of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) AesEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(gnad.ref.Append("aes_encryption"))
}

// BackupOperators returns a reference to field backup_operators of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) BackupOperators() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gnad.ref.Append("backup_operators"))
}

// CreateTime returns a reference to field create_time of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("create_time"))
}

// Description returns a reference to field description of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("description"))
}

// Dns returns a reference to field dns of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Dns() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("dns"))
}

// Domain returns a reference to field domain of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("domain"))
}

// EffectiveLabels returns a reference to field effective_labels of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnad.ref.Append("effective_labels"))
}

// EncryptDcConnections returns a reference to field encrypt_dc_connections of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) EncryptDcConnections() terra.BoolValue {
	return terra.ReferenceAsBool(gnad.ref.Append("encrypt_dc_connections"))
}

// Id returns a reference to field id of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("id"))
}

// KdcHostname returns a reference to field kdc_hostname of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) KdcHostname() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("kdc_hostname"))
}

// KdcIp returns a reference to field kdc_ip of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) KdcIp() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("kdc_ip"))
}

// Labels returns a reference to field labels of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnad.ref.Append("labels"))
}

// LdapSigning returns a reference to field ldap_signing of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) LdapSigning() terra.BoolValue {
	return terra.ReferenceAsBool(gnad.ref.Append("ldap_signing"))
}

// Location returns a reference to field location of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("location"))
}

// Name returns a reference to field name of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("name"))
}

// NetBiosPrefix returns a reference to field net_bios_prefix of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) NetBiosPrefix() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("net_bios_prefix"))
}

// NfsUsersWithLdap returns a reference to field nfs_users_with_ldap of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) NfsUsersWithLdap() terra.BoolValue {
	return terra.ReferenceAsBool(gnad.ref.Append("nfs_users_with_ldap"))
}

// OrganizationalUnit returns a reference to field organizational_unit of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) OrganizationalUnit() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("organizational_unit"))
}

// Password returns a reference to field password of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("password"))
}

// Project returns a reference to field project of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("project"))
}

// SecurityOperators returns a reference to field security_operators of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) SecurityOperators() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gnad.ref.Append("security_operators"))
}

// Site returns a reference to field site of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Site() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("site"))
}

// State returns a reference to field state of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("state"))
}

// StateDetails returns a reference to field state_details of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) StateDetails() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("state_details"))
}

// TerraformLabels returns a reference to field terraform_labels of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gnad.ref.Append("terraform_labels"))
}

// Username returns a reference to field username of google_netapp_active_directory.
func (gnad googleNetappActiveDirectoryAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(gnad.ref.Append("username"))
}

func (gnad googleNetappActiveDirectoryAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gnad.ref.Append("timeouts"))
}

type googleNetappActiveDirectoryState struct {
	AesEncryption        bool              `json:"aes_encryption"`
	BackupOperators      []string          `json:"backup_operators"`
	CreateTime           string            `json:"create_time"`
	Description          string            `json:"description"`
	Dns                  string            `json:"dns"`
	Domain               string            `json:"domain"`
	EffectiveLabels      map[string]string `json:"effective_labels"`
	EncryptDcConnections bool              `json:"encrypt_dc_connections"`
	Id                   string            `json:"id"`
	KdcHostname          string            `json:"kdc_hostname"`
	KdcIp                string            `json:"kdc_ip"`
	Labels               map[string]string `json:"labels"`
	LdapSigning          bool              `json:"ldap_signing"`
	Location             string            `json:"location"`
	Name                 string            `json:"name"`
	NetBiosPrefix        string            `json:"net_bios_prefix"`
	NfsUsersWithLdap     bool              `json:"nfs_users_with_ldap"`
	OrganizationalUnit   string            `json:"organizational_unit"`
	Password             string            `json:"password"`
	Project              string            `json:"project"`
	SecurityOperators    []string          `json:"security_operators"`
	Site                 string            `json:"site"`
	State                string            `json:"state"`
	StateDetails         string            `json:"state_details"`
	TerraformLabels      map[string]string `json:"terraform_labels"`
	Username             string            `json:"username"`
	Timeouts             *TimeoutsState    `json:"timeouts"`
}
