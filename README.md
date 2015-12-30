# gocollectanalytics

A library that supplies an http handler function to collect data according to a syntax based on the [Google Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters) and stores it in the specified type of store.

The idea here is to create a web app similar to the one in the examples folder and then write some js to send data to that app from a client site. 


## Build and test this project

    ./build

##  Implementation

Go get this project and import as a dependency to a web app. Implement as an http.HandlerFunc on a route like:

    logCollector, _ := gocollectanalytics.LogCollector()
    
    mux.HandleFunc("/collect", logCollector.CollectData)

A LogCollector just writes the sent data out as JSON to log.

##  Sending data to KeenIO

To send data to [KeenIO](https://keen.io/) you need to set up a Keen project and obtain a write key and a project ID. These need to be supplied as properties of a config object like this:

    
    kc := gocollectanalytics.KeenIOConfig{WriteKey: os.Getenv("KEENIO_WRITE_KEY"), ProjectID: os.Getenv("KEENIO_PROJECT_ID")}
    
    keenioCollector, err := gocollectanalytics.KeenIOCollector(kc)
      if err != nil {
      log.Fatal(err)
    }

    mux.HandleFunc("/collect", keenioCollector.CollectData)

At present this only stores KeenIO 'events' in a collection called 'Events'.

THIS PART NEEDS BREAKING INTO A SUBPROJECT SO THAT YOU DONT *HAVE* TO USE KEENIO.


## Sending data from a client

Following the [Google Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters):

* A version `v=1` must be supplied. 
* A site ID `tid=$string` must be supplied
* A hit type `t` must be supplied. Currently only `t=event` is supported
* Events must have a category `ec=$string` and an action `ea=$string`. They can optionally have a label and a value.
* Currently a client ID `cid` is NOT required as tracking the behaviour of individual users isn't my initial focus.

An example of a valid request would be 

    /collect?v=1&tid=test&t=event&ec=test&ea=test


