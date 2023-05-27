// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sccsource "github.com/golingon/terraproviders/google/4.66.0/sccsource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSccSource creates a new instance of [SccSource].
func NewSccSource(name string, args SccSourceArgs) *SccSource {
	return &SccSource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccSource)(nil)

// SccSource represents the Terraform resource google_scc_source.
type SccSource struct {
	Name      string
	Args      SccSourceArgs
	state     *sccSourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccSource].
func (ss *SccSource) Type() string {
	return "google_scc_source"
}

// LocalName returns the local name for [SccSource].
func (ss *SccSource) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SccSource].
func (ss *SccSource) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SccSource].
func (ss *SccSource) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SccSource] depends_on.
func (ss *SccSource) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccSource].
func (ss *SccSource) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SccSource].
func (ss *SccSource) Attributes() sccSourceAttributes {
	return sccSourceAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SccSource]'s state.
func (ss *SccSource) ImportState(av io.Reader) error {
	ss.state = &sccSourceState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccSource] has state.
func (ss *SccSource) State() (*sccSourceState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SccSource]. Panics if the state is nil.
func (ss *SccSource) StateMust() *sccSourceState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SccSourceArgs contains the configurations for google_scc_source.
type SccSourceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *sccsource.Timeouts `hcl:"timeouts,block"`
}
type sccSourceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_scc_source.
func (ss sccSourceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_scc_source.
func (ss sccSourceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("display_name"))
}

// Id returns a reference to field id of google_scc_source.
func (ss sccSourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Name returns a reference to field name of google_scc_source.
func (ss sccSourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// Organization returns a reference to field organization of google_scc_source.
func (ss sccSourceAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("organization"))
}

func (ss sccSourceAttributes) Timeouts() sccsource.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sccsource.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type sccSourceState struct {
	Description  string                   `json:"description"`
	DisplayName  string                   `json:"display_name"`
	Id           string                   `json:"id"`
	Name         string                   `json:"name"`
	Organization string                   `json:"organization"`
	Timeouts     *sccsource.TimeoutsState `json:"timeouts"`
}
