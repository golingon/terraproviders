// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataGuarddutyDetector creates a new instance of [DataGuarddutyDetector].
func NewDataGuarddutyDetector(name string, args DataGuarddutyDetectorArgs) *DataGuarddutyDetector {
	return &DataGuarddutyDetector{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGuarddutyDetector)(nil)

// DataGuarddutyDetector represents the Terraform data resource aws_guardduty_detector.
type DataGuarddutyDetector struct {
	Name string
	Args DataGuarddutyDetectorArgs
}

// DataSource returns the Terraform object type for [DataGuarddutyDetector].
func (gd *DataGuarddutyDetector) DataSource() string {
	return "aws_guardduty_detector"
}

// LocalName returns the local name for [DataGuarddutyDetector].
func (gd *DataGuarddutyDetector) LocalName() string {
	return gd.Name
}

// Configuration returns the configuration (args) for [DataGuarddutyDetector].
func (gd *DataGuarddutyDetector) Configuration() interface{} {
	return gd.Args
}

// Attributes returns the attributes for [DataGuarddutyDetector].
func (gd *DataGuarddutyDetector) Attributes() dataGuarddutyDetectorAttributes {
	return dataGuarddutyDetectorAttributes{ref: terra.ReferenceDataResource(gd)}
}

// DataGuarddutyDetectorArgs contains the configurations for aws_guardduty_detector.
type DataGuarddutyDetectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataGuarddutyDetectorAttributes struct {
	ref terra.Reference
}

// FindingPublishingFrequency returns a reference to field finding_publishing_frequency of aws_guardduty_detector.
func (gd dataGuarddutyDetectorAttributes) FindingPublishingFrequency() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("finding_publishing_frequency"))
}

// Id returns a reference to field id of aws_guardduty_detector.
func (gd dataGuarddutyDetectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("id"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_guardduty_detector.
func (gd dataGuarddutyDetectorAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("service_role_arn"))
}

// Status returns a reference to field status of aws_guardduty_detector.
func (gd dataGuarddutyDetectorAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("status"))
}