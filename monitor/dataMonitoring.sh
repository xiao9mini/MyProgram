#!/bin/bash

# 远程执行命令
# curl -O https://raw.githubusercontent.com/xiao9mini/MyProgram/main/monitor/dataMonitoring.sh && sh dataMonitoring.sh 0 1024 0

# 用于监控网卡流量，当流量达到设定值时，自动关机
din=${1:-1024}  # 进站 默认1024
dout=${2:-1024} # 出站 默认1024
dall=${3:-2048} # 进出 默认2048

din=$((din == 0 ? 99999 : din))
dout=$((dout == 0 ? 99999 : dout))
dall=$((dall == 0 ? 99999 : dall))

ens=$(ip link | grep ens | awk -F ':' '{ print $2}' | awk '{print $1}')

curPath=$(dirname $(readlink -f "$0")) # 文件所在目录
cd $curPath                            # CD文件所在目录
mkdir -p $curPath/data

# 初始化检测脚本
if ! dpkg -s vnstat bc >/dev/null 2>&1; then
    sudo apt update 2>&1 >/dev/null
    sudo apt install vnstat bc -y 2>&1 >/dev/null
    sudo sed -i 's/;UnitMode 0/UnitMode 1/' /etc/vnstat.conf
    sudo sed -i 's/;MonthRotate 1/MonthRotate 1/' /etc/vnstat.conf
    sudo systemctl restart vnstat
fi

vnstat -i $ens 2>&1 >/dev/null
result=$(vnstat --oneline | awk -F ';' -vOFS='|' '{print $9,$10,$11}')

inr=$(echo $result | awk -F '|' '{if($1~/GB/)print $1;else print 0}' | awk '{print int($1)}')
outr=$(echo $result | awk -F '|' '{if($2~/GB/)print $2;else print 0}' | awk '{print int($1)}')
allr=$(echo $result | awk -F '|' '{if($3~/GB/)print $3;else print 0}' | awk '{print int($1)}')

# 当前时间
now=$(date "+%Y-%m-%d %H:%M")
echo $now $inr $outr $allr
# if in > 0 and inr > in
if [ $inr -ge $din ]; then
    echo IN $din $inr GB
    sudo /usr/sbin/shutdown -h now
fi

# if out > 0 and outr > out
if [ $outr -ge $dout ]; then
    echo OUT $dout $outr GB
    sudo /usr/sbin/shutdown -h now
fi

# if all > 0 and allr > all
if [ $allr -ge $dall ]; then
    echo ALL $dall $allr GB
    sudo /usr/sbin/shutdown -h now
fi

# 添加定时任务
cmd="/bin/bash $curPath/dataMonitoring.sh"
cron_job="*/5 * * * * $cmd $din $dout $dall > $curPath/data/check.log 2>&1" # 每5分钟执行一次

# 添加新的cmd
crontab -l | grep -v -F "$cmd" | {
    cat
    echo "$cron_job"
} | crontab -
