// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace KomminarLabs.CrateDB
{
    /// <summary>
    /// Creates and manages an organization.
    /// </summary>
    [CrateDBResourceType("cratedb:index/organization:Organization")]
    public partial class Organization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The DublinCore of the organization.
        /// </summary>
        [Output("dc")]
        public Output<Outputs.OrganizationDc> Dc { get; private set; } = null!;

        /// <summary>
        /// The notification email used in the organization.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The name of the organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether notifications enabled for the organization.
        /// </summary>
        [Output("notificationsEnabled")]
        public Output<bool> NotificationsEnabled { get; private set; } = null!;

        /// <summary>
        /// The support plan type used in the organization.
        /// </summary>
        [Output("planType")]
        public Output<int> PlanType { get; private set; } = null!;

        /// <summary>
        /// The project count in the organization.
        /// </summary>
        [Output("projectCount")]
        public Output<int> ProjectCount { get; private set; } = null!;

        /// <summary>
        /// The role FQN.
        /// </summary>
        [Output("roleFqn")]
        public Output<string> RoleFqn { get; private set; } = null!;


        /// <summary>
        /// Create a Organization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Organization(string name, OrganizationArgs? args = null, CustomResourceOptions? options = null)
            : base("cratedb:index/organization:Organization", name, args ?? new OrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Organization(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
            : base("cratedb:index/organization:Organization", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/komminarlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Organization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Organization Get(string name, Input<string> id, OrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new Organization(name, id, state, options);
        }
    }

    public sealed class OrganizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OrganizationArgs()
        {
        }
        public static new OrganizationArgs Empty => new OrganizationArgs();
    }

    public sealed class OrganizationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DublinCore of the organization.
        /// </summary>
        [Input("dc")]
        public Input<Inputs.OrganizationDcGetArgs>? Dc { get; set; }

        /// <summary>
        /// The notification email used in the organization.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The name of the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether notifications enabled for the organization.
        /// </summary>
        [Input("notificationsEnabled")]
        public Input<bool>? NotificationsEnabled { get; set; }

        /// <summary>
        /// The support plan type used in the organization.
        /// </summary>
        [Input("planType")]
        public Input<int>? PlanType { get; set; }

        /// <summary>
        /// The project count in the organization.
        /// </summary>
        [Input("projectCount")]
        public Input<int>? ProjectCount { get; set; }

        /// <summary>
        /// The role FQN.
        /// </summary>
        [Input("roleFqn")]
        public Input<string>? RoleFqn { get; set; }

        public OrganizationState()
        {
        }
        public static new OrganizationState Empty => new OrganizationState();
    }
}
