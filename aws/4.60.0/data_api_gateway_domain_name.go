// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataapigatewaydomainname "github.com/golingon/terraproviders/aws/4.60.0/dataapigatewaydomainname"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataApiGatewayDomainName(name string, args DataApiGatewayDomainNameArgs) *DataApiGatewayDomainName {
	return &DataApiGatewayDomainName{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiGatewayDomainName)(nil)

type DataApiGatewayDomainName struct {
	Name string
	Args DataApiGatewayDomainNameArgs
}

func (agdn *DataApiGatewayDomainName) DataSource() string {
	return "aws_api_gateway_domain_name"
}

func (agdn *DataApiGatewayDomainName) LocalName() string {
	return agdn.Name
}

func (agdn *DataApiGatewayDomainName) Configuration() interface{} {
	return agdn.Args
}

func (agdn *DataApiGatewayDomainName) Attributes() dataApiGatewayDomainNameAttributes {
	return dataApiGatewayDomainNameAttributes{ref: terra.ReferenceDataResource(agdn)}
}

type DataApiGatewayDomainNameArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// EndpointConfiguration: min=0
	EndpointConfiguration []dataapigatewaydomainname.EndpointConfiguration `hcl:"endpoint_configuration,block" validate:"min=0"`
}
type dataApiGatewayDomainNameAttributes struct {
	ref terra.Reference
}

func (agdn dataApiGatewayDomainNameAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("arn"))
}

func (agdn dataApiGatewayDomainNameAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("certificate_arn"))
}

func (agdn dataApiGatewayDomainNameAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("certificate_name"))
}

func (agdn dataApiGatewayDomainNameAttributes) CertificateUploadDate() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("certificate_upload_date"))
}

func (agdn dataApiGatewayDomainNameAttributes) CloudfrontDomainName() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("cloudfront_domain_name"))
}

func (agdn dataApiGatewayDomainNameAttributes) CloudfrontZoneId() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("cloudfront_zone_id"))
}

func (agdn dataApiGatewayDomainNameAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("domain_name"))
}

func (agdn dataApiGatewayDomainNameAttributes) Id() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("id"))
}

func (agdn dataApiGatewayDomainNameAttributes) RegionalCertificateArn() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("regional_certificate_arn"))
}

func (agdn dataApiGatewayDomainNameAttributes) RegionalCertificateName() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("regional_certificate_name"))
}

func (agdn dataApiGatewayDomainNameAttributes) RegionalDomainName() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("regional_domain_name"))
}

func (agdn dataApiGatewayDomainNameAttributes) RegionalZoneId() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("regional_zone_id"))
}

func (agdn dataApiGatewayDomainNameAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceString(agdn.ref.Append("security_policy"))
}

func (agdn dataApiGatewayDomainNameAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](agdn.ref.Append("tags"))
}

func (agdn dataApiGatewayDomainNameAttributes) EndpointConfiguration() terra.ListValue[dataapigatewaydomainname.EndpointConfigurationAttributes] {
	return terra.ReferenceList[dataapigatewaydomainname.EndpointConfigurationAttributes](agdn.ref.Append("endpoint_configuration"))
}
