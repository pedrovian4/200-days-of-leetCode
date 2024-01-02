<?php

    function groupAnagrams($strs) {
        $map = [];
        foreach($strs as $str){
            $key = str_split($str);
            sort($key);
            $key = implode("",$key);
            $map[$key][]=$str;
        }
        return $map;
    }
