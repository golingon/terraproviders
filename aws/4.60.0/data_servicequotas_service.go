// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataServicequotasService(name string, args DataServicequotasServiceArgs) *DataServicequotasService {
	return &DataServicequotasService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicequotasService)(nil)

type DataServicequotasService struct {
	Name string
	Args DataServicequotasServiceArgs
}

func (ss *DataServicequotasService) DataSource() string {
	return "aws_servicequotas_service"
}

func (ss *DataServicequotasService) LocalName() string {
	return ss.Name
}

func (ss *DataServicequotasService) Configuration() interface{} {
	return ss.Args
}

func (ss *DataServicequotasService) Attributes() dataServicequotasServiceAttributes {
	return dataServicequotasServiceAttributes{ref: terra.ReferenceDataResource(ss)}
}

type DataServicequotasServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}
type dataServicequotasServiceAttributes struct {
	ref terra.Reference
}

func (ss dataServicequotasServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("id"))
}

func (ss dataServicequotasServiceAttributes) ServiceCode() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("service_code"))
}

func (ss dataServicequotasServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("service_name"))
}
