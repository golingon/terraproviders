// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataplextask

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExecutionStatus struct {
	// LatestJob: min=0
	LatestJob []LatestJob `hcl:"latest_job,block" validate:"min=0"`
}

type LatestJob struct{}

type ExecutionSpec struct {
	// Args: map of string, optional
	Args terra.MapValue[terra.StringValue] `hcl:"args,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// MaxJobExecutionLifetime: string, optional
	MaxJobExecutionLifetime terra.StringValue `hcl:"max_job_execution_lifetime,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceAccount: string, required
	ServiceAccount terra.StringValue `hcl:"service_account,attr" validate:"required"`
}

type Notebook struct {
	// ArchiveUris: list of string, optional
	ArchiveUris terra.ListValue[terra.StringValue] `hcl:"archive_uris,attr"`
	// FileUris: list of string, optional
	FileUris terra.ListValue[terra.StringValue] `hcl:"file_uris,attr"`
	// Notebook: string, required
	Notebook terra.StringValue `hcl:"notebook,attr" validate:"required"`
	// NotebookInfrastructureSpec: optional
	InfrastructureSpec *NotebookInfrastructureSpec `hcl:"infrastructure_spec,block"`
}

type NotebookInfrastructureSpec struct {
	// NotebookInfrastructureSpecBatch: optional
	Batch *NotebookInfrastructureSpecBatch `hcl:"batch,block"`
	// NotebookInfrastructureSpecContainerImage: optional
	ContainerImage *NotebookInfrastructureSpecContainerImage `hcl:"container_image,block"`
	// NotebookInfrastructureSpecVpcNetwork: optional
	VpcNetwork *NotebookInfrastructureSpecVpcNetwork `hcl:"vpc_network,block"`
}

type NotebookInfrastructureSpecBatch struct {
	// ExecutorsCount: number, optional
	ExecutorsCount terra.NumberValue `hcl:"executors_count,attr"`
	// MaxExecutorsCount: number, optional
	MaxExecutorsCount terra.NumberValue `hcl:"max_executors_count,attr"`
}

type NotebookInfrastructureSpecContainerImage struct {
	// Image: string, optional
	Image terra.StringValue `hcl:"image,attr"`
	// JavaJars: list of string, optional
	JavaJars terra.ListValue[terra.StringValue] `hcl:"java_jars,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// PythonPackages: list of string, optional
	PythonPackages terra.ListValue[terra.StringValue] `hcl:"python_packages,attr"`
}

type NotebookInfrastructureSpecVpcNetwork struct {
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkTags: list of string, optional
	NetworkTags terra.ListValue[terra.StringValue] `hcl:"network_tags,attr"`
	// SubNetwork: string, optional
	SubNetwork terra.StringValue `hcl:"sub_network,attr"`
}

type Spark struct {
	// ArchiveUris: list of string, optional
	ArchiveUris terra.ListValue[terra.StringValue] `hcl:"archive_uris,attr"`
	// FileUris: list of string, optional
	FileUris terra.ListValue[terra.StringValue] `hcl:"file_uris,attr"`
	// MainClass: string, optional
	MainClass terra.StringValue `hcl:"main_class,attr"`
	// MainJarFileUri: string, optional
	MainJarFileUri terra.StringValue `hcl:"main_jar_file_uri,attr"`
	// PythonScriptFile: string, optional
	PythonScriptFile terra.StringValue `hcl:"python_script_file,attr"`
	// SqlScript: string, optional
	SqlScript terra.StringValue `hcl:"sql_script,attr"`
	// SqlScriptFile: string, optional
	SqlScriptFile terra.StringValue `hcl:"sql_script_file,attr"`
	// SparkInfrastructureSpec: optional
	InfrastructureSpec *SparkInfrastructureSpec `hcl:"infrastructure_spec,block"`
}

type SparkInfrastructureSpec struct {
	// SparkInfrastructureSpecBatch: optional
	Batch *SparkInfrastructureSpecBatch `hcl:"batch,block"`
	// SparkInfrastructureSpecContainerImage: optional
	ContainerImage *SparkInfrastructureSpecContainerImage `hcl:"container_image,block"`
	// SparkInfrastructureSpecVpcNetwork: optional
	VpcNetwork *SparkInfrastructureSpecVpcNetwork `hcl:"vpc_network,block"`
}

type SparkInfrastructureSpecBatch struct {
	// ExecutorsCount: number, optional
	ExecutorsCount terra.NumberValue `hcl:"executors_count,attr"`
	// MaxExecutorsCount: number, optional
	MaxExecutorsCount terra.NumberValue `hcl:"max_executors_count,attr"`
}

type SparkInfrastructureSpecContainerImage struct {
	// Image: string, optional
	Image terra.StringValue `hcl:"image,attr"`
	// JavaJars: list of string, optional
	JavaJars terra.ListValue[terra.StringValue] `hcl:"java_jars,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// PythonPackages: list of string, optional
	PythonPackages terra.ListValue[terra.StringValue] `hcl:"python_packages,attr"`
}

