# Project Layout

- 📂 **ffrecipes**
  - 📄 [README.md](README.md) _General package info_
  - 📂 **cmd**
    - 📂 **ffrecipes**
      - 📄 [main.go](cmd/ffrecipes/main.go) _CLI entry point_
  - 📂 **docs**
    - 📄 [architecture.md](docs/architecture.md) _How FFrecipes functions_
    - 📄 [philosophy.md](docs/philosophy.md) _What is FFrecipes supposed to be_
    - 📄 [project\-layout.md](docs/project-layout.md) _Where is what_
    - 📄 [recipes.md](docs/recipes.md) _Catalog of available recipes_
    - 📄 [ui.md](docs/ui.md) _How the user interface is layered_
  - 📄 [go.mod](go.mod) _Go module definition and dependencies_
  - 📂 **internal**
    - 📂 **app**
      - 📄 [app.go](internal/app/app.go) _Application orchestration shared by frontends_
    - 📂 **files**
      - 📄 [files.go](internal/files/files.go) _Shared file existence checks for recipes_
      - 📄 [files_test.go](internal/files/files_test.go) _Tests for file checks_
    - 📂 **kitchen**
      - 📂 **ovens** _Contain all the tools available to do the actual work_
        - 📄 [ffmpeg.go](internal/kitchen/ovens/ffmpeg.go) _FFmpeg implementation of the runner interface_
      - 📄 [cookbook.go](internal/kitchen/cookbook.go) _Discovers and registers available recipes_
      - 📄 [logger.go](internal/kitchen/logger.go) _Defines the logging interface available to recipes_
      - 📄 [pantry.go](internal/kitchen/pantry.go) _Runtime resources available to recipes (inputs, logger, FFmpeg runner, ...)_
      - 📄 [recipe.go](internal/kitchen/recipe.go) _Defines what every recipe must implement_
      - 📄 [runner.go](internal/kitchen/runner.go) _Defines the external command runner interface_
    - 📂 **ui**
      - 📄 [logger.go](internal/ui/logger.go) _CLI implementation of the kitchen logging interface_
  - 📂 **recipes** _Recipe implementations_
    - 📄 [add\_cover\_to\_video.go](recipes/add_cover_to_video.go)
    - 📄 [add\_cover\_to\_video\_test.go](recipes/add_cover_to_video_test.go) _Tests for the add-cover-to-video recipe_
    - 📄 [video\_to\_png.go](recipes/video_to_png.go)
    - 📄 [video\_to\_png\_test.go](recipes/video_to_png_test.go) _Tests for the video-to-png recipe_
  - 📄 [test.mp4](test.mp4) _Local test file (ignored by Git)_

_Tree generated using [GitHub/michalbe/md-file-tree](https://github.com/michalbe/md-file-tree)_
