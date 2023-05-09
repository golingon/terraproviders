// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudassetresourcessearchall "github.com/golingon/terraproviders/googlebeta/4.64.0/datacloudassetresourcessearchall"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudAssetResourcesSearchAll creates a new instance of [DataCloudAssetResourcesSearchAll].
func NewDataCloudAssetResourcesSearchAll(name string, args DataCloudAssetResourcesSearchAllArgs) *DataCloudAssetResourcesSearchAll {
	return &DataCloudAssetResourcesSearchAll{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudAssetResourcesSearchAll)(nil)

// DataCloudAssetResourcesSearchAll represents the Terraform data resource google_cloud_asset_resources_search_all.
type DataCloudAssetResourcesSearchAll struct {
	Name string
	Args DataCloudAssetResourcesSearchAllArgs
}

// DataSource returns the Terraform object type for [DataCloudAssetResourcesSearchAll].
func (carsa *DataCloudAssetResourcesSearchAll) DataSource() string {
	return "google_cloud_asset_resources_search_all"
}

// LocalName returns the local name for [DataCloudAssetResourcesSearchAll].
func (carsa *DataCloudAssetResourcesSearchAll) LocalName() string {
	return carsa.Name
}

// Configuration returns the configuration (args) for [DataCloudAssetResourcesSearchAll].
func (carsa *DataCloudAssetResourcesSearchAll) Configuration() interface{} {
	return carsa.Args
}

// Attributes returns the attributes for [DataCloudAssetResourcesSearchAll].
func (carsa *DataCloudAssetResourcesSearchAll) Attributes() dataCloudAssetResourcesSearchAllAttributes {
	return dataCloudAssetResourcesSearchAllAttributes{ref: terra.ReferenceDataResource(carsa)}
}

// DataCloudAssetResourcesSearchAllArgs contains the configurations for google_cloud_asset_resources_search_all.
type DataCloudAssetResourcesSearchAllArgs struct {
	// AssetTypes: list of string, optional
	AssetTypes terra.ListValue[terra.StringValue] `hcl:"asset_types,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Query: string, optional
	Query terra.StringValue `hcl:"query,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Results: min=0
	Results []datacloudassetresourcessearchall.Results `hcl:"results,block" validate:"min=0"`
}
type dataCloudAssetResourcesSearchAllAttributes struct {
	ref terra.Reference
}

// AssetTypes returns a reference to field asset_types of google_cloud_asset_resources_search_all.
func (carsa dataCloudAssetResourcesSearchAllAttributes) AssetTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](carsa.ref.Append("asset_types"))
}

// Id returns a reference to field id of google_cloud_asset_resources_search_all.
func (carsa dataCloudAssetResourcesSearchAllAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(carsa.ref.Append("id"))
}

// Query returns a reference to field query of google_cloud_asset_resources_search_all.
func (carsa dataCloudAssetResourcesSearchAllAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(carsa.ref.Append("query"))
}

// Scope returns a reference to field scope of google_cloud_asset_resources_search_all.
func (carsa dataCloudAssetResourcesSearchAllAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(carsa.ref.Append("scope"))
}

func (carsa dataCloudAssetResourcesSearchAllAttributes) Results() terra.ListValue[datacloudassetresourcessearchall.ResultsAttributes] {
	return terra.ReferenceAsList[datacloudassetresourcessearchall.ResultsAttributes](carsa.ref.Append("results"))
}
