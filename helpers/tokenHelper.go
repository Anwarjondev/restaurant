package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Anwarjondev/restaurant.git/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct{
	Email string
	FirstName string
	LastName string
	Uid string
	jwt.StandartClaims
}


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")


func GenerateAllToken(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Uid: uid,
		StandartClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandartClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err 
}


func UpdateAllTokens(signedToken string, signedRefreshToken string, userID string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj =append(updateObj, bson.E{"token", signedRefreshToken})
	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true
	
	defer cancel()
}


func ValidateToken(signedToken string)(claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}


