// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlserver "github.com/golingon/terraproviders/azurerm/3.63.0/mssqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlServer creates a new instance of [MssqlServer].
func NewMssqlServer(name string, args MssqlServerArgs) *MssqlServer {
	return &MssqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServer)(nil)

// MssqlServer represents the Terraform resource azurerm_mssql_server.
type MssqlServer struct {
	Name      string
	Args      MssqlServerArgs
	state     *mssqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServer].
func (ms *MssqlServer) Type() string {
	return "azurerm_mssql_server"
}

// LocalName returns the local name for [MssqlServer].
func (ms *MssqlServer) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [MssqlServer].
func (ms *MssqlServer) Configuration() interface{} {
	return ms.Args
}

// DependOn is used for other resources to depend on [MssqlServer].
func (ms *MssqlServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ms)
}

// Dependencies returns the list of resources [MssqlServer] depends_on.
func (ms *MssqlServer) Dependencies() terra.Dependencies {
	return ms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServer].
func (ms *MssqlServer) LifecycleManagement() *terra.Lifecycle {
	return ms.Lifecycle
}

// Attributes returns the attributes for [MssqlServer].
func (ms *MssqlServer) Attributes() mssqlServerAttributes {
	return mssqlServerAttributes{ref: terra.ReferenceResource(ms)}
}

// ImportState imports the given attribute values into [MssqlServer]'s state.
func (ms *MssqlServer) ImportState(av io.Reader) error {
	ms.state = &mssqlServerState{}
	if err := json.NewDecoder(av).Decode(ms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ms.Type(), ms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServer] has state.
func (ms *MssqlServer) State() (*mssqlServerState, bool) {
	return ms.state, ms.state != nil
}

// StateMust returns the state for [MssqlServer]. Panics if the state is nil.
func (ms *MssqlServer) StateMust() *mssqlServerState {
	if ms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ms.Type(), ms.LocalName()))
	}
	return ms.state
}

// MssqlServerArgs contains the configurations for azurerm_mssql_server.
type MssqlServerArgs struct {
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
	AzureadAdministrator *mssqlserver.AzureadAdministrator `hcl:"azuread_administrator,block"`
	// Identity: optional
	Identity *mssqlserver.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *mssqlserver.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_mssql_server.
func (ms mssqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_mssql_server.
func (ms mssqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("administrator_login_password"))
}

// ConnectionPolicy returns a reference to field connection_policy of azurerm_mssql_server.
func (ms mssqlServerAttributes) ConnectionPolicy() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("connection_policy"))
}

// FullyQualifiedDomainName returns a reference to field fully_qualified_domain_name of azurerm_mssql_server.
func (ms mssqlServerAttributes) FullyQualifiedDomainName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("fully_qualified_domain_name"))
}

// Id returns a reference to field id of azurerm_mssql_server.
func (ms mssqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_server.
func (ms mssqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("location"))
}

// MinimumTlsVersion returns a reference to field minimum_tls_version of azurerm_mssql_server.
func (ms mssqlServerAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("minimum_tls_version"))
}

// Name returns a reference to field name of azurerm_mssql_server.
func (ms mssqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// OutboundNetworkRestrictionEnabled returns a reference to field outbound_network_restriction_enabled of azurerm_mssql_server.
func (ms mssqlServerAttributes) OutboundNetworkRestrictionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("outbound_network_restriction_enabled"))
}

// PrimaryUserAssignedIdentityId returns a reference to field primary_user_assigned_identity_id of azurerm_mssql_server.
func (ms mssqlServerAttributes) PrimaryUserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("primary_user_assigned_identity_id"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_mssql_server.
func (ms mssqlServerAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ms.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_server.
func (ms mssqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("resource_group_name"))
}

// RestorableDroppedDatabaseIds returns a reference to field restorable_dropped_database_ids of azurerm_mssql_server.
func (ms mssqlServerAttributes) RestorableDroppedDatabaseIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ms.ref.Append("restorable_dropped_database_ids"))
}

// Tags returns a reference to field tags of azurerm_mssql_server.
func (ms mssqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("tags"))
}

// TransparentDataEncryptionKeyVaultKeyId returns a reference to field transparent_data_encryption_key_vault_key_id of azurerm_mssql_server.
func (ms mssqlServerAttributes) TransparentDataEncryptionKeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("transparent_data_encryption_key_vault_key_id"))
}

// Version returns a reference to field version of azurerm_mssql_server.
func (ms mssqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("version"))
}

func (ms mssqlServerAttributes) AzureadAdministrator() terra.ListValue[mssqlserver.AzureadAdministratorAttributes] {
	return terra.ReferenceAsList[mssqlserver.AzureadAdministratorAttributes](ms.ref.Append("azuread_administrator"))
}

func (ms mssqlServerAttributes) Identity() terra.ListValue[mssqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[mssqlserver.IdentityAttributes](ms.ref.Append("identity"))
}

func (ms mssqlServerAttributes) Timeouts() mssqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlserver.TimeoutsAttributes](ms.ref.Append("timeouts"))
}

type mssqlServerState struct {
	AdministratorLogin                     string                                  `json:"administrator_login"`
	AdministratorLoginPassword             string                                  `json:"administrator_login_password"`
	ConnectionPolicy                       string                                  `json:"connection_policy"`
	FullyQualifiedDomainName               string                                  `json:"fully_qualified_domain_name"`
	Id                                     string                                  `json:"id"`
	Location                               string                                  `json:"location"`
	MinimumTlsVersion                      string                                  `json:"minimum_tls_version"`
	Name                                   string                                  `json:"name"`
	OutboundNetworkRestrictionEnabled      bool                                    `json:"outbound_network_restriction_enabled"`
	PrimaryUserAssignedIdentityId          string                                  `json:"primary_user_assigned_identity_id"`
	PublicNetworkAccessEnabled             bool                                    `json:"public_network_access_enabled"`
	ResourceGroupName                      string                                  `json:"resource_group_name"`
	RestorableDroppedDatabaseIds           []string                                `json:"restorable_dropped_database_ids"`
	Tags                                   map[string]string                       `json:"tags"`
	TransparentDataEncryptionKeyVaultKeyId string                                  `json:"transparent_data_encryption_key_vault_key_id"`
	Version                                string                                  `json:"version"`
	AzureadAdministrator                   []mssqlserver.AzureadAdministratorState `json:"azuread_administrator"`
	Identity                               []mssqlserver.IdentityState             `json:"identity"`
	Timeouts                               *mssqlserver.TimeoutsState              `json:"timeouts"`
}
