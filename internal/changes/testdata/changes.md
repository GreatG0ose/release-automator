Basket v2 is added and logger-core is removed.

## Added

*   Basket v2 support:
    *   Basket getters/setters: `totalValueGross`
    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`

## Breaking changes

*   Removed `log4j-core` from Maven dependencies. **Please, provide logger
    implementation on your own**
*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`
*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`

## Changed

*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)
*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)

## Deprecated

*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of
    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you
    rely on this method, please, replace it with `java.net.URL#URL(String)`
    constructor call
*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,
    `amountTotalDiscount`
*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,
    `amountPerUnit`, `amountNet`

## Fixed

*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was
    not formatted and contained `%s` instead
    of values

## Removed

*   Remove `log.error` in catch clause
    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the
    exception with exact same message is thrown after the `log.error` call.
