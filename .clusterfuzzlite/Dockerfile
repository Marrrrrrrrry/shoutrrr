FROM gcr.io/oss-fuzz-base/base-builder-go
RUN git clone --depth 1 https://github.com/marrrrrrrrry/shoutrrr
COPY . $SRC/shoutrrr
WORKDIR $SRC/shoutrrr
COPY ./clusterfuzzlite/build.sh $SRC/
