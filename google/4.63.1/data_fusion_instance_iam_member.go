// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datafusioninstanceiammember "github.com/golingon/terraproviders/google/4.63.1/datafusioninstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFusionInstanceIamMember creates a new instance of [DataFusionInstanceIamMember].
func NewDataFusionInstanceIamMember(name string, args DataFusionInstanceIamMemberArgs) *DataFusionInstanceIamMember {
	return &DataFusionInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFusionInstanceIamMember)(nil)

// DataFusionInstanceIamMember represents the Terraform resource google_data_fusion_instance_iam_member.
type DataFusionInstanceIamMember struct {
	Name      string
	Args      DataFusionInstanceIamMemberArgs
	state     *dataFusionInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) Type() string {
	return "google_data_fusion_instance_iam_member"
}

// LocalName returns the local name for [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) LocalName() string {
	return dfiim.Name
}

// Configuration returns the configuration (args) for [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) Configuration() interface{} {
	return dfiim.Args
}

// DependOn is used for other resources to depend on [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dfiim)
}

// Dependencies returns the list of resources [DataFusionInstanceIamMember] depends_on.
func (dfiim *DataFusionInstanceIamMember) Dependencies() terra.Dependencies {
	return dfiim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) LifecycleManagement() *terra.Lifecycle {
	return dfiim.Lifecycle
}

// Attributes returns the attributes for [DataFusionInstanceIamMember].
func (dfiim *DataFusionInstanceIamMember) Attributes() dataFusionInstanceIamMemberAttributes {
	return dataFusionInstanceIamMemberAttributes{ref: terra.ReferenceResource(dfiim)}
}

// ImportState imports the given attribute values into [DataFusionInstanceIamMember]'s state.
func (dfiim *DataFusionInstanceIamMember) ImportState(av io.Reader) error {
	dfiim.state = &dataFusionInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(dfiim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfiim.Type(), dfiim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFusionInstanceIamMember] has state.
func (dfiim *DataFusionInstanceIamMember) State() (*dataFusionInstanceIamMemberState, bool) {
	return dfiim.state, dfiim.state != nil
}

// StateMust returns the state for [DataFusionInstanceIamMember]. Panics if the state is nil.
func (dfiim *DataFusionInstanceIamMember) StateMust() *dataFusionInstanceIamMemberState {
	if dfiim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfiim.Type(), dfiim.LocalName()))
	}
	return dfiim.state
}

// DataFusionInstanceIamMemberArgs contains the configurations for google_data_fusion_instance_iam_member.
type DataFusionInstanceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *datafusioninstanceiammember.Condition `hcl:"condition,block"`
}
type dataFusionInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("member"))
}

// Name returns a reference to field name of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("name"))
}

// Project returns a reference to field project of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("project"))
}

// Region returns a reference to field region of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("region"))
}

// Role returns a reference to field role of google_data_fusion_instance_iam_member.
func (dfiim dataFusionInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dfiim.ref.Append("role"))
}

func (dfiim dataFusionInstanceIamMemberAttributes) Condition() terra.ListValue[datafusioninstanceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[datafusioninstanceiammember.ConditionAttributes](dfiim.ref.Append("condition"))
}

type dataFusionInstanceIamMemberState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Member    string                                       `json:"member"`
	Name      string                                       `json:"name"`
	Project   string                                       `json:"project"`
	Region    string                                       `json:"region"`
	Role      string                                       `json:"role"`
	Condition []datafusioninstanceiammember.ConditionState `json:"condition"`
}
