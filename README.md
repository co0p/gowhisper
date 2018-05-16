gowhisper
=========

[![Drone build status](http://ci.co0p.org/api/badges/co0p/gowhisper/status.svg)](http://ci.co0p.org/co0p/gowhisper)


A simple polling tool that checks http status and return value and reports any services missing via mail

Usage
-----

gowhisper needs the following command line arguments:

 * configurationFile string -- path/to/configuration file
 * pollingInterval int -- polling interval in seconds (10 - 360)


A call looks like:
```bash
./gowhisper -configurationFile whisper.json -pollingInterval 60
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
