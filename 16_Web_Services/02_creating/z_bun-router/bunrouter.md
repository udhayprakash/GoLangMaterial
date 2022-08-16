## BunRouter

    - extremely fast HTTP router for Go with unique combination of features:

    - Middlewares
        - allow to extract common operations from HTTP handlers into reusable functions.
    - Error handling
        - allows to further reduce the size of HTTP handlers by handling errors in middlewares.
    - Routes priority
        - enables meaningful matching priority for routing rules: first static nodes, then named nodes, lastly wildcard nodes.
    - net/http compatible API
        - which means using minimal API without constructing huge wrappers that try to do everything: from serving static files to XML generation (for example, gin.Context or echo.Context).

### Installation

    go get github.com/uptrace/bunrouter
