// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinecone

import (
	"context"
	"reflect"

	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type MetaDataConfig struct {
	Indexed []string `pulumi:"indexed"`
}

// MetaDataConfigInput is an input type that accepts MetaDataConfigArgs and MetaDataConfigOutput values.
// You can construct a concrete instance of `MetaDataConfigInput` via:
//
//	MetaDataConfigArgs{...}
type MetaDataConfigInput interface {
	pulumi.Input

	ToMetaDataConfigOutput() MetaDataConfigOutput
	ToMetaDataConfigOutputWithContext(context.Context) MetaDataConfigOutput
}

type MetaDataConfigArgs struct {
	Indexed pulumi.StringArrayInput `pulumi:"indexed"`
}

func (MetaDataConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetaDataConfig)(nil)).Elem()
}

func (i MetaDataConfigArgs) ToMetaDataConfigOutput() MetaDataConfigOutput {
	return i.ToMetaDataConfigOutputWithContext(context.Background())
}

func (i MetaDataConfigArgs) ToMetaDataConfigOutputWithContext(ctx context.Context) MetaDataConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetaDataConfigOutput)
}

func (i MetaDataConfigArgs) ToMetaDataConfigPtrOutput() MetaDataConfigPtrOutput {
	return i.ToMetaDataConfigPtrOutputWithContext(context.Background())
}

func (i MetaDataConfigArgs) ToMetaDataConfigPtrOutputWithContext(ctx context.Context) MetaDataConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetaDataConfigOutput).ToMetaDataConfigPtrOutputWithContext(ctx)
}

// MetaDataConfigPtrInput is an input type that accepts MetaDataConfigArgs, MetaDataConfigPtr and MetaDataConfigPtrOutput values.
// You can construct a concrete instance of `MetaDataConfigPtrInput` via:
//
//	        MetaDataConfigArgs{...}
//
//	or:
//
//	        nil
type MetaDataConfigPtrInput interface {
	pulumi.Input

	ToMetaDataConfigPtrOutput() MetaDataConfigPtrOutput
	ToMetaDataConfigPtrOutputWithContext(context.Context) MetaDataConfigPtrOutput
}

type metaDataConfigPtrType MetaDataConfigArgs

func MetaDataConfigPtr(v *MetaDataConfigArgs) MetaDataConfigPtrInput {
	return (*metaDataConfigPtrType)(v)
}

func (*metaDataConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetaDataConfig)(nil)).Elem()
}

func (i *metaDataConfigPtrType) ToMetaDataConfigPtrOutput() MetaDataConfigPtrOutput {
	return i.ToMetaDataConfigPtrOutputWithContext(context.Background())
}

func (i *metaDataConfigPtrType) ToMetaDataConfigPtrOutputWithContext(ctx context.Context) MetaDataConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetaDataConfigPtrOutput)
}

type MetaDataConfigOutput struct{ *pulumi.OutputState }

func (MetaDataConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetaDataConfig)(nil)).Elem()
}

func (o MetaDataConfigOutput) ToMetaDataConfigOutput() MetaDataConfigOutput {
	return o
}

func (o MetaDataConfigOutput) ToMetaDataConfigOutputWithContext(ctx context.Context) MetaDataConfigOutput {
	return o
}

func (o MetaDataConfigOutput) ToMetaDataConfigPtrOutput() MetaDataConfigPtrOutput {
	return o.ToMetaDataConfigPtrOutputWithContext(context.Background())
}

func (o MetaDataConfigOutput) ToMetaDataConfigPtrOutputWithContext(ctx context.Context) MetaDataConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetaDataConfig) *MetaDataConfig {
		return &v
	}).(MetaDataConfigPtrOutput)
}

func (o MetaDataConfigOutput) Indexed() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetaDataConfig) []string { return v.Indexed }).(pulumi.StringArrayOutput)
}

type MetaDataConfigPtrOutput struct{ *pulumi.OutputState }

func (MetaDataConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetaDataConfig)(nil)).Elem()
}

func (o MetaDataConfigPtrOutput) ToMetaDataConfigPtrOutput() MetaDataConfigPtrOutput {
	return o
}

func (o MetaDataConfigPtrOutput) ToMetaDataConfigPtrOutputWithContext(ctx context.Context) MetaDataConfigPtrOutput {
	return o
}

func (o MetaDataConfigPtrOutput) Elem() MetaDataConfigOutput {
	return o.ApplyT(func(v *MetaDataConfig) MetaDataConfig {
		if v != nil {
			return *v
		}
		var ret MetaDataConfig
		return ret
	}).(MetaDataConfigOutput)
}

