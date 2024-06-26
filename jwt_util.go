package main

import (
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
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

	v, err := jwk.FromRaw([]byte(secret))
	signedClaim, err := jwt.Sign(claim, jwt.WithKey(jwa.HS256, v))
	if err != nil {
		return nil, err
	}

	return signedClaim, nil
}

func ParseToken(token string) (jwt.Token, error) {
	v, err := jwk.FromRaw([]byte(secret))

	parse, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.HS256, v))
	if err != nil {
		return nil, err
	}

	return parse, err
}
