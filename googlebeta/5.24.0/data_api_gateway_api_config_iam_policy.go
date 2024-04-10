// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataApiGatewayApiConfigIamPolicy creates a new instance of [DataApiGatewayApiConfigIamPolicy].
func NewDataApiGatewayApiConfigIamPolicy(name string, args DataApiGatewayApiConfigIamPolicyArgs) *DataApiGatewayApiConfigIamPolicy {
	return &DataApiGatewayApiConfigIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiGatewayApiConfigIamPolicy)(nil)

// DataApiGatewayApiConfigIamPolicy represents the Terraform data resource google_api_gateway_api_config_iam_policy.
type DataApiGatewayApiConfigIamPolicy struct {
	Name string
	Args DataApiGatewayApiConfigIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataApiGatewayApiConfigIamPolicy].
func (agacip *DataApiGatewayApiConfigIamPolicy) DataSource() string {
	return "google_api_gateway_api_config_iam_policy"
}

// LocalName returns the local name for [DataApiGatewayApiConfigIamPolicy].
func (agacip *DataApiGatewayApiConfigIamPolicy) LocalName() string {
	return agacip.Name
}

// Configuration returns the configuration (args) for [DataApiGatewayApiConfigIamPolicy].
func (agacip *DataApiGatewayApiConfigIamPolicy) Configuration() interface{} {
	return agacip.Args
}

// Attributes returns the attributes for [DataApiGatewayApiConfigIamPolicy].
func (agacip *DataApiGatewayApiConfigIamPolicy) Attributes() dataApiGatewayApiConfigIamPolicyAttributes {
	return dataApiGatewayApiConfigIamPolicyAttributes{ref: terra.ReferenceDataResource(agacip)}
}

// DataApiGatewayApiConfigIamPolicyArgs contains the configurations for google_api_gateway_api_config_iam_policy.
type DataApiGatewayApiConfigIamPolicyArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// ApiConfig: string, required
	ApiConfig terra.StringValue `hcl:"api_config,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataApiGatewayApiConfigIamPolicyAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("api"))
}

// ApiConfig returns a reference to field api_config of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) ApiConfig() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("api_config"))
}

// Etag returns a reference to field etag of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_api_gateway_api_config_iam_policy.
func (agacip dataApiGatewayApiConfigIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agacip.ref.Append("project"))
}
