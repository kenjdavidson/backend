# Stream Pets Backend

## Contribution

### Local Development

1. Clone the project repository.  

If you wish rename development folder `backend` to `streampets-backend`:

```
/workspaces $ git clone git@github.com:StreamPets/backend.git streampets-backend
/workspaces $ cd streampets-backend
```

2. Open project

Within your IDE of choice, open the project.  Local development is usually performed in VSCode; with the recent addition of Dev Containers.  Working within dev containers the standard `.devcontainer/.env` should be defaulted to:

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_HOSTNAME=localhost
```

3. Setup local configuration:

Setup `.env` file with contents:

```
# Configure PostgreSQL environment
DB_HOST=localhost
DB_PORT=5432
DB_SSL_MODE=disable
DB_NAME=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres

# Fill with your Stream Pets developer (test) application
CLIENT_ID=<Titch Client Id>
CLIENT_SECRET=<Twitch Client Secrete>
```

4. Complete Feature/Fix and open Pull Request

Pull requests should contain:

- appropriate tests and documentation (where needed)
- appropriate commit message(s), noting whether merge should be squashed or not

### Running Tests

```shell
$ go test ./test/**
```

Examples:

```
vscode ➜ /workspaces/streampets-backend (kenjdavidson/issue-12) $ go test ./test/**
ok      github.com/streampets/backend/test/controllers  (cached)
ok      github.com/streampets/backend/test/repositories 0.873s
ok      github.com/streampets/backend/test/services     (cached)
```

### Running Server
