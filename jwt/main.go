package main

import "github.com/dgrijalva/jwt-go"
import "github.com/mendsley/gojwk"

var (
	//使用在API网关设置的keyId
	keyId = "uniq_key"
	//使用本文3.2节生成的Keypare
	privateKeyJson = "{\n" +
		"  \"kty\": \"RSA\",\n" +
		"  \"d\": " +
		"\"O9MJSOgcjjiVMNJ4jmBAh0mRHF_TlaVva70Imghtlgwxl8BLfcf1S8ueN1PD7xV6Cnq8YenSKsfiNOhC6yZ_fjW1syn5raWfj68eR7cjHWjLOvKjwVY33GBPNOvspNhVAFzeqfWneRTBbga53Agb6jjN0SUcZdJgnelzz5JNdOGaLzhacjH6YPJKpbuzCQYPkWtoZHDqWTzCSb4mJ3n0NRTsWy7Pm8LwG_Fd3pACl7JIY38IanPQDLoighFfo-Lriv5z3IdlhwbPnx0tk9sBwQBTRdZ8JkqqYkxUiB06phwr7mAnKEpQJ6HvhZBQ1cCnYZ_nIlrX9-I7qomrlE1UoQ\",\n" +
		"  \"e\": \"AQAB\",\n" +
		"  \"kid\": \"myJwtKey\",\n" +
		"  \"alg\": \"RS256\",\n" +
		"  \"n\": \"vCuB8MgwPZfziMSytEbBoOEwxsG7XI3MaVMoocziP4SjzU4IuWuE_DodbOHQwb_thUru57_Efe" +
		"--sfATHEa0Odv5ny3QbByqsvjyeHk6ZE4mSAV9BsHYa6GWAgEZtnDceeeDc0y76utXK2XHhC1Pysi2KG8KAzqDa099Yh7s31AyoueoMnrYTmWfEyDsQL_OAIiwgXakkS5U8QyXmWicCwXntDzkIMh8MjfPskesyli0XQD1AmCXVV3h2Opm1Amx0ggSOOiINUR5YRD6mKo49_cN-nrJWjtwSouqDdxHYP-4c7epuTcdS6kQHiQERBd1ejdpAxV4c0t0FHF7MOy9kw\"\n" +
		"}"
)

type Customer struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

func main() {
	println(privateKeyJson)
	// Audience  string `json:"aud,omitempty"`
	// ExpiresAt int64  `json:"exp,omitempty"`
	// Id        string `json:"jti,omitempty"`
	// IssuedAt  int64  `json:"iat,omitempty"`
	// Issuer    string `json:"iss,omitempty"`
	// NotBefore int64  `json:"nbf,omitempty"`
	// Subject   string `json:"sub,omitempty"`
	// {"jti":"s1m7pfZHUzrncg4LGDhwoA","iat":1562320980,"exp":1562328180,"nbf":1562320920,"sub":"","aud":"YOUR_AUDIENCE","userId":"1213234","email":"userEmail@youapp.com"}
	c := jwt.StandardClaims{
		Audience:  "YOUR_AUDIENCE",
		ExpiresAt: 1562329281,
		Id:        "s1m7pfZHUzrncg4LGDhwoA",
		IssuedAt:  1562322081,
		NotBefore: 1562322021,
		Subject:   "YOUR_SUBJECT",
	}
	user := Customer{
		StandardClaims: c,
		UserId:         "1213234",
		Email:          "userEmail@youapp.com",
	}
	key, err := gojwk.Unmarshal([]byte(privateKeyJson))
	if err!=nil{
		panic(err)
	}
	priKey,err:=key.DecodePrivateKey()
	if err!=nil{
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, user)

	result, err := token.SignedString(priKey)
	if err != nil {
		panic(err)
	}
	println(result)
}
