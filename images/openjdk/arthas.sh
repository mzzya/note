if [ ! -f "arthas-boot.jar" ]; then
    wget https://arthas.aliyun.com/arthas-boot.jar
fi

java -jar arthas-boot.jar
