package playlist

import (
	"fmt"
	"github.com/dima-kov/go-patterns/composite/playlists/play_component"
)

type playlist struct {
	playComponents []play_component.PlayComponent
}

func NewPlaylist(components ...play_component.PlayComponent) play_component.PlayComponent {
	return playlist{playComponents: components}
}

func (p playlist) Play() {
	fmt.Println("Playing playlist.")
	for _, item := range p.playComponents {
		//fmt.Println("Start playing first item in playlist")
		item.Play()
	}
}
