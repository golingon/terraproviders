// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataopensearchserverlesssecurityconfig "github.com/golingon/terraproviders/aws/5.10.0/dataopensearchserverlesssecurityconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOpensearchserverlessSecurityConfig creates a new instance of [DataOpensearchserverlessSecurityConfig].
func NewDataOpensearchserverlessSecurityConfig(name string, args DataOpensearchserverlessSecurityConfigArgs) *DataOpensearchserverlessSecurityConfig {
	return &DataOpensearchserverlessSecurityConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOpensearchserverlessSecurityConfig)(nil)

// DataOpensearchserverlessSecurityConfig represents the Terraform data resource aws_opensearchserverless_security_config.
type DataOpensearchserverlessSecurityConfig struct {
	Name string
	Args DataOpensearchserverlessSecurityConfigArgs
}

// DataSource returns the Terraform object type for [DataOpensearchserverlessSecurityConfig].
func (osc *DataOpensearchserverlessSecurityConfig) DataSource() string {
	return "aws_opensearchserverless_security_config"
}

// LocalName returns the local name for [DataOpensearchserverlessSecurityConfig].
func (osc *DataOpensearchserverlessSecurityConfig) LocalName() string {
	return osc.Name
}

// Configuration returns the configuration (args) for [DataOpensearchserverlessSecurityConfig].
func (osc *DataOpensearchserverlessSecurityConfig) Configuration() interface{} {
	return osc.Args
}

// Attributes returns the attributes for [DataOpensearchserverlessSecurityConfig].
func (osc *DataOpensearchserverlessSecurityConfig) Attributes() dataOpensearchserverlessSecurityConfigAttributes {
	return dataOpensearchserverlessSecurityConfigAttributes{ref: terra.ReferenceDataResource(osc)}
}

// DataOpensearchserverlessSecurityConfigArgs contains the configurations for aws_opensearchserverless_security_config.
type DataOpensearchserverlessSecurityConfigArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// SamlOptions: optional
	SamlOptions *dataopensearchserverlesssecurityconfig.SamlOptions `hcl:"saml_options,block"`
}
type dataOpensearchserverlessSecurityConfigAttributes struct {
	ref terra.Reference
}

// ConfigVersion returns a reference to field config_version of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) ConfigVersion() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("config_version"))
}

// CreatedDate returns a reference to field created_date of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("description"))
}

// Id returns a reference to field id of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("id"))
}

// LastModifiedDate returns a reference to field last_modified_date of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) LastModifiedDate() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("last_modified_date"))
}

// Type returns a reference to field type of aws_opensearchserverless_security_config.
func (osc dataOpensearchserverlessSecurityConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(osc.ref.Append("type"))
}

func (osc dataOpensearchserverlessSecurityConfigAttributes) SamlOptions() dataopensearchserverlesssecurityconfig.SamlOptionsAttributes {
	return terra.ReferenceAsSingle[dataopensearchserverlesssecurityconfig.SamlOptionsAttributes](osc.ref.Append("saml_options"))
}