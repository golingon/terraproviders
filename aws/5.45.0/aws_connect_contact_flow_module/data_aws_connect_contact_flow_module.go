// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_contact_flow_module

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_connect_contact_flow_module.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (accfm *DataSource) DataSource() string {
	return "aws_connect_contact_flow_module"
}

// LocalName returns the local name for [DataSource].
func (accfm *DataSource) LocalName() string {
	return accfm.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (accfm *DataSource) Configuration() interface{} {
	return accfm.Args
}

// Attributes returns the attributes for [DataSource].
func (accfm *DataSource) Attributes() dataAwsConnectContactFlowModuleAttributes {
	return dataAwsConnectContactFlowModuleAttributes{ref: terra.ReferenceDataSource(accfm)}
}

// DataArgs contains the configurations for aws_connect_contact_flow_module.
type DataArgs struct {
	// ContactFlowModuleId: string, optional
	ContactFlowModuleId terra.StringValue `hcl:"contact_flow_module_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsConnectContactFlowModuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("arn"))
}

// ContactFlowModuleId returns a reference to field contact_flow_module_id of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) ContactFlowModuleId() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("contact_flow_module_id"))
}

// Content returns a reference to field content of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("content"))
}

// Description returns a reference to field description of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("name"))
}

// State returns a reference to field state of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("state"))
}

// Status returns a reference to field status of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(accfm.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_connect_contact_flow_module.
func (accfm dataAwsConnectContactFlowModuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](accfm.ref.Append("tags"))
}
