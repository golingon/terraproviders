// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	quicksightaccountsubscription "github.com/golingon/terraproviders/aws/5.11.0/quicksightaccountsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightAccountSubscription creates a new instance of [QuicksightAccountSubscription].
func NewQuicksightAccountSubscription(name string, args QuicksightAccountSubscriptionArgs) *QuicksightAccountSubscription {
	return &QuicksightAccountSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightAccountSubscription)(nil)

// QuicksightAccountSubscription represents the Terraform resource aws_quicksight_account_subscription.
type QuicksightAccountSubscription struct {
	Name      string
	Args      QuicksightAccountSubscriptionArgs
	state     *quicksightAccountSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) Type() string {
	return "aws_quicksight_account_subscription"
}

// LocalName returns the local name for [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) LocalName() string {
	return qas.Name
}

// Configuration returns the configuration (args) for [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) Configuration() interface{} {
	return qas.Args
}

// DependOn is used for other resources to depend on [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(qas)
}

// Dependencies returns the list of resources [QuicksightAccountSubscription] depends_on.
func (qas *QuicksightAccountSubscription) Dependencies() terra.Dependencies {
	return qas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) LifecycleManagement() *terra.Lifecycle {
	return qas.Lifecycle
}

// Attributes returns the attributes for [QuicksightAccountSubscription].
func (qas *QuicksightAccountSubscription) Attributes() quicksightAccountSubscriptionAttributes {
	return quicksightAccountSubscriptionAttributes{ref: terra.ReferenceResource(qas)}
}

// ImportState imports the given attribute values into [QuicksightAccountSubscription]'s state.
func (qas *QuicksightAccountSubscription) ImportState(av io.Reader) error {
	qas.state = &quicksightAccountSubscriptionState{}
	if err := json.NewDecoder(av).Decode(qas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qas.Type(), qas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightAccountSubscription] has state.
func (qas *QuicksightAccountSubscription) State() (*quicksightAccountSubscriptionState, bool) {
	return qas.state, qas.state != nil
}

// StateMust returns the state for [QuicksightAccountSubscription]. Panics if the state is nil.
func (qas *QuicksightAccountSubscription) StateMust() *quicksightAccountSubscriptionState {
	if qas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qas.Type(), qas.LocalName()))
	}
	return qas.state
}

// QuicksightAccountSubscriptionArgs contains the configurations for aws_quicksight_account_subscription.
type QuicksightAccountSubscriptionArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// ActiveDirectoryName: string, optional
	ActiveDirectoryName terra.StringValue `hcl:"active_directory_name,attr"`
	// AdminGroup: list of string, optional
	AdminGroup terra.ListValue[terra.StringValue] `hcl:"admin_group,attr"`
	// AuthenticationMethod: string, required
	AuthenticationMethod terra.StringValue `hcl:"authentication_method,attr" validate:"required"`
	// AuthorGroup: list of string, optional
	AuthorGroup terra.ListValue[terra.StringValue] `hcl:"author_group,attr"`
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// ContactNumber: string, optional
	ContactNumber terra.StringValue `hcl:"contact_number,attr"`
	// DirectoryId: string, optional
	DirectoryId terra.StringValue `hcl:"directory_id,attr"`
	// Edition: string, required
	Edition terra.StringValue `hcl:"edition,attr" validate:"required"`
	// EmailAddress: string, optional
	EmailAddress terra.StringValue `hcl:"email_address,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// NotificationEmail: string, required
	NotificationEmail terra.StringValue `hcl:"notification_email,attr" validate:"required"`
	// ReaderGroup: list of string, optional
	ReaderGroup terra.ListValue[terra.StringValue] `hcl:"reader_group,attr"`
	// Realm: string, optional
	Realm terra.StringValue `hcl:"realm,attr"`
	// Timeouts: optional
	Timeouts *quicksightaccountsubscription.Timeouts `hcl:"timeouts,block"`
}
type quicksightAccountSubscriptionAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("account_name"))
}

// AccountSubscriptionStatus returns a reference to field account_subscription_status of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AccountSubscriptionStatus() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("account_subscription_status"))
}

// ActiveDirectoryName returns a reference to field active_directory_name of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) ActiveDirectoryName() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("active_directory_name"))
}

// AdminGroup returns a reference to field admin_group of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AdminGroup() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](qas.ref.Append("admin_group"))
}

// AuthenticationMethod returns a reference to field authentication_method of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AuthenticationMethod() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("authentication_method"))
}

// AuthorGroup returns a reference to field author_group of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AuthorGroup() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](qas.ref.Append("author_group"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("aws_account_id"))
}

// ContactNumber returns a reference to field contact_number of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) ContactNumber() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("contact_number"))
}

// DirectoryId returns a reference to field directory_id of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("directory_id"))
}

// Edition returns a reference to field edition of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("edition"))
}

// EmailAddress returns a reference to field email_address of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) EmailAddress() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("email_address"))
}

// FirstName returns a reference to field first_name of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("first_name"))
}

// Id returns a reference to field id of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("id"))
}

// LastName returns a reference to field last_name of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("last_name"))
}

// NotificationEmail returns a reference to field notification_email of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) NotificationEmail() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("notification_email"))
}

// ReaderGroup returns a reference to field reader_group of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) ReaderGroup() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](qas.ref.Append("reader_group"))
}

// Realm returns a reference to field realm of aws_quicksight_account_subscription.
func (qas quicksightAccountSubscriptionAttributes) Realm() terra.StringValue {
	return terra.ReferenceAsString(qas.ref.Append("realm"))
}

func (qas quicksightAccountSubscriptionAttributes) Timeouts() quicksightaccountsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[quicksightaccountsubscription.TimeoutsAttributes](qas.ref.Append("timeouts"))
}

type quicksightAccountSubscriptionState struct {
	AccountName               string                                       `json:"account_name"`
	AccountSubscriptionStatus string                                       `json:"account_subscription_status"`
	ActiveDirectoryName       string                                       `json:"active_directory_name"`
	AdminGroup                []string                                     `json:"admin_group"`
	AuthenticationMethod      string                                       `json:"authentication_method"`
	AuthorGroup               []string                                     `json:"author_group"`
	AwsAccountId              string                                       `json:"aws_account_id"`
	ContactNumber             string                                       `json:"contact_number"`
	DirectoryId               string                                       `json:"directory_id"`
	Edition                   string                                       `json:"edition"`
	EmailAddress              string                                       `json:"email_address"`
	FirstName                 string                                       `json:"first_name"`
	Id                        string                                       `json:"id"`
	LastName                  string                                       `json:"last_name"`
	NotificationEmail         string                                       `json:"notification_email"`
	ReaderGroup               []string                                     `json:"reader_group"`
	Realm                     string                                       `json:"realm"`
	Timeouts                  *quicksightaccountsubscription.TimeoutsState `json:"timeouts"`
}
