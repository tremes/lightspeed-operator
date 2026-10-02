package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// AlertsAdapterConfigSpec specifies runtime configuration for the alerts adapter.
// Creating an AlertsAdapterConfig named cluster activates the adapter.
// An empty spec uses the adapter's runtime defaults.
type AlertsAdapterConfigSpec struct {
	// pollPeriodSeconds is the time between alert polls.
	// When omitted, the current default is 30 seconds. This default may change.
	// +optional
	// +kubebuilder:validation:Minimum=1
	PollPeriodSeconds int32 `json:"pollPeriodSeconds,omitempty"`

	// preRunDelaySeconds is the delay before starting an AgenticRun.
	// When omitted, the current default is zero. This default may change.
	// +optional
	// +kubebuilder:validation:Minimum=0
	PreRunDelaySeconds int32 `json:"preRunDelaySeconds,omitempty"`

	// postRunDelaySeconds is the delay after an AgenticRun.
	// When omitted, the current default is 3600 seconds. This default may change.
	// Explicit zero disables the delay.
	// +optional
	// +kubebuilder:validation:Minimum=0
	PostRunDelaySeconds *int32 `json:"postRunDelaySeconds,omitempty"`

	// filtering configures which alert receivers the adapter accepts.
	// When omitted, the adapter uses its default filtering behavior.
	// +optional
	Filtering *AlertsAdapterFilteringConfig `json:"filtering,omitempty"`

	// deduplication configures labels ignored when deduplicating alerts.
	// When omitted, the adapter uses its default ignored labels.
	// +optional
	Deduplication *AlertsAdapterDeduplicationConfig `json:"deduplication,omitempty"`

	// tools configures skills supplied to AgenticRuns.
	// When omitted, no skills are configured.
	// +optional
	Tools *AlertsAdapterToolsConfig `json:"tools,omitempty"`

	// agent configures agent name overrides for workflow steps.
	// When omitted, the adapter uses its default agent names.
	// +optional
	Agent *AlertsAdapterAgentConfig `json:"agent,omitempty"`
}

// AlertsAdapterFilteringConfig configures alert receiver filtering.
// +kubebuilder:validation:MinProperties=1
type AlertsAdapterFilteringConfig struct {
	// allowedReceivers lists accepted receiver names. The adapter lowercases
	// these names. When omitted or empty, no receiver list is configured.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +listType=atomic
	AllowedReceivers []string `json:"allowedReceivers,omitempty"`
}

// AlertsAdapterDeduplicationConfig configures alert deduplication.
// +kubebuilder:validation:MinProperties=1
type AlertsAdapterDeduplicationConfig struct {
	// ignoredLabels lists labels excluded from deduplication. When omitted,
	// the current default is pod, instance, endpoint, and uid. This default
	// may change. An explicit empty list excludes no labels.
	// +optional
	// +kubebuilder:validation:MaxItems=100
	// +listType=atomic
	IgnoredLabels *[]string `json:"ignoredLabels,omitempty"`
}

// AlertsAdapterToolsConfig configures run-level skills.
// +kubebuilder:validation:MinProperties=1
type AlertsAdapterToolsConfig struct {
	// skills lists skill sources. When omitted or empty, no skills are configured.
	// +optional
	// +kubebuilder:validation:MaxItems=32
	// +listType=atomic
	Skills []AlertsAdapterSkillSource `json:"skills,omitempty"`
}

// AlertsAdapterSkillSource describes one skill source.
type AlertsAdapterSkillSource struct {
	// image is the skill image. It may be omitted or empty.
	// With the current adapter, an entry without an image is skipped.
	// +optional
	// +kubebuilder:validation:MaxLength=2048
	Image string `json:"image,omitempty"`

	// paths lists skill paths in the image. It may be omitted or empty.
	// With the current adapter, an entry without paths is skipped.
	// +optional
	// +kubebuilder:validation:MaxItems=64
	// +listType=atomic
	Paths []string `json:"paths,omitempty"`
}

// AlertsAdapterAgentConfig configures agent name overrides.
// +kubebuilder:validation:MinProperties=1
type AlertsAdapterAgentConfig struct {
	// defaultAgentName overrides the default agent name when non-empty.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	DefaultAgentName string `json:"defaultAgentName,omitempty"`

	// analysisAgentName overrides the analysis agent name when non-empty.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	AnalysisAgentName string `json:"analysisAgentName,omitempty"`

	// executionAgentName overrides the execution agent name when non-empty.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	ExecutionAgentName string `json:"executionAgentName,omitempty"`

	// verificationAgentName overrides the verification agent name when non-empty.
	// +optional
	// +kubebuilder:validation:MaxLength=253
	VerificationAgentName string `json:"verificationAgentName,omitempty"`
}

// AlertsAdapterConfigStatus describes the observed state of the adapter.
type AlertsAdapterConfigStatus struct {
	// conditions reports reconciliation and availability. The Ready condition
	// uses observedGeneration to identify the spec it reflects.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// AlertsAdapterConfig configures and activates the cluster's alerts adapter.
// Only the object named cluster is valid.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'cluster'",message=".metadata.name must be 'cluster'"
type AlertsAdapterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec is the desired runtime configuration. It may be empty.
	// +required
	Spec AlertsAdapterConfigSpec `json:"spec"`

	// status is the observed state, maintained by the controller.
	// +optional
	Status AlertsAdapterConfigStatus `json:"status,omitempty"`
}

// AlertsAdapterConfigList contains AlertsAdapterConfig objects.
// +kubebuilder:object:root=true
type AlertsAdapterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertsAdapterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlertsAdapterConfig{}, &AlertsAdapterConfigList{})
}
