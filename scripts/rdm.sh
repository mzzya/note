git clone --recursive https://github.com/uglide/RedisDesktopManager.git -b 2019 rdm && cd ./rdm
cd ./src && cp ./resources/Info.plist.sample ./resources/Info.plist
pip3 install -t ../bin/osx/release -r py/requirements.txt