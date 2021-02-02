package sjwt_test

import (
	"testing"

	sjwt "github.com/alseiitov/simple-jwt"
)

func TestCorrectlyGenerated(t *testing.T) {
	secret := "53cr3tk3y"

	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	payload := make(map[string]interface{})
	payload["user"] = "alseiitov"

	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWxzZWlpdG92In0.zkpqkk44FlOLmdC0Lvz-6PZVLX4pvx43_bCIyHrNM3M"

	jwt := sjwt.New(header, payload)
	token, err := jwt.Sign(secret)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if token != expected {
		t.Logf("\nExpected:	%s\nBut got:	%s\n", expected, token)
		t.Fail()
	}
}

func TestPassedVerification(t *testing.T) {
	secret := "53cr3tk3y"

	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	payload := make(map[string]interface{})
	payload["user"] = "alseiitov"

	jwt := sjwt.New(header, payload)
	token, err := jwt.Sign(secret)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	jwt2, _ := sjwt.Parse(token)
	if err := jwt2.Verify(token, secret); err != nil {
		t.Error(err)
		t.Fail()
	}
}
