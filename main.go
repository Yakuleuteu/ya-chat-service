package main

func main() {
	type JWTClaims struct {
		Role        string   `json:"role"`
		Authorities []string `json:"authorities"`
	}

}
