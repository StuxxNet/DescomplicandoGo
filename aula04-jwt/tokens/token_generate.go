package tokens

import (
	"fmt"

	jwt "github.com/cristalhq/jwt/v4"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Generate() {
	// create a Signer (HMAC in this example)
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	checkErr(err)

	// create claims (you can create your own, see: Example_BuildUserClaims)
	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       "random-unique-string",
	}

	// create a Builder
	builder := jwt.NewBuilder(signer)

	// and build a Token
	token, err := builder.Build(claims)
	checkErr(err)

	// here is token as a string
	var jwtToken string = token.String()
	fmt.Println(jwtToken)
}
