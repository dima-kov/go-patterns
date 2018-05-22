package main

import (
	"github.com/dima-kov/go-patterns/composite/playlists/playlist"
	"github.com/dima-kov/go-patterns/composite/playlists/song"
)

func main() {
	lalaSong := song.NewSong("la-la-la.mp3")
	abbaSong := song.NewSong("abba.mp3")

	museSong := song.NewSong("muse.mp3")
	nestedPlaylist := playlist.NewPlaylist(museSong)
	rootPlaylist := playlist.NewPlaylist(lalaSong, nestedPlaylist, abbaSong)

	rootPlaylist.Play()
}
