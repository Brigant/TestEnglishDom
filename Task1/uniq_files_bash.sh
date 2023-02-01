#!/bin/bash

LIST=$(find -type f -printf "%f\n")

declare -a array=($(echo $LIST | tr ";" "\n"))

# ====== firs method ==================

# declare -A uniq_tmp

# for i in "${array[@]}"; do
#     uniq_tmp[$i]=0 
# done

# for i in "${!uniq_tmp[@]}"; do
#     echo $i
# done 

#============ second method =============

uniqs_arr=($(for i in "${array[@]}"; do echo "${i}"; done | sort -u))

for i in "${uniqs_arr[@]}"; do
    echo $i
done 