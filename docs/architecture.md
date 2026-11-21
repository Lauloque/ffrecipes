# Architecture

FFrecipes is designed as a set of components. They each have a single responsibility and communicate with the others through well-defined interfaces. This separation makes it easier to test components independently and allows the application to evolve (for example from a CLI to a graphical interface) without changing the recipes themselves.

## Components overview

```mermaid
flowchart TD

Application["🖥️ Application"]

subgraph Kitchen["👨‍🍳 Kitchen"]
    Cook["👨‍🍳 Cook"]
    Cookbook["📖 Cookbook"]
    Recipe["📜 Recipe"]
    Pantry["🥫 Pantry"]

    subgraph Ovens["🔥 Ovens"]
        FFmpeg["🎬 FFmpeg"]
        Gifski["🖼️ Gifski"]
    end
end

Application --> Cook
Application --> Cookbook
Application --> Pantry

Cook --> Recipe
Cookbook --> Recipe
Pantry --> Recipe
Recipe --> FFmpeg
Recipe --> Gifski
```

The Application is the entry point between the user and FFrecipes. It currently uses a CLI, but should later be replaced or complemented by another interface (GUI or TUI).

The Kitchen contains the core cooking logic:

- Cook coordinates execution;
- Cookbook manages available recipes;
- Pantry provides recipes with their required resources;
- Recipes tell how to use the resources provided by the Pantry;
- Ovens perform the actual processing work.

Recipes log through the `kitchen.Logger` interface exposed by the Pantry. `internal/ui` provides the current CLI implementation; future TUI or GUI frontends can provide their own without changing recipe code.

```mermaid
sequenceDiagram
    participant User
    participant App as Application
    participant Cookbook
    participant Cook
    participant Pantry
    participant Recipe
    participant Oven as FFmpeg Oven

    User->>App: cook video-to-png movie.mp4

    App->>Cookbook: Find("video-to-png")
    Cookbook-->>App: Recipe

    App->>Pantry: Create execution context
    App->>Cook: Prepare(recipe, pantry)

    Cook->>Recipe: Cook(Pantry)
    Recipe->>Oven: Execute FFmpeg command
    Oven-->>Recipe: Result

    Recipe-->>Cook: Success / Error
    Cook-->>App: Success / Error
```

## Recipe principles

Every recipe should:

- describe itself;
- perform one focused task;
- be reusable;
- get its inputs through the Pantry;
- communicate to the user through the logger;
- never parse CLI arguments;
- never search for FFmpeg;
- never interact with the user;
- return an error instead of terminating the program;
- register itself in the Cookbook.
