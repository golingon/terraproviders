// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codegurureviewerrepositoryassociation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type S3RepositoryDetails struct {
	// CodeArtifacts: min=0
	CodeArtifacts []CodeArtifacts `hcl:"code_artifacts,block" validate:"min=0"`
}

type CodeArtifacts struct{}

type KmsKeyDetails struct {
	// EncryptionOption: string, optional
	EncryptionOption terra.StringValue `hcl:"encryption_option,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
}

type Repository struct {
	// Bitbucket: optional
	Bitbucket *Bitbucket `hcl:"bitbucket,block"`
	// Codecommit: optional
	Codecommit *Codecommit `hcl:"codecommit,block"`
	// GithubEnterpriseServer: optional
	GithubEnterpriseServer *GithubEnterpriseServer `hcl:"github_enterprise_server,block"`
	// S3Bucket: optional
	S3Bucket *S3Bucket `hcl:"s3_bucket,block"`
}

type Bitbucket struct {
	// ConnectionArn: string, required
	ConnectionArn terra.StringValue `hcl:"connection_arn,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
}

type Codecommit struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type GithubEnterpriseServer struct {
	// ConnectionArn: string, required
	ConnectionArn terra.StringValue `hcl:"connection_arn,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
}

type S3Bucket struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type S3RepositoryDetailsAttributes struct {
	ref terra.Reference
}

func (srd S3RepositoryDetailsAttributes) InternalRef() (terra.Reference, error) {
	return srd.ref, nil
}

func (srd S3RepositoryDetailsAttributes) InternalWithRef(ref terra.Reference) S3RepositoryDetailsAttributes {
	return S3RepositoryDetailsAttributes{ref: ref}
}

func (srd S3RepositoryDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return srd.ref.InternalTokens()
}

func (srd S3RepositoryDetailsAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(srd.ref.Append("bucket_name"))
}

func (srd S3RepositoryDetailsAttributes) CodeArtifacts() terra.ListValue[CodeArtifactsAttributes] {
	return terra.ReferenceAsList[CodeArtifactsAttributes](srd.ref.Append("code_artifacts"))
}

type CodeArtifactsAttributes struct {
	ref terra.Reference
}

func (ca CodeArtifactsAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca CodeArtifactsAttributes) InternalWithRef(ref terra.Reference) CodeArtifactsAttributes {
	return CodeArtifactsAttributes{ref: ref}
}

func (ca CodeArtifactsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca CodeArtifactsAttributes) BuildArtifactsObjectKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("build_artifacts_object_key"))
}

func (ca CodeArtifactsAttributes) SourceCodeArtifactsObjectKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("source_code_artifacts_object_key"))
}

type KmsKeyDetailsAttributes struct {
	ref terra.Reference
}

func (kkd KmsKeyDetailsAttributes) InternalRef() (terra.Reference, error) {
	return kkd.ref, nil
}

func (kkd KmsKeyDetailsAttributes) InternalWithRef(ref terra.Reference) KmsKeyDetailsAttributes {
	return KmsKeyDetailsAttributes{ref: ref}
}

func (kkd KmsKeyDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kkd.ref.InternalTokens()
}

func (kkd KmsKeyDetailsAttributes) EncryptionOption() terra.StringValue {
	return terra.ReferenceAsString(kkd.ref.Append("encryption_option"))
}

func (kkd KmsKeyDetailsAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(kkd.ref.Append("kms_key_id"))
}

type RepositoryAttributes struct {
	ref terra.Reference
}

func (r RepositoryAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RepositoryAttributes) InternalWithRef(ref terra.Reference) RepositoryAttributes {
	return RepositoryAttributes{ref: ref}
}

func (r RepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RepositoryAttributes) Bitbucket() terra.ListValue[BitbucketAttributes] {
	return terra.ReferenceAsList[BitbucketAttributes](r.ref.Append("bitbucket"))
}

func (r RepositoryAttributes) Codecommit() terra.ListValue[CodecommitAttributes] {
	return terra.ReferenceAsList[CodecommitAttributes](r.ref.Append("codecommit"))
}

func (r RepositoryAttributes) GithubEnterpriseServer() terra.ListValue[GithubEnterpriseServerAttributes] {
	return terra.ReferenceAsList[GithubEnterpriseServerAttributes](r.ref.Append("github_enterprise_server"))
}

func (r RepositoryAttributes) S3Bucket() terra.ListValue[S3BucketAttributes] {
	return terra.ReferenceAsList[S3BucketAttributes](r.ref.Append("s3_bucket"))
}

type BitbucketAttributes struct {
	ref terra.Reference
}

func (b BitbucketAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b BitbucketAttributes) InternalWithRef(ref terra.Reference) BitbucketAttributes {
	return BitbucketAttributes{ref: ref}
}

func (b BitbucketAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b BitbucketAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("connection_arn"))
}

func (b BitbucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("name"))
}

func (b BitbucketAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(b.ref.Append("owner"))
}

type CodecommitAttributes struct {
	ref terra.Reference
}

func (c CodecommitAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CodecommitAttributes) InternalWithRef(ref terra.Reference) CodecommitAttributes {
	return CodecommitAttributes{ref: ref}
}

func (c CodecommitAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CodecommitAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

type GithubEnterpriseServerAttributes struct {
	ref terra.Reference
}

func (ges GithubEnterpriseServerAttributes) InternalRef() (terra.Reference, error) {
	return ges.ref, nil
}

func (ges GithubEnterpriseServerAttributes) InternalWithRef(ref terra.Reference) GithubEnterpriseServerAttributes {
	return GithubEnterpriseServerAttributes{ref: ref}
}

func (ges GithubEnterpriseServerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ges.ref.InternalTokens()
}

func (ges GithubEnterpriseServerAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(ges.ref.Append("connection_arn"))
}

func (ges GithubEnterpriseServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ges.ref.Append("name"))
}

func (ges GithubEnterpriseServerAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(ges.ref.Append("owner"))
}

type S3BucketAttributes struct {
	ref terra.Reference
}

func (sb S3BucketAttributes) InternalRef() (terra.Reference, error) {
	return sb.ref, nil
}

func (sb S3BucketAttributes) InternalWithRef(ref terra.Reference) S3BucketAttributes {
	return S3BucketAttributes{ref: ref}
}

func (sb S3BucketAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sb.ref.InternalTokens()
}

func (sb S3BucketAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_name"))
}

func (sb S3BucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("name"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type S3RepositoryDetailsState struct {
	BucketName    string               `json:"bucket_name"`
	CodeArtifacts []CodeArtifactsState `json:"code_artifacts"`
}

type CodeArtifactsState struct {
	BuildArtifactsObjectKey      string `json:"build_artifacts_object_key"`
	SourceCodeArtifactsObjectKey string `json:"source_code_artifacts_object_key"`
}

type KmsKeyDetailsState struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKeyId         string `json:"kms_key_id"`
}

type RepositoryState struct {
	Bitbucket              []BitbucketState              `json:"bitbucket"`
	Codecommit             []CodecommitState             `json:"codecommit"`
	GithubEnterpriseServer []GithubEnterpriseServerState `json:"github_enterprise_server"`
	S3Bucket               []S3BucketState               `json:"s3_bucket"`
}

type BitbucketState struct {
	ConnectionArn string `json:"connection_arn"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
}

type CodecommitState struct {
	Name string `json:"name"`
}

type GithubEnterpriseServerState struct {
	ConnectionArn string `json:"connection_arn"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
}

type S3BucketState struct {
	BucketName string `json:"bucket_name"`
	Name       string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
