// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstation "github.com/golingon/terraproviders/googlebeta/4.73.1/workstationsworkstation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstation creates a new instance of [WorkstationsWorkstation].
func NewWorkstationsWorkstation(name string, args WorkstationsWorkstationArgs) *WorkstationsWorkstation {
	return &WorkstationsWorkstation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstation)(nil)

// WorkstationsWorkstation represents the Terraform resource google_workstations_workstation.
type WorkstationsWorkstation struct {
	Name      string
	Args      WorkstationsWorkstationArgs
	state     *workstationsWorkstationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) Type() string {
	return "google_workstations_workstation"
}

// LocalName returns the local name for [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) LocalName() string {
	return ww.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) Configuration() interface{} {
	return ww.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) DependOn() terra.Reference {
	return terra.ReferenceResource(ww)
}

// Dependencies returns the list of resources [WorkstationsWorkstation] depends_on.
func (ww *WorkstationsWorkstation) Dependencies() terra.Dependencies {
	return ww.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) LifecycleManagement() *terra.Lifecycle {
	return ww.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstation].
func (ww *WorkstationsWorkstation) Attributes() workstationsWorkstationAttributes {
	return workstationsWorkstationAttributes{ref: terra.ReferenceResource(ww)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstation]'s state.
func (ww *WorkstationsWorkstation) ImportState(av io.Reader) error {
	ww.state = &workstationsWorkstationState{}
	if err := json.NewDecoder(av).Decode(ww.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ww.Type(), ww.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstation] has state.
func (ww *WorkstationsWorkstation) State() (*workstationsWorkstationState, bool) {
	return ww.state, ww.state != nil
}

// StateMust returns the state for [WorkstationsWorkstation]. Panics if the state is nil.
func (ww *WorkstationsWorkstation) StateMust() *workstationsWorkstationState {
	if ww.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ww.Type(), ww.LocalName()))
	}
	return ww.state
}

// WorkstationsWorkstationArgs contains the configurations for google_workstations_workstation.
type WorkstationsWorkstationArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// WorkstationConfigId: string, required
	WorkstationConfigId terra.StringValue `hcl:"workstation_config_id,attr" validate:"required"`
	// WorkstationId: string, required
	WorkstationId terra.StringValue `hcl:"workstation_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *workstationsworkstation.Timeouts `hcl:"timeouts,block"`
}
type workstationsWorkstationAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ww.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("display_name"))
}

// Host returns a reference to field host of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("host"))
}

// Id returns a reference to field id of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("id"))
}

// Labels returns a reference to field labels of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ww.ref.Append("labels"))
}

// Location returns a reference to field location of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("location"))
}

// Name returns a reference to field name of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("name"))
}

// Project returns a reference to field project of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("project"))
}

// State returns a reference to field state of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("state"))
}

// Uid returns a reference to field uid of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("uid"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("workstation_config_id"))
}

// WorkstationId returns a reference to field workstation_id of google_workstations_workstation.
func (ww workstationsWorkstationAttributes) WorkstationId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("workstation_id"))
}

func (ww workstationsWorkstationAttributes) Timeouts() workstationsworkstation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workstationsworkstation.TimeoutsAttributes](ww.ref.Append("timeouts"))
}

type workstationsWorkstationState struct {
	Annotations          map[string]string                      `json:"annotations"`
	CreateTime           string                                 `json:"create_time"`
	DisplayName          string                                 `json:"display_name"`
	Host                 string                                 `json:"host"`
	Id                   string                                 `json:"id"`
	Labels               map[string]string                      `json:"labels"`
	Location             string                                 `json:"location"`
	Name                 string                                 `json:"name"`
	Project              string                                 `json:"project"`
	State                string                                 `json:"state"`
	Uid                  string                                 `json:"uid"`
	WorkstationClusterId string                                 `json:"workstation_cluster_id"`
	WorkstationConfigId  string                                 `json:"workstation_config_id"`
	WorkstationId        string                                 `json:"workstation_id"`
	Timeouts             *workstationsworkstation.TimeoutsState `json:"timeouts"`
}
