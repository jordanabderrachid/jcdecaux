# jcdecaux #

jcdecaux is a Go client library for accessing the [JCDecaux Bike API](https://developer.jcdecaux.com).

## Usage ##

```go
import "github.com/jordanabderrachid/jcdecaux"
```

First, construct a new JCDecaux client and provide your private API key. You will need to register on the JCDecaux
[website](https://developer.jcdecaux.com) to get your key. If you don't set your private key, the error `ErrUnsetAPIKey`
will be returned on performing requests.

```go
client := jcdecaux.Client{APIKey: "<put your private key here>"}
```

All realtime request are implemented by the library.

- Get the list of all contracts `GetContract()`
- Get the list of all stations `GetStations()`
- Get the list of all stations under a specific contract `GetStationsByContract(contractName)`
- Get a specific station `GetStation(stationNumber, contractName)`


## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE.md)
file.
