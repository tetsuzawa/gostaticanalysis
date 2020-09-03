package milestonequeue

import "go/types"

type MilestoneQueue []types.Object

func (q *MilestoneQueue) push(t types.Object) {
	*q = append(*q, t)
}

func (q *MilestoneQueue) pop() (t types.Object) {
	if q == nil {
		return nil
	}
	t = (*q)[0]
	*q = (*q)[1:]
	return t
}
