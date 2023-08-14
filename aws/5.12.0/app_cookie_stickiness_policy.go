// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppCookieStickinessPolicy creates a new instance of [AppCookieStickinessPolicy].
func NewAppCookieStickinessPolicy(name string, args AppCookieStickinessPolicyArgs) *AppCookieStickinessPolicy {
	return &AppCookieStickinessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppCookieStickinessPolicy)(nil)

// AppCookieStickinessPolicy represents the Terraform resource aws_app_cookie_stickiness_policy.
type AppCookieStickinessPolicy struct {
	Name      string
	Args      AppCookieStickinessPolicyArgs
	state     *appCookieStickinessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) Type() string {
	return "aws_app_cookie_stickiness_policy"
}

// LocalName returns the local name for [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) LocalName() string {
	return acsp.Name
}

// Configuration returns the configuration (args) for [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) Configuration() interface{} {
	return acsp.Args
}

// DependOn is used for other resources to depend on [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(acsp)
}

// Dependencies returns the list of resources [AppCookieStickinessPolicy] depends_on.
func (acsp *AppCookieStickinessPolicy) Dependencies() terra.Dependencies {
	return acsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) LifecycleManagement() *terra.Lifecycle {
	return acsp.Lifecycle
}

// Attributes returns the attributes for [AppCookieStickinessPolicy].
func (acsp *AppCookieStickinessPolicy) Attributes() appCookieStickinessPolicyAttributes {
	return appCookieStickinessPolicyAttributes{ref: terra.ReferenceResource(acsp)}
}

// ImportState imports the given attribute values into [AppCookieStickinessPolicy]'s state.
func (acsp *AppCookieStickinessPolicy) ImportState(av io.Reader) error {
	acsp.state = &appCookieStickinessPolicyState{}
	if err := json.NewDecoder(av).Decode(acsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acsp.Type(), acsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppCookieStickinessPolicy] has state.
func (acsp *AppCookieStickinessPolicy) State() (*appCookieStickinessPolicyState, bool) {
	return acsp.state, acsp.state != nil
}

// StateMust returns the state for [AppCookieStickinessPolicy]. Panics if the state is nil.
func (acsp *AppCookieStickinessPolicy) StateMust() *appCookieStickinessPolicyState {
	if acsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acsp.Type(), acsp.LocalName()))
	}
	return acsp.state
}

// AppCookieStickinessPolicyArgs contains the configurations for aws_app_cookie_stickiness_policy.
type AppCookieStickinessPolicyArgs struct {
	// CookieName: string, required
	CookieName terra.StringValue `hcl:"cookie_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbPort: number, required
	LbPort terra.NumberValue `hcl:"lb_port,attr" validate:"required"`
	// LoadBalancer: string, required
	LoadBalancer terra.StringValue `hcl:"load_balancer,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type appCookieStickinessPolicyAttributes struct {
	ref terra.Reference
}

// CookieName returns a reference to field cookie_name of aws_app_cookie_stickiness_policy.
func (acsp appCookieStickinessPolicyAttributes) CookieName() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("cookie_name"))
}

// Id returns a reference to field id of aws_app_cookie_stickiness_policy.
func (acsp appCookieStickinessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("id"))
}

// LbPort returns a reference to field lb_port of aws_app_cookie_stickiness_policy.
func (acsp appCookieStickinessPolicyAttributes) LbPort() terra.NumberValue {
	return terra.ReferenceAsNumber(acsp.ref.Append("lb_port"))
}

// LoadBalancer returns a reference to field load_balancer of aws_app_cookie_stickiness_policy.
func (acsp appCookieStickinessPolicyAttributes) LoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("load_balancer"))
}

// Name returns a reference to field name of aws_app_cookie_stickiness_policy.
func (acsp appCookieStickinessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acsp.ref.Append("name"))
}

type appCookieStickinessPolicyState struct {
	CookieName   string  `json:"cookie_name"`
	Id           string  `json:"id"`
	LbPort       float64 `json:"lb_port"`
	LoadBalancer string  `json:"load_balancer"`
	Name         string  `json:"name"`
}