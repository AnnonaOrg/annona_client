package main

import (
	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/repository"
	"github.com/AnnonaOrg/annona_client/internal/updates"
	"github.com/AnnonaOrg/osenv"
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

func init() {
	apiIdRaw := osenv.GetAppTelegramApiID()
	apiHash := osenv.GetAppTelegramApiHash()

	tdlibClient, err := api.ClientAuthorize(apiIdRaw, apiHash)
	if err != nil {
		log.Fatalf("api.ClientAuthorize(%s) error: %s", apiIdRaw, err)
	}
	repository.Tdlib = tdlibClient

	optionValue, err := client.GetOption(&client.GetOptionRequest{
		Name: "version",
	})
	if err != nil {
		log.Fatalf("GetOption error: %s", err)
	}
	log.Infof("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

	me, err := tdlibClient.GetMe()
	if err != nil {
		log.Fatalf("GetMe error: %s", err)
	}
	repository.Me = me

	log.Infof("Me: %s %s [%v]", repository.Me.FirstName, repository.Me.LastName, repository.Me.Usernames)
}
func listenTdlib() {
	listener := repository.Tdlib.GetListener()
	defer listener.Close()
	updates.HandleUpdates(listener)
}
