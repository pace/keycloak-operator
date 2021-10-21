package monkeypatch

import "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"

type KeycloakAPIClient struct {
	// Client ID. If not specified, automatically generated.
	// +optional
	ID string `json:"id,omitempty"`
	// Client ID.
	// +kubebuilder:validation:Required
	ClientID string `json:"clientId"`
	// Client name.
	// +optional
	Name string `json:"name,omitempty"`
	// Surrogate Authentication Required option.
	// +optional
	SurrogateAuthRequired bool `json:"surrogateAuthRequired,omitempty"`
	// Client enabled flag.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// What Client authentication type to use.
	// +optional
	ClientAuthenticatorType string `json:"clientAuthenticatorType,omitempty"`
	// Client Secret. The Operator will automatically create a Secret based on this value.
	// +optional
	Secret string `json:"secret,omitempty"`
	// Application base URL.
	// +optional
	BaseURL string `json:"baseUrl,omitempty"`
	// Application Admin URL.
	// +optional
	AdminURL string `json:"adminUrl,omitempty"`
	// Application root URL.
	// +optional
	RootURL string `json:"rootUrl,omitempty"`
	// Client description.
	// +optional
	Description string `json:"description,omitempty"`
	// Default Client roles.
	// +optional
	DefaultRoles []string `json:"defaultRoles,omitempty"`
	// Default Client scopes.
	// optional
	DefaultClientScopes []string `json:"defaultClientScopes,omitempty"`
	// A list of valid Redirection URLs.
	// +optional
	RedirectUris []string `json:"redirectUris"`
	// A list of valid Web Origins.
	// +optional
	WebOrigins []string `json:"webOrigins,omitempty"`
	// Not Before setting.
	// +optional
	NotBefore int `json:"notBefore,omitempty"`
	// True if a client supports only Bearer Tokens.
	// +optional
	BearerOnly bool `json:"bearerOnly"`
	// True if Consent Screen is required.
	// +optional
	ConsentRequired bool `json:"consentRequired"`
	// True if Standard flow is enabled.
	// +optional
	StandardFlowEnabled bool `json:"standardFlowEnabled"`
	// True if Implicit flow is enabled.
	// +optional
	ImplicitFlowEnabled bool `json:"implicitFlowEnabled"`
	// True if Direct Grant is enabled.
	// +optional
	DirectAccessGrantsEnabled bool `json:"directAccessGrantsEnabled"`
	// True if Service Accounts are enabled.
	// +optional
	ServiceAccountsEnabled bool `json:"serviceAccountsEnabled"`
	// True if this is a public Client.
	// +optional
	PublicClient bool `json:"publicClient"`
	// True if this client supports Front Channel logout.
	// +optional
	FrontchannelLogout bool `json:"frontchannelLogout,omitempty"`
	// Protocol used for this Client.
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// Client Attributes.
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`
	// True if Full Scope is allowed.
	// +optional
	FullScopeAllowed bool `json:"fullScopeAllowed,omitempty"`
	// Node registration timeout.
	// +optional
	NodeReRegistrationTimeout int `json:"nodeReRegistrationTimeout,omitempty"`
	// Protocol Mappers.
	// +optional
	ProtocolMappers []v1alpha1.KeycloakProtocolMapper `json:"protocolMappers,omitempty"`
	// True to use a Template Config.
	// +optional
	UseTemplateConfig bool `json:"useTemplateConfig,omitempty"`
	// True to use Template Scope.
	// +optional
	UseTemplateScope bool `json:"useTemplateScope,omitempty"`
	// True to use Template Mappers.
	// +optional
	UseTemplateMappers bool `json:"useTemplateMappers,omitempty"`
	// Access options.
	// +optional
	Access map[string]bool `json:"access,omitempty"`
}
