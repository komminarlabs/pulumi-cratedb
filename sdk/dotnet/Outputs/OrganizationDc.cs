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
    public sealed class OrganizationDc
    {
        /// <summary>
        /// The created time.
        /// </summary>
        public readonly string? Created;
        /// <summary>
        /// The modified time.
        /// </summary>
        public readonly string? Modified;

        [OutputConstructor]
        private OrganizationDc(
            string? created,

            string? modified)
        {
            Created = created;
            Modified = modified;
        }
    }
}