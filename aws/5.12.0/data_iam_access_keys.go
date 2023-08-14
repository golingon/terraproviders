// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataiamaccesskeys "github.com/golingon/terraproviders/aws/5.12.0/dataiamaccesskeys"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamAccessKeys creates a new instance of [DataIamAccessKeys].
func NewDataIamAccessKeys(name string, args DataIamAccessKeysArgs) *DataIamAccessKeys {
	return &DataIamAccessKeys{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamAccessKeys)(nil)

// DataIamAccessKeys represents the Terraform data resource aws_iam_access_keys.
type DataIamAccessKeys struct {
	Name string
	Args DataIamAccessKeysArgs
}

// DataSource returns the Terraform object type for [DataIamAccessKeys].
func (iak *DataIamAccessKeys) DataSource() string {
	return "aws_iam_access_keys"
}

// LocalName returns the local name for [DataIamAccessKeys].
func (iak *DataIamAccessKeys) LocalName() string {
	return iak.Name
}

// Configuration returns the configuration (args) for [DataIamAccessKeys].
func (iak *DataIamAccessKeys) Configuration() interface{} {
	return iak.Args
}

// Attributes returns the attributes for [DataIamAccessKeys].
func (iak *DataIamAccessKeys) Attributes() dataIamAccessKeysAttributes {
	return dataIamAccessKeysAttributes{ref: terra.ReferenceDataResource(iak)}
}

// DataIamAccessKeysArgs contains the configurations for aws_iam_access_keys.
type DataIamAccessKeysArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
	// AccessKeys: min=0
	AccessKeys []dataiamaccesskeys.AccessKeys `hcl:"access_keys,block" validate:"min=0"`
}
type dataIamAccessKeysAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iam_access_keys.
func (iak dataIamAccessKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("id"))
}

// User returns a reference to field user of aws_iam_access_keys.
func (iak dataIamAccessKeysAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("user"))
}

func (iak dataIamAccessKeysAttributes) AccessKeys() terra.SetValue[dataiamaccesskeys.AccessKeysAttributes] {
	return terra.ReferenceAsSet[dataiamaccesskeys.AccessKeysAttributes](iak.ref.Append("access_keys"))
}
