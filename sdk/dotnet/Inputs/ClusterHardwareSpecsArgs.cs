// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace KomminarLabs.CrateDB.Inputs
{

    public sealed class ClusterHardwareSpecsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cpus per node.
        /// </summary>
        [Input("cpusPerNode")]
        public Input<double>? CpusPerNode { get; set; }

        /// <summary>
        /// The disk size per node in bytes.
        /// </summary>
        [Input("diskSizePerNodeBytes")]
        public Input<int>? DiskSizePerNodeBytes { get; set; }

        /// <summary>
        /// The disk type.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The disks per node.
        /// </summary>
        [Input("disksPerNode")]
        public Input<double>? DisksPerNode { get; set; }

        /// <summary>
        /// The heap size in bytes.
        /// </summary>
        [Input("heapSizeBytes")]
        public Input<int>? HeapSizeBytes { get; set; }

        /// <summary>
        /// The memory per node in bytes.
        /// </summary>
        [Input("memoryPerNodeBytes")]
        public Input<int>? MemoryPerNodeBytes { get; set; }

        public ClusterHardwareSpecsArgs()
        {
        }
        public static new ClusterHardwareSpecsArgs Empty => new ClusterHardwareSpecsArgs();
    }
}
