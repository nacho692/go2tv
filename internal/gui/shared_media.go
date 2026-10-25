package gui

import (
	"path/filepath"
	"strings"

	"github.com/alexballas/refyne/v2"
)

func isTorrentDocument(uri fyne.URI, mime string) bool {
	if uri == nil {
		return false
	}
	mime, _, _ = strings.Cut(mime, ";")
	switch strings.ToLower(strings.TrimSpace(mime)) {
	case "application/x-bittorrent", "application/x-torrent":
		return true
	}
	return strings.EqualFold(filepath.Ext(uri.Name()), ".torrent")
}

// isRemoteMediaURI reports whether uri is a stream another app handed over by
// URL, e.g. Stremio's external player, rather than a document.
func isRemoteMediaURI(uri fyne.URI) bool {
	if uri == nil {
		return false
	}
	switch strings.ToLower(uri.Scheme()) {
	case "http", "https":
		return true
	}
	return false
}

// Match Android's registered document types even for explicit intents.
func acceptSharedURI(uri fyne.URI, mime string) bool {
	if isRemoteMediaURI(uri) {
		return !isTorrentDocument(uri, mime)
	}
	if uri == nil || !strings.EqualFold(uri.Scheme(), "content") {
		return false
	}
	if isTorrentDocument(uri, mime) {
		return true
	}
	mime, _, _ = strings.Cut(mime, ";")
	mime = strings.TrimSpace(mime)
	if mime == "" {
		return true
	}
	kind, _, _ := strings.Cut(mime, "/")
	switch strings.ToLower(kind) {
	case "video", "audio", "image":
		return true
	}
	return false
}
