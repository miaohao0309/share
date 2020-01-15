# -*- coding: utf-8 -*
#!/usr/bin/python3

import json
import urllib.request


def send_message(msg):

    url = 'https://qyapi.weixin.qq.com'
    corpid = 'wwf5a3cce9198caae5'
    api_secret = 'QypjRFHOBDpNYVHXripNjokGcFi4PNjSsTReEPQPFLc'
    token_url = '%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s' % (url, corpid, api_secret)
    token = json.loads(urllib.request.urlopen(token_url).read().decode())['access_token']
    send_url = '%s/cgi-bin/message/send?access_token=%s' % (url,token)
    
    data = {
        "touser": "@all",
        "msgtype": 'text',
	"agentid": 1000006,
	"text": {
            "content" : "msg"
            },
	"safe": 0
    }

    params = (bytes(json.dumps(data), 'utf-8'))
    r = urllib.request.urlopen(urllib.request.Request(url=send_url, data=params)).read()
    x = json.loads(r.decode())['errcode']
    if x == 0:
        print ("Send Succesfully")
    else:
        print (x)
        print ("Send Failed")

def main():
    a = "xxxxxxx"
    b = "yyyyyyy"
    c = "zzzzzzz"
    msg = a+b+c
    send_message(msg)	

if __name__ == '__main__':
    main() 

