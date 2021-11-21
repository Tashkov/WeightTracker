from flask import Flask
import pymongo


app = Flask(__name__)
client = MongoClient('mongodb://localhost:27017/')

@app.route("/")
def hello_world():
    client = pymongo.MongoClient("mongodb+srv://admin:<password>@cluster0test.xvts5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
    db = client.test

    return "<p>Hello, World!</p>" + db