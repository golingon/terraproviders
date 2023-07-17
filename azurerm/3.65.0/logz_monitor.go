// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logzmonitor "github.com/golingon/terraproviders/azurerm/3.65.0/logzmonitor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogzMonitor creates a new instance of [LogzMonitor].
func NewLogzMonitor(name string, args LogzMonitorArgs) *LogzMonitor {
	return &LogzMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogzMonitor)(nil)

// LogzMonitor represents the Terraform resource azurerm_logz_monitor.
type LogzMonitor struct {
	Name      string
	Args      LogzMonitorArgs
	state     *logzMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogzMonitor].
func (lm *LogzMonitor) Type() string {
	return "azurerm_logz_monitor"
}

// LocalName returns the local name for [LogzMonitor].
func (lm *LogzMonitor) LocalName() string {
	return lm.Name
}

// Configuration returns the configuration (args) for [LogzMonitor].
func (lm *LogzMonitor) Configuration() interface{} {
	return lm.Args
}

// DependOn is used for other resources to depend on [LogzMonitor].
func (lm *LogzMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(lm)
}

// Dependencies returns the list of resources [LogzMonitor] depends_on.
func (lm *LogzMonitor) Dependencies() terra.Dependencies {
	return lm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogzMonitor].
func (lm *LogzMonitor) LifecycleManagement() *terra.Lifecycle {
	return lm.Lifecycle
}

// Attributes returns the attributes for [LogzMonitor].
func (lm *LogzMonitor) Attributes() logzMonitorAttributes {
	return logzMonitorAttributes{ref: terra.ReferenceResource(lm)}
}

// ImportState imports the given attribute values into [LogzMonitor]'s state.
func (lm *LogzMonitor) ImportState(av io.Reader) error {
	lm.state = &logzMonitorState{}
	if err := json.NewDecoder(av).Decode(lm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lm.Type(), lm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogzMonitor] has state.
func (lm *LogzMonitor) State() (*logzMonitorState, bool) {
	return lm.state, lm.state != nil
}

// StateMust returns the state for [LogzMonitor]. Panics if the state is nil.
func (lm *LogzMonitor) StateMust() *logzMonitorState {
	if lm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lm.Type(), lm.LocalName()))
	}
	return lm.state
}

// LogzMonitorArgs contains the configurations for azurerm_logz_monitor.
type LogzMonitorArgs struct {
	// CompanyName: string, optional
	CompanyName terra.StringValue `hcl:"company_name,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EnterpriseAppId: string, optional
	EnterpriseAppId terra.StringValue `hcl:"enterprise_app_id,attr"`
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
	// Plan: required
	Plan *logzmonitor.Plan `hcl:"plan,block" validate:"required"`
	// Timeouts: optional
	Timeouts *logzmonitor.Timeouts `hcl:"timeouts,block"`
	// User: required
	User *logzmonitor.User `hcl:"user,block" validate:"required"`
}
type logzMonitorAttributes struct {
	ref terra.Reference
}

// CompanyName returns a reference to field company_name of azurerm_logz_monitor.
func (lm logzMonitorAttributes) CompanyName() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("company_name"))
}

// Enabled returns a reference to field enabled of azurerm_logz_monitor.
func (lm logzMonitorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lm.ref.Append("enabled"))
}

// EnterpriseAppId returns a reference to field enterprise_app_id of azurerm_logz_monitor.
func (lm logzMonitorAttributes) EnterpriseAppId() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("enterprise_app_id"))
}

// Id returns a reference to field id of azurerm_logz_monitor.
func (lm logzMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_logz_monitor.
func (lm logzMonitorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("location"))
}

// LogzOrganizationId returns a reference to field logz_organization_id of azurerm_logz_monitor.
func (lm logzMonitorAttributes) LogzOrganizationId() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("logz_organization_id"))
}

// Name returns a reference to field name of azurerm_logz_monitor.
func (lm logzMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logz_monitor.
func (lm logzMonitorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("resource_group_name"))
}

// SingleSignOnUrl returns a reference to field single_sign_on_url of azurerm_logz_monitor.
func (lm logzMonitorAttributes) SingleSignOnUrl() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("single_sign_on_url"))
}

// Tags returns a reference to field tags of azurerm_logz_monitor.
func (lm logzMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lm.ref.Append("tags"))
}

func (lm logzMonitorAttributes) Plan() terra.ListValue[logzmonitor.PlanAttributes] {
	return terra.ReferenceAsList[logzmonitor.PlanAttributes](lm.ref.Append("plan"))
}

func (lm logzMonitorAttributes) Timeouts() logzmonitor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logzmonitor.TimeoutsAttributes](lm.ref.Append("timeouts"))
}

func (lm logzMonitorAttributes) User() terra.ListValue[logzmonitor.UserAttributes] {
	return terra.ReferenceAsList[logzmonitor.UserAttributes](lm.ref.Append("user"))
}

type logzMonitorState struct {
	CompanyName        string                     `json:"company_name"`
	Enabled            bool                       `json:"enabled"`
	EnterpriseAppId    string                     `json:"enterprise_app_id"`
	Id                 string                     `json:"id"`
	Location           string                     `json:"location"`
	LogzOrganizationId string                     `json:"logz_organization_id"`
	Name               string                     `json:"name"`
	ResourceGroupName  string                     `json:"resource_group_name"`
	SingleSignOnUrl    string                     `json:"single_sign_on_url"`
	Tags               map[string]string          `json:"tags"`
	Plan               []logzmonitor.PlanState    `json:"plan"`
	Timeouts           *logzmonitor.TimeoutsState `json:"timeouts"`
	User               []logzmonitor.UserState    `json:"user"`
}
