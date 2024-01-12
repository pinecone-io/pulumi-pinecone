// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinecone

import (
	"context"
	"reflect"

	"github.com/pinecone-io/pulumi-pinecone/sdk/go/pinecone/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The result of a get operation on a Pinecone index.
func LookupPineconeIndex(ctx *pulumi.Context, args *LookupPineconeIndexArgs, opts ...pulumi.InvokeOption) (*LookupPineconeIndexResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPineconeIndexResult
	err := ctx.Invoke("pinecone:index:lookupPineconeIndex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPineconeIndexArgs struct {
	// The name of the Pinecone index.
	Name string `pulumi:"name"`
}

// The result of a get operation on a Pinecone index.
type LookupPineconeIndexResult struct {
	// The dimensions of the vectors in the index. Defaults to 1536.
	Dimension *int   `pulumi:"dimension"`
	Host      string `pulumi:"host"`
	// The metric used to compute the distance between vectors.
	Metric IndexMetric `pulumi:"metric"`
	// The name of the Pinecone index.
	Name string `pulumi:"name"`
	// Describe how the index should be deployed.
	Spec   PineconeSpec `pulumi:"spec"`
	Status bool         `pulumi:"status"`
}

func LookupPineconeIndexOutput(ctx *pulumi.Context, args LookupPineconeIndexOutputArgs, opts ...pulumi.InvokeOption) LookupPineconeIndexResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPineconeIndexResult, error) {
			args := v.(LookupPineconeIndexArgs)
			r, err := LookupPineconeIndex(ctx, &args, opts...)
			var s LookupPineconeIndexResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPineconeIndexResultOutput)
}

type LookupPineconeIndexOutputArgs struct {
	// The name of the Pinecone index.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupPineconeIndexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPineconeIndexArgs)(nil)).Elem()
}

// The result of a get operation on a Pinecone index.
type LookupPineconeIndexResultOutput struct{ *pulumi.OutputState }

func (LookupPineconeIndexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPineconeIndexResult)(nil)).Elem()
}

func (o LookupPineconeIndexResultOutput) ToLookupPineconeIndexResultOutput() LookupPineconeIndexResultOutput {
	return o
}

func (o LookupPineconeIndexResultOutput) ToLookupPineconeIndexResultOutputWithContext(ctx context.Context) LookupPineconeIndexResultOutput {
	return o
}

// The dimensions of the vectors in the index. Defaults to 1536.
func (o LookupPineconeIndexResultOutput) Dimension() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) *int { return v.Dimension }).(pulumi.IntPtrOutput)
}

func (o LookupPineconeIndexResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) string { return v.Host }).(pulumi.StringOutput)
}

// The metric used to compute the distance between vectors.
func (o LookupPineconeIndexResultOutput) Metric() IndexMetricOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) IndexMetric { return v.Metric }).(IndexMetricOutput)
}

// The name of the Pinecone index.
func (o LookupPineconeIndexResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describe how the index should be deployed.
func (o LookupPineconeIndexResultOutput) Spec() PineconeSpecOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) PineconeSpec { return v.Spec }).(PineconeSpecOutput)
}

func (o LookupPineconeIndexResultOutput) Status() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPineconeIndexResult) bool { return v.Status }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPineconeIndexResultOutput{})
}
