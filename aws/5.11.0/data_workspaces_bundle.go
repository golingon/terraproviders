// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataworkspacesbundle "github.com/golingon/terraproviders/aws/5.11.0/dataworkspacesbundle"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWorkspacesBundle creates a new instance of [DataWorkspacesBundle].
func NewDataWorkspacesBundle(name string, args DataWorkspacesBundleArgs) *DataWorkspacesBundle {
	return &DataWorkspacesBundle{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWorkspacesBundle)(nil)

// DataWorkspacesBundle represents the Terraform data resource aws_workspaces_bundle.
type DataWorkspacesBundle struct {
	Name string
	Args DataWorkspacesBundleArgs
}

// DataSource returns the Terraform object type for [DataWorkspacesBundle].
func (wb *DataWorkspacesBundle) DataSource() string {
	return "aws_workspaces_bundle"
}

// LocalName returns the local name for [DataWorkspacesBundle].
func (wb *DataWorkspacesBundle) LocalName() string {
	return wb.Name
}

// Configuration returns the configuration (args) for [DataWorkspacesBundle].
func (wb *DataWorkspacesBundle) Configuration() interface{} {
	return wb.Args
}

// Attributes returns the attributes for [DataWorkspacesBundle].
func (wb *DataWorkspacesBundle) Attributes() dataWorkspacesBundleAttributes {
	return dataWorkspacesBundleAttributes{ref: terra.ReferenceDataResource(wb)}
}

// DataWorkspacesBundleArgs contains the configurations for aws_workspaces_bundle.
type DataWorkspacesBundleArgs struct {
	// BundleId: string, optional
	BundleId terra.StringValue `hcl:"bundle_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// ComputeType: min=0
	ComputeType []dataworkspacesbundle.ComputeType `hcl:"compute_type,block" validate:"min=0"`
	// RootStorage: min=0
	RootStorage []dataworkspacesbundle.RootStorage `hcl:"root_storage,block" validate:"min=0"`
	// UserStorage: min=0
	UserStorage []dataworkspacesbundle.UserStorage `hcl:"user_storage,block" validate:"min=0"`
}
type dataWorkspacesBundleAttributes struct {
	ref terra.Reference
}

// BundleId returns a reference to field bundle_id of aws_workspaces_bundle.
func (wb dataWorkspacesBundleAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(wb.ref.Append("bundle_id"))
}

// Description returns a reference to field description of aws_workspaces_bundle.
func (wb dataWorkspacesBundleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wb.ref.Append("description"))
}

// Id returns a reference to field id of aws_workspaces_bundle.
func (wb dataWorkspacesBundleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wb.ref.Append("id"))
}

// Name returns a reference to field name of aws_workspaces_bundle.
func (wb dataWorkspacesBundleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wb.ref.Append("name"))
}

// Owner returns a reference to field owner of aws_workspaces_bundle.
func (wb dataWorkspacesBundleAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(wb.ref.Append("owner"))
}

func (wb dataWorkspacesBundleAttributes) ComputeType() terra.ListValue[dataworkspacesbundle.ComputeTypeAttributes] {
	return terra.ReferenceAsList[dataworkspacesbundle.ComputeTypeAttributes](wb.ref.Append("compute_type"))
}

func (wb dataWorkspacesBundleAttributes) RootStorage() terra.ListValue[dataworkspacesbundle.RootStorageAttributes] {
	return terra.ReferenceAsList[dataworkspacesbundle.RootStorageAttributes](wb.ref.Append("root_storage"))
}

func (wb dataWorkspacesBundleAttributes) UserStorage() terra.ListValue[dataworkspacesbundle.UserStorageAttributes] {
	return terra.ReferenceAsList[dataworkspacesbundle.UserStorageAttributes](wb.ref.Append("user_storage"))
}
