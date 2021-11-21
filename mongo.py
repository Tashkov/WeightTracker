

import requests
import json
import pymongo
import pprint
from datetime import datetime, timezone
from itertools import groupby
url = "https://data.mongodb-api.com/app/data-ljspn/endpoint/data/beta/action/findOne"
adminkey = "KUBweTjfexGq6nUi9jqrzivGdvSFoXl5m4aezeeHevuJ61i6fKnzHtHXXqlBmA9T"
connection_url = "https://data.mongodb-api.com/app/data-ljspn/endpoint/data/beta"


client = pymongo.MongoClient("mongodb+srv://admin:theadmin@cluster0test.xvts5.mongodb.net/test_db?retryWrites=true&w=majority")
db=client.sample_airbnb
collection = db.listingsAndReviews
date="2020-03-01"
date_dt3 = datetime.strptime(date, '%Y-%m-%d')

query, filter = {"calendar_last_scraped":{"$lt":date_dt3}}, {'calendar_last_scraped': 1, 'accommodates': 1, '_id': 0}

result = list(collection.find(query, filter))


result= list(collection.aggregate([{
    '$group':{
            'yearMonthDay': { $dateToString: { "format": "%Y-%m-%d", "date": "$calendar_last_scraped" } },
            'sum': {'$sum': 1}
            }}
            ]))

# for date_obj in result:
#     date_obj['calendar_last_scraped'] = date_obj['calendar_last_scraped'].strftime('%Y-%m-%d')



# result.sort(key=lambda x: x['calendar_last_scraped'])
# result = [{
#             'calendar_last_scraped': key ,
#             'accommodates': sum(int(item['accommodates']) for item in group)
#            } for key, group in groupby(result, key=lambda x: x['calendar_last_scraped'])]

print(result)

with open('data.json', 'w') as outfile:
    json.dump(result, outfile)
