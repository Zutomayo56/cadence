// routes.go
// Router.

package main

import (
	"net/http"

	"gopkg.in/antage/eventsource.v1"
)

var radiodata_sse = eventsource.New(nil, nil)

func routes() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/api/radiodata/sse", radiodata_sse)
	r.Handle("/api/search", Search())
	r.Handle("/api/request/id", RequestID())
	r.Handle("/api/request/bestmatch", RequestBestMatch())
	r.Handle("/api/nowplaying/metadata", NowPlayingMetadata())
	r.Handle("/api/nowplaying/albumart", NowPlayingAlbumArt())
	r.Handle("/api/history", History())
	r.Handle("/api/listenurl", ListenURL())
	r.Handle("/api/listeners", Listeners())
	r.Handle("/api/bitrate", Bitrate())
	r.Handle("/api/version", Version())
	r.Handle("/ready", Ready())
	if c.DevMode {
		r.Handle("/api/dev/skip", DevSkip())
	}
	r.Handle("/", http.FileServer(http.Dir(c.RootPath+"./public/")))
	return r
}
