package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Lauloque/ffrecipes/internal/app"
	"github.com/Lauloque/ffrecipes/internal/kitchen/ovens"
	"github.com/Lauloque/ffrecipes/internal/ui"
)

const (
	exitOK    = 0
	exitError = 1
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
		fmt.Fprintln(os.Stderr, "Missing recipe name")
		return "", false
	}
	return args[1], true
}

func parseArgs(args []string) ([]string, bool) {
	verbose := false

	for len(args) > 0 {
		switch args[0] {
		case "--verbose":
			verbose = true
			args = args[1:]
		default:
			return args, verbose
		}
	}
	return args, verbose
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	args, verbose := parseArgs(args)
	if len(args) == 0 {
		printHelp()
		return exitOK
	}

	application := app.New(ui.Logger{Verbose: verbose}, ovens.NewFFmpeg("ffmpeg"))

	switch args[0] {
	case "help":
		printHelp()
		return exitOK
	case "list":
		for _, recipe := range application.Recipes() {
			fmt.Println(recipe.Name())
		}
		return exitOK
	case "info":
		name, ok := getRecipeName(args)
		if !ok {
			return exitError
		}

		recipe, ok := application.Recipe(name)
		if !ok {
			fmt.Fprintln(os.Stderr, "Unknown recipe:", name)
			return exitError
		}

		fmt.Println(recipe.Name())
		fmt.Println(recipe.Description())
		return exitOK
	case "cook":
		name, ok := getRecipeName(args)
		if !ok {
			return exitError
		}

		recipe, ok := application.Recipe(name)
		if !ok {
			fmt.Fprintln(os.Stderr, "Unknown recipe:", name)
			return exitError
		}

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		if err := application.Cook(ctx, recipe, args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return exitError
		}
		return exitOK
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", args[0])
		return exitError
	}
}
