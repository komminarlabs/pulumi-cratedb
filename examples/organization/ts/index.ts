import * as cratedb from "@komminarlabs/cratedb";

export const organization = new cratedb.Organization("default", {
    name: "default",
});

export const organizationName = organization.name;

console.log(`Organization Name: {organizationName}`);
