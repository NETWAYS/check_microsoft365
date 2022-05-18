# check_microsoft365

check_microsoft365 is an Icinga check plugin, which is capable to check in the Microsoft 365 context via the Microsoft
Graph API.

In the current version check_microsoft365 supports the service health overview context.

## Usage

#### service health overview

Checks the health information of one or more services subscribed by a tenant.

```
Usage:
  check_microsoft365 servicehealth

Flags:
  -n, --servicenames strings      The name of one or more specific service/s
      --state-override strings    States to override (e.g. STATENAME=ok)
      --all                       Displays all services regardless of the status
  -M, --display-message           Displays the issue message to the specified service
  -i, --issue-start-time string   Displays only issue massages in the period of time given. Possbile values are e.G. '1h', '30m' (default "15m")
  -h, --help                      help for servicehealth


Global Flags:
  -c, --clientid string       The client id from your app registration (env: AZURE_CLIENT_ID)
  -s, --clientsecret string   The client secret from your app registration (env: AZURE_CLIENT_SECRET)
  -S, --scope string          Represents the definition of a delegated permission (default "https://graph.microsoft.com/.default")
  -T, --tenantid string       The tenant id from your app registration (env:AZURE_TENANT_ID)
  -t, --timeout int           Timeout for the check (default 30)
```

````
$ check_microsoft365 servicehealth --clientid "example_client_id" \
                                   --clientsecret "example_client_secret"
                                   --tenantid "example_client_tenant"
                               
CRITICAL - Found 27 services - 2 Critical -  0 Warning:
[CRITICAL] Exchange: Exchange Online SERVICEDEGRADATION
[CRITICAL] Microsoft365Defender: Microsoft 365 Defender EXTENDEDRECOVERY
 | total=27 warning=0 critical=2

````

```
$ check_microsoft365 servicehealth --clientid "example_client_id" \
                                   --clientsecret "example_client_secret"
                                   --tenantid "example_client_tenant"
                                   --state-override "SERVICEDEGRADATION=ok"
                                   --state-override "EXTENDEDRECOVERY=ok"
                                    
OK - Found 27 services - 0 Critical -  0 Warning:
All services are healthy | total=27 warning=0 critical=0
```

```
$ check_microsoft365 servicehealth --serivenames "Exchange Online"
                                   --serivenames "Microsoft 365 Defender"
                                   --state-override "EXTENDEDRECOVERY=warning"

CRITICAL - Found 2 services - 1 Critical -  1 Warning:
[CRITICAL] Exchange: Exchange Online SERVICEDEGRADATION
[WARNING] Microsoft365Defender: Microsoft 365 Defender EXTENDEDRECOVERY
 | total=2 warning=1 critical=1

```

## Setting up Access

In order to work correctly you need the correct permissions and configuration within Azure, to grant the plugin proper
read-only access to the resources.

The following step-by-step instructions will help you to setup this configuration.

### App Registration

In Azure, withing the Azure Active Directory, search for the key word **App registrations** and add a new registration
with a meaningful name for the app registration like `check_microsoft365`.

If the app registration was successfully, it should appear under the tab **Owned applications**. pen the app details and
navigate to the section **Certificates & secrets**, add a new client secret.

### Give app read access

Now the `check_microsoft365` **App Registration** needs *ServiceHealth.Read.All* access to fetch monitoring values via
the Microsoft Graph API.

In Azure, navigate to the previous registered App and search for the key word `API permissions`. Then click
on `Add a permission`
and select on the right-hand side `Microsoft Graph`. According to that click on `Application permissions` and
search/select for
`ServiceHealth.Read.All`, which allows the app to read your tenant's service health information, without a signed-in
user.

For more information, visit [How to register an app](https://docs.microsoft.com/en-us/graph/auth-register-app-v2)
, [How to find the tanant id](https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant)

## License

Copyright (C) 2022 [NETWAYS GmbH](mailto:info@netways)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
License as published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <http://www.gnu.org/licenses/>.
