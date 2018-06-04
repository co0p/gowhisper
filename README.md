gowhisper
=========

[![Drone build status](http://ci.co0p.org/api/badges/co0p/gowhisper/status.svg)](http://ci.co0p.org/co0p/gowhisper)


A dead simple website / service status checker for http status; nice ui included.

Usage
-----

gowhisper uses the following command line arguments:

 * configurationFile string -- path/to/configuration file (required)
 * pollingInterval int -- polling interval in seconds (10 - 360) (default: 60)
 * port int -- the port to serve the ui to (default: 8080)


A call looks like:
```bash
./gowhisper -configurationFile whisper.json -pollingInterval 60 -port 8000
```


Here is a sample configuration json:
```json
[
    {
        "Label":"Service1",
        "URL": "https://service1.de",
    },
    {
        "Label":"Service2",
        "URL": "https://service2.com/api/healthz",
    }
]
```
