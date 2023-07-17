// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerauthorizedorgsdesc "github.com/golingon/terraproviders/google/4.73.1/accesscontextmanagerauthorizedorgsdesc"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerAuthorizedOrgsDesc creates a new instance of [AccessContextManagerAuthorizedOrgsDesc].
func NewAccessContextManagerAuthorizedOrgsDesc(name string, args AccessContextManagerAuthorizedOrgsDescArgs) *AccessContextManagerAuthorizedOrgsDesc {
	return &AccessContextManagerAuthorizedOrgsDesc{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerAuthorizedOrgsDesc)(nil)

// AccessContextManagerAuthorizedOrgsDesc represents the Terraform resource google_access_context_manager_authorized_orgs_desc.
type AccessContextManagerAuthorizedOrgsDesc struct {
	Name      string
	Args      AccessContextManagerAuthorizedOrgsDescArgs
	state     *accessContextManagerAuthorizedOrgsDescState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) Type() string {
	return "google_access_context_manager_authorized_orgs_desc"
}

// LocalName returns the local name for [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) LocalName() string {
	return acmaod.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) Configuration() interface{} {
	return acmaod.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) DependOn() terra.Reference {
	return terra.ReferenceResource(acmaod)
}

// Dependencies returns the list of resources [AccessContextManagerAuthorizedOrgsDesc] depends_on.
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) Dependencies() terra.Dependencies {
	return acmaod.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) LifecycleManagement() *terra.Lifecycle {
	return acmaod.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerAuthorizedOrgsDesc].
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) Attributes() accessContextManagerAuthorizedOrgsDescAttributes {
	return accessContextManagerAuthorizedOrgsDescAttributes{ref: terra.ReferenceResource(acmaod)}
}

// ImportState imports the given attribute values into [AccessContextManagerAuthorizedOrgsDesc]'s state.
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) ImportState(av io.Reader) error {
	acmaod.state = &accessContextManagerAuthorizedOrgsDescState{}
	if err := json.NewDecoder(av).Decode(acmaod.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmaod.Type(), acmaod.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerAuthorizedOrgsDesc] has state.
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) State() (*accessContextManagerAuthorizedOrgsDescState, bool) {
	return acmaod.state, acmaod.state != nil
}

// StateMust returns the state for [AccessContextManagerAuthorizedOrgsDesc]. Panics if the state is nil.
func (acmaod *AccessContextManagerAuthorizedOrgsDesc) StateMust() *accessContextManagerAuthorizedOrgsDescState {
	if acmaod.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmaod.Type(), acmaod.LocalName()))
	}
	return acmaod.state
}

// AccessContextManagerAuthorizedOrgsDescArgs contains the configurations for google_access_context_manager_authorized_orgs_desc.
type AccessContextManagerAuthorizedOrgsDescArgs struct {
	// AssetType: string, optional
	AssetType terra.StringValue `hcl:"asset_type,attr"`
	// AuthorizationDirection: string, optional
	AuthorizationDirection terra.StringValue `hcl:"authorization_direction,attr"`
	// AuthorizationType: string, optional
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Orgs: list of string, optional
	Orgs terra.ListValue[terra.StringValue] `hcl:"orgs,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerauthorizedorgsdesc.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerAuthorizedOrgsDescAttributes struct {
	ref terra.Reference
}

// AssetType returns a reference to field asset_type of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) AssetType() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("asset_type"))
}

// AuthorizationDirection returns a reference to field authorization_direction of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) AuthorizationDirection() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("authorization_direction"))
}

// AuthorizationType returns a reference to field authorization_type of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("authorization_type"))
}

// CreateTime returns a reference to field create_time of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("create_time"))
}

// Id returns a reference to field id of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("name"))
}

// Orgs returns a reference to field orgs of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) Orgs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmaod.ref.Append("orgs"))
}

// Parent returns a reference to field parent of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("parent"))
}

// UpdateTime returns a reference to field update_time of google_access_context_manager_authorized_orgs_desc.
func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(acmaod.ref.Append("update_time"))
}

func (acmaod accessContextManagerAuthorizedOrgsDescAttributes) Timeouts() accesscontextmanagerauthorizedorgsdesc.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerauthorizedorgsdesc.TimeoutsAttributes](acmaod.ref.Append("timeouts"))
}

type accessContextManagerAuthorizedOrgsDescState struct {
	AssetType              string                                                `json:"asset_type"`
	AuthorizationDirection string                                                `json:"authorization_direction"`
	AuthorizationType      string                                                `json:"authorization_type"`
	CreateTime             string                                                `json:"create_time"`
	Id                     string                                                `json:"id"`
	Name                   string                                                `json:"name"`
	Orgs                   []string                                              `json:"orgs"`
	Parent                 string                                                `json:"parent"`
	UpdateTime             string                                                `json:"update_time"`
	Timeouts               *accesscontextmanagerauthorizedorgsdesc.TimeoutsState `json:"timeouts"`
}
