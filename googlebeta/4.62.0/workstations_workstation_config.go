// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstationconfig "github.com/golingon/terraproviders/googlebeta/4.62.0/workstationsworkstationconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstationConfig creates a new instance of [WorkstationsWorkstationConfig].
func NewWorkstationsWorkstationConfig(name string, args WorkstationsWorkstationConfigArgs) *WorkstationsWorkstationConfig {
	return &WorkstationsWorkstationConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstationConfig)(nil)

// WorkstationsWorkstationConfig represents the Terraform resource google_workstations_workstation_config.
type WorkstationsWorkstationConfig struct {
	Name      string
	Args      WorkstationsWorkstationConfigArgs
	state     *workstationsWorkstationConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) Type() string {
	return "google_workstations_workstation_config"
}

// LocalName returns the local name for [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) LocalName() string {
	return wwc.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) Configuration() interface{} {
	return wwc.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(wwc)
}

// Dependencies returns the list of resources [WorkstationsWorkstationConfig] depends_on.
func (wwc *WorkstationsWorkstationConfig) Dependencies() terra.Dependencies {
	return wwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) LifecycleManagement() *terra.Lifecycle {
	return wwc.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstationConfig].
func (wwc *WorkstationsWorkstationConfig) Attributes() workstationsWorkstationConfigAttributes {
	return workstationsWorkstationConfigAttributes{ref: terra.ReferenceResource(wwc)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstationConfig]'s state.
func (wwc *WorkstationsWorkstationConfig) ImportState(av io.Reader) error {
	wwc.state = &workstationsWorkstationConfigState{}
	if err := json.NewDecoder(av).Decode(wwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwc.Type(), wwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstationConfig] has state.
func (wwc *WorkstationsWorkstationConfig) State() (*workstationsWorkstationConfigState, bool) {
	return wwc.state, wwc.state != nil
}

// StateMust returns the state for [WorkstationsWorkstationConfig]. Panics if the state is nil.
func (wwc *WorkstationsWorkstationConfig) StateMust() *workstationsWorkstationConfigState {
	if wwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwc.Type(), wwc.LocalName()))
	}
	return wwc.state
}

// WorkstationsWorkstationConfigArgs contains the configurations for google_workstations_workstation_config.
type WorkstationsWorkstationConfigArgs struct {
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
	// Conditions: min=0
	Conditions []workstationsworkstationconfig.Conditions `hcl:"conditions,block" validate:"min=0"`
	// Container: optional
	Container *workstationsworkstationconfig.Container `hcl:"container,block"`
	// EncryptionKey: optional
	EncryptionKey *workstationsworkstationconfig.EncryptionKey `hcl:"encryption_key,block"`
	// Host: optional
	Host *workstationsworkstationconfig.Host `hcl:"host,block"`
	// PersistentDirectories: min=0
	PersistentDirectories []workstationsworkstationconfig.PersistentDirectories `hcl:"persistent_directories,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *workstationsworkstationconfig.Timeouts `hcl:"timeouts,block"`
}
type workstationsWorkstationConfigAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwc.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("create_time"))
}

// Degraded returns a reference to field degraded of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Degraded() terra.BoolValue {
	return terra.ReferenceAsBool(wwc.ref.Append("degraded"))
}

// DisplayName returns a reference to field display_name of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwc.ref.Append("labels"))
}

// Location returns a reference to field location of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("location"))
}

// Name returns a reference to field name of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("name"))
}

// Project returns a reference to field project of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("project"))
}

// Uid returns a reference to field uid of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("uid"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_config.
func (wwc workstationsWorkstationConfigAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("workstation_config_id"))
}

func (wwc workstationsWorkstationConfigAttributes) Conditions() terra.ListValue[workstationsworkstationconfig.ConditionsAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfig.ConditionsAttributes](wwc.ref.Append("conditions"))
}

func (wwc workstationsWorkstationConfigAttributes) Container() terra.ListValue[workstationsworkstationconfig.ContainerAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfig.ContainerAttributes](wwc.ref.Append("container"))
}

func (wwc workstationsWorkstationConfigAttributes) EncryptionKey() terra.ListValue[workstationsworkstationconfig.EncryptionKeyAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfig.EncryptionKeyAttributes](wwc.ref.Append("encryption_key"))
}

func (wwc workstationsWorkstationConfigAttributes) Host() terra.ListValue[workstationsworkstationconfig.HostAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfig.HostAttributes](wwc.ref.Append("host"))
}

func (wwc workstationsWorkstationConfigAttributes) PersistentDirectories() terra.ListValue[workstationsworkstationconfig.PersistentDirectoriesAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfig.PersistentDirectoriesAttributes](wwc.ref.Append("persistent_directories"))
}

func (wwc workstationsWorkstationConfigAttributes) Timeouts() workstationsworkstationconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workstationsworkstationconfig.TimeoutsAttributes](wwc.ref.Append("timeouts"))
}

type workstationsWorkstationConfigState struct {
	Annotations           map[string]string                                          `json:"annotations"`
	CreateTime            string                                                     `json:"create_time"`
	Degraded              bool                                                       `json:"degraded"`
	DisplayName           string                                                     `json:"display_name"`
	Etag                  string                                                     `json:"etag"`
	Id                    string                                                     `json:"id"`
	Labels                map[string]string                                          `json:"labels"`
	Location              string                                                     `json:"location"`
	Name                  string                                                     `json:"name"`
	Project               string                                                     `json:"project"`
	Uid                   string                                                     `json:"uid"`
	WorkstationClusterId  string                                                     `json:"workstation_cluster_id"`
	WorkstationConfigId   string                                                     `json:"workstation_config_id"`
	Conditions            []workstationsworkstationconfig.ConditionsState            `json:"conditions"`
	Container             []workstationsworkstationconfig.ContainerState             `json:"container"`
	EncryptionKey         []workstationsworkstationconfig.EncryptionKeyState         `json:"encryption_key"`
	Host                  []workstationsworkstationconfig.HostState                  `json:"host"`
	PersistentDirectories []workstationsworkstationconfig.PersistentDirectoriesState `json:"persistent_directories"`
	Timeouts              *workstationsworkstationconfig.TimeoutsState               `json:"timeouts"`
}
