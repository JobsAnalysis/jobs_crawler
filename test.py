from requests_html import HTMLSession
from time import sleep
from fake_useragent import UserAgent
import json
session = HTMLSession()
a = session.get(
    url='https://www.lagou.com/jobs/list_爬虫工程师?labelWords=&fromSearch=true&suginput=',
    headers={
        'useragent': UserAgent().random})
sleep(5)
url = 'https://www.lagou.com/jobs/positionAjax.json'
ua = UserAgent()
headers = {
    'user-agent': ua.random,
    'referer': 'https://www.lagou.com/jobs/list_%E6%95%B0%E6%8D%AE%E5%88%86%E6%9E%90?city=%E5%85%A8%E5%9B%BD&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot',
    'Origin': 'https://www.lagou.com',
    'Accept': 'application/json, text/javascript, */*; q=0.01',
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}
data = {
    'first': 'true',
    'pn': '1',
    'kd': '爬虫工程师',
    'city': '北京',
    'needAddtionalResult': 'false'
}


r = session.post(
    url=url,
    headers=headers,
    data=data
)
content = r.content

json_msg = json.loads(content)

print(type(json_msg))
with open('lagou.json', 'w') as fd:
    json.dump(json_msg, fd)