type SparkInfrastructureSpecVpcNetwork struct {
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkTags: list of string, optional
	NetworkTags terra.ListValue[terra.StringValue] `hcl:"network_tags,attr"`
	// SubNetwork: string, optional
	SubNetwork terra.StringValue `hcl:"sub_network,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TriggerSpec struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// MaxRetries: number, optional
	MaxRetries terra.NumberValue `hcl:"max_retries,attr"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ExecutionStatusAttributes struct {
	ref terra.Reference
}

func (es ExecutionStatusAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es ExecutionStatusAttributes) InternalWithRef(ref terra.Reference) ExecutionStatusAttributes {
	return ExecutionStatusAttributes{ref: ref}
}

func (es ExecutionStatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es ExecutionStatusAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("update_time"))
}

func (es ExecutionStatusAttributes) LatestJob() terra.ListValue[LatestJobAttributes] {
	return terra.ReferenceAsList[LatestJobAttributes](es.ref.Append("latest_job"))
}

type LatestJobAttributes struct {
	ref terra.Reference
}

func (lj LatestJobAttributes) InternalRef() (terra.Reference, error) {
	return lj.ref, nil
}

func (lj LatestJobAttributes) InternalWithRef(ref terra.Reference) LatestJobAttributes {
	return LatestJobAttributes{ref: ref}
}

func (lj LatestJobAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lj.ref.InternalTokens()
}

func (lj LatestJobAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("end_time"))
}

func (lj LatestJobAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("message"))
}

func (lj LatestJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("name"))
}

func (lj LatestJobAttributes) RetryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(lj.ref.Append("retry_count"))
}

func (lj LatestJobAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("service"))
}

func (lj LatestJobAttributes) ServiceJob() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("service_job"))
}

func (lj LatestJobAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("start_time"))
}

func (lj LatestJobAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("state"))
}

func (lj LatestJobAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(lj.ref.Append("uid"))
}

type ExecutionSpecAttributes struct {
	ref terra.Reference
}

func (es ExecutionSpecAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es ExecutionSpecAttributes) InternalWithRef(ref terra.Reference) ExecutionSpecAttributes {
	return ExecutionSpecAttributes{ref: ref}
}

func (es ExecutionSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es ExecutionSpecAttributes) Args() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("args"))
}

func (es ExecutionSpecAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("kms_key"))
}

func (es ExecutionSpecAttributes) MaxJobExecutionLifetime() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("max_job_execution_lifetime"))
}

func (es ExecutionSpecAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("project"))
}

func (es ExecutionSpecAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("service_account"))
}

type NotebookAttributes struct {
	ref terra.Reference
}

func (n NotebookAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NotebookAttributes) InternalWithRef(ref terra.Reference) NotebookAttributes {
	return NotebookAttributes{ref: ref}
}

func (n NotebookAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NotebookAttributes) ArchiveUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("archive_uris"))
}

func (n NotebookAttributes) FileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("file_uris"))
}

func (n NotebookAttributes) Notebook() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("notebook"))
}

func (n NotebookAttributes) InfrastructureSpec() terra.ListValue[NotebookInfrastructureSpecAttributes] {
	return terra.ReferenceAsList[NotebookInfrastructureSpecAttributes](n.ref.Append("infrastructure_spec"))
}

