# Smarty Street

## Installation

`go get github.com/smartystreets/smartystreets-go-sdk`

## Sample Code

```go

client := wireup.BuildUSStreetAPIClient(wireup.SecretKeyCredentail(key, pair))

lookup := &street.Lookup(Street: "1 Rosedale", LastLine: "Baltimore MD", MaxCandidates: 10)

if err := client.SendLookups(lookup); err != nil {
}

for c, candidate := range lookup.Results {
    candidate.DeliveryLine1, candidate.LastLine
}
```
