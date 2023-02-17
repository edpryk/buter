# Buter - Brute Force Tool

Works as other tools but multiple payload set can \
be used.

Docs:
```
Version:   0.0.1
Contact:   https://github.com/0x640x720x650x640x64
Author:    https://github.com/0x640x720x650x640x64

Started:   16:10:52
Error:     No AttackType provided

Usage: buter.exe
        Any payload/fuzzing position must be highlighted with ! char

        -u   <http://localhost?param1=!abc!&param_N=!efg!> (Url)
        -p   <payload-file_1> -p <payload-file_N> (Payload)
        -a   cluster/sniper/dos (AttackType)
        -t   5 (Max Concurrent Threads)
        -h   '{ "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/99.1" }' (Headers)
        -d   800 (Delay in milliseconds)
        -m   POST (HTTP method)
        -r   3 (Retries on request error)
        -rd  1000 (Retry delay in milliseconds)
        -b   {"email":"user_nameg@mail.com","password":"12345"} (request body)
        -T   10 (Request timeout in Seconds)
        -R   10 (request amount in DOS mode)
        -f   status:200,201;length:1553 (Output filters)
        -S   status:200 (Stop attack on event)

```

Example 1: 
```
buter \
-b '{\"email\":\"test@gmail.com\",\"password\":\"!abc!\"}' -h '{\"Cookie\": \"!some_value!\"}' \
-m POST \
-u "https://bla.bla.com?param=!some_value!" \
-p /path/payload-list1 \
-p /path/payload-list2 \
-p /path/payload-list3 \
-f "length:1337" \
-a cluster
```

Example 2: 
```
buter \
-m POST \
-b '{\"email\":\"admin@juice-sh.op\",\"password\":\"!admin!\"}' \
-u https://juice-shop.herokuapp.com/rest/user/login \
-p \Code\Src\SecLists\Passwords\cirt-default-passwords.txt \
-a sniper \
-d 50 \
-S "status:200"
```
Stopped on status 200 \

![Buter stop on status 200](/assessts/buter_stop_on_200.png)

Example 3 (simplified requst amount):
```
buter \
-u "https://juice-shop.herokuapp.com/rest/products/reviews"\
-m POST \
-h '{\"Authorization\": \"Bearer token\"}' \
-b '{\"id\":\"Xx2pFWBugt2HKPHrt\"}' \
-a dos
```
![Buter stop on status 200](/assessts/buter_dos_example.png)
