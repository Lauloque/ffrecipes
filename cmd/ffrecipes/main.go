package main

import (
	"fmt"
	"os"

	"github.com/Lauloque/ffrecipes/internal/kitchen"
	"github.com/Lauloque/ffrecipes/internal/ui"
	_ "github.com/Lauloque/ffrecipes/recipes"
)

func printHelp() {
	fmt.Println(`
ffrecipes - cook your FFmpeg recipes

Usage:
  ffrecipes [options] <command> [arguments]

Options:
  --verbose     Show detailed output

Commands:
  cook          Cook a recipe
  help			Show this help
  info          Show recipe information
  list          List available recipes

Examples:
  ffrecipes list
  ffrecipes info video-to-png
  ffrecipes --verbose cook video-to-png movie.mp4`)
}

func getRecipe(args []string) (kitchen.Recipe, bool) {
	if len(args) < 2 {
		fmt.Println(("Missing recipe name"))
		return nil, false
	}

	recipe, ok := kitchen.Find(args[1])
	if !ok {
		fmt.Println("Unknown recipe:", args[1])
		return nil, false
	}
	return recipe, true
}

func getArgsAndOptions() ([]string, kitchen.Options) {
	args := os.Args[1:]

	options := kitchen.Options{}

	for len(args) > 0 {
		switch args[0] {
		case "--verbose":
			options.Verbose = true
			args = args[1:]
		default:
			return args, options
		}
	}
	return args, options
}

func main() {

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	args, options := getArgsAndOptions()

	switch args[0] {
	case "help":
		printHelp()
	case "list":
		for _, recipe := range kitchen.List() {
			fmt.Println(recipe.Name())
		}
	case "info":
		recipe, ok := getRecipe(args)
		if !ok {
			return
		}

		fmt.Println(recipe.Name())
		fmt.Println(recipe.Description())
	case "cook":
		recipe, ok := getRecipe(args)
		if !ok {
			return
		}

		pantry := kitchen.Pantry{
			Inputs: args[2:],
			Log: ui.Logger{
				Verbose: options.Verbose,
			},
			Options: options,
		}
		err := recipe.Cook(&pantry)

		if err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		return
	}

	// logger := ui.Logger{}

	// pantry := kitchen.Pantry{
	// 	Inputs: []string{
	// 		"example.mp4",
	// 	},

	// 	FFmpegPath: "ffmpeg",

	// 	Logger: logger,
	// }

	// fmt.Println("ffrecipes starting")

	// _ = pantry

	// _ = os.Args
}
