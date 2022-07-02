#! /bin/bash 

HADOOP_VERSION=$1
if [ -z $HADOOP_VERSION ];then
    HADOOP_VERSION=3.3.2
fi

MIRROR=https://mirrors.tuna.tsinghua.edu.cn/apache/hadoop/common/hadoop-$HADOOP_VERSION



function obtain_hadoop_install_package() {
    if [ ! -f /soft/hadoop-$HADOOP_VERSION.tar.gz ];then
        echo "[WARNING] The file hadoop-$HADOOP_VERSION.tar.gz for hadoop installation not existed in local directory /soft." 
        echo "[WARNING] Start obtain installation package from Internet, It will take several minutes."
        echo -n "Downloading $MIRROR/hadoop-$HADOOP_VERSION.tar.gz to local directory ---> /soft ......"
        wget -q $MIRROR/hadoop-$HADOOP_VERSION.tar.gz -O /soft/hadoop-$HADOOP_VERSION.tar.gz
        echo "Done"
    fi
}


function display_help() {}

function distribute_install_file() {

}

function process_datanode(){}

function process_namenode(){}

function main() {
    for i in 
}