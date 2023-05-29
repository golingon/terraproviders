// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	analysisservicesserver "github.com/golingon/terraproviders/azurerm/3.55.0/analysisservicesserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAnalysisServicesServer creates a new instance of [AnalysisServicesServer].
func NewAnalysisServicesServer(name string, args AnalysisServicesServerArgs) *AnalysisServicesServer {
	return &AnalysisServicesServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AnalysisServicesServer)(nil)

// AnalysisServicesServer represents the Terraform resource azurerm_analysis_services_server.
type AnalysisServicesServer struct {
	Name      string
	Args      AnalysisServicesServerArgs
	state     *analysisServicesServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AnalysisServicesServer].
func (ass *AnalysisServicesServer) Type() string {
	return "azurerm_analysis_services_server"
}

// LocalName returns the local name for [AnalysisServicesServer].
func (ass *AnalysisServicesServer) LocalName() string {
	return ass.Name
}

// Configuration returns the configuration (args) for [AnalysisServicesServer].
func (ass *AnalysisServicesServer) Configuration() interface{} {
	return ass.Args
}

// DependOn is used for other resources to depend on [AnalysisServicesServer].
func (ass *AnalysisServicesServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ass)
}

// Dependencies returns the list of resources [AnalysisServicesServer] depends_on.
func (ass *AnalysisServicesServer) Dependencies() terra.Dependencies {
	return ass.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AnalysisServicesServer].
func (ass *AnalysisServicesServer) LifecycleManagement() *terra.Lifecycle {
	return ass.Lifecycle
}

// Attributes returns the attributes for [AnalysisServicesServer].
func (ass *AnalysisServicesServer) Attributes() analysisServicesServerAttributes {
	return analysisServicesServerAttributes{ref: terra.ReferenceResource(ass)}
}

// ImportState imports the given attribute values into [AnalysisServicesServer]'s state.
func (ass *AnalysisServicesServer) ImportState(av io.Reader) error {
	ass.state = &analysisServicesServerState{}
	if err := json.NewDecoder(av).Decode(ass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ass.Type(), ass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AnalysisServicesServer] has state.
func (ass *AnalysisServicesServer) State() (*analysisServicesServerState, bool) {
	return ass.state, ass.state != nil
}

// StateMust returns the state for [AnalysisServicesServer]. Panics if the state is nil.
func (ass *AnalysisServicesServer) StateMust() *analysisServicesServerState {
	if ass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ass.Type(), ass.LocalName()))
	}
	return ass.state
}

// AnalysisServicesServerArgs contains the configurations for azurerm_analysis_services_server.
type AnalysisServicesServerArgs struct {
	// AdminUsers: set of string, optional
	AdminUsers terra.SetValue[terra.StringValue] `hcl:"admin_users,attr"`
	// BackupBlobContainerUri: string, optional
	BackupBlobContainerUri terra.StringValue `hcl:"backup_blob_container_uri,attr"`
	// EnablePowerBiService: bool, optional
	EnablePowerBiService terra.BoolValue `hcl:"enable_power_bi_service,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// QuerypoolConnectionMode: string, optional
	QuerypoolConnectionMode terra.StringValue `hcl:"querypool_connection_mode,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ipv4FirewallRule: min=0
	Ipv4FirewallRule []analysisservicesserver.Ipv4FirewallRule `hcl:"ipv4_firewall_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *analysisservicesserver.Timeouts `hcl:"timeouts,block"`
}
type analysisServicesServerAttributes struct {
	ref terra.Reference
}

// AdminUsers returns a reference to field admin_users of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) AdminUsers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ass.ref.Append("admin_users"))
}

// BackupBlobContainerUri returns a reference to field backup_blob_container_uri of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) BackupBlobContainerUri() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("backup_blob_container_uri"))
}

// EnablePowerBiService returns a reference to field enable_power_bi_service of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) EnablePowerBiService() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("enable_power_bi_service"))
}

// Id returns a reference to field id of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name"))
}

// QuerypoolConnectionMode returns a reference to field querypool_connection_mode of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) QuerypoolConnectionMode() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("querypool_connection_mode"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("resource_group_name"))
}

// ServerFullName returns a reference to field server_full_name of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) ServerFullName() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("server_full_name"))
}

// Sku returns a reference to field sku of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_analysis_services_server.
func (ass analysisServicesServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags"))
}

func (ass analysisServicesServerAttributes) Ipv4FirewallRule() terra.SetValue[analysisservicesserver.Ipv4FirewallRuleAttributes] {
	return terra.ReferenceAsSet[analysisservicesserver.Ipv4FirewallRuleAttributes](ass.ref.Append("ipv4_firewall_rule"))
}

func (ass analysisServicesServerAttributes) Timeouts() analysisservicesserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[analysisservicesserver.TimeoutsAttributes](ass.ref.Append("timeouts"))
}

type analysisServicesServerState struct {
	AdminUsers              []string                                       `json:"admin_users"`
	BackupBlobContainerUri  string                                         `json:"backup_blob_container_uri"`
	EnablePowerBiService    bool                                           `json:"enable_power_bi_service"`
	Id                      string                                         `json:"id"`
	Location                string                                         `json:"location"`
	Name                    string                                         `json:"name"`
	QuerypoolConnectionMode string                                         `json:"querypool_connection_mode"`
	ResourceGroupName       string                                         `json:"resource_group_name"`
	ServerFullName          string                                         `json:"server_full_name"`
	Sku                     string                                         `json:"sku"`
	Tags                    map[string]string                              `json:"tags"`
	Ipv4FirewallRule        []analysisservicesserver.Ipv4FirewallRuleState `json:"ipv4_firewall_rule"`
	Timeouts                *analysisservicesserver.TimeoutsState          `json:"timeouts"`
}