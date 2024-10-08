"""A Python Pulumi program"""

import komminarlabs_cratedb as cratedb

organization = cratedb.Organization(
    "default",
    name="default",
)
