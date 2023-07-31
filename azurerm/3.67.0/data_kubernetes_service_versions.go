// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakubernetesserviceversions "github.com/golingon/terraproviders/azurerm/3.67.0/datakubernetesserviceversions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKubernetesServiceVersions creates a new instance of [DataKubernetesServiceVersions].
func NewDataKubernetesServiceVersions(name string, args DataKubernetesServiceVersionsArgs) *DataKubernetesServiceVersions {
	return &DataKubernetesServiceVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesServiceVersions)(nil)

// DataKubernetesServiceVersions represents the Terraform data resource azurerm_kubernetes_service_versions.
type DataKubernetesServiceVersions struct {
	Name string
	Args DataKubernetesServiceVersionsArgs
}

// DataSource returns the Terraform object type for [DataKubernetesServiceVersions].
func (ksv *DataKubernetesServiceVersions) DataSource() string {
	return "azurerm_kubernetes_service_versions"
}

// LocalName returns the local name for [DataKubernetesServiceVersions].
func (ksv *DataKubernetesServiceVersions) LocalName() string {
	return ksv.Name
}

// Configuration returns the configuration (args) for [DataKubernetesServiceVersions].
func (ksv *DataKubernetesServiceVersions) Configuration() interface{} {
	return ksv.Args
}

// Attributes returns the attributes for [DataKubernetesServiceVersions].
func (ksv *DataKubernetesServiceVersions) Attributes() dataKubernetesServiceVersionsAttributes {
	return dataKubernetesServiceVersionsAttributes{ref: terra.ReferenceDataResource(ksv)}
}

// DataKubernetesServiceVersionsArgs contains the configurations for azurerm_kubernetes_service_versions.
type DataKubernetesServiceVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludePreview: bool, optional
	IncludePreview terra.BoolValue `hcl:"include_preview,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// VersionPrefix: string, optional
	VersionPrefix terra.StringValue `hcl:"version_prefix,attr"`
	// Timeouts: optional
	Timeouts *datakubernetesserviceversions.Timeouts `hcl:"timeouts,block"`
}
type dataKubernetesServiceVersionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("id"))
}

// IncludePreview returns a reference to field include_preview of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) IncludePreview() terra.BoolValue {
	return terra.ReferenceAsBool(ksv.ref.Append("include_preview"))
}

// LatestVersion returns a reference to field latest_version of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) LatestVersion() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("latest_version"))
}

// Location returns a reference to field location of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("location"))
}

// VersionPrefix returns a reference to field version_prefix of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) VersionPrefix() terra.StringValue {
	return terra.ReferenceAsString(ksv.ref.Append("version_prefix"))
}

// Versions returns a reference to field versions of azurerm_kubernetes_service_versions.
func (ksv dataKubernetesServiceVersionsAttributes) Versions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ksv.ref.Append("versions"))
}

func (ksv dataKubernetesServiceVersionsAttributes) Timeouts() datakubernetesserviceversions.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakubernetesserviceversions.TimeoutsAttributes](ksv.ref.Append("timeouts"))
}
