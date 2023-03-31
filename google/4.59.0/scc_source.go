// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sccsource "github.com/golingon/terraproviders/google/4.59.0/sccsource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSccSource(name string, args SccSourceArgs) *SccSource {
	return &SccSource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccSource)(nil)

type SccSource struct {
	Name  string
	Args  SccSourceArgs
	state *sccSourceState
}

func (ss *SccSource) Type() string {
	return "google_scc_source"
}

func (ss *SccSource) LocalName() string {
	return ss.Name
}

func (ss *SccSource) Configuration() interface{} {
	return ss.Args
}

func (ss *SccSource) Attributes() sccSourceAttributes {
	return sccSourceAttributes{ref: terra.ReferenceResource(ss)}
}

func (ss *SccSource) ImportState(av io.Reader) error {
	ss.state = &sccSourceState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

func (ss *SccSource) State() (*sccSourceState, bool) {
	return ss.state, ss.state != nil
}

func (ss *SccSource) StateMust() *sccSourceState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

func (ss *SccSource) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

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
	// DependsOn contains resources that SccSource depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sccSourceAttributes struct {
	ref terra.Reference
}

func (ss sccSourceAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("description"))
}

func (ss sccSourceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("display_name"))
}

func (ss sccSourceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("id"))
}

func (ss sccSourceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("name"))
}

func (ss sccSourceAttributes) Organization() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("organization"))
}

func (ss sccSourceAttributes) Timeouts() sccsource.TimeoutsAttributes {
	return terra.ReferenceSingle[sccsource.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type sccSourceState struct {
	Description  string                   `json:"description"`
	DisplayName  string                   `json:"display_name"`
	Id           string                   `json:"id"`
	Name         string                   `json:"name"`
	Organization string                   `json:"organization"`
	Timeouts     *sccsource.TimeoutsState `json:"timeouts"`
}
