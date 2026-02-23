package common

import (
	"context"
	"fmt"
	"os"

	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// These kinds are not provided by the openshift api
const (
	RouteKind                 = "Route"
	JobKind                   = "Job"
	CronJobKind               = "CronJob"
	SecretKind                = "Secret"
	StatefulSetKind           = "StatefulSet"
	ServiceKind               = "Service"
	IngressKind               = "Ingress"
	DeploymentKind            = "Deployment"
	PersistentVolumeClaimKind = "PersistentVolumeClaim"
	PodDisruptionBudgetKind   = "PodDisruptionBudget"
)

func WatchSecondaryResource(c controller.Controller, controllerName string, resourceKind string, objectTypetoWatch runtime.Object, cr runtime.Object) error {
	stateManager := GetStateManager()
	stateFieldName := GetStateFieldName(controllerName, resourceKind)

	// Check to see if the watch exists for this resource type already for this controller, if it does, we return so we don't set up another watch
	watchExists, keyExists := stateManager.GetState(stateFieldName).(bool)
	if keyExists || watchExists {
		return nil
	}

	// Set up the actual watch
	err := c.Watch(&source.Kind{Type: objectTypetoWatch}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    cr,
	})

	// Retry on error
	if err != nil {
		log.Error(err, "error creating watch")
		stateManager.SetState(stateFieldName, false)
		return err
	}

	stateManager.SetState(stateFieldName, true)
	log.Info(fmt.Sprintf("Watch created for '%s' resource in '%s'", resourceKind, controllerName))
	return nil
}

func GetStateFieldName(controllerName string, kind string) string {
	return controllerName + "-watch-" + kind
}

// GetMatchingKeycloaks returns a synthetic Keycloak instance based on env vars
// KEYCLOAK_INTERNAL_URL and KEYCLOAK_CREDENTIAL_SECRET. The Keycloak CRD lives
// under k8s.keycloak.org/v2alpha1 (managed by the Quarkus operator), so we cannot
// query it from this legacy operator. Instead we read the connection info from the
// ExternalKeycloak CR via env vars injected into the deployment.
func GetMatchingKeycloaks(ctx context.Context, c client.Client, labelSelector *v1.LabelSelector) (v1alpha1.KeycloakList, error) {
	internalURL := os.Getenv("KEYCLOAK_INTERNAL_URL")
	credentialSecret := os.Getenv("KEYCLOAK_CREDENTIAL_SECRET")
	namespace := os.Getenv("KEYCLOAK_NAMESPACE")

	if internalURL == "" || credentialSecret == "" {
		return v1alpha1.KeycloakList{}, fmt.Errorf("KEYCLOAK_INTERNAL_URL and KEYCLOAK_CREDENTIAL_SECRET must be set")
	}
	if namespace == "" {
		namespace = "keycloak"
	}

	kc := v1alpha1.Keycloak{}
	kc.Namespace = namespace
	kc.Name = "external-keycloak"
	kc.Status.InternalURL = internalURL
	kc.Status.CredentialSecret = credentialSecret
	kc.Status.Ready = true

	return v1alpha1.KeycloakList{Items: []v1alpha1.Keycloak{kc}}, nil
}

// Try to get a list of keycloak instances that match the selector specified on the realm
func GetMatchingRealms(ctx context.Context, c client.Client, labelSelector *v1.LabelSelector) (v1alpha1.KeycloakRealmList, error) {
	var list v1alpha1.KeycloakRealmList
	opts := []client.ListOption{
		client.MatchingLabels(labelSelector.MatchLabels),
	}

	err := c.List(ctx, &list, opts...)
	if err != nil {
		return list, err
	}

	return list, nil
}
