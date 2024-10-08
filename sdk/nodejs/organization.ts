// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages an organization.
 */
export class Organization extends pulumi.CustomResource {
    /**
     * Get an existing Organization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationState, opts?: pulumi.CustomResourceOptions): Organization {
        return new Organization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'cratedb:index/organization:Organization';

    /**
     * Returns true if the given object is an instance of Organization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Organization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Organization.__pulumiType;
    }

    /**
     * The DublinCore of the organization.
     */
    public /*out*/ readonly dc!: pulumi.Output<outputs.OrganizationDc>;
    /**
     * The notification email used in the organization.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * The name of the organization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether notifications enabled for the organization.
     */
    public /*out*/ readonly notificationsEnabled!: pulumi.Output<boolean>;
    /**
     * The support plan type used in the organization.
     */
    public /*out*/ readonly planType!: pulumi.Output<number>;
    /**
     * The project count in the organization.
     */
    public /*out*/ readonly projectCount!: pulumi.Output<number>;
    /**
     * The role FQN.
     */
    public /*out*/ readonly roleFqn!: pulumi.Output<string>;

    /**
     * Create a Organization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrganizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationArgs | OrganizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationState | undefined;
            resourceInputs["dc"] = state ? state.dc : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationsEnabled"] = state ? state.notificationsEnabled : undefined;
            resourceInputs["planType"] = state ? state.planType : undefined;
            resourceInputs["projectCount"] = state ? state.projectCount : undefined;
            resourceInputs["roleFqn"] = state ? state.roleFqn : undefined;
        } else {
            const args = argsOrState as OrganizationArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["dc"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["notificationsEnabled"] = undefined /*out*/;
            resourceInputs["planType"] = undefined /*out*/;
            resourceInputs["projectCount"] = undefined /*out*/;
            resourceInputs["roleFqn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Organization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Organization resources.
 */
export interface OrganizationState {
    /**
     * The DublinCore of the organization.
     */
    dc?: pulumi.Input<inputs.OrganizationDc>;
    /**
     * The notification email used in the organization.
     */
    email?: pulumi.Input<string>;
    /**
     * The name of the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether notifications enabled for the organization.
     */
    notificationsEnabled?: pulumi.Input<boolean>;
    /**
     * The support plan type used in the organization.
     */
    planType?: pulumi.Input<number>;
    /**
     * The project count in the organization.
     */
    projectCount?: pulumi.Input<number>;
    /**
     * The role FQN.
     */
    roleFqn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Organization resource.
 */
export interface OrganizationArgs {
    /**
     * The name of the organization.
     */
    name?: pulumi.Input<string>;
}
