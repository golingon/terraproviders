// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	billingprojectinfo "github.com/golingon/terraproviders/google/5.24.0/billingprojectinfo"
	"io"
)

// NewBillingProjectInfo creates a new instance of [BillingProjectInfo].
func NewBillingProjectInfo(name string, args BillingProjectInfoArgs) *BillingProjectInfo {
	return &BillingProjectInfo{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingProjectInfo)(nil)

// BillingProjectInfo represents the Terraform resource google_billing_project_info.
type BillingProjectInfo struct {
	Name      string
	Args      BillingProjectInfoArgs
	state     *billingProjectInfoState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingProjectInfo].
func (bpi *BillingProjectInfo) Type() string {
	return "google_billing_project_info"
}

// LocalName returns the local name for [BillingProjectInfo].
func (bpi *BillingProjectInfo) LocalName() string {
	return bpi.Name
}

// Configuration returns the configuration (args) for [BillingProjectInfo].
func (bpi *BillingProjectInfo) Configuration() interface{} {
	return bpi.Args
}

// DependOn is used for other resources to depend on [BillingProjectInfo].
func (bpi *BillingProjectInfo) DependOn() terra.Reference {
	return terra.ReferenceResource(bpi)
}

// Dependencies returns the list of resources [BillingProjectInfo] depends_on.
func (bpi *BillingProjectInfo) Dependencies() terra.Dependencies {
	return bpi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingProjectInfo].
func (bpi *BillingProjectInfo) LifecycleManagement() *terra.Lifecycle {
	return bpi.Lifecycle
}

// Attributes returns the attributes for [BillingProjectInfo].
func (bpi *BillingProjectInfo) Attributes() billingProjectInfoAttributes {
	return billingProjectInfoAttributes{ref: terra.ReferenceResource(bpi)}
}

// ImportState imports the given attribute values into [BillingProjectInfo]'s state.
func (bpi *BillingProjectInfo) ImportState(av io.Reader) error {
	bpi.state = &billingProjectInfoState{}
	if err := json.NewDecoder(av).Decode(bpi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bpi.Type(), bpi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingProjectInfo] has state.
func (bpi *BillingProjectInfo) State() (*billingProjectInfoState, bool) {
	return bpi.state, bpi.state != nil
}

// StateMust returns the state for [BillingProjectInfo]. Panics if the state is nil.
func (bpi *BillingProjectInfo) StateMust() *billingProjectInfoState {
	if bpi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bpi.Type(), bpi.LocalName()))
	}
	return bpi.state
}

// BillingProjectInfoArgs contains the configurations for google_billing_project_info.
type BillingProjectInfoArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *billingprojectinfo.Timeouts `hcl:"timeouts,block"`
}
type billingProjectInfoAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_billing_project_info.
func (bpi billingProjectInfoAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(bpi.ref.Append("billing_account"))
}

// Id returns a reference to field id of google_billing_project_info.
func (bpi billingProjectInfoAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpi.ref.Append("id"))
}

// Project returns a reference to field project of google_billing_project_info.
func (bpi billingProjectInfoAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bpi.ref.Append("project"))
}

func (bpi billingProjectInfoAttributes) Timeouts() billingprojectinfo.TimeoutsAttributes {
	return terra.ReferenceAsSingle[billingprojectinfo.TimeoutsAttributes](bpi.ref.Append("timeouts"))
}

type billingProjectInfoState struct {
	BillingAccount string                            `json:"billing_account"`
	Id             string                            `json:"id"`
	Project        string                            `json:"project"`
	Timeouts       *billingprojectinfo.TimeoutsState `json:"timeouts"`
}
