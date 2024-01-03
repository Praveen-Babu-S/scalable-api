package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Praveen-Babu-S/scalable-api/models/dbmodels"
	"github.com/Praveen-Babu-S/scalable-api/pkg/common"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var JWTSecret = []byte("my-jwt-secret")

func (s *AuthServer) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user dbmodels.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user.Password = string(hashedPassword)
	s.db.Create(&user)

	common.RespondWithJSON(w, http.StatusCreated, "User registered successfully")
}

// LoginHandler handles user login and issues a JWT token.
func (s *AuthServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	var user dbmodels.User
	result := s.db.Where("username = ?", credentials.Username).First(&user)
	if result.Error != nil {
		common.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		common.RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	token, err := generateJWTToken(user.ID)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "Error generating JWT token")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// GenerateJWTToken generates a JWT token for the given user ID.
func generateJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // Token expires in 2 hours
	})

	return token.SignedString(JWTSecret)
}

// Middleware to authenticate requests using JWT token
func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})
		if err != nil || !token.Valid {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Pass the user ID to the next handler
		userID := uint(token.Claims.(jwt.MapClaims)["user_id"].(float64))
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
