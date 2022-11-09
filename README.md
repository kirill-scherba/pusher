# Pusher

Pusher send metrics to Prometheus Push Server using http post request. This
package has not any external dependencies and use only standard go librarie.

[![GoDoc](https://godoc.org/github.com/kirill-scherba/pusher?status.svg)](https://godoc.org/github.com/kirill-scherba/pusher/)
[![Go Report Card](https://goreportcard.com/badge/github.com/kirill-scherba/pusher)](https://goreportcard.com/report/github.com/kirill-scherba/pusher)

## Usage example

The sample code below show all this package capabilities. Just create new pusher
and send any methrics.

```go
// Initialize random
rand.Seed(time.Now().Unix())

// Get this host name
hostName, err := os.Hostname()
if err != nil {
    log.Fatalln("can't get host name error: ", err)
}

// Create publisher
pu := pusher.NewPusher("http://example.com:9091", "my_job", hostName)

// Push metrics every 15 seconds
for {
    m, err := pu.Push(
        pusher.Metric("my_job_couner_2", rand.Float64()*10),
        pusher.Metric("my_job_couner_3", rand.Float64()*10),
        pusher.Metric("my_job_couner_4{label=\"val1\"}", rand.Float64()*100),
        pusher.Metric("my_job_couner_4{label=\"val2\"}", rand.Float64()*100),
    )
    if err != nil {
        log.Println("push error: ", err)
    } else {
        log.Println("push metrics:\n" + m)
    }
    time.Sleep(15 * time.Second)
}
```

You can find complete packets documentation at: <https://pkg.go.dev/github.com/kirill-scherba/pusher>

-----------------------

## Licence

[BSD](LICENSE)
