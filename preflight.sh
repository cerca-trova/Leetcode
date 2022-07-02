# /bin/bash

function preflight() {
    # check whether jdk was installed and configured correctly
    # echo "Running preflight check on host $host..."
    # sleep 1
    if [ -z `echo $JAVA_HOME` ] || [ -z  `/usr/bin/command -v java` ];then
        echo "[Preflight] [Error] Java development kit was not installed or configured!"
        echo "Installation Aborted!"
        exit 1
    fi

    if [ ! -d /soft ];then
        echo -n "[preflight] [INFO] /soft directory does not existed,creating..."
        mkdir /soft
        sleep 1
        echo "Done"
    fi
    if [ ! -d /data1 ];then
        echo -n "[preflight] [INFO] /data1 directory does not existed,creating..."
        mkdir /data1
        sleep 1
        echo "Done"
    fi
    # echo "[preflight] Everything on host $host is OK!"

}

preflight