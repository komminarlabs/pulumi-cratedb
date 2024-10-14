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
    public static class GetOrganizations
    {
        /// <summary>
        /// To retrieve all organizations.
        /// </summary>
        public static Task<GetOrganizationsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationsResult>("cratedb:index/getOrganizations:getOrganizations", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// To retrieve all organizations.
        /// </summary>
        public static Output<GetOrganizationsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationsResult>("cratedb:index/getOrganizations:getOrganizations", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetOrganizationsOrganizationResult> Organizations;

        [OutputConstructor]
        private GetOrganizationsResult(
            string id,

            ImmutableArray<Outputs.GetOrganizationsOrganizationResult> organizations)
        {
            Id = id;
            Organizations = organizations;
        }
    }
}
