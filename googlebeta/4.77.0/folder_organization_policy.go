// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	folderorganizationpolicy "github.com/golingon/terraproviders/googlebeta/4.77.0/folderorganizationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolderOrganizationPolicy creates a new instance of [FolderOrganizationPolicy].
func NewFolderOrganizationPolicy(name string, args FolderOrganizationPolicyArgs) *FolderOrganizationPolicy {
	return &FolderOrganizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderOrganizationPolicy)(nil)

// FolderOrganizationPolicy represents the Terraform resource google_folder_organization_policy.
type FolderOrganizationPolicy struct {
	Name      string
	Args      FolderOrganizationPolicyArgs
	state     *folderOrganizationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) Type() string {
	return "google_folder_organization_policy"
}

// LocalName returns the local name for [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) LocalName() string {
	return fop.Name
}

// Configuration returns the configuration (args) for [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) Configuration() interface{} {
	return fop.Args
}

// DependOn is used for other resources to depend on [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(fop)
}

// Dependencies returns the list of resources [FolderOrganizationPolicy] depends_on.
func (fop *FolderOrganizationPolicy) Dependencies() terra.Dependencies {
	return fop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) LifecycleManagement() *terra.Lifecycle {
	return fop.Lifecycle
}

// Attributes returns the attributes for [FolderOrganizationPolicy].
func (fop *FolderOrganizationPolicy) Attributes() folderOrganizationPolicyAttributes {
	return folderOrganizationPolicyAttributes{ref: terra.ReferenceResource(fop)}
}

// ImportState imports the given attribute values into [FolderOrganizationPolicy]'s state.
func (fop *FolderOrganizationPolicy) ImportState(av io.Reader) error {
	fop.state = &folderOrganizationPolicyState{}
	if err := json.NewDecoder(av).Decode(fop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fop.Type(), fop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FolderOrganizationPolicy] has state.
func (fop *FolderOrganizationPolicy) State() (*folderOrganizationPolicyState, bool) {
	return fop.state, fop.state != nil
}

// StateMust returns the state for [FolderOrganizationPolicy]. Panics if the state is nil.
func (fop *FolderOrganizationPolicy) StateMust() *folderOrganizationPolicyState {
	if fop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fop.Type(), fop.LocalName()))
	}
	return fop.state
}

// FolderOrganizationPolicyArgs contains the configurations for google_folder_organization_policy.
type FolderOrganizationPolicyArgs struct {
	// Constraint: string, required
	Constraint terra.StringValue `hcl:"constraint,attr" validate:"required"`
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// BooleanPolicy: optional
	BooleanPolicy *folderorganizationpolicy.BooleanPolicy `hcl:"boolean_policy,block"`
	// ListPolicy: optional
	ListPolicy *folderorganizationpolicy.ListPolicy `hcl:"list_policy,block"`
	// RestorePolicy: optional
	RestorePolicy *folderorganizationpolicy.RestorePolicy `hcl:"restore_policy,block"`
	// Timeouts: optional
	Timeouts *folderorganizationpolicy.Timeouts `hcl:"timeouts,block"`
}
type folderOrganizationPolicyAttributes struct {
	ref terra.Reference
}

// Constraint returns a reference to field constraint of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) Constraint() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("constraint"))
}

// Etag returns a reference to field etag of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("id"))
}

// UpdateTime returns a reference to field update_time of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fop.ref.Append("update_time"))
}

// Version returns a reference to field version of google_folder_organization_policy.
func (fop folderOrganizationPolicyAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(fop.ref.Append("version"))
}

func (fop folderOrganizationPolicyAttributes) BooleanPolicy() terra.ListValue[folderorganizationpolicy.BooleanPolicyAttributes] {
	return terra.ReferenceAsList[folderorganizationpolicy.BooleanPolicyAttributes](fop.ref.Append("boolean_policy"))
}

func (fop folderOrganizationPolicyAttributes) ListPolicy() terra.ListValue[folderorganizationpolicy.ListPolicyAttributes] {
	return terra.ReferenceAsList[folderorganizationpolicy.ListPolicyAttributes](fop.ref.Append("list_policy"))
}

func (fop folderOrganizationPolicyAttributes) RestorePolicy() terra.ListValue[folderorganizationpolicy.RestorePolicyAttributes] {
	return terra.ReferenceAsList[folderorganizationpolicy.RestorePolicyAttributes](fop.ref.Append("restore_policy"))
}

func (fop folderOrganizationPolicyAttributes) Timeouts() folderorganizationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[folderorganizationpolicy.TimeoutsAttributes](fop.ref.Append("timeouts"))
}

type folderOrganizationPolicyState struct {
	Constraint    string                                        `json:"constraint"`
	Etag          string                                        `json:"etag"`
	Folder        string                                        `json:"folder"`
	Id            string                                        `json:"id"`
	UpdateTime    string                                        `json:"update_time"`
	Version       float64                                       `json:"version"`
	BooleanPolicy []folderorganizationpolicy.BooleanPolicyState `json:"boolean_policy"`
	ListPolicy    []folderorganizationpolicy.ListPolicyState    `json:"list_policy"`
	RestorePolicy []folderorganizationpolicy.RestorePolicyState `json:"restore_policy"`
	Timeouts      *folderorganizationpolicy.TimeoutsState       `json:"timeouts"`
}
