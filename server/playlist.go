package server

import (
	"github.com/otaconix/m3ufilter/m3u"
)

func updatePlaylist(conf *httpConfig) {
	if conf.lock {
		log.Info("Retrieval is locked, trying again next time...")
		return
	}

	if conf.appConfig.Core.AutoReloadConfig {
		conf.appConfig.Load()
	}

	conf.lock = true
	log.Info("Updating playlists")
	newPlaylists, allFailed := m3u.GetPlaylist(conf.appConfig)
	if allFailed {
		log.Info("Skipping updating playlist to server as all providers failed")
	} else {
		conf.playlists = &newPlaylists
	}
	log.Info("Done")
	conf.lock = false
}
