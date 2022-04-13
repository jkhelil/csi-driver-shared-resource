package csidriver

import (
	sharev1alpha1 "github.com/openshift/api/sharedresource/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime"
	admissionctl "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	admissionv1 "k8s.io/api/admission/v1"
)

const (
	WebhookName        string = "sharedresourcecsidriver-csidriver"
)

// Webhook interface
type Webhook interface {
	// Authorized will determine if the request is allowed
	Authorized(request admissionctl.Request) admissionctl.Response
	// GetURI returns the URI for the webhook
	GetURI() string
	// Validate will validate the incoming request
	Validate(admissionctl.Request) bool
	// Name is the name of the webhook
	Name() string
}

// SharedResourcesCSIDriverWebhook validates a Shared Resources CSI Driver change
type SharedResourcesCSIDriverWebhook struct{
	s runtime.Scheme
}

// Authorized implements Webhook interface
func (s *SharedResourcesCSIDriverWebhook) Authorized(request admissionctl.Request) admissionctl.Response {
	var ret admissionctl.Response
	ret = admissionctl.Allowed("Allowed to create SharedResource")
	return ret
}

// Validate if the incoming request even valid
func (s *SharedResourcesCSIDriverWebhook) Validate(req admissionctl.Request) bool {
	return true
}

// GetURI implements Webhook interface
func (s *SharedResourcesCSIDriverWebhook) GetURI() string { return "/" + WebhookName }

// Name implements Webhook interface
func (s *SharedResourcesCSIDriverWebhook) Name() string { return WebhookName }

// NewWebhook creates a new webhook
func NewWebhook() *SharedResourcesCSIDriverWebhook {
	scheme := runtime.NewScheme()
	admissionv1.AddToScheme(scheme)
	sharev1alpha1.AddToScheme(scheme)

	return &SharedResourcesCSIDriverWebhook{
		s: *scheme,
	}
}