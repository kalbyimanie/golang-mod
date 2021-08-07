#!/bin/bash
DATABASE_HOSTNAME=${DB_HOST}
DATABASE_USERNAME=${DB_USERNAME}
DATABASE_PASSWORD=${DB_PASS}
DATABASE_PORT=${DB_PORT}
DATABASE_NAME=${DB_NAME}

mysql -u ${DATABASE_USERNAME} -p"${DATABASE_PASSWORD}" \
      -h ${DATABASE_HOSTNAME} -P ${DATABASE_PORT} -D ${DATABASE_NAME} -e"quit"

while [[ $? -ne 0 ]];do
  mysql -u ${DATABASE_USERNAME} -p"${DATABASE_PASSWORD}" \
        -h ${DATABASE_HOSTNAME} -P ${DATABASE_PORT} -D ${DATABASE_NAME} -e"quit"; sleep 10s
done

if [[ $? == 0 ]];then
  STOPS=30
  SECONDS=0
  while [[ ${SECONDS} -lt ${STOPS} ]]; do
    sleep 5s
    echo "successfully connected to database ${DATABASE_NAME} in host ${DATABASE_HOSTNAME} using the given credentials"
  done
fi