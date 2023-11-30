// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinecone

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type PineconeIndexResource struct {
	pulumi.CustomResourceState

	IndexName pulumi.StringOutput `pulumi:"indexName"`
}

// NewPineconeIndexResource registers a new resource with the given unique name, arguments, and options.
func NewPineconeIndexResource(ctx *pulumi.Context,
	name string, args *PineconeIndexResourceArgs, opts ...pulumi.ResourceOption) (*PineconeIndexResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dimension == nil {
		return nil, errors.New("invalid value for required argument 'Dimension'")
	}
	if args.Metric == nil {
		return nil, errors.New("invalid value for required argument 'Metric'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.PodType == nil {
		return nil, errors.New("invalid value for required argument 'PodType'")
	}
	if args.Pods == nil {
		return nil, errors.New("invalid value for required argument 'Pods'")
	}
	if args.Replicas == nil {
		return nil, errors.New("invalid value for required argument 'Replicas'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PineconeIndexResource
	err := ctx.RegisterResource("pinecone:index:PineconeIndexResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPineconeIndexResource gets an existing PineconeIndexResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPineconeIndexResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PineconeIndexResourceState, opts ...pulumi.ResourceOption) (*PineconeIndexResource, error) {
	var resource PineconeIndexResource
	err := ctx.ReadResource("pinecone:index:PineconeIndexResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PineconeIndexResource resources.
type pineconeIndexResourceState struct {
}

type PineconeIndexResourceState struct {
}

func (PineconeIndexResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*pineconeIndexResourceState)(nil)).Elem()
}

type pineconeIndexResourceArgs struct {
	// The dimensions of the vectors in the index.
	Dimension int `pulumi:"dimension"`
	// The metric used to compute the distance between vectors.
	Metric string `pulumi:"metric"`
	// The name of the Pinecone index.
	Name string `pulumi:"name"`
	// The type of pods to use for the index.
	PodType string `pulumi:"podType"`
	// The number of pods to use for the index.
	Pods int `pulumi:"pods"`
	// The number of replicas to use for the index.
	Replicas int `pulumi:"replicas"`
}

// The set of arguments for constructing a PineconeIndexResource resource.
type PineconeIndexResourceArgs struct {
	// The dimensions of the vectors in the index.
	Dimension pulumi.IntInput
	// The metric used to compute the distance between vectors.
	Metric pulumi.StringInput
	// The name of the Pinecone index.
	Name pulumi.StringInput
	// The type of pods to use for the index.
	PodType pulumi.StringInput
	// The number of pods to use for the index.
	Pods pulumi.IntInput
	// The number of replicas to use for the index.
	Replicas pulumi.IntInput
}

func (PineconeIndexResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pineconeIndexResourceArgs)(nil)).Elem()
}

type PineconeIndexResourceInput interface {
	pulumi.Input

	ToPineconeIndexResourceOutput() PineconeIndexResourceOutput
	ToPineconeIndexResourceOutputWithContext(ctx context.Context) PineconeIndexResourceOutput
}

func (*PineconeIndexResource) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeIndexResource)(nil)).Elem()
}

func (i *PineconeIndexResource) ToPineconeIndexResourceOutput() PineconeIndexResourceOutput {
	return i.ToPineconeIndexResourceOutputWithContext(context.Background())
}

func (i *PineconeIndexResource) ToPineconeIndexResourceOutputWithContext(ctx context.Context) PineconeIndexResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PineconeIndexResourceOutput)
}

type PineconeIndexResourceOutput struct{ *pulumi.OutputState }

func (PineconeIndexResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PineconeIndexResource)(nil)).Elem()
}

func (o PineconeIndexResourceOutput) ToPineconeIndexResourceOutput() PineconeIndexResourceOutput {
	return o
}

func (o PineconeIndexResourceOutput) ToPineconeIndexResourceOutputWithContext(ctx context.Context) PineconeIndexResourceOutput {
	return o
}

func (o PineconeIndexResourceOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v *PineconeIndexResource) pulumi.StringOutput { return v.IndexName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PineconeIndexResourceInput)(nil)).Elem(), &PineconeIndexResource{})
	pulumi.RegisterOutputType(PineconeIndexResourceOutput{})
}
