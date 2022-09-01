# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres
to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [1.2.1.0][1.2.1.0]

### Added

* Add payment type Paylater Invoice. See more at [Unzer Docs](https://docs.unzer.com/payment-methods/unzer-invoice-upl/)

## [1.2.0.0][1.2.0.0]

Basket v2 is added and logger-core is removed.

### Breaking changes

*   Removed `log4j-core` from Maven dependencies. **Please, provide logger
    implementation on your own**

*   Renamed enum value `AbstractTransaction.Status.ERRROR` to `AbstractTransaction.Status.ERROR`

*   Renamed enum value `Paypage.Status.ERRROR` to `Paypage.Status.ERROR`

### Deprecated

*   `com.unzer.payment.service.UrlUtil#getUrl(String)` will not be part of
    java-sdk, because it has nothing to do with unzer/sdk specific logic. If you
    rely on this method, please, replace it with `java.net.URL#URL(String)`
    constructor call

*   Deprecated Basket getters/setters: `amountTotalVat`, `amountTotalGross`,
    `amountTotalDiscount`

*   Deprecated BasketItem getters/setters: `amountDiscount`, `amountGross`, `amountVat`,
    `amountPerUnit`, `amountNet`

### Added

*   Basket v2 support:

    *   Basket getters/setters: `totalValueGross`

    *   BasketItem getters/setters: `amountPerUnitGross`, `amountDiscountPerUnitGross`

### Removed

*   Remove `log.error` in catch clause
    `com.unzer.payment.service.PropertiesUtil#loadProperties()` because the
    exception with exact same message is thrown after the `log.error` call.

### Fixed

*   Fix log message of `com.unzer.payment.service.UrlUtil#getUrl(String)`. It was
    not formatted and contained `%s` instead
    of values

### Changed

*   Type of `BasketItem` type field (was `String`, become `BasketItem.Type`)

*   Type of `BasketItem` `vat` field (was `Integer`, become `BigDecimal`)

## [1.1.2.7][1.1.2.7]

### Changed

*   Upgrade of the used gson Dependencies from 2.8.6 to 2.8.9
*   Upgrade of the used faster-xml dependencies from 2.11.3 to 2.12.7
*   Upgrade of the used log4j dependencies from 2.17.1 to 2.18.0
