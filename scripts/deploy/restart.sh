serviceName=pinterest
authMsName=pinterest-auth
chatMsName=pinterest-chat
searchMsName=pinterest-search

sudo systemctl stop $serviceName || { echo "systemctl: can't stop $serviceName"  ; exit 1; }
sudo systemctl stop $authMsName || { echo "systemctl: can't stop $authMsName"  ; exit 1; }
sudo systemctl stop $chatMsName || { echo "systemctl: can't stop $chatMsName"  ; exit 1; }
sudo systemctl stop $searchMsName || { echo "systemctl: can't stop $searchMsName"  ; exit 1; }

sudo cp /home/ubuntu/pinterest-backend/bin-ci/api /bin/pinterest/api || { echo "cp error"  ; exit 1; }
sudo cp /home/ubuntu/pinterest-backend/bin-ci/chat /bin/pinterest/chat || { echo "cp error"  ; exit 1; }
sudo cp /home/ubuntu/pinterest-backend/bin-ci/auth /bin/pinterest/auth || { echo "cp error"  ; exit 1; }
sudo cp /home/ubuntu/pinterest-backend/bin-ci/search /bin/pinterest/search || { echo "cp error"  ; exit 1; }


sudo systemctl start $authMsName || { echo "systemctl: can't stop $authMsName"  ; exit 1; }
sudo systemctl start $chatMsName || { echo "systemctl: can't stop $chatMsName"  ; exit 1; }
sudo systemctl start $searchMsName || { echo "systemctl: can't stop $searchMsName"  ; exit 1; }
sudo systemctl start $serviceName || { echo "systemctl: can't start $serviceName"  ; exit 1; }