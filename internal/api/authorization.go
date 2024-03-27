package api

import (
	"path/filepath"
	"strconv"

	// log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

var tdlibClient *client.Client

func ClientAuthorize(apiIdRaw, apiHash string) (tClient *client.Client, err error) {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	apiId64, err := strconv.ParseInt(apiIdRaw, 10, 32)
	if err != nil {
		// log.Fatalf("strconv.Atoi error: %s", err)
		return nil, err
	}
	apiId := int32(apiId64)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})
	if err != nil {
		// log.Fatalf("SetLogVerbosityLevel error: %s", err)
		return nil, err
	}

	tdlibClient, err = client.NewClient(authorizer)
	return tdlibClient, err
}

// Authorize logs the bot into the provided account using tdlib.
func BotAuthorize(apiIdRaw, apiHash string, botToken string) (tClient *client.Client, err error) {
	authorizer := client.BotAuthorizer(botToken)

	apiId64, err := strconv.ParseInt(apiIdRaw, 10, 32)
	if err != nil {
		// log.Fatalf("strconv.Atoi error: %s", err)
		return nil, err
	}
	apiId := int32(apiId64)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:              false,
		DatabaseDirectory:      filepath.Join(".tdlib", "database"),
		FilesDirectory:         filepath.Join(".tdlib", "files"),
		UseFileDatabase:        true,
		UseChatInfoDatabase:    true,
		UseMessageDatabase:     true,
		UseSecretChats:         false,
		ApiId:                  apiId,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: true,
		IgnoreFileNames:        false,
	}

	logVerbosity := client.WithLogVerbosity(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 0,
	})

	tdlibClient, err = client.NewClient(authorizer, logVerbosity)
	return tdlibClient, err
}
