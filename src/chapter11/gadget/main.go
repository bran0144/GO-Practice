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

func main() {
	mixtape := []string{"My way", "Nice and easy", "Someone to watch over me"}
	var player Player = TapePlayer{}
	playList(player, mixtape)
	player = TapeRecorder{}
	playList(player, mixtape)
}
