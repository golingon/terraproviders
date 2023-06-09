// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computediskiammember "github.com/golingon/terraproviders/google/4.71.0/computediskiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeDiskIamMember creates a new instance of [ComputeDiskIamMember].
func NewComputeDiskIamMember(name string, args ComputeDiskIamMemberArgs) *ComputeDiskIamMember {
	return &ComputeDiskIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeDiskIamMember)(nil)

// ComputeDiskIamMember represents the Terraform resource google_compute_disk_iam_member.
type ComputeDiskIamMember struct {
	Name      string
	Args      ComputeDiskIamMemberArgs
	state     *computeDiskIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) Type() string {
	return "google_compute_disk_iam_member"
}

// LocalName returns the local name for [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) LocalName() string {
	return cdim.Name
}

// Configuration returns the configuration (args) for [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) Configuration() interface{} {
	return cdim.Args
}

// DependOn is used for other resources to depend on [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(cdim)
}

// Dependencies returns the list of resources [ComputeDiskIamMember] depends_on.
func (cdim *ComputeDiskIamMember) Dependencies() terra.Dependencies {
	return cdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) LifecycleManagement() *terra.Lifecycle {
	return cdim.Lifecycle
}

// Attributes returns the attributes for [ComputeDiskIamMember].
func (cdim *ComputeDiskIamMember) Attributes() computeDiskIamMemberAttributes {
	return computeDiskIamMemberAttributes{ref: terra.ReferenceResource(cdim)}
}

// ImportState imports the given attribute values into [ComputeDiskIamMember]'s state.
func (cdim *ComputeDiskIamMember) ImportState(av io.Reader) error {
	cdim.state = &computeDiskIamMemberState{}
	if err := json.NewDecoder(av).Decode(cdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdim.Type(), cdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeDiskIamMember] has state.
func (cdim *ComputeDiskIamMember) State() (*computeDiskIamMemberState, bool) {
	return cdim.state, cdim.state != nil
}

// StateMust returns the state for [ComputeDiskIamMember]. Panics if the state is nil.
func (cdim *ComputeDiskIamMember) StateMust() *computeDiskIamMemberState {
	if cdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdim.Type(), cdim.LocalName()))
	}
	return cdim.state
}

// ComputeDiskIamMemberArgs contains the configurations for google_compute_disk_iam_member.
type ComputeDiskIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Condition: optional
	Condition *computediskiammember.Condition `hcl:"condition,block"`
}
type computeDiskIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("role"))
}

// Zone returns a reference to field zone of google_compute_disk_iam_member.
func (cdim computeDiskIamMemberAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cdim.ref.Append("zone"))
}

func (cdim computeDiskIamMemberAttributes) Condition() terra.ListValue[computediskiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computediskiammember.ConditionAttributes](cdim.ref.Append("condition"))
}

type computeDiskIamMemberState struct {
	Etag      string                                `json:"etag"`
	Id        string                                `json:"id"`
	Member    string                                `json:"member"`
	Name      string                                `json:"name"`
	Project   string                                `json:"project"`
	Role      string                                `json:"role"`
	Zone      string                                `json:"zone"`
	Condition []computediskiammember.ConditionState `json:"condition"`
}
