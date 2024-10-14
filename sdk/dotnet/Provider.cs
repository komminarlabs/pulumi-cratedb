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
    /// The provider type for the cratedb package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [CrateDBResourceType("pulumi:providers:cratedb")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The API key
        /// </summary>
        [Output("apiKey")]
        public Output<string?> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The API secret
        /// </summary>
        [Output("apiSecret")]
        public Output<string?> ApiSecret { get; private set; } = null!;

        /// <summary>
        /// The CrateDB Cloud URL
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("cratedb", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
                    "apiKey",
                    "apiSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiKey")]
        private Input<string>? _apiKey;

        /// <summary>
        /// The API key
        /// </summary>
        public Input<string>? ApiKey
        {
            get => _apiKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("apiSecret")]
        private Input<string>? _apiSecret;

        /// <summary>
        /// The API secret
        /// </summary>
        public Input<string>? ApiSecret
        {
            get => _apiSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The CrateDB Cloud URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
