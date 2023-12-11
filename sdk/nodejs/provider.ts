// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'pinecone';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The API token for Pinecone.
     */
    public readonly apiToken!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The environment for the Pinecone API.
     */
    public readonly pineconeEnv!: pulumi.Output<string>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.apiToken === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiToken'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.pineconeEnv === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pineconeEnv'");
            }
            resourceInputs["apiToken"] = args?.apiToken ? pulumi.secret(args.apiToken) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pineconeEnv"] = args ? args.pineconeEnv : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The API token for Pinecone.
     */
    apiToken: pulumi.Input<string>;
    name: pulumi.Input<string>;
    /**
     * The environment for the Pinecone API.
     */
    pineconeEnv: pulumi.Input<string>;
}
