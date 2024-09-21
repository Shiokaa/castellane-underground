package music

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Fonction pour jouer la musique à partir d'un fichier donné
func PlayMusic(filename string) {
	// Ouvrir le fichier MP3
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Décoder le fichier MP3
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialiser le speaker avec le format du fichier
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Jouer la musique
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Attendre la fin de la lecture
	<-done
}
