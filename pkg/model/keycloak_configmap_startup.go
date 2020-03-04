package model

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)


func GetStartupScriptConfigMapContent(cr *kc.Keycloak) map[string]string {

	startupContent := map[string]string{}

	if cr.Spec.StartupScript.Enabled {
		startupContent["keycloak.cli"] =  cr.Spec.KeycloakCli.Content
	}

	if cr.Spec.KeycloakCli.Enabled {
		startupContent["keycloak.cli"] = GetKeycloakCliDefaultContent(cr.Spec.KeycloakCli.Content)
	}

	return startupContent
}

func GetKeycloakCliDefaultContent(customContent string) string {

	keycloakCliDefinitionTemplate := `
	embed-server --std-out=echo
	batch

	## Sets the node identifier to the node name (= pod name). Node identifiers have to be unique. They can have a
	## maximum length of 23 characters. Thus, the chart's fullname template truncates its length accordingly.
	/subsystem=transactions:write-attribute(name=node-identifier, value=${jboss.node.name})

	# Allow log level to be configured via environment variable
	/subsystem=logging/console-handler=CONSOLE:write-attribute(name=level, value=${env.WILDFLY_LOGLEVEL:INFO})
	/subsystem=logging/root-logger=ROOT:write-attribute(name=level, value=${env.WILDFLY_LOGLEVEL:INFO})
	
	# Add dedicated eventsListener config element to allow configuring elements.
	/subsystem=keycloak-server/spi=eventsListener:add()
	/subsystem=keycloak-server/spi=eventsListener/provider=jboss-logging:add(enabled=true)
	
	# Propagate success events to INFO instead of DEBUG, to expose successful logins for log analysis
	/subsystem=keycloak-server/spi=eventsListener/provider=jboss-logging:write-attribute(name=properties.success-level,value=info)
	/subsystem=keycloak-server/spi=eventsListener/provider=jboss-logging:write-attribute(name=properties.error-level,value=warn)

	# custom content
	{keycloakCliCustomContent}
	`
	FinalKeycloakCliDefinitionTemplate := strings.Replace(keycloakCliDefinitionTemplate, "{keycloakCliCustomContent}", customContent, -1)


	return FinalKeycloakCliDefinitionTemplate
}

func KeycloakConfigMapStartup(cr *kc.Keycloak) *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta {
			Name:        ApplicationName + "-startup",
			Namespace:   cr.Namespace,
		},
		Data: GetStartupScriptConfigMapContent(cr),
	}
}


func KeycloakConfigMapStartupReconiled(cr *kc.Keycloak, currentState *v1.ConfigMap) *v1.ConfigMap {
	reconciled := currentState.DeepCopy()
	reconciled.Data = GetStartupScriptConfigMapContent(cr)
	return reconciled
}

func KeycloakConfigMapStartupSelector(cr *kc.Keycloak) client.ObjectKey {
	return client.ObjectKey{
		Name:      ApplicationName + "-startup",
		Namespace: cr.Namespace,
	}
}