type NotebookInfrastructureSpecAttributes struct {
	ref terra.Reference
}

func (is NotebookInfrastructureSpecAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is NotebookInfrastructureSpecAttributes) InternalWithRef(ref terra.Reference) NotebookInfrastructureSpecAttributes {
	return NotebookInfrastructureSpecAttributes{ref: ref}
}

func (is NotebookInfrastructureSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is NotebookInfrastructureSpecAttributes) Batch() terra.ListValue[NotebookInfrastructureSpecBatchAttributes] {
	return terra.ReferenceAsList[NotebookInfrastructureSpecBatchAttributes](is.ref.Append("batch"))
}

func (is NotebookInfrastructureSpecAttributes) ContainerImage() terra.ListValue[NotebookInfrastructureSpecContainerImageAttributes] {
	return terra.ReferenceAsList[NotebookInfrastructureSpecContainerImageAttributes](is.ref.Append("container_image"))
}

func (is NotebookInfrastructureSpecAttributes) VpcNetwork() terra.ListValue[NotebookInfrastructureSpecVpcNetworkAttributes] {
	return terra.ReferenceAsList[NotebookInfrastructureSpecVpcNetworkAttributes](is.ref.Append("vpc_network"))
}

type NotebookInfrastructureSpecBatchAttributes struct {
	ref terra.Reference
}

func (b NotebookInfrastructureSpecBatchAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b NotebookInfrastructureSpecBatchAttributes) InternalWithRef(ref terra.Reference) NotebookInfrastructureSpecBatchAttributes {
	return NotebookInfrastructureSpecBatchAttributes{ref: ref}
}

func (b NotebookInfrastructureSpecBatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b NotebookInfrastructureSpecBatchAttributes) ExecutorsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("executors_count"))
}

func (b NotebookInfrastructureSpecBatchAttributes) MaxExecutorsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_executors_count"))
}

type NotebookInfrastructureSpecContainerImageAttributes struct {
	ref terra.Reference
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) InternalWithRef(ref terra.Reference) NotebookInfrastructureSpecContainerImageAttributes {
	return NotebookInfrastructureSpecContainerImageAttributes{ref: ref}
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("image"))
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) JavaJars() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("java_jars"))
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("properties"))
}

func (ci NotebookInfrastructureSpecContainerImageAttributes) PythonPackages() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("python_packages"))
}

type NotebookInfrastructureSpecVpcNetworkAttributes struct {
	ref terra.Reference
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) InternalRef() (terra.Reference, error) {
	return vn.ref, nil
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) InternalWithRef(ref terra.Reference) NotebookInfrastructureSpecVpcNetworkAttributes {
	return NotebookInfrastructureSpecVpcNetworkAttributes{ref: ref}
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vn.ref.InternalTokens()
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("network"))
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) NetworkTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("network_tags"))
}

func (vn NotebookInfrastructureSpecVpcNetworkAttributes) SubNetwork() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("sub_network"))
}

type SparkAttributes struct {
	ref terra.Reference
}

func (s SparkAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SparkAttributes) InternalWithRef(ref terra.Reference) SparkAttributes {
	return SparkAttributes{ref: ref}
}

func (s SparkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SparkAttributes) ArchiveUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("archive_uris"))
}

func (s SparkAttributes) FileUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("file_uris"))
}

func (s SparkAttributes) MainClass() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("main_class"))
}

func (s SparkAttributes) MainJarFileUri() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("main_jar_file_uri"))
}

func (s SparkAttributes) PythonScriptFile() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("python_script_file"))
}

func (s SparkAttributes) SqlScript() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("sql_script"))
}

func (s SparkAttributes) SqlScriptFile() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("sql_script_file"))
}

func (s SparkAttributes) InfrastructureSpec() terra.ListValue[SparkInfrastructureSpecAttributes] {
	return terra.ReferenceAsList[SparkInfrastructureSpecAttributes](s.ref.Append("infrastructure_spec"))
}

type SparkInfrastructureSpecAttributes struct {
	ref terra.Reference
}

func (is SparkInfrastructureSpecAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is SparkInfrastructureSpecAttributes) InternalWithRef(ref terra.Reference) SparkInfrastructureSpecAttributes {
	return SparkInfrastructureSpecAttributes{ref: ref}
}

func (is SparkInfrastructureSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is SparkInfrastructureSpecAttributes) Batch() terra.ListValue[SparkInfrastructureSpecBatchAttributes] {
	return terra.ReferenceAsList[SparkInfrastructureSpecBatchAttributes](is.ref.Append("batch"))
}

func (is SparkInfrastructureSpecAttributes) ContainerImage() terra.ListValue[SparkInfrastructureSpecContainerImageAttributes] {
	return terra.ReferenceAsList[SparkInfrastructureSpecContainerImageAttributes](is.ref.Append("container_image"))
}

func (is SparkInfrastructureSpecAttributes) VpcNetwork() terra.ListValue[SparkInfrastructureSpecVpcNetworkAttributes] {
	return terra.ReferenceAsList[SparkInfrastructureSpecVpcNetworkAttributes](is.ref.Append("vpc_network"))
}

type SparkInfrastructureSpecBatchAttributes struct {
	ref terra.Reference
}

func (b SparkInfrastructureSpecBatchAttributes) InternalRef() (terra.Reference, error) {
	return b.ref, nil
}

func (b SparkInfrastructureSpecBatchAttributes) InternalWithRef(ref terra.Reference) SparkInfrastructureSpecBatchAttributes {
	return SparkInfrastructureSpecBatchAttributes{ref: ref}
}

func (b SparkInfrastructureSpecBatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return b.ref.InternalTokens()
}

func (b SparkInfrastructureSpecBatchAttributes) ExecutorsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("executors_count"))
}

func (b SparkInfrastructureSpecBatchAttributes) MaxExecutorsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(b.ref.Append("max_executors_count"))
}

type SparkInfrastructureSpecContainerImageAttributes struct {
	ref terra.Reference
}

func (ci SparkInfrastructureSpecContainerImageAttributes) InternalRef() (terra.Reference, error) {
	return ci.ref, nil
}

func (ci SparkInfrastructureSpecContainerImageAttributes) InternalWithRef(ref terra.Reference) SparkInfrastructureSpecContainerImageAttributes {
	return SparkInfrastructureSpecContainerImageAttributes{ref: ref}
}

func (ci SparkInfrastructureSpecContainerImageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ci.ref.InternalTokens()
}

func (ci SparkInfrastructureSpecContainerImageAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("image"))
}

func (ci SparkInfrastructureSpecContainerImageAttributes) JavaJars() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("java_jars"))
}

func (ci SparkInfrastructureSpecContainerImageAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ci.ref.Append("properties"))
}

func (ci SparkInfrastructureSpecContainerImageAttributes) PythonPackages() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ci.ref.Append("python_packages"))
}

type SparkInfrastructureSpecVpcNetworkAttributes struct {
	ref terra.Reference
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) InternalRef() (terra.Reference, error) {
	return vn.ref, nil
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) InternalWithRef(ref terra.Reference) SparkInfrastructureSpecVpcNetworkAttributes {
	return SparkInfrastructureSpecVpcNetworkAttributes{ref: ref}
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vn.ref.InternalTokens()
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("network"))
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) NetworkTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("network_tags"))
}

func (vn SparkInfrastructureSpecVpcNetworkAttributes) SubNetwork() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("sub_network"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type TriggerSpecAttributes struct {
	ref terra.Reference
}

func (ts TriggerSpecAttributes) InternalRef() (terra.Reference, error) {
	return ts.ref, nil
}

func (ts TriggerSpecAttributes) InternalWithRef(ref terra.Reference) TriggerSpecAttributes {
	return TriggerSpecAttributes{ref: ref}
}

func (ts TriggerSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ts.ref.InternalTokens()
}

func (ts TriggerSpecAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(ts.ref.Append("disabled"))
}

func (ts TriggerSpecAttributes) MaxRetries() terra.NumberValue {
	return terra.ReferenceAsNumber(ts.ref.Append("max_retries"))
}

func (ts TriggerSpecAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("schedule"))
}

