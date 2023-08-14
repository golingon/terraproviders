// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package synapseworkspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AadAdmin struct {
	// Login: string, optional
	Login terra.StringValue `hcl:"login,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
}

type SqlAadAdmin struct {
	// Login: string, optional
	Login terra.StringValue `hcl:"login,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
}

type AzureDevopsRepo struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// BranchName: string, required
	BranchName terra.StringValue `hcl:"branch_name,attr" validate:"required"`
	// LastCommitId: string, optional
	LastCommitId terra.StringValue `hcl:"last_commit_id,attr"`
	// ProjectName: string, required
	ProjectName terra.StringValue `hcl:"project_name,attr" validate:"required"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// RootFolder: string, required
	RootFolder terra.StringValue `hcl:"root_folder,attr" validate:"required"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
}

type CustomerManagedKey struct {
	// KeyName: string, optional
	KeyName terra.StringValue `hcl:"key_name,attr"`
	// KeyVersionlessId: string, required
	KeyVersionlessId terra.StringValue `hcl:"key_versionless_id,attr" validate:"required"`
}

type GithubRepo struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// BranchName: string, required
	BranchName terra.StringValue `hcl:"branch_name,attr" validate:"required"`
	// GitUrl: string, optional
	GitUrl terra.StringValue `hcl:"git_url,attr"`
	// LastCommitId: string, optional
	LastCommitId terra.StringValue `hcl:"last_commit_id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// RootFolder: string, required
	RootFolder terra.StringValue `hcl:"root_folder,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AadAdminAttributes struct {
	ref terra.Reference
}

func (aa AadAdminAttributes) InternalRef() (terra.Reference, error) {
	return aa.ref, nil
}

func (aa AadAdminAttributes) InternalWithRef(ref terra.Reference) AadAdminAttributes {
	return AadAdminAttributes{ref: ref}
}

func (aa AadAdminAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aa.ref.InternalTokens()
}

func (aa AadAdminAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("login"))
}

func (aa AadAdminAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("object_id"))
}

func (aa AadAdminAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("tenant_id"))
}

type SqlAadAdminAttributes struct {
	ref terra.Reference
}

func (saa SqlAadAdminAttributes) InternalRef() (terra.Reference, error) {
	return saa.ref, nil
}

func (saa SqlAadAdminAttributes) InternalWithRef(ref terra.Reference) SqlAadAdminAttributes {
	return SqlAadAdminAttributes{ref: ref}
}

func (saa SqlAadAdminAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return saa.ref.InternalTokens()
}

func (saa SqlAadAdminAttributes) Login() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("login"))
}

func (saa SqlAadAdminAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("object_id"))
}

func (saa SqlAadAdminAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(saa.ref.Append("tenant_id"))
}

type AzureDevopsRepoAttributes struct {
	ref terra.Reference
}

func (adr AzureDevopsRepoAttributes) InternalRef() (terra.Reference, error) {
	return adr.ref, nil
}

func (adr AzureDevopsRepoAttributes) InternalWithRef(ref terra.Reference) AzureDevopsRepoAttributes {
	return AzureDevopsRepoAttributes{ref: ref}
}

func (adr AzureDevopsRepoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return adr.ref.InternalTokens()
}

func (adr AzureDevopsRepoAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("account_name"))
}

func (adr AzureDevopsRepoAttributes) BranchName() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("branch_name"))
}

func (adr AzureDevopsRepoAttributes) LastCommitId() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("last_commit_id"))
}

func (adr AzureDevopsRepoAttributes) ProjectName() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("project_name"))
}

func (adr AzureDevopsRepoAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("repository_name"))
}

func (adr AzureDevopsRepoAttributes) RootFolder() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("root_folder"))
}

func (adr AzureDevopsRepoAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(adr.ref.Append("tenant_id"))
}

type CustomerManagedKeyAttributes struct {
	ref terra.Reference
}

func (cmk CustomerManagedKeyAttributes) InternalRef() (terra.Reference, error) {
	return cmk.ref, nil
}

func (cmk CustomerManagedKeyAttributes) InternalWithRef(ref terra.Reference) CustomerManagedKeyAttributes {
	return CustomerManagedKeyAttributes{ref: ref}
}

func (cmk CustomerManagedKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cmk.ref.InternalTokens()
}

func (cmk CustomerManagedKeyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(cmk.ref.Append("key_name"))
}

func (cmk CustomerManagedKeyAttributes) KeyVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(cmk.ref.Append("key_versionless_id"))
}

type GithubRepoAttributes struct {
	ref terra.Reference
}

func (gr GithubRepoAttributes) InternalRef() (terra.Reference, error) {
	return gr.ref, nil
}

func (gr GithubRepoAttributes) InternalWithRef(ref terra.Reference) GithubRepoAttributes {
	return GithubRepoAttributes{ref: ref}
}

func (gr GithubRepoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gr.ref.InternalTokens()
}

func (gr GithubRepoAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("account_name"))
}

func (gr GithubRepoAttributes) BranchName() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("branch_name"))
}

func (gr GithubRepoAttributes) GitUrl() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("git_url"))
}

func (gr GithubRepoAttributes) LastCommitId() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("last_commit_id"))
}

func (gr GithubRepoAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("repository_name"))
}

func (gr GithubRepoAttributes) RootFolder() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("root_folder"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AadAdminState struct {
	Login    string `json:"login"`
	ObjectId string `json:"object_id"`
	TenantId string `json:"tenant_id"`
}

type SqlAadAdminState struct {
	Login    string `json:"login"`
	ObjectId string `json:"object_id"`
	TenantId string `json:"tenant_id"`
}

type AzureDevopsRepoState struct {
	AccountName    string `json:"account_name"`
	BranchName     string `json:"branch_name"`
	LastCommitId   string `json:"last_commit_id"`
	ProjectName    string `json:"project_name"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
	TenantId       string `json:"tenant_id"`
}

type CustomerManagedKeyState struct {
	KeyName          string `json:"key_name"`
	KeyVersionlessId string `json:"key_versionless_id"`
}

type GithubRepoState struct {
	AccountName    string `json:"account_name"`
	BranchName     string `json:"branch_name"`
	GitUrl         string `json:"git_url"`
	LastCommitId   string `json:"last_commit_id"`
	RepositoryName string `json:"repository_name"`
	RootFolder     string `json:"root_folder"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}