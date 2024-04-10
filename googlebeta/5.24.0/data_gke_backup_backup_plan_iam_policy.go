// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataGkeBackupBackupPlanIamPolicy creates a new instance of [DataGkeBackupBackupPlanIamPolicy].
func NewDataGkeBackupBackupPlanIamPolicy(name string, args DataGkeBackupBackupPlanIamPolicyArgs) *DataGkeBackupBackupPlanIamPolicy {
	return &DataGkeBackupBackupPlanIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGkeBackupBackupPlanIamPolicy)(nil)

// DataGkeBackupBackupPlanIamPolicy represents the Terraform data resource google_gke_backup_backup_plan_iam_policy.
type DataGkeBackupBackupPlanIamPolicy struct {
	Name string
	Args DataGkeBackupBackupPlanIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataGkeBackupBackupPlanIamPolicy].
func (gbbpip *DataGkeBackupBackupPlanIamPolicy) DataSource() string {
	return "google_gke_backup_backup_plan_iam_policy"
}

// LocalName returns the local name for [DataGkeBackupBackupPlanIamPolicy].
func (gbbpip *DataGkeBackupBackupPlanIamPolicy) LocalName() string {
	return gbbpip.Name
}

// Configuration returns the configuration (args) for [DataGkeBackupBackupPlanIamPolicy].
func (gbbpip *DataGkeBackupBackupPlanIamPolicy) Configuration() interface{} {
	return gbbpip.Args
}

// Attributes returns the attributes for [DataGkeBackupBackupPlanIamPolicy].
func (gbbpip *DataGkeBackupBackupPlanIamPolicy) Attributes() dataGkeBackupBackupPlanIamPolicyAttributes {
	return dataGkeBackupBackupPlanIamPolicyAttributes{ref: terra.ReferenceDataResource(gbbpip)}
}

// DataGkeBackupBackupPlanIamPolicyArgs contains the configurations for google_gke_backup_backup_plan_iam_policy.
type DataGkeBackupBackupPlanIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataGkeBackupBackupPlanIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("location"))
}

// Name returns a reference to field name of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_gke_backup_backup_plan_iam_policy.
func (gbbpip dataGkeBackupBackupPlanIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("project"))
}
