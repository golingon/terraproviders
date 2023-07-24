// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lakeformationpermissions "github.com/golingon/terraproviders/aws/5.9.0/lakeformationpermissions"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLakeformationPermissions creates a new instance of [LakeformationPermissions].
func NewLakeformationPermissions(name string, args LakeformationPermissionsArgs) *LakeformationPermissions {
	return &LakeformationPermissions{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LakeformationPermissions)(nil)

// LakeformationPermissions represents the Terraform resource aws_lakeformation_permissions.
type LakeformationPermissions struct {
	Name      string
	Args      LakeformationPermissionsArgs
	state     *lakeformationPermissionsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LakeformationPermissions].
func (lp *LakeformationPermissions) Type() string {
	return "aws_lakeformation_permissions"
}

// LocalName returns the local name for [LakeformationPermissions].
func (lp *LakeformationPermissions) LocalName() string {
	return lp.Name
}

// Configuration returns the configuration (args) for [LakeformationPermissions].
func (lp *LakeformationPermissions) Configuration() interface{} {
	return lp.Args
}

// DependOn is used for other resources to depend on [LakeformationPermissions].
func (lp *LakeformationPermissions) DependOn() terra.Reference {
	return terra.ReferenceResource(lp)
}

// Dependencies returns the list of resources [LakeformationPermissions] depends_on.
func (lp *LakeformationPermissions) Dependencies() terra.Dependencies {
	return lp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LakeformationPermissions].
func (lp *LakeformationPermissions) LifecycleManagement() *terra.Lifecycle {
	return lp.Lifecycle
}

// Attributes returns the attributes for [LakeformationPermissions].
func (lp *LakeformationPermissions) Attributes() lakeformationPermissionsAttributes {
	return lakeformationPermissionsAttributes{ref: terra.ReferenceResource(lp)}
}

// ImportState imports the given attribute values into [LakeformationPermissions]'s state.
func (lp *LakeformationPermissions) ImportState(av io.Reader) error {
	lp.state = &lakeformationPermissionsState{}
	if err := json.NewDecoder(av).Decode(lp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lp.Type(), lp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LakeformationPermissions] has state.
func (lp *LakeformationPermissions) State() (*lakeformationPermissionsState, bool) {
	return lp.state, lp.state != nil
}

// StateMust returns the state for [LakeformationPermissions]. Panics if the state is nil.
func (lp *LakeformationPermissions) StateMust() *lakeformationPermissionsState {
	if lp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lp.Type(), lp.LocalName()))
	}
	return lp.state
}

// LakeformationPermissionsArgs contains the configurations for aws_lakeformation_permissions.
type LakeformationPermissionsArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// CatalogResource: bool, optional
	CatalogResource terra.BoolValue `hcl:"catalog_resource,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Permissions: list of string, required
	Permissions terra.ListValue[terra.StringValue] `hcl:"permissions,attr" validate:"required"`
	// PermissionsWithGrantOption: list of string, optional
	PermissionsWithGrantOption terra.ListValue[terra.StringValue] `hcl:"permissions_with_grant_option,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// DataLocation: optional
	DataLocation *lakeformationpermissions.DataLocation `hcl:"data_location,block"`
	// Database: optional
	Database *lakeformationpermissions.Database `hcl:"database,block"`
	// LfTag: optional
	LfTag *lakeformationpermissions.LfTag `hcl:"lf_tag,block"`
	// LfTagPolicy: optional
	LfTagPolicy *lakeformationpermissions.LfTagPolicy `hcl:"lf_tag_policy,block"`
	// Table: optional
	Table *lakeformationpermissions.Table `hcl:"table,block"`
	// TableWithColumns: optional
	TableWithColumns *lakeformationpermissions.TableWithColumns `hcl:"table_with_columns,block"`
}
type lakeformationPermissionsAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("catalog_id"))
}

// CatalogResource returns a reference to field catalog_resource of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) CatalogResource() terra.BoolValue {
	return terra.ReferenceAsBool(lp.ref.Append("catalog_resource"))
}

// Id returns a reference to field id of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("id"))
}

// Permissions returns a reference to field permissions of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) Permissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lp.ref.Append("permissions"))
}

// PermissionsWithGrantOption returns a reference to field permissions_with_grant_option of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) PermissionsWithGrantOption() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lp.ref.Append("permissions_with_grant_option"))
}

// Principal returns a reference to field principal of aws_lakeformation_permissions.
func (lp lakeformationPermissionsAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("principal"))
}

func (lp lakeformationPermissionsAttributes) DataLocation() terra.ListValue[lakeformationpermissions.DataLocationAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.DataLocationAttributes](lp.ref.Append("data_location"))
}

func (lp lakeformationPermissionsAttributes) Database() terra.ListValue[lakeformationpermissions.DatabaseAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.DatabaseAttributes](lp.ref.Append("database"))
}

func (lp lakeformationPermissionsAttributes) LfTag() terra.ListValue[lakeformationpermissions.LfTagAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.LfTagAttributes](lp.ref.Append("lf_tag"))
}

func (lp lakeformationPermissionsAttributes) LfTagPolicy() terra.ListValue[lakeformationpermissions.LfTagPolicyAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.LfTagPolicyAttributes](lp.ref.Append("lf_tag_policy"))
}

func (lp lakeformationPermissionsAttributes) Table() terra.ListValue[lakeformationpermissions.TableAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.TableAttributes](lp.ref.Append("table"))
}

func (lp lakeformationPermissionsAttributes) TableWithColumns() terra.ListValue[lakeformationpermissions.TableWithColumnsAttributes] {
	return terra.ReferenceAsList[lakeformationpermissions.TableWithColumnsAttributes](lp.ref.Append("table_with_columns"))
}

type lakeformationPermissionsState struct {
	CatalogId                  string                                           `json:"catalog_id"`
	CatalogResource            bool                                             `json:"catalog_resource"`
	Id                         string                                           `json:"id"`
	Permissions                []string                                         `json:"permissions"`
	PermissionsWithGrantOption []string                                         `json:"permissions_with_grant_option"`
	Principal                  string                                           `json:"principal"`
	DataLocation               []lakeformationpermissions.DataLocationState     `json:"data_location"`
	Database                   []lakeformationpermissions.DatabaseState         `json:"database"`
	LfTag                      []lakeformationpermissions.LfTagState            `json:"lf_tag"`
	LfTagPolicy                []lakeformationpermissions.LfTagPolicyState      `json:"lf_tag_policy"`
	Table                      []lakeformationpermissions.TableState            `json:"table"`
	TableWithColumns           []lakeformationpermissions.TableWithColumnsState `json:"table_with_columns"`
}
