// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	organizationsorganizationalunit "github.com/golingon/terraproviders/aws/5.11.0/organizationsorganizationalunit"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrganizationsOrganizationalUnit creates a new instance of [OrganizationsOrganizationalUnit].
func NewOrganizationsOrganizationalUnit(name string, args OrganizationsOrganizationalUnitArgs) *OrganizationsOrganizationalUnit {
	return &OrganizationsOrganizationalUnit{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationsOrganizationalUnit)(nil)

// OrganizationsOrganizationalUnit represents the Terraform resource aws_organizations_organizational_unit.
type OrganizationsOrganizationalUnit struct {
	Name      string
	Args      OrganizationsOrganizationalUnitArgs
	state     *organizationsOrganizationalUnitState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) Type() string {
	return "aws_organizations_organizational_unit"
}

// LocalName returns the local name for [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) LocalName() string {
	return oou.Name
}

// Configuration returns the configuration (args) for [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) Configuration() interface{} {
	return oou.Args
}

// DependOn is used for other resources to depend on [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) DependOn() terra.Reference {
	return terra.ReferenceResource(oou)
}

// Dependencies returns the list of resources [OrganizationsOrganizationalUnit] depends_on.
func (oou *OrganizationsOrganizationalUnit) Dependencies() terra.Dependencies {
	return oou.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) LifecycleManagement() *terra.Lifecycle {
	return oou.Lifecycle
}

// Attributes returns the attributes for [OrganizationsOrganizationalUnit].
func (oou *OrganizationsOrganizationalUnit) Attributes() organizationsOrganizationalUnitAttributes {
	return organizationsOrganizationalUnitAttributes{ref: terra.ReferenceResource(oou)}
}

// ImportState imports the given attribute values into [OrganizationsOrganizationalUnit]'s state.
func (oou *OrganizationsOrganizationalUnit) ImportState(av io.Reader) error {
	oou.state = &organizationsOrganizationalUnitState{}
	if err := json.NewDecoder(av).Decode(oou.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oou.Type(), oou.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationsOrganizationalUnit] has state.
func (oou *OrganizationsOrganizationalUnit) State() (*organizationsOrganizationalUnitState, bool) {
	return oou.state, oou.state != nil
}

// StateMust returns the state for [OrganizationsOrganizationalUnit]. Panics if the state is nil.
func (oou *OrganizationsOrganizationalUnit) StateMust() *organizationsOrganizationalUnitState {
	if oou.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oou.Type(), oou.LocalName()))
	}
	return oou.state
}

// OrganizationsOrganizationalUnitArgs contains the configurations for aws_organizations_organizational_unit.
type OrganizationsOrganizationalUnitArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Accounts: min=0
	Accounts []organizationsorganizationalunit.Accounts `hcl:"accounts,block" validate:"min=0"`
}
type organizationsOrganizationalUnitAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("arn"))
}

// Id returns a reference to field id of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("id"))
}

// Name returns a reference to field name of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("name"))
}

// ParentId returns a reference to field parent_id of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("parent_id"))
}

// Tags returns a reference to field tags of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oou.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_organizations_organizational_unit.
func (oou organizationsOrganizationalUnitAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](oou.ref.Append("tags_all"))
}

func (oou organizationsOrganizationalUnitAttributes) Accounts() terra.ListValue[organizationsorganizationalunit.AccountsAttributes] {
	return terra.ReferenceAsList[organizationsorganizationalunit.AccountsAttributes](oou.ref.Append("accounts"))
}

type organizationsOrganizationalUnitState struct {
	Arn      string                                          `json:"arn"`
	Id       string                                          `json:"id"`
	Name     string                                          `json:"name"`
	ParentId string                                          `json:"parent_id"`
	Tags     map[string]string                               `json:"tags"`
	TagsAll  map[string]string                               `json:"tags_all"`
	Accounts []organizationsorganizationalunit.AccountsState `json:"accounts"`
}
