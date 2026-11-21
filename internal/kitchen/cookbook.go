package kitchen

import (
	"cmp"
	"slices"
)

var recipes = map[string]Recipe{}

func Register(recipe Recipe) {
	recipes[recipe.Name()] = recipe
}

func Find(name string) (Recipe, bool) {
	recipe, ok := recipes[name]
	return recipe, ok
}

func List() []Recipe {
	result := make([]Recipe, 0, len(recipes))

	for _, recipe := range recipes {
		result = append(result, recipe)
	}

	slices.SortFunc(result, func(a, b Recipe) int {
		return cmp.Compare(a.Name(), b.Name())
	})

	return result
}
