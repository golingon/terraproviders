// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDocdbOrderableDbInstance(name string, args DataDocdbOrderableDbInstanceArgs) *DataDocdbOrderableDbInstance {
	return &DataDocdbOrderableDbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDocdbOrderableDbInstance)(nil)

type DataDocdbOrderableDbInstance struct {
	Name string
	Args DataDocdbOrderableDbInstanceArgs
}

func (dodi *DataDocdbOrderableDbInstance) DataSource() string {
	return "aws_docdb_orderable_db_instance"
}

func (dodi *DataDocdbOrderableDbInstance) LocalName() string {
	return dodi.Name
}

func (dodi *DataDocdbOrderableDbInstance) Configuration() interface{} {
	return dodi.Args
}

func (dodi *DataDocdbOrderableDbInstance) Attributes() dataDocdbOrderableDbInstanceAttributes {
	return dataDocdbOrderableDbInstanceAttributes{ref: terra.ReferenceDataResource(dodi)}
}

type DataDocdbOrderableDbInstanceArgs struct {
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceClass: string, optional
	InstanceClass terra.StringValue `hcl:"instance_class,attr"`
	// LicenseModel: string, optional
	LicenseModel terra.StringValue `hcl:"license_model,attr"`
	// PreferredInstanceClasses: list of string, optional
	PreferredInstanceClasses terra.ListValue[terra.StringValue] `hcl:"preferred_instance_classes,attr"`
	// Vpc: bool, optional
	Vpc terra.BoolValue `hcl:"vpc,attr"`
}
type dataDocdbOrderableDbInstanceAttributes struct {
	ref terra.Reference
}

func (dodi dataDocdbOrderableDbInstanceAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dodi.ref.Append("availability_zones"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(dodi.ref.Append("engine"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceString(dodi.ref.Append("engine_version"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dodi.ref.Append("id"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceString(dodi.ref.Append("instance_class"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceString(dodi.ref.Append("license_model"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) PreferredInstanceClasses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dodi.ref.Append("preferred_instance_classes"))
}

func (dodi dataDocdbOrderableDbInstanceAttributes) Vpc() terra.BoolValue {
	return terra.ReferenceBool(dodi.ref.Append("vpc"))
}
