# Autenticação com OpenID Connect e Keycloak

  - [Golang](https://golang.org/)

  - [Docker](https://www.docker.com/)

  - [Keycloak](https://www.keycloak.org/)

  - [OpenID Connect](https://openid.net/connect/)


## Execução

1. Start container com o Keycloak

    $ docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:11.0.1

2. Login no [Administration Console](http://localhost:8080)

    user=admin, password=admin

3. Criar um realm
    
    - Clique "Add realm" no menu suspenso onde diz "Master"

    - Name: demo

    - Clique "Create"

    Verificar se está no realm recém criado

4. Criar client

    - Clique "Clients" \ "Create"

    - Preencher os campos

      - Client ID: app
      - Client Protocol: openid-connect
      - Root URL: http://localhost:8081

    - Clique "Save"  

    - Aba "Settings"

      - Access Type: confidential
      - Clique "Save"


5. Criar usuário

    - Clique "Users" \ "Add user"

    - Preencher os campos

      - Username: marcia
      - Email: marciafc.info@gmail.com
      - First Name: Marcia
      - Last Name: Castagna
      - User Enabled (usuário ativo): ON
      - Email Verified (usuário já verificado): ON      

    - Clique "Save"      

    - Clique "Credentials"
      - Preencher "Password"
      - Temporary (se o password é temporário): OFF
      - Clique "Set password"

6. Configurar os campos no arquivo main.go

    clientID: app

    clientSecret: "Secret" no Keycloak
    
    http://localhost:8080  \ Clients \ app \ Credentials \ Secret

7. Executar os comandos para subir a aplicação na porta 8081

    $ go mod init keycloak-exemplo

    $ go run client/main.go

8. Acessar o navegador http://localhost:8081/    