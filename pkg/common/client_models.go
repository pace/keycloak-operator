package common

type ClientScope struct {
	ID          string            `json:"id"`
	Name        string            `json:"name,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	Description string            `json:"description,omitempty"`
	Protocol    string            `json:"protocol,omitempty"`
	// TODO (JF) Add ProtocolMappers, https://www.keycloak.org/docs-api/8.0/rest-api/index.html#_clientscoperepresentation
}

type UpdateClientScopeRequest struct {
	Realm         string `json:"realm"`
	ClientID      string `json:"client"` // not to be confused with the oauth clientId
	ClientScopeID string `json:"clientScopeId"`
}
