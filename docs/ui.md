# User Interface

For now the UI is a basic CLI logger:

```mermaid
flowchart TD
    A[Recipe] --> B[Logger]
    B --> C[Terminal]
    B -->  |Future| D(GUI or TUI)
```
