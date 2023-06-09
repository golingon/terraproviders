// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDynamodbTableItem creates a new instance of [DataDynamodbTableItem].
func NewDataDynamodbTableItem(name string, args DataDynamodbTableItemArgs) *DataDynamodbTableItem {
	return &DataDynamodbTableItem{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDynamodbTableItem)(nil)

// DataDynamodbTableItem represents the Terraform data resource aws_dynamodb_table_item.
type DataDynamodbTableItem struct {
	Name string
	Args DataDynamodbTableItemArgs
}

// DataSource returns the Terraform object type for [DataDynamodbTableItem].
func (dti *DataDynamodbTableItem) DataSource() string {
	return "aws_dynamodb_table_item"
}

// LocalName returns the local name for [DataDynamodbTableItem].
func (dti *DataDynamodbTableItem) LocalName() string {
	return dti.Name
}

// Configuration returns the configuration (args) for [DataDynamodbTableItem].
func (dti *DataDynamodbTableItem) Configuration() interface{} {
	return dti.Args
}

// Attributes returns the attributes for [DataDynamodbTableItem].
func (dti *DataDynamodbTableItem) Attributes() dataDynamodbTableItemAttributes {
	return dataDynamodbTableItemAttributes{ref: terra.ReferenceDataResource(dti)}
}

// DataDynamodbTableItemArgs contains the configurations for aws_dynamodb_table_item.
type DataDynamodbTableItemArgs struct {
	// ExpressionAttributeNames: map of string, optional
	ExpressionAttributeNames terra.MapValue[terra.StringValue] `hcl:"expression_attribute_names,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ProjectionExpression: string, optional
	ProjectionExpression terra.StringValue `hcl:"projection_expression,attr"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
}
type dataDynamodbTableItemAttributes struct {
	ref terra.Reference
}

// ExpressionAttributeNames returns a reference to field expression_attribute_names of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) ExpressionAttributeNames() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dti.ref.Append("expression_attribute_names"))
}

// Id returns a reference to field id of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("id"))
}

// Item returns a reference to field item of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) Item() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("item"))
}

// Key returns a reference to field key of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("key"))
}

// ProjectionExpression returns a reference to field projection_expression of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) ProjectionExpression() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("projection_expression"))
}

// TableName returns a reference to field table_name of aws_dynamodb_table_item.
func (dti dataDynamodbTableItemAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dti.ref.Append("table_name"))
}
