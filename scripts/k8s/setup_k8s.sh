#!/bin/bash

# Centos X

RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"

# Setup Kubernetes
echo -e "${GREEN}Setup Kubernetes${ENDCOLOR}"

# Create kubernetes repo
echo -e "${GREEN}Create kubernetes repo${ENDCOLOR}"
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF

# Set SELinux in permissive mode (Disable effectively)
echo -e "${GREEN}Set SELinux in permissive mode (Disable effectively)${ENDCOLOR}"
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

# Install kubelet, kubeadm and kubectl
echo -e "${GREEN}Install kubelet, kubeadm and kubectl${ENDCOLOR}"
sudo yum install -y kubelet-1.26.0-0 kubeadm-1.26.0-0 kubectl-1.26.0-0 --disableexcludes=kubernetes

# Memory swapoff
echo -e "${GREEN}Memory swapoff${ENDCOLOR}"
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
sudo swapoff -a

# Enable Kubelet
echo -e "${GREEN}Enable Kubelet${ENDCOLOR}"
sudo systemctl enable kubelet

# Configure the firewall rules on the ports.
echo -e "${GREEN}Configure the firewall rules on the ports.${ENDCOLOR}"
firewall-cmd --permanent --add-port=6443/tcp
firewall-cmd --permanent --add-port=2379-2380/tcp
firewall-cmd --permanent --add-port=10250/tcp
firewall-cmd --permanent --add-port=10251/tcp
firewall-cmd --permanent --add-port=10252/tcp
firewall-cmd --permanent --add-port=10255/tcp
firewall-cmd --reload

# Add kernel modules
echo -e "${GREEN}Add kernel modules${ENDCOLOR}"
sudo modprobe overlay
sudo modprobe br_netfilter
sudo tee /etc/modules-load.d/containerd.conf <<EOF
overlay
br_netfilter
EOF

# Set the briged traffic for ip tables
echo -e "${GREEN}Set the briged traffic for ip tables${ENDCOLOR}"
sudo tee /etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF

# Reload configs
echo -e "${GREEN}Reload configs${ENDCOLOR}"
sysctl --system

# Check and install containerd
echo -e "${GREEN}Check and install containerd${ENDCOLOR}"
if (systemctl -q is-active containerd)
  then
      sed -i 's/disabled_plugins/#disabled_plugins/g' /etc/containerd/config.toml
      systemctl restart containerd
  else
    sudo yum install -y containerd.io
    mkdir -p /etc/containerd
    containerd config default>/etc/containerd/config.toml
    sed -i 's/disabled_plugins/#disabled_plugins/g' /etc/containerd/config.toml
    sudo systemctl restart containerd
    sudo systemctl enable containerd
fi

# Pull kubeadm config images
echo -e "${GREEN}Pull kubeadm config images${ENDCOLOR}"
sudo kubeadm config images pull --cri-socket unix:///run/containerd/containerd.sock --kubernetes-version v1.26.0