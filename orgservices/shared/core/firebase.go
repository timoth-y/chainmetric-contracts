package core

import (
	"context"

	"firebase.google.com/go/v4"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

var Firebase *firebase.App

func initFirebase() {
	var err error

	if !viper.GetBool("firebase_enabled") {
		return
	}

	if Firebase, err = firebase.NewApp(context.Background(), nil,
		option.WithCredentialsFile(viper.GetString("firebase_credentials")),
	); err != nil {
		Logger.Fatalf("failed to initialize Firebase client: %v", err)
	}
}
