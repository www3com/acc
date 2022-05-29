package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	duration := 1 * time.Hour
	token, err := GenerateToken(1, "200.22.11.31", duration)
	if err != nil {
		t.Errorf("产生token错误, %v", err)
		return
	}

	fmt.Printf("jwt: %s \n", token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImlwIjoiMjAwLjIyLjExLjMxIiwiZXhwIjoxNjUzNjY5ODkzLCJqdGkiOiI2Yzc1OGE4Yi0xOTdhLTQyOTEtOWEyYi1kYTZjMmEwNDI0ZjAifQ.VmiGq9IZW3xvBPCYYPwoY-V1sGQF8ZM96txuxYp3Sxw"
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("解析token错误, %v", err)
		return
	}

	fmt.Printf("uid: %v\n", claims.Uid)
	fmt.Printf("ip: %v\n", claims.IP)
	fmt.Printf("id: %v\n", claims.RegisteredClaims.ID)
	fmt.Printf("expire time: %v\n", claims.RegisteredClaims.ExpiresAt)
}
