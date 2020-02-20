package model

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)


func GetStartupScript(cr *kc.Keycloak) map[string]string {
	return  map[string]string{
		"mystartup.sh": cr.Spec.StartupScript.Content,
	}
}


func KeycloakConfigMapStartup(cr *kc.Keycloak) *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta {
			Name:        ApplicationName + "-startup",
			Namespace:   cr.Namespace,
		},
		Data: GetStartupScript(cr),
	}
}


func KeycloakConfigMapStartupReconiled(cr *kc.Keycloak, currentState *v1.ConfigMap) *v1.ConfigMap {
	reconciled := currentState.DeepCopy()
	reconciled.Data = GetStartupScript(cr)
	return reconciled
}

func KeycloakConfigMapStartupSelector(cr *kc.Keycloak) client.ObjectKey {
	return client.ObjectKey{
		Name:      ApplicationName + "-startup",
		Namespace: cr.Namespace,
	}
}
