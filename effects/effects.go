package effects

import (
	"time"

	"github.com/daanikus/golight/lights"
)

func Stream(lights *lights.Lights, palette *lights.Palette, duration time.Duration) {
	start := time.Now()
	for _, colour := range *palette {
		for i := 0; i < lights.Size(); i++ {
			if time.Since(start) > duration {
				return
			}
			lights.Set(i, colour)
			lights.Show()
			time.Sleep(100 * time.Millisecond)
		}
	}
}
