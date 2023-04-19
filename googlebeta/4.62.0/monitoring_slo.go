// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	monitoringslo "github.com/golingon/terraproviders/googlebeta/4.62.0/monitoringslo"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringSlo creates a new instance of [MonitoringSlo].
func NewMonitoringSlo(name string, args MonitoringSloArgs) *MonitoringSlo {
	return &MonitoringSlo{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringSlo)(nil)

// MonitoringSlo represents the Terraform resource google_monitoring_slo.
type MonitoringSlo struct {
	Name      string
	Args      MonitoringSloArgs
	state     *monitoringSloState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringSlo].
func (ms *MonitoringSlo) Type() string {
	return "google_monitoring_slo"
}

// LocalName returns the local name for [MonitoringSlo].
func (ms *MonitoringSlo) LocalName() string {
	return ms.Name
}

// Configuration returns the configuration (args) for [MonitoringSlo].
func (ms *MonitoringSlo) Configuration() interface{} {
	return ms.Args
}

// DependOn is used for other resources to depend on [MonitoringSlo].
func (ms *MonitoringSlo) DependOn() terra.Reference {
	return terra.ReferenceResource(ms)
}

// Dependencies returns the list of resources [MonitoringSlo] depends_on.
func (ms *MonitoringSlo) Dependencies() terra.Dependencies {
	return ms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringSlo].
func (ms *MonitoringSlo) LifecycleManagement() *terra.Lifecycle {
	return ms.Lifecycle
}

// Attributes returns the attributes for [MonitoringSlo].
func (ms *MonitoringSlo) Attributes() monitoringSloAttributes {
	return monitoringSloAttributes{ref: terra.ReferenceResource(ms)}
}

// ImportState imports the given attribute values into [MonitoringSlo]'s state.
func (ms *MonitoringSlo) ImportState(av io.Reader) error {
	ms.state = &monitoringSloState{}
	if err := json.NewDecoder(av).Decode(ms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ms.Type(), ms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringSlo] has state.
func (ms *MonitoringSlo) State() (*monitoringSloState, bool) {
	return ms.state, ms.state != nil
}

// StateMust returns the state for [MonitoringSlo]. Panics if the state is nil.
func (ms *MonitoringSlo) StateMust() *monitoringSloState {
	if ms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ms.Type(), ms.LocalName()))
	}
	return ms.state
}

// MonitoringSloArgs contains the configurations for google_monitoring_slo.
type MonitoringSloArgs struct {
	// CalendarPeriod: string, optional
	CalendarPeriod terra.StringValue `hcl:"calendar_period,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Goal: number, required
	Goal terra.NumberValue `hcl:"goal,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RollingPeriodDays: number, optional
	RollingPeriodDays terra.NumberValue `hcl:"rolling_period_days,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// SloId: string, optional
	SloId terra.StringValue `hcl:"slo_id,attr"`
	// UserLabels: map of string, optional
	UserLabels terra.MapValue[terra.StringValue] `hcl:"user_labels,attr"`
	// BasicSli: optional
	BasicSli *monitoringslo.BasicSli `hcl:"basic_sli,block"`
	// RequestBasedSli: optional
	RequestBasedSli *monitoringslo.RequestBasedSli `hcl:"request_based_sli,block"`
	// Timeouts: optional
	Timeouts *monitoringslo.Timeouts `hcl:"timeouts,block"`
	// WindowsBasedSli: optional
	WindowsBasedSli *monitoringslo.WindowsBasedSli `hcl:"windows_based_sli,block"`
}
type monitoringSloAttributes struct {
	ref terra.Reference
}

// CalendarPeriod returns a reference to field calendar_period of google_monitoring_slo.
func (ms monitoringSloAttributes) CalendarPeriod() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("calendar_period"))
}

// DisplayName returns a reference to field display_name of google_monitoring_slo.
func (ms monitoringSloAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("display_name"))
}

// Goal returns a reference to field goal of google_monitoring_slo.
func (ms monitoringSloAttributes) Goal() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("goal"))
}

// Id returns a reference to field id of google_monitoring_slo.
func (ms monitoringSloAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("id"))
}

// Name returns a reference to field name of google_monitoring_slo.
func (ms monitoringSloAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_slo.
func (ms monitoringSloAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("project"))
}

// RollingPeriodDays returns a reference to field rolling_period_days of google_monitoring_slo.
func (ms monitoringSloAttributes) RollingPeriodDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ms.ref.Append("rolling_period_days"))
}

// Service returns a reference to field service of google_monitoring_slo.
func (ms monitoringSloAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("service"))
}

// SloId returns a reference to field slo_id of google_monitoring_slo.
func (ms monitoringSloAttributes) SloId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("slo_id"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_slo.
func (ms monitoringSloAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ms.ref.Append("user_labels"))
}

func (ms monitoringSloAttributes) BasicSli() terra.ListValue[monitoringslo.BasicSliAttributes] {
	return terra.ReferenceAsList[monitoringslo.BasicSliAttributes](ms.ref.Append("basic_sli"))
}

func (ms monitoringSloAttributes) RequestBasedSli() terra.ListValue[monitoringslo.RequestBasedSliAttributes] {
	return terra.ReferenceAsList[monitoringslo.RequestBasedSliAttributes](ms.ref.Append("request_based_sli"))
}

func (ms monitoringSloAttributes) Timeouts() monitoringslo.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringslo.TimeoutsAttributes](ms.ref.Append("timeouts"))
}

func (ms monitoringSloAttributes) WindowsBasedSli() terra.ListValue[monitoringslo.WindowsBasedSliAttributes] {
	return terra.ReferenceAsList[monitoringslo.WindowsBasedSliAttributes](ms.ref.Append("windows_based_sli"))
}

type monitoringSloState struct {
	CalendarPeriod    string                               `json:"calendar_period"`
	DisplayName       string                               `json:"display_name"`
	Goal              float64                              `json:"goal"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Project           string                               `json:"project"`
	RollingPeriodDays float64                              `json:"rolling_period_days"`
	Service           string                               `json:"service"`
	SloId             string                               `json:"slo_id"`
	UserLabels        map[string]string                    `json:"user_labels"`
	BasicSli          []monitoringslo.BasicSliState        `json:"basic_sli"`
	RequestBasedSli   []monitoringslo.RequestBasedSliState `json:"request_based_sli"`
	Timeouts          *monitoringslo.TimeoutsState         `json:"timeouts"`
	WindowsBasedSli   []monitoringslo.WindowsBasedSliState `json:"windows_based_sli"`
}
