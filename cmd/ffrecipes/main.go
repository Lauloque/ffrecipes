package main

import (
	"fmt"
	"os"

	"github.com/Lauloque/ffrecipes/internal/app"
	"github.com/Lauloque/ffrecipes/internal/kitchen"
	"github.com/Lauloque/ffrecipes/internal/ui"
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

func getRecipeName(args []string) (string, bool) {
	if len(args) < 2 {
		fmt.Println("Missing recipe name")
		return "", false
	}
	return args[1], true
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
	if len(args) == 0 {
		printHelp()
		return
	}

	application := app.New(ui.Logger{Verbose: options.Verbose}, options)

	switch args[0] {
	case "help":
		printHelp()
	case "list":
		for _, recipe := range application.Recipes() {
			fmt.Println(recipe.Name())
		}
	case "info":
		name, ok := getRecipeName(args)
		if !ok {
			return
		}

		recipe, ok := application.Recipe(name)
		if !ok {
			fmt.Println("Unknown recipe:", name)
			return
		}

		fmt.Println(recipe.Name())
		fmt.Println(recipe.Description())
	case "cook":
		name, ok := getRecipeName(args)
		if !ok {
			return
		}

		recipe, ok := application.Recipe(name)
		if !ok {
			fmt.Println("Unknown recipe:", name)
			return
		}

		err := application.Cook(recipe, args[2:])
		if err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		return
	}
}
