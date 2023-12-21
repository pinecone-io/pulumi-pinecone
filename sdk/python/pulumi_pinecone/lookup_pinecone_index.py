# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'LookupPineconeIndexResult',
    'AwaitableLookupPineconeIndexResult',
    'lookup_pinecone_index',
    'lookup_pinecone_index_output',
]

@pulumi.output_type
class LookupPineconeIndexResult:
    """
    The result of a get operation on a Pinecone index.
    """
    def __init__(__self__, dimension=None, host=None, metric=None, name=None, spec=None, status=None):
        if dimension and not isinstance(dimension, int):
            raise TypeError("Expected argument 'dimension' to be a int")
        pulumi.set(__self__, "dimension", dimension)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if metric and not isinstance(metric, str):
            raise TypeError("Expected argument 'metric' to be a str")
        pulumi.set(__self__, "metric", metric)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if spec and not isinstance(spec, dict):
            raise TypeError("Expected argument 'spec' to be a dict")
        pulumi.set(__self__, "spec", spec)
        if status and not isinstance(status, bool):
            raise TypeError("Expected argument 'status' to be a bool")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def dimension(self) -> int:
        """
        The dimensions of the vectors in the index.
        """
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter
    def host(self) -> str:
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def metric(self) -> 'IndexMetric':
        """
        The metric used to compute the distance between vectors.
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Pinecone index.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.PineconeSpec':
        """
        Describe how the index should be deployed.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> bool:
        return pulumi.get(self, "status")


class AwaitableLookupPineconeIndexResult(LookupPineconeIndexResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupPineconeIndexResult(
            dimension=self.dimension,
            host=self.host,
            metric=self.metric,
            name=self.name,
            spec=self.spec,
            status=self.status)


def lookup_pinecone_index(name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLookupPineconeIndexResult:
    """
    The result of a get operation on a Pinecone index.


    :param str name: The name of the Pinecone index.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('pinecone:index:lookupPineconeIndex', __args__, opts=opts, typ=LookupPineconeIndexResult).value

    return AwaitableLookupPineconeIndexResult(
        dimension=pulumi.get(__ret__, 'dimension'),
        host=pulumi.get(__ret__, 'host'),
        metric=pulumi.get(__ret__, 'metric'),
        name=pulumi.get(__ret__, 'name'),
        spec=pulumi.get(__ret__, 'spec'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(lookup_pinecone_index)
def lookup_pinecone_index_output(name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LookupPineconeIndexResult]:
    """
    The result of a get operation on a Pinecone index.


    :param str name: The name of the Pinecone index.
    """
    ...
