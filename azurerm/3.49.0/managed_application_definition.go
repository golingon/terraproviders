// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managedapplicationdefinition "github.com/golingon/terraproviders/azurerm/3.49.0/managedapplicationdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewManagedApplicationDefinition(name string, args ManagedApplicationDefinitionArgs) *ManagedApplicationDefinition {
	return &ManagedApplicationDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagedApplicationDefinition)(nil)

type ManagedApplicationDefinition struct {
	Name  string
	Args  ManagedApplicationDefinitionArgs
	state *managedApplicationDefinitionState
}

func (mad *ManagedApplicationDefinition) Type() string {
	return "azurerm_managed_application_definition"
}

func (mad *ManagedApplicationDefinition) LocalName() string {
	return mad.Name
}

func (mad *ManagedApplicationDefinition) Configuration() interface{} {
	return mad.Args
}

func (mad *ManagedApplicationDefinition) Attributes() managedApplicationDefinitionAttributes {
	return managedApplicationDefinitionAttributes{ref: terra.ReferenceResource(mad)}
}

func (mad *ManagedApplicationDefinition) ImportState(av io.Reader) error {
	mad.state = &managedApplicationDefinitionState{}
	if err := json.NewDecoder(av).Decode(mad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mad.Type(), mad.LocalName(), err)
	}
	return nil
}

func (mad *ManagedApplicationDefinition) State() (*managedApplicationDefinitionState, bool) {
	return mad.state, mad.state != nil
}

func (mad *ManagedApplicationDefinition) StateMust() *managedApplicationDefinitionState {
	if mad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mad.Type(), mad.LocalName()))
	}
	return mad.state
}

func (mad *ManagedApplicationDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(mad)
}

type ManagedApplicationDefinitionArgs struct {
	// CreateUiDefinition: string, optional
	CreateUiDefinition terra.StringValue `hcl:"create_ui_definition,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LockLevel: string, required
	LockLevel terra.StringValue `hcl:"lock_level,attr" validate:"required"`
	// MainTemplate: string, optional
	MainTemplate terra.StringValue `hcl:"main_template,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PackageEnabled: bool, optional
	PackageEnabled terra.BoolValue `hcl:"package_enabled,attr"`
	// PackageFileUri: string, optional
	PackageFileUri terra.StringValue `hcl:"package_file_uri,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Authorization: min=0
	Authorization []managedapplicationdefinition.Authorization `hcl:"authorization,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *managedapplicationdefinition.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ManagedApplicationDefinition depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type managedApplicationDefinitionAttributes struct {
	ref terra.Reference
}

func (mad managedApplicationDefinitionAttributes) CreateUiDefinition() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("create_ui_definition"))
}

func (mad managedApplicationDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("description"))
}

func (mad managedApplicationDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("display_name"))
}

func (mad managedApplicationDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("id"))
}

func (mad managedApplicationDefinitionAttributes) Location() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("location"))
}

func (mad managedApplicationDefinitionAttributes) LockLevel() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("lock_level"))
}

func (mad managedApplicationDefinitionAttributes) MainTemplate() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("main_template"))
}

func (mad managedApplicationDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("name"))
}

func (mad managedApplicationDefinitionAttributes) PackageEnabled() terra.BoolValue {
	return terra.ReferenceBool(mad.ref.Append("package_enabled"))
}

func (mad managedApplicationDefinitionAttributes) PackageFileUri() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("package_file_uri"))
}

func (mad managedApplicationDefinitionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mad.ref.Append("resource_group_name"))
}

func (mad managedApplicationDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mad.ref.Append("tags"))
}

func (mad managedApplicationDefinitionAttributes) Authorization() terra.SetValue[managedapplicationdefinition.AuthorizationAttributes] {
	return terra.ReferenceSet[managedapplicationdefinition.AuthorizationAttributes](mad.ref.Append("authorization"))
}

func (mad managedApplicationDefinitionAttributes) Timeouts() managedapplicationdefinition.TimeoutsAttributes {
	return terra.ReferenceSingle[managedapplicationdefinition.TimeoutsAttributes](mad.ref.Append("timeouts"))
}

type managedApplicationDefinitionState struct {
	CreateUiDefinition string                                            `json:"create_ui_definition"`
	Description        string                                            `json:"description"`
	DisplayName        string                                            `json:"display_name"`
	Id                 string                                            `json:"id"`
	Location           string                                            `json:"location"`
	LockLevel          string                                            `json:"lock_level"`
	MainTemplate       string                                            `json:"main_template"`
	Name               string                                            `json:"name"`
	PackageEnabled     bool                                              `json:"package_enabled"`
	PackageFileUri     string                                            `json:"package_file_uri"`
	ResourceGroupName  string                                            `json:"resource_group_name"`
	Tags               map[string]string                                 `json:"tags"`
	Authorization      []managedapplicationdefinition.AuthorizationState `json:"authorization"`
	Timeouts           *managedapplicationdefinition.TimeoutsState       `json:"timeouts"`
}
