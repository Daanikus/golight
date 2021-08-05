package lights

import (
	"log"

	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

type Colour struct {
	Red   uint8
	Green uint8
	Blue  uint8
}
type Palette []Colour

var Palettes = map[string]Palette{
	"synthwave": {
		Colour{0x29, 0x0C, 0xFF},
		Colour{0x9B, 0x00, 0xE8},
		Colour{0xFF, 0x01, 0x9A},
	},
}

type Lights struct {
	conn      spi.Conn
	numPixels int
	buf       []byte
	p         spi.PortCloser
}

// New returns a new Lights instance connected with SPI to WS2801.
func New(numPixels int) (*Lights, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	p, err := spireg.Open("")
	if err != nil {
		return nil, err
	}

	c, err := p.Connect(physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		return nil, err
	}

	return &Lights{
		conn:      c,
		numPixels: numPixels,
		buf:       make([]byte, 32*3),
		p:         p,
	}, nil
}

// Set sets the LED at index to a colour. Call Show to write the buffer to the strip.
func (l *Lights) Set(index int, colour Colour) {
	offset := index * 3
	l.buf[offset] = colour.Red
	l.buf[offset+1] = colour.Green
	l.buf[offset+2] = colour.Blue
}

// Close closes the connection to the strip.
func (l *Lights) Close() error {
	return l.p.Close()
}

// Size returns the number of pixels in the strip.
func (l *Lights) Size() int {
	return l.numPixels
}

func (l *Lights) Off() {
	l.buf = make([]byte, l.numPixels*3)
	l.write()
}

func (l *Lights) Show() {
	l.write()
}

func (l *Lights) write() {
	// Must provide a read buffer with same length as write buffer
	read := make([]byte, len(l.buf))
	if err := l.conn.Tx(l.buf, read); err != nil {
		log.Fatal(err)
	}
}
