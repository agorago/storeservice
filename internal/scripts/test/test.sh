curl  -XGET -H "Key: hello" -H "Value: world" localhost:5000/storeService/store
curl  -XGET -H "Key: hello"  localhost:5000/storeService/retrieve