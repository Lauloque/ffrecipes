# Project Layout

- 📂 **ffrecipes**
  - 📄 [README.md](README.md) _General package info_
  - 📂 **cmd**
    - 📂 **ffrecipes**
      - 📄 [main.go](cmd/ffrecipes/main.go) _Program entry point_
  - 📂 **docs**
    - 📄 [architecture.md](docs/architecture.md) _How FFrecipes functions_
    - 📄 [philosophy.md](docs/philosophy.md) _What is FFrecipes supposed to be_
    - 📄 [project\-layout.md](docs/project-layout.md) _Where is what_
  - 📄 [go.mod](go.mod) _Go module definition and dependencies_
  - 📂 **internal**
    - 📂 **kitchen**
      - 📂 **ovens** _Contain all the tools available to do the actual work_
        - 📄 [ffmpeg.go](internal/kitchen/ovens/ffmpeg.go) _Interaction with FFmpeg_
      - 📄 [cook.go](internal/kitchen/cook.go) _TBD: Supposed to execute recipe using provided pantry, but for now that's done from main.go..._
      - 📄 [cookbook.go](internal/kitchen/cookbook.go) _Discovers and registers available recipes_
      - 📄 [options.go](internal/kitchen/options.go) _Runtime options provided from the application_
      - 📄 [pantry.go](internal/kitchen/pantry.go) _Runtime resources available to recipes (inputs, logger, FFmpeg path, options, ...)_
      - 📄 [recipe.go](internal/kitchen/recipe.go) _Defines what every recipe must implement_
    - 📂 **ui**
      - 📄 [logger.go](internal/ui/logger.go) _Logging abstraction used by recipes and application interfaces_
  - 📂 **recipes** _Recipe implementations_
    - 📄 [video\_to\_png.go](recipes/video_to_png.go)
  - 📄 [test.mp4](test.mp4) _Test file for convenience_

_Tree generated using [GitHub/michalbe/md-file-tree](https://github.com/michalbe/md-file-tree)_
