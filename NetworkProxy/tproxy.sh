#!/bin/bash
#
# Transparent proxy
#

if [[ $# -ne 2 ]]; then
    echo "Usage: $0 rhost rport"
    exit 0
fi

if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

host=$1
dport=$2
tmpport1=$RANDOM
tmpport2=$RANDOM

# create new queue
iptables -t nat -N proxyqueue

# redirect all incoming packets to new queue first
iptables -t nat -I OUTPUT 1 -j proxyqueue

for ip in $(dig +short $host); do
    # redirect desthost:destport -> localhost:tmpport1
    iptables -t nat -A proxyqueue -p tcp -d $ip --dport $dport -j DNAT --to-destination 127.0.0.1:$tmpport1
done

# redirect desthost:tmpport2 -> desthost:destport
ip=$(dig +short $host | head -1)
iptables -t nat -A proxyqueue -p tcp -d $ip --dport $tmpport2 -j DNAT --to-destination $ip:$dport

# start proxy
./proxy 127.0.0.1:$tmpport1 $ip:$tmpport2

# revert first rule
iptables -t nat -D OUTPUT 1

# flush queue
iptables -t nat -F proxyqueue

# delete queue
iptables -t nat -X proxyqueue
