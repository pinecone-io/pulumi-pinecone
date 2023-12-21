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
    'MetaDataConfig',
    'PineconePodSpec',
    'PineconeServerlessSpec',
    'PineconeSpec',
]

@pulumi.output_type
class MetaDataConfig(dict):
    def __init__(__self__, *,
                 indexed: Optional[Sequence[str]] = None):
        if indexed is not None:
            pulumi.set(__self__, "indexed", indexed)

    @property
    @pulumi.getter
    def indexed(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "indexed")


@pulumi.output_type
class PineconePodSpec(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "podType":
            suggest = "pod_type"
        elif key == "metaDataConfig":
            suggest = "meta_data_config"
        elif key == "sourceCollection":
            suggest = "source_collection"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PineconePodSpec. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PineconePodSpec.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PineconePodSpec.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 environment: str,
                 pod_type: str,
                 replicas: int,
                 meta_data_config: Optional['outputs.MetaDataConfig'] = None,
                 pods: Optional[int] = None,
                 shards: Optional[int] = None,
                 source_collection: Optional[str] = None):
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "pod_type", pod_type)
        pulumi.set(__self__, "replicas", replicas)
        if meta_data_config is not None:
            pulumi.set(__self__, "meta_data_config", meta_data_config)
        if pods is not None:
            pulumi.set(__self__, "pods", pods)
        if shards is not None:
            pulumi.set(__self__, "shards", shards)
        if source_collection is not None:
            pulumi.set(__self__, "source_collection", source_collection)

    @property
    @pulumi.getter
    def environment(self) -> str:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="podType")
    def pod_type(self) -> str:
        return pulumi.get(self, "pod_type")

    @property
    @pulumi.getter
    def replicas(self) -> int:
        return pulumi.get(self, "replicas")

    @property
    @pulumi.getter(name="metaDataConfig")
    def meta_data_config(self) -> Optional['outputs.MetaDataConfig']:
        return pulumi.get(self, "meta_data_config")

    @property
    @pulumi.getter
    def pods(self) -> Optional[int]:
        return pulumi.get(self, "pods")

    @property
    @pulumi.getter
    def shards(self) -> Optional[int]:
        return pulumi.get(self, "shards")

    @property
    @pulumi.getter(name="sourceCollection")
    def source_collection(self) -> Optional[str]:
        return pulumi.get(self, "source_collection")


@pulumi.output_type
class PineconeServerlessSpec(dict):
    def __init__(__self__, *,
                 cloud: 'ServerlessSpecCloud',
                 region: str):
        pulumi.set(__self__, "cloud", cloud)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def cloud(self) -> 'ServerlessSpecCloud':
        return pulumi.get(self, "cloud")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


@pulumi.output_type
class PineconeSpec(dict):
    def __init__(__self__, *,
                 pod: Optional['outputs.PineconePodSpec'] = None,
                 serverless: Optional['outputs.PineconeServerlessSpec'] = None):
        if pod is not None:
            pulumi.set(__self__, "pod", pod)
        if serverless is not None:
            pulumi.set(__self__, "serverless", serverless)

    @property
    @pulumi.getter
    def pod(self) -> Optional['outputs.PineconePodSpec']:
        return pulumi.get(self, "pod")

    @property
    @pulumi.getter
    def serverless(self) -> Optional['outputs.PineconeServerlessSpec']:
        return pulumi.get(self, "serverless")


