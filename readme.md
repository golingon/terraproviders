# Terra Providers

This repository contains popular Terraform providers generated using [Lingon](https://github.com/volvo-cars/lingon)'s `terragen`.

The current providers were created manually.
The plan is to add automation and do this on a regular interval.

```bash
export PROVIDER_NAME=aws
export PROVIDER_SOURCE=hashicorp/aws
export PROVIDER_VERSION=4.60.0
export PROVIDER=$PROVIDER_NAME=$PROVIDER_SOURCE:$PROVIDER_VERSION
export OUTDIR=$PROVIDER_NAME/$PROVIDER_VERSION
export OUTPKG=github.com/golingon/terraproviders/$PROVIDER_NAME/$PROVIDER_VERSION
terragen -out $OUTDIR -pkg $OUTPKG -provider $PROVIDER -force

cd $OUTDIR && go mod init $OUTPKG && go mod tidy && cd ../..
```
