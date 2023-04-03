// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	videoanalyzeredgemodule "github.com/golingon/terraproviders/azurerm/3.49.0/videoanalyzeredgemodule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVideoAnalyzerEdgeModule creates a new instance of [VideoAnalyzerEdgeModule].
func NewVideoAnalyzerEdgeModule(name string, args VideoAnalyzerEdgeModuleArgs) *VideoAnalyzerEdgeModule {
	return &VideoAnalyzerEdgeModule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VideoAnalyzerEdgeModule)(nil)

// VideoAnalyzerEdgeModule represents the Terraform resource azurerm_video_analyzer_edge_module.
type VideoAnalyzerEdgeModule struct {
	Name      string
	Args      VideoAnalyzerEdgeModuleArgs
	state     *videoAnalyzerEdgeModuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) Type() string {
	return "azurerm_video_analyzer_edge_module"
}

// LocalName returns the local name for [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) LocalName() string {
	return vaem.Name
}

// Configuration returns the configuration (args) for [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) Configuration() interface{} {
	return vaem.Args
}

// DependOn is used for other resources to depend on [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) DependOn() terra.Reference {
	return terra.ReferenceResource(vaem)
}

// Dependencies returns the list of resources [VideoAnalyzerEdgeModule] depends_on.
func (vaem *VideoAnalyzerEdgeModule) Dependencies() terra.Dependencies {
	return vaem.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) LifecycleManagement() *terra.Lifecycle {
	return vaem.Lifecycle
}

// Attributes returns the attributes for [VideoAnalyzerEdgeModule].
func (vaem *VideoAnalyzerEdgeModule) Attributes() videoAnalyzerEdgeModuleAttributes {
	return videoAnalyzerEdgeModuleAttributes{ref: terra.ReferenceResource(vaem)}
}

// ImportState imports the given attribute values into [VideoAnalyzerEdgeModule]'s state.
func (vaem *VideoAnalyzerEdgeModule) ImportState(av io.Reader) error {
	vaem.state = &videoAnalyzerEdgeModuleState{}
	if err := json.NewDecoder(av).Decode(vaem.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vaem.Type(), vaem.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VideoAnalyzerEdgeModule] has state.
func (vaem *VideoAnalyzerEdgeModule) State() (*videoAnalyzerEdgeModuleState, bool) {
	return vaem.state, vaem.state != nil
}

// StateMust returns the state for [VideoAnalyzerEdgeModule]. Panics if the state is nil.
func (vaem *VideoAnalyzerEdgeModule) StateMust() *videoAnalyzerEdgeModuleState {
	if vaem.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vaem.Type(), vaem.LocalName()))
	}
	return vaem.state
}

// VideoAnalyzerEdgeModuleArgs contains the configurations for azurerm_video_analyzer_edge_module.
type VideoAnalyzerEdgeModuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VideoAnalyzerName: string, required
	VideoAnalyzerName terra.StringValue `hcl:"video_analyzer_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *videoanalyzeredgemodule.Timeouts `hcl:"timeouts,block"`
}
type videoAnalyzerEdgeModuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_video_analyzer_edge_module.
func (vaem videoAnalyzerEdgeModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vaem.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_video_analyzer_edge_module.
func (vaem videoAnalyzerEdgeModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vaem.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_video_analyzer_edge_module.
func (vaem videoAnalyzerEdgeModuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vaem.ref.Append("resource_group_name"))
}

// VideoAnalyzerName returns a reference to field video_analyzer_name of azurerm_video_analyzer_edge_module.
func (vaem videoAnalyzerEdgeModuleAttributes) VideoAnalyzerName() terra.StringValue {
	return terra.ReferenceAsString(vaem.ref.Append("video_analyzer_name"))
}

func (vaem videoAnalyzerEdgeModuleAttributes) Timeouts() videoanalyzeredgemodule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[videoanalyzeredgemodule.TimeoutsAttributes](vaem.ref.Append("timeouts"))
}

type videoAnalyzerEdgeModuleState struct {
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	VideoAnalyzerName string                                 `json:"video_analyzer_name"`
	Timeouts          *videoanalyzeredgemodule.TimeoutsState `json:"timeouts"`
}
