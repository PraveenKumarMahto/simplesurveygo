from jsonschema import validate

schema = {
            "type" : "object",
            "properties" : { 
                    "colors" :[
                         {
                             "color" : "string",
                            "category" : "string",
                            "type" :"string",
                            "code" : {
                                "rgba" : [int]
                                "hex" : "string"
                                }
                         },
                         {
                             "color" : "string",
                             "category" : "string",
                             "code" : {
                                "rgba" : [int]
                                "hex" : "string"
                                }
                         }
                    ]
            }

def validator(payload):
    try:
        validate(payload,schema)
        return True
    except:
        return False


import json
from pprint import pprint

with open('data.json') as f:
    data = json.load(f)

pprint(data)
validator(schema):
