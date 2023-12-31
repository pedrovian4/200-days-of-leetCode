<?php

/**
 * @param Integer[] $nums
 * @param Integer $k
 * @return Integer[]
 */
function topKFrequent($nums, $k) {
   if (count($nums) <= 1) {
        return $nums;
    }

    $numsCounted = [];
    foreach ($nums as $n) {
        if (!isset($numsCounted[$n])) {
            $numsCounted[$n] = 0;
        }
        $numsCounted[$n]++;
    }

    arsort($numsCounted, SORT_NUMERIC);

    $frequent = array_slice(array_keys($numsCounted), 0, $k);

    return $frequent;
}