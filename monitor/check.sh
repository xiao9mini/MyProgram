
in=500      # 进站 默认500
out=200    # 出站 默认150
all=500    # 进出 默认500

vnstat -i ens4 2>&1 >/dev/null
result=$(vnstat --oneline | awk -F ';' -vOFS='|' '{print $9,$10,$11}')

inr=$(echo $result | awk -F '|' '{if($1~/GB/)print $1;else print 0}' | awk '{print int($1)}')
outr=$(echo $result | awk -F '|' '{if($2~/GB/)print $2;else print 0}' | awk '{print int($1)}')
allr=$(echo $result | awk -F '|' '{if($3~/GB/)print $3;else print 0}' | awk '{print int($1)}')

# 当前时间
now=$(date "+%Y-%m-%d %H:%M")
echo $now $inr $outr $allr
# if in > 0 and inr > in
if [ $inr -ge $in ]; then
    echo IN $in $inr GB
    sudo /usr/sbin/shutdown -h now
fi

# if out > 0 and outr > out
if [ $outr -ge $out ]; then
    echo OUT $out $outr GB
    sudo /usr/sbin/shutdown -h now
fi

# if all > 0 and allr > all
if [ $allr -ge $all ]; then
    echo ALL $all $allr GB
    sudo /usr/sbin/shutdown -h now
fi


