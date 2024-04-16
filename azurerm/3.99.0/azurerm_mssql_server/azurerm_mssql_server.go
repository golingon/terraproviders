// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mssql_server

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

// Resource represents the Terraform resource azurerm_mssql_server.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMssqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ams *Resource) Type() string {
	return "azurerm_mssql_server"
}

// LocalName returns the local name for [Resource].
func (ams *Resource) LocalName() string {
	return ams.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ams *Resource) Configuration() interface{} {
	return ams.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ams *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ams)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ams *Resource) Dependencies() terra.Dependencies {
	return ams.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ams *Resource) LifecycleManagement() *terra.Lifecycle {
	return ams.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ams *Resource) Attributes() azurermMssqlServerAttributes {
	return azurermMssqlServerAttributes{ref: terra.ReferenceResource(ams)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ams *Resource) ImportState(state io.Reader) error {
	ams.state = &azurermMssqlServerState{}
	if err := json.NewDecoder(state).Decode(ams.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ams.Type(), ams.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ams *Resource) State() (*azurermMssqlServerState, bool) {
	return ams.state, ams.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ams *Resource) StateMust() *azurermMssqlServerState {
	if ams.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ams.Type(), ams.LocalName()))
	}
	return ams.state
}

// Args contains the configurations for azurerm_mssql_server.
type Args struct {
	// AdministratorLogin: string, optional
	AdministratorLogin terra.StringValue `hcl:"administrator_login,attr"`
	// AdministratorLoginPassword: string, optional
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr"`
	// ConnectionPolicy: string, optional
	ConnectionPolicy terra.StringValue `hcl:"connection_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinimumTlsVersion: string, optional
	MinimumTlsVersion terra.StringValue `hcl:"minimum_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutboundNetworkRestrictionEnabled: bool, optional
	OutboundNetworkRestrictionEnabled terra.BoolValue `hcl:"outbound_network_restriction_enabled,attr"`
	// PrimaryUserAssignedIdentityId: string, optional
	PrimaryUserAssignedIdentityId terra.StringValue `hcl:"primary_user_assigned_identity_id,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransparentDataEncryptionKeyVaultKeyId: string, optional
	TransparentDataEncryptionKeyVaultKeyId terra.StringValue `hcl:"transparent_data_encryption_key_vault_key_id,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// AzureadAdministrator: optional
	AzureadAdministrator *AzureadAdministrator `hcl:"azuread_administrator,block"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMssqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("administrator_login_password"))
}

// ConnectionPolicy returns a reference to field connection_policy of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) ConnectionPolicy() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("connection_policy"))
}

// FullyQualifiedDomainName returns a reference to field fully_qualified_domain_name of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) FullyQualifiedDomainName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("fully_qualified_domain_name"))
}

// Id returns a reference to field id of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("name"))
}

// OutboundNetworkRestrictionEnabled returns a reference to field outbound_network_restriction_enabled of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) OutboundNetworkRestrictionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ams.ref.Append("outbound_network_restriction_enabled"))
}

// PrimaryUserAssignedIdentityId returns a reference to field primary_user_assigned_identity_id of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) PrimaryUserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("primary_user_assigned_identity_id"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ams.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("resource_group_name"))
}

// RestorableDroppedDatabaseIds returns a reference to field restorable_dropped_database_ids of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) RestorableDroppedDatabaseIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ams.ref.Append("restorable_dropped_database_ids"))
}

// Tags returns a reference to field tags of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ams.ref.Append("tags"))
}

// TransparentDataEncryptionKeyVaultKeyId returns a reference to field transparent_data_encryption_key_vault_key_id of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) TransparentDataEncryptionKeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("transparent_data_encryption_key_vault_key_id"))
}

// Version returns a reference to field version of azurerm_mssql_server.
func (ams azurermMssqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ams.ref.Append("version"))
}

func (ams azurermMssqlServerAttributes) AzureadAdministrator() terra.ListValue[AzureadAdministratorAttributes] {
	return terra.ReferenceAsList[AzureadAdministratorAttributes](ams.ref.Append("azuread_administrator"))
}

func (ams azurermMssqlServerAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](ams.ref.Append("identity"))
}

func (ams azurermMssqlServerAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ams.ref.Append("timeouts"))
}

type azurermMssqlServerState struct {
	AdministratorLogin                     string                      `json:"administrator_login"`
	AdministratorLoginPassword             string                      `json:"administrator_login_password"`
	ConnectionPolicy                       string                      `json:"connection_policy"`
	FullyQualifiedDomainName               string                      `json:"fully_qualified_domain_name"`
	Id                                     string                      `json:"id"`
	Location                               string                      `json:"location"`
	MinimumTlsVersion                      string                      `json:"minimum_tls_version"`
	Name                                   string                      `json:"name"`
	OutboundNetworkRestrictionEnabled      bool                        `json:"outbound_network_restriction_enabled"`
	PrimaryUserAssignedIdentityId          string                      `json:"primary_user_assigned_identity_id"`
	PublicNetworkAccessEnabled             bool                        `json:"public_network_access_enabled"`
	ResourceGroupName                      string                      `json:"resource_group_name"`
	RestorableDroppedDatabaseIds           []string                    `json:"restorable_dropped_database_ids"`
	Tags                                   map[string]string           `json:"tags"`
	TransparentDataEncryptionKeyVaultKeyId string                      `json:"transparent_data_encryption_key_vault_key_id"`
	Version                                string                      `json:"version"`
	AzureadAdministrator                   []AzureadAdministratorState `json:"azuread_administrator"`
	Identity                               []IdentityState             `json:"identity"`
	Timeouts                               *TimeoutsState              `json:"timeouts"`
}
