open Base
open Stdio

let file = "input.1.txt"

let map_letters_to_number = function
  | "one" -> "1"
  | "two" -> "2"
  | "three" -> "3"
  | "four" -> "4"
  | "five" -> "5"
  | "six" -> "6"
  | "seven" -> "7"
  | "eight" -> "8"
  | "nine" -> "9"
  | x -> x


let solve_line line ~regex =
    let _ = Str.search_forward regex line 0 in
    let first = Str.matched_string line in
    let _ = Str.search_backward regex line (String.length line) in
    let last = Str.matched_string line in
    (first |> map_letters_to_number) ^ (last |> map_letters_to_number)
  |> Int.of_string

let sum = List.fold ~init:0 ~f:( + )

let solve lines ~regex =
    List.map lines ~f:(solve_line ~regex)
  |> sum


let () = 
    let lines = In_channel.read_lines file in
    let regex = Str.regexp "[0-9]" in
    let res = solve lines ~regex:regex in
    print_endline (Int.to_string res)

