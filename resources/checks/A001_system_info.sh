#Collect system information

CPU_INFO=""
MEM_INFO=""
OS_INFO=""
DISK_INFO=""
CTL_INFO=""

if [[ "${SSH_SUPPORT}" = "false" ]]; then
  echo "SSH is not supported, skipping..." >&2
  exit 1
fi

function get_cpu_info() {
  local res=""
  local cpu_info="$(${CHECK_HOST_CMD} "lscpu")"
  cpu_info="${cpu_info/\"/\\\"}"
  local res_obj="{\"cmd2check\": \"lscpu\""
  while read -r line; do
    arg=$(echo "$line" | sed 's/:.*$//g' )
    value=$(echo "$line" | sed 's/^.*: *//g' )
    res_obj="$res_obj, \"$arg\": \"$value\""
  done <<< "$cpu_info"
  res_obj="$res_obj, \"raw\": \"$cpu_info\""
  res_obj="${res_obj} }"
  CPU_INFO=$res_obj #$(jq -n "$res_obj")
}

function get_mem_info() {
  local res=""
  local mem_info="$(${CHECK_HOST_CMD} "cat /proc/meminfo")"
  #local mem_info="$(cat /proc/meminfo)"
  mem_info="${mem_info/\"/\\\"}"
  local res_obj="{\"cmd2check\": \"cat /proc/meminfo\""
  while read -r line; do
    arg=$(echo "$line" | sed 's/:.*$//g' )
    value=$(echo "$line" | sed 's/^.*: *//g' )
    res_obj="$res_obj, \"$arg\": \"$value\""
  done <<< "$mem_info"
  res_obj="$res_obj, \"raw\": \"$mem_info\""
  res_obj="${res_obj} }"
  MEM_INFO=$res_obj #$(jq -n "$res_obj")
}

function get_system_info() {
  local uname_a_cmd='$(uname -a)'
  local uname_s_cmd='$(uname -s)'
  local uname_r_cmd='$(uname -r)'
  local uname_json_data=$(${CHECK_HOST_CMD} "echo '\"raw\":\"'$uname_a_cmd'\", \"operating_system\":\"'$uname_s_cmd'\", \"kernel_release\":\"'$uname_r_cmd'\"'")
  res_obj="{\"cmd2check\": \"uname -a\", $uname_json_data}"
  OS_INFO=$res_obj #$(jq -n "$res_obj")
}

function get_ctl_info() {
  local ctl_info="$(${CHECK_HOST_CMD} "hostnamectl status")"
  ctl_info="${ctl_info/\"/\\\"}"
  local res_obj="{\"cmd2check\": \"hostnamectl status\""
  while read -r line; do
    arg=$(echo "$line" | sed 's/:.*$//g' )
    value=$(echo "$line" | sed 's/^.*: *//g' )
    res_obj="$res_obj, \"$arg\": \"$value\""
  done <<< "$ctl_info"
  res_obj="$res_obj, \"raw\": \"$ctl_info\""
  res_obj="${res_obj} }"
  CTL_INFO=$res_obj #$(jq -n "$res_obj")
}

not_first=false
get_cpu_info
get_mem_info
get_system_info
get_ctl_info
host_obj="{\"cpu\": $CPU_INFO, \"ram\": $MEM_INFO, \"system\": $OS_INFO, \"virtualization\": $CTL_INFO }"
result="${host_obj}"

result=$(jq -n "$result")
echo "$result"
