// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexlakeiammember "github.com/golingon/terraproviders/google/4.66.0/dataplexlakeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexLakeIamMember creates a new instance of [DataplexLakeIamMember].
func NewDataplexLakeIamMember(name string, args DataplexLakeIamMemberArgs) *DataplexLakeIamMember {
	return &DataplexLakeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexLakeIamMember)(nil)

// DataplexLakeIamMember represents the Terraform resource google_dataplex_lake_iam_member.
type DataplexLakeIamMember struct {
	Name      string
	Args      DataplexLakeIamMemberArgs
	state     *dataplexLakeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) Type() string {
	return "google_dataplex_lake_iam_member"
}

// LocalName returns the local name for [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) LocalName() string {
	return dlim.Name
}

// Configuration returns the configuration (args) for [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) Configuration() interface{} {
	return dlim.Args
}

// DependOn is used for other resources to depend on [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dlim)
}

// Dependencies returns the list of resources [DataplexLakeIamMember] depends_on.
func (dlim *DataplexLakeIamMember) Dependencies() terra.Dependencies {
	return dlim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) LifecycleManagement() *terra.Lifecycle {
	return dlim.Lifecycle
}

// Attributes returns the attributes for [DataplexLakeIamMember].
func (dlim *DataplexLakeIamMember) Attributes() dataplexLakeIamMemberAttributes {
	return dataplexLakeIamMemberAttributes{ref: terra.ReferenceResource(dlim)}
}

// ImportState imports the given attribute values into [DataplexLakeIamMember]'s state.
func (dlim *DataplexLakeIamMember) ImportState(av io.Reader) error {
	dlim.state = &dataplexLakeIamMemberState{}
	if err := json.NewDecoder(av).Decode(dlim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlim.Type(), dlim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexLakeIamMember] has state.
func (dlim *DataplexLakeIamMember) State() (*dataplexLakeIamMemberState, bool) {
	return dlim.state, dlim.state != nil
}

// StateMust returns the state for [DataplexLakeIamMember]. Panics if the state is nil.
func (dlim *DataplexLakeIamMember) StateMust() *dataplexLakeIamMemberState {
	if dlim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlim.Type(), dlim.LocalName()))
	}
	return dlim.state
}

// DataplexLakeIamMemberArgs contains the configurations for google_dataplex_lake_iam_member.
type DataplexLakeIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataplexlakeiammember.Condition `hcl:"condition,block"`
}
type dataplexLakeIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_lake_iam_member.
func (dlim dataplexLakeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dlim.ref.Append("role"))
}

func (dlim dataplexLakeIamMemberAttributes) Condition() terra.ListValue[dataplexlakeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexlakeiammember.ConditionAttributes](dlim.ref.Append("condition"))
}

type dataplexLakeIamMemberState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Lake      string                                 `json:"lake"`
	Location  string                                 `json:"location"`
	Member    string                                 `json:"member"`
	Project   string                                 `json:"project"`
	Role      string                                 `json:"role"`
	Condition []dataplexlakeiammember.ConditionState `json:"condition"`
}
