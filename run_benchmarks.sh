#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

mkdir -p benchmark_results

get_benchmark_groups() {
	local test_files=$(find ./internal/basic-module -name "*_test.go" -type f)
	local all_benchmarks=""

	for file in $test_files; do
		if [ -f "$file" ]; then
			benchmarks=$(grep -o 'func Benchmark[^(]*' "$file" | sed 's/func Benchmark//' | sort -u)
			all_benchmarks="$all_benchmarks $benchmarks"
		fi
	done

	echo "$all_benchmarks" | tr ' ' '\n' | grep -v '^$' | while read benchmark; do
		group=$(echo "$benchmark" | sed -E 's/^([A-Za-z_]+).*$/\1/')
		echo "$group"
	done | sort -u
}

run_dynamic_benchmark_groups() {
	local groups=$(get_benchmark_groups)
	local color_index=0
	local colors=("$YELLOW" "$BLUE" "$PURPLE" "$CYAN" "$RED" "$GREEN" "$NC")

	for group in $groups; do
		if [ ! -z "$group" ]; then
			local color=${colors[$((color_index % ${#colors[@]}))]}
			run_benchmark_group "$group" "$group" "$color"
			((color_index++))
		fi
	done
}

run_benchmark_group() {
	local name=$1
	local pattern=$2
	local color=$3
	local output_file="benchmark_results/${name}_bench.log"
	local summary_file="benchmark_results/${name}_summary.log"

	echo -e "${color}========================================${NC}"
	echo -e "${color}=== Running $name Benchmarks ===${NC}"
	echo -e "${color}========================================${NC}"

	go test -bench="Benchmark$pattern" -benchmem ./internal/basic-module/... -v >"$output_file" 2>&1

	if [ -s "$output_file" ] && grep -q "Benchmark" "$output_file"; then
		echo -e "${CYAN}$name Benchmark Results:${NC}" >"$summary_file"
		echo "----------------------------------------" >>"$summary_file"

		grep "^Benchmark" "$output_file" | while read -r line; do
			if [[ $line =~ ^(Benchmark[^[:space:]]+)[[:space:]]+([0-9]+)[[:space:]]+([0-9.]+)[[:space:]]ns/op[[:space:]]+([0-9]+)[[:space:]]B/op[[:space:]]+([0-9]+)[[:space:]]allocs/op ]]; then
				name="${BASH_REMATCH[1]}"
				iterations="${BASH_REMATCH[2]}"
				ns_per_op="${BASH_REMATCH[3]}"
				bytes_per_op="${BASH_REMATCH[4]}"
				allocs_per_op="${BASH_REMATCH[5]}"

				if (($(echo "$ns_per_op > 1000000" | bc -l 2>/dev/null || echo "0"))); then
					time=$(echo "scale=2; $ns_per_op / 1000000" | bc 2>/dev/null || echo "0")
					time_unit="ms"
				elif (($(echo "$ns_per_op > 1000" | bc -l 2>/dev/null || echo "0"))); then
					time=$(echo "scale=2; $ns_per_op / 1000" | bc 2>/dev/null || echo "0")
					time_unit="μs"
				else
					time="$ns_per_op"
					time_unit="ns"
				fi

				printf "%-40s %12s %8s %8s B/op %5s allocs/op\n" \
					"$name" "${iterations} ops" "${time}${time_unit}" \
					"$bytes_per_op" "$allocs_per_op" >>"$summary_file"
			elif [[ $line =~ ^(Benchmark[^[:space:]]+)[[:space:]]+([0-9]+)[[:space:]]+([0-9.]+)[[:space:]]ns/op ]]; then
				name="${BASH_REMATCH[1]}"
				iterations="${BASH_REMATCH[2]}"
				ns_per_op="${BASH_REMATCH[3]}"

				if (($(echo "$ns_per_op > 1000000" | bc -l 2>/dev/null || echo "0"))); then
					time=$(echo "scale=2; $ns_per_op / 1000000" | bc 2>/dev/null || echo "0")
					time_unit="ms"
				elif (($(echo "$ns_per_op > 1000" | bc -l 2>/dev/null || echo "0"))); then
					time=$(echo "scale=2; $ns_per_op / 1000" | bc 2>/dev/null || echo "0")
					time_unit="μs"
				else
					time="$ns_per_op"
					time_unit="ns"
				fi

				printf "%-40s %12s %8s\n" \
					"$name" "${iterations} ops" "${time}${time_unit}" >>"$summary_file"
			fi
		done

		cat "$summary_file"
	else
		echo -e "${RED}No benchmarks found for pattern: $pattern${NC}"
		rm -f "$output_file"
	fi
	echo ""
}

create_summary_table() {
	local summary_file="benchmark_results/full_summary.md"

	echo "# Benchmark Results" >"$summary_file"
	echo "" >>"$summary_file"
	echo "## Performance Comparison" >>"$summary_file"
	echo "" >>"$summary_file"
	echo "| Benchmark | Time | Memory | Allocations |" >>"$summary_file"
	echo "|-----------|------|--------|-------------|" >>"$summary_file"

	for log in benchmark_results/*_summary.log; do
		if [ -f "$log" ]; then
			tail -n +3 "$log" | while IFS= read -r line; do
				if [[ ! -z "$line" ]]; then
					benchmark_name=$(echo "$line" | grep -o '^Benchmark[^0-9]*')
					time=$(echo "$line" | grep -o '[0-9.]\+[[:space:]]*\(ns\|μs\|ms\)' | head -1)
					memory=$(echo "$line" | grep -o '[0-9]\+[[:space:]]*B/op' | head -1)
					if [[ -z "$memory" ]]; then
						memory="0 B/op"
					fi
					allocations=$(echo "$line" | grep -o '[0-9]\+[[:space:]]*allocs/op' | head -1)
					if [[ -z "$allocations" ]]; then
						allocations="0 allocs/op"
					fi
					echo "| $benchmark_name |  $time | $memory | $allocations |" >>"$summary_file"
				fi
			done
		fi
	done
}

clean_old_results() {
	echo -e "${YELLOW}Cleaning old benchmark results...${NC}"
	rm -f benchmark_results/*.log benchmark_results/*.md
}

compare_with_previous() {
	if [ -f "benchmark_results/previous_full_summary.md" ]; then
		echo -e "${CYAN}=== Comparison with previous run ===${NC}"
		diff -u benchmark_results/previous_full_summary.md benchmark_results/full_summary.md | grep -E "^\+|^-" || echo "No significant changes"
	fi
	cp benchmark_results/full_summary.md benchmark_results/previous_full_summary.md
}

main() {
	clean_old_results

	echo -e "${GREEN}Starting benchmark suite...${NC}"
	echo ""

	run_dynamic_benchmark_groups

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
	if [ -f "benchmark_results/full_summary.md" ]; then
		cat benchmark_results/full_summary.md
	else
		echo -e "${RED}No summary file generated${NC}"
	fi
}

main
