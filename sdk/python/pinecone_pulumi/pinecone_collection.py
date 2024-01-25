# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PineconeCollectionArgs', 'PineconeCollection']

@pulumi.input_type
class PineconeCollectionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 source: pulumi.Input[str]):
        """
        The set of arguments for constructing a PineconeCollection resource.
        :param pulumi.Input[str] name: The name of the collection to be created.
        :param pulumi.Input[str] source: The name of the index to be used as the source for the collection.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the collection to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        The name of the index to be used as the source for the collection.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)


class PineconeCollection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PineconeCollection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the collection to be created.
        :param pulumi.Input[str] source: The name of the index to be used as the source for the collection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PineconeCollectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PineconeCollection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PineconeCollectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PineconeCollectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PineconeCollectionArgs.__new__(PineconeCollectionArgs)

            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["dimension"] = None
            __props__.__dict__["environment"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["vector_count"] = None
        super(PineconeCollection, __self__).__init__(
            'pinecone:index:PineconeCollection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PineconeCollection':
        """
        Get an existing PineconeCollection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PineconeCollectionArgs.__new__(PineconeCollectionArgs)

        __props__.__dict__["dimension"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["source"] = None
        __props__.__dict__["vector_count"] = None
        return PineconeCollection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dimension(self) -> pulumi.Output[int]:
        """
        The dimension of the vectors stored in each record held in the collection.
        """
        return pulumi.get(self, "dimension")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The environment where the collection is hosted.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the collection to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The size of the collection in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The name of the index to be used as the source for the collection.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="vectorCount")
    def vector_count(self) -> pulumi.Output[int]:
        """
        The number of records stored in the collection.
        """
        return pulumi.get(self, "vector_count")

