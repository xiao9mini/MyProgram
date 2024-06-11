
# 用于监控网卡流量，当流量达到设定值时，自动关机
in=${1:-500}  # 进站 默认500
out=${2:-150} # 出站 默认150
all=${3:-500} # 进出 默认500

ens=$(ip link | grep ens |   awk -F ':' '{ print $2}' | awk '{print $1}')

curPath=$(dirname $(readlink -f "$0")) # 文件所在目录
cd $(dirname $(readlink -f "$0"))      # CD文件所在目录

sudo apt install vnstat bc -y 2>&1 >/dev/null
sudo sed -i 's/;UnitMode 0/UnitMode 1/' /etc/vnstat.conf
sudo sed -i 's/;MonthRotate 1/MonthRotate 1/' /etc/vnstat.conf
sudo systemctl restart vnstat

echo "
in=$in      # 进站 默认500
out=$out    # 出站 默认150
all=$all    # 进出 默认500

vnstat -i $ens 2>&1 >/dev/null
result=\$(vnstat --oneline | awk -F ';' -vOFS='|' '{print \$9,\$10,\$11}')

inr=\$(echo \$result | awk -F '|' '{if(\$1~/GB/)print \$1;else print 0}' | awk '{print int(\$1)}')
outr=\$(echo \$result | awk -F '|' '{if(\$2~/GB/)print \$2;else print 0}' | awk '{print int(\$1)}')
allr=\$(echo \$result | awk -F '|' '{if(\$3~/GB/)print \$3;else print 0}' | awk '{print int(\$1)}')

# 当前时间
now=\$(date \"+%Y-%m-%d %H:%M\")
echo \$now \$inr \$outr \$allr
# if in > 0 and inr > in
if [ \$inr -ge \$in ]; then
    echo IN \$in \$inr GB
    sudo /usr/sbin/shutdown -h now
fi

# if out > 0 and outr > out
if [ \$outr -ge \$out ]; then
    echo OUT \$out \$outr GB
    sudo /usr/sbin/shutdown -h now
fi

# if all > 0 and allr > all
if [ \$allr -ge \$all ]; then
    echo ALL "\$all \$allr GB"
    sudo /usr/sbin/shutdown -h now
fi

" > "$curPath/check.sh"

cmd="/bin/bash $curPath/check.sh"
cron_job="*/5 * * * * $cmd > $curPath/check.log 2>&1"   # 每5分钟执行一次
# 如果存在则删除cmd
# echo "$current_crontab" | grep -v "$cmd" | crontab -

# 添加新的cmd
(crontab -l | grep -v -F "$cmd" ; echo "$cron_job") | crontab -

crontab -l
echo "监控网卡流量，当流量达到设定值时，自动关机。（5分钟）"
in=$in      # 进站 默认500
out=$out    # 出站 默认150
all=$all    # 进出 默认500"