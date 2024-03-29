package changelog

import (
	"github.com/GreatG0ose/release-automator/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractReleaseNotes(t *testing.T) {
	testCases := []struct {
		name     string
		cfg      config.Config
		version  string
		expected ReleaseChanges
	}{
		{
			name:    "full release notes",
			version: "1.2.0.0",
			cfg: config.Config{
				Project: config.Project{
					Name:          "Java SDK",
					ChangelogPath: "testdata/changelog.md",
				},
			},
			expected: ReleaseChanges{
				Summary: "Basket v2 is added and logger-core is removed.",
				Changes: []ChangeBlock{
					{"Added", "*   Basket v2 support:\n\n    *   Basket getters/setters: `totalValueGross`\n\n    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`"},
					{"Breaking changes", "*   Removed `log4j-core` from Maven dependencies. **Please, provide logger\n    implementation on your own**\n\n*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`\n\n*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`"},
					{"Changed", "*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)\n\n*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)"},
					{"Deprecated", "*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of\n    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you\n    rely on this method, please, replace it with `java.net.URL#URL(String)`\n    constructor call\n\n*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,\n    `amountTotalDiscount`\n\n*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,\n    `amountPerUnit`, `amountNet`"},
					{"Fixed", "*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was\n    not formatted and contained `%s` instead\n    of values"},
					{"Removed", "*   Remove `log.error` in catch clause\n    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the\n    exception with exact same message is thrown after the `log.error` call."},
				},
			},
		},

		{
			name:    "no release summary",
			version: "1.2.1.0",
			cfg: config.Config{
				Project: config.Project{
					Name:          "Java SDK",
					ChangelogPath: "testdata/no-summary.md",
				},
			},
			expected: ReleaseChanges{
				Changes: []ChangeBlock{
					{"Added", "* Add payment type Paylater Invoice. See more at [Unzer Docs](https://docs.unzer.com/payment-methods/unzer-invoice-upl/)"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ExtractReleaseChanges(tc.cfg, tc.version)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
