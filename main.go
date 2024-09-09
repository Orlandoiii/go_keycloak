package go_keycloak

import (
	"fmt"

	"github.com/Orlandoiii/go_keycloak/custom_jwt"
)

func TestGoCLoak() {
	fmt.Println("Este es mi gocloak")

	var test custom_jwt.RealmAccess = custom_jwt.RealmAccess{
		Roles: []string{"admin", "user"},
	}

	fmt.Println(test)
}
