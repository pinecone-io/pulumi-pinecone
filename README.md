# Pinecone Pulumi Provider

<img src="img/pinecone.svg" width="50%">

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/pinecone-io/pulumi-pinecone)

This Pulumi Pinecone Provider enables you to manage your [Pinecone](https://www.pinecone.io/) collections and indexes using any language of Pulumi Infrastructure as Code.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pinecone-database/pulumi
```

or `yarn`:

```bash
yarn add @pinecone-database/pinecone
```

### Python

To use from Python, install using `pip`:

```bash
pip install pinecone_pulumi
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pinecone-io/pulumi-pinecone/sdk
```

## Configuration

The following configuration points are available for the `pinecone` provider:

- `pinecone:APIKey` - This is the Pinecone API key. (environment: `PINECONE_API_KEY`)
