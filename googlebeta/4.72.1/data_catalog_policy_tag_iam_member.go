// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogpolicytagiammember "github.com/golingon/terraproviders/googlebeta/4.72.1/datacatalogpolicytagiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogPolicyTagIamMember creates a new instance of [DataCatalogPolicyTagIamMember].
func NewDataCatalogPolicyTagIamMember(name string, args DataCatalogPolicyTagIamMemberArgs) *DataCatalogPolicyTagIamMember {
	return &DataCatalogPolicyTagIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogPolicyTagIamMember)(nil)

// DataCatalogPolicyTagIamMember represents the Terraform resource google_data_catalog_policy_tag_iam_member.
type DataCatalogPolicyTagIamMember struct {
	Name      string
	Args      DataCatalogPolicyTagIamMemberArgs
	state     *dataCatalogPolicyTagIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) Type() string {
	return "google_data_catalog_policy_tag_iam_member"
}

// LocalName returns the local name for [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) LocalName() string {
	return dcptim.Name
}

// Configuration returns the configuration (args) for [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) Configuration() interface{} {
	return dcptim.Args
}

// DependOn is used for other resources to depend on [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dcptim)
}

// Dependencies returns the list of resources [DataCatalogPolicyTagIamMember] depends_on.
func (dcptim *DataCatalogPolicyTagIamMember) Dependencies() terra.Dependencies {
	return dcptim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) LifecycleManagement() *terra.Lifecycle {
	return dcptim.Lifecycle
}

// Attributes returns the attributes for [DataCatalogPolicyTagIamMember].
func (dcptim *DataCatalogPolicyTagIamMember) Attributes() dataCatalogPolicyTagIamMemberAttributes {
	return dataCatalogPolicyTagIamMemberAttributes{ref: terra.ReferenceResource(dcptim)}
}

// ImportState imports the given attribute values into [DataCatalogPolicyTagIamMember]'s state.
func (dcptim *DataCatalogPolicyTagIamMember) ImportState(av io.Reader) error {
	dcptim.state = &dataCatalogPolicyTagIamMemberState{}
	if err := json.NewDecoder(av).Decode(dcptim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcptim.Type(), dcptim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogPolicyTagIamMember] has state.
func (dcptim *DataCatalogPolicyTagIamMember) State() (*dataCatalogPolicyTagIamMemberState, bool) {
	return dcptim.state, dcptim.state != nil
}

// StateMust returns the state for [DataCatalogPolicyTagIamMember]. Panics if the state is nil.
func (dcptim *DataCatalogPolicyTagIamMember) StateMust() *dataCatalogPolicyTagIamMemberState {
	if dcptim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcptim.Type(), dcptim.LocalName()))
	}
	return dcptim.state
}

// DataCatalogPolicyTagIamMemberArgs contains the configurations for google_data_catalog_policy_tag_iam_member.
type DataCatalogPolicyTagIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// PolicyTag: string, required
	PolicyTag terra.StringValue `hcl:"policy_tag,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *datacatalogpolicytagiammember.Condition `hcl:"condition,block"`
}
type dataCatalogPolicyTagIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_policy_tag_iam_member.
func (dcptim dataCatalogPolicyTagIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcptim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_policy_tag_iam_member.
func (dcptim dataCatalogPolicyTagIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcptim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_catalog_policy_tag_iam_member.
func (dcptim dataCatalogPolicyTagIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dcptim.ref.Append("member"))
}

// PolicyTag returns a reference to field policy_tag of google_data_catalog_policy_tag_iam_member.
func (dcptim dataCatalogPolicyTagIamMemberAttributes) PolicyTag() terra.StringValue {
	return terra.ReferenceAsString(dcptim.ref.Append("policy_tag"))
}

// Role returns a reference to field role of google_data_catalog_policy_tag_iam_member.
func (dcptim dataCatalogPolicyTagIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcptim.ref.Append("role"))
}

func (dcptim dataCatalogPolicyTagIamMemberAttributes) Condition() terra.ListValue[datacatalogpolicytagiammember.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogpolicytagiammember.ConditionAttributes](dcptim.ref.Append("condition"))
}

type dataCatalogPolicyTagIamMemberState struct {
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Member    string                                         `json:"member"`
	PolicyTag string                                         `json:"policy_tag"`
	Role      string                                         `json:"role"`
	Condition []datacatalogpolicytagiammember.ConditionState `json:"condition"`
}
