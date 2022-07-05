Coding Task
===========

One of the application of the company dumps some data of network traffic into CSV file.
Write an application which converts the CSV file to NDJSON file after some modification and enrichment. 

Input CSV file:
---------------
Contains a timestamp (ISO 8601, like: 2022-04-21T10:13:31Z), a source IPv4 address, a target url and an integer represents the size of the traffic. 

Sample CSV file
```
2022-04-21T10:13:31Z,1.2.3.4,www.yahoo.com/abc,12000
2022-04-21T10:13:32Z,1.2.3.4,www.google.com/abc,1200
2022-04-21T10:13:30Z,1.2.3.4,radar.com/test,1201
```

Output NDJSON file
------------------
As an NDJSON file each line should be a valid JSON value. Since it will be loaded into a database the order of the records is not relevant (can be diffrent than it is in the original CSV)
The JSON record should contain the following data:
- ts : timestamp in EPOCH format
- source_ip: IPV4 source address (2rd column of CSV). Can contains only valid IPv4 address
- url: the parsed URL record (from 3rd column of CSV). Use https://github.com/goware/urlx to normalize the URL
    - Scheme
    - Host
    - Path
    - Opaque
- size: the size of traffic (4rd column of CSV)
- note: we add(enrich) some data to the record. For the sae of the simplicity in now we use a public API which returns a random text (http://numbersapi.com/random/math). This parameter should contain the result text of the API (this API should call on every line)


Sample NDJSON file
```
{"ts": 1650537316, "source_ip":"1.2.3.4", "url":  {"Scheme": "","Host": "www.yahoo.com", "Path":  "/abc", "Opaque":""}, "size": 12000, "note": "236 is the number of possible positions in Othello after 2 moves by both players."}
{"ts": 1650537317, "source_ip":"1.2.3.4", "url": {"Scheme": "","Host": "www.google.com", "Path":  "/abc", "Opaque":""}, "size": 1200, "note": "6615 is an odd abundant number."}
{"ts": 1650537315, "source_ip":"1.2.3.4", "url": {"Scheme": "","Host": "radar.com", "Path":  "/test", "Opaque":""}, "size": 1201, "note": "1060 is the sum of the primes less than 100."}
```


Additional requirements and notes
---------------------------------
- Implement it in Golang
- The conversion tool has at least 2 arguments: input file name, output file name
- The input file can be huge, the long processing tasks should be parallelized
- Invalid lines should be displayed in the console
- Use comments when and if necessary
- Unit tests are welcome

