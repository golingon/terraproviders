// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	organizationiambinding "github.com/golingon/terraproviders/googlebeta/4.64.0/organizationiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationIamBinding creates a new instance of [OrganizationIamBinding].
func NewOrganizationIamBinding(name string, args OrganizationIamBindingArgs) *OrganizationIamBinding {
	return &OrganizationIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationIamBinding)(nil)

// OrganizationIamBinding represents the Terraform resource google_organization_iam_binding.
type OrganizationIamBinding struct {
	Name      string
	Args      OrganizationIamBindingArgs
	state     *organizationIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationIamBinding].
func (oib *OrganizationIamBinding) Type() string {
	return "google_organization_iam_binding"
}

// LocalName returns the local name for [OrganizationIamBinding].
func (oib *OrganizationIamBinding) LocalName() string {
	return oib.Name
}

// Configuration returns the configuration (args) for [OrganizationIamBinding].
func (oib *OrganizationIamBinding) Configuration() interface{} {
	return oib.Args
}

// DependOn is used for other resources to depend on [OrganizationIamBinding].
func (oib *OrganizationIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(oib)
}

// Dependencies returns the list of resources [OrganizationIamBinding] depends_on.
func (oib *OrganizationIamBinding) Dependencies() terra.Dependencies {
	return oib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationIamBinding].
func (oib *OrganizationIamBinding) LifecycleManagement() *terra.Lifecycle {
	return oib.Lifecycle
}

// Attributes returns the attributes for [OrganizationIamBinding].
func (oib *OrganizationIamBinding) Attributes() organizationIamBindingAttributes {
	return organizationIamBindingAttributes{ref: terra.ReferenceResource(oib)}
}

// ImportState imports the given attribute values into [OrganizationIamBinding]'s state.
func (oib *OrganizationIamBinding) ImportState(av io.Reader) error {
	oib.state = &organizationIamBindingState{}
	if err := json.NewDecoder(av).Decode(oib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oib.Type(), oib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationIamBinding] has state.
func (oib *OrganizationIamBinding) State() (*organizationIamBindingState, bool) {
	return oib.state, oib.state != nil
}

// StateMust returns the state for [OrganizationIamBinding]. Panics if the state is nil.
func (oib *OrganizationIamBinding) StateMust() *organizationIamBindingState {
	if oib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oib.Type(), oib.LocalName()))
	}
	return oib.state
}

// OrganizationIamBindingArgs contains the configurations for google_organization_iam_binding.
type OrganizationIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *organizationiambinding.Condition `hcl:"condition,block"`
}
type organizationIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_organization_iam_binding.
func (oib organizationIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(oib.ref.Append("etag"))
}

// Id returns a reference to field id of google_organization_iam_binding.
func (oib organizationIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oib.ref.Append("id"))
}

// Members returns a reference to field members of google_organization_iam_binding.
func (oib organizationIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oib.ref.Append("members"))
}

// OrgId returns a reference to field org_id of google_organization_iam_binding.
func (oib organizationIamBindingAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(oib.ref.Append("org_id"))
}

// Role returns a reference to field role of google_organization_iam_binding.
func (oib organizationIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(oib.ref.Append("role"))
}

func (oib organizationIamBindingAttributes) Condition() terra.ListValue[organizationiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[organizationiambinding.ConditionAttributes](oib.ref.Append("condition"))
}

type organizationIamBindingState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Members   []string                                `json:"members"`
	OrgId     string                                  `json:"org_id"`
	Role      string                                  `json:"role"`
	Condition []organizationiambinding.ConditionState `json:"condition"`
}
