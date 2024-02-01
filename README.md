# Release Automator

Automates release actions:

* Send management sign-off message
* Send release email
* Generate tweet message
* Generate release notes

# Usage

## Commands

Syntax
```shell
release --version=<RELEASE_VERSION> [--config=<CONFIG_PATH>] <COMMAND>
```

### Arguments

| Argument  | Required | Description                                                            |
|-----------|----------|------------------------------------------------------------------------|
| --version | +        | Release version. For example, `--version=1.2.3`                        |
| --config  | -        | Path to release-automator configuration file. Default: `release.yaml`. |

### Send sign-off message

```shell
release --version="1.2.3" --config="java-sdk.yaml" signoff
```

### Send release email

Send release email to Outlook using Webhook API.

```shell
release --version="1.2.3" --config="java-sdk.yaml" mail
```

### Generate release notes

Generate markdown formatted release notes.

```shell
release --version="1.2.3" --config="java-sdk.yaml" notes
```

### Generate twitter text

Generate ready-to-post public release announce. Generated file

```shell
release --version="1.2.3" --config="java-sdk.yaml" tweet
```

## Configuration

See [release-config.yaml](release-config.yaml)

## Docker

Public images are available at [GitHub Registry](https://ghcr.io/greatg0ose/release-automator).

Run 
```shell
docker run -v $(pwd):/app ghcr.io/greatg0ose/release-automator:latest <arguments> <command>
```

Example: sending email using Docker container
```shell
docker run -v $(pwd):/app ghcr.io/greatg0ose/release-automator:latest --version="1.2.3" --config="java-sdk.yaml" signoff
```
