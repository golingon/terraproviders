// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIamAccessKey creates a new instance of [IamAccessKey].
func NewIamAccessKey(name string, args IamAccessKeyArgs) *IamAccessKey {
	return &IamAccessKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamAccessKey)(nil)

// IamAccessKey represents the Terraform resource aws_iam_access_key.
type IamAccessKey struct {
	Name      string
	Args      IamAccessKeyArgs
	state     *iamAccessKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamAccessKey].
func (iak *IamAccessKey) Type() string {
	return "aws_iam_access_key"
}

// LocalName returns the local name for [IamAccessKey].
func (iak *IamAccessKey) LocalName() string {
	return iak.Name
}

// Configuration returns the configuration (args) for [IamAccessKey].
func (iak *IamAccessKey) Configuration() interface{} {
	return iak.Args
}

// DependOn is used for other resources to depend on [IamAccessKey].
func (iak *IamAccessKey) DependOn() terra.Reference {
	return terra.ReferenceResource(iak)
}

// Dependencies returns the list of resources [IamAccessKey] depends_on.
func (iak *IamAccessKey) Dependencies() terra.Dependencies {
	return iak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamAccessKey].
func (iak *IamAccessKey) LifecycleManagement() *terra.Lifecycle {
	return iak.Lifecycle
}

// Attributes returns the attributes for [IamAccessKey].
func (iak *IamAccessKey) Attributes() iamAccessKeyAttributes {
	return iamAccessKeyAttributes{ref: terra.ReferenceResource(iak)}
}

// ImportState imports the given attribute values into [IamAccessKey]'s state.
func (iak *IamAccessKey) ImportState(av io.Reader) error {
	iak.state = &iamAccessKeyState{}
	if err := json.NewDecoder(av).Decode(iak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iak.Type(), iak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamAccessKey] has state.
func (iak *IamAccessKey) State() (*iamAccessKeyState, bool) {
	return iak.state, iak.state != nil
}

// StateMust returns the state for [IamAccessKey]. Panics if the state is nil.
func (iak *IamAccessKey) StateMust() *iamAccessKeyState {
	if iak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iak.Type(), iak.LocalName()))
	}
	return iak.state
}

// IamAccessKeyArgs contains the configurations for aws_iam_access_key.
type IamAccessKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PgpKey: string, optional
	PgpKey terra.StringValue `hcl:"pgp_key,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
}
type iamAccessKeyAttributes struct {
	ref terra.Reference
}

// CreateDate returns a reference to field create_date of aws_iam_access_key.
func (iak iamAccessKeyAttributes) CreateDate() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("create_date"))
}

// EncryptedSecret returns a reference to field encrypted_secret of aws_iam_access_key.
func (iak iamAccessKeyAttributes) EncryptedSecret() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("encrypted_secret"))
}

// EncryptedSesSmtpPasswordV4 returns a reference to field encrypted_ses_smtp_password_v4 of aws_iam_access_key.
func (iak iamAccessKeyAttributes) EncryptedSesSmtpPasswordV4() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("encrypted_ses_smtp_password_v4"))
}

// Id returns a reference to field id of aws_iam_access_key.
func (iak iamAccessKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("id"))
}

// KeyFingerprint returns a reference to field key_fingerprint of aws_iam_access_key.
func (iak iamAccessKeyAttributes) KeyFingerprint() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("key_fingerprint"))
}

// PgpKey returns a reference to field pgp_key of aws_iam_access_key.
func (iak iamAccessKeyAttributes) PgpKey() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("pgp_key"))
}

// Secret returns a reference to field secret of aws_iam_access_key.
func (iak iamAccessKeyAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("secret"))
}

// SesSmtpPasswordV4 returns a reference to field ses_smtp_password_v4 of aws_iam_access_key.
func (iak iamAccessKeyAttributes) SesSmtpPasswordV4() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("ses_smtp_password_v4"))
}

// Status returns a reference to field status of aws_iam_access_key.
func (iak iamAccessKeyAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("status"))
}

// User returns a reference to field user of aws_iam_access_key.
func (iak iamAccessKeyAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(iak.ref.Append("user"))
}

type iamAccessKeyState struct {
	CreateDate                 string `json:"create_date"`
	EncryptedSecret            string `json:"encrypted_secret"`
	EncryptedSesSmtpPasswordV4 string `json:"encrypted_ses_smtp_password_v4"`
	Id                         string `json:"id"`
	KeyFingerprint             string `json:"key_fingerprint"`
	PgpKey                     string `json:"pgp_key"`
	Secret                     string `json:"secret"`
	SesSmtpPasswordV4          string `json:"ses_smtp_password_v4"`
	Status                     string `json:"status"`
	User                       string `json:"user"`
}
