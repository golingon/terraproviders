// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apigeetargetserver "github.com/golingon/terraproviders/googlebeta/5.24.0/apigeetargetserver"
	"io"
)

// NewApigeeTargetServer creates a new instance of [ApigeeTargetServer].
func NewApigeeTargetServer(name string, args ApigeeTargetServerArgs) *ApigeeTargetServer {
	return &ApigeeTargetServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeTargetServer)(nil)

// ApigeeTargetServer represents the Terraform resource google_apigee_target_server.
type ApigeeTargetServer struct {
	Name      string
	Args      ApigeeTargetServerArgs
	state     *apigeeTargetServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeTargetServer].
func (ats *ApigeeTargetServer) Type() string {
	return "google_apigee_target_server"
}

// LocalName returns the local name for [ApigeeTargetServer].
func (ats *ApigeeTargetServer) LocalName() string {
	return ats.Name
}

// Configuration returns the configuration (args) for [ApigeeTargetServer].
func (ats *ApigeeTargetServer) Configuration() interface{} {
	return ats.Args
}

// DependOn is used for other resources to depend on [ApigeeTargetServer].
func (ats *ApigeeTargetServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ats)
}

// Dependencies returns the list of resources [ApigeeTargetServer] depends_on.
func (ats *ApigeeTargetServer) Dependencies() terra.Dependencies {
	return ats.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeTargetServer].
func (ats *ApigeeTargetServer) LifecycleManagement() *terra.Lifecycle {
	return ats.Lifecycle
}

// Attributes returns the attributes for [ApigeeTargetServer].
func (ats *ApigeeTargetServer) Attributes() apigeeTargetServerAttributes {
	return apigeeTargetServerAttributes{ref: terra.ReferenceResource(ats)}
}

// ImportState imports the given attribute values into [ApigeeTargetServer]'s state.
func (ats *ApigeeTargetServer) ImportState(av io.Reader) error {
	ats.state = &apigeeTargetServerState{}
	if err := json.NewDecoder(av).Decode(ats.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ats.Type(), ats.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeTargetServer] has state.
func (ats *ApigeeTargetServer) State() (*apigeeTargetServerState, bool) {
	return ats.state, ats.state != nil
}

// StateMust returns the state for [ApigeeTargetServer]. Panics if the state is nil.
func (ats *ApigeeTargetServer) StateMust() *apigeeTargetServerState {
	if ats.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ats.Type(), ats.LocalName()))
	}
	return ats.state
}

// ApigeeTargetServerArgs contains the configurations for google_apigee_target_server.
type ApigeeTargetServerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsEnabled: bool, optional
	IsEnabled terra.BoolValue `hcl:"is_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// SSlInfo: optional
	SSlInfo *apigeetargetserver.SSlInfo `hcl:"s_sl_info,block"`
	// Timeouts: optional
	Timeouts *apigeetargetserver.Timeouts `hcl:"timeouts,block"`
}
type apigeeTargetServerAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("description"))
}

// EnvId returns a reference to field env_id of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("env_id"))
}

// Host returns a reference to field host of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("host"))
}

// Id returns a reference to field id of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("id"))
}

// IsEnabled returns a reference to field is_enabled of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) IsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ats.ref.Append("is_enabled"))
}

// Name returns a reference to field name of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("name"))
}

// Port returns a reference to field port of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ats.ref.Append("port"))
}

// Protocol returns a reference to field protocol of google_apigee_target_server.
func (ats apigeeTargetServerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("protocol"))
}

func (ats apigeeTargetServerAttributes) SSlInfo() terra.ListValue[apigeetargetserver.SSlInfoAttributes] {
	return terra.ReferenceAsList[apigeetargetserver.SSlInfoAttributes](ats.ref.Append("s_sl_info"))
}

func (ats apigeeTargetServerAttributes) Timeouts() apigeetargetserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeetargetserver.TimeoutsAttributes](ats.ref.Append("timeouts"))
}

type apigeeTargetServerState struct {
	Description string                            `json:"description"`
	EnvId       string                            `json:"env_id"`
	Host        string                            `json:"host"`
	Id          string                            `json:"id"`
	IsEnabled   bool                              `json:"is_enabled"`
	Name        string                            `json:"name"`
	Port        float64                           `json:"port"`
	Protocol    string                            `json:"protocol"`
	SSlInfo     []apigeetargetserver.SSlInfoState `json:"s_sl_info"`
	Timeouts    *apigeetargetserver.TimeoutsState `json:"timeouts"`
}
