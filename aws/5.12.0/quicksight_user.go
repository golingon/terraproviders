// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightUser creates a new instance of [QuicksightUser].
func NewQuicksightUser(name string, args QuicksightUserArgs) *QuicksightUser {
	return &QuicksightUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightUser)(nil)

// QuicksightUser represents the Terraform resource aws_quicksight_user.
type QuicksightUser struct {
	Name      string
	Args      QuicksightUserArgs
	state     *quicksightUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightUser].
func (qu *QuicksightUser) Type() string {
	return "aws_quicksight_user"
}

// LocalName returns the local name for [QuicksightUser].
func (qu *QuicksightUser) LocalName() string {
	return qu.Name
}

// Configuration returns the configuration (args) for [QuicksightUser].
func (qu *QuicksightUser) Configuration() interface{} {
	return qu.Args
}

// DependOn is used for other resources to depend on [QuicksightUser].
func (qu *QuicksightUser) DependOn() terra.Reference {
	return terra.ReferenceResource(qu)
}

// Dependencies returns the list of resources [QuicksightUser] depends_on.
func (qu *QuicksightUser) Dependencies() terra.Dependencies {
	return qu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightUser].
func (qu *QuicksightUser) LifecycleManagement() *terra.Lifecycle {
	return qu.Lifecycle
}

// Attributes returns the attributes for [QuicksightUser].
func (qu *QuicksightUser) Attributes() quicksightUserAttributes {
	return quicksightUserAttributes{ref: terra.ReferenceResource(qu)}
}

// ImportState imports the given attribute values into [QuicksightUser]'s state.
func (qu *QuicksightUser) ImportState(av io.Reader) error {
	qu.state = &quicksightUserState{}
	if err := json.NewDecoder(av).Decode(qu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qu.Type(), qu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightUser] has state.
func (qu *QuicksightUser) State() (*quicksightUserState, bool) {
	return qu.state, qu.state != nil
}

// StateMust returns the state for [QuicksightUser]. Panics if the state is nil.
func (qu *QuicksightUser) StateMust() *quicksightUserState {
	if qu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qu.Type(), qu.LocalName()))
	}
	return qu.state
}

// QuicksightUserArgs contains the configurations for aws_quicksight_user.
type QuicksightUserArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// IamArn: string, optional
	IamArn terra.StringValue `hcl:"iam_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityType: string, required
	IdentityType terra.StringValue `hcl:"identity_type,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// SessionName: string, optional
	SessionName terra.StringValue `hcl:"session_name,attr"`
	// UserName: string, optional
	UserName terra.StringValue `hcl:"user_name,attr"`
	// UserRole: string, required
	UserRole terra.StringValue `hcl:"user_role,attr" validate:"required"`
}
type quicksightUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_user.
func (qu quicksightUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_user.
func (qu quicksightUserAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("aws_account_id"))
}

// Email returns a reference to field email of aws_quicksight_user.
func (qu quicksightUserAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("email"))
}

// IamArn returns a reference to field iam_arn of aws_quicksight_user.
func (qu quicksightUserAttributes) IamArn() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("iam_arn"))
}

// Id returns a reference to field id of aws_quicksight_user.
func (qu quicksightUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("id"))
}

// IdentityType returns a reference to field identity_type of aws_quicksight_user.
func (qu quicksightUserAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("identity_type"))
}

// Namespace returns a reference to field namespace of aws_quicksight_user.
func (qu quicksightUserAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("namespace"))
}

// SessionName returns a reference to field session_name of aws_quicksight_user.
func (qu quicksightUserAttributes) SessionName() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("session_name"))
}

// UserName returns a reference to field user_name of aws_quicksight_user.
func (qu quicksightUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("user_name"))
}

// UserRole returns a reference to field user_role of aws_quicksight_user.
func (qu quicksightUserAttributes) UserRole() terra.StringValue {
	return terra.ReferenceAsString(qu.ref.Append("user_role"))
}

type quicksightUserState struct {
	Arn          string `json:"arn"`
	AwsAccountId string `json:"aws_account_id"`
	Email        string `json:"email"`
	IamArn       string `json:"iam_arn"`
	Id           string `json:"id"`
	IdentityType string `json:"identity_type"`
	Namespace    string `json:"namespace"`
	SessionName  string `json:"session_name"`
	UserName     string `json:"user_name"`
	UserRole     string `json:"user_role"`
}