// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunv2jobiambinding "github.com/golingon/terraproviders/google/4.59.0/cloudrunv2jobiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudRunV2JobIamBinding(name string, args CloudRunV2JobIamBindingArgs) *CloudRunV2JobIamBinding {
	return &CloudRunV2JobIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2JobIamBinding)(nil)

type CloudRunV2JobIamBinding struct {
	Name  string
	Args  CloudRunV2JobIamBindingArgs
	state *cloudRunV2JobIamBindingState
}

func (crvjib *CloudRunV2JobIamBinding) Type() string {
	return "google_cloud_run_v2_job_iam_binding"
}

func (crvjib *CloudRunV2JobIamBinding) LocalName() string {
	return crvjib.Name
}

func (crvjib *CloudRunV2JobIamBinding) Configuration() interface{} {
	return crvjib.Args
}

func (crvjib *CloudRunV2JobIamBinding) Attributes() cloudRunV2JobIamBindingAttributes {
	return cloudRunV2JobIamBindingAttributes{ref: terra.ReferenceResource(crvjib)}
}

func (crvjib *CloudRunV2JobIamBinding) ImportState(av io.Reader) error {
	crvjib.state = &cloudRunV2JobIamBindingState{}
	if err := json.NewDecoder(av).Decode(crvjib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvjib.Type(), crvjib.LocalName(), err)
	}
	return nil
}

func (crvjib *CloudRunV2JobIamBinding) State() (*cloudRunV2JobIamBindingState, bool) {
	return crvjib.state, crvjib.state != nil
}

func (crvjib *CloudRunV2JobIamBinding) StateMust() *cloudRunV2JobIamBindingState {
	if crvjib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvjib.Type(), crvjib.LocalName()))
	}
	return crvjib.state
}

func (crvjib *CloudRunV2JobIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(crvjib)
}

type CloudRunV2JobIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudrunv2jobiambinding.Condition `hcl:"condition,block"`
	// DependsOn contains resources that CloudRunV2JobIamBinding depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudRunV2JobIamBindingAttributes struct {
	ref terra.Reference
}

func (crvjib cloudRunV2JobIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("etag"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("id"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("location"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](crvjib.ref.Append("members"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("name"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("project"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceString(crvjib.ref.Append("role"))
}

func (crvjib cloudRunV2JobIamBindingAttributes) Condition() terra.ListValue[cloudrunv2jobiambinding.ConditionAttributes] {
	return terra.ReferenceList[cloudrunv2jobiambinding.ConditionAttributes](crvjib.ref.Append("condition"))
}

type cloudRunV2JobIamBindingState struct {
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Location  string                                   `json:"location"`
	Members   []string                                 `json:"members"`
	Name      string                                   `json:"name"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	Condition []cloudrunv2jobiambinding.ConditionState `json:"condition"`
}
