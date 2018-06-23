Releases
========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

vNext
-----

- Changed: Moved `aom` CLI to its own GitHub repository: [https://github.com/philippgille/apiomat-cli](https://github.com/philippgille/apiomat-cli)

v0.7.0 (2018-06-17)
-------------------

### aomx

- Improved: GoDoc (added code example, improved existing comments)

### aomc

- Added: Type `dto.Backend`, which is the DTO for a Backend, which represents an ApiOmat Backend (sometimes called "Application")
- Added: Method `(client Client) GetRawBackends(customer string) ([]dto.Backend, error)` for getting backend DTOs for all backends a customer has READ permissions for
- Added: Method `(client Client) GetRawBackendByName(name string, customer string) (dto.Backend, error)` for getting a specific backend a customer has READ permissions for
- Added: Method `(client Client) GetRawClassByName(module string, name string) (dto.Class, error)` for getting a single class DTO with a given name
- Added: Method `(client Client) GetClassByName(module string, name string) (Class, error)` for getting a single class with a given name
- Improved: GoDoc (added code example, improved existing comments, fixed wording)

v0.6.0 (2018-06-02)
-------------------

### aoms

- Changed: Renamed package `aoms` to `aomx` to better reflect the cross cutting functionality

### aomc

- Added: Functions `ConvertRawClassesFromJSON(...)`, `ConvertRawClassFromJSON(...)`, `ConvertRawAttributesFromJSON(...)` and `ConvertRawAttributeFromJSON(...)` to wrap the current use of `json.Unmarshal(...)` to make this exchangeable in the future without having to change tests, but also for use by SDK users
- Changed / improved: Change type of `Class` fields `AuthImplStatus` (previously `UseOwnAuth`) and `UserRole` from string to their own type, with constants for use as an enum. They implement `String()` for conversion and printing.
- Changed: Renamed `Class` and `Attribute` struct fields to better reflect their meaning (the JSON attribute names from ApiOmat are inconsistent and not very expressive in some cases) and to adhere to Go's coding standards
- Changed: Renamed parameters of some functions and methods to adhere to Go's coding standards

v0.5.0 (2018-05-28)
-------------------

### aoms

- Added: Function `MustUrl(url *url.URL, err error) *url.URL` for parsing and dereferencing a URL in one line

### aomc

- Added: Package `dto` (`github.com/philippgille/apiomat-go/aomc/dto`)
- Added: Struct `dto.Attribute`, which represents a raw ApiOmat "MetaModelAttribute"
- Added: Method `(client Client) GetRawAttributes(module string, classId string) ([]dto.Attribute, error)` for fetching the raw attributes of a class
- Added: Struct `Class`, which is a convenience type which includes attribute objects and with better field names and types and other advantages over the raw one
- Added: Struct `Attribute`, which is a convenience type with better field names and types and other advantages over the raw one
- Added: Method `(client Client) GetClasses(module string) ([]Class, error)`, which fetches classes of a module and adds `Attribute` objects to them
- Added: Method `(client Client) GetAttributes(module string, classId string)` for fetching the attributes of a class
- Added: Function `ConvertClassFromDto(rawClass dto.Class) Class` for converting class DTOs into the `Class` convenience type
- Added: Function `ConvertAttributeFromDto(rawAttribute dto.Attribute) Attribute` for converting attribute DTOs into the `Attribute` convenience type
- Changed: Moved old / raw struct `Class` to `dto.Class`
- Changed: Renamed old method `GetClasses(...)` to `GetRawClasses(...)`

### aom

- Changed: Output of classes now is from printing the new `Class` instead of `dto.Class`, so it includes the attributes of the class and has better field names and types

v0.4.0 (2018-05-21)
-------------------

### aoms

- Improved: Pretty print (line break) the HTTP request when an error occurs

### aom

- Added: Debug flag for switching between printing a stack trace or not when an error occurs
- Added: Version flag for printing the version of the aom CLI (instead of the ApiOmat server, which "aom version" is for)
- Improved: Handle errors gracefully instead of panic
- Improved: Pretty print (line break) the class list output
- Improved: Move baseUrl, username, password and system flags from root to sub commands
- Fixed: Backslash meant to escape quotes is contained in description of the class command

v0.3.0 (2018-04-29)
-------------------

### aoms

- Improved: Enhanced errors with stack trace and custom descriptions

### aomc

- Improved: Enhanced errors with stack trace and custom descriptions
- Changed: Removed parameter `system` from method `GetClasses(...)`
- Fixed: Empty list of `Class` structs is returned even if an error occurs during unmarshalling the JSON body from the response

### aom

- Changed: Replaced CLI lib "flags" by "github.com/spf13/cobra", leading to a complete change of the CLI (different usage, but same functionality)

v0.2.0 (2018-04-26)
-------------------

- Changed: Renamed package `aomm` to `aoms`

### aoms

- Added: Interface `Client` - acts as interface for consuming packages
- Added: Constant `SdkVersion` - it indicates for which ApiOmat version the package was implemented
- Changed: Renamed struct `AomClient` to `DefaultClient` to indicate that `DefaultClient` is one implementation of the `Client` interface
- Changed: Field `baseUrl` of the type `DefaultClient` is not exported anymore

### aomc

- Added: Function `NewClient(client aoms.Client) Client` - creates a new `Client` that uses the given `aoms.Client` implementation as underlying ApiOmat HTTP client
- Improved: Added all JSON fields to the `Class` struct by generating it with `gojson`
- Changed: Renamed struct `DefaultClient` to `Client`

### aom

- Improved: Added proper formatting to output
- Improved: Removed direct dependency on package `aoms`

v0.1.0 (2018-04-15)
-------------------

### aomm

- Added: Basic ApiOmat client for sending HTTP GET requests to a given URL

### aomc

- Added: Basic ApiOmat client for fetching classes of a given module

### aom

- Added: Basic CLI with parameters that shows the version of the configured ApiOmat instance and classes of a given module
