package identity

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type identityManager struct {
	baseUrl             string
	realm               string
	restAPIClientID     string
	restAPIClientSecret string
}

func NewIdentityManger() *identityManager {
	return &identityManager{
		baseUrl:             viper.GetString("KeyCloak.BaseUrl"),
		realm:               viper.GetString("KeyCloak.Realm"),
		restAPIClientID:     viper.GetString("KeyCloak.RestAPI.ClientID"),
		restAPIClientSecret: viper.GetString("KeyCloak.RestAPI.ClientSecret"),
	}
}

func (im *identityManager) loginRestAPIClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.baseUrl)
	fmt.Println(im.restAPIClientID)
	token, err := client.LoginClient(ctx, im.restAPIClientID, im.restAPIClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login the rest client")
	}

	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error) {
	token, err := im.loginRestAPIClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)
	userID, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create the user")
	}
	err = client.SetPassword(ctx, token.AccessToken, userID, im.realm, password, false)
	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password for the user")
	}
	var roleNameLowerCase = strings.ToLower(role)
	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, roleNameLowerCase)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v'", roleNameLowerCase))
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userID, []gocloak.Role{*roleKeycloak})
	if err != nil {
		return nil, errors.Wrap(err, "unable to add realm role to the user")
	}
	userCloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return userCloak, nil
}
func (im *identityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	client := gocloak.NewClient(im.baseUrl)
	fmt.Println(client)
	rptResult, err := client.RetrospectToken(ctx, accessToken, im.restAPIClientID, im.restAPIClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect the token")
	}

	return rptResult, nil
}
