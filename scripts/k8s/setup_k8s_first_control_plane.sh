#!/bin/bash

# Centos X

RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"

# Setup Kubernetes Control Plane
echo -e "${GREEN}Setup Kubernetes Control Plane${ENDCOLOR}"

# Create kubernetes cluster
echo -e "${GREEN}Create kubernetes cluster${ENDCOLOR}"
# Execute only on the first control plane. Use join command to join to cluster.
# You will find to join command at the result of the below command on the control plane server.
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --upload-certs --kubernetes-version=v1.26.0 --control-plane-endpoint=$(hostname) --ignore-preflight-errors=all --cri-socket unix:///run/containerd/containerd.sock
# Enable local user to access cluster info
echo -e "${GREEN}Enable local user to access cluster info${ENDCOLOR}"
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=/etc/kubernetes/admin.conf >> ~/.bashrc
source ~/.bashrc
# Apply the CNI
kubectl apply -f https://github.com/coreos/flannel/raw/master/Documentation/kube-flannel.yml
# Disable pod schedule on control plane
echo -e "${GREEN}Disable pod schedule on master and control plane${ENDCOLOR}"
kubectl taint node $(hostname) node-role.kubernetes.io/control-plane:NoSchedule-

# Add control plane label for role
kubectl label nodes $(hostname -s) "kubernetes.io/role=control-plane"

# Create setup_kubernetes_control_planes.sh and setup_kubernetes_workers.sh files
# Add kubeadm join and the other commands(enable local user and tain for cp) for these files.
# How do I find the join command for kubeadm on the master?
# Create join command and certificate key
echo -e "${GREEN}Create join command and certificate key${ENDCOLOR}"
JOINCOMMAND=$(kubeadm token create --print-join-command)
CERTIFICATEKEY=$(kubeadm init phase upload-certs --upload-certs | grep -vw -e certificate -e Namespace)

# Create control plane script to join the kubernetes cluster as control plane role.
echo -e "${GREEN}Create control plane script to join de kubernetes cluster as control plane${ENDCOLOR}"
sudo tee setup_k8s_control_plane.sh <<EOF
#!/bin/bash
# Centos X
echo ${JOINCOMMAND} --control-plane --certificate-key ${CERTIFICATEKEY} \
  --node-name $(hostname -s) \
  --node-labels "kubernetes.io/role=control-plane"
EOF

# Create worker script to join de kubernetes cluster as worker role
echo -e "${GREEN}Create worker script to join de kubernetes cluster as worker role${ENDCOLOR}"
sudo tee setup_k8s_worker.sh <<EOF
#!/bin/bash
# Centos X
echo ${JOINCOMMAND} \
  --node-name $(hostname -s) \
  --node-labels "kubernetes.io/role=worker"
EOF