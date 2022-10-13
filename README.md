# GetNetblock
GetNetblock checks the ASN owner and netblock for a given IP address.

You will need to register for a WhoisXMLAPI API key here: https://ip-netblocks.whoisxmlapi.com/api/signup

This allows for up to 1000 free requests a month.

Add your key to main.go here:

```
//add your API key here (free from https://ip-netblocks.whoisxmlapi.com/api/signup)
var apiKey = "<your_WhoisXMLAPI_key>"
```

## Usage:
```
$ ./GetNetblock --help
Usage of ./GetNetblock:
  -ip string
        IP address.
```

## Example:
```
./GetNetblock -ip 8.8.8.8
ASN is owned by:  GOOGLE
Netblock for IP:  8.8.8.0 - 8.8.8.255
```
