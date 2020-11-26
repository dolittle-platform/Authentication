#! /bin/bash

PARSED=$(kubectl get svc -A -o json | jq -r '.items[] | select(.spec.type =="LoadBalancer") | .spec.ports[] | [.port, .nodePort] | join("-")' )
MINIKUBE_IP=$(minikube ip)

PIDS=""

for k in $PARSED; do
    PORTS=($(echo $k | tr "-" "\n"))

    echo 'Starting socat ' "TCP-LISTEN:${PORTS[0]},fork" "TCP:$MINIKUBE_IP:${PORTS[1]}"
    socat "TCP-LISTEN:${PORTS[0]},fork" "TCP:$MINIKUBE_IP:${PORTS[1]}" &
    PIDS="$PIDS $!"
done

echo 'socats started. good luck lol'

sleep infinity
echo 'we done'
kill $PIDS
