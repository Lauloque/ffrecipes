# FFrecipes

**A cookbook of reusable FFmpeg recipes.**

FFrecipes is a collection of small, focused media-processing workflows
for tasks that are useful in real projects but too specific to fit
comfortably into a general-purpose encoder.

Sometimes you need to extract a PNG image sequence from a video. Other
times you need to stack videos from different folders, correct rotation
metadata, burn frame numbers with a particular offset, or handle a
one-off conversion that turns out to be useful more than once.

FFrecipes aims to make those workflows easy to find, run, reuse, and
contribute.

## What it is (and isn't)

FFrecipes is:

-   a collection of focused FFmpeg workflows;
-   a way to discover what each workflow does and what it needs;
-   a common structure for writing and sharing additional recipes;
-   intended to work across supported operating systems without
    requiring users to manage a collection of separate scripts and their
    runtimes.

FFrecipes is **not** intended to be another full-featured graphical
media encoder. Tools such as HandBrake already serve that purpose.
FFrecipes focuses on small, specific jobs and repeatable workflows.

## The cooking metaphor

The project uses cooking terminology to describe its components:

  Term          Responsibility
  ------------- -----------------------------------------------------------
  Application   The user-facing entry point. Currently a CLI.
  Kitchen       The core execution package.
  Cookbook      Registers and finds available recipes.
  Recipe        Performs one focused media-processing task.
  Pantry        Provides recipes with their inputs and runtime resources.
  Oven          An external processing tool used by a recipe.

FFmpeg is the primary oven. Other tools may be integrated when a
particular workflow genuinely needs them.

For more detail, see [docs/architecture.md](docs/architecture.md).

## Current status

FFrecipes is under active development and its architecture and
command-line interface may change. The initial work is focused on
establishing the recipe system and getting the first recipes working end
to end.

The `video-to-png` recipe is the initial working example. It extracts
every frame as `<input>_%04d.png` next to the input file. Check the
available commands and recipe information in the application itself for
the current behaviour.

## Usage

The current command structure is:

``` sh
ffrecipes [options] <command> [arguments]
```

Examples:

``` sh
ffrecipes list
ffrecipes info video-to-png
ffrecipes --verbose cook video-to-png movie.mp4
```

-   `list` displays the available recipes.
-   `info` displays information about a recipe.
-   `cook` runs a recipe. `video-to-png` writes an image sequence next to
    each input file.
-   `--verbose` enables more detailed output.

These examples describe the current prototype interface; run
`ffrecipes help` for the command list.

## Requirements

-   Go, to build the application from source.
-   FFmpeg, for recipes that use the FFmpeg oven.

Additional external tools may be required by particular recipes. Their
requirements should be documented by those recipes. FFrecipes does not
imply that every optional tool is required for every task.

## Build from source

Clone the repository and build the application:

``` sh
git clone https://github.com/Lauloque/ffrecipes.git
cd ffrecipes
go build -o ffrecipes ./cmd/ffrecipes
```

You can then run it from the repository directory:

``` sh
./ffrecipes list
```

## Contributing

New recipes are welcome. Before adding one, take a look at the existing
implementation and the project architecture.

As a general rule, a recipe should:

-   describe itself clearly;
-   perform one focused task;
-   be reusable;
-   receive inputs and runtime resources through the Pantry;
-   avoid parsing command-line arguments;
-   use the runner provided by the Pantry instead of calling external
    tools directly;
-   honor context cancellation for long-running work;
-   report progress through the project's logging abstraction;
-   return errors to the caller rather than terminating the application;
-   register itself with the Cookbook, following the existing
    convention.

Keep recipes independent of the current user interface so they can be
run through other interfaces in the future.

See:

-   [Architecture](docs/architecture.md)
-   [Project layout](docs/project-layout.md)
-   [AGENTS.md](AGENTS.md), for coding-agent and contributor guidance

## License

FFrecipes is distributed under the GNU General Public License v3.0. See
[LICENSE](LICENSE) for the full license text.