func (o MetaDataConfigPtrOutput) Indexed() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetaDataConfig) []string {
		if v == nil {
			return nil
		}
		return v.Indexed
	}).(pulumi.StringArrayOutput)
}

type PineconePodSpec struct {
	Environment      string          `pulumi:"environment"`
	MetaDataConfig   *MetaDataConfig `pulumi:"metaDataConfig"`
	PodType          string          `pulumi:"podType"`
	Pods             *int            `pulumi:"pods"`
	Replicas         int             `pulumi:"replicas"`
	Shards           *int            `pulumi:"shards"`
	SourceCollection *string         `pulumi:"sourceCollection"`
}

// PineconePodSpecInput is an input type that accepts PineconePodSpecArgs and PineconePodSpecOutput values.
// You can construct a concrete instance of `PineconePodSpecInput` via:
//
//	PineconePodSpecArgs{...}
type PineconePodSpecInput interface {
	pulumi.Input

	ToPineconePodSpecOutput() PineconePodSpecOutput
	ToPineconePodSpecOutputWithContext(context.Context) PineconePodSpecOutput
}

type PineconePodSpecArgs struct {
	Environment      pulumi.StringInput     `pulumi:"environment"`
	MetaDataConfig   MetaDataConfigPtrInput `pulumi:"metaDataConfig"`
	PodType          pulumi.StringInput     `pulumi:"podType"`
	Pods             pulumi.IntPtrInput     `pulumi:"pods"`
	Replicas         pulumi.IntInput        `pulumi:"replicas"`
	Shards           pulumi.IntPtrInput     `pulumi:"shards"`
	SourceCollection pulumi.StringPtrInput  `pulumi:"sourceCollection"`
}

func (PineconePodSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconePodSpec)(nil)).Elem()
}

func (i PineconePodSpecArgs) ToPineconePodSpecOutput() PineconePodSpecOutput {
	return i.ToPineconePodSpecOutputWithContext(context.Background())
}

func (i PineconePodSpecArgs) ToPineconePodSpecOutputWithContext(ctx context.Context) PineconePodSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconePodSpecOutput)
}

func (i PineconePodSpecArgs) ToPineconePodSpecPtrOutput() PineconePodSpecPtrOutput {
	return i.ToPineconePodSpecPtrOutputWithContext(context.Background())
}

func (i PineconePodSpecArgs) ToPineconePodSpecPtrOutputWithContext(ctx context.Context) PineconePodSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconePodSpecOutput).ToPineconePodSpecPtrOutputWithContext(ctx)
}

// PineconePodSpecPtrInput is an input type that accepts PineconePodSpecArgs, PineconePodSpecPtr and PineconePodSpecPtrOutput values.
// You can construct a concrete instance of `PineconePodSpecPtrInput` via:
//
//	        PineconePodSpecArgs{...}
//
//	or:
//
//	        nil
type PineconePodSpecPtrInput interface {
	pulumi.Input

	ToPineconePodSpecPtrOutput() PineconePodSpecPtrOutput
	ToPineconePodSpecPtrOutputWithContext(context.Context) PineconePodSpecPtrOutput
}

type pineconePodSpecPtrType PineconePodSpecArgs

func PineconePodSpecPtr(v *PineconePodSpecArgs) PineconePodSpecPtrInput {
	return (*pineconePodSpecPtrType)(v)
}

func (*pineconePodSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconePodSpec)(nil)).Elem()
}

func (i *pineconePodSpecPtrType) ToPineconePodSpecPtrOutput() PineconePodSpecPtrOutput {
	return i.ToPineconePodSpecPtrOutputWithContext(context.Background())
}

func (i *pineconePodSpecPtrType) ToPineconePodSpecPtrOutputWithContext(ctx context.Context) PineconePodSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconePodSpecPtrOutput)
}

type PineconePodSpecOutput struct{ *pulumi.OutputState }

func (PineconePodSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconePodSpec)(nil)).Elem()
}

func (o PineconePodSpecOutput) ToPineconePodSpecOutput() PineconePodSpecOutput {
	return o
}

func (o PineconePodSpecOutput) ToPineconePodSpecOutputWithContext(ctx context.Context) PineconePodSpecOutput {
	return o
}

func (o PineconePodSpecOutput) ToPineconePodSpecPtrOutput() PineconePodSpecPtrOutput {
	return o.ToPineconePodSpecPtrOutputWithContext(context.Background())
}

