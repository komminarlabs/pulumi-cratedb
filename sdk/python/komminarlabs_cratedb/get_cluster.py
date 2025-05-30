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
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
    'get_cluster_output',
]

@pulumi.output_type
class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, allow_custom_storage=None, allow_suspend=None, backup_schedule=None, channel=None, crate_version=None, dc=None, deletion_protected=None, external_ip=None, fqdn=None, gc_available=None, hardware_specs=None, health=None, id=None, ip_whitelists=None, last_async_operation=None, name=None, num_nodes=None, origin=None, password=None, product_name=None, product_tier=None, product_unit=None, project_id=None, subscription_id=None, suspended=None, url=None, username=None):
        if allow_custom_storage and not isinstance(allow_custom_storage, bool):
            raise TypeError("Expected argument 'allow_custom_storage' to be a bool")
        pulumi.set(__self__, "allow_custom_storage", allow_custom_storage)
        if allow_suspend and not isinstance(allow_suspend, bool):
            raise TypeError("Expected argument 'allow_suspend' to be a bool")
        pulumi.set(__self__, "allow_suspend", allow_suspend)
        if backup_schedule and not isinstance(backup_schedule, str):
            raise TypeError("Expected argument 'backup_schedule' to be a str")
        pulumi.set(__self__, "backup_schedule", backup_schedule)
        if channel and not isinstance(channel, str):
            raise TypeError("Expected argument 'channel' to be a str")
        pulumi.set(__self__, "channel", channel)
        if crate_version and not isinstance(crate_version, str):
            raise TypeError("Expected argument 'crate_version' to be a str")
        pulumi.set(__self__, "crate_version", crate_version)
        if dc and not isinstance(dc, dict):
            raise TypeError("Expected argument 'dc' to be a dict")
        pulumi.set(__self__, "dc", dc)
        if deletion_protected and not isinstance(deletion_protected, bool):
            raise TypeError("Expected argument 'deletion_protected' to be a bool")
        pulumi.set(__self__, "deletion_protected", deletion_protected)
        if external_ip and not isinstance(external_ip, str):
            raise TypeError("Expected argument 'external_ip' to be a str")
        pulumi.set(__self__, "external_ip", external_ip)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if gc_available and not isinstance(gc_available, bool):
            raise TypeError("Expected argument 'gc_available' to be a bool")
        pulumi.set(__self__, "gc_available", gc_available)
        if hardware_specs and not isinstance(hardware_specs, dict):
            raise TypeError("Expected argument 'hardware_specs' to be a dict")
        pulumi.set(__self__, "hardware_specs", hardware_specs)
        if health and not isinstance(health, dict):
            raise TypeError("Expected argument 'health' to be a dict")
        pulumi.set(__self__, "health", health)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_whitelists and not isinstance(ip_whitelists, list):
            raise TypeError("Expected argument 'ip_whitelists' to be a list")
        pulumi.set(__self__, "ip_whitelists", ip_whitelists)
        if last_async_operation and not isinstance(last_async_operation, dict):
            raise TypeError("Expected argument 'last_async_operation' to be a dict")
        pulumi.set(__self__, "last_async_operation", last_async_operation)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if num_nodes and not isinstance(num_nodes, int):
            raise TypeError("Expected argument 'num_nodes' to be a int")
        pulumi.set(__self__, "num_nodes", num_nodes)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if product_name and not isinstance(product_name, str):
            raise TypeError("Expected argument 'product_name' to be a str")
        pulumi.set(__self__, "product_name", product_name)
        if product_tier and not isinstance(product_tier, str):
            raise TypeError("Expected argument 'product_tier' to be a str")
        pulumi.set(__self__, "product_tier", product_tier)
        if product_unit and not isinstance(product_unit, int):
            raise TypeError("Expected argument 'product_unit' to be a int")
        pulumi.set(__self__, "product_unit", product_unit)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        pulumi.set(__self__, "subscription_id", subscription_id)
        if suspended and not isinstance(suspended, bool):
            raise TypeError("Expected argument 'suspended' to be a bool")
        pulumi.set(__self__, "suspended", suspended)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="allowCustomStorage")
    def allow_custom_storage(self) -> bool:
        """
        The allow custom storage flag.
        """
        return pulumi.get(self, "allow_custom_storage")

    @property
    @pulumi.getter(name="allowSuspend")
    def allow_suspend(self) -> bool:
        """
        The allow suspend flag.
        """
        return pulumi.get(self, "allow_suspend")

    @property
    @pulumi.getter(name="backupSchedule")
    def backup_schedule(self) -> str:
        """
        The backup schedule.
        """
        return pulumi.get(self, "backup_schedule")

    @property
    @pulumi.getter
    def channel(self) -> str:
        """
        The channel of the cluster.
        """
        return pulumi.get(self, "channel")

    @property
    @pulumi.getter(name="crateVersion")
    def crate_version(self) -> str:
        """
        The CrateDB version of the cluster.
        """
        return pulumi.get(self, "crate_version")

    @property
    @pulumi.getter
    def dc(self) -> 'outputs.GetClusterDcResult':
        """
        The DublinCore of the cluster.
        """
        return pulumi.get(self, "dc")

    @property
    @pulumi.getter(name="deletionProtected")
    def deletion_protected(self) -> bool:
        """
        The deletion protected flag.
        """
        return pulumi.get(self, "deletion_protected")

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> str:
        """
        The external IP address.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        The Fully Qualified Domain Name.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="gcAvailable")
    def gc_available(self) -> bool:
        """
        The garbage collection available flag.
        """
        return pulumi.get(self, "gc_available")

    @property
    @pulumi.getter(name="hardwareSpecs")
    def hardware_specs(self) -> 'outputs.GetClusterHardwareSpecsResult':
        """
        The hardware specs of the cluster.
        """
        return pulumi.get(self, "hardware_specs")

    @property
    @pulumi.getter
    def health(self) -> 'outputs.GetClusterHealthResult':
        """
        The health of the cluster.
        """
        return pulumi.get(self, "health")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipWhitelists")
    def ip_whitelists(self) -> Sequence['outputs.GetClusterIpWhitelistResult']:
        """
        The IP whitelist of the cluster.
        """
        return pulumi.get(self, "ip_whitelists")

    @property
    @pulumi.getter(name="lastAsyncOperation")
    def last_async_operation(self) -> 'outputs.GetClusterLastAsyncOperationResult':
        """
        The last async operation of the cluster.
        """
        return pulumi.get(self, "last_async_operation")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numNodes")
    def num_nodes(self) -> int:
        """
        The number of nodes in the cluster.
        """
        return pulumi.get(self, "num_nodes")

    @property
    @pulumi.getter
    def origin(self) -> str:
        """
        The origin of the cluster.
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The password of the cluster.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="productName")
    def product_name(self) -> str:
        """
        The product name of the cluster.
        """
        return pulumi.get(self, "product_name")

    @property
    @pulumi.getter(name="productTier")
    def product_tier(self) -> str:
        """
        The product tier of the cluster.
        """
        return pulumi.get(self, "product_tier")

    @property
    @pulumi.getter(name="productUnit")
    def product_unit(self) -> int:
        """
        The product unit of the cluster.
        """
        return pulumi.get(self, "product_unit")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project id of the cluster.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> str:
        """
        The subscription id of the cluster.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter
    def suspended(self) -> bool:
        """
        The suspended flag.
        """
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL of the cluster.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username of the cluster.
        """
        return pulumi.get(self, "username")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            allow_custom_storage=self.allow_custom_storage,
            allow_suspend=self.allow_suspend,
            backup_schedule=self.backup_schedule,
            channel=self.channel,
            crate_version=self.crate_version,
            dc=self.dc,
            deletion_protected=self.deletion_protected,
            external_ip=self.external_ip,
            fqdn=self.fqdn,
            gc_available=self.gc_available,
            hardware_specs=self.hardware_specs,
            health=self.health,
            id=self.id,
            ip_whitelists=self.ip_whitelists,
            last_async_operation=self.last_async_operation,
            name=self.name,
            num_nodes=self.num_nodes,
            origin=self.origin,
            password=self.password,
            product_name=self.product_name,
            product_tier=self.product_tier,
            product_unit=self.product_unit,
            project_id=self.project_id,
            subscription_id=self.subscription_id,
            suspended=self.suspended,
            url=self.url,
            username=self.username)


