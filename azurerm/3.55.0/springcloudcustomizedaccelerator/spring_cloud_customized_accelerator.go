// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package springcloudcustomizedaccelerator

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GitRepository struct {
	// Branch: string, optional
	Branch terra.StringValue `hcl:"branch,attr"`
	// Commit: string, optional
	Commit terra.StringValue `hcl:"commit,attr"`
	// GitTag: string, optional
	GitTag terra.StringValue `hcl:"git_tag,attr"`
	// IntervalInSeconds: number, optional
	IntervalInSeconds terra.NumberValue `hcl:"interval_in_seconds,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
	// BasicAuth: optional
	BasicAuth *BasicAuth `hcl:"basic_auth,block"`
	// SshAuth: optional
	SshAuth *SshAuth `hcl:"ssh_auth,block"`
}

type BasicAuth struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type SshAuth struct {
	// HostKey: string, optional
	HostKey terra.StringValue `hcl:"host_key,attr"`
	// HostKeyAlgorithm: string, optional
	HostKeyAlgorithm terra.StringValue `hcl:"host_key_algorithm,attr"`
	// PrivateKey: string, required
	PrivateKey terra.StringValue `hcl:"private_key,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type GitRepositoryAttributes struct {
	ref terra.Reference
}

func (gr GitRepositoryAttributes) InternalRef() (terra.Reference, error) {
	return gr.ref, nil
}

func (gr GitRepositoryAttributes) InternalWithRef(ref terra.Reference) GitRepositoryAttributes {
	return GitRepositoryAttributes{ref: ref}
}

func (gr GitRepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gr.ref.InternalTokens()
}

func (gr GitRepositoryAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("branch"))
}

func (gr GitRepositoryAttributes) Commit() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("commit"))
}

func (gr GitRepositoryAttributes) GitTag() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("git_tag"))
}

func (gr GitRepositoryAttributes) IntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(gr.ref.Append("interval_in_seconds"))
}

func (gr GitRepositoryAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(gr.ref.Append("url"))
}

func (gr GitRepositoryAttributes) BasicAuth() terra.ListValue[BasicAuthAttributes] {
	return terra.ReferenceAsList[BasicAuthAttributes](gr.ref.Append("basic_auth"))
}

func (gr GitRepositoryAttributes) SshAuth() terra.ListValue[SshAuthAttributes] {
	return terra.ReferenceAsList[SshAuthAttributes](gr.ref.Append("ssh_auth"))
}

type BasicAuthAttributes struct {
	ref terra.Reference
}

func (ba BasicAuthAttributes) InternalRef() (terra.Reference, error) {
	return ba.ref, nil
}

func (ba BasicAuthAttributes) InternalWithRef(ref terra.Reference) BasicAuthAttributes {
	return BasicAuthAttributes{ref: ref}
}

func (ba BasicAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ba.ref.InternalTokens()
}

func (ba BasicAuthAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("password"))
}

func (ba BasicAuthAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("username"))
}

type SshAuthAttributes struct {
	ref terra.Reference
}

func (sa SshAuthAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa SshAuthAttributes) InternalWithRef(ref terra.Reference) SshAuthAttributes {
	return SshAuthAttributes{ref: ref}
}

func (sa SshAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa SshAuthAttributes) HostKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key"))
}

func (sa SshAuthAttributes) HostKeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key_algorithm"))
}

func (sa SshAuthAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("private_key"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type GitRepositoryState struct {
	Branch            string           `json:"branch"`
	Commit            string           `json:"commit"`
	GitTag            string           `json:"git_tag"`
	IntervalInSeconds float64          `json:"interval_in_seconds"`
	Url               string           `json:"url"`
	BasicAuth         []BasicAuthState `json:"basic_auth"`
	SshAuth           []SshAuthState   `json:"ssh_auth"`
}

type BasicAuthState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SshAuthState struct {
	HostKey          string `json:"host_key"`
	HostKeyAlgorithm string `json:"host_key_algorithm"`
	PrivateKey       string `json:"private_key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
