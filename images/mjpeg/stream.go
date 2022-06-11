// Package mjpeg implements a simple MJPEG streamer.
//
// Stream objects implement the http.Handler interface, allowing to use them with the net/http package like so:
//	stream = mjpeg.NewStream()
//	http.Handle("/camera", stream)
// Then push new JPEG frames to the connected clients using stream.UpdateJPEG().
package mjpeg

import (
	"bytes"
	"fmt"
	"image"
	j "image/jpeg"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/image/draw"
)

// Stream represents a single video feed.
type Stream struct {
	m             map[chan []byte]bool
	frame         []byte
	lock          sync.Mutex
	FrameInterval time.Duration
}

const boundaryWord = "MJPEGBOUNDARY"
const headerf = "\r\n" +
	"--" + boundaryWord + "\r\n" +
	"Content-Type: image/jpeg\r\n" +
	"Content-Length: %d\r\n" +
	"X-Timestamp: 0.000000\r\n" +
	"\r\n"

// ServeHTTP responds to HTTP requests with the MJPEG stream, implementing the http.Handler interface.
func (s *Stream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Stream:", r.RemoteAddr, "connected")
	w.Header().Add("Content-Type", "multipart/x-mixed-replace;boundary="+boundaryWord)

	c := make(chan []byte)
	s.lock.Lock()
	s.m[c] = true
	s.lock.Unlock()

	for {
		time.Sleep(s.FrameInterval)
		b := <-c
		_, err := w.Write(b)
		if err != nil {
			break
		}
	}

	s.lock.Lock()
	delete(s.m, c)
	s.lock.Unlock()
	log.Println("Stream:", r.RemoteAddr, "disconnected")
}

// UpdateJPEG pushes a new JPEG frame onto the clients.
func (s *Stream) UpdateJPEG(jpeg []byte) {
	header := fmt.Sprintf(headerf, len(jpeg))
	if len(s.frame) < len(jpeg)+len(header) {
		s.frame = make([]byte, (len(jpeg)+len(header))*2)
	}

	// Decode the image (from PNG to image.Image):
	src, _ := j.Decode(bytes.NewReader(jpeg))

	// Set the expected size that you want:
	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/3, src.Bounds().Max.Y/3))

	// Resize:
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	buf := new(bytes.Buffer)

	// Encode to `buf`:
	j.Encode(buf, dst, nil)

	copy(s.frame, header)
	//	copy(s.frame[len(header):], jpeg)
	copy(s.frame[len(header):], buf.Bytes())

	s.lock.Lock()
	for c := range s.m {
		// Select to skip streams which are sleeping to drop frames.
		// This might need more thought.
		select {
		case c <- s.frame:
		default:
		}
	}
	s.lock.Unlock()
}

// NewStream initializes and returns a new Stream.
func NewStream() *Stream {
	return &Stream{
		m:             make(map[chan []byte]bool),
		frame:         make([]byte, len(headerf)),
		FrameInterval: 50 * time.Millisecond,
	}
}
