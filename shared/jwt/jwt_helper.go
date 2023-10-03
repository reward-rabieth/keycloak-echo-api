package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/goutil"
	"strings"
)

type JwtHelper struct {
	claims       jwt.MapClaims
	realmRoles   []string
	accountRoles []string
	scopes       []string
}

func NewJwtHelper(claims jwt.MapClaims) *JwtHelper {
	return &JwtHelper{
		claims:       claims,
		realmRoles:   parseRealmRoles(claims),
		accountRoles: parseAccountRoles(claims),
		scopes:       parseScope(claims),
	}
}

func parseRealmRoles(claims jwt.MapClaims) []string {
	var realmRoles = make([]string, 0)
	if claim, ok := claims["realm_access"]; ok {
		if roles, ok := claim.(map[string]interface{})["roles"]; ok {
			for _, role := range roles.([]interface{}) {
				realmRoles = append(realmRoles, role.(string))
			}
		}

	}
	return realmRoles
}

func parseAccountRoles(claims jwt.MapClaims) []string {
	var accountRoles = make([]string, 0)
	if acc, ok := claims["resource_access"]; ok {
		if roles, ok := acc.(map[string]interface{})["roles"]; ok {
			for _, role := range roles.([]interface{}) {
				accountRoles = append(accountRoles, role.(string))
			}

		}
	}
	return accountRoles
}

func parseScope(claims jwt.MapClaims) []string {
	scopeStr, err := parseString(claims, "scope")
	if err != nil {
		return make([]string, 0)
	}
	scopes := strings.Split(scopeStr, "")

	return scopes
}

func (j *JwtHelper) TokenHasScope(scope string) bool {
	return goutil.Contains(j.scopes, scope)
}

func (j *JwtHelper) GetUserId() (string, error) {
	return j.claims.GetSubject()
}
func (j *JwtHelper) IsUserInRealmRole(role string) bool {
	return goutil.Contains(j.realmRoles, role)
}

func parseString(claims jwt.MapClaims, key string) (string, error) {
	var (
		ok  bool
		raw interface{}
		iss string
	)

	raw, ok = claims[key]
	if !ok {
		return "", nil
	}

	iss, ok = raw.(string)
	if !ok {
		return "", fmt.Errorf("key %s is invalid ", key)
	}

	return iss, nil
}
