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
    public sealed class GetClusterHardwareSpecsResult
    {
        /// <summary>
        /// The cpus per node.
        /// </summary>
        public readonly double CpusPerNode;
        /// <summary>
        /// The disk size per node in bytes.
        /// </summary>
        public readonly int DiskSizePerNodeBytes;
        /// <summary>
        /// The disk type.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// The disks per node.
        /// </summary>
        public readonly double DisksPerNode;
        /// <summary>
        /// The heap size in bytes.
        /// </summary>
        public readonly int HeapSizeBytes;
        /// <summary>
        /// The memory per node in bytes.
        /// </summary>
        public readonly int MemoryPerNodeBytes;

        [OutputConstructor]
        private GetClusterHardwareSpecsResult(
            double cpusPerNode,

            int diskSizePerNodeBytes,

            string diskType,

            double disksPerNode,

            int heapSizeBytes,

            int memoryPerNodeBytes)
        {
            CpusPerNode = cpusPerNode;
            DiskSizePerNodeBytes = diskSizePerNodeBytes;
            DiskType = diskType;
            DisksPerNode = disksPerNode;
            HeapSizeBytes = heapSizeBytes;
            MemoryPerNodeBytes = memoryPerNodeBytes;
        }
    }
}
