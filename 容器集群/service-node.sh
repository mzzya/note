ips=(10.101.1.79 10.101.1.62 10.101.1.78 10.101.1.63 10.101.1.36 10.101.1.77)

for ip in "${ips[@]}"
do
    wget $ip:31560
done

for ip in "${ips[@]}"
do
    wget $ip:32180
done