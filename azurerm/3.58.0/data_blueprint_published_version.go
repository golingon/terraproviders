// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datablueprintpublishedversion "github.com/golingon/terraproviders/azurerm/3.58.0/datablueprintpublishedversion"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBlueprintPublishedVersion creates a new instance of [DataBlueprintPublishedVersion].
func NewDataBlueprintPublishedVersion(name string, args DataBlueprintPublishedVersionArgs) *DataBlueprintPublishedVersion {
	return &DataBlueprintPublishedVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBlueprintPublishedVersion)(nil)

// DataBlueprintPublishedVersion represents the Terraform data resource azurerm_blueprint_published_version.
type DataBlueprintPublishedVersion struct {
	Name string
	Args DataBlueprintPublishedVersionArgs
}

// DataSource returns the Terraform object type for [DataBlueprintPublishedVersion].
func (bpv *DataBlueprintPublishedVersion) DataSource() string {
	return "azurerm_blueprint_published_version"
}

// LocalName returns the local name for [DataBlueprintPublishedVersion].
func (bpv *DataBlueprintPublishedVersion) LocalName() string {
	return bpv.Name
}

// Configuration returns the configuration (args) for [DataBlueprintPublishedVersion].
func (bpv *DataBlueprintPublishedVersion) Configuration() interface{} {
	return bpv.Args
}

// Attributes returns the attributes for [DataBlueprintPublishedVersion].
func (bpv *DataBlueprintPublishedVersion) Attributes() dataBlueprintPublishedVersionAttributes {
	return dataBlueprintPublishedVersionAttributes{ref: terra.ReferenceDataResource(bpv)}
}

// DataBlueprintPublishedVersionArgs contains the configurations for azurerm_blueprint_published_version.
type DataBlueprintPublishedVersionArgs struct {
	// BlueprintName: string, required
	BlueprintName terra.StringValue `hcl:"blueprint_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ScopeId: string, required
	ScopeId terra.StringValue `hcl:"scope_id,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datablueprintpublishedversion.Timeouts `hcl:"timeouts,block"`
}
type dataBlueprintPublishedVersionAttributes struct {
	ref terra.Reference
}

// BlueprintName returns a reference to field blueprint_name of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) BlueprintName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("blueprint_name"))
}

// Description returns a reference to field description of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("id"))
}

// LastModified returns a reference to field last_modified of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("last_modified"))
}

// ScopeId returns a reference to field scope_id of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("scope_id"))
}

// TargetScope returns a reference to field target_scope of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) TargetScope() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("target_scope"))
}

// TimeCreated returns a reference to field time_created of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("time_created"))
}

// Type returns a reference to field type of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("type"))
}

// Version returns a reference to field version of azurerm_blueprint_published_version.
func (bpv dataBlueprintPublishedVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("version"))
}

func (bpv dataBlueprintPublishedVersionAttributes) Timeouts() datablueprintpublishedversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datablueprintpublishedversion.TimeoutsAttributes](bpv.ref.Append("timeouts"))
}