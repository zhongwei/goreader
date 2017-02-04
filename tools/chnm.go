package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "sort"
    "strconv"
    "strings"
)

const (
    separator = string(os.PathSeparator)
)

func init() {
    flag.Parse()
}

func main() {
    top_path := flag.Arg(0)
    dirs := sorted_file_names(top_path)

    for _, current_path := range dirs {
        process_path := top_path + separator + current_path
        sortednames := sorted_file_names(process_path)
        change_names(process_path, sortednames)
    }
}


func change_names(current_path string, sortednames []string) {
    for i, name := range sortednames {
        newname := get_new_name(i, name)
        ch_name(current_path + separator + name, current_path + separator + newname)
    }
}

func get_new_name(i int, name string) (newname string) {
        lastdot := strings.LastIndex(name, ".")
        if i < 9 {
           newname = "0" + strconv.Itoa(i + 1) + name[lastdot:]
        } else {
           newname = strconv.Itoa(i + 1) + name[lastdot:]
        }
        return newname
}

func sorted_file_names(dir string) (filenames []string) {
    files, err := ioutil.ReadDir(dir)

    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        filename := file.Name()
        if !is_cover(filename) {
            filenames = append(filenames, filename)
        }
    }

    sort.Strings(filenames)
    return filenames
}

func is_cover(filename string) bool {
    return strings.EqualFold(filename, "0.jpg") || strings.EqualFold(filename, "0.jpeg") || strings.EqualFold(filename, "0.png")
}

func ch_name(s, t string) {
    err := os.Rename(s, t)

    if err != nil {
        fmt.Println(err)
        return
    }
}

