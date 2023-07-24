// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	containeranalysisoccurrence "github.com/golingon/terraproviders/googlebeta/4.74.0/containeranalysisoccurrence"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAnalysisOccurrence creates a new instance of [ContainerAnalysisOccurrence].
func NewContainerAnalysisOccurrence(name string, args ContainerAnalysisOccurrenceArgs) *ContainerAnalysisOccurrence {
	return &ContainerAnalysisOccurrence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAnalysisOccurrence)(nil)

// ContainerAnalysisOccurrence represents the Terraform resource google_container_analysis_occurrence.
type ContainerAnalysisOccurrence struct {
	Name      string
	Args      ContainerAnalysisOccurrenceArgs
	state     *containerAnalysisOccurrenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) Type() string {
	return "google_container_analysis_occurrence"
}

// LocalName returns the local name for [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) LocalName() string {
	return cao.Name
}

// Configuration returns the configuration (args) for [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) Configuration() interface{} {
	return cao.Args
}

// DependOn is used for other resources to depend on [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) DependOn() terra.Reference {
	return terra.ReferenceResource(cao)
}

// Dependencies returns the list of resources [ContainerAnalysisOccurrence] depends_on.
func (cao *ContainerAnalysisOccurrence) Dependencies() terra.Dependencies {
	return cao.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) LifecycleManagement() *terra.Lifecycle {
	return cao.Lifecycle
}

// Attributes returns the attributes for [ContainerAnalysisOccurrence].
func (cao *ContainerAnalysisOccurrence) Attributes() containerAnalysisOccurrenceAttributes {
	return containerAnalysisOccurrenceAttributes{ref: terra.ReferenceResource(cao)}
}

// ImportState imports the given attribute values into [ContainerAnalysisOccurrence]'s state.
func (cao *ContainerAnalysisOccurrence) ImportState(av io.Reader) error {
	cao.state = &containerAnalysisOccurrenceState{}
	if err := json.NewDecoder(av).Decode(cao.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cao.Type(), cao.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAnalysisOccurrence] has state.
func (cao *ContainerAnalysisOccurrence) State() (*containerAnalysisOccurrenceState, bool) {
	return cao.state, cao.state != nil
}

// StateMust returns the state for [ContainerAnalysisOccurrence]. Panics if the state is nil.
func (cao *ContainerAnalysisOccurrence) StateMust() *containerAnalysisOccurrenceState {
	if cao.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cao.Type(), cao.LocalName()))
	}
	return cao.state
}

// ContainerAnalysisOccurrenceArgs contains the configurations for google_container_analysis_occurrence.
type ContainerAnalysisOccurrenceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NoteName: string, required
	NoteName terra.StringValue `hcl:"note_name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Remediation: string, optional
	Remediation terra.StringValue `hcl:"remediation,attr"`
	// ResourceUri: string, required
	ResourceUri terra.StringValue `hcl:"resource_uri,attr" validate:"required"`
	// Attestation: required
	Attestation *containeranalysisoccurrence.Attestation `hcl:"attestation,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containeranalysisoccurrence.Timeouts `hcl:"timeouts,block"`
}
type containerAnalysisOccurrenceAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("create_time"))
}

// Id returns a reference to field id of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("id"))
}

// Kind returns a reference to field kind of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("kind"))
}

// Name returns a reference to field name of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("name"))
}

// NoteName returns a reference to field note_name of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) NoteName() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("note_name"))
}

// Project returns a reference to field project of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("project"))
}

// Remediation returns a reference to field remediation of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) Remediation() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("remediation"))
}

// ResourceUri returns a reference to field resource_uri of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) ResourceUri() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("resource_uri"))
}

// UpdateTime returns a reference to field update_time of google_container_analysis_occurrence.
func (cao containerAnalysisOccurrenceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cao.ref.Append("update_time"))
}

func (cao containerAnalysisOccurrenceAttributes) Attestation() terra.ListValue[containeranalysisoccurrence.AttestationAttributes] {
	return terra.ReferenceAsList[containeranalysisoccurrence.AttestationAttributes](cao.ref.Append("attestation"))
}

func (cao containerAnalysisOccurrenceAttributes) Timeouts() containeranalysisoccurrence.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containeranalysisoccurrence.TimeoutsAttributes](cao.ref.Append("timeouts"))
}

type containerAnalysisOccurrenceState struct {
	CreateTime  string                                         `json:"create_time"`
	Id          string                                         `json:"id"`
	Kind        string                                         `json:"kind"`
	Name        string                                         `json:"name"`
	NoteName    string                                         `json:"note_name"`
	Project     string                                         `json:"project"`
	Remediation string                                         `json:"remediation"`
	ResourceUri string                                         `json:"resource_uri"`
	UpdateTime  string                                         `json:"update_time"`
	Attestation []containeranalysisoccurrence.AttestationState `json:"attestation"`
	Timeouts    *containeranalysisoccurrence.TimeoutsState     `json:"timeouts"`
}
