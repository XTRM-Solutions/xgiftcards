# XGIFTCARDS

A simple tool to fetch gift cards and available amounts with SKU for XTRM gift cards.

## Getting Started as a Developer

Written in [go](https://go.dev) for cross-platform convenience, the simplest way
to fetch this program is via [go](https://go.dev) itself. Install go on the
system (please refer to the instructions at [download and install go](https://go.dev/doc/install)
for your operating system. Create a directory for the project, `cd` into that directory, and
[go](https://go.dev) will get the code for you:
```
go get github.com/XTRM-Solutions/xtrm_giftcards
```
## Download and install an executable on your system
Again, the simplest way to do this is to install [go](https://go.dev) on the system,
and to set the system environment variable `GOBIN` to a directory on the system
that is also in the executable `PATH`. [Go](https://go.dev) uses the `GOBIN` location
to install the executable file on the system with the command:
```
go install github.com/XTRM-Solutions/xtrm_giftcards@latest
```

There is also a zipped 65-bit windows executable in the releases section. Download, unzip, and
put the executable file in the path.

### Prerequisites

* A working [go](https://go.dev) installation
* A sandbox developer account on [XTRM](https://sandbox.xtrm.com)
    * API access in your account
    * An allow-listed IP address from which to make calls (this may be requested from the console application after the sandbox account is created.)
    * An `xtrm.ini` file in the working directory (first choice), or the same directory as the executable (second choice).
        * Please note that the following fields differ from account to account. This file contains dummy values.
            * `xClient`
            * `xSecret`
            * `xDefaultWallet` (this value is not used in this program)
* This code will work if directed at a live, production account, but please **do not create test data on [XTRM](https://www.xtrm.com)'s production servers**.
* This code produces a log file in the working directory `<program_name>log`; the log file is overwritten with each program run.

#### SAMPLE `XTRM.INI` FILE
```
[DEFAULT]
currentSection = INITIAL

[INITIAL]
apiAuthorizeUrl = https://xapisandbox.xtrm.com/oAuth/token
xIssuerID       = SPN89754632
xClient         = 89754632_API_User
xSecret         = PyhxJkl4f89b/2opDOKcuqqfbrXvnvId7/s319u1==
xUrl            = https://xapisandbox.xtrm.com/API/V4
xDefaultWallet  = 111203
```

`currentSection` : This app is single-threaded, and retains state information about API keys in the .ini file. It's not the most secure thing in the world, and improving the security aspect of key storage is a long-term goal. Regardless, the `currentSection` is the presently used section. The tool will at some point support a master section (for the managing account), and subsections for managed remitter keys.

`apiAuthorizeUrl` : the endpoint for authorization. This sample file points to the sandbox.

`xIssuerID` : SPN for the company.

`xClient` : This is the API client ID

`xSecret` : This is the API secret corresponding to the client ID. Not securely stored, but having it as an environment variable is **not** safer; it just **looks** safer. Pretending unsafe things are safe or safer than the alternative is also a poor security strategy.

`xUrl` : The endpoint prefix for API calls. In this case, it points to the sandbox and API version 4.

`xDefaultWallet` : This value is deprecated, and will be removed someday.



## Usage

Here are a few command lines to show the program

* `xtrm_giftcards --all`
    * Print out all gift cards with SKUs, brands, and values
* `xtrm_giftcards --currency USD`
    * Print out all gift cards denominated in US dollars
* `xtrm_giftcards --help`
    * Print usage information
