# WEB 服务

## 简单的 WEB 服务

```shell
python -m SimpleHTTPServer
```

## 延时测试 WEB 服务

sleep.py

```python
#! /usr/bin/env python
# -*- coding: utf-8 -*-
import time
import sys
import urlparse
from SocketServer import ThreadingMixIn
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer, test

sleep = 3


class ThreadedHTTPServer(ThreadingMixIn, HTTPServer):
    pass


class SleepHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        parsed_path = urlparse.urlparse(self.path)
        try:
            params = dict([p.split('=') for p in parsed_path[4].split('&')])
        except:
            params = {}
        s = sleep
        if 'sleep' in params:
            s = int(params['sleep'])
        print("start sleep: %ds" % s)
        time.sleep(s)
        print("stop sleep.")

        self.wfile.write("request sleep: %ds ok.\n" % s)


if __name__ == '__main__':
    if len(sys.argv) > 2:
        sleep = int(sys.argv[2])
    print("python sleep.py [端口号] [默认延时,单位秒]")
    print("request sleep default: %ds.\n" % sleep)
    print("例子:")
    print("curl http://localhost:8000\?sleep\=10")

    test(SleepHandler, ThreadedHTTPServer)
```
