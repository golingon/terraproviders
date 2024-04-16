// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataplex_datascan_iam_member

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_dataplex_datascan_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataplexDatascanIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gddim *Resource) Type() string {
	return "google_dataplex_datascan_iam_member"
}

// LocalName returns the local name for [Resource].
func (gddim *Resource) LocalName() string {
	return gddim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gddim *Resource) Configuration() interface{} {
	return gddim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gddim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gddim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gddim *Resource) Dependencies() terra.Dependencies {
	return gddim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gddim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gddim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gddim *Resource) Attributes() googleDataplexDatascanIamMemberAttributes {
	return googleDataplexDatascanIamMemberAttributes{ref: terra.ReferenceResource(gddim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gddim *Resource) ImportState(state io.Reader) error {
	gddim.state = &googleDataplexDatascanIamMemberState{}
	if err := json.NewDecoder(state).Decode(gddim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gddim.Type(), gddim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gddim *Resource) State() (*googleDataplexDatascanIamMemberState, bool) {
	return gddim.state, gddim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gddim *Resource) StateMust() *googleDataplexDatascanIamMemberState {
	if gddim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gddim.Type(), gddim.LocalName()))
	}
	return gddim.state
}

// Args contains the configurations for google_dataplex_datascan_iam_member.
type Args struct {
	// DataScanId: string, required
	DataScanId terra.StringValue `hcl:"data_scan_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleDataplexDatascanIamMemberAttributes struct {
	ref terra.Reference
}

// DataScanId returns a reference to field data_scan_id of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) DataScanId() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("data_scan_id"))
}

// Etag returns a reference to field etag of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("id"))
}

// Location returns a reference to field location of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_datascan_iam_member.
func (gddim googleDataplexDatascanIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gddim.ref.Append("role"))
}

func (gddim googleDataplexDatascanIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gddim.ref.Append("condition"))
}

type googleDataplexDatascanIamMemberState struct {
	DataScanId string           `json:"data_scan_id"`
	Etag       string           `json:"etag"`
	Id         string           `json:"id"`
	Location   string           `json:"location"`
	Member     string           `json:"member"`
	Project    string           `json:"project"`
	Role       string           `json:"role"`
	Condition  []ConditionState `json:"condition"`
}
