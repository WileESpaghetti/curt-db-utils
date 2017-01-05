# Curt Labs database utils
## Not ready for production

## Features

### API boilerplate

#### Database initialization and session creation
Database initialization and session creation can
now be reused instead of copy/pasted across
multiple projects. (the database helper snippet)

#### Common API models
Structs that are used throughout Curt APIs are
now bundled in this library so that we avoid
duplication across the various APIs.

A common database access layer has also been
created so that we can get the data for the
models in a consistant way reguardless of which
database we use.

### Utility functions
* database connection builders
    * Create MySQL and MongoDB connection strings
    easily from environment variables or hard-coded
    values