func (ts TriggerSpecAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("start_time"))
}

func (ts TriggerSpecAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("type"))
}

type ExecutionStatusState struct {
	UpdateTime string           `json:"update_time"`
	LatestJob  []LatestJobState `json:"latest_job"`
}

type LatestJobState struct {
	EndTime    string  `json:"end_time"`
	Message    string  `json:"message"`
	Name       string  `json:"name"`
	RetryCount float64 `json:"retry_count"`
	Service    string  `json:"service"`
	ServiceJob string  `json:"service_job"`
	StartTime  string  `json:"start_time"`
	State      string  `json:"state"`
	Uid        string  `json:"uid"`
}

type ExecutionSpecState struct {
	Args                    map[string]string `json:"args"`
	KmsKey                  string            `json:"kms_key"`
	MaxJobExecutionLifetime string            `json:"max_job_execution_lifetime"`
	Project                 string            `json:"project"`
	ServiceAccount          string            `json:"service_account"`
}

type NotebookState struct {
	ArchiveUris        []string                          `json:"archive_uris"`
	FileUris           []string                          `json:"file_uris"`
	Notebook           string                            `json:"notebook"`
	InfrastructureSpec []NotebookInfrastructureSpecState `json:"infrastructure_spec"`
}

type NotebookInfrastructureSpecState struct {
	Batch          []NotebookInfrastructureSpecBatchState          `json:"batch"`
	ContainerImage []NotebookInfrastructureSpecContainerImageState `json:"container_image"`
	VpcNetwork     []NotebookInfrastructureSpecVpcNetworkState     `json:"vpc_network"`
}

type NotebookInfrastructureSpecBatchState struct {
	ExecutorsCount    float64 `json:"executors_count"`
	MaxExecutorsCount float64 `json:"max_executors_count"`
}

type NotebookInfrastructureSpecContainerImageState struct {
	Image          string            `json:"image"`
	JavaJars       []string          `json:"java_jars"`
	Properties     map[string]string `json:"properties"`
	PythonPackages []string          `json:"python_packages"`
}

type NotebookInfrastructureSpecVpcNetworkState struct {
	Network     string   `json:"network"`
	NetworkTags []string `json:"network_tags"`
	SubNetwork  string   `json:"sub_network"`
}

type SparkState struct {
	ArchiveUris        []string                       `json:"archive_uris"`
	FileUris           []string                       `json:"file_uris"`
	MainClass          string                         `json:"main_class"`
	MainJarFileUri     string                         `json:"main_jar_file_uri"`
	PythonScriptFile   string                         `json:"python_script_file"`
	SqlScript          string                         `json:"sql_script"`
	SqlScriptFile      string                         `json:"sql_script_file"`
	InfrastructureSpec []SparkInfrastructureSpecState `json:"infrastructure_spec"`
}

type SparkInfrastructureSpecState struct {
	Batch          []SparkInfrastructureSpecBatchState          `json:"batch"`
	ContainerImage []SparkInfrastructureSpecContainerImageState `json:"container_image"`
	VpcNetwork     []SparkInfrastructureSpecVpcNetworkState     `json:"vpc_network"`
}

type SparkInfrastructureSpecBatchState struct {
	ExecutorsCount    float64 `json:"executors_count"`
	MaxExecutorsCount float64 `json:"max_executors_count"`
}

type SparkInfrastructureSpecContainerImageState struct {
	Image          string            `json:"image"`
	JavaJars       []string          `json:"java_jars"`
	Properties     map[string]string `json:"properties"`
	PythonPackages []string          `json:"python_packages"`
}

type SparkInfrastructureSpecVpcNetworkState struct {
	Network     string   `json:"network"`
	NetworkTags []string `json:"network_tags"`
	SubNetwork  string   `json:"sub_network"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TriggerSpecState struct {
	Disabled   bool    `json:"disabled"`
	MaxRetries float64 `json:"max_retries"`
	Schedule   string  `json:"schedule"`
	StartTime  string  `json:"start_time"`
	Type       string  `json:"type"`
}