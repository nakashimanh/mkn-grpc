package jwt

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

const authorizationKey = "authorization"

type jwtConfig struct {
	Sub      string
	Name     string
	Iat      int64
	Exp      int64
	Role     string
	AppName  string
	DeviceID string
	Secret   string
}

func getConfig(role string, appName string, deviceID string) jwtConfig {
	return jwtConfig{
		"app-dummy-token",
		"zfb system",
		time.Now().Unix(),
		time.Now().Add(time.Hour * 3).Unix(),
		role,
		appName,
		deviceID,
		"secret",
	}
}

// ExtractAuthorization extracts token from metadata
func extractAuthorization(ctx context.Context) (string, error) {
	return fromMeta(ctx, authorizationKey)
}

func fromMeta(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("not found metadata")
	}
	vs := md[key]
	if len(vs) == 0 {
		return "", fmt.Errorf("not found %s in metadata", key)
	}
	return vs[0], nil
}

// GenerateToken is function which generates a jwt
func GenerateToken(operatorID string, deviceID string) (string, error) {
	c := getConfig("operator", "android", deviceID)
	//secret := os.Getenv("SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        c.Sub,
		"name":       c.Name,
		"iat":        c.Iat,
		"exp":        c.Exp,
		"role":       c.Role,
		"zfb-app-id": c.DeviceID,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}

// TokenVerifyMiddleWare is a fucntion verify token
func TokenVerifyMiddleWare(ctx context.Context) error {
	var err error
	//meta, _ := metadata.FromIncomingContext(ctx)
	authHeader, _ := extractAuthorization(ctx)
	log.Printf("authHeader = %v\n", authHeader)
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		authToken := bearerToken[1]
		//log.Println(authToken)
		token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error: %s", "There was an error")

			}

			return []byte("secret"), nil
			//return []byte(os.Getenv("SIGNINGKEY")), nil
		})

		if error != nil {
			err = error
			log.Printf("Parse error = %v\n", error.Error())
			return err
		}

		if token.Valid {
			log.Printf("Token is OK %v", token.Valid)
		}
	} else {

		log.Printf("Invalid token.")
		return fmt.Errorf("Error: %s", "Invalid token")
	}

	return err
}
