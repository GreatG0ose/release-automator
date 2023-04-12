package release_notes

import (
	"github.com/GreatG0ose/release-automator/internal/changelog"
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/GreatG0ose/release-automator/internal/release"
	"github.com/GreatG0ose/release-automator/internal/test_utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateVersionChangesFile(t *testing.T) {
	r := release.Release{
		Version: "1.2.0.0",
		Changes: changelog.ReleaseChanges{
			Summary: "Basket v2 is added and logger-core is removed.",
			Changes: map[string]string{
				"Breaking changes": "*   Removed `log4j-core` from Maven dependencies. **Please, provide logger\n    implementation on your own**\n*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`\n*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`",
				"Deprecated":       "*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of\n    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you\n    rely on this method, please, replace it with `java.net.URL#URL(String)`\n    constructor call\n*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,\n    `amountTotalDiscount`\n*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,\n    `amountPerUnit`, `amountNet`",
				"Added":            "*   Basket v2 support:\n    *   Basket getters/setters: `totalValueGross`\n    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`",
				"Removed":          "*   Remove `log.error` in catch clause\n    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the\n    exception with exact same message is thrown after the `log.error` call.",
				"Changed":          "*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)\n*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)",
				"Fixed":            "*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was\n    not formatted and contained `%s` instead\n    of values",
			},
		},
	}

	tmpdir, err := os.MkdirTemp("", "")
	require.NoError(t, err)

	cfg := config.Config{
		Project: config.Project{
			Name:          "Java SDK",
			ChangelogPath: "testdata/changelog.md",
		},
		Output: tmpdir,
	}
	assert.NoError(t, err)

	err = Generate(zerolog.Nop(), cfg, r)
	assert.NoError(t, err)

	test_utils.FilesEqual(t, filepath.Join("testdata", "changes.md"), filepath.Join(tmpdir, "ReleaseNotes.md"))
}
