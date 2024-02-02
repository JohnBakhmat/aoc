open Stdio

type round_result = Win | Loose | Tie

let round_resolution a b =
  match (a, b) with
  | "A", "Y" | "B", "Z" | "C", "X" -> Win
  | "A", "Z" | "B", "X" | "C", "Y" -> Loose
  | _ -> Tie

let move_to_score m = match m with "X" -> 1 | "Y" -> 2 | "Z" -> 3 | _ -> 0
let result_to_score r = match r with Win -> 6 | Tie -> 3 | Loose -> 0

let rec sum l =
  match l with
  | [] -> 0
  | x :: tail -> ( match x with Some n -> n + sum tail | None -> 0)

let () =
  let results =
    In_channel.read_all "../input.1.txt"
    |> String.split_on_char '\n'
    |> List.filter (fun s -> s <> "")
    |> List.map (fun s ->
           let moves = String.split_on_char ' ' s in
           match moves with
           | a :: b :: _ ->
               Some (result_to_score (round_resolution a b) + move_to_score b)
           | _ -> None)
  in

  print_int (sum results)
