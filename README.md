# cert-manager webhook for Namecheap

# Instructions for use with Let's Encrypt

Use helm to deploy this into your `cert-manager` namespace:

``` sh
# Make sure you're in the right context:
# kubectl config use-context mycontext

# cert-manager is by default in the cert-manager context
helm install -n cert-manager namecheap-webhook deploy/cert-manager-webhook-namecheap/
```

Create the cluster issuers:

``` sh
helm install --set email=yourname@example.com -n cert-manager letsencrypt-namecheap-issuer deploy/letsencrypt-namecheap-issuer/
```

Go to [namecheap](https://www.namecheap.com/myaccount/login/) and set up your API key (note that you'll need to whitelist the public IP of the k8s cluster to use the webhook), and set the secret:

``` yaml
apiVersion: v1
kind: Secret
metadata:
  name: namecheap-credentials
  namespace: cert-manager
type: Opaque
stringData:
  apiUser: <not base64 encoded>my_username_from_namecheap</not base64 encoded>
  apiKey: <not base64 encoded>my_api_key_from_namecheap</not base64 encoded>
```

## Examples: 

<details>

<summary>Example with a wildcard certificate:</summary>

Now you can create a certificate in _staging_ for testing:

``` yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: wildcard-cert-stage
  namespace: default
spec:
  secretName: wildcard-cert-stage
  commonName: "*.<domain>"
  issuerRef:
    kind: ClusterIssuer
    name: letsencrypt-stage
  dnsNames:
  - "*.<domain>"
```

And now validate that it worked:

``` sh
kubectl get certificates -n default
kubectl describe certificate wildcard-cert-stage
```

And finally, create your _production_ cert, and it'll be ready to use in the `wildcard-cert-prod` secret.

``` yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: wildcard-cert-prod
  namespace: default
spec:
  secretName: wildcard-cert-prod
  commonName: "*.<domain>"
  issuerRef:
    kind: ClusterIssuer
    name: letsencrypt-prod
  dnsNames:
  - "*.<domain>"
```

And now validate that it worked:

``` sh
kubectl get certificates -n default
kubectl describe certificate service-cert-prod
```

</details>

<details>

<summary>Example with non-wildcard certificate:</summary>


Now you can create a certificate in _staging_ for testing:

``` yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: service-cert-stage
  namespace: default
spec:
  secretName: service-cert-stage
  commonName: "service.<domain>"
  issuerRef:
    kind: ClusterIssuer
    name: letsencrypt-stage
  dnsNames:
  - "service.<domain>"
```

And now validate that it worked:

``` sh
kubectl get certificates -n default
kubectl describe certificate service-cert-stage
```

And finally, create your _production_ cert, and it'll be ready to use in the `wildcard-cert-prod` secret.

``` yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: service-cert-prod
  namespace: default
spec:
  secretName: service-cert-prod
  commonName: "service.<domain>"
  issuerRef:
    kind: ClusterIssuer
    name: letsencrypt-prod
  dnsNames:
  - "service.<domain>"
```

And now validate that it worked:

``` sh
kubectl get certificates -n default
kubectl describe certificate service-cert-prod
```
  
</details>

## Traefik example

Hint: https://github.com/cert-manager/cert-manager/issues/3501#issuecomment-826132642

Full example can be found in [testService.yaml](testService.yaml)

```yaml
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
 name: testservice
 namespace: testservice
 annotations:
   cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
 tls:
   - hosts:
       - testservice.vps9.hochguertel.work
     secretName: testservice.vps9.hochguertel.work-prod
 rules:
   - host: testservice.vps9.hochguertel.work
     http:
       paths:
         - path: /
           pathType: Prefix
           backend:
             service:
               name: testservice
               port:
                 name: web
```

> [!IMPORTANT]
> Important is that you **use** the annotation `cert-manager.io/cluster-issuer` and not `cert-manager.io/issuer`!


```yaml
  annotations:
   cert-manager.io/cluster-issuer: "letsencrypt-prod"
```

### TODO:

- [ ] add simple nginx example to test that it works

### Running the test suite

All DNS providers **must** run the DNS01 provider conformance testing suite,
else they will have undetermined behaviour when used with cert-manager.

**It is essential that you configure and run the test suite when creating a
DNS01 webhook.**

An example Go test file has been provided in [main_test.go](https://github.com/jetstack/cert-manager-webhook-example/blob/master/main_test.go).

You can run the test suite with:

```bash
$ TEST_ZONE_NAME=example.com. make test
```

The example file has a number of areas you must fill in and replace with your
own options in order for tests to pass.

## Thanks

- Thanks to [Addison van den Hoeven](https://github.com/Addyvan), from https://github.com/jetstack/cert-manager/issues/646
- Thanks to [Kelvie Wong](https://github.com/kelvie/cert-manager-webhook-namecheap), for creating the code of this repository and the docker container of the issuer.
