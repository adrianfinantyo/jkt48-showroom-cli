# create build command

# see the following for more specific info
# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

# ask for the target
echo "What is the target OS?"
read target

# ask for the arch
echo "What is the target arch?"
read arch

# ask for the source
echo "What is the source?"
read source

# ask for the output
echo "What is the output?"
read output

GOOS=$target GOARCH=$arch go build -o $output $source