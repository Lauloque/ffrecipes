package dishes

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Lauloque/ffrecipes/internal/kitchen"
)

type VideoToPNG struct{}

func (VideoToPNG) Name() string {
	return "video-to-png"
}

func (VideoToPNG) Description() string {
	return "Convert video files to PNG image sequences."
}

func (VideoToPNG) Cook(pantry *kitchen.Pantry) error {

	pantry.Log.Info("Let's pretend I am cooking!")

	for _, input := range pantry.Inputs {
		msg := fmt.Sprintf("  Found '%s'...", input)
		pantry.Log.Info(msg)

		base := strings.TrimSuffix(input, filepath.Ext(input))
		new_name := base + ".png"
		msg2 := fmt.Sprintf("  Will bake into '%s'!", new_name)
		pantry.Log.Info(msg2)

		pantry.Log.Debug("  Heating up oven...")
		pantry.Log.Warning("  Summer heat detected, might impact baking!")
		msg3 := fmt.Sprintf("  Cooked up '%s'!", new_name)
		pantry.Log.Debug(msg3)

	}
	pantry.Log.Success("Done!")
	return nil
}

func init() {
	kitchen.Register(VideoToPNG{})
}