func (o PineconePodSpecOutput) ToPineconePodSpecPtrOutputWithContext(ctx context.Context) PineconePodSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PineconePodSpec) *PineconePodSpec {
		return &v
	}).(PineconePodSpecPtrOutput)
}

func (o PineconePodSpecOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v PineconePodSpec) string { return v.Environment }).(pulumi.StringOutput)
}

func (o PineconePodSpecOutput) MetaDataConfig() MetaDataConfigPtrOutput {
	return o.ApplyT(func(v PineconePodSpec) *MetaDataConfig { return v.MetaDataConfig }).(MetaDataConfigPtrOutput)
}

func (o PineconePodSpecOutput) PodType() pulumi.StringOutput {
	return o.ApplyT(func(v PineconePodSpec) string { return v.PodType }).(pulumi.StringOutput)
}

func (o PineconePodSpecOutput) Pods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PineconePodSpec) *int { return v.Pods }).(pulumi.IntPtrOutput)
}

func (o PineconePodSpecOutput) Replicas() pulumi.IntOutput {
	return o.ApplyT(func(v PineconePodSpec) int { return v.Replicas }).(pulumi.IntOutput)
}

func (o PineconePodSpecOutput) Shards() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PineconePodSpec) *int { return v.Shards }).(pulumi.IntPtrOutput)
}

func (o PineconePodSpecOutput) SourceCollection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PineconePodSpec) *string { return v.SourceCollection }).(pulumi.StringPtrOutput)
}

type PineconePodSpecPtrOutput struct{ *pulumi.OutputState }

func (PineconePodSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconePodSpec)(nil)).Elem()
}

func (o PineconePodSpecPtrOutput) ToPineconePodSpecPtrOutput() PineconePodSpecPtrOutput {
	return o
}

func (o PineconePodSpecPtrOutput) ToPineconePodSpecPtrOutputWithContext(ctx context.Context) PineconePodSpecPtrOutput {
	return o
}

func (o PineconePodSpecPtrOutput) Elem() PineconePodSpecOutput {
	return o.ApplyT(func(v *PineconePodSpec) PineconePodSpec {
		if v != nil {
			return *v
		}
		var ret PineconePodSpec
		return ret
	}).(PineconePodSpecOutput)
}

func (o PineconePodSpecPtrOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *string {
		if v == nil {
			return nil
		}
		return &v.Environment
	}).(pulumi.StringPtrOutput)
}

func (o PineconePodSpecPtrOutput) MetaDataConfig() MetaDataConfigPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *MetaDataConfig {
		if v == nil {
			return nil
		}
		return v.MetaDataConfig
	}).(MetaDataConfigPtrOutput)
}

func (o PineconePodSpecPtrOutput) PodType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *string {
		if v == nil {
			return nil
		}
		return &v.PodType
	}).(pulumi.StringPtrOutput)
}

func (o PineconePodSpecPtrOutput) Pods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *int {
		if v == nil {
			return nil
		}
		return v.Pods
	}).(pulumi.IntPtrOutput)
}

func (o PineconePodSpecPtrOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *int {
		if v == nil {
			return nil
		}
		return &v.Replicas
	}).(pulumi.IntPtrOutput)
}

func (o PineconePodSpecPtrOutput) Shards() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *int {
		if v == nil {
			return nil
		}
		return v.Shards
	}).(pulumi.IntPtrOutput)
}

func (o PineconePodSpecPtrOutput) SourceCollection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PineconePodSpec) *string {
		if v == nil {
			return nil
		}
		return v.SourceCollection
	}).(pulumi.StringPtrOutput)
}

type PineconeServerlessSpec struct {
	// he public cloud where you would like your index hosted
	Cloud ServerlessSpecCloud `pulumi:"cloud"`
	// The region where you would like your index to be created. Different cloud providers have different regions available.
	Region string `pulumi:"region"`
}

// PineconeServerlessSpecInput is an input type that accepts PineconeServerlessSpecArgs and PineconeServerlessSpecOutput values.
// You can construct a concrete instance of `PineconeServerlessSpecInput` via:
//
//	PineconeServerlessSpecArgs{...}
type PineconeServerlessSpecInput interface {
	pulumi.Input

	ToPineconeServerlessSpecOutput() PineconeServerlessSpecOutput
	ToPineconeServerlessSpecOutputWithContext(context.Context) PineconeServerlessSpecOutput
}

