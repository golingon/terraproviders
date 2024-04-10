// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dataplexdatascaniammember "github.com/golingon/terraproviders/googlebeta/5.24.0/dataplexdatascaniammember"
	"io"
)

// NewDataplexDatascanIamMember creates a new instance of [DataplexDatascanIamMember].
func NewDataplexDatascanIamMember(name string, args DataplexDatascanIamMemberArgs) *DataplexDatascanIamMember {
	return &DataplexDatascanIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexDatascanIamMember)(nil)

// DataplexDatascanIamMember represents the Terraform resource google_dataplex_datascan_iam_member.
type DataplexDatascanIamMember struct {
	Name      string
	Args      DataplexDatascanIamMemberArgs
	state     *dataplexDatascanIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) Type() string {
	return "google_dataplex_datascan_iam_member"
}

// LocalName returns the local name for [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) LocalName() string {
	return ddim.Name
}

// Configuration returns the configuration (args) for [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) Configuration() interface{} {
	return ddim.Args
}

// DependOn is used for other resources to depend on [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ddim)
}

// Dependencies returns the list of resources [DataplexDatascanIamMember] depends_on.
func (ddim *DataplexDatascanIamMember) Dependencies() terra.Dependencies {
	return ddim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) LifecycleManagement() *terra.Lifecycle {
	return ddim.Lifecycle
}

// Attributes returns the attributes for [DataplexDatascanIamMember].
func (ddim *DataplexDatascanIamMember) Attributes() dataplexDatascanIamMemberAttributes {
	return dataplexDatascanIamMemberAttributes{ref: terra.ReferenceResource(ddim)}
}

// ImportState imports the given attribute values into [DataplexDatascanIamMember]'s state.
func (ddim *DataplexDatascanIamMember) ImportState(av io.Reader) error {
	ddim.state = &dataplexDatascanIamMemberState{}
	if err := json.NewDecoder(av).Decode(ddim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ddim.Type(), ddim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexDatascanIamMember] has state.
func (ddim *DataplexDatascanIamMember) State() (*dataplexDatascanIamMemberState, bool) {
	return ddim.state, ddim.state != nil
}

// StateMust returns the state for [DataplexDatascanIamMember]. Panics if the state is nil.
func (ddim *DataplexDatascanIamMember) StateMust() *dataplexDatascanIamMemberState {
	if ddim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ddim.Type(), ddim.LocalName()))
	}
	return ddim.state
}

// DataplexDatascanIamMemberArgs contains the configurations for google_dataplex_datascan_iam_member.
type DataplexDatascanIamMemberArgs struct {
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
	Condition *dataplexdatascaniammember.Condition `hcl:"condition,block"`
}
type dataplexDatascanIamMemberAttributes struct {
	ref terra.Reference
}

// DataScanId returns a reference to field data_scan_id of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) DataScanId() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("data_scan_id"))
}

// Etag returns a reference to field etag of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("id"))
}

// Location returns a reference to field location of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_datascan_iam_member.
func (ddim dataplexDatascanIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ddim.ref.Append("role"))
}

func (ddim dataplexDatascanIamMemberAttributes) Condition() terra.ListValue[dataplexdatascaniammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexdatascaniammember.ConditionAttributes](ddim.ref.Append("condition"))
}

type dataplexDatascanIamMemberState struct {
	DataScanId string                                     `json:"data_scan_id"`
	Etag       string                                     `json:"etag"`
	Id         string                                     `json:"id"`
	Location   string                                     `json:"location"`
	Member     string                                     `json:"member"`
	Project    string                                     `json:"project"`
	Role       string                                     `json:"role"`
	Condition  []dataplexdatascaniammember.ConditionState `json:"condition"`
}
