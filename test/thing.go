package main

// +test slice:"Any,All,Count,Distinct,DistinctBy,Each,First,MaxBy,MinBy,Single,Where,SortBy,SortByDesc,IsSortedBy,IsSortedByDesc,Aggregate[Other],Average[Other],GroupBy[Other],Max[Other],Min[Other],Select[Other],Shuffle,Sum[Other]"
type Thing struct {
	Name   string
	Number Other
}

// +test * slice:"Any,All,Count,Distinct,DistinctBy,Each,First,MaxBy,MinBy,Single,Where,SortBy,SortByDesc,IsSortedBy,IsSortedByDesc,Aggregate[Other],Average[Other],GroupBy[Other],Max[Other],Min[Other],Select[Other],Shuffle,Sum[Other]"
type PointerThing struct {
	Name   string
	Number Other
}
