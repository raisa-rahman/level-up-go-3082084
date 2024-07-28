package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`

	// Additional fields to keep track of album and song positions
	AlbumCount, SongCount int
}

// PlaylistHeap is a max-heap of Song entries.
type PlaylistHeap []Song

// Len returns the length of the heap
func (h PlaylistHeap) Len() int {
	return len(h)
}

// Less compares two heap elements (songs) and returns true if the first has a higher play count than the second
func (h PlaylistHeap) Less(i, j int) bool {
	// We want Pop to return the highest play count.
	return h[i].PlayCount > h[j].PlayCount
}

// Swap swaps two elements in the heap
func (h PlaylistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap
func (h *PlaylistHeap) Push(x any) {
	*h = append(*h, x.(Song))
}

// Pop removes and returns the last element from the heap
func (h *PlaylistHeap) Pop() any {
	original := *h
	n := len(original)
	x := original[n-1]
	*h = original[0 : n-1]
	return x
}

// makePlaylist merges and sorts the given albums into a single playlist
func makePlaylist(albums [][]Song) []Song {
	var playlist []Song
	pHeap := &PlaylistHeap{}
	if len(albums) == 0 {
		return playlist
	}

	// Initialize the heap and add the first song of each album
	heap.Init(pHeap)
	for i, f := range albums {
		firstSong := f[0]
		firstSong.AlbumCount, firstSong.SongCount = i, 0
		heap.Push(pHeap, firstSong)
	}

	// Process the heap until it's empty
	for pHeap.Len() != 0 {
		// Pop the song with the highest play count from the heap
		p := heap.Pop(pHeap)
		song := p.(Song)
		playlist = append(playlist, song)

		// If there are more songs in the current album, push the next song onto the heap
		if song.SongCount < len(albums[song.AlbumCount])-1 {
			nextSong := albums[song.AlbumCount][song.SongCount+1]
			nextSong.AlbumCount, nextSong.SongCount = song.AlbumCount, song.SongCount+1
			heap.Push(pHeap, nextSong)
		}
	}

	return playlist
}

func main() {
	// Import the album data from the JSON file
	albums := importData()
	// Generate the playlist and print it as a table
	printTable(makePlaylist(albums))
}

// printTable prints the merged playlist in a formatted table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()
}

// importData reads the input data from the file and unmarshals it into a 2D slice of Songs
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}