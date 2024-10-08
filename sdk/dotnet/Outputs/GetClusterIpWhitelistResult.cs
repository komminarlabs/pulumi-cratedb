// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace KomminarLabs.cratedb.Outputs
{

    [OutputType]
    public sealed class GetClusterIpWhitelistResult
    {
        /// <summary>
        /// The CIDR.
        /// </summary>
        public readonly string Cidr;
        /// <summary>
        /// The description.
        /// </summary>
        public readonly string Description;

        [OutputConstructor]
        private GetClusterIpWhitelistResult(
            string cidr,

            string description)
        {
            Cidr = cidr;
            Description = description;
        }
    }
}
