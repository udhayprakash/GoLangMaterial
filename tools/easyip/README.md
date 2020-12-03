# EasyIP
Easily get your public IPv4 address with this program!

### Usage
```
easyip [-i] [-n] [-q]
  -i    Print information about the script.
  -n    Don't print a line feed at the end. (Only affects Unix systems)
  -q    Quieter, print only the IP. (Except for errors)
```


### How it works
EasyIP will attempt a connection to the first host in the list below. If the connection succeeds, the public IP is displayed and the program exits.
If the connection fails, it will attempt a connection to the next host and so on, until it gets a successful connection, otherwise it will display an error and exit.


### Host list
The list of hosts EasyIP will try connecting to:
```
http://ifconfig.me/ip
http://api.ipify.org
http://ipecho.net/plain
https://ip.seeip.org
```