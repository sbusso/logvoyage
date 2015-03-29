# LogVoyage - fast and simple open-source logging service

LogVoyage is front-end for ElasticSearch. It allows you to store and explore your logs in real-time with friendly web ui.

Note: This is only beta version.

![Dashboard](https://raw.githubusercontent.com/firstrow/logvoyage/master/screenshots/dashboard.png)
![Live logs](https://raw.githubusercontent.com/firstrow/logvoyage/master/screenshots/live-logs.png)

## Installation

### Pre-Requirements.
- ElasticSearch
- Redis

### Installing
Installing LogVoyage is as easy as installing any other go package:
``` bash
go get github.com/firstrow/logvoyage
cd $GOPATH/src/github.com/firstrow/logvoyage
go get ./...
go install
logvoyage create_users_index
```

## Usage
Once you installed LogVoyage you need to start backend and web servers.
``` bash
logvoyage start-all
```
Or you can start/stop servers separately
``` bash
logvoyage backend
logvoyage web
```
Once server started you can access it at [http://localhost:3000](http://localhost:3000).
Execute `logvoyage help` for more info about available commands.

### Sending data to storage
TODO

## Contribution
TODO

## License
TODO

# Roadmap v0.1
TODO

## Front-end development
### Bower
To manage 3rd-party libraries simply add it to static/bower.json and run
```
bower install
```

### Building
We are using grunt to build project js and css files.
Execute next commands to setup environment:
```
npm install
grunt
```
After grunt is done, you can find result files in static/build directory.

### Auto rebuild
To automatically rebuild js, css, coffee, less files simply run in console
```
grunt watch
```

### WebSocket messages
``` coffee
// Sample coffescript code
PubSub.subscribe "log_message", (type, data) ->
  console.log data.message
```

Sample messages:

``` json
{
	"type": "log_message",
	"log_type": "nginx_access",
	"message": "test received log message goes here..."
}
```

``` json
{
	"type": "logs_per_second",
	"count": 5
}
```

<a href='https://pledgie.com/campaigns/28740'><img alt='Click here to lend your support to: LogVoyage and make a donation at pledgie.com !' src='https://pledgie.com/campaigns/28740.png?skin_name=chrome' border='0' ></a>
