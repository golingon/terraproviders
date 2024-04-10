// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"github.com/golingon/lingon/pkg/terra"
	datacloudquotasquotainfos "github.com/golingon/terraproviders/googlebeta/5.24.0/datacloudquotasquotainfos"
)

// NewDataCloudQuotasQuotaInfos creates a new instance of [DataCloudQuotasQuotaInfos].
func NewDataCloudQuotasQuotaInfos(name string, args DataCloudQuotasQuotaInfosArgs) *DataCloudQuotasQuotaInfos {
	return &DataCloudQuotasQuotaInfos{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudQuotasQuotaInfos)(nil)

// DataCloudQuotasQuotaInfos represents the Terraform data resource google_cloud_quotas_quota_infos.
type DataCloudQuotasQuotaInfos struct {
	Name string
	Args DataCloudQuotasQuotaInfosArgs
}

// DataSource returns the Terraform object type for [DataCloudQuotasQuotaInfos].
func (cqqi *DataCloudQuotasQuotaInfos) DataSource() string {
	return "google_cloud_quotas_quota_infos"
}

// LocalName returns the local name for [DataCloudQuotasQuotaInfos].
func (cqqi *DataCloudQuotasQuotaInfos) LocalName() string {
	return cqqi.Name
}

// Configuration returns the configuration (args) for [DataCloudQuotasQuotaInfos].
func (cqqi *DataCloudQuotasQuotaInfos) Configuration() interface{} {
	return cqqi.Args
}

// Attributes returns the attributes for [DataCloudQuotasQuotaInfos].
func (cqqi *DataCloudQuotasQuotaInfos) Attributes() dataCloudQuotasQuotaInfosAttributes {
	return dataCloudQuotasQuotaInfosAttributes{ref: terra.ReferenceDataResource(cqqi)}
}

// DataCloudQuotasQuotaInfosArgs contains the configurations for google_cloud_quotas_quota_infos.
type DataCloudQuotasQuotaInfosArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// QuotaInfos: min=0
	QuotaInfos []datacloudquotasquotainfos.QuotaInfos `hcl:"quota_infos,block" validate:"min=0"`
}
type dataCloudQuotasQuotaInfosAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_cloud_quotas_quota_infos.
func (cqqi dataCloudQuotasQuotaInfosAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cqqi.ref.Append("id"))
}

// Parent returns a reference to field parent of google_cloud_quotas_quota_infos.
func (cqqi dataCloudQuotasQuotaInfosAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(cqqi.ref.Append("parent"))
}

// Service returns a reference to field service of google_cloud_quotas_quota_infos.
func (cqqi dataCloudQuotasQuotaInfosAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(cqqi.ref.Append("service"))
}

func (cqqi dataCloudQuotasQuotaInfosAttributes) QuotaInfos() terra.ListValue[datacloudquotasquotainfos.QuotaInfosAttributes] {
	return terra.ReferenceAsList[datacloudquotasquotainfos.QuotaInfosAttributes](cqqi.ref.Append("quota_infos"))
}
