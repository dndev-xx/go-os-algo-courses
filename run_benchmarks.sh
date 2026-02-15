#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

mkdir -p benchmark_results

run_benchmark_group() {
	local name=$1
	local pattern=$2
	local color=$3
	local output_file="benchmark_results/${name}_bench.log"
	local summary_file="benchmark_results/${name}_summary.log"

	echo -e "${color}========================================${NC}"
	echo -e "${color}=== Running $name Benchmarks ===${NC}"
	echo -e "${color}========================================${NC}"

	go test -bench="$pattern" -benchmem ./internal/basic-module/... -v >"$output_file"

	echo -e "${CYAN}$name Benchmark Results:${NC}" >"$summary_file"
	echo "----------------------------------------" >>"$summary_file"
	grep "^Benchmark" "$output_file" | while read -r line; do
		if [[ $line =~ ^(Benchmark[^[:space:]]+)[[:space:]]+([0-9]+)[[:space:]]+([0-9.]+)[[:space:]]ns/op[[:space:]]+([0-9]+)[[:space:]]B/op[[:space:]]+([0-9]+)[[:space:]]allocs/op ]]; then
			name="${BASH_REMATCH[1]}"
			iterations="${BASH_REMATCH[2]}"
			ns_per_op="${BASH_REMATCH[3]}"
			bytes_per_op="${BASH_REMATCH[4]}"
			allocs_per_op="${BASH_REMATCH[5]}"

			if (($(echo "$ns_per_op > 1000" | bc -l))); then
				time=$(echo "scale=2; $ns_per_op / 1000" | bc)
				time_unit="μs"
			else
				time="$ns_per_op"
				time_unit="ns"
			fi

			printf "%-40s %12s %8s %8s B/op %5s allocs/op\n" \
				"$name" "${iterations} ops" "${time}${time_unit}" \
				"$bytes_per_op" "$allocs_per_op" >>"$summary_file"
		fi
	done

	cat "$summary_file"
	echo ""
}

create_summary_table() {
	local summary_file="benchmark_results/full_summary.md"

	echo "# Benchmark Results" >"$summary_file"
	echo "" >>"$summary_file"
	echo "## Performance Comparison" >>"$summary_file"
	echo "" >>"$summary_file"
	echo "| Benchmark | Iterations | Time | Memory | Allocations |" >>"$summary_file"
	echo "|-----------|------------|------|--------|-------------|" >>"$summary_file"

	for log in benchmark_results/*_summary.log; do
		if [ -f "$log" ]; then
			tail -n +3 "$log" | while IFS= read -r line; do
				if [[ ! -z "$line" ]]; then
					benchmark_name=$(echo "$line" | grep -o '^Benchmark[^0-9]*')
					iterations=$(echo "$line" | grep -o '[0-9]\+[[:space:]]*[kMG]?ops' | head -1)
					time=$(echo "$line" | grep -o '[0-9.]\+[[:space:]]*\(ns\|μs\|ms\)' | head -1)
					memory=$(echo "$line" | grep -o '[0-9]\+[[:space:]]*B/op' | head -1)
					if [[ -z "$memory" ]]; then
						memory="0 B/op"
					fi
					allocations=$(echo "$line" | grep -o '[0-9]\+[[:space:]]*allocs/op' | head -1)
					if [[ -z "$allocations" ]]; then
						allocations="0 allocs/op"
					fi
					echo "| $benchmark_name | $iterations | $time | $memory | $allocations |" >>"$summary_file"
				fi
			done
		fi
	done
}

main() {
	echo -e "${GREEN}Starting benchmark suite...${NC}"
	echo ""

	run_benchmark_group "Fibonacci" "Fibonacci" "$YELLOW"
	run_benchmark_group "Power_Iterative" "IteractivePow" "$BLUE"
	run_benchmark_group "Power_Binary" "BinaryExpansionPow" "$PURPLE"
	run_benchmark_group "Power_Math" "MathPow" "$CYAN"
	run_benchmark_group "Large_Exponent" "LargeExponent" "$RED"

	create_summary_table

	echo -e "${GREEN}========================================${NC}"
	echo -e "${GREEN}=== All benchmarks completed! ===${NC}"
	echo -e "${GREEN}========================================${NC}"
	echo ""
	echo -e "${CYAN}Results saved in:${NC}"
	ls -la benchmark_results/
	echo ""
	echo -e "${CYAN}Full summary: benchmark_results/full_summary.md${NC}"

	echo ""
	echo -e "${PURPLE}=== Performance Summary ===${NC}"
	cat benchmark_results/full_summary.md
}
main
