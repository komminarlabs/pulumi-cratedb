// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * To retrieve a project.
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("cratedb:index/getProject:getProject", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The id of the project.
     */
    id: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * The DublinCore of the project.
     */
    readonly dc: outputs.GetProjectDc;
    /**
     * The id of the project.
     */
    readonly id: string;
    /**
     * The name of the project.
     */
    readonly name: string;
    /**
     * The organization id of the project.
     */
    readonly organizationId: string;
    /**
     * The region of the project.
     */
    readonly region: string;
}
/**
 * To retrieve a project.
 */
export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * The id of the project.
     */
    id: pulumi.Input<string>;
}