type PineconeServerlessSpecArgs struct {
	// he public cloud where you would like your index hosted
	Cloud ServerlessSpecCloudInput `pulumi:"cloud"`
	// The region where you would like your index to be created. Different cloud providers have different regions available.
	Region pulumi.StringInput `pulumi:"region"`
}

func (PineconeServerlessSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconeServerlessSpec)(nil)).Elem()
}

func (i PineconeServerlessSpecArgs) ToPineconeServerlessSpecOutput() PineconeServerlessSpecOutput {
	return i.ToPineconeServerlessSpecOutputWithContext(context.Background())
}

func (i PineconeServerlessSpecArgs) ToPineconeServerlessSpecOutputWithContext(ctx context.Context) PineconeServerlessSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeServerlessSpecOutput)
}

func (i PineconeServerlessSpecArgs) ToPineconeServerlessSpecPtrOutput() PineconeServerlessSpecPtrOutput {
	return i.ToPineconeServerlessSpecPtrOutputWithContext(context.Background())
}

func (i PineconeServerlessSpecArgs) ToPineconeServerlessSpecPtrOutputWithContext(ctx context.Context) PineconeServerlessSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeServerlessSpecOutput).ToPineconeServerlessSpecPtrOutputWithContext(ctx)
}

// PineconeServerlessSpecPtrInput is an input type that accepts PineconeServerlessSpecArgs, PineconeServerlessSpecPtr and PineconeServerlessSpecPtrOutput values.
// You can construct a concrete instance of `PineconeServerlessSpecPtrInput` via:
//
//	        PineconeServerlessSpecArgs{...}
//
//	or:
//
//	        nil
type PineconeServerlessSpecPtrInput interface {
	pulumi.Input

	ToPineconeServerlessSpecPtrOutput() PineconeServerlessSpecPtrOutput
	ToPineconeServerlessSpecPtrOutputWithContext(context.Context) PineconeServerlessSpecPtrOutput
}

type pineconeServerlessSpecPtrType PineconeServerlessSpecArgs

func PineconeServerlessSpecPtr(v *PineconeServerlessSpecArgs) PineconeServerlessSpecPtrInput {
	return (*pineconeServerlessSpecPtrType)(v)
}

func (*pineconeServerlessSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeServerlessSpec)(nil)).Elem()
}

func (i *pineconeServerlessSpecPtrType) ToPineconeServerlessSpecPtrOutput() PineconeServerlessSpecPtrOutput {
	return i.ToPineconeServerlessSpecPtrOutputWithContext(context.Background())
}

func (i *pineconeServerlessSpecPtrType) ToPineconeServerlessSpecPtrOutputWithContext(ctx context.Context) PineconeServerlessSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeServerlessSpecPtrOutput)
}

type PineconeServerlessSpecOutput struct{ *pulumi.OutputState }

func (PineconeServerlessSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconeServerlessSpec)(nil)).Elem()
}

func (o PineconeServerlessSpecOutput) ToPineconeServerlessSpecOutput() PineconeServerlessSpecOutput {
	return o
}

func (o PineconeServerlessSpecOutput) ToPineconeServerlessSpecOutputWithContext(ctx context.Context) PineconeServerlessSpecOutput {
	return o
}

func (o PineconeServerlessSpecOutput) ToPineconeServerlessSpecPtrOutput() PineconeServerlessSpecPtrOutput {
	return o.ToPineconeServerlessSpecPtrOutputWithContext(context.Background())
}

func (o PineconeServerlessSpecOutput) ToPineconeServerlessSpecPtrOutputWithContext(ctx context.Context) PineconeServerlessSpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PineconeServerlessSpec) *PineconeServerlessSpec {
		return &v
	}).(PineconeServerlessSpecPtrOutput)
}

// he public cloud where you would like your index hosted
func (o PineconeServerlessSpecOutput) Cloud() ServerlessSpecCloudOutput {
	return o.ApplyT(func(v PineconeServerlessSpec) ServerlessSpecCloud { return v.Cloud }).(ServerlessSpecCloudOutput)
}

// The region where you would like your index to be created. Different cloud providers have different regions available.
func (o PineconeServerlessSpecOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v PineconeServerlessSpec) string { return v.Region }).(pulumi.StringOutput)
}

type PineconeServerlessSpecPtrOutput struct{ *pulumi.OutputState }

func (PineconeServerlessSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeServerlessSpec)(nil)).Elem()
}

func (o PineconeServerlessSpecPtrOutput) ToPineconeServerlessSpecPtrOutput() PineconeServerlessSpecPtrOutput {
	return o
}

