package constant

type PredefinedStatus struct {
	ID      uint32
	Title   string
	HexCode string
}

var PredefinedStatusMap map[uint32]PredefinedStatus = map[uint32]PredefinedStatus{
	1: {1, "Open", "3D4AFF"},
	2: {2, "Waiting", "FF913D"},
	3: {3, "In Progress", "913DFF"},
	4: {4, "Declined", "FF3D4A"},
	5: {5, "Reopened", "FF3D4A"},
	6: {6, "Stuck", "F23DFF"},
	7: {7, "Idle", "FF3DAB"},
	8: {8, "Waiting Dependency", "FF913D"},
}
