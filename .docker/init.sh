#!/bin/bash
set -e

appConnect () {
  # Run Server
  ${LILYPAD_INSTALL_DIR}/${LILYPAD_CONNECT_BIN}
}

appProxy () {
  # When using --link to connect, set address in config to this
  if [[ -n ${CONNECT_PORT_5091_TCP_ADDR} && -f "proxy.yml" ]]; then
    sed -i -E 's|address: (.*)+$|address: '${CONNECT_PORT_5091_TCP_ADDR:-connect}':'${CONNECT_PORT_5091_TCP_PORT:-5091}'|' proxy.yml
  fi
  # Run server
  ${LILYPAD_INSTALL_DIR}/${LILYPAD_PROXY_BIN}
}

appHelp () {
  echo "Available options:"
  echo " connect            - Starts LilyPad Connect"
  echo " proxy              - Starts LilyPad Proxy"
  echo " app:help           - Displays the help (default)"
  echo " [command]          - Execute the specified linux command eg. bash."
}

case "$1" in
  connect)
    appConnect
    ;;
  proxy)
    appProxy
    ;;
  app:help)
    appHelp
    ;;
  *)
    if [ -x $1 ]; then
      $1
    else
      prog=$(which $1)
      if [ -n "${prog}" ] ; then
        shift 1
        $prog $@
      else
        appHelp
      fi
    fi
    ;;
esac

exit 0
