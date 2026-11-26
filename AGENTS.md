# AGENTS.md

## Project context

FFrecipes is an open-source collection of reusable, focused FFmpeg
workflows, presented through a cooking metaphor. It is intended to make
small, sometimes highly specific media-processing tasks easy to
discover, reuse, and run.

FFrecipes is **not** intended to become a general-purpose graphical
media encoder. Avoid feature expansion that would turn it into a
replacement for tools such as HandBrake.

The project is FFmpeg-first. Other processing tools may be added when a
recipe has a concrete need for them, but do not introduce abstractions
for hypothetical tools or workflows.

## Current implementation state

FFrecipes is an early prototype, and parts of the documented design are
not implemented yet. Check the code before relying on an assumption:

-   `cmd/ffrecipes/main.go` parses options, creates a signal-aware
    context, and delegates to `internal/app`, which looks up the recipe,
    builds a Pantry, and calls `Recipe.Cook`.
-   The `recipes` directory contains two working recipes:
    `video-to-png` extracts every frame as `<input>_%04d.png` next to each
    input, and `add-cover-to-video` attaches a JPG or PNG cover to MP4 or
    MKV files as `<input>.cover.<ext>`.
-   `internal/kitchen/ovens/ffmpeg.go` implements `kitchen.Runner` by
    executing FFmpeg. It captures FFmpeg output and includes it in
    returned errors.
-   The application wires `ovens.NewFFmpeg("ffmpeg")` into the Pantry.
-   Shared input-file validation lives in `internal/files` and is used by
    both recipes.
-   Tests cover both recipes and the shared file helper; most packages
    still have no tests.
-   `test.mp4` in the repository root is a local sample input, ignored
    by Git. Generated `test_*.png` and `test.cover.*` files are ignored
    too. Leave the sample input in place.

Keep this list accurate when the implementation changes.

## Design principles

-   Keep each recipe focused on one task.
-   Keep responsibilities separated and dependencies explicit.
-   Prefer small, understandable changes over broad refactors.
-   Keep recipes independent of the way the application receives user
    input.
-   Keep shared orchestration in `internal/app`; frontends should only
    parse input and present results.
-   Recipes must receive their inputs and runtime resources through the
    Pantry.
-   Recipes must not parse command-line arguments or inspect `os.Args`.
-   Recipes must not locate or install external executables themselves;
    executable paths come from the Pantry, which the application is
    responsible for populating.
-   Recipes must run external tools through `pantry.FFmpeg`; do not call
    `os/exec` from recipe code.
-   Pass `context.Context` through recipe and runner calls and honor
    cancellation.
-   Recipes must not interact directly with the user through terminal
    input or output. Use the project's logging abstraction and return
    errors.
-   Return errors to the caller; do not terminate the process from
    recipe code.
-   Register recipes with the Cookbook using the existing project
    convention.
-   Keep external-tool integration in `internal/kitchen/ovens`.
-   Do not assume every recipe needs every oven.

## Project vocabulary

Use the cooking terminology consistently where it clarifies the design:

-   **Application**: the user-facing entry point. The CLI is the current
    interface, not a core kitchen component. Shared orchestration lives
    in `internal/app`.
-   **Kitchen**: the core package and execution environment.
-   **Cookbook**: registers and finds available recipes.
-   **Recipe**: one focused, reusable workflow.
-   **Pantry**: the inputs, logger, and other resources made available
    to a recipe.
-   **Oven**: an external processing tool integrated for recipes to use.
-   **FFmpeg oven**: the primary processing tool, implemented in
    `internal/kitchen/ovens`.

Do not force the metaphor onto concepts where it becomes confusing. For
example, ordinary supporting utilities do not need to be renamed as
kitchen equipment.

## Code layout and recipe conventions

-   The module is `github.com/Lauloque/ffrecipes` and `go.mod` targets
    Go 1.26.4. It has no third-party dependencies; keep it that way
    unless a dependency clearly pays off.
-   A recipe is a struct implementing `kitchen.Recipe`: `Name() string`,
    `Description() string`, and
    `Cook(ctx context.Context, pantry *kitchen.Pantry) error`.
-   Put each recipe in its own file under `recipes/`, named after the
    recipe (for example `video_to_png.go` for `video-to-png`).
-   The `recipes/` directory uses `package recipes`.
-   Register recipes in `init()` with `kitchen.Register`. The
    `internal/app` package blank-imports the recipes package, so there is
    no central list to edit and no explicit registration step to add.
