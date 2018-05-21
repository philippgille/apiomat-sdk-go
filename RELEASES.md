Releases
========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

vNext
-----

v0.4.0 (2018-05-21)
-------------------

### aom

- Added: Debug flag for switching between printing a stack trace or not when an error occurs
- Added: Version flag for printing the version of the aom CLI (instead of the ApiOmat server, which "aom version" is for)
- Improved: Handle errors gracefully instead of panic
- Improved: Pretty print (line break) the class list output
- Improved: Move baseUrl, username, password and system flags from root to sub commands
- Improved: Pretty print (line break) the request when an error occurs
- Fixed: Backslash meant to escape quotes is contained in description of the class command

v0.3.0 (2018-04-29)
-------------------

### aoms

- Improved: Enhanced errors with stack trace and custom descriptions

### aomc

- Improved: Enhanced errors with stack trace and custom descriptions
- Changed: Removed parameter "system" from method GetClasses()
- Fixed: Empty list of classes is returned even if an error occurs during unmarshalling the JSON body from the response

### aom

- Changed: Replaced CLI lib "flags" by "github.com/spf13/cobra", leading to a complete change of the CLI (different usage, but same functionality)

v0.2.0 (2018-04-26)
-------------------

- Changed: Renamed package aomm to aoms

### aoms

- Added: Interface Client - acts as interface for consuming packages
- Added: Constant SdkVersion - it indicates for which ApiOmat version the package was implemented
- Changed: Renamed struct AomClient to DefaultClient to indicate that DefaultClient is one implementation of the Client interface
- Changed: Field baseUrl of the type DefaultClient is not exported anymore

### aomc

- Added: method NewClient(client aoms.Client) - creates a new Client that uses the given aoms.Client implementation as underlying ApiOmat HTTP client
- Improved: Added all JSON fields to the class struct by generating it with `gojson`
- Changed: Renamed struct DefaultClient to Client

### aom

- Improved: Added proper formatting to output
- Improved: Removed direct dependency on package aoms

v0.1.0 (2018-04-15)
-------------------

### aomm

- Added: Basic ApiOmat client for sending HTTP GET requests to a given URL

### aomc

- Added: Basic ApiOmat client for fetching classes of a given module

### aom

- Added: Basic CLI with parameters that shows the version of the configured ApiOmat instance and classes of a given module
