package main

import (
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"time"
)

const secret = "Wm44UTV0eVovRzFNSGx0YzRGL2dUa1ZKTWxyYktpWnQ="

func NewClaim(userId, email string) (jwt.Token, error) {
	return jwt.NewBuilder().
		Expiration(time.Now().Add(time.Hour*24)).
		Subject(userId).
		Claim("email", email).
		Build()
}

func SignClaimToken(claim jwt.Token) ([]byte, error) {
	signedClaim, err := jwt.Sign(claim, jwt.WithKey(jwa.HS256, secret))
	if err != nil {
		return nil, err
	}

	return signedClaim, nil
}
