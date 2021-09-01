tag=$1
if [ -z "$tag" ]
then
    tag="latest"
fi
docker build -t 676007236/bush:$tag .
docker push 676007236/bush:$tag