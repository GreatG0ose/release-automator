package ms_teams

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderSignOffMessage(t *testing.T) {
	expectedText := `# Request for management sign-off: Java-SDK Release Version v1.2.0.0

Oleksii Ishchenko Rene Salecker Tess Akerlund Timo Seifert Simon Gabriel 

This release introduces a brilliant feature A to replace feature C

Release will take place provided the Go/NoGo meeting resulted in Go.

# ChangeLog

## Deprecated

* Old feature B
* Another old feature C
	* with subfeature 1
	* what use instead

## Added

* New feature A
* Another new feature B

## Fixed

* Fixed old bug
* Fixed not so old bug

Release checklist for Java SDK v1.2.0.0
`
	testRelease := release.Release{
		ProjectName: "Java-SDK",
		Version:     "1.2.0.0",
		Changelog: changelog.ReleaseChangelog{
			Summary: "This release introduces a brilliant feature A to replace feature C",
			Changes: `## Deprecated

* Old feature B
* Another old feature C
	* with subfeature 1
	* what use instead

## Added

* New feature A
* Another new feature B

## Fixed

* Fixed old bug
* Fixed not so old bug`,
		},
	}
	mentions := []string{"Oleksii Ishchenko", "Rene Salecker", "Tess Akerlund", "Timo Seifert", "Simon Gabriel"}

	actual, err := RenderSignOffMessage(testRelease, mentions)
	assert.NoError(t, err)
	assert.Equal(t, expectedText, actual)
}
