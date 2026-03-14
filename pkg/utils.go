package pkg

import (
    "os"
    "time"

    jwtlib "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// ── Hash password ──────────────────────────────────────────

func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed), err
}

func CheckPassword(hashed, plain string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

// ── JWT ────────────────────────────────────────────────────

type JWTClaims struct {
    UserID uint `json:"user_id"`
    jwtlib.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
    claims := JWTClaims{
        UserID: userID,
        RegisteredClaims: jwtlib.RegisteredClaims{
            ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwtlib.NewNumericDate(time.Now()),
        },
    }
    token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenStr string) (*JWTClaims, error) {
    token, err := jwtlib.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwtlib.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return token.Claims.(*JWTClaims), nil
}