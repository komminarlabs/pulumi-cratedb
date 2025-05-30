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
    /// Creates and manages a cluster.
    /// </summary>
    [CrateDBResourceType("cratedb:index/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allow custom storage flag.
        /// </summary>
        [Output("allowCustomStorage")]
        public Output<bool> AllowCustomStorage { get; private set; } = null!;

        /// <summary>
        /// The allow suspend flag.
        /// </summary>
        [Output("allowSuspend")]
        public Output<bool> AllowSuspend { get; private set; } = null!;

        /// <summary>
        /// The backup schedule.
        /// </summary>
        [Output("backupSchedule")]
        public Output<string> BackupSchedule { get; private set; } = null!;

        /// <summary>
        /// The channel of the cluster. Default is `stable`.
        /// </summary>
        [Output("channel")]
        public Output<string> Channel { get; private set; } = null!;

        /// <summary>
        /// The CrateDB version of the cluster.
        /// </summary>
        [Output("crateVersion")]
        public Output<string> CrateVersion { get; private set; } = null!;

        /// <summary>
        /// The DublinCore of the cluster.
        /// </summary>
        [Output("dc")]
        public Output<Outputs.ClusterDc> Dc { get; private set; } = null!;

        /// <summary>
        /// The deletion protected flag.
        /// </summary>
        [Output("deletionProtected")]
        public Output<bool> DeletionProtected { get; private set; } = null!;

        /// <summary>
        /// The external IP address.
        /// </summary>
        [Output("externalIp")]
        public Output<string> ExternalIp { get; private set; } = null!;

        /// <summary>
        /// The Fully Qualified Domain Name.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The garbage collection available flag.
        /// </summary>
        [Output("gcAvailable")]
        public Output<bool> GcAvailable { get; private set; } = null!;

        /// <summary>
        /// The hardware specs of the cluster.
        /// </summary>
        [Output("hardwareSpecs")]
        public Output<Outputs.ClusterHardwareSpecs> HardwareSpecs { get; private set; } = null!;

        /// <summary>
        /// The health of the cluster.
        /// </summary>
        [Output("health")]
        public Output<Outputs.ClusterHealth> Health { get; private set; } = null!;

        /// <summary>
        /// The IP whitelist of the cluster.
        /// </summary>
        [Output("ipWhitelists")]
        public Output<ImmutableArray<Outputs.ClusterIpWhitelist>> IpWhitelists { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the cluster.
        /// </summary>
        [Output("numNodes")]
        public Output<int> NumNodes { get; private set; } = null!;

        /// <summary>
        /// The organization id of the cluster.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The origin of the cluster.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// The password of the cluster.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The product name of the cluster.
        /// </summary>
        [Output("productName")]
        public Output<string> ProductName { get; private set; } = null!;

        /// <summary>
        /// The product tier of the cluster.
        /// </summary>
        [Output("productTier")]
        public Output<string> ProductTier { get; private set; } = null!;

        /// <summary>
        /// The product unit of the cluster. Default is `0`.
        /// </summary>
        [Output("productUnit")]
        public Output<int> ProductUnit { get; private set; } = null!;

        /// <summary>
        /// The project id of the cluster.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The subscription id of the cluster.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// The suspended flag.
        /// </summary>
        [Output("suspended")]
        public Output<bool> Suspended { get; private set; } = null!;

        /// <summary>
        /// The URL of the cluster.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The username of the cluster.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("cratedb:index/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("cratedb:index/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/komminarlabs",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The channel of the cluster. Default is `stable`.
        /// </summary>
        [Input("channel")]
        public Input<string>? Channel { get; set; }

        /// <summary>
        /// The CrateDB version of the cluster.
        /// </summary>
        [Input("crateVersion", required: true)]
        public Input<string> CrateVersion { get; set; } = null!;

        /// <summary>
        /// The hardware specs of the cluster.
        /// </summary>
        [Input("hardwareSpecs")]
        public Input<Inputs.ClusterHardwareSpecsArgs>? HardwareSpecs { get; set; }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization id of the cluster.
        /// </summary>
        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password of the cluster.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The product name of the cluster.
        /// </summary>
        [Input("productName", required: true)]
        public Input<string> ProductName { get; set; } = null!;

        /// <summary>
        /// The product tier of the cluster.
        /// </summary>
        [Input("productTier", required: true)]
        public Input<string> ProductTier { get; set; } = null!;

        /// <summary>
        /// The product unit of the cluster. Default is `0`.
        /// </summary>
        [Input("productUnit")]
        public Input<int>? ProductUnit { get; set; }

        /// <summary>
        /// The project id of the cluster.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The subscription id of the cluster.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        /// <summary>
        /// The username of the cluster.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allow custom storage flag.
        /// </summary>
        [Input("allowCustomStorage")]
        public Input<bool>? AllowCustomStorage { get; set; }

        /// <summary>
        /// The allow suspend flag.
        /// </summary>
        [Input("allowSuspend")]
        public Input<bool>? AllowSuspend { get; set; }

        /// <summary>
        /// The backup schedule.
        /// </summary>
        [Input("backupSchedule")]
        public Input<string>? BackupSchedule { get; set; }

        /// <summary>
        /// The channel of the cluster. Default is `stable`.
        /// </summary>
        [Input("channel")]
        public Input<string>? Channel { get; set; }

        /// <summary>
        /// The CrateDB version of the cluster.
        /// </summary>
        [Input("crateVersion")]
        public Input<string>? CrateVersion { get; set; }

        /// <summary>
        /// The DublinCore of the cluster.
        /// </summary>
        [Input("dc")]
        public Input<Inputs.ClusterDcGetArgs>? Dc { get; set; }

        /// <summary>
        /// The deletion protected flag.
        /// </summary>
        [Input("deletionProtected")]
        public Input<bool>? DeletionProtected { get; set; }

        /// <summary>
        /// The external IP address.
        /// </summary>
        [Input("externalIp")]
        public Input<string>? ExternalIp { get; set; }

        /// <summary>
        /// The Fully Qualified Domain Name.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The garbage collection available flag.
        /// </summary>
        [Input("gcAvailable")]
        public Input<bool>? GcAvailable { get; set; }

        /// <summary>
        /// The hardware specs of the cluster.
        /// </summary>
        [Input("hardwareSpecs")]
        public Input<Inputs.ClusterHardwareSpecsGetArgs>? HardwareSpecs { get; set; }

        /// <summary>
        /// The health of the cluster.
        /// </summary>
        [Input("health")]
        public Input<Inputs.ClusterHealthGetArgs>? Health { get; set; }

        [Input("ipWhitelists")]
        private InputList<Inputs.ClusterIpWhitelistGetArgs>? _ipWhitelists;

        /// <summary>
        /// The IP whitelist of the cluster.
        /// </summary>
        public InputList<Inputs.ClusterIpWhitelistGetArgs> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<Inputs.ClusterIpWhitelistGetArgs>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of nodes in the cluster.
        /// </summary>
        [Input("numNodes")]
        public Input<int>? NumNodes { get; set; }

        /// <summary>
        /// The organization id of the cluster.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The origin of the cluster.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the cluster.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The product name of the cluster.
        /// </summary>
        [Input("productName")]
        public Input<string>? ProductName { get; set; }

        /// <summary>
        /// The product tier of the cluster.
        /// </summary>
        [Input("productTier")]
        public Input<string>? ProductTier { get; set; }

        /// <summary>
        /// The product unit of the cluster. Default is `0`.
        /// </summary>
        [Input("productUnit")]
        public Input<int>? ProductUnit { get; set; }

        /// <summary>
        /// The project id of the cluster.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The subscription id of the cluster.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        /// <summary>
        /// The suspended flag.
        /// </summary>
        [Input("suspended")]
        public Input<bool>? Suspended { get; set; }

        /// <summary>
        /// The URL of the cluster.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The username of the cluster.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
