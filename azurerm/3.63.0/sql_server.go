// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sqlserver "github.com/golingon/terraproviders/azurerm/3.63.0/sqlserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqlServer creates a new instance of [SqlServer].
func NewSqlServer(name string, args SqlServerArgs) *SqlServer {
	return &SqlServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqlServer)(nil)

// SqlServer represents the Terraform resource azurerm_sql_server.
type SqlServer struct {
	Name      string
	Args      SqlServerArgs
	state     *sqlServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqlServer].
func (ss *SqlServer) Type() string {
	return "azurerm_sql_server"
}

// LocalName returns the local name for [SqlServer].
func (ss *SqlServer) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SqlServer].
func (ss *SqlServer) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SqlServer].
func (ss *SqlServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SqlServer] depends_on.
func (ss *SqlServer) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqlServer].
func (ss *SqlServer) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SqlServer].
func (ss *SqlServer) Attributes() sqlServerAttributes {
	return sqlServerAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SqlServer]'s state.
func (ss *SqlServer) ImportState(av io.Reader) error {
	ss.state = &sqlServerState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqlServer] has state.
func (ss *SqlServer) State() (*sqlServerState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SqlServer]. Panics if the state is nil.
func (ss *SqlServer) StateMust() *sqlServerState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SqlServerArgs contains the configurations for azurerm_sql_server.
type SqlServerArgs struct {
	// AdministratorLogin: string, required
	AdministratorLogin terra.StringValue `hcl:"administrator_login,attr" validate:"required"`
	// AdministratorLoginPassword: string, required
	AdministratorLoginPassword terra.StringValue `hcl:"administrator_login_password,attr" validate:"required"`
	// ConnectionPolicy: string, optional
	ConnectionPolicy terra.StringValue `hcl:"connection_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Identity: optional
	Identity *sqlserver.Identity `hcl:"identity,block"`
	// ThreatDetectionPolicy: optional
	ThreatDetectionPolicy *sqlserver.ThreatDetectionPolicy `hcl:"threat_detection_policy,block"`
	// Timeouts: optional
	Timeouts *sqlserver.Timeouts `hcl:"timeouts,block"`
}
type sqlServerAttributes struct {
	ref terra.Reference
}

// AdministratorLogin returns a reference to field administrator_login of azurerm_sql_server.
func (ss sqlServerAttributes) AdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("administrator_login"))
}

// AdministratorLoginPassword returns a reference to field administrator_login_password of azurerm_sql_server.
func (ss sqlServerAttributes) AdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("administrator_login_password"))
}

// ConnectionPolicy returns a reference to field connection_policy of azurerm_sql_server.
func (ss sqlServerAttributes) ConnectionPolicy() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("connection_policy"))
}

// FullyQualifiedDomainName returns a reference to field fully_qualified_domain_name of azurerm_sql_server.
func (ss sqlServerAttributes) FullyQualifiedDomainName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("fully_qualified_domain_name"))
}

// Id returns a reference to field id of azurerm_sql_server.
func (ss sqlServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_sql_server.
func (ss sqlServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_sql_server.
func (ss sqlServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sql_server.
func (ss sqlServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_sql_server.
func (ss sqlServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_sql_server.
func (ss sqlServerAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("version"))
}

func (ss sqlServerAttributes) Identity() terra.ListValue[sqlserver.IdentityAttributes] {
	return terra.ReferenceAsList[sqlserver.IdentityAttributes](ss.ref.Append("identity"))
}

func (ss sqlServerAttributes) ThreatDetectionPolicy() terra.ListValue[sqlserver.ThreatDetectionPolicyAttributes] {
	return terra.ReferenceAsList[sqlserver.ThreatDetectionPolicyAttributes](ss.ref.Append("threat_detection_policy"))
}

func (ss sqlServerAttributes) Timeouts() sqlserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sqlserver.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type sqlServerState struct {
	AdministratorLogin         string                                 `json:"administrator_login"`
	AdministratorLoginPassword string                                 `json:"administrator_login_password"`
	ConnectionPolicy           string                                 `json:"connection_policy"`
	FullyQualifiedDomainName   string                                 `json:"fully_qualified_domain_name"`
	Id                         string                                 `json:"id"`
	Location                   string                                 `json:"location"`
	Name                       string                                 `json:"name"`
	ResourceGroupName          string                                 `json:"resource_group_name"`
	Tags                       map[string]string                      `json:"tags"`
	Version                    string                                 `json:"version"`
	Identity                   []sqlserver.IdentityState              `json:"identity"`
	ThreatDetectionPolicy      []sqlserver.ThreatDetectionPolicyState `json:"threat_detection_policy"`
	Timeouts                   *sqlserver.TimeoutsState               `json:"timeouts"`
}