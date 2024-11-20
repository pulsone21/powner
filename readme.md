# POWNER

lets product own our teams and keep track on skillsets from thier members.
This is a side project, which i develop to make my live easier as an product owner.

## Starting the Application

### ENV Variables

There are a few env variables that you might want to set.

```.env
PROTOCOL=http
URL=localhost
PORT=5550
DB=./db/test.db # where you SQLite DB file is
```

Also there are some optional variables which you can set:

#### Logging

```.env
LOGGER_LEVEL=debug # default is info (Options: debug, error, warning)
LOGGER_OUTPUT=file # default is Stdout (Options: File)
LOGGER_SOURCE=true # default is false (Options: true, false) - Enables Slog Source logging
LOGGER_TYPE=TEXT   # default is text (Options: TEXT, JSON) - Which Output format you want
```

### Running the Application

by side the normal stuff like `go run ....` you can just use

```bash
just dev
```

to start the application
or

```bash
just watch_dev
```

that it is in a watch mode for any tailwind or templ changes.
