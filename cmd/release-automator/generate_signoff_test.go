package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed testdata/changelog.md
var testChangelog string

func TestGenerateSignOffMessage(t *testing.T) {
	expectedMsg := "# Request for management sign-off: Java SDK Release Version v1.2.0.0\n\nSimon Gabriel Oleksii Ishchenko \n\nBasket v2 is added and logger-core is removed.\n\nRelease will take place provided the Go/NoGo meeting resulted in Go.\n\n# ChangeLog\n\n### Breaking changes\n\n*   Removed `log4j-core` from Maven dependencies. **Please, provide logger\n    implementation on your own**\n\n*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`\n\n*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`\n\n### Deprecated\n\n*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of\n    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you\n    rely on this method, please, replace it with `java.net.URL#URL(String)`\n    constructor call\n\n*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,\n    `amountTotalDiscount`\n\n*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,\n    `amountPerUnit`, `amountNet`\n\n### Added\n\n*   Basket v2 support:\n\n    *   Basket getters/setters: `totalValueGross`\n\n    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`\n\n### Removed\n\n*   Remove `log.error` in catch clause\n    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the\n    exception with exact same message is thrown after the `log.error` call.\n\n### Fixed\n\n*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was\n    not formatted and contained `%s` instead\n    of values\n\n### Changed\n\n*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)\n\n*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)\n\nRelease checklist for Java SDK v1.2.0.0\n"
	msg, err := generateSignoffMessage(
		"Java SDK",
		testChangelog,
		"1.2.0.0",
		[]string{"Simon Gabriel", "Oleksii Ishchenko"},
	)

	assert.NoError(t, err)
	assert.Equal(t, expectedMsg, msg)
}
