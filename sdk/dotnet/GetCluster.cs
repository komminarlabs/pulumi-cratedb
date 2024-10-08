// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace KomminarLabs.cratedb
{
    public static class GetCluster
    {
        /// <summary>
        /// To retrieve a cluster.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("cratedb:index/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// To retrieve a cluster.
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("cratedb:index/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the last async operation.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the last async operation.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// The allow custom storage flag.
        /// </summary>
        public readonly bool AllowCustomStorage;
        /// <summary>
        /// The allow suspend flag.
        /// </summary>
        public readonly bool AllowSuspend;
        /// <summary>
        /// The backup schedule.
        /// </summary>
        public readonly string BackupSchedule;
        /// <summary>
        /// The channel of the cluster.
        /// </summary>
        public readonly string Channel;
        /// <summary>
        /// The CrateDB version of the cluster.
        /// </summary>
        public readonly string CrateVersion;
        /// <summary>
        /// The DublinCore of the cluster.
        /// </summary>
        public readonly Outputs.GetClusterDcResult Dc;
        /// <summary>
        /// The deletion protected flag.
        /// </summary>
        public readonly bool DeletionProtected;
        /// <summary>
        /// The external IP address.
        /// </summary>
        public readonly string ExternalIp;
        /// <summary>
        /// The Fully Qualified Domain Name.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The garbage collection available flag.
        /// </summary>
        public readonly bool GcAvailable;
        /// <summary>
        /// The hardware specs of the cluster.
        /// </summary>
        public readonly Outputs.GetClusterHardwareSpecsResult HardwareSpecs;
        /// <summary>
        /// The health of the cluster.
        /// </summary>
        public readonly Outputs.GetClusterHealthResult Health;
        /// <summary>
        /// The id of the cluster.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP whitelist of the cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClusterIpWhitelistResult> IpWhitelists;
        /// <summary>
        /// The last async operation of the cluster.
        /// </summary>
        public readonly Outputs.GetClusterLastAsyncOperationResult LastAsyncOperation;
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of nodes in the cluster.
        /// </summary>
        public readonly double NumNodes;
        /// <summary>
        /// The origin of the cluster.
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// The password of the cluster.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The product name of the cluster.
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// The product tier of the cluster.
        /// </summary>
        public readonly string ProductTier;
        /// <summary>
        /// The product unit of the cluster.
        /// </summary>
        public readonly double ProductUnit;
        /// <summary>
        /// The project id of the cluster.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The subscription id of the cluster.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// The suspended flag.
        /// </summary>
        public readonly bool Suspended;
        /// <summary>
        /// The URL of the cluster.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The username of the cluster.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetClusterResult(
            bool allowCustomStorage,

            bool allowSuspend,

            string backupSchedule,

            string channel,

            string crateVersion,

            Outputs.GetClusterDcResult dc,

            bool deletionProtected,

            string externalIp,

            string fqdn,

            bool gcAvailable,

            Outputs.GetClusterHardwareSpecsResult hardwareSpecs,

            Outputs.GetClusterHealthResult health,

            string id,

            ImmutableArray<Outputs.GetClusterIpWhitelistResult> ipWhitelists,

            Outputs.GetClusterLastAsyncOperationResult lastAsyncOperation,

            string name,

            double numNodes,

            string origin,

            string password,

            string productName,

            string productTier,

            double productUnit,

            string projectId,

            string subscriptionId,

            bool suspended,

            string url,

            string username)
        {
            AllowCustomStorage = allowCustomStorage;
            AllowSuspend = allowSuspend;
            BackupSchedule = backupSchedule;
            Channel = channel;
            CrateVersion = crateVersion;
            Dc = dc;
            DeletionProtected = deletionProtected;
            ExternalIp = externalIp;
            Fqdn = fqdn;
            GcAvailable = gcAvailable;
            HardwareSpecs = hardwareSpecs;
            Health = health;
            Id = id;
            IpWhitelists = ipWhitelists;
            LastAsyncOperation = lastAsyncOperation;
            Name = name;
            NumNodes = numNodes;
            Origin = origin;
            Password = password;
            ProductName = productName;
            ProductTier = productTier;
            ProductUnit = productUnit;
            ProjectId = projectId;
            SubscriptionId = subscriptionId;
            Suspended = suspended;
            Url = url;
            Username = username;
        }
    }
}
