// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	iapbrand "github.com/golingon/terraproviders/google/4.71.0/iapbrand"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapBrand creates a new instance of [IapBrand].
func NewIapBrand(name string, args IapBrandArgs) *IapBrand {
	return &IapBrand{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapBrand)(nil)

// IapBrand represents the Terraform resource google_iap_brand.
type IapBrand struct {
	Name      string
	Args      IapBrandArgs
	state     *iapBrandState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapBrand].
func (ib *IapBrand) Type() string {
	return "google_iap_brand"
}

// LocalName returns the local name for [IapBrand].
func (ib *IapBrand) LocalName() string {
	return ib.Name
}

// Configuration returns the configuration (args) for [IapBrand].
func (ib *IapBrand) Configuration() interface{} {
	return ib.Args
}

// DependOn is used for other resources to depend on [IapBrand].
func (ib *IapBrand) DependOn() terra.Reference {
	return terra.ReferenceResource(ib)
}

// Dependencies returns the list of resources [IapBrand] depends_on.
func (ib *IapBrand) Dependencies() terra.Dependencies {
	return ib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapBrand].
func (ib *IapBrand) LifecycleManagement() *terra.Lifecycle {
	return ib.Lifecycle
}

// Attributes returns the attributes for [IapBrand].
func (ib *IapBrand) Attributes() iapBrandAttributes {
	return iapBrandAttributes{ref: terra.ReferenceResource(ib)}
}

// ImportState imports the given attribute values into [IapBrand]'s state.
func (ib *IapBrand) ImportState(av io.Reader) error {
	ib.state = &iapBrandState{}
	if err := json.NewDecoder(av).Decode(ib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ib.Type(), ib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapBrand] has state.
func (ib *IapBrand) State() (*iapBrandState, bool) {
	return ib.state, ib.state != nil
}

// StateMust returns the state for [IapBrand]. Panics if the state is nil.
func (ib *IapBrand) StateMust() *iapBrandState {
	if ib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ib.Type(), ib.LocalName()))
	}
	return ib.state
}

// IapBrandArgs contains the configurations for google_iap_brand.
type IapBrandArgs struct {
	// ApplicationTitle: string, required
	ApplicationTitle terra.StringValue `hcl:"application_title,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SupportEmail: string, required
	SupportEmail terra.StringValue `hcl:"support_email,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iapbrand.Timeouts `hcl:"timeouts,block"`
}
type iapBrandAttributes struct {
	ref terra.Reference
}

// ApplicationTitle returns a reference to field application_title of google_iap_brand.
func (ib iapBrandAttributes) ApplicationTitle() terra.StringValue {
	return terra.ReferenceAsString(ib.ref.Append("application_title"))
}

// Id returns a reference to field id of google_iap_brand.
func (ib iapBrandAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ib.ref.Append("id"))
}

// Name returns a reference to field name of google_iap_brand.
func (ib iapBrandAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ib.ref.Append("name"))
}

// OrgInternalOnly returns a reference to field org_internal_only of google_iap_brand.
func (ib iapBrandAttributes) OrgInternalOnly() terra.BoolValue {
	return terra.ReferenceAsBool(ib.ref.Append("org_internal_only"))
}

// Project returns a reference to field project of google_iap_brand.
func (ib iapBrandAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ib.ref.Append("project"))
}

// SupportEmail returns a reference to field support_email of google_iap_brand.
func (ib iapBrandAttributes) SupportEmail() terra.StringValue {
	return terra.ReferenceAsString(ib.ref.Append("support_email"))
}

func (ib iapBrandAttributes) Timeouts() iapbrand.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iapbrand.TimeoutsAttributes](ib.ref.Append("timeouts"))
}

type iapBrandState struct {
	ApplicationTitle string                  `json:"application_title"`
	Id               string                  `json:"id"`
	Name             string                  `json:"name"`
	OrgInternalOnly  bool                    `json:"org_internal_only"`
	Project          string                  `json:"project"`
	SupportEmail     string                  `json:"support_email"`
	Timeouts         *iapbrand.TimeoutsState `json:"timeouts"`
}
