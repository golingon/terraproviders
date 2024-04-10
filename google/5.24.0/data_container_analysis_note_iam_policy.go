// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataContainerAnalysisNoteIamPolicy creates a new instance of [DataContainerAnalysisNoteIamPolicy].
func NewDataContainerAnalysisNoteIamPolicy(name string, args DataContainerAnalysisNoteIamPolicyArgs) *DataContainerAnalysisNoteIamPolicy {
	return &DataContainerAnalysisNoteIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerAnalysisNoteIamPolicy)(nil)

// DataContainerAnalysisNoteIamPolicy represents the Terraform data resource google_container_analysis_note_iam_policy.
type DataContainerAnalysisNoteIamPolicy struct {
	Name string
	Args DataContainerAnalysisNoteIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataContainerAnalysisNoteIamPolicy].
func (canip *DataContainerAnalysisNoteIamPolicy) DataSource() string {
	return "google_container_analysis_note_iam_policy"
}

// LocalName returns the local name for [DataContainerAnalysisNoteIamPolicy].
func (canip *DataContainerAnalysisNoteIamPolicy) LocalName() string {
	return canip.Name
}

// Configuration returns the configuration (args) for [DataContainerAnalysisNoteIamPolicy].
func (canip *DataContainerAnalysisNoteIamPolicy) Configuration() interface{} {
	return canip.Args
}

// Attributes returns the attributes for [DataContainerAnalysisNoteIamPolicy].
func (canip *DataContainerAnalysisNoteIamPolicy) Attributes() dataContainerAnalysisNoteIamPolicyAttributes {
	return dataContainerAnalysisNoteIamPolicyAttributes{ref: terra.ReferenceDataResource(canip)}
}

// DataContainerAnalysisNoteIamPolicyArgs contains the configurations for google_container_analysis_note_iam_policy.
type DataContainerAnalysisNoteIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Note: string, required
	Note terra.StringValue `hcl:"note,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataContainerAnalysisNoteIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_container_analysis_note_iam_policy.
func (canip dataContainerAnalysisNoteIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_analysis_note_iam_policy.
func (canip dataContainerAnalysisNoteIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("id"))
}

// Note returns a reference to field note of google_container_analysis_note_iam_policy.
func (canip dataContainerAnalysisNoteIamPolicyAttributes) Note() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("note"))
}

// PolicyData returns a reference to field policy_data of google_container_analysis_note_iam_policy.
func (canip dataContainerAnalysisNoteIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_container_analysis_note_iam_policy.
func (canip dataContainerAnalysisNoteIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("project"))
}
