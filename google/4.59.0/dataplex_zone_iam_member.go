// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexzoneiammember "github.com/golingon/terraproviders/google/4.59.0/dataplexzoneiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexZoneIamMember creates a new instance of [DataplexZoneIamMember].
func NewDataplexZoneIamMember(name string, args DataplexZoneIamMemberArgs) *DataplexZoneIamMember {
	return &DataplexZoneIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexZoneIamMember)(nil)

// DataplexZoneIamMember represents the Terraform resource google_dataplex_zone_iam_member.
type DataplexZoneIamMember struct {
	Name      string
	Args      DataplexZoneIamMemberArgs
	state     *dataplexZoneIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) Type() string {
	return "google_dataplex_zone_iam_member"
}

// LocalName returns the local name for [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) LocalName() string {
	return dzim.Name
}

// Configuration returns the configuration (args) for [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) Configuration() interface{} {
	return dzim.Args
}

// DependOn is used for other resources to depend on [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dzim)
}

// Dependencies returns the list of resources [DataplexZoneIamMember] depends_on.
func (dzim *DataplexZoneIamMember) Dependencies() terra.Dependencies {
	return dzim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) LifecycleManagement() *terra.Lifecycle {
	return dzim.Lifecycle
}

// Attributes returns the attributes for [DataplexZoneIamMember].
func (dzim *DataplexZoneIamMember) Attributes() dataplexZoneIamMemberAttributes {
	return dataplexZoneIamMemberAttributes{ref: terra.ReferenceResource(dzim)}
}

// ImportState imports the given attribute values into [DataplexZoneIamMember]'s state.
func (dzim *DataplexZoneIamMember) ImportState(av io.Reader) error {
	dzim.state = &dataplexZoneIamMemberState{}
	if err := json.NewDecoder(av).Decode(dzim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dzim.Type(), dzim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexZoneIamMember] has state.
func (dzim *DataplexZoneIamMember) State() (*dataplexZoneIamMemberState, bool) {
	return dzim.state, dzim.state != nil
}

// StateMust returns the state for [DataplexZoneIamMember]. Panics if the state is nil.
func (dzim *DataplexZoneIamMember) StateMust() *dataplexZoneIamMemberState {
	if dzim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dzim.Type(), dzim.LocalName()))
	}
	return dzim.state
}

// DataplexZoneIamMemberArgs contains the configurations for google_dataplex_zone_iam_member.
type DataplexZoneIamMemberArgs struct {
	// DataplexZone: string, required
	DataplexZone terra.StringValue `hcl:"dataplex_zone,attr" validate:"required"`
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
	Condition *dataplexzoneiammember.Condition `hcl:"condition,block"`
}
type dataplexZoneIamMemberAttributes struct {
	ref terra.Reference
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_zone_iam_member.
func (dzim dataplexZoneIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dzim.ref.Append("role"))
}

func (dzim dataplexZoneIamMemberAttributes) Condition() terra.ListValue[dataplexzoneiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexzoneiammember.ConditionAttributes](dzim.ref.Append("condition"))
}

type dataplexZoneIamMemberState struct {
	DataplexZone string                                 `json:"dataplex_zone"`
	Etag         string                                 `json:"etag"`
	Id           string                                 `json:"id"`
	Lake         string                                 `json:"lake"`
	Location     string                                 `json:"location"`
	Member       string                                 `json:"member"`
	Project      string                                 `json:"project"`
	Role         string                                 `json:"role"`
	Condition    []dataplexzoneiammember.ConditionState `json:"condition"`
}
