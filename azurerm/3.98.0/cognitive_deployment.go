// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	cognitivedeployment "github.com/golingon/terraproviders/azurerm/3.98.0/cognitivedeployment"
	"io"
)

// NewCognitiveDeployment creates a new instance of [CognitiveDeployment].
func NewCognitiveDeployment(name string, args CognitiveDeploymentArgs) *CognitiveDeployment {
	return &CognitiveDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitiveDeployment)(nil)

// CognitiveDeployment represents the Terraform resource azurerm_cognitive_deployment.
type CognitiveDeployment struct {
	Name      string
	Args      CognitiveDeploymentArgs
	state     *cognitiveDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitiveDeployment].
func (cd *CognitiveDeployment) Type() string {
	return "azurerm_cognitive_deployment"
}

// LocalName returns the local name for [CognitiveDeployment].
func (cd *CognitiveDeployment) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [CognitiveDeployment].
func (cd *CognitiveDeployment) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [CognitiveDeployment].
func (cd *CognitiveDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [CognitiveDeployment] depends_on.
func (cd *CognitiveDeployment) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitiveDeployment].
func (cd *CognitiveDeployment) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [CognitiveDeployment].
func (cd *CognitiveDeployment) Attributes() cognitiveDeploymentAttributes {
	return cognitiveDeploymentAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [CognitiveDeployment]'s state.
func (cd *CognitiveDeployment) ImportState(av io.Reader) error {
	cd.state = &cognitiveDeploymentState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitiveDeployment] has state.
func (cd *CognitiveDeployment) State() (*cognitiveDeploymentState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [CognitiveDeployment]. Panics if the state is nil.
func (cd *CognitiveDeployment) StateMust() *cognitiveDeploymentState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// CognitiveDeploymentArgs contains the configurations for azurerm_cognitive_deployment.
type CognitiveDeploymentArgs struct {
	// CognitiveAccountId: string, required
	CognitiveAccountId terra.StringValue `hcl:"cognitive_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RaiPolicyName: string, optional
	RaiPolicyName terra.StringValue `hcl:"rai_policy_name,attr"`
	// VersionUpgradeOption: string, optional
	VersionUpgradeOption terra.StringValue `hcl:"version_upgrade_option,attr"`
	// Model: required
	Model *cognitivedeployment.Model `hcl:"model,block" validate:"required"`
	// Scale: required
	Scale *cognitivedeployment.Scale `hcl:"scale,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cognitivedeployment.Timeouts `hcl:"timeouts,block"`
}
type cognitiveDeploymentAttributes struct {
	ref terra.Reference
}

// CognitiveAccountId returns a reference to field cognitive_account_id of azurerm_cognitive_deployment.
func (cd cognitiveDeploymentAttributes) CognitiveAccountId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("cognitive_account_id"))
}

// Id returns a reference to field id of azurerm_cognitive_deployment.
func (cd cognitiveDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cognitive_deployment.
func (cd cognitiveDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

// RaiPolicyName returns a reference to field rai_policy_name of azurerm_cognitive_deployment.
func (cd cognitiveDeploymentAttributes) RaiPolicyName() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("rai_policy_name"))
}

// VersionUpgradeOption returns a reference to field version_upgrade_option of azurerm_cognitive_deployment.
func (cd cognitiveDeploymentAttributes) VersionUpgradeOption() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("version_upgrade_option"))
}

func (cd cognitiveDeploymentAttributes) Model() terra.ListValue[cognitivedeployment.ModelAttributes] {
	return terra.ReferenceAsList[cognitivedeployment.ModelAttributes](cd.ref.Append("model"))
}

func (cd cognitiveDeploymentAttributes) Scale() terra.ListValue[cognitivedeployment.ScaleAttributes] {
	return terra.ReferenceAsList[cognitivedeployment.ScaleAttributes](cd.ref.Append("scale"))
}

func (cd cognitiveDeploymentAttributes) Timeouts() cognitivedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cognitivedeployment.TimeoutsAttributes](cd.ref.Append("timeouts"))
}

type cognitiveDeploymentState struct {
	CognitiveAccountId   string                             `json:"cognitive_account_id"`
	Id                   string                             `json:"id"`
	Name                 string                             `json:"name"`
	RaiPolicyName        string                             `json:"rai_policy_name"`
	VersionUpgradeOption string                             `json:"version_upgrade_option"`
	Model                []cognitivedeployment.ModelState   `json:"model"`
	Scale                []cognitivedeployment.ScaleState   `json:"scale"`
	Timeouts             *cognitivedeployment.TimeoutsState `json:"timeouts"`
}
