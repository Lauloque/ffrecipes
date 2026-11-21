# Recipes

This catalog describes what each recipe does, independently of the interface
used to run it. CLI examples are included where they help, but the behavior
below applies to any frontend.

The catalog is intentionally small while FFrecipes is in development.

## video-to-png

Extract every frame from one or more video files as PNG images.

For each input file, the recipe writes a numbered image sequence next to the
input:

```text
<input>_%04d.png
```

For example, `movie.mp4` produces `movie_0001.png`, `movie_0002.png`, and
so on.

### Inputs

- One or more video files.

### Output

- One PNG image per decoded frame, named with a zero-padded four-digit
  sequence number.
- Existing files matching the output pattern are overwritten.

### Behavior

- Uses FFmpeg with `-fps_mode passthrough`, so frames are not duplicated or
  dropped to fit a fixed frame rate.
- Processes inputs in the order they are given.
- Stops on the first input that fails.

### Requirements

- FFmpeg must be available to the application.

### CLI example

``` sh
ffrecipes cook video-to-png movie.mp4
```
