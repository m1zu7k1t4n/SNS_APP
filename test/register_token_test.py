import requests
import urllib
import random
from pprint import pprint
import asyncio
import logging

url = 'http://localhost:8080/api/v1/account/'

def test_func(i):
  base = random.random()
  username = str(base + 1)
  password = str(base + 10)
  item_data = {
    "email": username+"@"+password+".com",
    "username": username,
    "password": password
  }

  item_data = urllib.parse.urlencode(item_data)

  headers = {"Content-Type": "application/x-www-form-urlencoded"}
  res1 = requests.post(url+"register", headers=headers, data=item_data)
  res2 = requests.post(url+"token", headers=headers, data=item_data)
  print(str(i) + "\n" + str(res1.text) + "\n" + str(res2.text) + "\n" + "-----------------------------------------------------------------")
  return i


async def run(loop):
  async def run_req(i):
    return await loop.run_in_executor(None, test_func, i)

  tasks = [run_req(i) for i in range(50)]
  return await asyncio.gather(*tasks)


for _ in range(100):
  loop = asyncio.get_event_loop()
  print(loop.run_until_complete(run(loop)))
