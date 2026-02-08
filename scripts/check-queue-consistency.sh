#!/usr/bin/env bash
set -euo pipefail

# check-queue-consistency.sh
# Compares queue file declared statuses against actual tasks.md completion states.
# Exit 0 = all consistent, Exit 1 = inconsistencies found.

ERRORS=0
CHECKED=0

trim() {
  local s="$1"
  s="${s#"${s%%[![:space:]]*}"}"
  s="${s%"${s##*[![:space:]]}"}"
  s="${s//\`/}"
  printf '%s' "$s"
}

count_status() {
  local file="$1" pattern="$2"
  local n
  n=$(grep -c "$pattern" "$file" 2>/dev/null || true)
  printf '%s' "${n:-0}" | tr -d '[:space:]'
}

# Parse a markdown table row into an array of cell values.
# Usage: parse_row "| a | b | c |" cells
parse_row() {
  local line="$1"
  local -n _arr=$2
  _arr=()
  # Remove leading/trailing pipe and split
  line="${line#|}"
  line="${line%|}"
  while IFS='|' read -r -d '|' cell; do
    _arr+=("$(trim "$cell")")
  done <<< "${line}|"
}

check_queue_file() {
  local queue_file="$1"
  local label="$2"

  if [ ! -f "$queue_file" ]; then
    echo "[WARN] $queue_file not found, skipping"
    return
  fi

  local header_found=0
  local col_qid=-1 col_tasks=-1 col_status=-1

  while IFS= read -r line; do
    # Skip non-table lines
    [[ "$line" != "|"* ]] && continue
    # Skip separator lines
    [[ "$line" =~ ^\|[[:space:]]*-+[[:space:]]*\| ]] && continue

    local cells=()
    parse_row "$line" cells
    local ncols=${#cells[@]}

    # Detect header row — only for the first valid table
    if [ "$header_found" -eq 0 ]; then
      for i in $(seq 0 $((ncols - 1))); do
        local val="${cells[$i]}"
        case "$val" in
          "Queue ID") col_qid=$i ;;
          "Tasks File") col_tasks=$i ;;
          "Spec")
            # In issue-priority-queue, "Spec" column holds the tasks path
            if [ "$col_tasks" -eq -1 ]; then
              col_tasks=$i
            fi
            ;;
          "Status") col_status=$i ;;
        esac
      done
      if [ "$col_qid" -ge 0 ] && [ "$col_tasks" -ge 0 ] && [ "$col_status" -ge 0 ]; then
        header_found=1
      else
        # Reset for next table
        col_qid=-1; col_tasks=-1; col_status=-1
      fi
      continue
    fi

    # Once header is found, if we hit a row that doesn't start with QUEUE-/PHASE-
    # in the qid column, we've left the Queue Items table — stop processing.
    [ "$ncols" -le "$col_qid" ] && continue
    [ "$ncols" -le "$col_tasks" ] && continue
    [ "$ncols" -le "$col_status" ] && continue
    local qid="${cells[$col_qid]}"
    [[ -z "$qid" ]] && continue
    # Only process QUEUE-* or PHASE-* rows
    [[ "$qid" != QUEUE-* ]] && [[ "$qid" != PHASE-* ]] && continue

    local tasks_path="${cells[$col_tasks]}"
    local status="${cells[$col_status]}"

    # Skip rows where tasks_path doesn't look like a file path
    [[ "$tasks_path" != *".md"* ]] && continue

    if [ ! -f "$tasks_path" ]; then
      echo "[SKIP] $label/$qid: $tasks_path not found"
      continue
    fi

    local total completed pending in_prog
    total=$(count_status "$tasks_path" '^\*\*Status:\*\*')
    completed=$(count_status "$tasks_path" '^\*\*Status:\*\* completed')
    pending=$(count_status "$tasks_path" '^\*\*Status:\*\* pending')
    in_prog=$(count_status "$tasks_path" '^\*\*Status:\*\* in_progress')

    CHECKED=$((CHECKED + 1))

    if [ "$status" = "completed" ]; then
      if [ "$total" -gt 0 ] && [ "$completed" -ne "$total" ]; then
        echo "[ERROR] $label/$qid: marked 'completed' but tasks=$completed/$total (pending=$pending, in_progress=$in_prog)"
        ERRORS=$((ERRORS + 1))
      else
        echo "[OK]    $label/$qid: completed ($total/$total)"
      fi
    elif [ "$status" = "pending" ] || [ "$status" = "in_progress" ]; then
      if [ "$total" -gt 0 ] && [ "$completed" -eq "$total" ]; then
        echo "[ERROR] $label/$qid: marked '$status' but all $total tasks completed"
        ERRORS=$((ERRORS + 1))
      else
        echo "[OK]    $label/$qid: $status ($completed/$total completed)"
      fi
    else
      echo "[OK]    $label/$qid: status='$status' ($completed/$total completed)"
    fi

  done < "$queue_file"
}

echo "=== Queue Consistency Check ==="
echo ""

echo "--- task-queue.md ---"
check_queue_file ".specs/task-queue.md" "task-queue"
echo ""

echo "--- issue-priority-queue.md ---"
check_queue_file ".specs/issue-priority-queue.md" "issue-priority"
echo ""

echo "--- v1-release-queue.md ---"
check_queue_file ".specs/v1-release-queue.md" "v1-release"
echo ""

echo "Checked: $CHECKED queue items"
if [ "$ERRORS" -gt 0 ]; then
  echo "Result: FAILED ($ERRORS inconsistencies found)"
  exit 1
fi
echo "Result: All queues consistent"
exit 0
