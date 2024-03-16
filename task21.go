package main

import (
	"fmt"
)

// Интерфейс, который мы хотим использовать
type MediaPlayer interface {
	playAudio()
}

// Структура, представляющая плеер, который мы хотим использовать
type MusicPlayer struct{}

func (m *MusicPlayer) playAudio() {
	fmt.Println("Воспроизведение аудио")
}

// Интерфейс адаптера
type AdvancedMediaPlayer interface {
	playVideo()
}

// Структура, представляющая продвинутый плеер, который мы хотим адаптировать
type VideoPlayer struct{}

func (v *VideoPlayer) playVideo() {
	fmt.Println("Воспроизведение видео")
}

// Адаптер, преобразующий интерфейс AdvancedMediaPlayer в интерфейс MediaPlayer
type MediaPlayerAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func (a *MediaPlayerAdapter) playAudio() {
	a.advancedPlayer.playVideo()
}

func main() {
	musicPlayer := &MusicPlayer{}
	musicPlayer.playAudio()

	videoPlayer := &VideoPlayer{}

	adapter := &MediaPlayerAdapter{advancedPlayer: videoPlayer}
	adapter.playAudio()
}
