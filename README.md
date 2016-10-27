# Example of two microservices on framework go-micro for perfomance test
### How use
- `git clone https://github.com/dzen-it/testmicro.git`
- `cd firstservice && go install`
- `cd secondservice && go install`
- [install a consul](https://www.consul.io/downloads.html)
- run `consul agent -dev -server -client=0.0.0.0`
- run:</br>
 `secondservice --selector=cache --client_pool_size=4000`</br>
 `firstservice --selector=cache --client_pool_size=4000`  
- Get `http://localhost:8088?name=Hello`

### How to test
Download and install [Vegeta](https://github.com/tsenart/vegeta) </br>
In console: `echo "GET http://dockerhost:8088/?name=Hello" > test.json` </br>
Run: `vegeta attack -targets test.json -connections 10 -rate 2000 -duration 60s | vegeta report -output test.html -reporter plot` </br>
Look `test.html`
