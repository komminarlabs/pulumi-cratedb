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
from . import outputs

__all__ = [
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, dc=None, id=None, name=None, organization_id=None, region=None):
        if dc and not isinstance(dc, dict):
            raise TypeError("Expected argument 'dc' to be a dict")
        pulumi.set(__self__, "dc", dc)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def dc(self) -> 'outputs.GetProjectDcResult':
        """
        The DublinCore of the project.
        """
        return pulumi.get(self, "dc")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the project.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The organization id of the project.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region of the project.
        """
        return pulumi.get(self, "region")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            dc=self.dc,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            region=self.region)


def get_project(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    To retrieve a project.


    :param str id: The id of the project.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cratedb:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        dc=pulumi.get(__ret__, 'dc'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        region=pulumi.get(__ret__, 'region'))
def get_project_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectResult]:
    """
    To retrieve a project.


    :param str id: The id of the project.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cratedb:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult)
    return __ret__.apply(lambda __response__: GetProjectResult(
        dc=pulumi.get(__response__, 'dc'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        region=pulumi.get(__response__, 'region')))
