package song

import (
	"fmt"
	"github.com/dima-kov/go-patterns/composite/playlists/play_component"
)

type song struct {
	file string
}

func NewSong(file string) play_component.PlayComponent {
	return song{file}
}

func (s song) Play() {
	fmt.Printf("Start play %s song \n", s.file)
}
