// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace KomminarLabs.cratedb.Inputs
{

    public sealed class ProjectDcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The created time.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The modified time.
        /// </summary>
        [Input("modified")]
        public Input<string>? Modified { get; set; }

        public ProjectDcArgs()
        {
        }
        public static new ProjectDcArgs Empty => new ProjectDcArgs();
    }
}
