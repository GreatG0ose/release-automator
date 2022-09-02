package config

// Config is used by release-automator tools
type Config struct {
	Project          Project          `yaml:"project"` // Project is meta information of the target project
	SignOff          SignOff          `yaml:"signOff"` // SignOff used for sending sign-off messages to MS Teams
	FullReleaseEmail FullReleaseEmail `yaml:"email"`   // FullReleaseEmail configures full-release message sending
}

// Project meta information and path to changelog file
type Project struct {
	Name          string `yaml:"name"`          // Name of the project
	ChangelogPath string `yaml:"changelogPath"` // ChangelogPath is path to changelog.md file of the project
}

// SignOff contains Webhook URL to connector app and list of people to mention
type SignOff struct {
	TeamsWebhook string    `yaml:"teamsWebhook"` // TeamsWebhook can be obtained from Teams Connector. See more: https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook
	Mentions     []Mention `yaml:"mentions"`     // Mentions is list of contacts to mention in Signoff message
}

// Mention is used for mentioning a person in MS Teams message
type Mention struct {
	Name    string `yaml:"name"`    // Name is visible name of a mentioned person
	TeamsID string `yaml:"teamsID"` // TeamsID usually is Microsoft email
}

// FullReleaseEmail configures full-release message sending. It consists of webhook url.
type FullReleaseEmail struct {
	OutlookWebhook string `yaml:"outlookWebhook"` // OutlookWebhook used for sending message to Outlook connector. See more how to obtain webhook: https://docs.microsoft.com/en-us/outlook/actionable-messages/send-via-connectors
}
