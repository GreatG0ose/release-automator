# Release Automator

Automates release workflow steps

# Usage

## Commands

### Send sign-off message

```shell
release-automator --version="1.2.3.0" --config-path="local_config/java-sdk.yaml" signoff
```

### Send release email

```shell
release-automator --version="1.2.3.0" --config-path="local_config/java-sdk.yaml" mail
```

## Configuration

```yaml
project:
  name: "<Project Name>"
  changelogPath: "CHANGELOG.md"
signOff:
  teamsWebhook: "https://webhook.office.com/webhookb2/<SOME-MAGIC-NUMBERS>"
  mentions:
    - name: "Firstname Lastname"
      teamsID: "firstname.lastname@company.com"
    - name: "Other Person"
      teamsID: "other.person@company.com"
email:
  outlookWebhook: "https://webhook.office.com/webhookb2/<SOME-MAGIC-NUMBERS>"
confluence:
  endpoint: "https://<COMPANY>.atlassian.net/wiki/rest/api"
  releasesPageId: ""
  credentials:
    username: "your@company.com"
    accessToken: "access-token-obtained-from-confluence"
```
