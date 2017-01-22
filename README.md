This script updates your /etc/hosts file with the latest version from
http://someonewhocares.org/hosts/ which provides values to prevent
your computer from connecting to selected internet hosts in order to
block spam, spyware, etc.

A backup of your old hosts file will be created.

Building
========

Tested with Go 1.7, but this is a very simple script
and should also work with older versions.

```
go build .
```

Usage
========

```
sudo ./update_hosts_file
```
