package middlewares

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/reward-rabieth/Authclk/shared/enums"
	"github.com/spf13/viper"
)

type TokenRetroSpector interface {
	RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
}

func CheckToken(tokenRetroSpector TokenRetroSpector) echo.MiddlewareFunc {
	base64Str := viper.GetString("KeyCloak.RealmRS256PublicKey")
	publicKey, err := parseKeycloakRSAPublicKey(base64Str)
	if err != nil {
		panic(err)
	}

	return echojwt.WithConfig(echojwt.Config{

		SigningKey:    publicKey,
		SigningMethod: jwt.SigningMethodRS256.Name,
		SuccessHandler: func(c echo.Context) {
			//var tokenRetroSpector TokenRetroSpector

			jwtToken, ok := c.Get("user").(*jwt.Token)
			if !ok {
				log.Println("cannot cast")
			}

			claims := jwtToken.Claims.(jwt.MapClaims)

			// Get the user context from the Echo context
			ctx := c.Request().Context()
			//Create a new context with the claims value
			contextWithClaims := context.WithValue(ctx, enums.ContextKeyClaims, claims)
			//Set the updated user context in the Echo context
			c.Set("context", contextWithClaims)

			rpResult, err := tokenRetroSpector.RetrospectToken(ctx, jwtToken.Raw)

			if err != nil {
				log.Println(err)
			}
			if !*rpResult.Active {
				log.Println(err)
			}

		},
	},
	)

}

func parseKeycloakRSAPublicKey(base64Str string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, nil
	}

	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
