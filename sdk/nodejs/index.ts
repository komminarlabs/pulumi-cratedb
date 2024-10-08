// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetOrganizationArgs, GetOrganizationResult, GetOrganizationOutputArgs } from "./getOrganization";
export const getOrganization: typeof import("./getOrganization").getOrganization = null as any;
export const getOrganizationOutput: typeof import("./getOrganization").getOrganizationOutput = null as any;
utilities.lazyLoad(exports, ["getOrganization","getOrganizationOutput"], () => require("./getOrganization"));

export { GetOrganizationsResult } from "./getOrganizations";
export const getOrganizations: typeof import("./getOrganizations").getOrganizations = null as any;
export const getOrganizationsOutput: typeof import("./getOrganizations").getOrganizationsOutput = null as any;
utilities.lazyLoad(exports, ["getOrganizations","getOrganizationsOutput"], () => require("./getOrganizations"));

export { GetProjectArgs, GetProjectResult, GetProjectOutputArgs } from "./getProject";
export const getProject: typeof import("./getProject").getProject = null as any;
export const getProjectOutput: typeof import("./getProject").getProjectOutput = null as any;
utilities.lazyLoad(exports, ["getProject","getProjectOutput"], () => require("./getProject"));

export { OrganizationArgs, OrganizationState } from "./organization";
export type Organization = import("./organization").Organization;
export const Organization: typeof import("./organization").Organization = null as any;
utilities.lazyLoad(exports, ["Organization"], () => require("./organization"));

export { ProjectArgs, ProjectState } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cratedb:index/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "cratedb:index/organization:Organization":
                return new Organization(name, <any>undefined, { urn })
            case "cratedb:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cratedb", "index/cluster", _module)
pulumi.runtime.registerResourceModule("cratedb", "index/organization", _module)
pulumi.runtime.registerResourceModule("cratedb", "index/project", _module)
pulumi.runtime.registerResourcePackage("cratedb", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:cratedb") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});