func (o PineconeServerlessSpecPtrOutput) ToPineconeServerlessSpecPtrOutputWithContext(ctx context.Context) PineconeServerlessSpecPtrOutput {
	return o
}

func (o PineconeServerlessSpecPtrOutput) Elem() PineconeServerlessSpecOutput {
	return o.ApplyT(func(v *PineconeServerlessSpec) PineconeServerlessSpec {
		if v != nil {
			return *v
		}
		var ret PineconeServerlessSpec
		return ret
	}).(PineconeServerlessSpecOutput)
}

// he public cloud where you would like your index hosted
func (o PineconeServerlessSpecPtrOutput) Cloud() ServerlessSpecCloudPtrOutput {
	return o.ApplyT(func(v *PineconeServerlessSpec) *ServerlessSpecCloud {
		if v == nil {
			return nil
		}
		return &v.Cloud
	}).(ServerlessSpecCloudPtrOutput)
}

// The region where you would like your index to be created. Different cloud providers have different regions available.
func (o PineconeServerlessSpecPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PineconeServerlessSpec) *string {
		if v == nil {
			return nil
		}
		return &v.Region
	}).(pulumi.StringPtrOutput)
}

type PineconeSpec struct {
	// Configuration needed to deploy a pod index.
	Pod *PineconePodSpec `pulumi:"pod"`
	// Configuration needed to deploy a serverless index.
	Serverless *PineconeServerlessSpec `pulumi:"serverless"`
}

// PineconeSpecInput is an input type that accepts PineconeSpecArgs and PineconeSpecOutput values.
// You can construct a concrete instance of `PineconeSpecInput` via:
//
//	PineconeSpecArgs{...}
type PineconeSpecInput interface {
	pulumi.Input

	ToPineconeSpecOutput() PineconeSpecOutput
	ToPineconeSpecOutputWithContext(context.Context) PineconeSpecOutput
}

type PineconeSpecArgs struct {
	// Configuration needed to deploy a pod index.
	Pod PineconePodSpecPtrInput `pulumi:"pod"`
	// Configuration needed to deploy a serverless index.
	Serverless PineconeServerlessSpecPtrInput `pulumi:"serverless"`
}

func (PineconeSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconeSpec)(nil)).Elem()
}

func (i PineconeSpecArgs) ToPineconeSpecOutput() PineconeSpecOutput {
	return i.ToPineconeSpecOutputWithContext(context.Background())
}

func (i PineconeSpecArgs) ToPineconeSpecOutputWithContext(ctx context.Context) PineconeSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeSpecOutput)
}

type PineconeSpecOutput struct{ *pulumi.OutputState }

func (PineconeSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PineconeSpec)(nil)).Elem()
}

func (o PineconeSpecOutput) ToPineconeSpecOutput() PineconeSpecOutput {
	return o
}

func (o PineconeSpecOutput) ToPineconeSpecOutputWithContext(ctx context.Context) PineconeSpecOutput {
	return o
}

// Configuration needed to deploy a pod index.
func (o PineconeSpecOutput) Pod() PineconePodSpecPtrOutput {
	return o.ApplyT(func(v PineconeSpec) *PineconePodSpec { return v.Pod }).(PineconePodSpecPtrOutput)
}

// Configuration needed to deploy a serverless index.
func (o PineconeSpecOutput) Serverless() PineconeServerlessSpecPtrOutput {
	return o.ApplyT(func(v PineconeSpec) *PineconeServerlessSpec { return v.Serverless }).(PineconeServerlessSpecPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetaDataConfigInput)(nil)).Elem(), MetaDataConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetaDataConfigPtrInput)(nil)).Elem(), MetaDataConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconePodSpecInput)(nil)).Elem(), PineconePodSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconePodSpecPtrInput)(nil)).Elem(), PineconePodSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeServerlessSpecInput)(nil)).Elem(), PineconeServerlessSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeServerlessSpecPtrInput)(nil)).Elem(), PineconeServerlessSpecArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeSpecInput)(nil)).Elem(), PineconeSpecArgs{})
	pulumi.RegisterOutputType(MetaDataConfigOutput{})
	pulumi.RegisterOutputType(MetaDataConfigPtrOutput{})
	pulumi.RegisterOutputType(PineconePodSpecOutput{})
	pulumi.RegisterOutputType(PineconePodSpecPtrOutput{})
	pulumi.RegisterOutputType(PineconeServerlessSpecOutput{})
	pulumi.RegisterOutputType(PineconeServerlessSpecPtrOutput{})
	pulumi.RegisterOutputType(PineconeSpecOutput{})
}
