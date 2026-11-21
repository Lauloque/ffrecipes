package kitchen

import (
	"context"
	"testing"
)

type testRecipe struct {
	name string
}

func (r testRecipe) Name() string                        { return r.name }
func (r testRecipe) Description() string                 { return "" }
func (r testRecipe) Cook(context.Context, *Pantry) error { return nil }

func TestListSortsByName(t *testing.T) {
	original := recipes
	t.Cleanup(func() { recipes = original })

	recipes = map[string]Recipe{
		"zucchini": testRecipe{name: "zucchini"},
		"apple":    testRecipe{name: "apple"},
		"mango":    testRecipe{name: "mango"},
	}

	got := List()
	want := []string{"apple", "mango", "zucchini"}

	if len(got) != len(want) {
		t.Fatalf("got %d recipes, want %d", len(got), len(want))
	}

	for i, recipe := range got {
		if recipe.Name() != want[i] {
			t.Errorf("recipe %d = %q, want %q", i, recipe.Name(), want[i])
		}
	}
}
