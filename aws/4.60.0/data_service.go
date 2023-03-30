// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataService(name string, args DataServiceArgs) *DataService {
	return &DataService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataService)(nil)

type DataService struct {
	Name string
	Args DataServiceArgs
}

func (s *DataService) DataSource() string {
	return "aws_service"
}

func (s *DataService) LocalName() string {
	return s.Name
}

func (s *DataService) Configuration() interface{} {
	return s.Args
}

func (s *DataService) Attributes() dataServiceAttributes {
	return dataServiceAttributes{ref: terra.ReferenceDataResource(s)}
}

type DataServiceArgs struct {
	// DnsName: string, optional
	DnsName terra.StringValue `hcl:"dns_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReverseDnsName: string, optional
	ReverseDnsName terra.StringValue `hcl:"reverse_dns_name,attr"`
	// ReverseDnsPrefix: string, optional
	ReverseDnsPrefix terra.StringValue `hcl:"reverse_dns_prefix,attr"`
	// ServiceId: string, optional
	ServiceId terra.StringValue `hcl:"service_id,attr"`
}
type dataServiceAttributes struct {
	ref terra.Reference
}

func (s dataServiceAttributes) DnsName() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("dns_name"))
}

func (s dataServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("id"))
}

func (s dataServiceAttributes) Partition() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("partition"))
}

func (s dataServiceAttributes) Region() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("region"))
}

func (s dataServiceAttributes) ReverseDnsName() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("reverse_dns_name"))
}

func (s dataServiceAttributes) ReverseDnsPrefix() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("reverse_dns_prefix"))
}

func (s dataServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("service_id"))
}

func (s dataServiceAttributes) Supported() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("supported"))
}
