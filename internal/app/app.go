package app

import (
	"github.com/Lauloque/ffrecipes/internal/kitchen"
	_ "github.com/Lauloque/ffrecipes/recipes" // Register built-in recipes.
)

// App coordinates recipe lookup and execution for a user interface.
type App struct {
	log kitchen.Logger
}

// New returns an App that logs through log.
func New(log kitchen.Logger) *App {
	return &App{
		log: log,
	}
}

// Recipes returns the registered recipes.
func (a *App) Recipes() []kitchen.Recipe {
	return kitchen.List()
}

// Recipe looks up a recipe by name.
func (a *App) Recipe(name string) (kitchen.Recipe, bool) {
	return kitchen.Find(name)
}

// Cook runs recipe with the provided inputs.
func (a *App) Cook(recipe kitchen.Recipe, inputs []string) error {
	pantry := kitchen.Pantry{
		Inputs: inputs,
		Log:    a.log,
	}

	return recipe.Cook(&pantry)
}
