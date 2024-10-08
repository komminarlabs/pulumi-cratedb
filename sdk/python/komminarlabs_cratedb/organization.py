# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OrganizationArgs', 'Organization']

@pulumi.input_type
class OrganizationArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Organization resource.
        :param pulumi.Input[str] name: The name of the organization.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OrganizationState:
    def __init__(__self__, *,
                 dc: Optional[pulumi.Input['OrganizationDcArgs']] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications_enabled: Optional[pulumi.Input[bool]] = None,
                 plan_type: Optional[pulumi.Input[float]] = None,
                 project_count: Optional[pulumi.Input[float]] = None,
                 role_fqn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Organization resources.
        :param pulumi.Input['OrganizationDcArgs'] dc: The DublinCore of the organization.
        :param pulumi.Input[str] email: The notification email used in the organization.
        :param pulumi.Input[str] name: The name of the organization.
        :param pulumi.Input[bool] notifications_enabled: Whether notifications enabled for the organization.
        :param pulumi.Input[float] plan_type: The support plan type used in the organization.
        :param pulumi.Input[float] project_count: The project count in the organization.
        :param pulumi.Input[str] role_fqn: The role FQN.
        """
        if dc is not None:
            pulumi.set(__self__, "dc", dc)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications_enabled is not None:
            pulumi.set(__self__, "notifications_enabled", notifications_enabled)
        if plan_type is not None:
            pulumi.set(__self__, "plan_type", plan_type)
        if project_count is not None:
            pulumi.set(__self__, "project_count", project_count)
        if role_fqn is not None:
            pulumi.set(__self__, "role_fqn", role_fqn)

    @property
    @pulumi.getter
    def dc(self) -> Optional[pulumi.Input['OrganizationDcArgs']]:
        """
        The DublinCore of the organization.
        """
        return pulumi.get(self, "dc")

    @dc.setter
    def dc(self, value: Optional[pulumi.Input['OrganizationDcArgs']]):
        pulumi.set(self, "dc", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The notification email used in the organization.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationsEnabled")
    def notifications_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether notifications enabled for the organization.
        """
        return pulumi.get(self, "notifications_enabled")

    @notifications_enabled.setter
    def notifications_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notifications_enabled", value)

    @property
    @pulumi.getter(name="planType")
    def plan_type(self) -> Optional[pulumi.Input[float]]:
        """
        The support plan type used in the organization.
        """
        return pulumi.get(self, "plan_type")

    @plan_type.setter
    def plan_type(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "plan_type", value)

    @property
    @pulumi.getter(name="projectCount")
    def project_count(self) -> Optional[pulumi.Input[float]]:
        """
        The project count in the organization.
        """
        return pulumi.get(self, "project_count")

    @project_count.setter
    def project_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "project_count", value)

    @property
    @pulumi.getter(name="roleFqn")
    def role_fqn(self) -> Optional[pulumi.Input[str]]:
        """
        The role FQN.
        """
        return pulumi.get(self, "role_fqn")

    @role_fqn.setter
    def role_fqn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_fqn", value)


class Organization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages an organization.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the organization.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OrganizationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages an organization.

        :param str resource_name: The name of the resource.
        :param OrganizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationArgs.__new__(OrganizationArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["dc"] = None
            __props__.__dict__["email"] = None
            __props__.__dict__["notifications_enabled"] = None
            __props__.__dict__["plan_type"] = None
            __props__.__dict__["project_count"] = None
            __props__.__dict__["role_fqn"] = None
        super(Organization, __self__).__init__(
            'cratedb:index/organization:Organization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dc: Optional[pulumi.Input[pulumi.InputType['OrganizationDcArgs']]] = None,
            email: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications_enabled: Optional[pulumi.Input[bool]] = None,
            plan_type: Optional[pulumi.Input[float]] = None,
            project_count: Optional[pulumi.Input[float]] = None,
            role_fqn: Optional[pulumi.Input[str]] = None) -> 'Organization':
        """
        Get an existing Organization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OrganizationDcArgs']] dc: The DublinCore of the organization.
        :param pulumi.Input[str] email: The notification email used in the organization.
        :param pulumi.Input[str] name: The name of the organization.
        :param pulumi.Input[bool] notifications_enabled: Whether notifications enabled for the organization.
        :param pulumi.Input[float] plan_type: The support plan type used in the organization.
        :param pulumi.Input[float] project_count: The project count in the organization.
        :param pulumi.Input[str] role_fqn: The role FQN.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationState.__new__(_OrganizationState)

        __props__.__dict__["dc"] = dc
        __props__.__dict__["email"] = email
        __props__.__dict__["name"] = name
        __props__.__dict__["notifications_enabled"] = notifications_enabled
        __props__.__dict__["plan_type"] = plan_type
        __props__.__dict__["project_count"] = project_count
        __props__.__dict__["role_fqn"] = role_fqn
        return Organization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dc(self) -> pulumi.Output['outputs.OrganizationDc']:
        """
        The DublinCore of the organization.
        """
        return pulumi.get(self, "dc")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The notification email used in the organization.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationsEnabled")
    def notifications_enabled(self) -> pulumi.Output[bool]:
        """
        Whether notifications enabled for the organization.
        """
        return pulumi.get(self, "notifications_enabled")

    @property
    @pulumi.getter(name="planType")
    def plan_type(self) -> pulumi.Output[float]:
        """
        The support plan type used in the organization.
        """
        return pulumi.get(self, "plan_type")

    @property
    @pulumi.getter(name="projectCount")
    def project_count(self) -> pulumi.Output[float]:
        """
        The project count in the organization.
        """
        return pulumi.get(self, "project_count")

    @property
    @pulumi.getter(name="roleFqn")
    def role_fqn(self) -> pulumi.Output[str]:
        """
        The role FQN.
        """
        return pulumi.get(self, "role_fqn")
