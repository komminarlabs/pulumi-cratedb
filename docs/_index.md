---
title: CrateDB
meta_desc: Provides an overview of the CrateDB Provider for Pulumi.
layout: package
---

The CrateDB provider for Pulumi can be used to provision the resources available in [CrateDB](https://www.cratedb.com/).

The CrateDB provider must be configured with credentials to deploy and update resources in CrateDB; see [Installation & Configuration](./installation-configuration) for instructions.

## Example

{{< chooser language "typescript,python,go" >}}
{{% choosable language typescript %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as cratedb from "@komminarlabs/cratedb";

// Create a new Bucket
export const database = new cratedb.Database("signals", {
    name: "signals",
    retentionPeriod: 604800,
});

// Get the id of the new bucket as an output
export const databaseId = database.id;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import komminarlabs_cratedb as cratedb

database = cratedb.Database(
    "signals",
    name="signals",
    retention_period=604800,
)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/komminarlabs/pulumi-cratedb/sdk/go/cratedb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		signals, err := cratedb.NewDatabase(ctx, "signals", &cratedb.DatabaseArgs{
			Name:            pulumi.String("signals"),
			RetentionPeriod: pulumi.Int(604800),
		})
		if err != nil {
			return err
		}

		ctx.Export("databaseId", signals.ID())
		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}
