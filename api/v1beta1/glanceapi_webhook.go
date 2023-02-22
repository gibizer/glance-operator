/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// we can make this its own separate struct type if necessary later on
var glanceAPIDefaults GlanceDefaults

// log is for logging in this package.
var glanceapilog = logf.Log.WithName("glanceapi-resource")

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *GlanceAPI) SetupWebhookWithManager(mgr ctrl.Manager, defaults GlanceDefaults) error {
	glanceAPIDefaults = defaults

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-glance-openstack-org-v1beta1-glanceapi,mutating=true,failurePolicy=fail,sideEffects=None,groups=glance.openstack.org,resources=glanceapis,verbs=create;update,versions=v1beta1,name=mglanceapi.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &GlanceAPI{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *GlanceAPI) Default() {
	glanceapilog.Info("default", "name", r.Name)

	if r.Spec.ContainerImage == "" {
		r.Spec.ContainerImage = glanceAPIDefaults.ContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-glance-openstack-org-v1beta1-glanceapi,mutating=false,failurePolicy=fail,sideEffects=None,groups=glance.openstack.org,resources=glanceapis,verbs=create;update,versions=v1beta1,name=vglanceapi.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &GlanceAPI{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *GlanceAPI) ValidateCreate() error {
	glanceapilog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *GlanceAPI) ValidateUpdate(old runtime.Object) error {
	glanceapilog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *GlanceAPI) ValidateDelete() error {
	glanceapilog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
