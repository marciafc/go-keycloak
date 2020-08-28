package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = "app"
	clientSecret = "14ec4396-d000-4096-a5aa-3c1dbae54092"
)

func main() {

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{

		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	// TODO Configurar de uma variável de ambiente
	state := "magica"

	// Redireciona para o Keycloak autenticar
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	// Retorno do Keycloak
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {

		// Valida se quando recebe do Keycloak, a String é a mesma
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)

			// Não passou pela autenticação
			return
		}

		// Troca o code pelo access token
		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "failed to excahge token", http.StatusBadRequest)
			return
		}

		// id_token é o token de quem está autenticado
		// oauth2Token (acima) é o token para AUTORIZAÇÃO
		// rawIDToken é token para AUTENTICAÇÃO
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "no id_token", http.StatusBadRequest)
			return
		}

		// Estrutura para exibir os dados
		resp := struct {
			OAuth2Token *oauth2.Token
			RawIDToken  string
		}{
			oauth2Token, rawIDToken,
		}

		// Converte o struct em json
		data, err := json.MarshalIndent(resp, "", "   ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Exibe o json
		w.Write(data)

	})

	// Subindo servidor na 8081
	log.Fatal(http.ListenAndServe(":8081", nil))

}
