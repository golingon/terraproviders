// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataWafregionalWebAcl creates a new instance of [DataWafregionalWebAcl].
func NewDataWafregionalWebAcl(name string, args DataWafregionalWebAclArgs) *DataWafregionalWebAcl {
	return &DataWafregionalWebAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafregionalWebAcl)(nil)

// DataWafregionalWebAcl represents the Terraform data resource aws_wafregional_web_acl.
type DataWafregionalWebAcl struct {
	Name string
	Args DataWafregionalWebAclArgs
}

// DataSource returns the Terraform object type for [DataWafregionalWebAcl].
func (wwa *DataWafregionalWebAcl) DataSource() string {
	return "aws_wafregional_web_acl"
}

// LocalName returns the local name for [DataWafregionalWebAcl].
func (wwa *DataWafregionalWebAcl) LocalName() string {
	return wwa.Name
}

// Configuration returns the configuration (args) for [DataWafregionalWebAcl].
func (wwa *DataWafregionalWebAcl) Configuration() interface{} {
	return wwa.Args
}

// Attributes returns the attributes for [DataWafregionalWebAcl].
func (wwa *DataWafregionalWebAcl) Attributes() dataWafregionalWebAclAttributes {
	return dataWafregionalWebAclAttributes{ref: terra.ReferenceDataResource(wwa)}
}

// DataWafregionalWebAclArgs contains the configurations for aws_wafregional_web_acl.
type DataWafregionalWebAclArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataWafregionalWebAclAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_web_acl.
func (wwa dataWafregionalWebAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_web_acl.
func (wwa dataWafregionalWebAclAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwa.ref.Append("name"))
}