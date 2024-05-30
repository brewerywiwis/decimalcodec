# Shopspring Decimal codec for MongoDB

## Usage

### Example 1

```
import "github.com/brewerywiwis/decimalcodec"

func NewMongo() {
    ...
    opts := options.Client()

    registry := bson.DefaultRegistry // Or another *bsoncodec.Registry

    decimalcodec.Register(registry)

    opts.SetRegistry(registry)
}
```

### Example 2

```
import "github.com/brewerywiwis/decimalcodec"

func NewMongo() {
    ...
    opts := options.Client()

    registry := decimalcodec.NewDefaultRegistryWithDecimalCodec()

    opts.SetRegistry(registry)
}
```
