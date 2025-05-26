package api

import (
	"fmt"
	"github.com/AnnonaOrg/annona_client/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
	"path/filepath"
	"strconv"
	"time"
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

	// New value of the verbosity level for logging.
	//Value 0 corresponds to fatal errors,
	//value 1 corresponds to errors,
	//value 2 corresponds to warnings and debug warnings,
	//value 3 corresponds to informational,
	//value 4 corresponds to debug,
	//value 5 corresponds to verbose debug,
	//value greater than 5 and up to 1023 can be used to enable even more logging
	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})

	if err != nil {
		// log.Fatalf("SetLogVerbosityLevel error: %s", err)
		log.Errorf("SetLogVerbosityLevel(): %v", err)
		return nil, err
	}

	//proxy := client.WithProxy(&client.AddProxyRequest{
	//	Server: osenv.GetSocks5ProxyServer(),
	//	Port:   int32(osenv.GetSocks5ProxyPortInt()),
	//	Enable: osenv.IsEnableSocks5Proxy(),
	//	Type: &client.ProxyTypeSocks5{
	//		Username: osenv.GetSocks5ProxyUsername(),
	//		Password: osenv.GetSocks5ProxyPassword(),
	//	},
	//})
	catchTimeout := client.WithCatchTimeout(2 * time.Minute)
	//tdlibClient, err = client.NewClient(authorizer, proxy, catchTimeout)
	tdlibClient, err = client.NewClient(authorizer, catchTimeout)
	if err != nil {
		log.Errorf("NewClient(): %v", err)
	}
	return tdlibClient, err
}
