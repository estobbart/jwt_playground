package main

import(
  "fmt"
  "time"

  "gopkg.in/square/go-jose.v2/jwt"
  "gopkg.in/square/go-jose.v2"
)

func main(){
  key := []byte("secret")
  sig, err := jose.NewSigner(jose.SigningKey{
                              Algorithm: jose.HS256,
                              Key: key },
                            (&jose.SignerOptions{}).WithType("JWT"))
  if err != nil {
      panic(err)
  }

  cl := jwt.Claims{
      Subject:   "subject",
      Issuer:    "issuer",
      NotBefore: jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
      Audience:  jwt.Audience{"foo", "bar"},
  }
  raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
  if err != nil {
      panic(err)
  }

  fmt.Println(raw)
}
