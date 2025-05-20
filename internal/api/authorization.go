package api

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/AnnonaOrg/annona_client/utils"
	// log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

var tdlibClient *client.Client

func ClientAuthorize(apiIdRaw, apiHash string) (tClient *client.Client, err error) {
	apiId64, err := strconv.ParseInt(apiIdRaw, 10, 32)
	if err != nil {
		// log.Fatalf("strconv.Atoi error: %s", err)
		return nil, err
	}
	apiId := int32(apiId64)
	// 清理历史文件 .tdlib/database/db.sqlite "/app/.tdlib/database/db.sqlite"
	if err := utils.Remove(filepath.Join(".tdlib", "database", "db.sqlite")); err != nil {
		fmt.Printf("Remove(): %v\n", err)
	}
	tdlibParameters := &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               apiId,
		ApiHash:             apiHash,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}
	authorizer := client.ClientAuthorizer(tdlibParameters)
	go client.CliInteractor(authorizer)

	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		// log.Fatalf("SetLogVerbosityLevel error: %s", err)
		return nil, err
	}

	tdlibClient, err = client.NewClient(authorizer)
	return tdlibClient, err
}
