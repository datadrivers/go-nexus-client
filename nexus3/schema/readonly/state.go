package readonly

type State struct {
	SummaryReason   string `json:"summaryReason"`
	SystemInitiated bool   `json:"systemInitiated"`
	Frozen          bool   `json:"frozen"`
}
