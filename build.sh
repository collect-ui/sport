rm -rf dist/*
go build -o dist/bin -v -x
# windows
# go build -o dist/bin.exe -v -x
cp -r static dist
cp -r conf dist
cp -r collect dist
cp -r frontend dist
cp -r conf dist
cp -r shutdown.sh dist
cp -r startup.sh dist
cp -r database dist
