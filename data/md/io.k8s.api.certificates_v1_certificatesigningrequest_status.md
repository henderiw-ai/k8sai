# openAPI schema status for io.k8s.api.certificates.v1.CertificateSigningRequest

## description

CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.

Kubelets use this API to obtain:
 1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
 2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).

This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.

## schema

```yaml
|
  description: |-
    CertificateSigningRequest objects provide a mechanism to obtain x509 certificates by submitting a certificate signing request, and having it asynchronously approved and issued.

    Kubelets use this API to obtain:
     1. client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client-kubelet" signerName).
     2. serving certificates for TLS endpoints kube-apiserver can connect to securely (with the "kubernetes.io/kubelet-serving" signerName).

    This API can be used to request client certificates to authenticate to kube-apiserver (with the "kubernetes.io/kube-apiserver-client" signerName), or to obtain certificates from custom non-Kubernetes signers.
  properties:
    apiVersion:
      description: 'APIVersion defines the versioned schema of this representation of
        an object. Servers should convert recognized schemas to the latest internal
        value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
      type: string
    kind:
      description: 'Kind is a string value representing the REST resource this object
        represents. Servers may infer this from the endpoint the client submits requests
        to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
      type: string
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta'
    spec:
      $ref: '#/definitions/io.k8s.api.certificates.v1.CertificateSigningRequestSpec'
      description: spec contains the certificate request, and is immutable after creation.
        Only the request, signerName, expirationSeconds, and usages fields can be set
        on creation. Other fields are derived by Kubernetes and cannot be modified by
        users.
    status:
      $ref: '#/definitions/io.k8s.api.certificates.v1.CertificateSigningRequestStatus'
      description: status contains information about whether the request is approved
        or denied, and the certificate issued by the signer, or the failure condition
        indicating signer failure.
  required:
  - spec
  type: object
  x-kubernetes-group-version-kind:
  - group: certificates.k8s.io
    kind: CertificateSigningRequest
    version: v1

```