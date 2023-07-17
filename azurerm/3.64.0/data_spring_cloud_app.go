// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataspringcloudapp "github.com/golingon/terraproviders/azurerm/3.64.0/dataspringcloudapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSpringCloudApp creates a new instance of [DataSpringCloudApp].
func NewDataSpringCloudApp(name string, args DataSpringCloudAppArgs) *DataSpringCloudApp {
	return &DataSpringCloudApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSpringCloudApp)(nil)

// DataSpringCloudApp represents the Terraform data resource azurerm_spring_cloud_app.
type DataSpringCloudApp struct {
	Name string
	Args DataSpringCloudAppArgs
}

// DataSource returns the Terraform object type for [DataSpringCloudApp].
func (sca *DataSpringCloudApp) DataSource() string {
	return "azurerm_spring_cloud_app"
}

// LocalName returns the local name for [DataSpringCloudApp].
func (sca *DataSpringCloudApp) LocalName() string {
	return sca.Name
}

// Configuration returns the configuration (args) for [DataSpringCloudApp].
func (sca *DataSpringCloudApp) Configuration() interface{} {
	return sca.Args
}

// Attributes returns the attributes for [DataSpringCloudApp].
func (sca *DataSpringCloudApp) Attributes() dataSpringCloudAppAttributes {
	return dataSpringCloudAppAttributes{ref: terra.ReferenceDataResource(sca)}
}

// DataSpringCloudAppArgs contains the configurations for azurerm_spring_cloud_app.
type DataSpringCloudAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Identity: min=0
	Identity []dataspringcloudapp.Identity `hcl:"identity,block" validate:"min=0"`
	// PersistentDisk: min=0
	PersistentDisk []dataspringcloudapp.PersistentDisk `hcl:"persistent_disk,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataspringcloudapp.Timeouts `hcl:"timeouts,block"`
}
type dataSpringCloudAppAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("fqdn"))
}

// HttpsOnly returns a reference to field https_only of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("id"))
}

// IsPublic returns a reference to field is_public of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) IsPublic() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("is_public"))
}

// Name returns a reference to field name of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("resource_group_name"))
}

// ServiceName returns a reference to field service_name of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("service_name"))
}

// TlsEnabled returns a reference to field tls_enabled of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("tls_enabled"))
}

// Url returns a reference to field url of azurerm_spring_cloud_app.
func (sca dataSpringCloudAppAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("url"))
}

func (sca dataSpringCloudAppAttributes) Identity() terra.ListValue[dataspringcloudapp.IdentityAttributes] {
	return terra.ReferenceAsList[dataspringcloudapp.IdentityAttributes](sca.ref.Append("identity"))
}

func (sca dataSpringCloudAppAttributes) PersistentDisk() terra.ListValue[dataspringcloudapp.PersistentDiskAttributes] {
	return terra.ReferenceAsList[dataspringcloudapp.PersistentDiskAttributes](sca.ref.Append("persistent_disk"))
}

func (sca dataSpringCloudAppAttributes) Timeouts() dataspringcloudapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataspringcloudapp.TimeoutsAttributes](sca.ref.Append("timeouts"))
}