// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	spannerinstance "github.com/golingon/terraproviders/googlebeta/4.63.1/spannerinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpannerInstance creates a new instance of [SpannerInstance].
func NewSpannerInstance(name string, args SpannerInstanceArgs) *SpannerInstance {
	return &SpannerInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpannerInstance)(nil)

// SpannerInstance represents the Terraform resource google_spanner_instance.
type SpannerInstance struct {
	Name      string
	Args      SpannerInstanceArgs
	state     *spannerInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpannerInstance].
func (si *SpannerInstance) Type() string {
	return "google_spanner_instance"
}

// LocalName returns the local name for [SpannerInstance].
func (si *SpannerInstance) LocalName() string {
	return si.Name
}

// Configuration returns the configuration (args) for [SpannerInstance].
func (si *SpannerInstance) Configuration() interface{} {
	return si.Args
}

// DependOn is used for other resources to depend on [SpannerInstance].
func (si *SpannerInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(si)
}

// Dependencies returns the list of resources [SpannerInstance] depends_on.
func (si *SpannerInstance) Dependencies() terra.Dependencies {
	return si.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpannerInstance].
func (si *SpannerInstance) LifecycleManagement() *terra.Lifecycle {
	return si.Lifecycle
}

// Attributes returns the attributes for [SpannerInstance].
func (si *SpannerInstance) Attributes() spannerInstanceAttributes {
	return spannerInstanceAttributes{ref: terra.ReferenceResource(si)}
}

// ImportState imports the given attribute values into [SpannerInstance]'s state.
func (si *SpannerInstance) ImportState(av io.Reader) error {
	si.state = &spannerInstanceState{}
	if err := json.NewDecoder(av).Decode(si.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", si.Type(), si.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpannerInstance] has state.
func (si *SpannerInstance) State() (*spannerInstanceState, bool) {
	return si.state, si.state != nil
}

// StateMust returns the state for [SpannerInstance]. Panics if the state is nil.
func (si *SpannerInstance) StateMust() *spannerInstanceState {
	if si.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", si.Type(), si.LocalName()))
	}
	return si.state
}

// SpannerInstanceArgs contains the configurations for google_spanner_instance.
type SpannerInstanceArgs struct {
	// Config: string, required
	Config terra.StringValue `hcl:"config,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NumNodes: number, optional
	NumNodes terra.NumberValue `hcl:"num_nodes,attr"`
	// ProcessingUnits: number, optional
	ProcessingUnits terra.NumberValue `hcl:"processing_units,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *spannerinstance.Timeouts `hcl:"timeouts,block"`
}
type spannerInstanceAttributes struct {
	ref terra.Reference
}

// Config returns a reference to field config of google_spanner_instance.
func (si spannerInstanceAttributes) Config() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("config"))
}

// DisplayName returns a reference to field display_name of google_spanner_instance.
func (si spannerInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("display_name"))
}

// ForceDestroy returns a reference to field force_destroy of google_spanner_instance.
func (si spannerInstanceAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(si.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_spanner_instance.
func (si spannerInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("id"))
}

// Labels returns a reference to field labels of google_spanner_instance.
func (si spannerInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("labels"))
}

// Name returns a reference to field name of google_spanner_instance.
func (si spannerInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("name"))
}

// NumNodes returns a reference to field num_nodes of google_spanner_instance.
func (si spannerInstanceAttributes) NumNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("num_nodes"))
}

// ProcessingUnits returns a reference to field processing_units of google_spanner_instance.
func (si spannerInstanceAttributes) ProcessingUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("processing_units"))
}

// Project returns a reference to field project of google_spanner_instance.
func (si spannerInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("project"))
}

// State returns a reference to field state of google_spanner_instance.
func (si spannerInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("state"))
}

func (si spannerInstanceAttributes) Timeouts() spannerinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[spannerinstance.TimeoutsAttributes](si.ref.Append("timeouts"))
}

type spannerInstanceState struct {
	Config          string                         `json:"config"`
	DisplayName     string                         `json:"display_name"`
	ForceDestroy    bool                           `json:"force_destroy"`
	Id              string                         `json:"id"`
	Labels          map[string]string              `json:"labels"`
	Name            string                         `json:"name"`
	NumNodes        float64                        `json:"num_nodes"`
	ProcessingUnits float64                        `json:"processing_units"`
	Project         string                         `json:"project"`
	State           string                         `json:"state"`
	Timeouts        *spannerinstance.TimeoutsState `json:"timeouts"`
}
