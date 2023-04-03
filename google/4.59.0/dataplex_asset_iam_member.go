// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataplexassetiammember "github.com/golingon/terraproviders/google/4.59.0/dataplexassetiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexAssetIamMember creates a new instance of [DataplexAssetIamMember].
func NewDataplexAssetIamMember(name string, args DataplexAssetIamMemberArgs) *DataplexAssetIamMember {
	return &DataplexAssetIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexAssetIamMember)(nil)

// DataplexAssetIamMember represents the Terraform resource google_dataplex_asset_iam_member.
type DataplexAssetIamMember struct {
	Name      string
	Args      DataplexAssetIamMemberArgs
	state     *dataplexAssetIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) Type() string {
	return "google_dataplex_asset_iam_member"
}

// LocalName returns the local name for [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) LocalName() string {
	return daim.Name
}

// Configuration returns the configuration (args) for [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) Configuration() interface{} {
	return daim.Args
}

// DependOn is used for other resources to depend on [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(daim)
}

// Dependencies returns the list of resources [DataplexAssetIamMember] depends_on.
func (daim *DataplexAssetIamMember) Dependencies() terra.Dependencies {
	return daim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) LifecycleManagement() *terra.Lifecycle {
	return daim.Lifecycle
}

// Attributes returns the attributes for [DataplexAssetIamMember].
func (daim *DataplexAssetIamMember) Attributes() dataplexAssetIamMemberAttributes {
	return dataplexAssetIamMemberAttributes{ref: terra.ReferenceResource(daim)}
}

// ImportState imports the given attribute values into [DataplexAssetIamMember]'s state.
func (daim *DataplexAssetIamMember) ImportState(av io.Reader) error {
	daim.state = &dataplexAssetIamMemberState{}
	if err := json.NewDecoder(av).Decode(daim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", daim.Type(), daim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexAssetIamMember] has state.
func (daim *DataplexAssetIamMember) State() (*dataplexAssetIamMemberState, bool) {
	return daim.state, daim.state != nil
}

// StateMust returns the state for [DataplexAssetIamMember]. Panics if the state is nil.
func (daim *DataplexAssetIamMember) StateMust() *dataplexAssetIamMemberState {
	if daim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", daim.Type(), daim.LocalName()))
	}
	return daim.state
}

// DataplexAssetIamMemberArgs contains the configurations for google_dataplex_asset_iam_member.
type DataplexAssetIamMemberArgs struct {
	// Asset: string, required
	Asset terra.StringValue `hcl:"asset,attr" validate:"required"`
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
	Condition *dataplexassetiammember.Condition `hcl:"condition,block"`
}
type dataplexAssetIamMemberAttributes struct {
	ref terra.Reference
}

// Asset returns a reference to field asset of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Asset() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("asset"))
}

// DataplexZone returns a reference to field dataplex_zone of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) DataplexZone() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("dataplex_zone"))
}

// Etag returns a reference to field etag of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_asset_iam_member.
func (daim dataplexAssetIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(daim.ref.Append("role"))
}

func (daim dataplexAssetIamMemberAttributes) Condition() terra.ListValue[dataplexassetiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataplexassetiammember.ConditionAttributes](daim.ref.Append("condition"))
}

type dataplexAssetIamMemberState struct {
	Asset        string                                  `json:"asset"`
	DataplexZone string                                  `json:"dataplex_zone"`
	Etag         string                                  `json:"etag"`
	Id           string                                  `json:"id"`
	Lake         string                                  `json:"lake"`
	Location     string                                  `json:"location"`
	Member       string                                  `json:"member"`
	Project      string                                  `json:"project"`
	Role         string                                  `json:"role"`
	Condition    []dataplexassetiammember.ConditionState `json:"condition"`
}
