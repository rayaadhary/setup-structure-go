package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rayaadhary/social-go/internal/service"
	"github.com/rayaadhary/social-go/internal/users"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(u *service.UserService) *AuthHandler {
	return &AuthHandler{UserService: u}
}

// GenerateJWT bikin token buat user
func GenerateJWT(userID int64) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 hari
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// simpen user id di context kalau butuh
		ctx := r.Context()
		ctx = WithUserID(ctx, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds users.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.Authenticate(r.Context(), creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users.LoginResponse{Token: token})
}
