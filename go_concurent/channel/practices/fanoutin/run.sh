n=0
while (($n <10 ))
do
    commend="BenchmarkPipelineFanBuffered_"$n
    echo $commend
    go test -test.bench = $commend
    echo ""
    let "n++"
done