def get_cluster(id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    To retrieve a cluster.


    :param str id: The id of the cluster.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('cratedb:index/getCluster:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        allow_custom_storage=pulumi.get(__ret__, 'allow_custom_storage'),
        allow_suspend=pulumi.get(__ret__, 'allow_suspend'),
        backup_schedule=pulumi.get(__ret__, 'backup_schedule'),
        channel=pulumi.get(__ret__, 'channel'),
        crate_version=pulumi.get(__ret__, 'crate_version'),
        dc=pulumi.get(__ret__, 'dc'),
        deletion_protected=pulumi.get(__ret__, 'deletion_protected'),
        external_ip=pulumi.get(__ret__, 'external_ip'),
        fqdn=pulumi.get(__ret__, 'fqdn'),
        gc_available=pulumi.get(__ret__, 'gc_available'),
        hardware_specs=pulumi.get(__ret__, 'hardware_specs'),
        health=pulumi.get(__ret__, 'health'),
        id=pulumi.get(__ret__, 'id'),
        ip_whitelists=pulumi.get(__ret__, 'ip_whitelists'),
        last_async_operation=pulumi.get(__ret__, 'last_async_operation'),
        name=pulumi.get(__ret__, 'name'),
        num_nodes=pulumi.get(__ret__, 'num_nodes'),
        origin=pulumi.get(__ret__, 'origin'),
        password=pulumi.get(__ret__, 'password'),
        product_name=pulumi.get(__ret__, 'product_name'),
        product_tier=pulumi.get(__ret__, 'product_tier'),
        product_unit=pulumi.get(__ret__, 'product_unit'),
        project_id=pulumi.get(__ret__, 'project_id'),
        subscription_id=pulumi.get(__ret__, 'subscription_id'),
        suspended=pulumi.get(__ret__, 'suspended'),
        url=pulumi.get(__ret__, 'url'),
        username=pulumi.get(__ret__, 'username'))
def get_cluster_output(id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetClusterResult]:
    """
    To retrieve a cluster.


    :param str id: The id of the cluster.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('cratedb:index/getCluster:getCluster', __args__, opts=opts, typ=GetClusterResult)
    return __ret__.apply(lambda __response__: GetClusterResult(
        allow_custom_storage=pulumi.get(__response__, 'allow_custom_storage'),
        allow_suspend=pulumi.get(__response__, 'allow_suspend'),
        backup_schedule=pulumi.get(__response__, 'backup_schedule'),
        channel=pulumi.get(__response__, 'channel'),
        crate_version=pulumi.get(__response__, 'crate_version'),
        dc=pulumi.get(__response__, 'dc'),
        deletion_protected=pulumi.get(__response__, 'deletion_protected'),
        external_ip=pulumi.get(__response__, 'external_ip'),
        fqdn=pulumi.get(__response__, 'fqdn'),
        gc_available=pulumi.get(__response__, 'gc_available'),
        hardware_specs=pulumi.get(__response__, 'hardware_specs'),
        health=pulumi.get(__response__, 'health'),
        id=pulumi.get(__response__, 'id'),
        ip_whitelists=pulumi.get(__response__, 'ip_whitelists'),
        last_async_operation=pulumi.get(__response__, 'last_async_operation'),
        name=pulumi.get(__response__, 'name'),
        num_nodes=pulumi.get(__response__, 'num_nodes'),
        origin=pulumi.get(__response__, 'origin'),
        password=pulumi.get(__response__, 'password'),
        product_name=pulumi.get(__response__, 'product_name'),
        product_tier=pulumi.get(__response__, 'product_tier'),
        product_unit=pulumi.get(__response__, 'product_unit'),
        project_id=pulumi.get(__response__, 'project_id'),
        subscription_id=pulumi.get(__response__, 'subscription_id'),
        suspended=pulumi.get(__response__, 'suspended'),
        url=pulumi.get(__response__, 'url'),
        username=pulumi.get(__response__, 'username')))
