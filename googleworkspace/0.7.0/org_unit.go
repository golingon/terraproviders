// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrgUnit creates a new instance of [OrgUnit].
func NewOrgUnit(name string, args OrgUnitArgs) *OrgUnit {
	return &OrgUnit{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrgUnit)(nil)

// OrgUnit represents the Terraform resource googleworkspace_org_unit.
type OrgUnit struct {
	Name      string
	Args      OrgUnitArgs
	state     *orgUnitState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrgUnit].
func (ou *OrgUnit) Type() string {
	return "googleworkspace_org_unit"
}

// LocalName returns the local name for [OrgUnit].
func (ou *OrgUnit) LocalName() string {
	return ou.Name
}

// Configuration returns the configuration (args) for [OrgUnit].
func (ou *OrgUnit) Configuration() interface{} {
	return ou.Args
}

// DependOn is used for other resources to depend on [OrgUnit].
func (ou *OrgUnit) DependOn() terra.Reference {
	return terra.ReferenceResource(ou)
}

// Dependencies returns the list of resources [OrgUnit] depends_on.
func (ou *OrgUnit) Dependencies() terra.Dependencies {
	return ou.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrgUnit].
func (ou *OrgUnit) LifecycleManagement() *terra.Lifecycle {
	return ou.Lifecycle
}

// Attributes returns the attributes for [OrgUnit].
func (ou *OrgUnit) Attributes() orgUnitAttributes {
	return orgUnitAttributes{ref: terra.ReferenceResource(ou)}
}

// ImportState imports the given attribute values into [OrgUnit]'s state.
func (ou *OrgUnit) ImportState(av io.Reader) error {
	ou.state = &orgUnitState{}
	if err := json.NewDecoder(av).Decode(ou.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ou.Type(), ou.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrgUnit] has state.
func (ou *OrgUnit) State() (*orgUnitState, bool) {
	return ou.state, ou.state != nil
}

// StateMust returns the state for [OrgUnit]. Panics if the state is nil.
func (ou *OrgUnit) StateMust() *orgUnitState {
	if ou.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ou.Type(), ou.LocalName()))
	}
	return ou.state
}

// OrgUnitArgs contains the configurations for googleworkspace_org_unit.
type OrgUnitArgs struct {
	// BlockInheritance: bool, optional
	BlockInheritance terra.BoolValue `hcl:"block_inheritance,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentOrgUnitId: string, optional
	ParentOrgUnitId terra.StringValue `hcl:"parent_org_unit_id,attr"`
	// ParentOrgUnitPath: string, optional
	ParentOrgUnitPath terra.StringValue `hcl:"parent_org_unit_path,attr"`
}
type orgUnitAttributes struct {
	ref terra.Reference
}

// BlockInheritance returns a reference to field block_inheritance of googleworkspace_org_unit.
func (ou orgUnitAttributes) BlockInheritance() terra.BoolValue {
	return terra.ReferenceAsBool(ou.ref.Append("block_inheritance"))
}

// Description returns a reference to field description of googleworkspace_org_unit.
func (ou orgUnitAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("description"))
}

// Etag returns a reference to field etag of googleworkspace_org_unit.
func (ou orgUnitAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_org_unit.
func (ou orgUnitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("id"))
}

// Name returns a reference to field name of googleworkspace_org_unit.
func (ou orgUnitAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("name"))
}

// OrgUnitId returns a reference to field org_unit_id of googleworkspace_org_unit.
func (ou orgUnitAttributes) OrgUnitId() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("org_unit_id"))
}

// OrgUnitPath returns a reference to field org_unit_path of googleworkspace_org_unit.
func (ou orgUnitAttributes) OrgUnitPath() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("org_unit_path"))
}

// ParentOrgUnitId returns a reference to field parent_org_unit_id of googleworkspace_org_unit.
func (ou orgUnitAttributes) ParentOrgUnitId() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("parent_org_unit_id"))
}

// ParentOrgUnitPath returns a reference to field parent_org_unit_path of googleworkspace_org_unit.
func (ou orgUnitAttributes) ParentOrgUnitPath() terra.StringValue {
	return terra.ReferenceAsString(ou.ref.Append("parent_org_unit_path"))
}

type orgUnitState struct {
	BlockInheritance  bool   `json:"block_inheritance"`
	Description       string `json:"description"`
	Etag              string `json:"etag"`
	Id                string `json:"id"`
	Name              string `json:"name"`
	OrgUnitId         string `json:"org_unit_id"`
	OrgUnitPath       string `json:"org_unit_path"`
	ParentOrgUnitId   string `json:"parent_org_unit_id"`
	ParentOrgUnitPath string `json:"parent_org_unit_path"`
}
