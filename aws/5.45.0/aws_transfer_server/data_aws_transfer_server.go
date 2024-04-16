// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_transfer_server

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_transfer_server.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ats *DataSource) DataSource() string {
	return "aws_transfer_server"
}

// LocalName returns the local name for [DataSource].
func (ats *DataSource) LocalName() string {
	return ats.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ats *DataSource) Configuration() interface{} {
	return ats.Args
}

// Attributes returns the attributes for [DataSource].
func (ats *DataSource) Attributes() dataAwsTransferServerAttributes {
	return dataAwsTransferServerAttributes{ref: terra.ReferenceDataSource(ats)}
}

// DataArgs contains the configurations for aws_transfer_server.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
}

type dataAwsTransferServerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("certificate"))
}

// Domain returns a reference to field domain of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("domain"))
}

// Endpoint returns a reference to field endpoint of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("endpoint"))
}

// EndpointType returns a reference to field endpoint_type of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("endpoint_type"))
}

// Id returns a reference to field id of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("id"))
}

// IdentityProviderType returns a reference to field identity_provider_type of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) IdentityProviderType() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("identity_provider_type"))
}

// InvocationRole returns a reference to field invocation_role of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) InvocationRole() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("invocation_role"))
}

// LoggingRole returns a reference to field logging_role of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) LoggingRole() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("logging_role"))
}

// Protocols returns a reference to field protocols of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ats.ref.Append("protocols"))
}

// SecurityPolicyName returns a reference to field security_policy_name of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) SecurityPolicyName() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("security_policy_name"))
}

// ServerId returns a reference to field server_id of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("server_id"))
}

// StructuredLogDestinations returns a reference to field structured_log_destinations of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) StructuredLogDestinations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ats.ref.Append("structured_log_destinations"))
}

// Url returns a reference to field url of aws_transfer_server.
func (ats dataAwsTransferServerAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ats.ref.Append("url"))
}
