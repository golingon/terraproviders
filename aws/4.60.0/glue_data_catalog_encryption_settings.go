// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluedatacatalogencryptionsettings "github.com/golingon/terraproviders/aws/4.60.0/gluedatacatalogencryptionsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueDataCatalogEncryptionSettings creates a new instance of [GlueDataCatalogEncryptionSettings].
func NewGlueDataCatalogEncryptionSettings(name string, args GlueDataCatalogEncryptionSettingsArgs) *GlueDataCatalogEncryptionSettings {
	return &GlueDataCatalogEncryptionSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueDataCatalogEncryptionSettings)(nil)

// GlueDataCatalogEncryptionSettings represents the Terraform resource aws_glue_data_catalog_encryption_settings.
type GlueDataCatalogEncryptionSettings struct {
	Name      string
	Args      GlueDataCatalogEncryptionSettingsArgs
	state     *glueDataCatalogEncryptionSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) Type() string {
	return "aws_glue_data_catalog_encryption_settings"
}

// LocalName returns the local name for [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) LocalName() string {
	return gdces.Name
}

// Configuration returns the configuration (args) for [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) Configuration() interface{} {
	return gdces.Args
}

// DependOn is used for other resources to depend on [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(gdces)
}

// Dependencies returns the list of resources [GlueDataCatalogEncryptionSettings] depends_on.
func (gdces *GlueDataCatalogEncryptionSettings) Dependencies() terra.Dependencies {
	return gdces.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) LifecycleManagement() *terra.Lifecycle {
	return gdces.Lifecycle
}

// Attributes returns the attributes for [GlueDataCatalogEncryptionSettings].
func (gdces *GlueDataCatalogEncryptionSettings) Attributes() glueDataCatalogEncryptionSettingsAttributes {
	return glueDataCatalogEncryptionSettingsAttributes{ref: terra.ReferenceResource(gdces)}
}

// ImportState imports the given attribute values into [GlueDataCatalogEncryptionSettings]'s state.
func (gdces *GlueDataCatalogEncryptionSettings) ImportState(av io.Reader) error {
	gdces.state = &glueDataCatalogEncryptionSettingsState{}
	if err := json.NewDecoder(av).Decode(gdces.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdces.Type(), gdces.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueDataCatalogEncryptionSettings] has state.
func (gdces *GlueDataCatalogEncryptionSettings) State() (*glueDataCatalogEncryptionSettingsState, bool) {
	return gdces.state, gdces.state != nil
}

// StateMust returns the state for [GlueDataCatalogEncryptionSettings]. Panics if the state is nil.
func (gdces *GlueDataCatalogEncryptionSettings) StateMust() *glueDataCatalogEncryptionSettingsState {
	if gdces.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdces.Type(), gdces.LocalName()))
	}
	return gdces.state
}

// GlueDataCatalogEncryptionSettingsArgs contains the configurations for aws_glue_data_catalog_encryption_settings.
type GlueDataCatalogEncryptionSettingsArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DataCatalogEncryptionSettings: required
	DataCatalogEncryptionSettings *gluedatacatalogencryptionsettings.DataCatalogEncryptionSettings `hcl:"data_catalog_encryption_settings,block" validate:"required"`
}
type glueDataCatalogEncryptionSettingsAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_glue_data_catalog_encryption_settings.
func (gdces glueDataCatalogEncryptionSettingsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gdces.ref.Append("catalog_id"))
}

// Id returns a reference to field id of aws_glue_data_catalog_encryption_settings.
func (gdces glueDataCatalogEncryptionSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdces.ref.Append("id"))
}

func (gdces glueDataCatalogEncryptionSettingsAttributes) DataCatalogEncryptionSettings() terra.ListValue[gluedatacatalogencryptionsettings.DataCatalogEncryptionSettingsAttributes] {
	return terra.ReferenceAsList[gluedatacatalogencryptionsettings.DataCatalogEncryptionSettingsAttributes](gdces.ref.Append("data_catalog_encryption_settings"))
}

type glueDataCatalogEncryptionSettingsState struct {
	CatalogId                     string                                                                 `json:"catalog_id"`
	Id                            string                                                                 `json:"id"`
	DataCatalogEncryptionSettings []gluedatacatalogencryptionsettings.DataCatalogEncryptionSettingsState `json:"data_catalog_encryption_settings"`
}
