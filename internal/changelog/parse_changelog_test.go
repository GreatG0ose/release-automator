package changelog

import (
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractReleaseNotes(t *testing.T) {
	expectedReleaseChangeLog := ReleaseChanges{
		Summary: "Basket v2 is added and logger-core is removed.",
		Changes: map[string]string{
			"Breaking changes": "*   Removed `log4j-core` from Maven dependencies. **Please, provide logger\n    implementation on your own**\n\n*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`\n\n*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`",
			"Deprecated":       "*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of\n    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you\n    rely on this method, please, replace it with `java.net.URL#URL(String)`\n    constructor call\n\n*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,\n    `amountTotalDiscount`\n\n*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,\n    `amountPerUnit`, `amountNet`",
			"Added":            "*   Basket v2 support:\n\n    *   Basket getters/setters: `totalValueGross`\n\n    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`",
			"Removed":          "*   Remove `log.error` in catch clause\n    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the\n    exception with exact same message is thrown after the `log.error` call.",
			"Changed":          "*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)\n\n*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)",
			"Fixed":            "*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was\n    not formatted and contained `%s` instead\n    of values",
		},
	}

	cfg := config.Config{
		Project: config.Project{
			Name:          "Java SDK",
			ChangelogPath: "testdata/changelog.md",
		},
	}

	actualReleaseChangelog, err := ExtractReleaseChanges(cfg, "1.2.0.0")

	assert.NoError(t, err)

	assert.Equal(t, expectedReleaseChangeLog.Summary, actualReleaseChangelog.Summary)
	assert.Equal(t, expectedReleaseChangeLog.Changes, actualReleaseChangelog.Changes)
}
