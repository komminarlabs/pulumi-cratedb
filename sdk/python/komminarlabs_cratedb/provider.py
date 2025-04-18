# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_key: The API key
        :param pulumi.Input[str] api_secret: The API secret
        :param pulumi.Input[str] url: The CrateDB Cloud URL
        """
        if api_key is not None:
            pulumi.set(__self__, "api_key", api_key)
        if api_secret is not None:
            pulumi.set(__self__, "api_secret", api_secret)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[pulumi.Input[str]]:
        """
        The API key
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter(name="apiSecret")
    def api_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The API secret
        """
        return pulumi.get(self, "api_secret")

    @api_secret.setter
    def api_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_secret", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The CrateDB Cloud URL
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the cratedb package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: The API key
        :param pulumi.Input[str] api_secret: The API secret
        :param pulumi.Input[str] url: The CrateDB Cloud URL
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the cratedb package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 api_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["api_key"] = None if api_key is None else pulumi.Output.secret(api_key)
            __props__.__dict__["api_secret"] = None if api_secret is None else pulumi.Output.secret(api_secret)
            __props__.__dict__["url"] = url
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiKey", "apiSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'cratedb',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[Optional[str]]:
        """
        The API key
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="apiSecret")
    def api_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The API secret
        """
        return pulumi.get(self, "api_secret")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The CrateDB Cloud URL
        """
        return pulumi.get(self, "url")

