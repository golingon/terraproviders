// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sesv2_account_vdm_attributes

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

// Resource represents the Terraform resource aws_sesv2_account_vdm_attributes.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSesv2AccountVdmAttributesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asava *Resource) Type() string {
	return "aws_sesv2_account_vdm_attributes"
}

// LocalName returns the local name for [Resource].
func (asava *Resource) LocalName() string {
	return asava.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asava *Resource) Configuration() interface{} {
	return asava.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asava *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asava)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asava *Resource) Dependencies() terra.Dependencies {
	return asava.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asava *Resource) LifecycleManagement() *terra.Lifecycle {
	return asava.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asava *Resource) Attributes() awsSesv2AccountVdmAttributesAttributes {
	return awsSesv2AccountVdmAttributesAttributes{ref: terra.ReferenceResource(asava)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asava *Resource) ImportState(state io.Reader) error {
	asava.state = &awsSesv2AccountVdmAttributesState{}
	if err := json.NewDecoder(state).Decode(asava.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asava.Type(), asava.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asava *Resource) State() (*awsSesv2AccountVdmAttributesState, bool) {
	return asava.state, asava.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asava *Resource) StateMust() *awsSesv2AccountVdmAttributesState {
	if asava.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asava.Type(), asava.LocalName()))
	}
	return asava.state
}

// Args contains the configurations for aws_sesv2_account_vdm_attributes.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VdmEnabled: string, required
	VdmEnabled terra.StringValue `hcl:"vdm_enabled,attr" validate:"required"`
	// DashboardAttributes: optional
	DashboardAttributes *DashboardAttributes `hcl:"dashboard_attributes,block"`
	// GuardianAttributes: optional
	GuardianAttributes *GuardianAttributes `hcl:"guardian_attributes,block"`
}

type awsSesv2AccountVdmAttributesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_sesv2_account_vdm_attributes.
func (asava awsSesv2AccountVdmAttributesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asava.ref.Append("id"))
}

// VdmEnabled returns a reference to field vdm_enabled of aws_sesv2_account_vdm_attributes.
func (asava awsSesv2AccountVdmAttributesAttributes) VdmEnabled() terra.StringValue {
	return terra.ReferenceAsString(asava.ref.Append("vdm_enabled"))
}

func (asava awsSesv2AccountVdmAttributesAttributes) DashboardAttributes() terra.ListValue[DashboardAttributesAttributes] {
	return terra.ReferenceAsList[DashboardAttributesAttributes](asava.ref.Append("dashboard_attributes"))
}

func (asava awsSesv2AccountVdmAttributesAttributes) GuardianAttributes() terra.ListValue[GuardianAttributesAttributes] {
	return terra.ReferenceAsList[GuardianAttributesAttributes](asava.ref.Append("guardian_attributes"))
}

type awsSesv2AccountVdmAttributesState struct {
	Id                  string                     `json:"id"`
	VdmEnabled          string                     `json:"vdm_enabled"`
	DashboardAttributes []DashboardAttributesState `json:"dashboard_attributes"`
	GuardianAttributes  []GuardianAttributesState  `json:"guardian_attributes"`
}
