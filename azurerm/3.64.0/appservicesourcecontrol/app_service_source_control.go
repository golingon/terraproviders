// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appservicesourcecontrol

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GithubActionConfiguration struct {
	// GenerateWorkflowFile: bool, optional
	GenerateWorkflowFile terra.BoolValue `hcl:"generate_workflow_file,attr"`
	// CodeConfiguration: optional
	CodeConfiguration *CodeConfiguration `hcl:"code_configuration,block"`
	// ContainerConfiguration: optional
	ContainerConfiguration *ContainerConfiguration `hcl:"container_configuration,block"`
}

type CodeConfiguration struct {
	// RuntimeStack: string, required
	RuntimeStack terra.StringValue `hcl:"runtime_stack,attr" validate:"required"`
	// RuntimeVersion: string, required
	RuntimeVersion terra.StringValue `hcl:"runtime_version,attr" validate:"required"`
}

type ContainerConfiguration struct {
	// ImageName: string, required
	ImageName terra.StringValue `hcl:"image_name,attr" validate:"required"`
	// RegistryPassword: string, optional
	RegistryPassword terra.StringValue `hcl:"registry_password,attr"`
	// RegistryUrl: string, required
	RegistryUrl terra.StringValue `hcl:"registry_url,attr" validate:"required"`
	// RegistryUsername: string, optional
	RegistryUsername terra.StringValue `hcl:"registry_username,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type GithubActionConfigurationAttributes struct {
	ref terra.Reference
}

func (gac GithubActionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return gac.ref, nil
}

func (gac GithubActionConfigurationAttributes) InternalWithRef(ref terra.Reference) GithubActionConfigurationAttributes {
	return GithubActionConfigurationAttributes{ref: ref}
}

func (gac GithubActionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gac.ref.InternalTokens()
}

func (gac GithubActionConfigurationAttributes) GenerateWorkflowFile() terra.BoolValue {
	return terra.ReferenceAsBool(gac.ref.Append("generate_workflow_file"))
}

func (gac GithubActionConfigurationAttributes) LinuxAction() terra.BoolValue {
	return terra.ReferenceAsBool(gac.ref.Append("linux_action"))
}

func (gac GithubActionConfigurationAttributes) CodeConfiguration() terra.ListValue[CodeConfigurationAttributes] {
	return terra.ReferenceAsList[CodeConfigurationAttributes](gac.ref.Append("code_configuration"))
}

func (gac GithubActionConfigurationAttributes) ContainerConfiguration() terra.ListValue[ContainerConfigurationAttributes] {
	return terra.ReferenceAsList[ContainerConfigurationAttributes](gac.ref.Append("container_configuration"))
}

type CodeConfigurationAttributes struct {
	ref terra.Reference
}

func (cc CodeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CodeConfigurationAttributes) InternalWithRef(ref terra.Reference) CodeConfigurationAttributes {
	return CodeConfigurationAttributes{ref: ref}
}

func (cc CodeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CodeConfigurationAttributes) RuntimeStack() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("runtime_stack"))
}

func (cc CodeConfigurationAttributes) RuntimeVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("runtime_version"))
}

type ContainerConfigurationAttributes struct {
	ref terra.Reference
}

func (cc ContainerConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc ContainerConfigurationAttributes) InternalWithRef(ref terra.Reference) ContainerConfigurationAttributes {
	return ContainerConfigurationAttributes{ref: ref}
}

func (cc ContainerConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc ContainerConfigurationAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("image_name"))
}

func (cc ContainerConfigurationAttributes) RegistryPassword() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("registry_password"))
}

func (cc ContainerConfigurationAttributes) RegistryUrl() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("registry_url"))
}

func (cc ContainerConfigurationAttributes) RegistryUsername() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("registry_username"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type GithubActionConfigurationState struct {
	GenerateWorkflowFile   bool                          `json:"generate_workflow_file"`
	LinuxAction            bool                          `json:"linux_action"`
	CodeConfiguration      []CodeConfigurationState      `json:"code_configuration"`
	ContainerConfiguration []ContainerConfigurationState `json:"container_configuration"`
}

type CodeConfigurationState struct {
	RuntimeStack   string `json:"runtime_stack"`
	RuntimeVersion string `json:"runtime_version"`
}

type ContainerConfigurationState struct {
	ImageName        string `json:"image_name"`
	RegistryPassword string `json:"registry_password"`
	RegistryUrl      string `json:"registry_url"`
	RegistryUsername string `json:"registry_username"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
