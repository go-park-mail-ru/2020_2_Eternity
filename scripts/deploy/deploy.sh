#!/bin/bash

numParams=1

serviceName=pinterest
authMsName=pinterest-auth
chatMsName=pinterest-chat
searchMsName=pinterest-search

usage="Usage. Run from root project dir:
>>>  ./scripts/deploy/deploy.sh [branch-name] (service-name)

* branch-name - name of git branch, which last commit will be build for deploy
* service-name - name of backend service, default 'pinterest'"

if [ -n "$2" ]
then
  serviceName=$2
fi



if [ $# -lt  $numParams ]
then
  echo "No parameters found. "
  echo "$usage"


else
  git fetch || { echo "Fetch error" ; exit 1; }
  git checkout $1 || { echo "Can't checkout branch $1"  ; exit 1; }
  git pull || { echo "Pull error"  ; exit 1; }
  make dbsetup || { echo "make: database setup error"  ; exit 1; }
  make build || { echo "make: build error"  ; exit 1; }

  df=$(diff /etc/pinterest/config.yaml configs/yaml/config.yaml)
  if [ -n "$df" ]
  then
    echo "
    Config files are different
    "
    diff /etc/pinterest/config.yaml configs/yaml/config.yaml
    echo "
    Would you like to continue? (y/n)"

    read line
    if [ $line != "y" ] && [  $line != "Y" ]
    then
      echo "Interrupting. Now you have your actual binary built into path,
      specified by 'make build'. And that's all. Server was not restarted."
      exit 0
    fi
  fi

  sudo systemctl stop $serviceName || { echo "systemctl: can't stop $serviceName"  ; exit 1; }
  sudo systemctl stop $authMsName || { echo "systemctl: can't stop $authMsName"  ; exit 1; }
  sudo systemctl stop $chatMsName || { echo "systemctl: can't stop $chatMsName"  ; exit 1; }
  sudo systemctl stop $searchMsName || { echo "systemctl: can't stop $searchMsName"  ; exit 1; }

  sudo cp build/bin/api /bin/pinterest/api || { echo "cp error"  ; exit 1; }
  sudo cp build/bin/chat /bin/pinterest/chat || { echo "cp error"  ; exit 1; }
  sudo cp build/bin/auth /bin/pinterest/auth || { echo "cp error"  ; exit 1; }
  sudo cp build/bin/search /bin/pinterest/search || { echo "cp error"  ; exit 1; }


  sudo systemctl start $authMsName || { echo "systemctl: can't stop $authMsName"  ; exit 1; }
  sudo systemctl start $chatMsName || { echo "systemctl: can't stop $chatMsName"  ; exit 1; }
  sudo systemctl start $searchMsName || { echo "systemctl: can't stop $searchMsName"  ; exit 1; }
  sudo systemctl start $serviceName || { echo "systemctl: can't start $serviceName"  ; exit 1; }

  echo "Success! New implementation of service is running"
fi