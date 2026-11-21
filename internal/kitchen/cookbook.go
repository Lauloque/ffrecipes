package kitchen

var recipes = map[string]Recipe{}

func Register(recipe Recipe) {
	recipes[recipe.Name()] = recipe
}

func Find(name string) (Recipe, bool) {
	recipe, ok := recipes[name]
	return recipe, ok
}

func List() []Recipe {
	result := []Recipe{}

	for _, recipe := range recipes {
		result = append(result, recipe)
	}
	return result
}
