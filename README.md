# gocollectanalytics

A library that supplies an http handler function to collect data according to a syntax based on the [Google Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters) and stores it in the specified type of store.


## Build and test this project

    ./build

##  Implementation

Go get this project and import as a dependency to a web app. Implement as an http.HandlerFunc on a route like:

    logCollector, _ := gocollectanalytics.LogCollector()
    
    mux.HandleFunc("/collect", logCollector.CollectData)

