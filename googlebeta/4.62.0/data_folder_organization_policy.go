// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datafolderorganizationpolicy "github.com/golingon/terraproviders/googlebeta/4.62.0/datafolderorganizationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataFolderOrganizationPolicy creates a new instance of [DataFolderOrganizationPolicy].
func NewDataFolderOrganizationPolicy(name string, args DataFolderOrganizationPolicyArgs) *DataFolderOrganizationPolicy {
	return &DataFolderOrganizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFolderOrganizationPolicy)(nil)

// DataFolderOrganizationPolicy represents the Terraform data resource google_folder_organization_policy.
type DataFolderOrganizationPolicy struct {
	Name string
	Args DataFolderOrganizationPolicyArgs
}

// DataSource returns the Terraform object type for [DataFolderOrganizationPolicy].
func (fop *DataFolderOrganizationPolicy) DataSource() string {
	return "google_folder_organization_policy"
}

// LocalName returns the local name for [DataFolderOrganizationPolicy].
func (fop *DataFolderOrganizationPolicy) LocalName() string {
	return fop.Name
}

// Configuration returns the configuration (args) for [DataFolderOrganizationPolicy].
func (fop *DataFolderOrganizationPolicy) Configuration() interface{} {
	return fop.Args
}

// Attributes returns the attributes for [DataFolderOrganizationPolicy].
func (fop *DataFolderOrganizationPolicy) Attributes() dataFolderOrganizationPolicyAttributes {
	return dataFolderOrganizationPolicyAttributes{ref: terra.ReferenceDataResource(fop)}
}

// DataFolderOrganizationPolicyArgs contains the configurations for google_folder_organization_policy.
type DataFolderOrganizationPolicyArgs struct {
	// Constraint: string, required
	Constraint terra.StringValue `hcl:"constraint,attr" validate:"required"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// BooleanPolicy: min=0
	BooleanPolicy []datafolderorganizationpolicy.BooleanPolicy `hcl:"boolean_policy,block" validate:"min=0"`
	// ListPolicy: min=0
	ListPolicy []datafolderorganizationpolicy.ListPolicy `hcl:"list_policy,block" validate:"min=0"`
	// RestorePolicy: min=0
	RestorePolicy []datafolderorganizationpolicy.RestorePolicy `hcl:"restore_policy,block" validate:"min=0"`
}
type dataFolderOrganizationPolicyAttributes struct {
	ref terra.Reference
}

// Constraint returns a reference to field constraint of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) Constraint() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("constraint"))
}

// Etag returns a reference to field etag of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("id"))
}

// UpdateTime returns a reference to field update_time of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("update_time"))
}

// Version returns a reference to field version of google_folder_organization_policy.
func (fop dataFolderOrganizationPolicyAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(fop.ref.Append("version"))
}

func (fop dataFolderOrganizationPolicyAttributes) BooleanPolicy() terra.ListValue[datafolderorganizationpolicy.BooleanPolicyAttributes] {
	return terra.ReferenceAsList[datafolderorganizationpolicy.BooleanPolicyAttributes](fop.ref.Append("boolean_policy"))
}

func (fop dataFolderOrganizationPolicyAttributes) ListPolicy() terra.ListValue[datafolderorganizationpolicy.ListPolicyAttributes] {
	return terra.ReferenceAsList[datafolderorganizationpolicy.ListPolicyAttributes](fop.ref.Append("list_policy"))
}

func (fop dataFolderOrganizationPolicyAttributes) RestorePolicy() terra.ListValue[datafolderorganizationpolicy.RestorePolicyAttributes] {
	return terra.ReferenceAsList[datafolderorganizationpolicy.RestorePolicyAttributes](fop.ref.Append("restore_policy"))
}
