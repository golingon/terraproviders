// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	privatecacapool "github.com/golingon/terraproviders/googlebeta/4.64.0/privatecacapool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivatecaCaPool creates a new instance of [PrivatecaCaPool].
func NewPrivatecaCaPool(name string, args PrivatecaCaPoolArgs) *PrivatecaCaPool {
	return &PrivatecaCaPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCaPool)(nil)

// PrivatecaCaPool represents the Terraform resource google_privateca_ca_pool.
type PrivatecaCaPool struct {
	Name      string
	Args      PrivatecaCaPoolArgs
	state     *privatecaCaPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) Type() string {
	return "google_privateca_ca_pool"
}

// LocalName returns the local name for [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) LocalName() string {
	return pcp.Name
}

// Configuration returns the configuration (args) for [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) Configuration() interface{} {
	return pcp.Args
}

// DependOn is used for other resources to depend on [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) DependOn() terra.Reference {
	return terra.ReferenceResource(pcp)
}

// Dependencies returns the list of resources [PrivatecaCaPool] depends_on.
func (pcp *PrivatecaCaPool) Dependencies() terra.Dependencies {
	return pcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) LifecycleManagement() *terra.Lifecycle {
	return pcp.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCaPool].
func (pcp *PrivatecaCaPool) Attributes() privatecaCaPoolAttributes {
	return privatecaCaPoolAttributes{ref: terra.ReferenceResource(pcp)}
}

// ImportState imports the given attribute values into [PrivatecaCaPool]'s state.
func (pcp *PrivatecaCaPool) ImportState(av io.Reader) error {
	pcp.state = &privatecaCaPoolState{}
	if err := json.NewDecoder(av).Decode(pcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pcp.Type(), pcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCaPool] has state.
func (pcp *PrivatecaCaPool) State() (*privatecaCaPoolState, bool) {
	return pcp.state, pcp.state != nil
}

// StateMust returns the state for [PrivatecaCaPool]. Panics if the state is nil.
func (pcp *PrivatecaCaPool) StateMust() *privatecaCaPoolState {
	if pcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pcp.Type(), pcp.LocalName()))
	}
	return pcp.state
}

// PrivatecaCaPoolArgs contains the configurations for google_privateca_ca_pool.
type PrivatecaCaPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
	// IssuancePolicy: optional
	IssuancePolicy *privatecacapool.IssuancePolicy `hcl:"issuance_policy,block"`
	// PublishingOptions: optional
	PublishingOptions *privatecacapool.PublishingOptions `hcl:"publishing_options,block"`
	// Timeouts: optional
	Timeouts *privatecacapool.Timeouts `hcl:"timeouts,block"`
}
type privatecaCaPoolAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pcp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pcp.ref.Append("labels"))
}

// Location returns a reference to field location of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pcp.ref.Append("location"))
}

// Name returns a reference to field name of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pcp.ref.Append("name"))
}

// Project returns a reference to field project of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pcp.ref.Append("project"))
}

// Tier returns a reference to field tier of google_privateca_ca_pool.
func (pcp privatecaCaPoolAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(pcp.ref.Append("tier"))
}

func (pcp privatecaCaPoolAttributes) IssuancePolicy() terra.ListValue[privatecacapool.IssuancePolicyAttributes] {
	return terra.ReferenceAsList[privatecacapool.IssuancePolicyAttributes](pcp.ref.Append("issuance_policy"))
}

func (pcp privatecaCaPoolAttributes) PublishingOptions() terra.ListValue[privatecacapool.PublishingOptionsAttributes] {
	return terra.ReferenceAsList[privatecacapool.PublishingOptionsAttributes](pcp.ref.Append("publishing_options"))
}

func (pcp privatecaCaPoolAttributes) Timeouts() privatecacapool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatecacapool.TimeoutsAttributes](pcp.ref.Append("timeouts"))
}

type privatecaCaPoolState struct {
	Id                string                                   `json:"id"`
	Labels            map[string]string                        `json:"labels"`
	Location          string                                   `json:"location"`
	Name              string                                   `json:"name"`
	Project           string                                   `json:"project"`
	Tier              string                                   `json:"tier"`
	IssuancePolicy    []privatecacapool.IssuancePolicyState    `json:"issuance_policy"`
	PublishingOptions []privatecacapool.PublishingOptionsState `json:"publishing_options"`
	Timeouts          *privatecacapool.TimeoutsState           `json:"timeouts"`
}
