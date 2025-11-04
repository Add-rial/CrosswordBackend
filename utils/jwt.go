package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "errors"
    "CrosswordBackend/config"
)

type Claims struct {
    UserID uint
    jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(config.JwtKey)
}

func ValidateJWT(tokenString string) (uint, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
        return config.JwtKey, nil
    })

    if err != nil || !token.Valid {
        return 0, errors.New("invalid token")
    }

    return claims.UserID, nil
}
