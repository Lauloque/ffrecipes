# FFrecipes

[![GitHub license](https://img.shields.io/github/license/Lauloque/ffrecipes?style=for-the-badge&labelColor=rgb(64,64,64))](https://github.com/Lauloque/ffrecipes/blob/master/LICENSE)  [![Development status](https://img.shields.io/badge/status-in%20development-orange?style=for-the-badge&labelColor=rgb(64,64,64))](https://github.com/Lauloque/ffrecipes/issues)  [![ko-fi](https://img.shields.io/badge/Support%20me%20on-KoFi-red?style=for-the-badge&logo=kofi&labelColor=cyan&logoColor=red)](https://ko-fi.com/H2H818FHX)

FFrecipes is a collection of small, reusable FFmpeg recipes for everyday
media-processing tasks. It makes those workflows easy to find, run, and
share.

## Motivation

FFmpeg can do almost anything, but its options are not always easy to
remember or reuse. Small, useful jobs often end up as one-off commands or
scripts scattered across machines, with no clear way to find them later
or share them with someone else.

I built FFrecipes to give those jobs a stable home: one recipe per task,
with a consistent name, documented inputs, and a common way to run it.
That makes a workflow easier to reuse months later, adapt when needs
change, or contribute so others can benefit from it. It is not meant to
replace full-featured graphical encoders such as HandBrake; it focuses on
small, specific jobs and repeatable workflows.

If you keep reaching for FFmpeg to do a specific job, FFrecipes gives you
a place to keep that work small, discoverable, and independent of any
single interface.

## 🚀 Quickstart

You need Go and FFmpeg installed and available on your `PATH`.

### 1. Install FFrecipes

```bash
go install github.com/Lauloque/ffrecipes/cmd/ffrecipes@latest
```

### 2. Cook a recipe

```bash
# Extract every frame from a video as PNG files
ffrecipes cook video-to-png movie.mp4
```

This writes `movie_0001.png`, `movie_0002.png`, and so on next to
`movie.mp4`.

### 3. Explore the CLI

``` sh
ffrecipes list
ffrecipes info video-to-png
ffrecipes --verbose cook video-to-png movie.mp4
```

-   `list` shows the available recipes.
-   `info` describes a recipe.
-   `cook` runs a recipe.
-   `--verbose` enables more detailed output. It must come before the
    command.

Run `ffrecipes help` for the command list. For details on what each
recipe does, see the [recipe catalog](docs/recipes.md).

## 🤝 Contributing

You need [Go](https://go.dev/) and FFmpeg installed to work on FFrecipes.

### Clone the repository

``` sh
git clone https://github.com/Lauloque/ffrecipes.git
cd ffrecipes
```

### Run the project locally

`go run` is the quickest way to try the project or your changes without
building a binary:

``` sh
go run ./cmd/ffrecipes list
go run ./cmd/ffrecipes info video-to-png
go run ./cmd/ffrecipes cook video-to-png /path/to/video.mp4
```

### Build the binary

``` sh
go build -o ffrecipes ./cmd/ffrecipes
./ffrecipes help
```

The built binary is written to the repository root and is ignored by Git.

### Run the checks

Before submitting a change, format the code and run the test suite:

``` sh
gofmt -w .
go build ./...
go vet ./...
go test ./...
```

### Add a recipe

New recipes are welcome. A recipe should:

-   describe itself clearly and do one focused task;
-   receive inputs and runtime resources through the Pantry;
-   use `pantry.FFmpeg` instead of calling external tools directly;
-   accept a context and honor cancellation;
-   report progress through `pantry.Log` and return errors instead of
    terminating the application;
-   register itself with the Cookbook;
-   include tests for input validation and argument building.

Keep recipes independent of the current user interface so they can be
run through other interfaces in the future. See
[Architecture](docs/architecture.md) and the
[recipe catalog](docs/recipes.md) for more detail.

### Submit a pull request

Fork the repository, create a branch, and open a pull request against
`main`. Keep changes focused, and include what you changed and how you
tested it.

Coding agents should also read [AGENTS.md](AGENTS.md).

## License

FFrecipes is distributed under the GNU General Public License v3.0. See
[LICENSE](LICENSE) for the full license text.
