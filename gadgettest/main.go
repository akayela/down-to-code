package main

import "github.com/akayela/gadget"

type Player interface {
    Play(string)
    Stop()
}

func TryOut(player Player) {
    player.Play("Destiny")
    player.Stop()
    recorder, ok := player.(gadget.TapeRecorder)
    if ok {
        recorder.Record()
    }        
}

func playList(device Player, songs []string) {
    for _, song := range songs {
        device.Play(song)
    }
    device.Stop()
}

func main() {
    mixtape := []string{"Redemption song", "Iron lion zion", "Destiny", "Untold stories"}
    var player Player = gadget.TapePlayer{}
    playList(player, mixtape)
    player = gadget.TapeRecorder{}
    playList(player, mixtape)

    TryOut(gadget.TapeRecorder{})
    TryOut(gadget.TapePlayer{})
}
