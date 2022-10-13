# GetNetblock
GetNetblock checks the ASN owner and netblock for a given IP address or a given organisation name.

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
  -org string
        Organisation name.
```

## Example:
Query by IP address:
```
./GetNetblock -ip 213.152.228.7
ASN is owned by:  ZSCALER-EMEA
Netblock for IP:  213.152.228.0 - 213.152.228.255
CIDR equivelant:  213.152.228.0/24

```
Query by organisation name:
```
./GetNetblock -org zscaler
CIDR ranges associated with zscaler :
213.152.228.0/24
72.37.128.0/17
199.168.148.0/24
...
```
