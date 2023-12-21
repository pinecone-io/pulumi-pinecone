// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinecone

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-pinecone/sdk/go/pinecone/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PineconeCollection struct {
	pulumi.CustomResourceState

	// The dimension of the vectors stored in each record held in the collection.
	Dimension pulumi.IntOutput `pulumi:"dimension"`
	// The environment where the collection is hosted.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// The name of the collection to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of records stored in the collection.
	RecordCount pulumi.IntOutput `pulumi:"recordCount"`
	// The size of the collection in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// The name of the index to be used as the source for the collection.
	Source pulumi.StringOutput `pulumi:"source"`
}

// NewPineconeCollection registers a new resource with the given unique name, arguments, and options.
func NewPineconeCollection(ctx *pulumi.Context,
	name string, args *PineconeCollectionArgs, opts ...pulumi.ResourceOption) (*PineconeCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PineconeCollection
	err := ctx.RegisterResource("pinecone:index:PineconeCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPineconeCollection gets an existing PineconeCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPineconeCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PineconeCollectionState, opts ...pulumi.ResourceOption) (*PineconeCollection, error) {
	var resource PineconeCollection
	err := ctx.ReadResource("pinecone:index:PineconeCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PineconeCollection resources.
type pineconeCollectionState struct {
}

type PineconeCollectionState struct {
}

func (PineconeCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*pineconeCollectionState)(nil)).Elem()
}

type pineconeCollectionArgs struct {
	// The name of the collection to be created.
	Name string `pulumi:"name"`
	// The name of the index to be used as the source for the collection.
	Source string `pulumi:"source"`
}

// The set of arguments for constructing a PineconeCollection resource.
type PineconeCollectionArgs struct {
	// The name of the collection to be created.
	Name pulumi.StringInput
	// The name of the index to be used as the source for the collection.
	Source pulumi.StringInput
}

func (PineconeCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pineconeCollectionArgs)(nil)).Elem()
}

type PineconeCollectionInput interface {
	pulumi.Input

	ToPineconeCollectionOutput() PineconeCollectionOutput
	ToPineconeCollectionOutputWithContext(ctx context.Context) PineconeCollectionOutput
}

func (*PineconeCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeCollection)(nil)).Elem()
}

func (i *PineconeCollection) ToPineconeCollectionOutput() PineconeCollectionOutput {
	return i.ToPineconeCollectionOutputWithContext(context.Background())
}

func (i *PineconeCollection) ToPineconeCollectionOutputWithContext(ctx context.Context) PineconeCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeCollectionOutput)
}

// PineconeCollectionArrayInput is an input type that accepts PineconeCollectionArray and PineconeCollectionArrayOutput values.
// You can construct a concrete instance of `PineconeCollectionArrayInput` via:
//
//	PineconeCollectionArray{ PineconeCollectionArgs{...} }
type PineconeCollectionArrayInput interface {
	pulumi.Input

	ToPineconeCollectionArrayOutput() PineconeCollectionArrayOutput
	ToPineconeCollectionArrayOutputWithContext(context.Context) PineconeCollectionArrayOutput
}

type PineconeCollectionArray []PineconeCollectionInput

func (PineconeCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PineconeCollection)(nil)).Elem()
}

func (i PineconeCollectionArray) ToPineconeCollectionArrayOutput() PineconeCollectionArrayOutput {
	return i.ToPineconeCollectionArrayOutputWithContext(context.Background())
}

func (i PineconeCollectionArray) ToPineconeCollectionArrayOutputWithContext(ctx context.Context) PineconeCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeCollectionArrayOutput)
}

// PineconeCollectionMapInput is an input type that accepts PineconeCollectionMap and PineconeCollectionMapOutput values.
// You can construct a concrete instance of `PineconeCollectionMapInput` via:
//
//	PineconeCollectionMap{ "key": PineconeCollectionArgs{...} }
type PineconeCollectionMapInput interface {
	pulumi.Input

	ToPineconeCollectionMapOutput() PineconeCollectionMapOutput
	ToPineconeCollectionMapOutputWithContext(context.Context) PineconeCollectionMapOutput
}

type PineconeCollectionMap map[string]PineconeCollectionInput

func (PineconeCollectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PineconeCollection)(nil)).Elem()
}

func (i PineconeCollectionMap) ToPineconeCollectionMapOutput() PineconeCollectionMapOutput {
	return i.ToPineconeCollectionMapOutputWithContext(context.Background())
}

func (i PineconeCollectionMap) ToPineconeCollectionMapOutputWithContext(ctx context.Context) PineconeCollectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeCollectionMapOutput)
}

type PineconeCollectionOutput struct{ *pulumi.OutputState }

func (PineconeCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeCollection)(nil)).Elem()
}

func (o PineconeCollectionOutput) ToPineconeCollectionOutput() PineconeCollectionOutput {
	return o
}

func (o PineconeCollectionOutput) ToPineconeCollectionOutputWithContext(ctx context.Context) PineconeCollectionOutput {
	return o
}

// The dimension of the vectors stored in each record held in the collection.
func (o PineconeCollectionOutput) Dimension() pulumi.IntOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.IntOutput { return v.Dimension }).(pulumi.IntOutput)
}

// The environment where the collection is hosted.
func (o PineconeCollectionOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// The name of the collection to be created.
func (o PineconeCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of records stored in the collection.
func (o PineconeCollectionOutput) RecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.IntOutput { return v.RecordCount }).(pulumi.IntOutput)
}

// The size of the collection in bytes.
func (o PineconeCollectionOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The name of the index to be used as the source for the collection.
func (o PineconeCollectionOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *PineconeCollection) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

type PineconeCollectionArrayOutput struct{ *pulumi.OutputState }

func (PineconeCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PineconeCollection)(nil)).Elem()
}

func (o PineconeCollectionArrayOutput) ToPineconeCollectionArrayOutput() PineconeCollectionArrayOutput {
	return o
}

func (o PineconeCollectionArrayOutput) ToPineconeCollectionArrayOutputWithContext(ctx context.Context) PineconeCollectionArrayOutput {
	return o
}

func (o PineconeCollectionArrayOutput) Index(i pulumi.IntInput) PineconeCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PineconeCollection {
		return vs[0].([]*PineconeCollection)[vs[1].(int)]
	}).(PineconeCollectionOutput)
}

type PineconeCollectionMapOutput struct{ *pulumi.OutputState }

func (PineconeCollectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PineconeCollection)(nil)).Elem()
}

func (o PineconeCollectionMapOutput) ToPineconeCollectionMapOutput() PineconeCollectionMapOutput {
	return o
}

func (o PineconeCollectionMapOutput) ToPineconeCollectionMapOutputWithContext(ctx context.Context) PineconeCollectionMapOutput {
	return o
}

func (o PineconeCollectionMapOutput) MapIndex(k pulumi.StringInput) PineconeCollectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PineconeCollection {
		return vs[0].(map[string]*PineconeCollection)[vs[1].(string)]
	}).(PineconeCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeCollectionInput)(nil)).Elem(), &PineconeCollection{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeCollectionArrayInput)(nil)).Elem(), PineconeCollectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeCollectionMapInput)(nil)).Elem(), PineconeCollectionMap{})
	pulumi.RegisterOutputType(PineconeCollectionOutput{})
	pulumi.RegisterOutputType(PineconeCollectionArrayOutput{})
	pulumi.RegisterOutputType(PineconeCollectionMapOutput{})
}
