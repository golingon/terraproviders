// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datawafv2regexpatternset "github.com/golingon/terraproviders/aws/5.0.1/datawafv2regexpatternset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWafv2RegexPatternSet creates a new instance of [DataWafv2RegexPatternSet].
func NewDataWafv2RegexPatternSet(name string, args DataWafv2RegexPatternSetArgs) *DataWafv2RegexPatternSet {
	return &DataWafv2RegexPatternSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafv2RegexPatternSet)(nil)

// DataWafv2RegexPatternSet represents the Terraform data resource aws_wafv2_regex_pattern_set.
type DataWafv2RegexPatternSet struct {
	Name string
	Args DataWafv2RegexPatternSetArgs
}

// DataSource returns the Terraform object type for [DataWafv2RegexPatternSet].
func (wrps *DataWafv2RegexPatternSet) DataSource() string {
	return "aws_wafv2_regex_pattern_set"
}

// LocalName returns the local name for [DataWafv2RegexPatternSet].
func (wrps *DataWafv2RegexPatternSet) LocalName() string {
	return wrps.Name
}

// Configuration returns the configuration (args) for [DataWafv2RegexPatternSet].
func (wrps *DataWafv2RegexPatternSet) Configuration() interface{} {
	return wrps.Args
}

// Attributes returns the attributes for [DataWafv2RegexPatternSet].
func (wrps *DataWafv2RegexPatternSet) Attributes() dataWafv2RegexPatternSetAttributes {
	return dataWafv2RegexPatternSetAttributes{ref: terra.ReferenceDataResource(wrps)}
}

// DataWafv2RegexPatternSetArgs contains the configurations for aws_wafv2_regex_pattern_set.
type DataWafv2RegexPatternSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// RegularExpression: min=0
	RegularExpression []datawafv2regexpatternset.RegularExpression `hcl:"regular_expression,block" validate:"min=0"`
}
type dataWafv2RegexPatternSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafv2_regex_pattern_set.
func (wrps dataWafv2RegexPatternSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("arn"))
}

// Description returns a reference to field description of aws_wafv2_regex_pattern_set.
func (wrps dataWafv2RegexPatternSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_regex_pattern_set.
func (wrps dataWafv2RegexPatternSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafv2_regex_pattern_set.
func (wrps dataWafv2RegexPatternSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_regex_pattern_set.
func (wrps dataWafv2RegexPatternSetAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(wrps.ref.Append("scope"))
}

func (wrps dataWafv2RegexPatternSetAttributes) RegularExpression() terra.SetValue[datawafv2regexpatternset.RegularExpressionAttributes] {
	return terra.ReferenceAsSet[datawafv2regexpatternset.RegularExpressionAttributes](wrps.ref.Append("regular_expression"))
}
