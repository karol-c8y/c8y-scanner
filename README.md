# c8y-scanner
PoC of security/AV scanner created in innovation week 2023. Also a demonstration of using go for microservices.

Go client from https://github.com/reubenmiller/go-c8y made by Reuben Miller 

Go version 1.19 is required

This short PoC was made to demonstrate a way to integrate antivirus tool into Cumulocity. The app is very simple and integrates ClamAV into it. We did no evaluation of any anti-viruses, ClamAV was simply chosen because it does work and fits the PoC, but not necessarily Cumulocity.

Its working for its limited scope. Raises an alarm if there is a vulnerability detected. The scan is triggered manually, via rest API.
```
PUT /service/c8y-scanner/scan/{binary-id}
```
The signature database is downloaded at the start of microservice, which takes about 1 minute.

TODO next list:
- select more/better scanners to be used. This microservice can work either with more scanners or be deployed as one per scanner.
- implement some error handling/tighter integration with Cumulocity. Current implementation is very basic
- implement microservice as MULTI_TENANT. Current PER_TENANT is probably a bit waste.
- implement real-time scanning of files.
- consider what to do with detected malware (remove? quarantine? etc)
