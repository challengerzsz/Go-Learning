package mp

import (
	"fmt"
	"mlib"
)

type Player interface {
	Play(source string)
}

func Play(source string, mtype string) {

	var p Player
	switch mtype {
	case mlib.MP3:
		p = &MP3Player{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
