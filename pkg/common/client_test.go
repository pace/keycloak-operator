package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
)

const (
	RealmsGetPath    = "/auth/admin/realms/%s"
	RealmsCreatePath = "/auth/admin/realms"
	RealmsDeletePath = "/auth/admin/realms/%s"
	UserCreatePath   = "/auth/admin/realms/%s/users"
	UserDeletePath   = "/auth/admin/realms/%s/users/%s"
	TokenPath        = "/auth/realms/master/protocol/openid-connect/token" // nolint
)

func getDummyRealm() *v1alpha1.KeycloakRealm {
	return &v1alpha1.KeycloakRealm{
		Spec: v1alpha1.KeycloakRealmSpec{
			Realm: &v1alpha1.KeycloakAPIRealm{
				ID:          "dummy",
				Realm:       "dummy",
				Enabled:     false,
				DisplayName: "dummy",
			},
		},
	}
}

func getDummyUser() *v1alpha1.KeycloakAPIUser {
	return &v1alpha1.KeycloakAPIUser{
		ID:            "dummy",
		UserName:      "dummy",
		FirstName:     "dummy",
		LastName:      "dummy",
		EmailVerified: false,
		Enabled:       false,
	}
}

func TestClient_CreateRealm(t *testing.T) {
	// given
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, RealmsCreatePath, req.URL.Path)
		w.WriteHeader(201)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	realm := getDummyRealm()

	// when
	err := client.CreateRealm(realm)

	// then
	// no error expected
	// correct path expected on httptest server
	assert.NoError(t, err)
}

func TestClient_DeleteRealmRealm(t *testing.T) {
	// given
	realm := getDummyRealm()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, fmt.Sprintf(RealmsDeletePath, realm.Spec.Realm.Realm), req.URL.Path)
		w.WriteHeader(204)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	err := client.DeleteRealm(realm.Spec.Realm.Realm)

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
}

func TestClient_CreateUser(t *testing.T) {
	// given
	user := getDummyUser()
	realm := getDummyRealm()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, fmt.Sprintf(UserCreatePath, realm.Spec.Realm.Realm), req.URL.Path)
		w.WriteHeader(201)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	err := client.CreateUser(user, realm.Spec.Realm.Realm)

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
}

func TestClient_DeleteUser(t *testing.T) {
	// given
	user := getDummyUser()
	realm := getDummyRealm()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, fmt.Sprintf(UserDeletePath, realm.Spec.Realm.Realm, user.ID), req.URL.Path)
		w.WriteHeader(204)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	err := client.DeleteUser(user.ID, realm.Spec.Realm.Realm)

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
}

func TestClient_GetRealm(t *testing.T) {
	// given
	realm := getDummyRealm()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, fmt.Sprintf(RealmsGetPath, realm.Spec.Realm.Realm), req.URL.Path)
		assert.Equal(t, req.Method, http.MethodGet)
		json, err := jsoniter.Marshal(realm.Spec.Realm)
		assert.NoError(t, err)

		size, err := w.Write(json)
		assert.NoError(t, err)
		assert.Equal(t, size, len(json))

		w.WriteHeader(204)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	newRealm, err := client.GetRealm(realm.Spec.Realm.Realm)

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)

	// returned realm must equal dummy realm
	assert.Equal(t, realm.Spec.Realm.Realm, newRealm.Spec.Realm.Realm)
}

func TestClient_ListRealms(t *testing.T) {
	// given
	realm := getDummyRealm()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, RealmsCreatePath, req.URL.Path)
		assert.Equal(t, req.Method, http.MethodGet)
		var list []*v1alpha1.KeycloakRealm
		list = append(list, realm)
		json, err := jsoniter.Marshal(list)
		assert.NoError(t, err)

		size, err := w.Write(json)
		assert.NoError(t, err)
		assert.Equal(t, size, len(json))

		w.WriteHeader(204)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	realms, err := client.ListRealms()

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)

	// exactly one realms must be returned
	assert.Len(t, realms, 1)
}

func TestClient_login(t *testing.T) {
	// given
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, TokenPath, req.URL.Path)
		assert.Equal(t, req.Method, http.MethodPost)

		response := v1alpha1.TokenResponse{
			AccessToken: "dummy",
		}

		json, err := jsoniter.Marshal(response)
		assert.NoError(t, err)

		size, err := w.Write(json)
		assert.NoError(t, err)
		assert.Equal(t, size, len(json))

		w.WriteHeader(204)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "not set",
	}

	// when
	err := client.login("dummy", "dummy")

	// then
	// token must be set on the client now
	assert.NoError(t, err)
	assert.Equal(t, client.token, "dummy")
}

func listAllClientScopesHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		require.Equal(t, req.Method, http.MethodGet)
		require.True(t, strings.HasSuffix(req.URL.Path, "realms/testrealm/client-scopes"), req.URL.Path)
		list := []*ClientScope{
			&ClientScope{
				ID:   "12345-6789",
				Name: "scope:a:b:c",
			},
			&ClientScope{
				ID:   "6789-12345",
				Name: "scope:d:e:f",
			},
			&ClientScope{
				ID:   "12334-6789-12345",
				Name: "scope:k:t:c",
			},
			&ClientScope{
				ID:   "888888-6789-12345",
				Name: "scope:u:r:k:t:c",
			},
		}
		json, err := jsoniter.Marshal(list)
		assert.NoError(t, err)

		size, err := w.Write(json)
		assert.NoError(t, err)
		assert.Equal(t, size, len(json))

		w.WriteHeader(200)
	})
}

func TestClient_ListAllClientScopes(t *testing.T) {
	handler := listAllClientScopesHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	scopes, err := client.ListAllClientScopes("testrealm")

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
	assert.Len(t, scopes, 4)
	assert.Equal(t, "12345-6789", scopes[0].ID)
	assert.Equal(t, "scope:a:b:c", scopes[0].Name)
}

func listClientScopesForClientHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		require.Equal(t, req.Method, http.MethodGet)
		assert.True(t, strings.HasSuffix(req.URL.Path, "realms/testrealm/clients/testclient/default-client-scopes"), req.URL.Path)
		list := []*ClientScope{
			&ClientScope{
				ID:   "12345-6789",
				Name: "scope:a:b:c",
			},
			&ClientScope{
				ID:   "6789-12345",
				Name: "scope:d:e:f",
			},
			&ClientScope{
				ID:   "888888-6789-12345",
				Name: "scope:u:r:k:t:c",
			},
		}
		json, err := jsoniter.Marshal(list)
		assert.NoError(t, err)

		size, err := w.Write(json)
		assert.NoError(t, err)
		assert.Equal(t, size, len(json))

		w.WriteHeader(200)
	})
}

func TestClient_ListClientScopesForClient(t *testing.T) {
	handler := listClientScopesForClientHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	scopes, err := client.ListClientScopesForClient("testclient", "testrealm")

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
	assert.Len(t, scopes, 3)
	assert.Equal(t, "12345-6789", scopes[0].ID)
	assert.Equal(t, "scope:a:b:c", scopes[0].Name)
}

func TestClient_UpdateClientScopes(t *testing.T) {
	removalCount := 0
	additionCount := 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ListAllClientScopes
		if strings.HasSuffix(req.URL.Path, "realms/testrealm/client-scopes") {
			listAllClientScopesHandler(t)(w, req)
			return
		}

		// ListClientScopesForClient
		if strings.HasSuffix(req.URL.Path, "realms/testrealm/clients/testclient/default-client-scopes") &&
			req.Method == http.MethodGet {
			listClientScopesForClientHandler(t)(w, req)
			return
		}

		if strings.Contains(req.URL.Path, "/auth/admin/realms/testrealm/clients/testclient/default-client-scopes") &&
			req.Method != http.MethodGet {
			if req.Method == http.MethodPut {
				additionCount++
			}
			if req.Method == http.MethodDelete {
				removalCount++
			}
			w.WriteHeader(204)
			return
		}
		w.WriteHeader(500)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	client := Client{
		requester: server.Client(),
		URL:       server.URL,
		token:     "dummy",
	}

	// when
	spec := &v1alpha1.KeycloakAPIClient{
		ID:       "testclient",
		ClientID: "testclientid",
		DefaultClientScopes: []string{
			"scope:a:b:c", // have, present in the response of listClientScopesForClientHandler
			"scope:d:e:f", // have, ^^^^^^^^^^^^
			"scope:k:t:c", // want, missing in the response of listClientScopesForClientHandler
			// scope:u:r:k:t:c is being returned by listClientScopesForClientHandler but is not
			// present in the resource and should therefore be removed.
		},
	}
	err := client.UpdateClientScopes(spec, "testrealm")

	// then
	// correct path expected on httptest server
	assert.NoError(t, err)
	assert.Equal(t, 1, additionCount)
	assert.Equal(t, 1, removalCount)
}
