# Notes

## Install operator-sdk

```sh
RELEASE_VERSION=v0.8.1
curl -OJL https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
chmod +x operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu && sudo cp operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
```

Ref: https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md   

## References
- https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md
- https://medium.com/@shubhomoybiswas/writing-kubernetes-operator-using-operator-sdk-c2e7f845163a
- https://github.com/operator-framework/operator-sdk-samples/tree/master/memcached-operator