-   Recipe names are kebab-case and are the public CLI identifiers.
    Treat renaming one as a user-visible change.
-   Inputs reach a recipe as `pantry.Inputs`, the positional arguments
    after the recipe name. The recipe decides how to interpret them.
-   Run external tools through `pantry.FFmpeg`, which is the
    `kitchen.Runner` interface. `internal/kitchen/ovens` provides the
    FFmpeg implementation; recipe tests can supply a fake runner.
-   Log through `pantry.Log`, which is the `kitchen.Logger` interface
    (`Info`, `Success`, `Warning`, `Error`, `Debug`). `internal/ui`
    provides the CLI implementation. `Debug` prints only when
    `--verbose` is set. Do not print directly from recipe code.
-   The CLI supports `list`, `info <recipe>`,
    `cook <recipe> [inputs...]`, and `help`. Options are parsed only
    before the command, so `--verbose` must precede the subcommand.
-   CLI commands return exit status 0 on success and 1 on errors, and
    write errors to stderr.

## Go conventions

-   Follow idiomatic Go and the conventions already present in the
    repository.
-   Use `gofmt` for Go source files.
-   Keep package names short, lowercase, and consistent with Go
    conventions.
-   Avoid unnecessary dependencies. Prefer the standard library unless a
    dependency provides a clear benefit.
-   Handle errors explicitly and wrap them with useful context when
    appropriate.
-   Avoid panics for expected runtime failures, such as missing inputs
    or an unavailable executable.
-   Keep exported identifiers documented when Go documentation
    conventions require it. Most existing exported identifiers are not
    documented; add doc comments to new ones rather than imitating the
    current gaps.
-   Do not add comments that merely restate the code.
-   The existing code has a few naming deviations, such as `new_name`
    in `recipes/video_to_png.go`. Use idiomatic names in new code and
    fix deviations in code you touch.

## Working on the codebase

Before changing code:

1.  Inspect the relevant files and their callers.
2.  Check the existing interfaces and package boundaries rather than
    assuming that an earlier design proposal has already been
    implemented.
3.  Make the smallest change that solves the requested problem.
4.  Update documentation when a change affects project structure, usage,
    architecture, or contributor workflows.

Do not silently implement speculative features. If a requested change
would require a significant architectural decision (for example,
changing the `kitchen.Runner` interface or adding a new external tool),
explain the trade-offs and ask before committing to one.

## Git workflow

-   Leave changes unstaged for review unless the user asks you to stage
    them.
-   Never run `git commit` without explicit approval. You may stage
    changes and prepare a commit, but wait for approval before
    committing.
-   Never run `git push` without a specific request. Approval to commit
    does not imply approval to push.

## Testing and validation

After changing Go code, run the relevant checks when possible:

``` sh
gofmt -w <changed-go-files>
go build ./...
go vet ./...
go test ./...
```

Tests currently cover the `video-to-png` recipe. Add tests when you
introduce logic worth testing, such as input validation or argument
building, rather than treating a passing `go test ./...` as meaningful
coverage.

For CLI changes, also exercise the affected command manually, including
missing arguments, unknown recipes, and invalid input where relevant.
Run cooking tests in a temporary directory so generated frames do not
pollute the repository:

``` sh
mkdir -p /tmp/ffrecipes-test
cp test.mp4 /tmp/ffrecipes-test/
go run ./cmd/ffrecipes list
go run ./cmd/ffrecipes info video-to-png
go run ./cmd/ffrecipes --verbose cook video-to-png /tmp/ffrecipes-test/test.mp4
```

Check that invalid commands and failing cooks return a non-zero exit
status.

Do not claim that tests or commands passed unless they were actually
run. If a check cannot be run, say so.

## Documentation

-   Keep `README.md` approachable for users who may not know Go or
    software architecture terminology, and aligned with the actual CLI.
-   Keep `docs/architecture.md` focused on responsibilities and
    interactions.
-   Keep `docs/philosophy.md` focused on purpose and non-goals.
-   Keep `docs/ui.md` aligned with the current logger and interface.
-   Keep `docs/project-layout.md` synchronized with the actual
    repository tree.
-   Explain important design decisions and their reasons, not just what
    a file contains.
-   Prefer accurate descriptions of the current implementation over
    promises about planned functionality. Where the README describes
    intended behavior, keep that distinction explicit.
-   Preserve the cooking metaphor without sacrificing technical clarity.
