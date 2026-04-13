# Changelog

## 0.2.5 (2026-04-13)

Full Changelog: [v0.2.4...v0.2.5](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.2.4...v0.2.5)

## 0.2.4 (2026-04-13)

Full Changelog: [v0.2.3...v0.2.4](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.2.3...v0.2.4)

### Chores

* remove old list_attachments. ([93d1c40](https://github.com/MercuryTechnologies/mercury-cli/commit/93d1c405ab56bf79459cf799cd55932ab9cbf473))


### Refactors

* consolidate money movement into single payments resource ([a7b610a](https://github.com/MercuryTechnologies/mercury-cli/commit/a7b610a8ad2f2f73720f0d25e30f41951ea5e54e))
* make attachments a sub resource ([e9b8414](https://github.com/MercuryTechnologies/mercury-cli/commit/e9b8414564623d4e802654ee4ef281204851a1fb))
* move ar attachments as a sub resrource ([a667236](https://github.com/MercuryTechnologies/mercury-cli/commit/a667236dfeb3aadf073f1c18f1da6a0b0015c8d2))
* rename upload_attachment -&gt; attach ([503dc81](https://github.com/MercuryTechnologies/mercury-cli/commit/503dc81d79f996b7f6cf78d3def0ce5c8cfb1209))
* unify statements into single resource with account/treasury subresources ([3ed5a0b](https://github.com/MercuryTechnologies/mercury-cli/commit/3ed5a0bc0b178b8b75f1a14a49fb0878f70a1228))
* unify transactions listing into single resource ([6ebff83](https://github.com/MercuryTechnologies/mercury-cli/commit/6ebff83fef8adffac20706406bd6ceaed6893589))

## 0.2.3 (2026-04-10)

Full Changelog: [v0.2.2...v0.2.3](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.2.2...v0.2.3)

## 0.2.2 (2026-04-10)

Full Changelog: [v0.2.1...v0.2.2](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.2.1...v0.2.2)

## 0.2.1 (2026-04-10)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.2.0...v0.2.1)

### Bug Fixes

* **cli:** fix incompatible Go types for flag generated as array of maps ([cadc5dd](https://github.com/MercuryTechnologies/mercury-cli/commit/cadc5dd76c3e9685912e9f2ce2b2597840e9029e))
* fix for failing to drop invalid module replace in link script ([ba1a25e](https://github.com/MercuryTechnologies/mercury-cli/commit/ba1a25e1955da4a575e37dc351a9ba3e46e08f45))

## 0.2.0 (2026-04-10)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.1.1...v0.2.0)

### ⚠ BREAKING CHANGES

* remove account.get_transaction

### Features

* remove account.get_transaction ([c6bc7c3](https://github.com/MercuryTechnologies/mercury-cli/commit/c6bc7c3e4a4dcd55144c028b5622a33ce7b54c90))


### Bug Fixes

* get events typo ([24cd315](https://github.com/MercuryTechnologies/mercury-cli/commit/24cd3157982b05b3f545e73810c3407db51469b9))


### Chores

* **cli:** additional test cases for `ShowJSONIterator` ([777e10c](https://github.com/MercuryTechnologies/mercury-cli/commit/777e10cad687ff4e35bfffcba7631436791b310d))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([d835896](https://github.com/MercuryTechnologies/mercury-cli/commit/d8358961bd2309d523d77f4d8837aecc5f2f297c))
* fix attachments typo ([8629d7d](https://github.com/MercuryTechnologies/mercury-cli/commit/8629d7dbba404424aa47af2ac4ebae93c232b109))
* **internal:** codegen related update ([97c5744](https://github.com/MercuryTechnologies/mercury-cli/commit/97c5744bc85b6910421ce86a044c1a73d131646a))


### Documentation

* remove fake api token example ([d0bbe9d](https://github.com/MercuryTechnologies/mercury-cli/commit/d0bbe9d4419a864bf4f5b13a6ca429f47eb66fe8))


### Styles

* rename "retrieve" to "get" ([570eaea](https://github.com/MercuryTechnologies/mercury-cli/commit/570eaea86443294b09e2a871f023abb1879656a6))
* rename `organization` to `org` ([e236ed8](https://github.com/MercuryTechnologies/mercury-cli/commit/e236ed8eb66f42055e3407ef3cb7d99d016ed020))
* use `download` as method ([5d4398a](https://github.com/MercuryTechnologies/mercury-cli/commit/5d4398a5696c2a2f142a341bbe85339a8ab36060))


### Refactors

* create cards resource ([3f908bd](https://github.com/MercuryTechnologies/mercury-cli/commit/3f908bdadb8048c99f9eb9dedaf82c68fd569b67))
* make `customers` top level resource ([c9fecfc](https://github.com/MercuryTechnologies/mercury-cli/commit/c9fecfc9e6404e63324038f284523245b5b09946))
* make `invoices` top level resource ([027f3b4](https://github.com/MercuryTechnologies/mercury-cli/commit/027f3b45f6fbe557869f5d5ab262e7d58dc5f61e))

## 0.1.1 (2026-04-08)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.1.0...v0.1.1)

### Bug Fixes

* remove basicAuth ([ff77abc](https://github.com/MercuryTechnologies/mercury-cli/commit/ff77abc7b05dcec7fd1dae8891c8fe06c4cc1cbc))

## 0.1.0 (2026-04-08)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/MercuryTechnologies/mercury-cli/compare/v0.0.1...v0.1.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([9a328e6](https://github.com/MercuryTechnologies/mercury-cli/commit/9a328e6b1e133353e8fe6a4518ab54c541def136))
* add default description for enum CLI flags without an explicit description ([9ef297b](https://github.com/MercuryTechnologies/mercury-cli/commit/9ef297ba79693f9abb3437382cb8892c57db3587))
* add support for file downloads from binary response endpoints ([00f7fa0](https://github.com/MercuryTechnologies/mercury-cli/commit/00f7fa0930990f87f3ead2bd0d0e9703d169f8d2))
* allow `-` as value representing stdin to binary-only file parameters in CLIs ([18237f7](https://github.com/MercuryTechnologies/mercury-cli/commit/18237f7ba547f392bbe81965b31dcc0fed751978))
* **api:** cursor_id_request_send_money ([858e7b4](https://github.com/MercuryTechnologies/mercury-cli/commit/858e7b4e256aa0c48c42d607ba64f8f012d2f79b))
* **api:** manual updates ([ec941a1](https://github.com/MercuryTechnologies/mercury-cli/commit/ec941a1aee318337e0619af7490e452447f1b8ab))
* **api:** manual updates ([b6dcca0](https://github.com/MercuryTechnologies/mercury-cli/commit/b6dcca03662331f2c587996da3fb93c87974d646))
* **api:** manual updates ([dc4bf65](https://github.com/MercuryTechnologies/mercury-cli/commit/dc4bf6501a03f52e0b48363ac8bc1d9948fd39fe))
* **api:** manual updates ([ec5c2c5](https://github.com/MercuryTechnologies/mercury-cli/commit/ec5c2c54a2e33aa9a9be7db1501ea258b2ff36f9))
* **api:** manual updates ([3ae5336](https://github.com/MercuryTechnologies/mercury-cli/commit/3ae533632df242cf7cd922929e1c18ce11b020b1))
* **api:** manual updates ([386aaa9](https://github.com/MercuryTechnologies/mercury-cli/commit/386aaa9fb0c673756e2c5873df2b05c4ef37cbcb))
* **api:** manual updates ([cc6d089](https://github.com/MercuryTechnologies/mercury-cli/commit/cc6d0898d02bf3dce7ada356003c4eb2e1bda991))
* **api:** manual updates ([3830d03](https://github.com/MercuryTechnologies/mercury-cli/commit/3830d03ae30c4998925a70f8a3afdcb07c12b1b1))
* **api:** manual updates ([fdf9d35](https://github.com/MercuryTechnologies/mercury-cli/commit/fdf9d35f14b6b8504ecace2b7f10b20d6840c3a9))
* **api:** manual updates ([ce4b5eb](https://github.com/MercuryTechnologies/mercury-cli/commit/ce4b5eb702c1dd23ab6a5805ad5ef82ee61057a3))
* **api:** rename list_send_money_requests ([16a6bc1](https://github.com/MercuryTechnologies/mercury-cli/commit/16a6bc1dc1d4614e4aff6bf2d44a07a97de017f8))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([4e5762d](https://github.com/MercuryTechnologies/mercury-cli/commit/4e5762d33798461605c30476e84908a56471f3c2))
* binary-only parameters become CLI flags that take filenames only ([71d6303](https://github.com/MercuryTechnologies/mercury-cli/commit/71d63032afe5e1c283959ca8c97b69b99f16457d))
* improved documentation and flags for client options ([d0e9f2e](https://github.com/MercuryTechnologies/mercury-cli/commit/d0e9f2e65f88e87ee1018cc153317edd1b825b5a))
* set CLI flag constant values automatically where `x-stainless-const` is set ([5d32cf5](https://github.com/MercuryTechnologies/mercury-cli/commit/5d32cf5d76885aba6c91e12c837faa39ac1e07d5))
* support passing required body params through pipes ([12c290d](https://github.com/MercuryTechnologies/mercury-cli/commit/12c290d239763618c9fcbd0cb6e3f53cadda0280))


### Bug Fixes

* add missing example parameters for test cases ([5552400](https://github.com/MercuryTechnologies/mercury-cli/commit/55524007383f17e5fbbe6b8d7786793764bee279))
* avoid printing usage errors twice ([83cb7ef](https://github.com/MercuryTechnologies/mercury-cli/commit/83cb7efd73b22b4db34ca6e2e9420784f4660ccb))
* avoid reading from stdin unless request body is form encoded or json ([060b66f](https://github.com/MercuryTechnologies/mercury-cli/commit/060b66f96f9e5b346a812b1a23f9d2e1a7089d3c))
* better support passing client args in any position ([fc33727](https://github.com/MercuryTechnologies/mercury-cli/commit/fc33727ce97f39783d670705320d500c65076f55))
* cli no longer hangs when stdin is attached to a pipe with empty input ([c10c8b0](https://github.com/MercuryTechnologies/mercury-cli/commit/c10c8b08b41d932ded99feed2b26583266db16d0))
* fall back to main branch if linking fails in CI ([2542ae1](https://github.com/MercuryTechnologies/mercury-cli/commit/2542ae173a1a78f35bb8a08bc4c27b0f80e310be))
* fix for encoding arrays with `any` type items ([7feca68](https://github.com/MercuryTechnologies/mercury-cli/commit/7feca68dd866eed942f608b8430b1816a2d1293a))
* fix for off-by-one error in pagination logic ([4c048fd](https://github.com/MercuryTechnologies/mercury-cli/commit/4c048fdfafb360c1c2ed986024c2ea136a51d897))
* fix for test cases with newlines in YAML and better error reporting ([988658c](https://github.com/MercuryTechnologies/mercury-cli/commit/988658c0ee850d9512ba578fb0bb712dd37ecb6d))
* fix quoting typo ([98119c0](https://github.com/MercuryTechnologies/mercury-cli/commit/98119c0dd549856e17fddfbf3d3f4d9fd9c133a2))
* handle empty data set using `--format explore` ([5decbc5](https://github.com/MercuryTechnologies/mercury-cli/commit/5decbc5cef1e1e9f97fb5133de550015a2109452))
* improve linking behavior when developing on a branch not in the Go SDK ([22c2566](https://github.com/MercuryTechnologies/mercury-cli/commit/22c2566c19e3ffe0ddb36bb6f2a33757df9b876d))
* improved workflow for developing on branches ([abab4d1](https://github.com/MercuryTechnologies/mercury-cli/commit/abab4d1a14b36a027235f7a8fcd13efbf062ba11))
* no longer require an API key when building on production repos ([4e1785f](https://github.com/MercuryTechnologies/mercury-cli/commit/4e1785facf50fe7fd4b665e696ccd31e9627e3ba))
* only set client options when the corresponding CLI flag or env var is explicitly set ([3d3daec](https://github.com/MercuryTechnologies/mercury-cli/commit/3d3daec0b163e2378b4ba06291b7a6a35cea8386))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([1d0a1b4](https://github.com/MercuryTechnologies/mercury-cli/commit/1d0a1b4cb5e3a4f03421dae8af9b80f884fc2b12))


### Chores

* **ci:** skip lint on metadata-only changes ([99e6f7e](https://github.com/MercuryTechnologies/mercury-cli/commit/99e6f7e30d2fdd511c5072deb2cae3640e0f7478))
* **ci:** skip uploading artifacts on stainless-internal branches ([a49176b](https://github.com/MercuryTechnologies/mercury-cli/commit/a49176b779b737d3c2bca8d1970cb9437959b7f2))
* configure new SDK language ([9948c69](https://github.com/MercuryTechnologies/mercury-cli/commit/9948c699d53fac564415e791da1c134c53180f56))
* configure new SDK language ([423139c](https://github.com/MercuryTechnologies/mercury-cli/commit/423139cc9a7b7af0f354cb3ce46d504a806b19dc))
* **internal:** codegen related update ([faf6f32](https://github.com/MercuryTechnologies/mercury-cli/commit/faf6f32cfcee8083b8604058af9e4678c81c6c9d))
* **internal:** codegen related update ([8a2d4be](https://github.com/MercuryTechnologies/mercury-cli/commit/8a2d4be91c6947b22a27b4e2923c06a3b4aa335f))
* **internal:** codegen related update ([87660e7](https://github.com/MercuryTechnologies/mercury-cli/commit/87660e7072b54463437185f6efa714856ef74b6f))
* **internal:** codegen related update ([0f01ae6](https://github.com/MercuryTechnologies/mercury-cli/commit/0f01ae676b408e264c4525a4b539518a2b573575))
* **internal:** codegen related update ([f1cc7de](https://github.com/MercuryTechnologies/mercury-cli/commit/f1cc7de3beba48878c834cbe85d0394d4d54795d))
* **internal:** tweak CI branches ([0e68be9](https://github.com/MercuryTechnologies/mercury-cli/commit/0e68be99b4d668c02f714f9e9aea579ebef01754))
* **internal:** update gitignore ([3b8ccf2](https://github.com/MercuryTechnologies/mercury-cli/commit/3b8ccf2fdbbc97472ad8e9921ef435de36b26879))
* mark all CLI-related tests in Go with `t.Parallel()` ([4df9734](https://github.com/MercuryTechnologies/mercury-cli/commit/4df97340aa4044fd77be706008cb6784a0bba4a8))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([48ace8c](https://github.com/MercuryTechnologies/mercury-cli/commit/48ace8c8c8975b4d3da885dac2c1310a83f29ac3))
* omit full usage information when missing required CLI parameters ([205a9c0](https://github.com/MercuryTechnologies/mercury-cli/commit/205a9c03f74fa0aa3fd824f17ac3b51f719f38a9))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([914f53c](https://github.com/MercuryTechnologies/mercury-cli/commit/914f53c3509091a37a14fc6f4a63b6f88eed7287))
* update OpenAPI spec from 3c46d6dcf348406bf338585fa316bce0918e8ad3 ([51c49e9](https://github.com/MercuryTechnologies/mercury-cli/commit/51c49e9b98781f07639afc58e2da8bb581532a3f))
* update OpenAPI spec from 72108644355aea945490035c66e83cc5f2344938 ([4e697c4](https://github.com/MercuryTechnologies/mercury-cli/commit/4e697c4a7e22150b8feec6865259aca054a0c6d4))
* update OpenAPI spec from a45da942ea4858e86718b1a132081277dbf81343 ([f5b64ee](https://github.com/MercuryTechnologies/mercury-cli/commit/f5b64ee28c76889ca05997389886fd1aec66240f))
* update OpenAPI spec from ef842a3c9c4d9d635ed6ab91cc18c629b01f01b7 ([cba380b](https://github.com/MercuryTechnologies/mercury-cli/commit/cba380b89cb6127e1a3a310b5f5b33f59169d4c1))
* update OpenAPI spec from v2026.03.02.11 ([3cab552](https://github.com/MercuryTechnologies/mercury-cli/commit/3cab552e80dca7875b69c8a00adf7de32946e379))
* update OpenAPI spec from v2026.03.02.13 ([3de0b91](https://github.com/MercuryTechnologies/mercury-cli/commit/3de0b91a2429bd8b14a9d6aae343fb0066cf6833))
* update OpenAPI spec from v2026.03.03.01 ([77547a2](https://github.com/MercuryTechnologies/mercury-cli/commit/77547a2550a1fe06dc1239e59cc8febc72021569))
* update OpenAPI spec from v2026.03.03.02 ([623719b](https://github.com/MercuryTechnologies/mercury-cli/commit/623719b409e26a255b1237dddf206597fe6915ae))
* update OpenAPI spec from v2026.03.03.10 ([5289d31](https://github.com/MercuryTechnologies/mercury-cli/commit/5289d31b66f1cf1221ecfe220b0560e171d1d820))
* update OpenAPI spec from v2026.03.03.16 ([1bf9e49](https://github.com/MercuryTechnologies/mercury-cli/commit/1bf9e49601626eae94ec5be067c0424ca636de9f))
* update OpenAPI spec from v2026.03.04.16 ([2143d9f](https://github.com/MercuryTechnologies/mercury-cli/commit/2143d9f5f60add692706514d5495ddbde04656ba))
* update SDK settings ([8e5d6e6](https://github.com/MercuryTechnologies/mercury-cli/commit/8e5d6e6e01d998272a9b8044004aaf2e6644ed1e))
* update SDK settings ([0776f0c](https://github.com/MercuryTechnologies/mercury-cli/commit/0776f0cfdf4ecbd92cced7d6ab9a231a0c0fe1ac))
* update SDK settings ([cb44652](https://github.com/MercuryTechnologies/mercury-cli/commit/cb446528f7984fa32996e8a90ff72c6c0ac6b6c5))
* update SDK settings ([1e2744e](https://github.com/MercuryTechnologies/mercury-cli/commit/1e2744e973c8cc3bd0f7d82dde24c15f6a4eae12))
* update SDK settings ([55708d2](https://github.com/MercuryTechnologies/mercury-cli/commit/55708d2900bce258e30cfbc54c807140b69fd9ee))
