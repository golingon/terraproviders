// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vmwareengine_cluster

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_vmwareengine_cluster.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gvc *DataSource) DataSource() string {
	return "google_vmwareengine_cluster"
}

// LocalName returns the local name for [DataSource].
func (gvc *DataSource) LocalName() string {
	return gvc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gvc *DataSource) Configuration() interface{} {
	return gvc.Args
}

// Attributes returns the attributes for [DataSource].
func (gvc *DataSource) Attributes() dataGoogleVmwareengineClusterAttributes {
	return dataGoogleVmwareengineClusterAttributes{ref: terra.ReferenceDataSource(gvc)}
}

// DataArgs contains the configurations for google_vmwareengine_cluster.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
}

type dataGoogleVmwareengineClusterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvc.ref.Append("id"))
}

// Management returns a reference to field management of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) Management() terra.BoolValue {
	return terra.ReferenceAsBool(gvc.ref.Append("management"))
}

// Name returns a reference to field name of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvc.ref.Append("name"))
}

// Parent returns a reference to field parent of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(gvc.ref.Append("parent"))
}

// State returns a reference to field state of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gvc.ref.Append("state"))
}

// Uid returns a reference to field uid of google_vmwareengine_cluster.
func (gvc dataGoogleVmwareengineClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gvc.ref.Append("uid"))
}

func (gvc dataGoogleVmwareengineClusterAttributes) NodeTypeConfigs() terra.SetValue[DataNodeTypeConfigsAttributes] {
	return terra.ReferenceAsSet[DataNodeTypeConfigsAttributes](gvc.ref.Append("node_type_configs"))
}
