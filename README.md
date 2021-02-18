# Authentication
To get users / or add teants to a user [Go here](./Documentation/user-management.md).

## Paths

```
/ -> apiserver-proxy:80/
/.ory/kratos/public -> kratos:4433/
/oauth2 -> hydra:4444/oauth2
/.well-known -> hydra:4444/well-known

k8s.dolittle.studio/ -> "k8 apiserver proxy path"

/ -> studio
/.auth/select-tenant -> "select tenant page"
/.auth/login -> "select login provider page"

/.auth/initiate -> "cookie thing"
/.auth/callback -> "cookie thing"

/.openid -> hydra:public/

/.ory/kratos/public -> kratos:public/
```

## Running locally

```shell
# not needed anymore
kubectl -n system-authentication port-forward <postgres-pod> 8080:80
```

Create/register the OAuth 2.0 client in hydra. Make sure the go code oath client is the exact same.
```shell
kubectl -n system-authentication exec $(kubectl get pod -l "component=Browser" -o name -n system-authentication) -- hydra --endpoint http://localhost:4445 clients create --id client-id --secret client-secret -c http://local.dolittle.studio:8080/.auth/cookies/callback
```

List out your clients:
```shell
kubectl -n system-authentication exec $(kubectl get pod -l "component=Browser" -o name -n system-authentication) -- hydra --endpoint http://localhost:4445 clients list
```

Add to your /etc/hosts (bottom is a good idea)
```
127.0.0.1 oidc-provider.oidc-provider.svc.cluster.local
```

Also in `chrome://flags` disable __"Cookies without SameSite must be secure"__

To get users / or add teants to a user [Go here](./Documentation/user-management.md).


## Minikube port-forwarding
Set the socats to do the port forwarding for load balancers:
```shell
./socat.sh
```

## Linux `host.docker.internal` fix
Modify the coredns like this with uor own minikube ip (with the `.1` ending as we anna modify the dns gateawy stuff) so that we can keep using `host.docker.internal` on Linux system:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        health {
           lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
           ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
           max_concurrent 1000
        }
        file /etc/coredns/docker.internal.db docker.internal
        cache 30
        loop
        reload
        loadbalance
    }
  docker.internal.db: |
    docker.internal.        IN      SOA     sns.dns.icann.org. noc.dns.icann.org. 2015082541 7200 3600 1209600 3600
    docker.internal.        IN      A       192.168.49.1
    host.docker.internal.   IN      A       192.168.49.1

  example.db: |
      ; example.org test file
      example.org.            IN      SOA     sns.dns.icann.org. noc.dns.icann.org. 2015082541 7200 3600 1209600 3600
      example.org.            IN      NS      b.iana-servers.net.
      example.org.            IN      NS      a.iana-servers.net.
      example.org.            IN      A       127.0.0.1
      a.b.c.w.example.org.    IN      TXT     "Not a wildcard"
      cname.example.org.      IN      CNAME   www.example.net.

      service.example.org.    IN      SRV     8080 10 10 example.org.
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    k8s-app: kube-dns
  name: coredns
  namespace: kube-system
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      k8s-app: kube-dns
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        k8s-app: kube-dns
    spec:
      containers:
      - args:
        - -conf
        - /etc/coredns/Corefile
        image: k8s.gcr.io/coredns:1.7.0
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 5
        name: coredns
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        - containerPort: 9153
          name: metrics
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /ready
            port: 8181
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        resources:
          limits:
            memory: 170Mi
          requests:
            cpu: 100m
            memory: 70Mi
        securityContext:
          allowPrivilegeEscalation: false
          capabilities:
            add:
            - NET_BIND_SERVICE
            drop:
            - all
          readOnlyRootFilesystem: true
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /etc/coredns
          name: config-volume
          readOnly: true
      dnsPolicy: Default
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: coredns
      serviceAccountName: coredns
      terminationGracePeriodSeconds: 30
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: Corefile
            path: Corefile
          - key: docker.internal.db
            path: docker.internal.db
          name: coredns
        name: config-volume

```

## `kubectl` HTTP token fix
As `kubectl` [silenty doesn't transmit tokens over http](https://github.com/kubernetes/kubectl/issues/744) we need to do do some cert stuff in our localhost or kluster. WIP


## Local browser.yml changes
Add `127.0.0.1 local.dolittle.studio` to your local _/etc/hosts_.

Change the `hostAliases.ip` of hydra in browser.yml to point the _local.dolittle.studio_ to your local cluster host network, eg `docker.internal` ip or `minikube ip` minus one usually.
