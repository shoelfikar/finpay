package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var signedKey = []byte(viper.GetString("JWT_SECRET")) // Kunci rahasia untuk tanda tangan JWT


// JWTPayload Struktur payload JWT berisi info user dan klaim standar JWT
type JWTPayload struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT is func to generate the token
func GenerateJWT(duration time.Duration) (*string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		log.Errorf("error generate uuid token %v", err)
		return nil, err
	}

	now := time.Now()
	exp := now.Add(duration)

	payload := JWTPayload{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "user_login",            // Pengeluarnya token
			Subject:   "go-escape",             // Subjek token
			ID:        tokenID.String(),        // ID unik token
			IssuedAt:  jwt.NewNumericDate(now), // Waktu dibuat token
			NotBefore: jwt.NewNumericDate(now), // Token berlaku sejak waktu ini
			ExpiresAt: jwt.NewNumericDate(exp), // Waktu kadaluarsa token
		},
	}


	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(signedKey)
	if err != nil {
		log.Errorf("[SERVER ERROR] %v", err)
		return nil, err
	}

	log.Info("success generate jwt token")
	return &token, nil
}

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	
	// Check if header is empty or doesn't start with "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Error("Missing or invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
	}

	// Extract token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	// If token is invalid
	if err != nil || !token.Valid {
		log.Error("Invalid token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Extract claims (Assuming it's a standard JWT)
	claims, ok := token.Claims.(*JWTPayload)
	if !ok {
		log.Error("Invalid token claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	// Store claims in Fiber locals
	c.Locals("claims", claims)

	return c.Next()
}