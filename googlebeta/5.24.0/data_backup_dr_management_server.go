// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"github.com/golingon/lingon/pkg/terra"
	databackupdrmanagementserver "github.com/golingon/terraproviders/googlebeta/5.24.0/databackupdrmanagementserver"
)

// NewDataBackupDrManagementServer creates a new instance of [DataBackupDrManagementServer].
func NewDataBackupDrManagementServer(name string, args DataBackupDrManagementServerArgs) *DataBackupDrManagementServer {
	return &DataBackupDrManagementServer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBackupDrManagementServer)(nil)

// DataBackupDrManagementServer represents the Terraform data resource google_backup_dr_management_server.
type DataBackupDrManagementServer struct {
	Name string
	Args DataBackupDrManagementServerArgs
}

// DataSource returns the Terraform object type for [DataBackupDrManagementServer].
func (bdms *DataBackupDrManagementServer) DataSource() string {
	return "google_backup_dr_management_server"
}

// LocalName returns the local name for [DataBackupDrManagementServer].
func (bdms *DataBackupDrManagementServer) LocalName() string {
	return bdms.Name
}

// Configuration returns the configuration (args) for [DataBackupDrManagementServer].
func (bdms *DataBackupDrManagementServer) Configuration() interface{} {
	return bdms.Args
}

// Attributes returns the attributes for [DataBackupDrManagementServer].
func (bdms *DataBackupDrManagementServer) Attributes() dataBackupDrManagementServerAttributes {
	return dataBackupDrManagementServerAttributes{ref: terra.ReferenceDataResource(bdms)}
}

// DataBackupDrManagementServerArgs contains the configurations for google_backup_dr_management_server.
type DataBackupDrManagementServerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagementUri: min=0
	ManagementUri []databackupdrmanagementserver.ManagementUri `hcl:"management_uri,block" validate:"min=0"`
	// Networks: min=0
	Networks []databackupdrmanagementserver.Networks `hcl:"networks,block" validate:"min=0"`
}
type dataBackupDrManagementServerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("id"))
}

// Location returns a reference to field location of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("location"))
}

// Name returns a reference to field name of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("name"))
}

// Oauth2ClientId returns a reference to field oauth2_client_id of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Oauth2ClientId() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("oauth2_client_id"))
}

// Project returns a reference to field project of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("project"))
}

// Type returns a reference to field type of google_backup_dr_management_server.
func (bdms dataBackupDrManagementServerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bdms.ref.Append("type"))
}

func (bdms dataBackupDrManagementServerAttributes) ManagementUri() terra.ListValue[databackupdrmanagementserver.ManagementUriAttributes] {
	return terra.ReferenceAsList[databackupdrmanagementserver.ManagementUriAttributes](bdms.ref.Append("management_uri"))
}

func (bdms dataBackupDrManagementServerAttributes) Networks() terra.ListValue[databackupdrmanagementserver.NetworksAttributes] {
	return terra.ReferenceAsList[databackupdrmanagementserver.NetworksAttributes](bdms.ref.Append("networks"))
}
