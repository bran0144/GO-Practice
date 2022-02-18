package gadget

type Player interface {
	Play(string)
	Stop()
}
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func TryOut(player Player) {
	player.Play("test track")
	player.Stop()
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	mixtape := []string{"My way", "Nice and easy", "Someone to watch over me"}
	var player Player = TapePlayer{}
	playList(player, mixtape)
	player = TapeRecorder{}
	playList(player, mixtape)
	TryOut(TapeRecorder{})
	TryOut(TapePlayer{})
}
