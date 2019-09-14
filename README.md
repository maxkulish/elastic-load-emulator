# elastic-load-emulator
Utility to put documents into Elastic index and emulate workload

## How to install
Linux
```bash
wget https://github.com/maxkulish/elastic-load-emulator/raw/master/bin/Linux/elasticLoad
chmod +x elasticLoad
./elasticLoad -example
```


MacOS
```bash
wget https://github.com/maxkulish/elastic-load-emulator/raw/master/bin/macOS/elasticLoad
chmod +x elasticLoad
./elasticLoad -example
```

## How to use

1. Create index.json file.

This json file will be put in ElasticSearch index.
You can use any json structure or generate an example.

```bash
./elasticLoad -example
```
As example index.json I've used [Twitter example](https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/intro-to-tweet-json.html) twit
```json
{
 "created_at": "Wed Oct 10 20:19:24 +0000 2018",
 "id": 1050118621198921728,
 "id_str": "1050118621198921728",
 "text": "To make room for more expression, we will now count all emojis as equal—including those with gender‍‍‍ ‍‍and skin t… https://t.co/MkGjXf9aXm",
 "user": {},  
 "entities": {}
}
```

Second example file is config.yml
```yaml
# Example ElasticSearch config. Put your data
elastic:
  proto: "http"
  host: "1.1.1.1"
  port: "9200"
  username: ""
  password: ""
  index: "test-index"
```

Change 'host' to ElasticSearch real IP

Before start your working directory should looks like this:

```bash
your_folder
   ├── config.yml
   ├── elasticLoad
   └── index.json
```

#### Run Elastic load emulation
```bash
./elasticLoad -start=1 -stop=1000
```
This command will put 1,000 documents to index "test-index"

If you get error
```bash
elastic connection error: Head http://1.1.1.1:9200: context deadline exceeded
```
you started load emulation and example config was used. Change host to ElasticSearch IP address.

If everything is OK you should get output:
```bash
2019/09/14 10:49:18 Elasticsearch returned with code 200 and version 7.3.1
2019/09/14 10:49:18 sent index: 1
2019/09/14 10:49:18 sent index: 2
2019/09/14 10:49:18 sent index: 3
2019/09/14 10:49:18 sent index: 4
2019/09/14 10:49:18 sent index: 5
...
2019/09/14 10:50:49 Done! Spent: 5.527262 seconds
2019/09/14 10:50:49  Started: 1, finished: 100

```