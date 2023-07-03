// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapclient "github.com/golingon/terraproviders/google/4.71.0/iapclient"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapClient creates a new instance of [IapClient].
func NewIapClient(name string, args IapClientArgs) *IapClient {
	return &IapClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapClient)(nil)

// IapClient represents the Terraform resource google_iap_client.
type IapClient struct {
	Name      string
	Args      IapClientArgs
	state     *iapClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapClient].
func (ic *IapClient) Type() string {
	return "google_iap_client"
}

// LocalName returns the local name for [IapClient].
func (ic *IapClient) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [IapClient].
func (ic *IapClient) Configuration() interface{} {
	return ic.Args
}

// DependOn is used for other resources to depend on [IapClient].
func (ic *IapClient) DependOn() terra.Reference {
	return terra.ReferenceResource(ic)
}

// Dependencies returns the list of resources [IapClient] depends_on.
func (ic *IapClient) Dependencies() terra.Dependencies {
	return ic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapClient].
func (ic *IapClient) LifecycleManagement() *terra.Lifecycle {
	return ic.Lifecycle
}

// Attributes returns the attributes for [IapClient].
func (ic *IapClient) Attributes() iapClientAttributes {
	return iapClientAttributes{ref: terra.ReferenceResource(ic)}
}

// ImportState imports the given attribute values into [IapClient]'s state.
func (ic *IapClient) ImportState(av io.Reader) error {
	ic.state = &iapClientState{}
	if err := json.NewDecoder(av).Decode(ic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ic.Type(), ic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapClient] has state.
func (ic *IapClient) State() (*iapClientState, bool) {
	return ic.state, ic.state != nil
}

// StateMust returns the state for [IapClient]. Panics if the state is nil.
func (ic *IapClient) StateMust() *iapClientState {
	if ic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ic.Type(), ic.LocalName()))
	}
	return ic.state
}

// IapClientArgs contains the configurations for google_iap_client.
type IapClientArgs struct {
	// Brand: string, required
	Brand terra.StringValue `hcl:"brand,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *iapclient.Timeouts `hcl:"timeouts,block"`
}
type iapClientAttributes struct {
	ref terra.Reference
}

// Brand returns a reference to field brand of google_iap_client.
func (ic iapClientAttributes) Brand() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("brand"))
}

// ClientId returns a reference to field client_id of google_iap_client.
func (ic iapClientAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("client_id"))
}

// DisplayName returns a reference to field display_name of google_iap_client.
func (ic iapClientAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iap_client.
func (ic iapClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// Secret returns a reference to field secret of google_iap_client.
func (ic iapClientAttributes) Secret() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("secret"))
}

func (ic iapClientAttributes) Timeouts() iapclient.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iapclient.TimeoutsAttributes](ic.ref.Append("timeouts"))
}

type iapClientState struct {
	Brand       string                   `json:"brand"`
	ClientId    string                   `json:"client_id"`
	DisplayName string                   `json:"display_name"`
	Id          string                   `json:"id"`
	Secret      string                   `json:"secret"`
	Timeouts    *iapclient.TimeoutsState `json:"timeouts"`
}
