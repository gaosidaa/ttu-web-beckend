FROM --platform=$TARGETPLATFORM alpine

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR  /app

ADD resource $WORKDIR/
ADD manifest/config/config.yaml $WORKDIR/config.yaml
ADD ttu_backend_arm_linux $WORKDIR/main

RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
