# NetCLI
A Lightweight network tool, which helps you do multiple things.

## Features
> Some features are still under development.
- [x] DNS Lookup.
- [x] Shows your public IP.
- [x] Internet speed tester. (Thanks to [speedtest-go](https://github.com/showwin/speedtest-go))
- [x] Find subdomains.
- [x] Whois Lookup.
- [x] URL Shortener.
- [ ] Email Checker(check validity and reachability).
- [ ] HTTP Header information.
- [ ] Subnet Calculator.

## Usage
Simply type `netcli -h` and you'll get to know what options are available and how to use them.
```sh
COMMANDS:
   nameserver, ns    Print the nameserver of given hostname.(NS RECORD)
   mailserver, mx    Print the mail server of given hostname.(MX RECORD)
   a                 Print the host ip of given hostname.(A RECORD)
   cname             Print the canonical name of given hostname.(CNAME)
   myip              Print your public IP Address.
   speedtest, speed  Do an internet speed test.
   subdomain, subd   Scans an entire domain to find as many subdomains as possible.
   whois             Get whois information of Domain Name or IP Addres.
   shorten, short    URL shortener to reduce a long link.
   help, h           Shows a list of commands or help for one command
```