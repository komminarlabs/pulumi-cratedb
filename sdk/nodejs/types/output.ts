// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ClusterDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface ClusterHardwareSpecs {
    /**
     * The cpus per node.
     */
    cpusPerNode: number;
    /**
     * The disk size per node in bytes.
     */
    diskSizePerNodeBytes: number;
    /**
     * The disk type.
     */
    diskType: string;
    /**
     * The disks per node.
     */
    disksPerNode: number;
    /**
     * The heap size in bytes.
     */
    heapSizeBytes: number;
    /**
     * The memory per node in bytes.
     */
    memoryPerNodeBytes: number;
}

export interface ClusterHealth {
    /**
     * The health status of the cluster.
     */
    status: string;
}

export interface ClusterIpWhitelist {
    /**
     * The CIDR.
     */
    cidr: string;
    /**
     * The description.
     */
    description: string;
}

export interface GetClusterDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface GetClusterHardwareSpecs {
    /**
     * The cpus per node.
     */
    cpusPerNode: number;
    /**
     * The disk size per node in bytes.
     */
    diskSizePerNodeBytes: number;
    /**
     * The disk type.
     */
    diskType: string;
    /**
     * The disks per node.
     */
    disksPerNode: number;
    /**
     * The heap size in bytes.
     */
    heapSizeBytes: number;
    /**
     * The memory per node in bytes.
     */
    memoryPerNodeBytes: number;
}

export interface GetClusterHealth {
    /**
     * The last seen time.
     */
    lastSeen: string;
    /**
     * The type of the currently running operation. Returns an empty string if there is no operation in progress.
     */
    runningOperation: string;
    /**
     * The health status of the cluster.
     */
    status: string;
}

export interface GetClusterIpWhitelist {
    /**
     * The CIDR.
     */
    cidr: string;
    /**
     * The description.
     */
    description: string;
}

export interface GetClusterLastAsyncOperation {
    /**
     * The DublinCore of the cluster.
     */
    dc: outputs.GetClusterLastAsyncOperationDc;
    /**
     * The id of the last async operation.
     */
    id: string;
    /**
     * The status of the last async operation.
     */
    status: string;
    /**
     * The type of the last async operation.
     */
    type: string;
}

export interface GetClusterLastAsyncOperationDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface GetOrganizationDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface GetOrganizationsOrganization {
    /**
     * The DublinCore of the organization.
     */
    dc: outputs.GetOrganizationsOrganizationDc;
    /**
     * The notification email used in the organization.
     */
    email: string;
    /**
     * The id of the organization.
     */
    id: string;
    /**
     * The name of the organization.
     */
    name: string;
    /**
     * Whether notifications enabled for the organization.
     */
    notificationsEnabled: boolean;
    /**
     * The support plan type used in the organization.
     */
    planType: number;
    /**
     * The project count in the organization.
     */
    projectCount: number;
    /**
     * The role FQN.
     */
    roleFqn: string;
}

export interface GetOrganizationsOrganizationDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface GetProjectDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface OrganizationDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}

export interface ProjectDc {
    /**
     * The created time.
     */
    created: string;
    /**
     * The modified time.
     */
    modified: string;
}
