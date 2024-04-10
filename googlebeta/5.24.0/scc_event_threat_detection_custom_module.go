// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	scceventthreatdetectioncustommodule "github.com/golingon/terraproviders/googlebeta/5.24.0/scceventthreatdetectioncustommodule"
	"io"
)

// NewSccEventThreatDetectionCustomModule creates a new instance of [SccEventThreatDetectionCustomModule].
func NewSccEventThreatDetectionCustomModule(name string, args SccEventThreatDetectionCustomModuleArgs) *SccEventThreatDetectionCustomModule {
	return &SccEventThreatDetectionCustomModule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccEventThreatDetectionCustomModule)(nil)

// SccEventThreatDetectionCustomModule represents the Terraform resource google_scc_event_threat_detection_custom_module.
type SccEventThreatDetectionCustomModule struct {
	Name      string
	Args      SccEventThreatDetectionCustomModuleArgs
	state     *sccEventThreatDetectionCustomModuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) Type() string {
	return "google_scc_event_threat_detection_custom_module"
}

// LocalName returns the local name for [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) LocalName() string {
	return setdcm.Name
}

// Configuration returns the configuration (args) for [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) Configuration() interface{} {
	return setdcm.Args
}

// DependOn is used for other resources to depend on [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) DependOn() terra.Reference {
	return terra.ReferenceResource(setdcm)
}

// Dependencies returns the list of resources [SccEventThreatDetectionCustomModule] depends_on.
func (setdcm *SccEventThreatDetectionCustomModule) Dependencies() terra.Dependencies {
	return setdcm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) LifecycleManagement() *terra.Lifecycle {
	return setdcm.Lifecycle
}

// Attributes returns the attributes for [SccEventThreatDetectionCustomModule].
func (setdcm *SccEventThreatDetectionCustomModule) Attributes() sccEventThreatDetectionCustomModuleAttributes {
	return sccEventThreatDetectionCustomModuleAttributes{ref: terra.ReferenceResource(setdcm)}
}

// ImportState imports the given attribute values into [SccEventThreatDetectionCustomModule]'s state.
func (setdcm *SccEventThreatDetectionCustomModule) ImportState(av io.Reader) error {
	setdcm.state = &sccEventThreatDetectionCustomModuleState{}
	if err := json.NewDecoder(av).Decode(setdcm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", setdcm.Type(), setdcm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccEventThreatDetectionCustomModule] has state.
func (setdcm *SccEventThreatDetectionCustomModule) State() (*sccEventThreatDetectionCustomModuleState, bool) {
	return setdcm.state, setdcm.state != nil
}

// StateMust returns the state for [SccEventThreatDetectionCustomModule]. Panics if the state is nil.
func (setdcm *SccEventThreatDetectionCustomModule) StateMust() *sccEventThreatDetectionCustomModuleState {
	if setdcm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", setdcm.Type(), setdcm.LocalName()))
	}
	return setdcm.state
}

// SccEventThreatDetectionCustomModuleArgs contains the configurations for google_scc_event_threat_detection_custom_module.
type SccEventThreatDetectionCustomModuleArgs struct {
	// Config: string, required
	Config terra.StringValue `hcl:"config,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EnablementState: string, required
	EnablementState terra.StringValue `hcl:"enablement_state,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *scceventthreatdetectioncustommodule.Timeouts `hcl:"timeouts,block"`
}
type sccEventThreatDetectionCustomModuleAttributes struct {
	ref terra.Reference
}

// Config returns a reference to field config of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) Config() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("config"))
}

// DisplayName returns a reference to field display_name of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("display_name"))
}

// EnablementState returns a reference to field enablement_state of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) EnablementState() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("enablement_state"))
}

// Id returns a reference to field id of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("id"))
}

// LastEditor returns a reference to field last_editor of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) LastEditor() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("last_editor"))
}

// Name returns a reference to field name of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("name"))
}

// Organization returns a reference to field organization of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("organization"))
}

// Type returns a reference to field type of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_scc_event_threat_detection_custom_module.
func (setdcm sccEventThreatDetectionCustomModuleAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(setdcm.ref.Append("update_time"))
}

func (setdcm sccEventThreatDetectionCustomModuleAttributes) Timeouts() scceventthreatdetectioncustommodule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[scceventthreatdetectioncustommodule.TimeoutsAttributes](setdcm.ref.Append("timeouts"))
}

type sccEventThreatDetectionCustomModuleState struct {
	Config          string                                             `json:"config"`
	DisplayName     string                                             `json:"display_name"`
	EnablementState string                                             `json:"enablement_state"`
	Id              string                                             `json:"id"`
	LastEditor      string                                             `json:"last_editor"`
	Name            string                                             `json:"name"`
	Organization    string                                             `json:"organization"`
	Type            string                                             `json:"type"`
	UpdateTime      string                                             `json:"update_time"`
	Timeouts        *scceventthreatdetectioncustommodule.TimeoutsState `json:"timeouts"`
}
